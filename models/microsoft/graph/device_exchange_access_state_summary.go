package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceExchangeAccessStateSummary struct {
    additionalData map[string]interface{};
    allowedDeviceCount *int32;
    blockedDeviceCount *int32;
    quarantinedDeviceCount *int32;
    unavailableDeviceCount *int32;
    unknownDeviceCount *int32;
}
func NewDeviceExchangeAccessStateSummary()(*DeviceExchangeAccessStateSummary) {
    m := &DeviceExchangeAccessStateSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceExchangeAccessStateSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceExchangeAccessStateSummary) GetAllowedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.allowedDeviceCount
    }
}
func (m *DeviceExchangeAccessStateSummary) GetBlockedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.blockedDeviceCount
    }
}
func (m *DeviceExchangeAccessStateSummary) GetQuarantinedDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.quarantinedDeviceCount
    }
}
func (m *DeviceExchangeAccessStateSummary) GetUnavailableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unavailableDeviceCount
    }
}
func (m *DeviceExchangeAccessStateSummary) GetUnknownDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownDeviceCount
    }
}
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
func (m *DeviceExchangeAccessStateSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceExchangeAccessStateSummary) SetAllowedDeviceCount(value *int32)() {
    m.allowedDeviceCount = value
}
func (m *DeviceExchangeAccessStateSummary) SetBlockedDeviceCount(value *int32)() {
    m.blockedDeviceCount = value
}
func (m *DeviceExchangeAccessStateSummary) SetQuarantinedDeviceCount(value *int32)() {
    m.quarantinedDeviceCount = value
}
func (m *DeviceExchangeAccessStateSummary) SetUnavailableDeviceCount(value *int32)() {
    m.unavailableDeviceCount = value
}
func (m *DeviceExchangeAccessStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
