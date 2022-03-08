package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VisualInfoable 
type VisualInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAttribution()(ImageInfoable)
    GetBackgroundColor()(*string)
    GetContent()(Jsonable)
    GetDescription()(*string)
    GetDisplayText()(*string)
    SetAttribution(value ImageInfoable)()
    SetBackgroundColor(value *string)()
    SetContent(value Jsonable)()
    SetDescription(value *string)()
    SetDisplayText(value *string)()
}
