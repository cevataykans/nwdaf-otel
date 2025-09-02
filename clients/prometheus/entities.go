package prometheus

import (
	"encoding/json"
	"fmt"
	"time"
)

func CreateESAvgQuery(service string, start, end time.Time) (ESQuery, error) {
	if service == UPFContainer {
		service = UPFPod
	}
	startMicro := start.UnixNano() / int64(time.Microsecond)
	endMicro := end.UnixNano() / int64(time.Microsecond)

	wildcard, err := json.Marshal(WildcardQuery{
		Wildcard: map[string]WildcardField{
			"process.serviceName": {Value: fmt.Sprintf("%s*", service)},
		},
	})
	if err != nil {
		return ESQuery{}, fmt.Errorf("error, cannot serialize wildcard query into raw bytes")
	}

	rangeQ, err := json.Marshal(RangeQuery{
		Range: map[string]RangeField{
			"startTime": {
				Gte: startMicro,
				Lte: endMicro,
			},
		},
	})
	if err != nil {
		return ESQuery{}, fmt.Errorf("error, cannot serialize range query into raw bytes")
	}

	return ESQuery{
		Size: 0,
		Query: SearchQuery{
			Bool: BoolQuery{
				Must: []json.RawMessage{wildcard, rangeQ},
			},
		},
		Aggs: Aggregations{
			AvgDuration: DurationQuery{
				Avg: AvgAgg{
					AvgField: "duration",
				},
			},
		},
	}, nil
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
	Must []json.RawMessage `json:"must"`
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
