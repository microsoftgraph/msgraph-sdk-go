package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedAppPolicyDeploymentSummaryPerApp struct {
    additionalData map[string]interface{};
    configurationAppliedUserCount *int32;
    mobileAppIdentifier *MobileAppIdentifier;
}
func NewManagedAppPolicyDeploymentSummaryPerApp()(*ManagedAppPolicyDeploymentSummaryPerApp) {
    m := &ManagedAppPolicyDeploymentSummaryPerApp{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetConfigurationAppliedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationAppliedUserCount
    }
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetMobileAppIdentifier()(*MobileAppIdentifier) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppIdentifier
    }
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["configurationAppliedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfigurationAppliedUserCount(val)
        return nil
    }
    res["mobileAppIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppIdentifier() })
        if err != nil {
            return err
        }
        m.SetMobileAppIdentifier(val.(*MobileAppIdentifier))
        return nil
    }
    return res
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) IsNil()(bool) {
    return m == nil
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("configurationAppliedUserCount", m.GetConfigurationAppliedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mobileAppIdentifier", m.GetMobileAppIdentifier())
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
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetConfigurationAppliedUserCount(value *int32)() {
    m.configurationAppliedUserCount = value
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetMobileAppIdentifier(value *MobileAppIdentifier)() {
    m.mobileAppIdentifier = value
}
