package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintDocumentUploadProperties 
type PrintDocumentUploadProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The document's content (MIME) type.
    contentType *string;
    // The document's name.
    documentName *string;
    // The document's size in bytes.
    size *int64;
}
// NewPrintDocumentUploadProperties instantiates a new printDocumentUploadProperties and sets the default values.
func NewPrintDocumentUploadProperties()(*PrintDocumentUploadProperties) {
    m := &PrintDocumentUploadProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintDocumentUploadProperties) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentType gets the contentType property value. The document's content (MIME) type.
func (m *PrintDocumentUploadProperties) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetDocumentName gets the documentName property value. The document's name.
func (m *PrintDocumentUploadProperties) GetDocumentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentName
    }
}
// GetSize gets the size property value. The document's size in bytes.
func (m *PrintDocumentUploadProperties) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintDocumentUploadProperties) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["documentName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentName(val)
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
func (m *PrintDocumentUploadProperties) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintDocumentUploadProperties) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentType sets the contentType property value. The document's content (MIME) type.
func (m *PrintDocumentUploadProperties) SetContentType(value *string)() {
    if m != nil {
        m.contentType = value
    }
}
// SetDocumentName sets the documentName property value. The document's name.
func (m *PrintDocumentUploadProperties) SetDocumentName(value *string)() {
    if m != nil {
        m.documentName = value
    }
}
// SetSize sets the size property value. The document's size in bytes.
func (m *PrintDocumentUploadProperties) SetSize(value *int64)() {
    if m != nil {
        m.size = value
    }
}
