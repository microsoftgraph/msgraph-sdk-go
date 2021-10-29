package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceExchangeAccessStateSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Total count of devices with Exchange Access State: Allowed.
    allowedDeviceCount *int32;
    // Total count of devices with Exchange Access State: Blocked.
    blockedDeviceCount *int32;
    // Total count of devices with Exchange Access State: Quarantined.
    quarantinedDeviceCount *int32;
    // Total count of devices for which no Exchange Access State could be found.
    unavailableDeviceCount *int32;
    // Total count of devices with Exchange Access State: Unknown.
    unknownDeviceCount *int32;
}
// Instantiates a new deviceExchangeAccessStateSummary and sets the default values.
func NewDeviceExchangeAccessStateSummary()(*DeviceExchangeAccessStateSummary) {
    m := &DeviceExchangeAccessStateSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceExchangeAccessStateSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowedDeviceCount property value. Total count of devices with Exchange Access State: Allowed.
func (m *DeviceExchangeAccessStateSummary) GetAllowedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.allowedDeviceCount
    }
}
// Gets the blockedDeviceCount property value. Total count of devices with Exchange Access State: Blocked.
func (m *DeviceExchangeAccessStateSummary) GetBlockedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.blockedDeviceCount
    }
}
// Gets the quarantinedDeviceCount property value. Total count of devices with Exchange Access State: Quarantined.
func (m *DeviceExchangeAccessStateSummary) GetQuarantinedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.quarantinedDeviceCount
    }
}
// Gets the unavailableDeviceCount property value. Total count of devices for which no Exchange Access State could be found.
func (m *DeviceExchangeAccessStateSummary) GetUnavailableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unavailableDeviceCount
    }
}
// Gets the unknownDeviceCount property value. Total count of devices with Exchange Access State: Unknown.
func (m *DeviceExchangeAccessStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
// The deserialization information for the current model
func (m *DeviceExchangeAccessStateSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAllowedDeviceCount(val)
        return nil
    }
    res["blockedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBlockedDeviceCount(val)
        return nil
    }
    res["quarantinedDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetQuarantinedDeviceCount(val)
        return nil
    }
    res["unavailableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnavailableDeviceCount(val)
        return nil
    }
    res["unknownDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownDeviceCount(val)
        return nil
    }
    return res
}
func (m *DeviceExchangeAccessStateSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceExchangeAccessStateSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("allowedDeviceCount", m.GetAllowedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("blockedDeviceCount", m.GetBlockedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("quarantinedDeviceCount", m.GetQuarantinedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unavailableDeviceCount", m.GetUnavailableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
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
func (m *DeviceExchangeAccessStateSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowedDeviceCount property value. Total count of devices with Exchange Access State: Allowed.
// Parameters:
//  - value : Value to set for the allowedDeviceCount property.
func (m *DeviceExchangeAccessStateSummary) SetAllowedDeviceCount(value *int32)() {
    m.allowedDeviceCount = value
}
// Sets the blockedDeviceCount property value. Total count of devices with Exchange Access State: Blocked.
// Parameters:
//  - value : Value to set for the blockedDeviceCount property.
func (m *DeviceExchangeAccessStateSummary) SetBlockedDeviceCount(value *int32)() {
    m.blockedDeviceCount = value
}
// Sets the quarantinedDeviceCount property value. Total count of devices with Exchange Access State: Quarantined.
// Parameters:
//  - value : Value to set for the quarantinedDeviceCount property.
func (m *DeviceExchangeAccessStateSummary) SetQuarantinedDeviceCount(value *int32)() {
    m.quarantinedDeviceCount = value
}
// Sets the unavailableDeviceCount property value. Total count of devices for which no Exchange Access State could be found.
// Parameters:
//  - value : Value to set for the unavailableDeviceCount property.
func (m *DeviceExchangeAccessStateSummary) SetUnavailableDeviceCount(value *int32)() {
    m.unavailableDeviceCount = value
}
// Sets the unknownDeviceCount property value. Total count of devices with Exchange Access State: Unknown.
// Parameters:
//  - value : Value to set for the unknownDeviceCount property.
func (m *DeviceExchangeAccessStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
