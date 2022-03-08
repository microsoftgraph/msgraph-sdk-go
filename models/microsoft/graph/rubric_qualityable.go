package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RubricQualityable 
type RubricQualityable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCriteria()([]RubricCriterionable)
    GetDescription()(EducationItemBodyable)
    GetDisplayName()(*string)
    GetQualityId()(*string)
    GetWeight()(*float32)
    SetCriteria(value []RubricCriterionable)()
    SetDescription(value EducationItemBodyable)()
    SetDisplayName(value *string)()
    SetQualityId(value *string)()
    SetWeight(value *float32)()
}
