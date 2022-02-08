package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttachmentItem 
type AttachmentItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of attachment. Possible values are: file, item, reference. Required.
    attachmentType *AttachmentType;
    // The CID or Content-Id of the attachment for referencing in case of in-line attachments using <img src='cid:contentId'> tag in HTML messages. Optional.
    contentId *string;
    // The nature of the data in the attachment. Optional.
    contentType *string;
    // true if the attachment is an inline attachment; otherwise, false. Optional.
    isInline *bool;
    // The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
    name *string;
    // The length of the attachment in bytes. Required.
    size *int64;
}
// NewAttachmentItem instantiates a new attachmentItem and sets the default values.
func NewAttachmentItem()(*AttachmentItem) {
    m := &AttachmentItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttachmentItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttachmentType gets the attachmentType property value. The type of attachment. Possible values are: file, item, reference. Required.
func (m *AttachmentItem) GetAttachmentType()(*AttachmentType) {
    if m == nil {
        return nil
    } else {
        return m.attachmentType
    }
}
// GetContentId gets the contentId property value. The CID or Content-Id of the attachment for referencing in case of in-line attachments using <img src='cid:contentId'> tag in HTML messages. Optional.
func (m *AttachmentItem) GetContentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentId
    }
}
// GetContentType gets the contentType property value. The nature of the data in the attachment. Optional.
func (m *AttachmentItem) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetIsInline gets the isInline property value. true if the attachment is an inline attachment; otherwise, false. Optional.
func (m *AttachmentItem) GetIsInline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInline
    }
}
// GetName gets the name property value. The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
func (m *AttachmentItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSize gets the size property value. The length of the attachment in bytes. Required.
func (m *AttachmentItem) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttachmentItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attachmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttachmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentType(val.(*AttachmentType))
        }
        return nil
    }
    res["contentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentId(val)
        }
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["isInline"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInline(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
func (m *AttachmentItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttachmentItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAttachmentType() != nil {
        cast := (*m.GetAttachmentType()).String()
        err := writer.WriteStringValue("attachmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentId", m.GetContentId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isInline", m.GetIsInline())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttachmentItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttachmentType sets the attachmentType property value. The type of attachment. Possible values are: file, item, reference. Required.
func (m *AttachmentItem) SetAttachmentType(value *AttachmentType)() {
    if m != nil {
        m.attachmentType = value
    }
}
// SetContentId sets the contentId property value. The CID or Content-Id of the attachment for referencing in case of in-line attachments using <img src='cid:contentId'> tag in HTML messages. Optional.
func (m *AttachmentItem) SetContentId(value *string)() {
    if m != nil {
        m.contentId = value
    }
}
// SetContentType sets the contentType property value. The nature of the data in the attachment. Optional.
func (m *AttachmentItem) SetContentType(value *string)() {
    if m != nil {
        m.contentType = value
    }
}
// SetIsInline sets the isInline property value. true if the attachment is an inline attachment; otherwise, false. Optional.
func (m *AttachmentItem) SetIsInline(value *bool)() {
    if m != nil {
        m.isInline = value
    }
}
// SetName sets the name property value. The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
func (m *AttachmentItem) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSize sets the size property value. The length of the attachment in bytes. Required.
func (m *AttachmentItem) SetSize(value *int64)() {
    if m != nil {
        m.size = value
    }
}
