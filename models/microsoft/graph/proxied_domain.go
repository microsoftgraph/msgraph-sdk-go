package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ProxiedDomain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The IP address or FQDN
    ipAddressOrFQDN *string;
    // Proxy IP or FQDN
    proxy *string;
}
// Instantiates a new proxiedDomain and sets the default values.
func NewProxiedDomain()(*ProxiedDomain) {
    m := &ProxiedDomain{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProxiedDomain) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the ipAddressOrFQDN property value. The IP address or FQDN
func (m *ProxiedDomain) GetIpAddressOrFQDN()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressOrFQDN
    }
}
// Gets the proxy property value. Proxy IP or FQDN
func (m *ProxiedDomain) GetProxy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.proxy
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ProxiedDomain) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the ipAddressOrFQDN property value. The IP address or FQDN
// Parameters:
//  - value : Value to set for the ipAddressOrFQDN property.
func (m *ProxiedDomain) SetIpAddressOrFQDN(value *string)() {
    m.ipAddressOrFQDN = value
}
// Sets the proxy property value. Proxy IP or FQDN
// Parameters:
//  - value : Value to set for the proxy property.
func (m *ProxiedDomain) SetProxy(value *string)() {
    m.proxy = value
}
