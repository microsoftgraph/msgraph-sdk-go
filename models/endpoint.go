package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Endpoint 
type Endpoint struct {
    DirectoryObject
    // Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.) Not nullable. Read-only.
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
// CreateEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEndpoint(), nil
}
// GetCapability gets the capability property value. Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.) Not nullable. Read-only.
func (m *Endpoint) GetCapability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.capability
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Endpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["capability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapability(val)
        }
        return nil
    }
    res["providerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderId(val)
        }
        return nil
    }
    res["providerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderName(val)
        }
        return nil
    }
    res["providerResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderResourceId(val)
        }
        return nil
    }
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *Endpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetCapability sets the capability property value. Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.) Not nullable. Read-only.
func (m *Endpoint) SetCapability(value *string)() {
    if m != nil {
        m.capability = value
    }
}
// SetProviderId sets the providerId property value. Application id of the publishing underlying service. Not nullable. Read-only.
func (m *Endpoint) SetProviderId(value *string)() {
    if m != nil {
        m.providerId = value
    }
}
// SetProviderName sets the providerName property value. Name of the publishing underlying service. Read-only.
func (m *Endpoint) SetProviderName(value *string)() {
    if m != nil {
        m.providerName = value
    }
}
// SetProviderResourceId sets the providerResourceId property value. For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only.
func (m *Endpoint) SetProviderResourceId(value *string)() {
    if m != nil {
        m.providerResourceId = value
    }
}
// SetUri sets the uri property value. URL of the published resource. Not nullable. Read-only.
func (m *Endpoint) SetUri(value *string)() {
    if m != nil {
        m.uri = value
    }
}
