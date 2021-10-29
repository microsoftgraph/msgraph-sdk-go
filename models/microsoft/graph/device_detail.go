package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates the browser information of the used for signing in.
    browser *string;
    // Refers to the UniqueID of the device used for signing in.
    deviceId *string;
    // Refers to the name of the device used for signing in.
    displayName *string;
    // Indicates whether the device is compliant.
    isCompliant *bool;
    // Indicates whether the device is managed.
    isManaged *bool;
    // Indicates the operating system name and version used for signing in.
    operatingSystem *string;
    // Provides information about whether the signed-in device is Workplace Joined, AzureAD Joined, Domain Joined.
    trustType *string;
}
// Instantiates a new deviceDetail and sets the default values.
func NewDeviceDetail()(*DeviceDetail) {
    m := &DeviceDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the browser property value. Indicates the browser information of the used for signing in.
func (m *DeviceDetail) GetBrowser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.browser
    }
}
// Gets the deviceId property value. Refers to the UniqueID of the device used for signing in.
func (m *DeviceDetail) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the displayName property value. Refers to the name of the device used for signing in.
func (m *DeviceDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isCompliant property value. Indicates whether the device is compliant.
func (m *DeviceDetail) GetIsCompliant()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCompliant
    }
}
// Gets the isManaged property value. Indicates whether the device is managed.
func (m *DeviceDetail) GetIsManaged()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isManaged
    }
}
// Gets the operatingSystem property value. Indicates the operating system name and version used for signing in.
func (m *DeviceDetail) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// Gets the trustType property value. Provides information about whether the signed-in device is Workplace Joined, AzureAD Joined, Domain Joined.
func (m *DeviceDetail) GetTrustType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.trustType
    }
}
// The deserialization information for the current model
func (m *DeviceDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["browser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBrowser(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isCompliant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCompliant(val)
        return nil
    }
    res["isManaged"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsManaged(val)
        return nil
    }
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystem(val)
        return nil
    }
    res["trustType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTrustType(val)
        return nil
    }
    return res
}
func (m *DeviceDetail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("browser", m.GetBrowser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCompliant", m.GetIsCompliant())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isManaged", m.GetIsManaged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trustType", m.GetTrustType())
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
func (m *DeviceDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the browser property value. Indicates the browser information of the used for signing in.
// Parameters:
//  - value : Value to set for the browser property.
func (m *DeviceDetail) SetBrowser(value *string)() {
    m.browser = value
}
// Sets the deviceId property value. Refers to the UniqueID of the device used for signing in.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *DeviceDetail) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the displayName property value. Refers to the name of the device used for signing in.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isCompliant property value. Indicates whether the device is compliant.
// Parameters:
//  - value : Value to set for the isCompliant property.
func (m *DeviceDetail) SetIsCompliant(value *bool)() {
    m.isCompliant = value
}
// Sets the isManaged property value. Indicates whether the device is managed.
// Parameters:
//  - value : Value to set for the isManaged property.
func (m *DeviceDetail) SetIsManaged(value *bool)() {
    m.isManaged = value
}
// Sets the operatingSystem property value. Indicates the operating system name and version used for signing in.
// Parameters:
//  - value : Value to set for the operatingSystem property.
func (m *DeviceDetail) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// Sets the trustType property value. Provides information about whether the signed-in device is Workplace Joined, AzureAD Joined, Domain Joined.
// Parameters:
//  - value : Value to set for the trustType property.
func (m *DeviceDetail) SetTrustType(value *string)() {
    m.trustType = value
}
