package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubjectRightsRequestDetailable 
type SubjectRightsRequestDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExcludedItemCount()(*int32)
    GetInsightCounts()([]KeyValuePairable)
    GetItemCount()(*int32)
    GetItemNeedReview()(*int32)
    GetProductItemCounts()([]KeyValuePairable)
    GetSignedOffItemCount()(*int32)
    GetTotalItemSize()(*int32)
    SetExcludedItemCount(value *int32)()
    SetInsightCounts(value []KeyValuePairable)()
    SetItemCount(value *int32)()
    SetItemNeedReview(value *int32)()
    SetProductItemCounts(value []KeyValuePairable)()
    SetSignedOffItemCount(value *int32)()
    SetTotalItemSize(value *int32)()
}
