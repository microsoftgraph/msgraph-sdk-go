package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type File struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Hashes of the file's binary content, if available. Read-only.
    hashes *Hashes;
    // The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
    mimeType *string;
    // 
    processingMetadata *bool;
}
// Instantiates a new file and sets the default values.
func NewFile()(*File) {
    m := &File{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *File) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the hashes property value. Hashes of the file's binary content, if available. Read-only.
func (m *File) GetHashes()(*Hashes) {
    if m == nil {
        return nil
    } else {
        return m.hashes
    }
}
// Gets the mimeType property value. The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
func (m *File) GetMimeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mimeType
    }
}
// Gets the processingMetadata property value. 
func (m *File) GetProcessingMetadata()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.processingMetadata
    }
}
// The deserialization information for the current model
func (m *File) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hashes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHashes() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashes(val.(*Hashes))
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
func (m *File) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *File) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the hashes property value. Hashes of the file's binary content, if available. Read-only.
// Parameters:
//  - value : Value to set for the hashes property.
func (m *File) SetHashes(value *Hashes)() {
    m.hashes = value
}
// Sets the mimeType property value. The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
// Parameters:
//  - value : Value to set for the mimeType property.
func (m *File) SetMimeType(value *string)() {
    m.mimeType = value
}
// Sets the processingMetadata property value. 
// Parameters:
//  - value : Value to set for the processingMetadata property.
func (m *File) SetProcessingMetadata(value *bool)() {
    m.processingMetadata = value
}
