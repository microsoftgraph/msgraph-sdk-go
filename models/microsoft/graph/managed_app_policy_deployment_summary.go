package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppPolicyDeploymentSummary provides operations to manage the deviceAppManagement singleton.
type ManagedAppPolicyDeploymentSummary struct {
    Entity
    // Not yet documented
    configurationDeployedUserCount *int32;
    // Not yet documented
    configurationDeploymentSummaryPerApp []ManagedAppPolicyDeploymentSummaryPerAppable;
    // Not yet documented
    displayName *string;
    // Not yet documented
    lastRefreshTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Version of the entity.
    version *string;
}
// NewManagedAppPolicyDeploymentSummary instantiates a new managedAppPolicyDeploymentSummary and sets the default values.
func NewManagedAppPolicyDeploymentSummary()(*ManagedAppPolicyDeploymentSummary) {
    m := &ManagedAppPolicyDeploymentSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewManagedAppPolicyDeploymentSummary(), nil
}
// GetConfigurationDeployedUserCount gets the configurationDeployedUserCount property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) GetConfigurationDeployedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationDeployedUserCount
    }
}
// GetConfigurationDeploymentSummaryPerApp gets the configurationDeploymentSummaryPerApp property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) GetConfigurationDeploymentSummaryPerApp()([]ManagedAppPolicyDeploymentSummaryPerAppable) {
    if m == nil {
        return nil
    } else {
        return m.configurationDeploymentSummaryPerApp
    }
}
// GetDisplayName gets the displayName property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppPolicyDeploymentSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationDeployedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationDeployedUserCount(val)
        }
        return nil
    }
    res["configurationDeploymentSummaryPerApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppPolicyDeploymentSummaryPerAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicyDeploymentSummaryPerAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppPolicyDeploymentSummaryPerAppable)
            }
            m.SetConfigurationDeploymentSummaryPerApp(res)
        }
        return nil
    }
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
    res["lastRefreshTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshTime(val)
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
// GetLastRefreshTime gets the lastRefreshTime property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) GetLastRefreshTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshTime
    }
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedAppPolicyDeploymentSummary) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *ManagedAppPolicyDeploymentSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetConfigurationDeploymentSummaryPerApp() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConfigurationDeploymentSummaryPerApp()))
        for i, v := range m.GetConfigurationDeploymentSummaryPerApp() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetConfigurationDeployedUserCount sets the configurationDeployedUserCount property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) SetConfigurationDeployedUserCount(value *int32)() {
    if m != nil {
        m.configurationDeployedUserCount = value
    }
}
// SetConfigurationDeploymentSummaryPerApp sets the configurationDeploymentSummaryPerApp property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) SetConfigurationDeploymentSummaryPerApp(value []ManagedAppPolicyDeploymentSummaryPerAppable)() {
    if m != nil {
        m.configurationDeploymentSummaryPerApp = value
    }
}
// SetDisplayName sets the displayName property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastRefreshTime sets the lastRefreshTime property value. Not yet documented
func (m *ManagedAppPolicyDeploymentSummary) SetLastRefreshTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRefreshTime = value
    }
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedAppPolicyDeploymentSummary) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
