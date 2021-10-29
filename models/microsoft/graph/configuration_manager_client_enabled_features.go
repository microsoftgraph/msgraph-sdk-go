package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new configurationManagerClientEnabledFeatures and sets the default values.
func NewConfigurationManagerClientEnabledFeatures()(*ConfigurationManagerClientEnabledFeatures) {
    m := &ConfigurationManagerClientEnabledFeatures{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientEnabledFeatures) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the compliancePolicy property value. Whether compliance policy is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetCompliancePolicy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicy
    }
}
// Gets the deviceConfiguration property value. Whether device configuration is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetDeviceConfiguration()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfiguration
    }
}
// Gets the inventory property value. Whether inventory is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetInventory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.inventory
    }
}
// Gets the modernApps property value. Whether modern application is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetModernApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.modernApps
    }
}
// Gets the resourceAccess property value. Whether resource access is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetResourceAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccess
    }
}
// Gets the windowsUpdateForBusiness property value. Whether Windows Update for Business is managed by Intune
func (m *ConfigurationManagerClientEnabledFeatures) GetWindowsUpdateForBusiness()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsUpdateForBusiness
    }
}
// The deserialization information for the current model
func (m *ConfigurationManagerClientEnabledFeatures) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compliancePolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCompliancePolicy(val)
        return nil
    }
    res["deviceConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDeviceConfiguration(val)
        return nil
    }
    res["inventory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetInventory(val)
        return nil
    }
    res["modernApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetModernApps(val)
        return nil
    }
    res["resourceAccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetResourceAccess(val)
        return nil
    }
    res["windowsUpdateForBusiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetWindowsUpdateForBusiness(val)
        return nil
    }
    return res
}
func (m *ConfigurationManagerClientEnabledFeatures) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ConfigurationManagerClientEnabledFeatures) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the compliancePolicy property value. Whether compliance policy is managed by Intune
// Parameters:
//  - value : Value to set for the compliancePolicy property.
func (m *ConfigurationManagerClientEnabledFeatures) SetCompliancePolicy(value *bool)() {
    m.compliancePolicy = value
}
// Sets the deviceConfiguration property value. Whether device configuration is managed by Intune
// Parameters:
//  - value : Value to set for the deviceConfiguration property.
func (m *ConfigurationManagerClientEnabledFeatures) SetDeviceConfiguration(value *bool)() {
    m.deviceConfiguration = value
}
// Sets the inventory property value. Whether inventory is managed by Intune
// Parameters:
//  - value : Value to set for the inventory property.
func (m *ConfigurationManagerClientEnabledFeatures) SetInventory(value *bool)() {
    m.inventory = value
}
// Sets the modernApps property value. Whether modern application is managed by Intune
// Parameters:
//  - value : Value to set for the modernApps property.
func (m *ConfigurationManagerClientEnabledFeatures) SetModernApps(value *bool)() {
    m.modernApps = value
}
// Sets the resourceAccess property value. Whether resource access is managed by Intune
// Parameters:
//  - value : Value to set for the resourceAccess property.
func (m *ConfigurationManagerClientEnabledFeatures) SetResourceAccess(value *bool)() {
    m.resourceAccess = value
}
// Sets the windowsUpdateForBusiness property value. Whether Windows Update for Business is managed by Intune
// Parameters:
//  - value : Value to set for the windowsUpdateForBusiness property.
func (m *ConfigurationManagerClientEnabledFeatures) SetWindowsUpdateForBusiness(value *bool)() {
    m.windowsUpdateForBusiness = value
}
