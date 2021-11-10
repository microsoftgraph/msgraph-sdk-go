package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DetectedApp struct {
    Entity
    // The number of devices that have installed this application
    deviceCount *int32;
    // Name of the discovered application. Read-only
    displayName *string;
    // The devices that have the discovered application installed
    managedDevices []ManagedDevice;
    // Discovered application size in bytes. Read-only
    sizeInByte *int64;
    // Version of the discovered application. Read-only
    version *string;
}
// Instantiates a new detectedApp and sets the default values.
func NewDetectedApp()(*DetectedApp) {
    m := &DetectedApp{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceCount property value. The number of devices that have installed this application
func (m *DetectedApp) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// Gets the displayName property value. Name of the discovered application. Read-only
func (m *DetectedApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the managedDevices property value. The devices that have the discovered application installed
func (m *DetectedApp) GetManagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
// Gets the sizeInByte property value. Discovered application size in bytes. Read-only
func (m *DetectedApp) GetSizeInByte()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sizeInByte
    }
}
// Gets the version property value. Version of the discovered application. Read-only
func (m *DetectedApp) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DetectedApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["managedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDevice, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDevice))
            }
            m.SetManagedDevices(res)
        }
        return nil
    }
    res["sizeInByte"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInByte(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *DetectedApp) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DetectedApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeInByte", m.GetSizeInByte())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceCount property value. The number of devices that have installed this application
// Parameters:
//  - value : Value to set for the deviceCount property.
func (m *DetectedApp) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
// Sets the displayName property value. Name of the discovered application. Read-only
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DetectedApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the managedDevices property value. The devices that have the discovered application installed
// Parameters:
//  - value : Value to set for the managedDevices property.
func (m *DetectedApp) SetManagedDevices(value []ManagedDevice)() {
    m.managedDevices = value
}
// Sets the sizeInByte property value. Discovered application size in bytes. Read-only
// Parameters:
//  - value : Value to set for the sizeInByte property.
func (m *DetectedApp) SetSizeInByte(value *int64)() {
    m.sizeInByte = value
}
// Sets the version property value. Version of the discovered application. Read-only
// Parameters:
//  - value : Value to set for the version property.
func (m *DetectedApp) SetVersion(value *string)() {
    m.version = value
}
