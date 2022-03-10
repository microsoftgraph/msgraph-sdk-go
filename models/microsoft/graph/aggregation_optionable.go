package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AggregationOptionable 
type AggregationOptionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBucketDefinition()(BucketAggregationDefinitionable)
    GetField()(*string)
    GetSize()(*int32)
    SetBucketDefinition(value BucketAggregationDefinitionable)()
    SetField(value *string)()
    SetSize(value *int32)()
}
