package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttachmentItem struct {
    additionalData map[string]interface{};
    attachmentType *AttachmentType;
    contentType *string;
    isInline *bool;
    name *string;
    size *int64;
}
func NewAttachmentItem()(*AttachmentItem) {
    m := &AttachmentItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttachmentItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AttachmentItem) GetAttachmentType()(*AttachmentType) {
    if m == nil {
        return nil
    } else {
        return m.attachmentType
    }
}
func (m *AttachmentItem) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
func (m *AttachmentItem) GetIsInline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInline
    }
}
func (m *AttachmentItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *AttachmentItem) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
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
func (m *AttachmentItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttachmentItem) SetAttachmentType(value *AttachmentType)() {
    m.attachmentType = value
}
func (m *AttachmentItem) SetContentType(value *string)() {
    m.contentType = value
}
func (m *AttachmentItem) SetIsInline(value *bool)() {
    m.isInline = value
}
func (m *AttachmentItem) SetName(value *string)() {
    m.name = value
}
func (m *AttachmentItem) SetSize(value *int64)() {
    m.size = value
}
