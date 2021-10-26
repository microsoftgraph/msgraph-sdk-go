package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new media and sets the default values.
func NewMedia()(*Media) {
    m := &Media{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Media) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the calleeDevice property value. Device information associated with the callee endpoint of this media.
func (m *Media) GetCalleeDevice()(*DeviceInfo) {
    if m == nil {
        return nil
    } else {
        return m.calleeDevice
    }
}
// Gets the calleeNetwork property value. Network information associated with the callee endpoint of this media.
func (m *Media) GetCalleeNetwork()(*NetworkInfo) {
    if m == nil {
        return nil
    } else {
        return m.calleeNetwork
    }
}
// Gets the callerDevice property value. Device information associated with the caller endpoint of this media.
func (m *Media) GetCallerDevice()(*DeviceInfo) {
    if m == nil {
        return nil
    } else {
        return m.callerDevice
    }
}
// Gets the callerNetwork property value. Network information associated with the caller endpoint of this media.
func (m *Media) GetCallerNetwork()(*NetworkInfo) {
    if m == nil {
        return nil
    } else {
        return m.callerNetwork
    }
}
// Gets the label property value. How the media was identified during media negotiation stage.
func (m *Media) GetLabel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.label
    }
}
// Gets the streams property value. Network streams associated with this media.
func (m *Media) GetStreams()([]MediaStream) {
    if m == nil {
        return nil
    } else {
        return m.streams
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Media) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the calleeDevice property value. Device information associated with the callee endpoint of this media.
// Parameters:
//  - value : Value to set for the calleeDevice property.
func (m *Media) SetCalleeDevice(value *DeviceInfo)() {
    m.calleeDevice = value
}
// Sets the calleeNetwork property value. Network information associated with the callee endpoint of this media.
// Parameters:
//  - value : Value to set for the calleeNetwork property.
func (m *Media) SetCalleeNetwork(value *NetworkInfo)() {
    m.calleeNetwork = value
}
// Sets the callerDevice property value. Device information associated with the caller endpoint of this media.
// Parameters:
//  - value : Value to set for the callerDevice property.
func (m *Media) SetCallerDevice(value *DeviceInfo)() {
    m.callerDevice = value
}
// Sets the callerNetwork property value. Network information associated with the caller endpoint of this media.
// Parameters:
//  - value : Value to set for the callerNetwork property.
func (m *Media) SetCallerNetwork(value *NetworkInfo)() {
    m.callerNetwork = value
}
// Sets the label property value. How the media was identified during media negotiation stage.
// Parameters:
//  - value : Value to set for the label property.
func (m *Media) SetLabel(value *string)() {
    m.label = value
}
// Sets the streams property value. Network streams associated with this media.
// Parameters:
//  - value : Value to set for the streams property.
func (m *Media) SetStreams(value []MediaStream)() {
    m.streams = value
}
