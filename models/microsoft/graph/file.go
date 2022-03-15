package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// File provides operations to manage the collection of drive entities.
type File struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Hashes of the file's binary content, if available. Read-only.
    hashes Hashesable;
    // The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
    mimeType *string;
    // 
    processingMetadata *bool;
}
// NewFile instantiates a new file and sets the default values.
func NewFile()(*File) {
    m := &File{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *File) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *File) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hashes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateHashesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashes(val.(Hashesable))
        }
        return nil
    }
    res["mimeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMimeType(val)
        }
        return nil
    }
    res["processingMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingMetadata(val)
        }
        return nil
    }
    return res
}
// GetHashes gets the hashes property value. Hashes of the file's binary content, if available. Read-only.
func (m *File) GetHashes()(Hashesable) {
    if m == nil {
        return nil
    } else {
        return m.hashes
    }
}
// GetMimeType gets the mimeType property value. The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
func (m *File) GetMimeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mimeType
    }
}
// GetProcessingMetadata gets the processingMetadata property value. 
func (m *File) GetProcessingMetadata()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.processingMetadata
    }
}
func (m *File) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *File) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("hashes", m.GetHashes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mimeType", m.GetMimeType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("processingMetadata", m.GetProcessingMetadata())
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
func (m *File) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHashes sets the hashes property value. Hashes of the file's binary content, if available. Read-only.
func (m *File) SetHashes(value Hashesable)() {
    if m != nil {
        m.hashes = value
    }
}
// SetMimeType sets the mimeType property value. The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
func (m *File) SetMimeType(value *string)() {
    if m != nil {
        m.mimeType = value
    }
}
// SetProcessingMetadata sets the processingMetadata property value. 
func (m *File) SetProcessingMetadata(value *bool)() {
    if m != nil {
        m.processingMetadata = value
    }
}
