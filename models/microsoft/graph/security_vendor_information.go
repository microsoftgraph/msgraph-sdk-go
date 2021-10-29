package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SecurityVendorInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specific provider (product/service - not vendor company); for example, WindowsDefenderATP.
    provider *string;
    // Version of the provider or subprovider, if it exists, that generated the alert. Required
    providerVersion *string;
    // Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen.
    subProvider *string;
    // Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required
    vendor_escaped *string;
}
// Instantiates a new securityVendorInformation and sets the default values.
func NewSecurityVendorInformation()(*SecurityVendorInformation) {
    m := &SecurityVendorInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityVendorInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the provider property value. Specific provider (product/service - not vendor company); for example, WindowsDefenderATP.
func (m *SecurityVendorInformation) GetProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
// Gets the providerVersion property value. Version of the provider or subprovider, if it exists, that generated the alert. Required
func (m *SecurityVendorInformation) GetProviderVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerVersion
    }
}
// Gets the subProvider property value. Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen.
func (m *SecurityVendorInformation) GetSubProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subProvider
    }
}
// Gets the vendor_escaped property value. Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required
func (m *SecurityVendorInformation) GetVendor_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor_escaped
    }
}
// The deserialization information for the current model
func (m *SecurityVendorInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["provider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProvider(val)
        return nil
    }
    res["providerVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProviderVersion(val)
        return nil
    }
    res["subProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubProvider(val)
        return nil
    }
    res["vendor_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendor_escaped(val)
        return nil
    }
    return res
}
func (m *SecurityVendorInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SecurityVendorInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("providerVersion", m.GetProviderVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subProvider", m.GetSubProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendor_escaped", m.GetVendor_escaped())
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
func (m *SecurityVendorInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the provider property value. Specific provider (product/service - not vendor company); for example, WindowsDefenderATP.
// Parameters:
//  - value : Value to set for the provider property.
func (m *SecurityVendorInformation) SetProvider(value *string)() {
    m.provider = value
}
// Sets the providerVersion property value. Version of the provider or subprovider, if it exists, that generated the alert. Required
// Parameters:
//  - value : Value to set for the providerVersion property.
func (m *SecurityVendorInformation) SetProviderVersion(value *string)() {
    m.providerVersion = value
}
// Sets the subProvider property value. Specific subprovider (under aggregating provider); for example, WindowsDefenderATP.SmartScreen.
// Parameters:
//  - value : Value to set for the subProvider property.
func (m *SecurityVendorInformation) SetSubProvider(value *string)() {
    m.subProvider = value
}
// Sets the vendor_escaped property value. Name of the alert vendor (for example, Microsoft, Dell, FireEye). Required
// Parameters:
//  - value : Value to set for the vendor_escaped property.
func (m *SecurityVendorInformation) SetVendor_escaped(value *string)() {
    m.vendor_escaped = value
}
