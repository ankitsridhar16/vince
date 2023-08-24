package entry

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/apache/arrow/go/v14/arrow/compute"
	"github.com/apache/arrow/go/v14/arrow/memory"
	"github.com/apache/arrow/go/v14/parquet"
	"github.com/apache/arrow/go/v14/parquet/compress"
	"github.com/apache/arrow/go/v14/parquet/file"
	"github.com/apache/arrow/go/v14/parquet/schema"
	"github.com/cespare/xxhash/v2"
	"github.com/vinceanalytics/vince/internal/must"
	v1 "github.com/vinceanalytics/vince/proto/v1"
	"golang.org/x/exp/slices"
)

type Entry struct {
	Bounce         int64
	Session        int64
	Browser        string
	BrowserVersion string
	City           string
	Country        string
	Domain         string
	Duration       time.Duration
	EntryPage      string
	ExitPage       string
	Host           string
	ID             uint64
	Event          string
	Os             string
	OsVersion      string
	Path           string
	Referrer       string
	ReferrerSource string
	Region         string
	Screen         string
	Timestamp      int64
	UtmCampaign    string
	UtmContent     string
	UtmMedium      string
	UtmSource      string
	UtmTerm        string
}

type BufWrite interface {
	Append(any)
	Write(file.ColumnChunkWriter, *roaring64.Bitmap, v1.Column)
	Reset()
}

type ByteArray struct {
	hash xxhash.Digest
	buf  []parquet.ByteArray
}

var _ BufWrite = (*ByteArray)(nil)

func NewByteArray() *ByteArray {
	return &ByteArray{
		buf: make([]parquet.ByteArray, 0, 1<<10),
	}
}

func (b *ByteArray) Append(a any) {
	b.buf = append(b.buf, parquet.ByteArray(a.(string)))
}

func (b *ByteArray) Reset() {
	b.buf = b.buf[:0]
	b.hash.Reset()
}

func (b *ByteArray) Write(g file.ColumnChunkWriter, r *roaring64.Bitmap, field v1.Column) {
	w := g.(*file.ByteArrayColumnChunkWriter)
	must.Must(w.WriteBatch(b.buf, nil, nil))(
		"failed writing int64 column to parquet",
	)
	w.Close()
	if r == nil {
		return
	}
	column := []byte(field.String())
	for i := range b.buf {
		if len(b.buf[i]) == 0 {
			continue
		}
		b.hash.Reset()
		b.hash.Write(column)
		b.hash.Write(b.buf[i])
		sum := b.hash.Sum64()
		if !r.Contains(sum) {
			r.Add(sum)
		}
	}
}

type Int64Array struct {
	buf []int64
}

var _ BufWrite = (*Int64Array)(nil)

func NewInt64Array() *Int64Array {
	return &Int64Array{
		buf: make([]int64, 0, 1<<10),
	}
}

func (b *Int64Array) First() int64 {
	if len(b.buf) == 0 {
		return 0
	}
	return b.buf[0]
}

func (b *Int64Array) Last() int64 {
	if len(b.buf) == 0 {
		return 0
	}
	return b.buf[len(b.buf)-1]
}

func (b *Int64Array) Append(v any) {
	b.buf = append(b.buf, v.(int64))
}

func (b *Int64Array) Reset() {
	b.buf = b.buf[:0]
}

func (b *Int64Array) Write(g file.ColumnChunkWriter, _ *roaring64.Bitmap, _ v1.Column) {
	w := g.(*file.Int64ColumnChunkWriter)
	must.Must(w.WriteBatch(b.buf, nil, nil))(
		"failed writing int64 column to parquet",
	)
	w.Close()
}

type MultiEntry struct {
	mu      sync.Mutex
	columns [v1.Column_utm_term + 1]BufWrite
}

func (m *MultiEntry) Boundary() (lo, hi int64) {
	ts := m.columns[v1.Column_timestamp].(*Int64Array)
	lo = ts.First()
	hi = ts.Last()
	return
}

func (m *MultiEntry) setup() {
	for i := range m.columns {
		col := v1.Column(i)
		switch col {
		case v1.Column_bounce,
			v1.Column_duration,
			v1.Column_id,
			v1.Column_session,
			v1.Column_timestamp:
			m.columns[i] = NewInt64Array()
		default:
			m.columns[i] = NewByteArray()
		}
	}
}

var multiPool = &sync.Pool{
	New: func() any {
		var m MultiEntry
		m.setup()
		return &m
	},
}

func NewMulti() *MultiEntry {
	return multiPool.Get().(*MultiEntry)
}

