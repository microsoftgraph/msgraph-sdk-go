package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new endpoint and sets the default values.
func NewEndpoint()(*Endpoint) {
    m := &Endpoint{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the capability property value. Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.)  Not nullable. Read-only.
func (m *Endpoint) GetCapability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capability
    }
}
// Gets the providerId property value. Application id of the publishing underlying service. Not nullable. Read-only.
func (m *Endpoint) GetProviderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerId
    }
}
// Gets the providerName property value. Name of the publishing underlying service. Read-only.
func (m *Endpoint) GetProviderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerName
    }
}
// Gets the providerResourceId property value. For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only.
func (m *Endpoint) GetProviderResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerResourceId
    }
}
// Gets the uri property value. URL of the published resource. Not nullable. Read-only.
func (m *Endpoint) GetUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uri
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the capability property value. Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.)  Not nullable. Read-only.
// Parameters:
//  - value : Value to set for the capability property.
func (m *Endpoint) SetCapability(value *string)() {
    m.capability = value
}
// Sets the providerId property value. Application id of the publishing underlying service. Not nullable. Read-only.
// Parameters:
//  - value : Value to set for the providerId property.
func (m *Endpoint) SetProviderId(value *string)() {
    m.providerId = value
}
// Sets the providerName property value. Name of the publishing underlying service. Read-only.
// Parameters:
//  - value : Value to set for the providerName property.
func (m *Endpoint) SetProviderName(value *string)() {
    m.providerName = value
}
// Sets the providerResourceId property value. For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only.
// Parameters:
//  - value : Value to set for the providerResourceId property.
func (m *Endpoint) SetProviderResourceId(value *string)() {
    m.providerResourceId = value
}
// Sets the uri property value. URL of the published resource. Not nullable. Read-only.
// Parameters:
//  - value : Value to set for the uri property.
func (m *Endpoint) SetUri(value *string)() {
    m.uri = value
}
