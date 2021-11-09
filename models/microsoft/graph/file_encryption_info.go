package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type FileEncryptionInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The key used to encrypt the file content.
    encryptionKey []byte;
    // The file digest prior to encryption.
    fileDigest []byte;
    // The file digest algorithm.
    fileDigestAlgorithm *string;
    // The initialization vector used for the encryption algorithm.
    initializationVector []byte;
    // The hash of the encrypted file content + IV (content hash).
    mac []byte;
    // The key used to get mac.
    macKey []byte;
    // The the profile identifier.
    profileIdentifier *string;
}
// Instantiates a new fileEncryptionInfo and sets the default values.
func NewFileEncryptionInfo()(*FileEncryptionInfo) {
    m := &FileEncryptionInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileEncryptionInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the encryptionKey property value. The key used to encrypt the file content.
func (m *FileEncryptionInfo) GetEncryptionKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.encryptionKey
    }
}
// Gets the fileDigest property value. The file digest prior to encryption.
func (m *FileEncryptionInfo) GetFileDigest()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.fileDigest
    }
}
// Gets the fileDigestAlgorithm property value. The file digest algorithm.
func (m *FileEncryptionInfo) GetFileDigestAlgorithm()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileDigestAlgorithm
    }
}
// Gets the initializationVector property value. The initialization vector used for the encryption algorithm.
func (m *FileEncryptionInfo) GetInitializationVector()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.initializationVector
    }
}
// Gets the mac property value. The hash of the encrypted file content + IV (content hash).
func (m *FileEncryptionInfo) GetMac()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
// Gets the macKey property value. The key used to get mac.
func (m *FileEncryptionInfo) GetMacKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.macKey
    }
}
// Gets the profileIdentifier property value. The the profile identifier.
func (m *FileEncryptionInfo) GetProfileIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileIdentifier
    }
}
// The deserialization information for the current model
func (m *FileEncryptionInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["encryptionKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionKey(val)
        }
        return nil
    }
    res["fileDigest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileDigest(val)
        }
        return nil
    }
    res["fileDigestAlgorithm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileDigestAlgorithm(val)
        }
        return nil
    }
    res["initializationVector"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitializationVector(val)
        }
        return nil
    }
    res["mac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMac(val)
        }
        return nil
    }
    res["macKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacKey(val)
        }
        return nil
    }
    res["profileIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileIdentifier(val)
        }
        return nil
    }
    return res
}
func (m *FileEncryptionInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *FileEncryptionInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("encryptionKey", m.GetEncryptionKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("fileDigest", m.GetFileDigest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileDigestAlgorithm", m.GetFileDigestAlgorithm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("initializationVector", m.GetInitializationVector())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("mac", m.GetMac())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("macKey", m.GetMacKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("profileIdentifier", m.GetProfileIdentifier())
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
func (m *FileEncryptionInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the encryptionKey property value. The key used to encrypt the file content.
// Parameters:
//  - value : Value to set for the encryptionKey property.
func (m *FileEncryptionInfo) SetEncryptionKey(value []byte)() {
    m.encryptionKey = value
}
// Sets the fileDigest property value. The file digest prior to encryption.
// Parameters:
//  - value : Value to set for the fileDigest property.
func (m *FileEncryptionInfo) SetFileDigest(value []byte)() {
    m.fileDigest = value
}
// Sets the fileDigestAlgorithm property value. The file digest algorithm.
// Parameters:
//  - value : Value to set for the fileDigestAlgorithm property.
func (m *FileEncryptionInfo) SetFileDigestAlgorithm(value *string)() {
    m.fileDigestAlgorithm = value
}
// Sets the initializationVector property value. The initialization vector used for the encryption algorithm.
// Parameters:
//  - value : Value to set for the initializationVector property.
func (m *FileEncryptionInfo) SetInitializationVector(value []byte)() {
    m.initializationVector = value
}
// Sets the mac property value. The hash of the encrypted file content + IV (content hash).
// Parameters:
//  - value : Value to set for the mac property.
func (m *FileEncryptionInfo) SetMac(value []byte)() {
    m.mac = value
}
// Sets the macKey property value. The key used to get mac.
// Parameters:
//  - value : Value to set for the macKey property.
func (m *FileEncryptionInfo) SetMacKey(value []byte)() {
    m.macKey = value
}
// Sets the profileIdentifier property value. The the profile identifier.
// Parameters:
//  - value : Value to set for the profileIdentifier property.
func (m *FileEncryptionInfo) SetProfileIdentifier(value *string)() {
    m.profileIdentifier = value
}
