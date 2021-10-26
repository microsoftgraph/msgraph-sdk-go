package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedAppPolicyDeploymentSummary struct {
    Entity
    // Not yet documented
    configurationDeployedUserCount *int32;
    // Not yet documented
    configurationDeploymentSummaryPerApp []ManagedAppPolicyDeploymentSummaryPerApp;
    // Not yet documented
    displayName *string;
    // Not yet documented
    lastRefreshTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Version of the entity.
    version *string;
}
// Instantiates a new managedAppPolicyDeploymentSummary and sets the default values.
func NewManagedAppPolicyDeploymentSummary()(*ManagedAppPolicyDeploymentSummary) {
    m := &ManagedAppPolicyDeploymentSummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the configurationDeployedUserCount property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationDeployedUserCount
    }
}
// Gets the configurationDeploymentSummaryPerApp property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerApp()([]ManagedAppPolicyDeploymentSummaryPerApp) {
    if m == nil {
        return nil
    } else {
        return m.configurationDeploymentSummaryPerApp
    }
}
// Gets the displayName property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastRefreshTime property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) GetLastRefreshTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshTime
    }
}
// Gets the version property value. Version of the entity.
func (m *ManagedAppPolicyDeploymentSummary) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *ManagedAppPolicyDeploymentSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationDeployedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfigurationDeployedUserCount(val)
        return nil
    }
    res["configurationDeploymentSummaryPerApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicyDeploymentSummaryPerApp() })
        if err != nil {
            return err
        }
        res := make([]ManagedAppPolicyDeploymentSummaryPerApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppPolicyDeploymentSummaryPerApp))
        }
        m.SetConfigurationDeploymentSummaryPerApp(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastRefreshTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastRefreshTime(val)
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
func (m *ManagedAppPolicyDeploymentSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedAppPolicyDeploymentSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("configurationDeployedUserCount", m.GetConfigurationDeployedUserCount())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConfigurationDeploymentSummaryPerApp()))
        for i, v := range m.GetConfigurationDeploymentSummaryPerApp() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("configurationDeploymentSummaryPerApp", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshTime", m.GetLastRefreshTime())
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
// Sets the configurationDeployedUserCount property value. Not yet documented
// Parameters:
//  - value : Value to set for the configurationDeployedUserCount property.
func (m *ManagedAppPolicyDeploymentSummary) SetConfigurationDeployedUserCount(value *int32)() {
    m.configurationDeployedUserCount = value
}
// Sets the configurationDeploymentSummaryPerApp property value. Not yet documented
// Parameters:
//  - value : Value to set for the configurationDeploymentSummaryPerApp property.
func (m *ManagedAppPolicyDeploymentSummary) SetConfigurationDeploymentSummaryPerApp(value []ManagedAppPolicyDeploymentSummaryPerApp)() {
    m.configurationDeploymentSummaryPerApp = value
}
// Sets the displayName property value. Not yet documented
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ManagedAppPolicyDeploymentSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastRefreshTime property value. Not yet documented
// Parameters:
//  - value : Value to set for the lastRefreshTime property.
func (m *ManagedAppPolicyDeploymentSummary) SetLastRefreshTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshTime = value
}
// Sets the version property value. Version of the entity.
// Parameters:
//  - value : Value to set for the version property.
func (m *ManagedAppPolicyDeploymentSummary) SetVersion(value *string)() {
    m.version = value
}
