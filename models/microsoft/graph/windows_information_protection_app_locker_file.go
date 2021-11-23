package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// windowsInformationProtectionAppLockerFile 
type WindowsInformationProtectionAppLockerFile struct {
    Entity
    // The friendly name
    displayName *string;
    // File as a byte array
    file []byte;
    // SHA256 hash of the file
    fileHash *string;
    // Version of the entity.
    version *string;
}
// NewWindowsInformationProtectionAppLockerFile instantiates a new windowsInformationProtectionAppLockerFile and sets the default values.
func NewWindowsInformationProtectionAppLockerFile()(*WindowsInformationProtectionAppLockerFile) {
    m := &WindowsInformationProtectionAppLockerFile{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The friendly name
func (m *WindowsInformationProtectionAppLockerFile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFile gets the file property value. File as a byte array
func (m *WindowsInformationProtectionAppLockerFile) GetFile()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
// GetFileHash gets the fileHash property value. SHA256 hash of the file
func (m *WindowsInformationProtectionAppLockerFile) GetFileHash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileHash
    }
}
// GetVersion gets the version property value. Version of the entity.
func (m *WindowsInformationProtectionAppLockerFile) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionAppLockerFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val)
        }
        return nil
    }
    res["fileHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHash(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionAppLockerFile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDisplayName sets the displayName property value. The friendly name
func (m *WindowsInformationProtectionAppLockerFile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFile sets the file property value. File as a byte array
func (m *WindowsInformationProtectionAppLockerFile) SetFile(value []byte)() {
    m.file = value
}
// SetFileHash sets the fileHash property value. SHA256 hash of the file
func (m *WindowsInformationProtectionAppLockerFile) SetFileHash(value *string)() {
    m.fileHash = value
}
// SetVersion sets the version property value. Version of the entity.
func (m *WindowsInformationProtectionAppLockerFile) SetVersion(value *string)() {
    m.version = value
}
