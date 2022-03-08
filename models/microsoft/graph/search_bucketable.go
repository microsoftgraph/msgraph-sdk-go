package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchBucketable 
type SearchBucketable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAggregationFilterToken()(*string)
    GetCount()(*int32)
    GetKey()(*string)
    SetAggregationFilterToken(value *string)()
    SetCount(value *int32)()
    SetKey(value *string)()
}
