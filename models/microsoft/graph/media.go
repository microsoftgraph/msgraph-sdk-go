package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Media 
type Media struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Device information associated with the callee endpoint of this media.
    calleeDevice *DeviceInfo;
    // Network information associated with the callee endpoint of this media.
    calleeNetwork *NetworkInfo;
    // Device information associated with the caller endpoint of this media.
    callerDevice *DeviceInfo;
    // Network information associated with the caller endpoint of this media.
    callerNetwork *NetworkInfo;
    // How the media was identified during media negotiation stage.
    label *string;
    // Network streams associated with this media.
    streams []MediaStream;
}
// NewMedia instantiates a new media and sets the default values.
func NewMedia()(*Media) {
    m := &Media{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Media) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCalleeDevice gets the calleeDevice property value. Device information associated with the callee endpoint of this media.
func (m *Media) GetCalleeDevice()(*DeviceInfo) {
    if m == nil {
        return nil
    } else {
        return m.calleeDevice
    }
}
// GetCalleeNetwork gets the calleeNetwork property value. Network information associated with the callee endpoint of this media.
func (m *Media) GetCalleeNetwork()(*NetworkInfo) {
    if m == nil {
        return nil
    } else {
        return m.calleeNetwork
    }
}
// GetCallerDevice gets the callerDevice property value. Device information associated with the caller endpoint of this media.
func (m *Media) GetCallerDevice()(*DeviceInfo) {
    if m == nil {
        return nil
    } else {
        return m.callerDevice
    }
}
// GetCallerNetwork gets the callerNetwork property value. Network information associated with the caller endpoint of this media.
func (m *Media) GetCallerNetwork()(*NetworkInfo) {
    if m == nil {
        return nil
    } else {
        return m.callerNetwork
    }
}
// GetLabel gets the label property value. How the media was identified during media negotiation stage.
func (m *Media) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// GetStreams gets the streams property value. Network streams associated with this media.
func (m *Media) GetStreams()([]MediaStream) {
    if m == nil {
        return nil
    } else {
        return m.streams
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Media) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calleeDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalleeDevice(val.(*DeviceInfo))
        }
        return nil
    }
    res["calleeNetwork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalleeNetwork(val.(*NetworkInfo))
        }
        return nil
    }
    res["callerDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallerDevice(val.(*DeviceInfo))
        }
        return nil
    }
    res["callerNetwork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNetworkInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallerNetwork(val.(*NetworkInfo))
        }
        return nil
    }
    res["label"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["streams"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMediaStream() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MediaStream, len(val))
            for i, v := range val {
                res[i] = *(v.(*MediaStream))
            }
            m.SetStreams(res)
        }
        return nil
    }
    return res
}
func (m *Media) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Media) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCalleeDevice sets the calleeDevice property value. Device information associated with the callee endpoint of this media.
func (m *Media) SetCalleeDevice(value *DeviceInfo)() {
    if m != nil {
        m.calleeDevice = value
    }
}
// SetCalleeNetwork sets the calleeNetwork property value. Network information associated with the callee endpoint of this media.
func (m *Media) SetCalleeNetwork(value *NetworkInfo)() {
    if m != nil {
        m.calleeNetwork = value
    }
}
// SetCallerDevice sets the callerDevice property value. Device information associated with the caller endpoint of this media.
func (m *Media) SetCallerDevice(value *DeviceInfo)() {
    if m != nil {
        m.callerDevice = value
    }
}
// SetCallerNetwork sets the callerNetwork property value. Network information associated with the caller endpoint of this media.
func (m *Media) SetCallerNetwork(value *NetworkInfo)() {
    if m != nil {
        m.callerNetwork = value
    }
}
// SetLabel sets the label property value. How the media was identified during media negotiation stage.
func (m *Media) SetLabel(value *string)() {
    if m != nil {
        m.label = value
    }
}
// SetStreams sets the streams property value. Network streams associated with this media.
func (m *Media) SetStreams(value []MediaStream)() {
    if m != nil {
        m.streams = value
    }
}
