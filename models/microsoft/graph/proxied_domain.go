package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProxiedDomain struct {
    additionalData map[string]interface{};
    ipAddressOrFQDN *string;
    proxy *string;
}
func NewProxiedDomain()(*ProxiedDomain) {
    m := &ProxiedDomain{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ProxiedDomain) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ProxiedDomain) GetIpAddressOrFQDN()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressOrFQDN
    }
}
func (m *ProxiedDomain) GetProxy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proxy
    }
}
func (m *ProxiedDomain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ipAddressOrFQDN"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpAddressOrFQDN(val)
        return nil
    }
    res["proxy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProxy(val)
        return nil
    }
    return res
}
func (m *ProxiedDomain) IsNil()(bool) {
    return m == nil
}
func (m *ProxiedDomain) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("ipAddressOrFQDN", m.GetIpAddressOrFQDN())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proxy", m.GetProxy())
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
func (m *ProxiedDomain) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ProxiedDomain) SetIpAddressOrFQDN(value *string)() {
    m.ipAddressOrFQDN = value
}
func (m *ProxiedDomain) SetProxy(value *string)() {
    m.proxy = value
}
