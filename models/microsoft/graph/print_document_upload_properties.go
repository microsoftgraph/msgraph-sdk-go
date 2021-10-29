package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new printDocumentUploadProperties and sets the default values.
func NewPrintDocumentUploadProperties()(*PrintDocumentUploadProperties) {
    m := &PrintDocumentUploadProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintDocumentUploadProperties) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the contentType property value. The document's content (MIME) type.
func (m *PrintDocumentUploadProperties) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// Gets the documentName property value. The document's name.
func (m *PrintDocumentUploadProperties) GetDocumentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.documentName
    }
}
// Gets the size property value. The document's size in bytes.
func (m *PrintDocumentUploadProperties) GetSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PrintDocumentUploadProperties) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the contentType property value. The document's content (MIME) type.
// Parameters:
//  - value : Value to set for the contentType property.
func (m *PrintDocumentUploadProperties) SetContentType(value *string)() {
    m.contentType = value
}
// Sets the documentName property value. The document's name.
// Parameters:
//  - value : Value to set for the documentName property.
func (m *PrintDocumentUploadProperties) SetDocumentName(value *string)() {
    m.documentName = value
}
// Sets the size property value. The document's size in bytes.
// Parameters:
//  - value : Value to set for the size property.
func (m *PrintDocumentUploadProperties) SetSize(value *int64)() {
    m.size = value
}
