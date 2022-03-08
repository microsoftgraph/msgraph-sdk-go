package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchHitable 
type SearchHitable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetContentSource()(*string)
    GetHitId()(*string)
    GetRank()(*int32)
    GetResource()(Entityable)
    GetResultTemplateId()(*string)
    GetSummary()(*string)
    SetContentSource(value *string)()
    SetHitId(value *string)()
    SetRank(value *int32)()
    SetResource(value Entityable)()
    SetResultTemplateId(value *string)()
    SetSummary(value *string)()
}
