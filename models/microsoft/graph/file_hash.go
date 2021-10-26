package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type FileHash struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // File hash type. Possible values are: unknown, sha1, sha256, md5, authenticodeHash256, lsHash, ctph, peSha1, peSha256.
    hashType *FileHashType;
    // Value of the file hash.
    hashValue *string;
}
// Instantiates a new fileHash and sets the default values.
func NewFileHash()(*FileHash) {
    m := &FileHash{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileHash) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the hashType property value. File hash type. Possible values are: unknown, sha1, sha256, md5, authenticodeHash256, lsHash, ctph, peSha1, peSha256.
func (m *FileHash) GetHashType()(*FileHashType) {
    if m == nil {
        return nil
    } else {
        return m.hashType
    }
}
// Gets the hashValue property value. Value of the file hash.
func (m *FileHash) GetHashValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hashValue
    }
}
// The deserialization information for the current model
func (m *FileHash) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hashType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileHashType)
        if err != nil {
            return err
        }
        cast := val.(FileHashType)
        m.SetHashType(&cast)
        return nil
    }
    res["hashValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHashValue(val)
        return nil
    }
    return res
}
func (m *FileHash) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *FileHash) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetHashType() != nil {
        cast := m.GetHashType().String()
        err := writer.WriteStringValue("hashType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hashValue", m.GetHashValue())
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
func (m *FileHash) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the hashType property value. File hash type. Possible values are: unknown, sha1, sha256, md5, authenticodeHash256, lsHash, ctph, peSha1, peSha256.
// Parameters:
//  - value : Value to set for the hashType property.
func (m *FileHash) SetHashType(value *FileHashType)() {
    m.hashType = value
}
// Sets the hashValue property value. Value of the file hash.
// Parameters:
//  - value : Value to set for the hashValue property.
func (m *FileHash) SetHashValue(value *string)() {
    m.hashValue = value
}
