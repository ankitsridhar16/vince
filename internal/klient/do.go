package klient

import (
	"bytes"
	"context"
	"net/http"

	"github.com/vinceanalytics/vince/internal/must"
	"github.com/vinceanalytics/vince/internal/pj"
	v1 "github.com/vinceanalytics/vince/proto/v1"
	"google.golang.org/protobuf/proto"
)

var client = &http.Client{}

type Input interface {
	*v1.Token_Request
}

type Output interface {
	*v1.Client_Auth
}

func POST[I Input, O Output](ctx context.Context, uri string, in I, out O) *v1.Error {
	return Do(ctx, http.MethodPost, uri, in, out)
}

func Do[I Input, O Output](ctx context.Context, method, uri string, in I, out O) *v1.Error {
	data := must.Must(pj.Marshal(any(in).(proto.Message)))(
		"failed encoding api request object",
	)
	r := must.Must(http.NewRequestWithContext(ctx, method, uri, bytes.NewReader(data)))(
		"failed creating api request",
	)
	r.Header.Set("Accept", "application/json")
	r.Header.Set("content-type", "application/json")
	res := must.Must(client.Do(r))(
		"failed sending api request", "uri", uri,
	)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var out v1.Error
		must.One(pj.UnmarshalDefault(&out, res.Body))(
			"failed decoding api error",
		)
		return &out
	}
	must.One(pj.UnmarshalDefault(any(out).(proto.Message), res.Body))(
		"failed decoding api result",
	)
	return nil
}
