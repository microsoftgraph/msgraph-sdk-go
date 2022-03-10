package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// NumberColumnable 
type NumberColumnable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDecimalPlaces()(*string)
    GetDisplayAs()(*string)
    GetMaximum()(*float64)
    GetMinimum()(*float64)
    SetDecimalPlaces(value *string)()
    SetDisplayAs(value *string)()
    SetMaximum(value *float64)()
    SetMinimum(value *float64)()
}
