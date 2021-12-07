package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProxiedDomain 
type ProxiedDomain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The IP address or FQDN
    ipAddressOrFQDN *string;
    // Proxy IP or FQDN
    proxy *string;
}
// NewProxiedDomain instantiates a new proxiedDomain and sets the default values.
func NewProxiedDomain()(*ProxiedDomain) {
    m := &ProxiedDomain{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProxiedDomain) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIpAddressOrFQDN gets the ipAddressOrFQDN property value. The IP address or FQDN
func (m *ProxiedDomain) GetIpAddressOrFQDN()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressOrFQDN
    }
}
// GetProxy gets the proxy property value. Proxy IP or FQDN
func (m *ProxiedDomain) GetProxy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proxy
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProxiedDomain) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ipAddressOrFQDN"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressOrFQDN(val)
        }
        return nil
    }
    res["proxy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProxy(val)
        }
        return nil
    }
    return res
}
func (m *ProxiedDomain) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProxiedDomain) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIpAddressOrFQDN sets the ipAddressOrFQDN property value. The IP address or FQDN
func (m *ProxiedDomain) SetIpAddressOrFQDN(value *string)() {
    if m != nil {
        m.ipAddressOrFQDN = value
    }
}
// SetProxy sets the proxy property value. Proxy IP or FQDN
func (m *ProxiedDomain) SetProxy(value *string)() {
    if m != nil {
        m.proxy = value
    }
}
