package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EBookInstallSummary struct {
    Entity
    failedDeviceCount *int32;
    failedUserCount *int32;
    installedDeviceCount *int32;
    installedUserCount *int32;
    notInstalledDeviceCount *int32;
    notInstalledUserCount *int32;
}
func NewEBookInstallSummary()(*EBookInstallSummary) {
    m := &EBookInstallSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EBookInstallSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
func (m *EBookInstallSummary) GetFailedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUserCount
    }
}
func (m *EBookInstallSummary) GetInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedDeviceCount
    }
}
func (m *EBookInstallSummary) GetInstalledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedUserCount
    }
}
func (m *EBookInstallSummary) GetNotInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledDeviceCount
    }
}
func (m *EBookInstallSummary) GetNotInstalledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledUserCount
    }
}
func (m *EBookInstallSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedDeviceCount(val)
        return nil
    }
    res["failedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedUserCount(val)
        return nil
    }
    res["installedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInstalledDeviceCount(val)
        return nil
    }
    res["installedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInstalledUserCount(val)
        return nil
    }
    res["notInstalledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotInstalledDeviceCount(val)
        return nil
    }
    res["notInstalledUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNotInstalledUserCount(val)
        return nil
    }
    return res
}
func (m *EBookInstallSummary) IsNil()(bool) {
    return m == nil
}
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
func (m *EBookInstallSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
func (m *EBookInstallSummary) SetFailedUserCount(value *int32)() {
    m.failedUserCount = value
}
func (m *EBookInstallSummary) SetInstalledDeviceCount(value *int32)() {
    m.installedDeviceCount = value
}
func (m *EBookInstallSummary) SetInstalledUserCount(value *int32)() {
    m.installedUserCount = value
}
func (m *EBookInstallSummary) SetNotInstalledDeviceCount(value *int32)() {
    m.notInstalledDeviceCount = value
}
func (m *EBookInstallSummary) SetNotInstalledUserCount(value *int32)() {
    m.notInstalledUserCount = value
}
