package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceMobileAppConfigurationDeviceSummary 
type ManagedDeviceMobileAppConfigurationDeviceSummary struct {
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
// NewManagedDeviceMobileAppConfigurationDeviceSummary instantiates a new managedDeviceMobileAppConfigurationDeviceSummary and sets the default values.
func NewManagedDeviceMobileAppConfigurationDeviceSummary()(*ManagedDeviceMobileAppConfigurationDeviceSummary) {
    m := &ManagedDeviceMobileAppConfigurationDeviceSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetConfigurationVersion gets the configurationVersion property value. Version of the policy for that overview
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) GetConfigurationVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationVersion
    }
}
// GetErrorCount gets the errorCount property value. Number of error devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) GetErrorCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCount
    }
}
// GetFailedCount gets the failedCount property value. Number of failed devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) GetFailedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedCount
    }
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. Last update time
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) GetNotApplicableCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.notApplicableCount
    }
}
// GetPendingCount gets the pendingCount property value. Number of pending devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) GetPendingCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.pendingCount
    }
}
// GetSuccessCount gets the successCount property value. Number of succeeded devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) GetSuccessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetConfigurationVersion sets the configurationVersion property value. Version of the policy for that overview
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) SetConfigurationVersion(value *int32)() {
    if m != nil {
        m.configurationVersion = value
    }
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) SetErrorCount(value *int32)() {
    if m != nil {
        m.errorCount = value
    }
}
// SetFailedCount sets the failedCount property value. Number of failed devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) SetFailedCount(value *int32)() {
    if m != nil {
        m.failedCount = value
    }
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. Last update time
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdateDateTime = value
    }
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) SetNotApplicableCount(value *int32)() {
    if m != nil {
        m.notApplicableCount = value
    }
}
// SetPendingCount sets the pendingCount property value. Number of pending devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) SetPendingCount(value *int32)() {
    if m != nil {
        m.pendingCount = value
    }
}
// SetSuccessCount sets the successCount property value. Number of succeeded devices
func (m *ManagedDeviceMobileAppConfigurationDeviceSummary) SetSuccessCount(value *int32)() {
    if m != nil {
        m.successCount = value
    }
}
