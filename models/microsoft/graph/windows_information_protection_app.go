package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WindowsInformationProtectionApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If true, app is denied protection or exemption.
    denied *bool;
    // The app's description.
    description *string;
    // App display name.
    displayName *string;
    // The product name.
    productName *string;
    // The publisher name
    publisherName *string;
}
// Instantiates a new windowsInformationProtectionApp and sets the default values.
func NewWindowsInformationProtectionApp()(*WindowsInformationProtectionApp) {
    m := &WindowsInformationProtectionApp{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the denied property value. If true, app is denied protection or exemption.
func (m *WindowsInformationProtectionApp) GetDenied()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.denied
    }
}
// Gets the description property value. The app's description.
func (m *WindowsInformationProtectionApp) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. App display name.
func (m *WindowsInformationProtectionApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the productName property value. The product name.
func (m *WindowsInformationProtectionApp) GetProductName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productName
    }
}
// Gets the publisherName property value. The publisher name
func (m *WindowsInformationProtectionApp) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// The deserialization information for the current model
func (m *WindowsInformationProtectionApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["denied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDenied(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["productName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["publisherName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionApp) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WindowsInformationProtectionApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("denied", m.GetDenied())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisherName", m.GetPublisherName())
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
func (m *WindowsInformationProtectionApp) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the denied property value. If true, app is denied protection or exemption.
// Parameters:
//  - value : Value to set for the denied property.
func (m *WindowsInformationProtectionApp) SetDenied(value *bool)() {
    m.denied = value
}
// Sets the description property value. The app's description.
// Parameters:
//  - value : Value to set for the description property.
func (m *WindowsInformationProtectionApp) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. App display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *WindowsInformationProtectionApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the productName property value. The product name.
// Parameters:
//  - value : Value to set for the productName property.
func (m *WindowsInformationProtectionApp) SetProductName(value *string)() {
    m.productName = value
}
// Sets the publisherName property value. The publisher name
// Parameters:
//  - value : Value to set for the publisherName property.
func (m *WindowsInformationProtectionApp) SetPublisherName(value *string)() {
    m.publisherName = value
}
