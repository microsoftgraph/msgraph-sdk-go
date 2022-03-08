package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttachmentItemable 
type AttachmentItemable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAttachmentType()(*AttachmentType)
    GetContentId()(*string)
    GetContentType()(*string)
    GetIsInline()(*bool)
    GetName()(*string)
    GetSize()(*int64)
    SetAttachmentType(value *AttachmentType)()
    SetContentId(value *string)()
    SetContentType(value *string)()
    SetIsInline(value *bool)()
    SetName(value *string)()
    SetSize(value *int64)()
}