func (m *MultiEntry) Release() {
	m.Reset()
	multiPool.Put(m)
}

func (m *MultiEntry) Reset() {
	for i := range m.columns {
		m.columns[i].Reset()
	}
}

func (m *MultiEntry) Append(e *Entry) {
	m.mu.Lock()
	m.columns[v1.Column_bounce].Append(e.Bounce)
	m.columns[v1.Column_browser].Append(e.Browser)
	m.columns[v1.Column_browser_version].Append(e.BrowserVersion)
	m.columns[v1.Column_city].Append(e.City)
	m.columns[v1.Column_country].Append(e.Country)
	m.columns[v1.Column_duration].Append(int64(e.Duration))
	m.columns[v1.Column_entry_page].Append(e.EntryPage)
	m.columns[v1.Column_event].Append(e.Event)
	m.columns[v1.Column_exit_page].Append(e.ExitPage)
	m.columns[v1.Column_host].Append(e.Host)
	m.columns[v1.Column_id].Append(int64(e.ID))
	m.columns[v1.Column_os].Append(e.Os)
	m.columns[v1.Column_os_version].Append(e.OsVersion)
	m.columns[v1.Column_path].Append(e.Path)
	m.columns[v1.Column_referrer].Append(e.Referrer)
	m.columns[v1.Column_referrer_source].Append(e.ReferrerSource)
	m.columns[v1.Column_region].Append(e.Region)
	m.columns[v1.Column_screen].Append(e.Screen)
	m.columns[v1.Column_session].Append(e.Session)
	m.columns[v1.Column_timestamp].Append(e.Timestamp)
	m.columns[v1.Column_utm_campaign].Append(e.UtmCampaign)
	m.columns[v1.Column_utm_content].Append(e.UtmContent)
	m.columns[v1.Column_utm_medium].Append(e.UtmMedium)
	m.columns[v1.Column_utm_source].Append(e.UtmSource)
	m.columns[v1.Column_utm_term].Append(e.UtmTerm)
	m.mu.Unlock()
}

func (m *MultiEntry) Write(f *file.Writer, r *roaring64.Bitmap) {
	g := f.AppendRowGroup()
	next := func() file.ColumnChunkWriter {
		return must.Must(g.NextColumn())("failed getting next column")
	}
	for i := range m.columns {
		m.columns[i].Write(next(), r, v1.Column(i))
	}
	must.One(g.Close())("failed closing row group writer")
}

