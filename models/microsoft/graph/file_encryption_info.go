package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// fileEncryptionInfo 
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
// NewFileEncryptionInfo instantiates a new fileEncryptionInfo and sets the default values.
func NewFileEncryptionInfo()(*FileEncryptionInfo) {
    m := &FileEncryptionInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileEncryptionInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEncryptionKey gets the encryptionKey property value. The key used to encrypt the file content.
func (m *FileEncryptionInfo) GetEncryptionKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.encryptionKey
    }
}
// GetFileDigest gets the fileDigest property value. The file digest prior to encryption.
func (m *FileEncryptionInfo) GetFileDigest()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.fileDigest
    }
}
// GetFileDigestAlgorithm gets the fileDigestAlgorithm property value. The file digest algorithm.
func (m *FileEncryptionInfo) GetFileDigestAlgorithm()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileDigestAlgorithm
    }
}
// GetInitializationVector gets the initializationVector property value. The initialization vector used for the encryption algorithm.
func (m *FileEncryptionInfo) GetInitializationVector()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.initializationVector
    }
}
// GetMac gets the mac property value. The hash of the encrypted file content + IV (content hash).
func (m *FileEncryptionInfo) GetMac()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
// GetMacKey gets the macKey property value. The key used to get mac.
func (m *FileEncryptionInfo) GetMacKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.macKey
    }
}
// GetProfileIdentifier gets the profileIdentifier property value. The the profile identifier.
func (m *FileEncryptionInfo) GetProfileIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileIdentifier
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileEncryptionInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEncryptionKey sets the encryptionKey property value. The key used to encrypt the file content.
func (m *FileEncryptionInfo) SetEncryptionKey(value []byte)() {
    m.encryptionKey = value
}
// SetFileDigest sets the fileDigest property value. The file digest prior to encryption.
func (m *FileEncryptionInfo) SetFileDigest(value []byte)() {
    m.fileDigest = value
}
// SetFileDigestAlgorithm sets the fileDigestAlgorithm property value. The file digest algorithm.
func (m *FileEncryptionInfo) SetFileDigestAlgorithm(value *string)() {
    m.fileDigestAlgorithm = value
}
// SetInitializationVector sets the initializationVector property value. The initialization vector used for the encryption algorithm.
func (m *FileEncryptionInfo) SetInitializationVector(value []byte)() {
    m.initializationVector = value
}
// SetMac sets the mac property value. The hash of the encrypted file content + IV (content hash).
func (m *FileEncryptionInfo) SetMac(value []byte)() {
    m.mac = value
}
// SetMacKey sets the macKey property value. The key used to get mac.
func (m *FileEncryptionInfo) SetMacKey(value []byte)() {
    m.macKey = value
}
// SetProfileIdentifier sets the profileIdentifier property value. The the profile identifier.
func (m *FileEncryptionInfo) SetProfileIdentifier(value *string)() {
    m.profileIdentifier = value
}
