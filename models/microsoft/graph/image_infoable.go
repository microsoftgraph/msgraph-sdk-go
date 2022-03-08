package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImageInfoable 
type ImageInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAddImageQuery()(*bool)
    GetAlternateText()(*string)
    GetAlternativeText()(*string)
    GetIconUrl()(*string)
    SetAddImageQuery(value *bool)()
    SetAlternateText(value *string)()
    SetAlternativeText(value *string)()
    SetIconUrl(value *string)()
}
