package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConfigurationManagerClientEnabledFeatures provides operations to manage the deviceManagement singleton.
type ConfigurationManagerClientEnabledFeatures struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Whether compliance policy is managed by Intune
    compliancePolicy *bool;
    // Whether device configuration is managed by Intune
    deviceConfiguration *bool;
    // Whether inventory is managed by Intune
    inventory *bool;
    // Whether modern application is managed by Intune
    modernApps *bool;
    // Whether resource access is managed by Intune
    resourceAccess *bool;
    // Whether Windows Update for Business is managed by Intune
    windowsUpdateForBusiness *bool;
}
// NewConfigurationManagerClientEnabledFeatures instantiates a new configurationManagerClientEnabledFeatures and sets the default values.
func NewConfigurationManagerClientEnabledFeatures()(*ConfigurationManagerClientEnabledFeatures) {
    m := &ConfigurationManagerClientEnabledFeatures{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewConfigurationManagerClientEnabledFeatures(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientEnabledFeatures) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompliancePolicy gets the compliancePolicy property value. Whether compliance policy is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetCompliancePolicy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicy
    }
}
// GetDeviceConfiguration gets the deviceConfiguration property value. Whether device configuration is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetDeviceConfiguration()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfiguration
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerClientEnabledFeatures) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compliancePolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicy(val)
        }
        return nil
    }
    res["deviceConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfiguration(val)
        }
        return nil
    }
    res["inventory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventory(val)
        }
        return nil
    }
    res["modernApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModernApps(val)
        }
        return nil
    }
    res["resourceAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAccess(val)
        }
        return nil
    }
    res["windowsUpdateForBusiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsUpdateForBusiness(val)
        }
        return nil
    }
    return res
}
// GetInventory gets the inventory property value. Whether inventory is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetInventory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.inventory
    }
}
// GetModernApps gets the modernApps property value. Whether modern application is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetModernApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.modernApps
    }
}
// GetResourceAccess gets the resourceAccess property value. Whether resource access is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetResourceAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccess
    }
}
// GetWindowsUpdateForBusiness gets the windowsUpdateForBusiness property value. Whether Windows Update for Business is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetWindowsUpdateForBusiness()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsUpdateForBusiness
    }
}
func (m *ConfigurationManagerClientEnabledFeatures) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConfigurationManagerClientEnabledFeatures) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("compliancePolicy", m.GetCompliancePolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deviceConfiguration", m.GetDeviceConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inventory", m.GetInventory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("modernApps", m.GetModernApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("resourceAccess", m.GetResourceAccess())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("windowsUpdateForBusiness", m.GetWindowsUpdateForBusiness())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientEnabledFeatures) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompliancePolicy sets the compliancePolicy property value. Whether compliance policy is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetCompliancePolicy(value *bool)() {
    if m != nil {
        m.compliancePolicy = value
    }
}
// SetDeviceConfiguration sets the deviceConfiguration property value. Whether device configuration is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetDeviceConfiguration(value *bool)() {
    if m != nil {
        m.deviceConfiguration = value
    }
}
// SetInventory sets the inventory property value. Whether inventory is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetInventory(value *bool)() {
    if m != nil {
        m.inventory = value
    }
}
// SetModernApps sets the modernApps property value. Whether modern application is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetModernApps(value *bool)() {
    if m != nil {
        m.modernApps = value
    }
}
// SetResourceAccess sets the resourceAccess property value. Whether resource access is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetResourceAccess(value *bool)() {
    if m != nil {
        m.resourceAccess = value
    }
}
// SetWindowsUpdateForBusiness sets the windowsUpdateForBusiness property value. Whether Windows Update for Business is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) SetWindowsUpdateForBusiness(value *bool)() {
    if m != nil {
        m.windowsUpdateForBusiness = value
    }
}
