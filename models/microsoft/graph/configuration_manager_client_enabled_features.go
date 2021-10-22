package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConfigurationManagerClientEnabledFeatures struct {
    additionalData map[string]interface{};
    compliancePolicy *bool;
    deviceConfiguration *bool;
    inventory *bool;
    modernApps *bool;
    resourceAccess *bool;
    windowsUpdateForBusiness *bool;
}
func NewConfigurationManagerClientEnabledFeatures()(*ConfigurationManagerClientEnabledFeatures) {
    m := &ConfigurationManagerClientEnabledFeatures{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConfigurationManagerClientEnabledFeatures) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConfigurationManagerClientEnabledFeatures) GetCompliancePolicy()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicy
    }
}
func (m *ConfigurationManagerClientEnabledFeatures) GetDeviceConfiguration()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfiguration
    }
}
func (m *ConfigurationManagerClientEnabledFeatures) GetInventory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.inventory
    }
}
func (m *ConfigurationManagerClientEnabledFeatures) GetModernApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.modernApps
    }
}
func (m *ConfigurationManagerClientEnabledFeatures) GetResourceAccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccess
    }
}
func (m *ConfigurationManagerClientEnabledFeatures) GetWindowsUpdateForBusiness()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.windowsUpdateForBusiness
    }
}
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
func (m *ConfigurationManagerClientEnabledFeatures) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConfigurationManagerClientEnabledFeatures) SetCompliancePolicy(value *bool)() {
    m.compliancePolicy = value
}
func (m *ConfigurationManagerClientEnabledFeatures) SetDeviceConfiguration(value *bool)() {
    m.deviceConfiguration = value
}
func (m *ConfigurationManagerClientEnabledFeatures) SetInventory(value *bool)() {
    m.inventory = value
}
func (m *ConfigurationManagerClientEnabledFeatures) SetModernApps(value *bool)() {
    m.modernApps = value
}
func (m *ConfigurationManagerClientEnabledFeatures) SetResourceAccess(value *bool)() {
    m.resourceAccess = value
}
func (m *ConfigurationManagerClientEnabledFeatures) SetWindowsUpdateForBusiness(value *bool)() {
    m.windowsUpdateForBusiness = value
}
