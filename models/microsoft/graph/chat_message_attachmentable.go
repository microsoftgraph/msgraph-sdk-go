package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessageAttachmentable 
type ChatMessageAttachmentable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetContent()(*string)
    GetContentType()(*string)
    GetContentUrl()(*string)
    GetId()(*string)
    GetName()(*string)
    GetThumbnailUrl()(*string)
    SetContent(value *string)()
    SetContentType(value *string)()
    SetContentUrl(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetThumbnailUrl(value *string)()
}
