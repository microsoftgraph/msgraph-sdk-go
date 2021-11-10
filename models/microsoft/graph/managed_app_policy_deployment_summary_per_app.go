package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedAppPolicyDeploymentSummaryPerApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of users the policy is applied.
    configurationAppliedUserCount *int32;
    // Deployment of an app.
    mobileAppIdentifier *MobileAppIdentifier;
}
// Instantiates a new managedAppPolicyDeploymentSummaryPerApp and sets the default values.
func NewManagedAppPolicyDeploymentSummaryPerApp()(*ManagedAppPolicyDeploymentSummaryPerApp) {
    m := &ManagedAppPolicyDeploymentSummaryPerApp{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the configurationAppliedUserCount property value. Number of users the policy is applied.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetConfigurationAppliedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationAppliedUserCount
    }
}
// Gets the mobileAppIdentifier property value. Deployment of an app.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetMobileAppIdentifier()(*MobileAppIdentifier) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppIdentifier
    }
}
// The deserialization information for the current model
func (m *ManagedAppPolicyDeploymentSummaryPerApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["configurationAppliedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationAppliedUserCount(val)
        }
        return nil
    }
    res["mobileAppIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppIdentifier() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileAppIdentifier(val.(*MobileAppIdentifier))
        }
        return nil
    }
    return res
}
func (m *ManagedAppPolicyDeploymentSummaryPerApp) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the configurationAppliedUserCount property value. Number of users the policy is applied.
// Parameters:
//  - value : Value to set for the configurationAppliedUserCount property.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetConfigurationAppliedUserCount(value *int32)() {
    m.configurationAppliedUserCount = value
}
// Sets the mobileAppIdentifier property value. Deployment of an app.
// Parameters:
//  - value : Value to set for the mobileAppIdentifier property.
func (m *ManagedAppPolicyDeploymentSummaryPerApp) SetMobileAppIdentifier(value *MobileAppIdentifier)() {
    m.mobileAppIdentifier = value
}
