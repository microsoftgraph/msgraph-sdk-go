package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WindowsInformationProtectionProxiedDomainCollection struct {
    additionalData map[string]interface{};
    displayName *string;
    proxiedDomains []ProxiedDomain;
}
func NewWindowsInformationProtectionProxiedDomainCollection()(*WindowsInformationProtectionProxiedDomainCollection) {
    m := &WindowsInformationProtectionProxiedDomainCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WindowsInformationProtectionProxiedDomainCollection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WindowsInformationProtectionProxiedDomainCollection) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *WindowsInformationProtectionProxiedDomainCollection) GetProxiedDomains()([]ProxiedDomain) {
    if m == nil {
        return nil
    } else {
        return m.proxiedDomains
    }
}
func (m *WindowsInformationProtectionProxiedDomainCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["proxiedDomains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProxiedDomain() })
        if err != nil {
            return err
        }
        res := make([]ProxiedDomain, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProxiedDomain))
        }
        m.SetProxiedDomains(res)
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionProxiedDomainCollection) IsNil()(bool) {
    return m == nil
}
func (m *WindowsInformationProtectionProxiedDomainCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProxiedDomains()))
        for i, v := range m.GetProxiedDomains() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("proxiedDomains", cast)
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
func (m *WindowsInformationProtectionProxiedDomainCollection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WindowsInformationProtectionProxiedDomainCollection) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *WindowsInformationProtectionProxiedDomainCollection) SetProxiedDomains(value []ProxiedDomain)() {
    m.proxiedDomains = value
}
