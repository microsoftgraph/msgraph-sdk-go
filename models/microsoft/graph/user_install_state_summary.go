package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// userInstallStateSummary 
type UserInstallStateSummary struct {
    Entity
    // The install state of the eBook.
    deviceStates []DeviceInstallState;
    // Failed Device Count.
    failedDeviceCount *int32;
    // Installed Device Count.
    installedDeviceCount *int32;
    // Not installed device count.
    notInstalledDeviceCount *int32;
    // User name.
    userName *string;
}
// NewUserInstallStateSummary instantiates a new userInstallStateSummary and sets the default values.
func NewUserInstallStateSummary()(*UserInstallStateSummary) {
    m := &UserInstallStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeviceStates gets the deviceStates property value. The install state of the eBook.
func (m *UserInstallStateSummary) GetDeviceStates()([]DeviceInstallState) {
    if m == nil {
        return nil
    } else {
        return m.deviceStates
    }
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Failed Device Count.
func (m *UserInstallStateSummary) GetFailedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedDeviceCount
    }
}
// GetInstalledDeviceCount gets the installedDeviceCount property value. Installed Device Count.
func (m *UserInstallStateSummary) GetInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.installedDeviceCount
    }
}
// GetNotInstalledDeviceCount gets the notInstalledDeviceCount property value. Not installed device count.
func (m *UserInstallStateSummary) GetNotInstalledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notInstalledDeviceCount
    }
}
// GetUserName gets the userName property value. User name.
func (m *UserInstallStateSummary) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserInstallStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceInstallState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceInstallState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceInstallState))
            }
            m.SetDeviceStates(res)
        }
        return nil
    }
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
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
func (m *UserInstallStateSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserInstallStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
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
        err = writer.WriteInt32Value("notInstalledDeviceCount", m.GetNotInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceStates sets the deviceStates property value. The install state of the eBook.
func (m *UserInstallStateSummary) SetDeviceStates(value []DeviceInstallState)() {
    m.deviceStates = value
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Failed Device Count.
func (m *UserInstallStateSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
// SetInstalledDeviceCount sets the installedDeviceCount property value. Installed Device Count.
func (m *UserInstallStateSummary) SetInstalledDeviceCount(value *int32)() {
    m.installedDeviceCount = value
}
// SetNotInstalledDeviceCount sets the notInstalledDeviceCount property value. Not installed device count.
func (m *UserInstallStateSummary) SetNotInstalledDeviceCount(value *int32)() {
    m.notInstalledDeviceCount = value
}
// SetUserName sets the userName property value. User name.
func (m *UserInstallStateSummary) SetUserName(value *string)() {
    m.userName = value
}
