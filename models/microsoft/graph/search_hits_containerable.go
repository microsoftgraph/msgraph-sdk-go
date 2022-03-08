package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchHitsContainerable 
type SearchHitsContainerable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAggregations()([]SearchAggregationable)
    GetHits()([]SearchHitable)
    GetMoreResultsAvailable()(*bool)
    GetTotal()(*int32)
    SetAggregations(value []SearchAggregationable)()
    SetHits(value []SearchHitable)()
    SetMoreResultsAvailable(value *bool)()
    SetTotal(value *int32)()
}
