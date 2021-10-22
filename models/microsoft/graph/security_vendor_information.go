package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SecurityVendorInformation struct {
    additionalData map[string]interface{};
    provider *string;
    providerVersion *string;
    subProvider *string;
    vendor *string;
}
func NewSecurityVendorInformation()(*SecurityVendorInformation) {
    m := &SecurityVendorInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SecurityVendorInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SecurityVendorInformation) GetProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
func (m *SecurityVendorInformation) GetProviderVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerVersion
    }
}
func (m *SecurityVendorInformation) GetSubProvider()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subProvider
    }
}
func (m *SecurityVendorInformation) GetVendor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vendor
    }
}
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
    res["vendor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVendor(val)
        return nil
    }
    return res
}
func (m *SecurityVendorInformation) IsNil()(bool) {
    return m == nil
}
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
        err := writer.WriteStringValue("vendor", m.GetVendor())
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
func (m *SecurityVendorInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SecurityVendorInformation) SetProvider(value *string)() {
    m.provider = value
}
func (m *SecurityVendorInformation) SetProviderVersion(value *string)() {
    m.providerVersion = value
}
func (m *SecurityVendorInformation) SetSubProvider(value *string)() {
    m.subProvider = value
}
func (m *SecurityVendorInformation) SetVendor(value *string)() {
    m.vendor = value
}
