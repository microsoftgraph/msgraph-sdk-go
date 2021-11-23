package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EBookInstallSummary 
type EBookInstallSummary struct {
    Entity
    // Number of Devices that have failed to install this book.
    failedDeviceCount *int32;
    // Number of Users that have 1 or more device that failed to install this book.
    failedUserCount *int32;
    // Number of Devices that have successfully installed this book.
    installedDeviceCount *int32;
    // Number of Users whose devices have all succeeded to install this book.
    installedUserCount *int32;
    // Number of Devices that does not have this book installed.
    notInstalledDeviceCount *int32;
    // Number of Users that did not install this book.
    notInstalledUserCount *int32;
}
// NewEBookInstallSummary instantiates a new eBookInstallSummary and sets the default values.
func NewEBookInstallSummary()(*EBookInstallSummary) {
    m := &EBookInstallSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Number of Devices that have failed to install this book.
func (m *EBookInstallSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// GetFailedUserCount gets the failedUserCount property value. Number of Users that have 1 or more device that failed to install this book.
func (m *EBookInstallSummary) GetFailedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUserCount
    }
}
// GetInstalledDeviceCount gets the installedDeviceCount property value. Number of Devices that have successfully installed this book.
func (m *EBookInstallSummary) GetInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedDeviceCount
    }
}
// GetInstalledUserCount gets the installedUserCount property value. Number of Users whose devices have all succeeded to install this book.
func (m *EBookInstallSummary) GetInstalledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedUserCount
    }
}
// GetNotInstalledDeviceCount gets the notInstalledDeviceCount property value. Number of Devices that does not have this book installed.
func (m *EBookInstallSummary) GetNotInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledDeviceCount
    }
}
// GetNotInstalledUserCount gets the notInstalledUserCount property value. Number of Users that did not install this book.
func (m *EBookInstallSummary) GetNotInstalledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledUserCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EBookInstallSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    res["failedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUserCount(val)
        }
        return nil
    }
    res["installedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledDeviceCount(val)
        }
        return nil
    }
    res["installedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledUserCount(val)
        }
        return nil
    }
    res["notInstalledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledDeviceCount(val)
        }
        return nil
    }
    res["notInstalledUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledUserCount(val)
        }
        return nil
    }
    return res
}
func (m *EBookInstallSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EBookInstallSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedUserCount", m.GetFailedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installedDeviceCount", m.GetInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installedUserCount", m.GetInstalledUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notInstalledDeviceCount", m.GetNotInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notInstalledUserCount", m.GetNotInstalledUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Number of Devices that have failed to install this book.
func (m *EBookInstallSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
// SetFailedUserCount sets the failedUserCount property value. Number of Users that have 1 or more device that failed to install this book.
func (m *EBookInstallSummary) SetFailedUserCount(value *int32)() {
    m.failedUserCount = value
}
// SetInstalledDeviceCount sets the installedDeviceCount property value. Number of Devices that have successfully installed this book.
func (m *EBookInstallSummary) SetInstalledDeviceCount(value *int32)() {
    m.installedDeviceCount = value
}
// SetInstalledUserCount sets the installedUserCount property value. Number of Users whose devices have all succeeded to install this book.
func (m *EBookInstallSummary) SetInstalledUserCount(value *int32)() {
    m.installedUserCount = value
}
// SetNotInstalledDeviceCount sets the notInstalledDeviceCount property value. Number of Devices that does not have this book installed.
func (m *EBookInstallSummary) SetNotInstalledDeviceCount(value *int32)() {
    m.notInstalledDeviceCount = value
}
// SetNotInstalledUserCount sets the notInstalledUserCount property value. Number of Users that did not install this book.
func (m *EBookInstallSummary) SetNotInstalledUserCount(value *int32)() {
    m.notInstalledUserCount = value
}
