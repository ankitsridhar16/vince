package timeseries

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/gernest/vince/caches"
	"github.com/gernest/vince/cities"
	"github.com/gernest/vince/log"
	"github.com/gernest/vince/store"
	"github.com/gernest/vince/system"
	"github.com/gernest/vince/ua"
)

type aggregate struct {
	u, s *roaring64.Bitmap
	sum  store.Sum
}

var groupPool = &sync.Pool{
	New: func() any {
		return &aggregate{
			u: roaring64.New(),
			s: roaring64.New(),
		}
	},
}

func (g *aggregate) Reset() {
	g.u.Clear()
	g.s.Clear()
	g.sum = store.Sum{}
}

func (g *aggregate) Save(el EntryList) {
	g.Reset()
	el.Count(g.u, g.s, &g.sum)
}

func (g *aggregate) Prop(el EntryList, by PROPS, f func(key string, sum *store.Sum) error) error {
	g.Reset()
	return el.EmitProp(g.u, g.s, by, &g.sum, f)
}

func (g *aggregate) City(el EntryList, f func(key uint32, sum *store.Sum) error) error {
	g.Reset()
	return el.EmitCity(g.u, g.s, &g.sum, f)
}

func (g *aggregate) Release() {
	g.Reset()
	groupPool.Put(g)
}

func Save(ctx context.Context, b *Buffer) {
	start := time.Now()
	defer system.SaveDuration.UpdateDuration(start)

	db := GetMike(ctx)
	defer b.Release()
	group := groupPool.Get().(*aggregate)
	ent := EntryList(b.entries)
	id := newID()
	defer id.Release()
	meta := newMetaKey()
	defer meta.Release()

	sid := b.SID()
	uid := b.UID()

	id.SetSiteID(sid)
	id.SetUserID(uid)
	meta.SetSiteID(sid)
	meta.SetUserID(uid)

	// Guarantee that aggregates are on per hour windows.
	ent.Emit(func(el EntryList) {
		defer group.Reset()
		group.Save(el)
		ts := time.Unix(el[0].Timestamp, 0)
		err := db.Update(func(txn *badger.Txn) error {
			id.Year(ts).SetTable(byte(TABLE_YEAR))
			return errors.Join(
				updateCalendar(ctx, txn, ts, id[:], &group.sum),
				updateMeta(ctx, txn, el, group, meta, ts),
			)
		})
		if err != nil {
			log.Get(ctx).Err(err).
				Uint64("uid", uid).
				Uint64("sid", sid).
				Msg("failed to save hourly stats ")
		}
	})
}

func updateMeta(ctx context.Context, txn *badger.Txn, el EntryList, g *aggregate, x *MetaKey, ts time.Time) error {
	errs := make([]error, 0, PROPS_CITY)
	for i := 1; i <= int(PROPS_CITY); i++ {
		p := PROPS(i)
		errs = append(errs, p.Save(ctx, txn, el, g, x, ts))
	}
	return errors.Join(errs...)
}

func (p PROPS) Save(ctx context.Context, txn *badger.Txn, el EntryList, g *aggregate, x *MetaKey, ts time.Time) error {
	switch p {
	case PROPS_NAME, PROPS_PAGE, PROPS_ENTRY_PAGE, PROPS_EXIT_PAGE,
		PROPS_REFERRER, PROPS_UTM_MEDIUM, PROPS_UTM_SOURCE,
		PROPS_UTM_CAMPAIGN, PROPS_UTM_CONTENT, PROPS_UTM_TERM:
		return g.Prop(el, p, func(key string, sum *store.Sum) error {
			return updateCalendar(ctx, txn, ts, x.SetMeta(byte(p)).String(key), sum)
		})
		// properties from user agent
	case PROPS_UTM_DEVICE, PROPS_UTM_BROWSER, PROPS_BROWSER_VERSION, PROPS_OS, PROPS_OS_VERSION:
		return g.Prop(el, p, func(key string, sum *store.Sum) error {
			return updateCalendar(ctx, txn, ts, x.SetMeta(byte(p)).HashU16(ua.ToIndex(key)), sum)
		})
	case PROPS_COUNTRY, PROPS_REGION:
		return g.Prop(el, p, func(key string, sum *store.Sum) error {
			return updateCalendar(ctx, txn, ts, x.SetMeta(byte(p)).HashU16(cities.ToIndex(key)), sum)
		})
	case PROPS_CITY:
		return g.City(el, func(key uint32, sum *store.Sum) error {
			return updateCalendar(ctx, txn, ts, x.SetMeta(byte(p)).HashU32(key), sum)
		})
	default:
		return nil
	}
}
func (p PROPS) StringKey(x *MetaKey, key string) []byte {
	switch p {
	case PROPS_NAME, PROPS_PAGE, PROPS_ENTRY_PAGE, PROPS_EXIT_PAGE,
		PROPS_REFERRER, PROPS_UTM_MEDIUM, PROPS_UTM_SOURCE,
		PROPS_UTM_CAMPAIGN, PROPS_UTM_CONTENT, PROPS_UTM_TERM:
		return x.SetMeta(byte(p)).String(key)
	default:
		return nil
	}
}

func (p PROPS) UAKeys(x *MetaKey, key string) []byte {
	switch p {
	case PROPS_UTM_DEVICE, PROPS_UTM_BROWSER, PROPS_BROWSER_VERSION, PROPS_OS, PROPS_OS_VERSION:
		return x.SetMeta(byte(p)).HashU16(ua.ToIndex(key))
	default:
		return nil
	}
}

func (p PROPS) CountryKey(x *MetaKey, key string) []byte {
	switch p {
	case PROPS_COUNTRY, PROPS_REGION:
		return x.SetMeta(byte(p)).HashU16(cities.ToIndex(key))
	case PROPS_CITY:
		return x.SetMeta(byte(p)).HashU32(uint32(cities.ToIndex(key)))
	default:
		return nil
	}
}
func (p PROPS) CityKey(x *MetaKey, key uint32) []byte {
	switch p {
	case PROPS_CITY:
		return x.SetMeta(byte(p)).HashU32(key)
	default:
		return nil
	}
}

func updateCalendar(ctx context.Context, txn *badger.Txn, ts time.Time, key []byte, a *store.Sum) error {
	cache := caches.Calendar(ctx)
	hash := xxhash.Sum64(key)
	if x, ok := cache.Get(hash); ok {
		// The calendar was in cache we update it and save.
		cal := x.(*store.Calendar)
		a.UpdateCalendar(ts, cal)
		b, err := cal.Message().MarshalPacked()
		if err != nil {
			return fmt.Errorf("failed to marshal calendar %v", err)
		}
		return txn.Set(key, b)
	}
	x, err := txn.Get(key)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			// new entry we store a
			cal, err := store.ZeroCalendar(ts, a)
			if err != nil {
				return err
			}
			defer cache.Set(hash, &cal, store.CacheCost)
			b, err := cal.Message().MarshalPacked()
			if err != nil {
				return fmt.Errorf("failed to marshal calendar %v", err)
			}
			return txn.Set(key, b)
		}
		return err
	}
	return x.Value(func(val []byte) error {
		cal, err := store.CalendarFromBytes(val)
		if err != nil {
			return err
		}
		defer cache.Set(hash, &cal, store.CacheCost)
		a.UpdateCalendar(ts, &cal)
		b, err := cal.Message().MarshalPacked()
		if err != nil {
			return fmt.Errorf("failed to marshal calendar %v", err)
		}
		return txn.Set(key, b)
	})
}
