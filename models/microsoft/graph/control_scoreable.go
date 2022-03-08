package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ControlScoreable 
type ControlScoreable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetControlCategory()(*string)
    GetControlName()(*string)
    GetDescription()(*string)
    GetScore()(*float64)
    SetControlCategory(value *string)()
    SetControlName(value *string)()
    SetDescription(value *string)()
    SetScore(value *float64)()
}
