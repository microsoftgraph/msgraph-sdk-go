package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsInformationProtectionAppLockerFile struct {
    Entity
    displayName *string;
    file []byte;
    fileHash *string;
    version *string;
}
func NewWindowsInformationProtectionAppLockerFile()(*WindowsInformationProtectionAppLockerFile) {
    m := &WindowsInformationProtectionAppLockerFile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WindowsInformationProtectionAppLockerFile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *WindowsInformationProtectionAppLockerFile) GetFile()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
func (m *WindowsInformationProtectionAppLockerFile) GetFileHash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileHash
    }
}
func (m *WindowsInformationProtectionAppLockerFile) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *WindowsInformationProtectionAppLockerFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetFile(val)
        return nil
    }
    res["fileHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileHash(val)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionAppLockerFile) IsNil()(bool) {
    return m == nil
}
func (m *WindowsInformationProtectionAppLockerFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileHash", m.GetFileHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WindowsInformationProtectionAppLockerFile) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *WindowsInformationProtectionAppLockerFile) SetFile(value []byte)() {
    m.file = value
}
func (m *WindowsInformationProtectionAppLockerFile) SetFileHash(value *string)() {
    m.fileHash = value
}
func (m *WindowsInformationProtectionAppLockerFile) SetVersion(value *string)() {
    m.version = value
}
