package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Thumbnailable 
type Thumbnailable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContent()([]byte)
    GetHeight()(*int32)
    GetSourceItemId()(*string)
    GetUrl()(*string)
    GetWidth()(*int32)
    SetContent(value []byte)()
    SetHeight(value *int32)()
    SetSourceItemId(value *string)()
    SetUrl(value *string)()
    SetWidth(value *int32)()
}
