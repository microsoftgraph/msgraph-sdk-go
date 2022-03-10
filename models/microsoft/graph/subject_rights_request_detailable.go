package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubjectRightsRequestDetailable 
type SubjectRightsRequestDetailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExcludedItemCount()(*int64)
    GetInsightCounts()([]KeyValuePairable)
    GetItemCount()(*int64)
    GetItemNeedReview()(*int64)
    GetProductItemCounts()([]KeyValuePairable)
    GetSignedOffItemCount()(*int64)
    GetTotalItemSize()(*int64)
    SetExcludedItemCount(value *int64)()
    SetInsightCounts(value []KeyValuePairable)()
    SetItemCount(value *int64)()
    SetItemNeedReview(value *int64)()
    SetProductItemCounts(value []KeyValuePairable)()
    SetSignedOffItemCount(value *int64)()
    SetTotalItemSize(value *int64)()
}
