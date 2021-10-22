package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Endpoint struct {
    DirectoryObject
    capability *string;
    providerId *string;
    providerName *string;
    providerResourceId *string;
    uri *string;
}
func NewEndpoint()(*Endpoint) {
    m := &Endpoint{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
func (m *Endpoint) GetCapability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capability
    }
}
func (m *Endpoint) GetProviderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerId
    }
}
func (m *Endpoint) GetProviderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerName
    }
}
func (m *Endpoint) GetProviderResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerResourceId
    }
}
func (m *Endpoint) GetUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uri
    }
}
func (m *Endpoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["capability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCapability(val)
        return nil
    }
    res["providerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProviderId(val)
        return nil
    }
    res["providerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProviderName(val)
        return nil
    }
    res["providerResourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProviderResourceId(val)
        return nil
    }
    res["uri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUri(val)
        return nil
    }
    return res
}
func (m *Endpoint) IsNil()(bool) {
    return m == nil
}
func (m *Endpoint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("capability", m.GetCapability())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerId", m.GetProviderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerName", m.GetProviderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerResourceId", m.GetProviderResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uri", m.GetUri())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Endpoint) SetCapability(value *string)() {
    m.capability = value
}
func (m *Endpoint) SetProviderId(value *string)() {
    m.providerId = value
}
func (m *Endpoint) SetProviderName(value *string)() {
    m.providerName = value
}
func (m *Endpoint) SetProviderResourceId(value *string)() {
    m.providerResourceId = value
}
func (m *Endpoint) SetUri(value *string)() {
    m.uri = value
}
