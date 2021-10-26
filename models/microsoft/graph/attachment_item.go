package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttachmentItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of attachment. Possible values are: file, item, reference. Required.
    attachmentType *AttachmentType;
    // The nature of the data in the attachment. Optional.
    contentType *string;
    // true if the attachment is an inline attachment; otherwise, false. Optional.
    isInline *bool;
    // The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
    name *string;
    // The length of the attachment in bytes. Required.
    size *int64;
}
// Instantiates a new attachmentItem and sets the default values.
func NewAttachmentItem()(*AttachmentItem) {
    m := &AttachmentItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttachmentItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attachmentType property value. The type of attachment. Possible values are: file, item, reference. Required.
func (m *AttachmentItem) GetAttachmentType()(*AttachmentType) {
    if m == nil {
        return nil
    } else {
        return m.attachmentType
    }
}
// Gets the contentType property value. The nature of the data in the attachment. Optional.
func (m *AttachmentItem) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// Gets the isInline property value. true if the attachment is an inline attachment; otherwise, false. Optional.
func (m *AttachmentItem) GetIsInline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInline
    }
}
// Gets the name property value. The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
func (m *AttachmentItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the size property value. The length of the attachment in bytes. Required.
func (m *AttachmentItem) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// The deserialization information for the current model
func (m *AttachmentItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attachmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttachmentType)
        if err != nil {
            return err
        }
        cast := val.(AttachmentType)
        m.SetAttachmentType(&cast)
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentType(val)
        return nil
    }
    res["isInline"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInline(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSize(val)
        return nil
    }
    return res
}
func (m *AttachmentItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AttachmentItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAttachmentType() != nil {
        cast := m.GetAttachmentType().String()
        err := writer.WriteStringValue("attachmentType", &cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AttachmentItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attachmentType property value. The type of attachment. Possible values are: file, item, reference. Required.
// Parameters:
//  - value : Value to set for the attachmentType property.
func (m *AttachmentItem) SetAttachmentType(value *AttachmentType)() {
    m.attachmentType = value
}
// Sets the contentType property value. The nature of the data in the attachment. Optional.
// Parameters:
//  - value : Value to set for the contentType property.
func (m *AttachmentItem) SetContentType(value *string)() {
    m.contentType = value
}
// Sets the isInline property value. true if the attachment is an inline attachment; otherwise, false. Optional.
// Parameters:
//  - value : Value to set for the isInline property.
func (m *AttachmentItem) SetIsInline(value *bool)() {
    m.isInline = value
}
// Sets the name property value. The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
// Parameters:
//  - value : Value to set for the name property.
func (m *AttachmentItem) SetName(value *string)() {
    m.name = value
}
// Sets the size property value. The length of the attachment in bytes. Required.
// Parameters:
//  - value : Value to set for the size property.
func (m *AttachmentItem) SetSize(value *int64)() {
    m.size = value
}
