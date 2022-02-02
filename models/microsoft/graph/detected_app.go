package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DetectedApp 
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
// NewDetectedApp instantiates a new detectedApp and sets the default values.
func NewDetectedApp()(*DetectedApp) {
    m := &DetectedApp{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeviceCount gets the deviceCount property value. The number of devices that have installed this application
func (m *DetectedApp) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// GetDisplayName gets the displayName property value. Name of the discovered application. Read-only
func (m *DetectedApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetManagedDevices gets the managedDevices property value. The devices that have the discovered application installed
func (m *DetectedApp) GetManagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
// GetSizeInByte gets the sizeInByte property value. Discovered application size in bytes. Read-only
func (m *DetectedApp) GetSizeInByte()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sizeInByte
    }
}
// GetVersion gets the version property value. Version of the discovered application. Read-only
func (m *DetectedApp) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
    if m.GetManagedDevices() != nil {
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
// SetDeviceCount sets the deviceCount property value. The number of devices that have installed this application
func (m *DetectedApp) SetDeviceCount(value *int32)() {
    if m != nil {
        m.deviceCount = value
    }
}
// SetDisplayName sets the displayName property value. Name of the discovered application. Read-only
func (m *DetectedApp) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetManagedDevices sets the managedDevices property value. The devices that have the discovered application installed
func (m *DetectedApp) SetManagedDevices(value []ManagedDevice)() {
    if m != nil {
        m.managedDevices = value
    }
}
// SetSizeInByte sets the sizeInByte property value. Discovered application size in bytes. Read-only
func (m *DetectedApp) SetSizeInByte(value *int64)() {
    if m != nil {
        m.sizeInByte = value
    }
}
// SetVersion sets the version property value. Version of the discovered application. Read-only
func (m *DetectedApp) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
