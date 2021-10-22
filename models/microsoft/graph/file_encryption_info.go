package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FileEncryptionInfo struct {
    additionalData map[string]interface{};
    encryptionKey []byte;
    fileDigest []byte;
    fileDigestAlgorithm *string;
    initializationVector []byte;
    mac []byte;
    macKey []byte;
    profileIdentifier *string;
}
func NewFileEncryptionInfo()(*FileEncryptionInfo) {
    m := &FileEncryptionInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FileEncryptionInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FileEncryptionInfo) GetEncryptionKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.encryptionKey
    }
}
func (m *FileEncryptionInfo) GetFileDigest()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.fileDigest
    }
}
func (m *FileEncryptionInfo) GetFileDigestAlgorithm()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileDigestAlgorithm
    }
}
func (m *FileEncryptionInfo) GetInitializationVector()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.initializationVector
    }
}
func (m *FileEncryptionInfo) GetMac()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
func (m *FileEncryptionInfo) GetMacKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.macKey
    }
}
func (m *FileEncryptionInfo) GetProfileIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileIdentifier
    }
}
func (m *FileEncryptionInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["encryptionKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetEncryptionKey(val)
        return nil
    }
    res["fileDigest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetFileDigest(val)
        return nil
    }
    res["fileDigestAlgorithm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileDigestAlgorithm(val)
        return nil
    }
    res["initializationVector"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetInitializationVector(val)
        return nil
    }
    res["mac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetMac(val)
        return nil
    }
    res["macKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetMacKey(val)
        return nil
    }
    res["profileIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProfileIdentifier(val)
        return nil
    }
    return res
}
func (m *FileEncryptionInfo) IsNil()(bool) {
    return m == nil
}
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
func (m *FileEncryptionInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FileEncryptionInfo) SetEncryptionKey(value []byte)() {
    m.encryptionKey = value
}
func (m *FileEncryptionInfo) SetFileDigest(value []byte)() {
    m.fileDigest = value
}
func (m *FileEncryptionInfo) SetFileDigestAlgorithm(value *string)() {
    m.fileDigestAlgorithm = value
}
func (m *FileEncryptionInfo) SetInitializationVector(value []byte)() {
    m.initializationVector = value
}
func (m *FileEncryptionInfo) SetMac(value []byte)() {
    m.mac = value
}
func (m *FileEncryptionInfo) SetMacKey(value []byte)() {
    m.macKey = value
}
func (m *FileEncryptionInfo) SetProfileIdentifier(value *string)() {
    m.profileIdentifier = value
}
