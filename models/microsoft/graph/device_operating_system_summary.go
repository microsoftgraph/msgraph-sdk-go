package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceOperatingSystemSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of android device count.
    androidCount *int32;
    // Number of iOS device count.
    iosCount *int32;
    // Number of Mac OS X device count.
    macOSCount *int32;
    // Number of unknown device count.
    unknownCount *int32;
    // Number of Windows device count.
    windowsCount *int32;
    // Number of Windows mobile device count.
    windowsMobileCount *int32;
}
// Instantiates a new deviceOperatingSystemSummary and sets the default values.
func NewDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    m := &DeviceOperatingSystemSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceOperatingSystemSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) GetAndroidCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidCount
    }
}
// Gets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) GetIosCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iosCount
    }
}
// Gets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) GetMacOSCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.macOSCount
    }
}
// Gets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) GetUnknownCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownCount
    }
}
// Gets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) GetWindowsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsCount
    }
}
// Gets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) GetWindowsMobileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsMobileCount
    }
}
// The deserialization information for the current model
func (m *DeviceOperatingSystemSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["androidCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAndroidCount(val)
        return nil
    }
    res["iosCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIosCount(val)
        return nil
    }
    res["macOSCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMacOSCount(val)
        return nil
    }
    res["unknownCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUnknownCount(val)
        return nil
    }
    res["windowsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWindowsCount(val)
        return nil
    }
    res["windowsMobileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWindowsMobileCount(val)
        return nil
    }
    return res
}
func (m *DeviceOperatingSystemSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceOperatingSystemSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("androidCount", m.GetAndroidCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("iosCount", m.GetIosCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("macOSCount", m.GetMacOSCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownCount", m.GetUnknownCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsCount", m.GetWindowsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsMobileCount", m.GetWindowsMobileCount())
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
func (m *DeviceOperatingSystemSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the androidCount property value. Number of android device count.
// Parameters:
//  - value : Value to set for the androidCount property.
func (m *DeviceOperatingSystemSummary) SetAndroidCount(value *int32)() {
    m.androidCount = value
}
// Sets the iosCount property value. Number of iOS device count.
// Parameters:
//  - value : Value to set for the iosCount property.
func (m *DeviceOperatingSystemSummary) SetIosCount(value *int32)() {
    m.iosCount = value
}
// Sets the macOSCount property value. Number of Mac OS X device count.
// Parameters:
//  - value : Value to set for the macOSCount property.
func (m *DeviceOperatingSystemSummary) SetMacOSCount(value *int32)() {
    m.macOSCount = value
}
// Sets the unknownCount property value. Number of unknown device count.
// Parameters:
//  - value : Value to set for the unknownCount property.
func (m *DeviceOperatingSystemSummary) SetUnknownCount(value *int32)() {
    m.unknownCount = value
}
// Sets the windowsCount property value. Number of Windows device count.
// Parameters:
//  - value : Value to set for the windowsCount property.
func (m *DeviceOperatingSystemSummary) SetWindowsCount(value *int32)() {
    m.windowsCount = value
}
// Sets the windowsMobileCount property value. Number of Windows mobile device count.
// Parameters:
//  - value : Value to set for the windowsMobileCount property.
func (m *DeviceOperatingSystemSummary) SetWindowsMobileCount(value *int32)() {
    m.windowsMobileCount = value
}
