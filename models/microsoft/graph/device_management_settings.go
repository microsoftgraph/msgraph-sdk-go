package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettings 
type DeviceManagementSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of days a device is allowed to go without checking in to remain compliant.
    deviceComplianceCheckinThresholdDays *int32;
    // Is feature enabled or not for scheduled action for rule.
    isScheduledActionEnabled *bool;
    // Device should be noncompliant when there is no compliance policy targeted when this is true
    secureByDefault *bool;
}
// NewDeviceManagementSettings instantiates a new deviceManagementSettings and sets the default values.
func NewDeviceManagementSettings()(*DeviceManagementSettings) {
    m := &DeviceManagementSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceComplianceCheckinThresholdDays gets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) GetDeviceComplianceCheckinThresholdDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceCheckinThresholdDays
    }
}
// GetIsScheduledActionEnabled gets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) GetIsScheduledActionEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isScheduledActionEnabled
    }
}
// GetSecureByDefault gets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) GetSecureByDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.secureByDefault
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceComplianceCheckinThresholdDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceCheckinThresholdDays(val)
        }
        return nil
    }
    res["isScheduledActionEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScheduledActionEnabled(val)
        }
        return nil
    }
    res["secureByDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureByDefault(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceComplianceCheckinThresholdDays sets the deviceComplianceCheckinThresholdDays property value. The number of days a device is allowed to go without checking in to remain compliant.
func (m *DeviceManagementSettings) SetDeviceComplianceCheckinThresholdDays(value *int32)() {
    if m != nil {
        m.deviceComplianceCheckinThresholdDays = value
    }
}
// SetIsScheduledActionEnabled sets the isScheduledActionEnabled property value. Is feature enabled or not for scheduled action for rule.
func (m *DeviceManagementSettings) SetIsScheduledActionEnabled(value *bool)() {
    if m != nil {
        m.isScheduledActionEnabled = value
    }
}
// SetSecureByDefault sets the secureByDefault property value. Device should be noncompliant when there is no compliance policy targeted when this is true
func (m *DeviceManagementSettings) SetSecureByDefault(value *bool)() {
    if m != nil {
        m.secureByDefault = value
    }
}
