package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Media struct {
    additionalData map[string]interface{};
    calleeDevice *DeviceInfo;
    calleeNetwork *NetworkInfo;
    callerDevice *DeviceInfo;
    callerNetwork *NetworkInfo;
    label *string;
    streams []MediaStream;
}
func NewMedia()(*Media) {
    m := &Media{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Media) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Media) GetCalleeDevice()(*DeviceInfo) {
    if m == nil {
        return nil
    } else {
        return m.calleeDevice
    }
}
func (m *Media) GetCalleeNetwork()(*NetworkInfo) {
    if m == nil {
        return nil
    } else {
        return m.calleeNetwork
    }
}
func (m *Media) GetCallerDevice()(*DeviceInfo) {
    if m == nil {
        return nil
    } else {
        return m.callerDevice
    }
}
func (m *Media) GetCallerNetwork()(*NetworkInfo) {
    if m == nil {
        return nil
    } else {
        return m.callerNetwork
    }
}
func (m *Media) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
func (m *Media) GetStreams()([]MediaStream) {
    if m == nil {
        return nil
    } else {
        return m.streams
    }
}
func (m *Media) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calleeDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceInfo() })
        if err != nil {
            return err
        }
        m.SetCalleeDevice(val.(*DeviceInfo))
        return nil
    }
    res["calleeNetwork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkInfo() })
        if err != nil {
            return err
        }
        m.SetCalleeNetwork(val.(*NetworkInfo))
        return nil
    }
    res["callerDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceInfo() })
        if err != nil {
            return err
        }
        m.SetCallerDevice(val.(*DeviceInfo))
        return nil
    }
    res["callerNetwork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkInfo() })
        if err != nil {
            return err
        }
        m.SetCallerNetwork(val.(*NetworkInfo))
        return nil
    }
    res["label"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLabel(val)
        return nil
    }
    res["streams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMediaStream() })
        if err != nil {
            return err
        }
        res := make([]MediaStream, len(val))
        for i, v := range val {
            res[i] = *(v.(*MediaStream))
        }
        m.SetStreams(res)
        return nil
    }
    return res
}
func (m *Media) IsNil()(bool) {
    return m == nil
}
func (m *Media) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("calleeDevice", m.GetCalleeDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("calleeNetwork", m.GetCalleeNetwork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("callerDevice", m.GetCallerDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("callerNetwork", m.GetCallerNetwork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStreams()))
        for i, v := range m.GetStreams() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("streams", cast)
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
func (m *Media) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Media) SetCalleeDevice(value *DeviceInfo)() {
    m.calleeDevice = value
}
func (m *Media) SetCalleeNetwork(value *NetworkInfo)() {
    m.calleeNetwork = value
}
func (m *Media) SetCallerDevice(value *DeviceInfo)() {
    m.callerDevice = value
}
func (m *Media) SetCallerNetwork(value *NetworkInfo)() {
    m.callerNetwork = value
}
func (m *Media) SetLabel(value *string)() {
    m.label = value
}
func (m *Media) SetStreams(value []MediaStream)() {
    m.streams = value
}
