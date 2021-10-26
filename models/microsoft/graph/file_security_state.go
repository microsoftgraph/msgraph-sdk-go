package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type FileSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Complex type containing file hashes (cryptographic and location-sensitive).
    fileHash *FileHash;
    // File name (without path).
    name *string;
    // Full file path of the file/imageFile.
    path *string;
    // Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a percentage.
    riskScore *string;
}
// Instantiates a new fileSecurityState and sets the default values.
func NewFileSecurityState()(*FileSecurityState) {
    m := &FileSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
func (m *FileSecurityState) GetFileHash()(*FileHash) {
    if m == nil {
        return nil
    } else {
        return m.fileHash
    }
}
// Gets the name property value. File name (without path).
func (m *FileSecurityState) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the path property value. Full file path of the file/imageFile.
func (m *FileSecurityState) GetPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.path
    }
}
// Gets the riskScore property value. Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a percentage.
func (m *FileSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// The deserialization information for the current model
func (m *FileSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["fileHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileHash() })
        if err != nil {
            return err
        }
        m.SetFileHash(val.(*FileHash))
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
    res["path"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPath(val)
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskScore(val)
        return nil
    }
    return res
}
func (m *FileSecurityState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *FileSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("fileHash", m.GetFileHash())
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
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
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
func (m *FileSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
// Parameters:
//  - value : Value to set for the fileHash property.
func (m *FileSecurityState) SetFileHash(value *FileHash)() {
    m.fileHash = value
}
// Sets the name property value. File name (without path).
// Parameters:
//  - value : Value to set for the name property.
func (m *FileSecurityState) SetName(value *string)() {
    m.name = value
}
// Sets the path property value. Full file path of the file/imageFile.
// Parameters:
//  - value : Value to set for the path property.
func (m *FileSecurityState) SetPath(value *string)() {
    m.path = value
}
// Sets the riskScore property value. Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a percentage.
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *FileSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
