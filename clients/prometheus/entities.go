package prometheus

import (
	"fmt"
	"time"
)

func CreateESAvgQuery(service string, start, end time.Time) ESQuery {
	if service == UPFContainer {
		service = UPFPod
	}
	startMicro := start.UnixNano() / int64(time.Microsecond)
	endMicro := end.UnixNano() / int64(time.Microsecond)
	return ESQuery{
		Size: 0,
		Query: SearchQuery{
			Bool: BoolQuery{
				Must: []interface{}{
					WildcardQuery{
						Wildcard: map[string]WildcardField{
							"process.serviceName": {Value: fmt.Sprintf("%s*", service)},
						},
					},
					RangeQuery{
						Range: map[string]RangeField{
							"startTime": {
								Gte: startMicro,
								Lte: endMicro,
							},
						},
					},
				},
			},
		},
		Aggs: Aggregations{
			AvgDuration: DurationQuery{
				Avg: AvgAgg{
					AvgField: "duration",
				},
			},
		},
	}
}

type ESQuery struct {
	Size  int          `json:"size"`
	Query SearchQuery  `json:"query"`
	Aggs  Aggregations `json:"aggs"`
}

type SearchQuery struct {
	Bool BoolQuery `json:"bool"`
}

type BoolQuery struct {
	Must []interface{} `json:"must"`
}

type WildcardQuery struct {
	Wildcard map[string]WildcardField `json:"wildcard"`
}

type WildcardField struct {
	Value string `json:"value"`
}

type RangeQuery struct {
	Range map[string]RangeField `json:"range"`
}

type RangeField struct {
	Gte int64 `json:"gte"`
	Lte int64 `json:"lte"`
}

type Aggregations struct {
	AvgDuration DurationQuery `json:"avg_duration"`
}

type DurationQuery struct {
	Avg AvgAgg `json:"avg"`
}

type AvgAgg struct {
	AvgField string `json:"field"`
}

// Response received from AVG duration for service traces
type ElasticsearchResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Hits     struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
	} `json:"hits"`
	Aggregations struct {
		AvgDuration struct {
			Value *float64 `json:"value"`
		} `json:"avg_duration"`
	} `json:"aggregations"`
}
