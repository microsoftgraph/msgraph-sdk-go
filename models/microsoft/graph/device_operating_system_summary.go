package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceOperatingSystemSummary provides operations to manage the deviceManagement singleton.
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
// NewDeviceOperatingSystemSummary instantiates a new deviceOperatingSystemSummary and sets the default values.
func NewDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    m := &DeviceOperatingSystemSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceOperatingSystemSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceOperatingSystemSummaryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceOperatingSystemSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceOperatingSystemSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAndroidCount gets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) GetAndroidCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceOperatingSystemSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["androidCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidCount(val)
        }
        return nil
    }
    res["iosCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosCount(val)
        }
        return nil
    }
    res["macOSCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOSCount(val)
        }
        return nil
    }
    res["unknownCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownCount(val)
        }
        return nil
    }
    res["windowsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsCount(val)
        }
        return nil
    }
    res["windowsMobileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsMobileCount(val)
        }
        return nil
    }
    return res
}
// GetIosCount gets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) GetIosCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iosCount
    }
}
// GetMacOSCount gets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) GetMacOSCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.macOSCount
    }
}
// GetUnknownCount gets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) GetUnknownCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unknownCount
    }
}
// GetWindowsCount gets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) GetWindowsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsCount
    }
}
// GetWindowsMobileCount gets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) GetWindowsMobileCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsMobileCount
    }
}
func (m *DeviceOperatingSystemSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceOperatingSystemSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAndroidCount sets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) SetAndroidCount(value *int32)() {
    if m != nil {
        m.androidCount = value
    }
}
// SetIosCount sets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) SetIosCount(value *int32)() {
    if m != nil {
        m.iosCount = value
    }
}
// SetMacOSCount sets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) SetMacOSCount(value *int32)() {
    if m != nil {
        m.macOSCount = value
    }
}
// SetUnknownCount sets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) SetUnknownCount(value *int32)() {
    if m != nil {
        m.unknownCount = value
    }
}
// SetWindowsCount sets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) SetWindowsCount(value *int32)() {
    if m != nil {
        m.windowsCount = value
    }
}
// SetWindowsMobileCount sets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) SetWindowsMobileCount(value *int32)() {
    if m != nil {
        m.windowsMobileCount = value
    }
}