// Fields for constructing arrow schema on Entry.
func Fields() []arrow.Field {
	return []arrow.Field{
		{Name: "bounce", Type: arrow.PrimitiveTypes.Int64},
		{Name: "browser", Type: arrow.BinaryTypes.String},
		{Name: "browser_version", Type: arrow.BinaryTypes.String},
		{Name: "city", Type: arrow.BinaryTypes.String},
		{Name: "country", Type: arrow.BinaryTypes.String},
		{Name: "duration", Type: arrow.PrimitiveTypes.Int64},
		{Name: "entry_page", Type: arrow.BinaryTypes.String},
		{Name: "event", Type: arrow.BinaryTypes.String},
		{Name: "exit_page", Type: arrow.BinaryTypes.String},
		{Name: "host", Type: arrow.BinaryTypes.String},
		{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		{Name: "os", Type: arrow.BinaryTypes.String},
		{Name: "os_version", Type: arrow.BinaryTypes.String},
		{Name: "path", Type: arrow.BinaryTypes.String},
		{Name: "referrer", Type: arrow.BinaryTypes.String},
		{Name: "referrer_source", Type: arrow.BinaryTypes.String},
		{Name: "region", Type: arrow.BinaryTypes.String},
		{Name: "screen", Type: arrow.BinaryTypes.String},
		{Name: "session", Type: arrow.PrimitiveTypes.Int64},
		{Name: "timestamp", Type: arrow.PrimitiveTypes.Int64},
		{Name: "utm_campaign", Type: arrow.BinaryTypes.String},
		{Name: "utm_content", Type: arrow.BinaryTypes.String},
		{Name: "utm_medium", Type: arrow.BinaryTypes.String},
		{Name: "utm_source", Type: arrow.BinaryTypes.String},
		{Name: "utm_term", Type: arrow.BinaryTypes.String},
	}
}

var All = Fields()

var Index = func() (m map[string]int) {
	m = make(map[string]int)
	for i := range All {
		m[All[i].Name] = i
	}
	return
}()

var ParquetSchema = parquetSchema()

func parquetSchema() *schema.Schema {
	f := Fields()
	nodes := make(schema.FieldList, len(f))
	for i := range nodes {
		x := &f[i]
		var logicalType schema.LogicalType
		var typ parquet.Type
		switch f[i].Type.ID() {
		case arrow.STRING:
			logicalType = schema.StringLogicalType{}
			typ = parquet.Types.ByteArray

		default:
			typ = parquet.Types.Int64
			logicalType = schema.NewIntLogicalType(64, true)
		}
		nodes[i] = must.Must(
			schema.NewPrimitiveNodeLogical(x.Name,
				parquet.Repetitions.Required,
				logicalType,
				typ, -1, -1),
		)("schema.NewPrimitiveNodeLogical")
	}
	root := must.Must(
		schema.NewGroupNode(parquet.DefaultRootName,
			parquet.Repetitions.Required, nodes, -1),
	)("schema.NewGroupNode")
	return schema.NewSchema(root)
}

func NewFileWriter(w io.Writer) *file.Writer {
	return file.NewParquetWriter(w,
		ParquetSchema.Root(),
		file.WithWriterProps(
			parquet.NewWriterProperties(
				parquet.WithAllocator(Pool),
				parquet.WithCompression(compress.Codecs.Zstd),
				parquet.WithCompressionLevel(10),
			),
		),
	)
}

func NewFileReader(r parquet.ReaderAtSeeker) *file.Reader {
	return must.Must(
		file.NewParquetReader(r),
	)("failed creating new parquet file reader")
}

var Schema = arrow.NewSchema(All, nil)

var Pool = memory.NewGoAllocator()

var entryPool = &sync.Pool{
	New: func() any {
		return new(Entry)
	},
}

func NewEntry() *Entry {
	return entryPool.Get().(*Entry)
}

func (e *Entry) Clone() *Entry {
	o := NewEntry()
	*o = *e
	return o
}

func (e *Entry) Release() {
	*e = Entry{}
	entryPool.Put(e)
}

func (e *Entry) Hit() {
	e.EntryPage = e.Path
	e.Bounce = 1
	e.Session = 1
}

func (s *Entry) Update(e *Entry) {
	if s.Bounce == 1 {
		s.Bounce, e.Bounce = -1, -1
	}
	e.ExitPage = e.Path
	e.Session = 0
	e.Duration = time.UnixMilli(e.Timestamp).Sub(time.UnixMilli(s.Timestamp))
	s.Timestamp = e.Timestamp
}

func Context(ctx ...context.Context) context.Context {
	if len(ctx) > 0 {
		return compute.WithAllocator(ctx[0], Pool)
	}
	return compute.WithAllocator(context.Background(), Pool)
}

type Reader struct {
	strings []parquet.ByteArray
	ints    []int64
	b       *array.RecordBuilder
}

func NewReader() *Reader {
	return readerPool.Get().(*Reader)
}

var readerPool = &sync.Pool{
	New: func() any {
		return &Reader{
			strings: make([]parquet.ByteArray, 1<<10),
			ints:    make([]int64, 1<<10),
			b:       array.NewRecordBuilder(Pool, Schema),
		}
	},
}

func (b *Reader) Release() {
	b.strings = b.strings[:0]
	b.ints = b.ints[:0]
	readerPool.Put(b)
}

func (b *Reader) Read(r *file.Reader, groups []int) {
	x := b.b.Schema()
	if groups == nil {
		groups = make([]int, r.NumRowGroups())
		for i := range groups {
			groups[i] = i
		}
	}
	for i := range groups {
		g := r.RowGroup(groups[i])
		n := g.NumRows()
		for f := 0; f < x.NumFields(); f++ {
			b.read(f, n, must.Must(g.Column(f))("failed getting a RowGroup column"))
		}
	}
}

func (b *Reader) read(f int, rows int64, chunk file.ColumnChunkReader) {
	switch e := b.b.Field(f).(type) {
	case *array.StringBuilder:
		r := chunk.(*file.ByteArrayColumnChunkReader)
		b.strings = slices.Grow(b.strings, int(rows))[:rows]
		r.ReadBatch(rows, b.strings, nil, nil)
		e.Reserve(int(rows))
		for i := range b.strings {
			e.UnsafeAppend(b.strings[i])
		}
	case *array.Int64Builder:
		r := chunk.(*file.Int64ColumnChunkReader)
		b.ints = slices.Grow(b.ints, int(rows))[:rows]
		r.ReadBatch(rows, b.ints, nil, nil)
		e.AppendValues(b.ints, nil)
	default:
		must.AssertFMT(false)("unsupported arrow builder type %T", e)
	}
}

func (b *Reader) Record() arrow.Record {
	return b.b.NewRecord()
}
