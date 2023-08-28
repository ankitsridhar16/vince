package config

import (
	"context"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
	"github.com/vinceanalytics/vince/internal/must"
	"github.com/vinceanalytics/vince/internal/pj"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

type configKey struct{}

func Get(ctx context.Context) *Options {
	return ctx.Value(configKey{}).(*Options)
}

func Load(base *Options, x *cli.Context) (context.Context, error) {
	root, err := filepath.Abs(x.Args().First())
	if err != nil {
		return nil, err
	}
	base.SyncInterval = durationpb.New(x.Duration("sync-interval"))
	file := filepath.Join(root, FILE)
	b := must.Must(os.ReadFile(file))(
		"called vince on non vince project, call vince init and try again",
	)
	var f Options
	must.One(pj.Unmarshal(b, &f))("invalid configuration file")
	proto.Merge(base, &f)

	base.DbPath = resolve(root, base.DbPath)
	base.BlocksPath = resolve(root, base.BlocksPath)
	base.RaftPath = resolve(root, base.RaftPath)
	baseCtx := context.WithValue(context.Background(), configKey{}, base)
	return baseCtx, nil
}

func resolve(root, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(root, filepath.Clean(path))
}
