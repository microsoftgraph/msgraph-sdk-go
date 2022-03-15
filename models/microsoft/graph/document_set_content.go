package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DocumentSetContent provides operations to manage the collection of drive entities.
type DocumentSetContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Content type information of the file.
    contentType ContentTypeInfoable;
    // Name of the file in resource folder that should be added as a default content or a template in the document set.
    fileName *string;
    // Folder name in which the file will be placed when a new document set is created in the library.
    folderName *string;
}
// NewDocumentSetContent instantiates a new documentSetContent and sets the default values.
func NewDocumentSetContent()(*DocumentSetContent) {
    m := &DocumentSetContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDocumentSetContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentSetContentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDocumentSetContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DocumentSetContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentType gets the contentType property value. Content type information of the file.
func (m *DocumentSetContent) GetContentType()(ContentTypeInfoable) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentSetContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentTypeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(ContentTypeInfoable))
        }
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["folderName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderName(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. Name of the file in resource folder that should be added as a default content or a template in the document set.
func (m *DocumentSetContent) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetFolderName gets the folderName property value. Folder name in which the file will be placed when a new document set is created in the library.
func (m *DocumentSetContent) GetFolderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.folderName
    }
}
func (m *DocumentSetContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DocumentSetContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("folderName", m.GetFolderName())
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
func (m *DocumentSetContent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentType sets the contentType property value. Content type information of the file.
func (m *DocumentSetContent) SetContentType(value ContentTypeInfoable)() {
    if m != nil {
        m.contentType = value
    }
}
// SetFileName sets the fileName property value. Name of the file in resource folder that should be added as a default content or a template in the document set.
func (m *DocumentSetContent) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetFolderName sets the folderName property value. Folder name in which the file will be placed when a new document set is created in the library.
func (m *DocumentSetContent) SetFolderName(value *string)() {
    if m != nil {
        m.folderName = value
    }
}
