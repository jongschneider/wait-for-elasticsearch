// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

// ValueCountAggregation is a single-value metrics aggregation that counts
// the number of values that are extracted from the aggregated documents.
// These values can be extracted either from specific fields in the documents,
// or be generated by a provided script. Typically, this aggregator will be
// used in conjunction with other single-value aggregations.
// For example, when computing the avg one might be interested in the
// number of values the average is computed over.
// See: https://www.elastic.co/guide/en/elasticsearch/reference/6.8/search-aggregations-metrics-valuecount-aggregation.html
type ValueCountAggregation struct {
	field           string
	script          *Script
	format          string
	subAggregations map[string]Aggregation
	meta            map[string]interface{}
}

func NewValueCountAggregation() *ValueCountAggregation {
	return &ValueCountAggregation{
		subAggregations: make(map[string]Aggregation),
	}
}

func (a *ValueCountAggregation) Field(field string) *ValueCountAggregation {
	a.field = field
	return a
}

func (a *ValueCountAggregation) Script(script *Script) *ValueCountAggregation {
	a.script = script
	return a
}

func (a *ValueCountAggregation) Format(format string) *ValueCountAggregation {
	a.format = format
	return a
}

func (a *ValueCountAggregation) SubAggregation(name string, subAggregation Aggregation) *ValueCountAggregation {
	a.subAggregations[name] = subAggregation
	return a
}

// Meta sets the meta data to be included in the aggregation response.
func (a *ValueCountAggregation) Meta(metaData map[string]interface{}) *ValueCountAggregation {
	a.meta = metaData
	return a
}

func (a *ValueCountAggregation) Source() (interface{}, error) {
	// Example:
	//	{
	//    "aggs" : {
	//      "grades_count" : { "value_count" : { "field" : "grade" } }
	//    }
	//	}
	// This method returns only the { "value_count" : { "field" : "grade" } } part.

	source := make(map[string]interface{})
	opts := make(map[string]interface{})
	source["value_count"] = opts

	// ValuesSourceAggregationBuilder
	if a.field != "" {
		opts["field"] = a.field
	}
	if a.script != nil {
		src, err := a.script.Source()
		if err != nil {
			return nil, err
		}
		opts["script"] = src
	}
	if a.format != "" {
		opts["format"] = a.format
	}

	// AggregationBuilder (SubAggregations)
	if len(a.subAggregations) > 0 {
		aggsMap := make(map[string]interface{})
		source["aggregations"] = aggsMap
		for name, aggregate := range a.subAggregations {
			src, err := aggregate.Source()
			if err != nil {
				return nil, err
			}
			aggsMap[name] = src
		}
	}

	// Add Meta data if available
	if len(a.meta) > 0 {
		source["meta"] = a.meta
	}

	return source, nil
}
