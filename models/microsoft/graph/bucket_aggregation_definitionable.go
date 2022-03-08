package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BucketAggregationDefinitionable 
type BucketAggregationDefinitionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIsDescending()(*bool)
    GetMinimumCount()(*int32)
    GetPrefixFilter()(*string)
    GetRanges()([]BucketAggregationRangeable)
    GetSortBy()(*BucketAggregationSortProperty)
    SetIsDescending(value *bool)()
    SetMinimumCount(value *int32)()
    SetPrefixFilter(value *string)()
    SetRanges(value []BucketAggregationRangeable)()
    SetSortBy(value *BucketAggregationSortProperty)()
}
