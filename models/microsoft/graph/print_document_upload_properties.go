package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintDocumentUploadProperties struct {
    additionalData map[string]interface{};
    contentType *string;
    documentName *string;
    size *int64;
}
func NewPrintDocumentUploadProperties()(*PrintDocumentUploadProperties) {
    m := &PrintDocumentUploadProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrintDocumentUploadProperties) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrintDocumentUploadProperties) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
func (m *PrintDocumentUploadProperties) GetDocumentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentName
    }
}
func (m *PrintDocumentUploadProperties) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *PrintDocumentUploadProperties) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentType(val)
        return nil
    }
    res["documentName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDocumentName(val)
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
func (m *PrintDocumentUploadProperties) IsNil()(bool) {
    return m == nil
}
func (m *PrintDocumentUploadProperties) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("documentName", m.GetDocumentName())
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
func (m *PrintDocumentUploadProperties) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrintDocumentUploadProperties) SetContentType(value *string)() {
    m.contentType = value
}
func (m *PrintDocumentUploadProperties) SetDocumentName(value *string)() {
    m.documentName = value
}
func (m *PrintDocumentUploadProperties) SetSize(value *int64)() {
    m.size = value
}
