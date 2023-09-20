package engine

import (
	"fmt"
	"io"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/parquet-go/parquet-go"
	blocksv1 "github.com/vinceanalytics/vince/gen/proto/go/vince/blocks/v1"
	storev1 "github.com/vinceanalytics/vince/gen/proto/go/vince/store/v1"
	"github.com/vinceanalytics/vince/internal/db"
	"github.com/vinceanalytics/vince/internal/keys"
	"github.com/vinceanalytics/vince/internal/px"
)

type Table struct {
	Context
	name        string
	schema      tableSchema
	projections []string
}

var _ sql.Table = (*Table)(nil)
var _ sql.ProjectedTable = (*Table)(nil)

func (t *Table) Name() string {
	return t.name
}

func (t *Table) String() string {
	return t.name
}

func (t *Table) Schema() sql.Schema {
	return t.schema.sql
}

func (t *Table) Collation() sql.CollationID {
	return sql.Collation_Default
}

func (t *Table) Partitions(ctx *sql.Context) (sql.PartitionIter, error) {
	hints := GetIndexHint(ctx)
	txn := t.DB.NewTransaction(false)
	it := txn.Iter(db.IterOpts{
		Prefix:         keys.BlockMetadata(t.name, ""),
		PrefetchValues: false,
	})
	it.Rewind()
	return &partitionIter{
		it:  it,
		txn: txn,
		partition: Partition{
			Hints: hints,
		},
	}, nil
}

func (t *Table) PartitionRows(ctx *sql.Context, partition sql.Partition) (sql.RowIter, error) {
	x := partition.(*Partition)
	var record arrow.Record
	err := t.Reader.Read(ctx, x.Key(), func(f io.ReaderAt, size int64) error {
		r, err := parquet.OpenFile(f, size)
		if err != nil {
			return err
		}
		record, err = t.schema.read(ctx, r)
		return err
	})
	if err != nil {
		return nil, err
	}
	return newRecordIter(record), nil
}

func (t *Table) WithProjections(colNames []string) sql.Table {
	m := make([]storev1.Column, len(colNames))
	for i := range colNames {
		m[i] = storev1.Column(storev1.Column_value[colNames[i]])
	}
	return &Table{Context: t.Context,
		name:        t.name,
		schema:      createSchema(t.name, m),
		projections: colNames,
	}
}

func (t *Table) Projections() (o []string) {
	return t.projections
}

type Partition struct {
	Info  blocksv1.BlockInfo
	Hints *IndexHint
}

func (p *Partition) Key() []byte { return []byte(p.Info.Id) }

type partitionIter struct {
	it        db.Iter
	txn       db.Txn
	partition Partition
	started   bool
}

func (p *partitionIter) Next(*sql.Context) (sql.Partition, error) {
	if !p.started {
		p.started = true
		p.it.Rewind()
	} else {
		p.it.Next()
	}
	if !p.it.Valid() {
		return nil, io.EOF
	}
	err := p.it.Value(px.Decode(&p.partition.Info))
	if err != nil {
		return nil, fmt.Errorf("failed decoding block index err:%v", err)
	}
	return &p.partition, nil
}

func (p *partitionIter) Close(*sql.Context) error {
	p.it.Close()
	p.txn.Close()
	return nil
}
