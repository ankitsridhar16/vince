package api

import (
	"crypto/ed25519"
	"errors"
	"net/http"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vinceanalytics/vince/internal/config"
	"github.com/vinceanalytics/vince/internal/core"
	"github.com/vinceanalytics/vince/internal/db"
	"github.com/vinceanalytics/vince/internal/keys"
	"github.com/vinceanalytics/vince/internal/must"
	"github.com/vinceanalytics/vince/internal/pj"
	"github.com/vinceanalytics/vince/internal/render"
	"github.com/vinceanalytics/vince/internal/secrets"
	"github.com/vinceanalytics/vince/internal/tokens"
	v1 "github.com/vinceanalytics/vince/proto/v1"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slog"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

var privateKey = secrets.ED25519Raw()
var publicKey = privateKey.Public()

func Token(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var tr v1.Token_Request
	err := pj.UnmarshalDefault(&tr, r.Body)
	if err != nil ||
		((tr.Name == "" ||
			tr.Password == "") ||
			(!tr.Generate &&
				tr.PublicKey == nil ||
				tr.Token == "")) {
		render.ERROR(w, http.StatusBadRequest)
		return
	}

	if tr.Ttl == nil {
		tr.Ttl = durationpb.New(30 * 24 * time.Hour)
	}
	var claims *jwt.RegisteredClaims
	pub := publicKey
	if tr.Generate {
		// By default we generate short lived tokens.
		tr.Token, claims = tokens.Generate(ctx, privateKey, v1.Token_SERVER,
			tr.Name, core.Now(ctx).Add(tr.Ttl.AsDuration()))
	} else {
		if len(tr.PublicKey) != ed25519.PublicKeySize {
			render.ERROR(w, http.StatusBadRequest,
				"invalid public key")
			return
		}
		pub = ed25519.PublicKey(tr.PublicKey)
		claims = &jwt.RegisteredClaims{}
		tok, err := jwt.ParseWithClaims(tr.Token, claims, func(t *jwt.Token) (interface{}, error) {
			return pub, nil
		})
		if err != nil || !tok.Valid {
			if err != nil {
				slog.Error("invalid token", "err", err)
			}
			render.ERROR(w, http.StatusBadRequest, "invalid token")
			return
		}
	}

	var a v1.Account
	err = db.Get(ctx).View(func(txn *badger.Txn) error {
		key := keys.Account(tr.Name)
		it, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return it.Value(func(val []byte) error {
			return proto.Unmarshal(val, &a)
		})
	})

	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			render.ERROR(w, http.StatusBadRequest, "account does not exist")
			return
		}
		slog.Error("failed reading account", "err", err)
		render.ERROR(w, http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword(a.Password, []byte(tr.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			render.ERROR(w, http.StatusBadRequest, "invalid password")
			return
		}
		slog.Error("failed comparing passwords", "err", err)
		render.ERROR(w, http.StatusInternalServerError)
		return
	}

	err = db.Get(ctx).Update(func(txn *badger.Txn) error {
		key := keys.Token(tr.Name, claims.ID)
		tok := must.Must(
			proto.Marshal(&v1.Token{
				PubKey: pub.(ed25519.PublicKey),
			}),
		)("failed encoding token")
		e := badger.NewEntry([]byte(key), tok)
		issuer := v1.Token_Issuer_value[claims.Issuer]
		e.WithMeta(byte(issuer)).
			WithTTL(tr.Ttl.AsDuration())
		return txn.SetEntry(e)
	})
	if err != nil {
		slog.Error("failed saving token", "err", err)
		render.ERROR(w, http.StatusInternalServerError)
		return
	}
	o := config.Get(ctx)
	render.JSON(w, http.StatusOK, &v1.Client_Auth{
		Name:  tr.Name,
		Token: tr.Token,
		Api:   o.ListenAddress,
		Mysql: o.MysqlListenAddress,
	})
}