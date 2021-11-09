package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceConfigurationDeviceOverview struct {
    Entity
    // Version of the policy for that overview
    configurationVersion *int32;
    // Number of error devices
    errorCount *int32;
    // Number of failed devices
    failedCount *int32;
    // Last update time
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of not applicable devices
    notApplicableCount *int32;
    // Number of pending devices
    pendingCount *int32;
    // Number of succeeded devices
    successCount *int32;
}
// Instantiates a new deviceConfigurationDeviceOverview and sets the default values.
func NewDeviceConfigurationDeviceOverview()(*DeviceConfigurationDeviceOverview) {
    m := &DeviceConfigurationDeviceOverview{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the configurationVersion property value. Version of the policy for that overview
func (m *DeviceConfigurationDeviceOverview) GetConfigurationVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationVersion
    }
}
// Gets the errorCount property value. Number of error devices
func (m *DeviceConfigurationDeviceOverview) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
}
// Gets the failedCount property value. Number of failed devices
func (m *DeviceConfigurationDeviceOverview) GetFailedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedCount
    }
}
// Gets the lastUpdateDateTime property value. Last update time
func (m *DeviceConfigurationDeviceOverview) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// Gets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceConfigurationDeviceOverview) GetNotApplicableCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
// Gets the pendingCount property value. Number of pending devices
func (m *DeviceConfigurationDeviceOverview) GetPendingCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingCount
    }
}
// Gets the successCount property value. Number of succeeded devices
func (m *DeviceConfigurationDeviceOverview) GetSuccessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successCount
    }
}
// The deserialization information for the current model
func (m *DeviceConfigurationDeviceOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationVersion(val)
        }
        return nil
    }
    res["errorCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["failedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedCount(val)
        }
        return nil
    }
    res["lastUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
        }
        return nil
    }
    res["notApplicableCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableCount(val)
        }
        return nil
    }
    res["pendingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingCount(val)
        }
        return nil
    }
    res["successCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessCount(val)
        }
        return nil
    }
    return res
}
func (m *DeviceConfigurationDeviceOverview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceConfigurationDeviceOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("configurationVersion", m.GetConfigurationVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCount", m.GetErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedCount", m.GetFailedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableCount", m.GetNotApplicableCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingCount", m.GetPendingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successCount", m.GetSuccessCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the configurationVersion property value. Version of the policy for that overview
// Parameters:
//  - value : Value to set for the configurationVersion property.
func (m *DeviceConfigurationDeviceOverview) SetConfigurationVersion(value *int32)() {
    m.configurationVersion = value
}
// Sets the errorCount property value. Number of error devices
// Parameters:
//  - value : Value to set for the errorCount property.
func (m *DeviceConfigurationDeviceOverview) SetErrorCount(value *int32)() {
    m.errorCount = value
}
// Sets the failedCount property value. Number of failed devices
// Parameters:
//  - value : Value to set for the failedCount property.
func (m *DeviceConfigurationDeviceOverview) SetFailedCount(value *int32)() {
    m.failedCount = value
}
// Sets the lastUpdateDateTime property value. Last update time
// Parameters:
//  - value : Value to set for the lastUpdateDateTime property.
func (m *DeviceConfigurationDeviceOverview) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
// Sets the notApplicableCount property value. Number of not applicable devices
// Parameters:
//  - value : Value to set for the notApplicableCount property.
func (m *DeviceConfigurationDeviceOverview) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
// Sets the pendingCount property value. Number of pending devices
// Parameters:
//  - value : Value to set for the pendingCount property.
func (m *DeviceConfigurationDeviceOverview) SetPendingCount(value *int32)() {
    m.pendingCount = value
}
// Sets the successCount property value. Number of succeeded devices
// Parameters:
//  - value : Value to set for the successCount property.
func (m *DeviceConfigurationDeviceOverview) SetSuccessCount(value *int32)() {
    m.successCount = value
}
