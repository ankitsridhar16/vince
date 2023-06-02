package timeseries

import (
	"context"
	"math"
	"strconv"
	"time"
)

type Stats struct {
	Timestamps []int64
	Aggregate  Aggregate
	Timeseries []float64
}

type Aggregate struct {
	Metrics map[string]FloatValue
	Props   map[string][]StatValue
}

type StatValue struct {
	Key   string
	Value FloatValue
}

func sum(ls []float64) (o float64) {
	for _, v := range ls {
		o += v
	}
	return
}

func RootQuery(ctx context.Context,
	uid, sid uint64, selectedMetric Metric,
	selectedProp Property,
	key string,
	offset time.Duration,
) (o Stats) {
	q := Query(ctx, QueryRequest{
		UserID: uid,
		SiteID: sid,
		BaseQuery: BaseQuery{
			Offset:  offset,
			Metrics: allMetrics,
			Filters: allProperties(selectedMetric, selectedProp, key),
		},
	})
	o.Timestamps = q.Timestamps
	o.Aggregate = Aggregate{
		Metrics: make(map[string]FloatValue),
		Props:   make(map[string][]StatValue),
	}
	// calculate base stats
	base := q.Result[selectedProp.String()]
	for i := Visitors; i <= VisitDurations; i++ {
		o.Aggregate.Metrics[i.String()] = FloatValue(sum(base[i.String()][key]))
	}
	o.Timeseries = base[selectedMetric.String()][key]
	if len(o.Timeseries) == 0 {
		// no key was found, make sure time/value aligns for the graph.
		o.Timeseries = make([]float64, len(o.Timestamps))
	}
	for i := Event; i <= City; i++ {
		for k, v := range q.Result[i.String()][selectedMetric.String()] {
			o.Aggregate.Props[i.String()] = append(o.Aggregate.Props[i.String()], StatValue{
				Key:   k,
				Value: FloatValue(sum(v)),
			})
		}
	}

	// calculate bounce rate
	visits := o.Aggregate.Metrics[Visits.String()]
	if visits != 0 {
		// avoid dividing by zero, thats why whe check visits != 0
		o.Aggregate.Metrics[BounceRates.String()] =
			(o.Aggregate.Metrics[BounceRates.String()] / visits) * 100
	}
	return
}

var allMetrics = []Metric{
	Visitors,
	Views,
	Events,
	Visits,
	BounceRates,
	VisitDurations,
}

func allProperties(selected Metric, selectedProp Property, key string) []*Filter {
	if selectedProp != Base {
		// No need to select other properties if its not for the base. This is the
		// case when we are searching based on a key
		return []*Filter{
			{Property: selectedProp, Expr: FilterExpr{
				Text: key,
			}},
		}
	}
	o := make([]*Filter, City+1)
	a := make([]Metric, 0, VisitDurations+1)
	for i := Visitors; i <= VisitDurations; i++ {
		if i == selected {
			continue
		}
		a = append(a, i)
	}

	for i := range o {
		if i == int(Base) {
			o[i] = &Filter{
				Property: Property(i),
				Expr: FilterExpr{
					Text: BaseKey,
				},
			}
		} else {
			o[i] = &Filter{
				Property: Property(i),
				Omit:     a,
				Expr: FilterExpr{
					Text: "*",
				},
			}
		}
	}
	return o
}

type FloatValue float64

func (v FloatValue) String() string {
	f := float64(v)
	p := math.Floor(math.Log10(math.Abs(f)))
	if p <= 2 {
		return strconv.FormatFloat(f, 'f', -1, 64)
	}
	x := math.Floor(p / 3)
	s := math.Pow(10, p-x*3) * +(f / math.Pow(10, p))
	s = math.Round(s*100) / 100
	k := []string{"", "K", "M", "B", "T"}
	return strconv.FormatFloat(s, 'f', -1, 64) + k[int(x)]
}