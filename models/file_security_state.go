package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileSecurityState 
type FileSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Complex type containing file hashes (cryptographic and location-sensitive).
    fileHash FileHashable;
    // File name (without path).
    name *string;
    // Full file path of the file/imageFile.
    path *string;
    // Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a percentage.
    riskScore *string;
}
// NewFileSecurityState instantiates a new fileSecurityState and sets the default values.
func NewFileSecurityState()(*FileSecurityState) {
    m := &FileSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFileSecurityStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileSecurityStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileSecurityState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fileHash"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileHashFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHash(val.(FileHashable))
        }
        return nil
    }
    res["name"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["path"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["riskScore"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    return res
}
// GetFileHash gets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
func (m *FileSecurityState) GetFileHash()(FileHashable) {
    if m == nil {
        return nil
    } else {
        return m.fileHash
    }
}
// GetName gets the name property value. File name (without path).
func (m *FileSecurityState) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetPath gets the path property value. Full file path of the file/imageFile.
func (m *FileSecurityState) GetPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.path
    }
}
// GetRiskScore gets the riskScore property value. Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a percentage.
func (m *FileSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// Serialize serializes information the current object
func (m *FileSecurityState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileSecurityState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFileHash sets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
func (m *FileSecurityState) SetFileHash(value FileHashable)() {
    if m != nil {
        m.fileHash = value
    }
}
// SetName sets the name property value. File name (without path).
func (m *FileSecurityState) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetPath sets the path property value. Full file path of the file/imageFile.
func (m *FileSecurityState) SetPath(value *string)() {
    if m != nil {
        m.path = value
    }
}
// SetRiskScore sets the riskScore property value. Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a percentage.
func (m *FileSecurityState) SetRiskScore(value *string)() {
    if m != nil {
        m.riskScore = value
    }
}
