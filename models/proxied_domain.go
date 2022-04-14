package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProxiedDomain proxied Domain
type ProxiedDomain struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The IP address or FQDN
    ipAddressOrFQDN *string
    // Proxy IP or FQDN
    proxy *string
}
// NewProxiedDomain instantiates a new proxiedDomain and sets the default values.
func NewProxiedDomain()(*ProxiedDomain) {
    m := &ProxiedDomain{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateProxiedDomainFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProxiedDomainFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProxiedDomain(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProxiedDomain) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProxiedDomain) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ipAddressOrFQDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressOrFQDN(val)
        }
        return nil
    }
    res["proxy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *ProxiedDomain) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
