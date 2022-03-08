package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityProvider provides operations to manage the identityContainer singleton.
type IdentityProvider struct {
    Entity
    // The client ID for the application. This is the client ID obtained when registering the application with the identity provider. Required. Not nullable.
    clientId *string;
    // The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.  Required. Not nullable.
    clientSecret *string;
    // The display name of the identity provider. Not nullable.
    name *string;
    // The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect. Not nullable.
    type_escaped *string;
}
// NewIdentityProvider instantiates a new identityProvider and sets the default values.
func NewIdentityProvider()(*IdentityProvider) {
    m := &IdentityProvider{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIdentityProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityProviderFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewIdentityProvider(), nil
}
// GetClientId gets the clientId property value. The client ID for the application. This is the client ID obtained when registering the application with the identity provider. Required. Not nullable.
func (m *IdentityProvider) GetClientId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientId
    }
}
// GetClientSecret gets the clientSecret property value. The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.  Required. Not nullable.
func (m *IdentityProvider) GetClientSecret()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientSecret
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityProvider) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The display name of the identity provider. Not nullable.
func (m *IdentityProvider) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetType gets the type property value. The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect. Not nullable.
func (m *IdentityProvider) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *IdentityProvider) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IdentityProvider) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientId sets the clientId property value. The client ID for the application. This is the client ID obtained when registering the application with the identity provider. Required. Not nullable.
func (m *IdentityProvider) SetClientId(value *string)() {
    if m != nil {
        m.clientId = value
    }
}
// SetClientSecret sets the clientSecret property value. The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.  Required. Not nullable.
func (m *IdentityProvider) SetClientSecret(value *string)() {
    if m != nil {
        m.clientSecret = value
    }
}
// SetName sets the name property value. The display name of the identity provider. Not nullable.
func (m *IdentityProvider) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetType sets the type property value. The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect. Not nullable.
func (m *IdentityProvider) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
