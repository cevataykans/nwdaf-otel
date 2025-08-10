package prometheus

import (
	"fmt"
	"time"
)

func CreateESAvgQuery(service string, start, end time.Time) ESQuery {
	if service == UPFContainer {
		service = UPFPod
	}
	return ESQuery{
		Size: 0,
		Aggs: Aggregations{
			Duration: DurationAgg{
				Filter: BoolFilterAgg{
					Bool: BoolQuery{
						Must: []interface{}{
							WildcardQuery{
								Wildcard: map[string]WildcardField{
									"process.serviceName": {Value: fmt.Sprintf("%s*", service)},
								},
							},
							RangeQuery{
								Range: map[string]RangeField{
									"startTimeMillis": {
										Gte: start.Unix() * 1000,
										Lte: end.Unix() * 1000,
									},
								},
							},
						},
					},
				},
				Aggs: AvgAggs{
					AvgDuration: AvgField{
						Avg: AvgFieldInner{
							Field: "duration",
						},
					},
				},
			},
		},
	}
}

type ESQuery struct {
	Size int          `json:"size"`
	Aggs Aggregations `json:"aggs"`
}

type Aggregations struct {
	Duration DurationAgg `json:"duration"`
}

type DurationAgg struct {
	Filter BoolFilterAgg `json:"filter"`
	Aggs   AvgAggs       `json:"aggs"`
}

type BoolFilterAgg struct {
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

type AvgAggs struct {
	AvgDuration AvgField `json:"avg_duration"`
}

type AvgField struct {
	Avg AvgFieldInner `json:"avg"`
}

type AvgFieldInner struct {
	Field string `json:"field"`
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
		DurationAgg struct {
			DocCount    int `json:"doc_count"`
			AvgDuration struct {
				Value *float64 `json:"value"`
			} `json:"avg_duration"`
		} `json:"duration"`
	} `json:"aggregations"`
}
