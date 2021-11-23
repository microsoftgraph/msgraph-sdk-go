package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// endpoint 
type Endpoint struct {
    DirectoryObject
    // Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.)  Not nullable. Read-only.
    capability *string;
    // Application id of the publishing underlying service. Not nullable. Read-only.
    providerId *string;
    // Name of the publishing underlying service. Read-only.
    providerName *string;
    // For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only.
    providerResourceId *string;
    // URL of the published resource. Not nullable. Read-only.
    uri *string;
}
// NewEndpoint instantiates a new endpoint and sets the default values.
func NewEndpoint()(*Endpoint) {
    m := &Endpoint{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetCapability gets the capability property value. Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.)  Not nullable. Read-only.
func (m *Endpoint) GetCapability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capability
    }
}
// GetProviderId gets the providerId property value. Application id of the publishing underlying service. Not nullable. Read-only.
func (m *Endpoint) GetProviderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerId
    }
}
// GetProviderName gets the providerName property value. Name of the publishing underlying service. Read-only.
func (m *Endpoint) GetProviderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerName
    }
}
// GetProviderResourceId gets the providerResourceId property value. For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only.
func (m *Endpoint) GetProviderResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerResourceId
    }
}
// GetUri gets the uri property value. URL of the published resource. Not nullable. Read-only.
func (m *Endpoint) GetUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uri
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Endpoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["capability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapability(val)
        }
        return nil
    }
    res["providerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderId(val)
        }
        return nil
    }
    res["providerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderName(val)
        }
        return nil
    }
    res["providerResourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderResourceId(val)
        }
        return nil
    }
    res["uri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
func (m *Endpoint) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCapability sets the capability property value. Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.)  Not nullable. Read-only.
func (m *Endpoint) SetCapability(value *string)() {
    m.capability = value
}
// SetProviderId sets the providerId property value. Application id of the publishing underlying service. Not nullable. Read-only.
func (m *Endpoint) SetProviderId(value *string)() {
    m.providerId = value
}
// SetProviderName sets the providerName property value. Name of the publishing underlying service. Read-only.
func (m *Endpoint) SetProviderName(value *string)() {
    m.providerName = value
}
// SetProviderResourceId sets the providerResourceId property value. For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only.
func (m *Endpoint) SetProviderResourceId(value *string)() {
    m.providerResourceId = value
}
// SetUri sets the uri property value. URL of the published resource. Not nullable. Read-only.
func (m *Endpoint) SetUri(value *string)() {
    m.uri = value
}
