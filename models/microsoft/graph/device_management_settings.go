package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementSettings struct {
    additionalData map[string]interface{};
    deviceComplianceCheckinThresholdDays *int32;
    isScheduledActionEnabled *bool;
    secureByDefault *bool;
}
func NewDeviceManagementSettings()(*DeviceManagementSettings) {
    m := &DeviceManagementSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceManagementSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceCheckinThresholdDays
    }
}
func (m *DeviceManagementSettings) GetIsScheduledActionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isScheduledActionEnabled
    }
}
func (m *DeviceManagementSettings) GetSecureByDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.secureByDefault
    }
}
func (m *DeviceManagementSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceComplianceCheckinThresholdDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceComplianceCheckinThresholdDays(val)
        return nil
    }
    res["isScheduledActionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsScheduledActionEnabled(val)
        return nil
    }
    res["secureByDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSecureByDefault(val)
        return nil
    }
    return res
}
func (m *DeviceManagementSettings) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("deviceComplianceCheckinThresholdDays", m.GetDeviceComplianceCheckinThresholdDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isScheduledActionEnabled", m.GetIsScheduledActionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("secureByDefault", m.GetSecureByDefault())
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
func (m *DeviceManagementSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(value *int32)() {
    m.deviceComplianceCheckinThresholdDays = value
}
func (m *DeviceManagementSettings) SetIsScheduledActionEnabled(value *bool)() {
    m.isScheduledActionEnabled = value
}
func (m *DeviceManagementSettings) SetSecureByDefault(value *bool)() {
    m.secureByDefault = value
}
