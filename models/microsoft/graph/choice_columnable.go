package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChoiceColumnable 
type ChoiceColumnable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAllowTextEntry()(*bool)
    GetChoices()([]string)
    GetDisplayAs()(*string)
    SetAllowTextEntry(value *bool)()
    SetChoices(value []string)()
    SetDisplayAs(value *string)()
}
