package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IdentityApiConnector struct {
    Entity
    // The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client certificate are supported.
    authenticationConfiguration *ApiAuthenticationConfigurationBase;
    // The name of the API connector.
    displayName *string;
    // The URL of the API endpoint to call.
    targetUrl *string;
}
// Instantiates a new identityApiConnector and sets the default values.
func NewIdentityApiConnector()(*IdentityApiConnector) {
    m := &IdentityApiConnector{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the authenticationConfiguration property value. The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client certificate are supported.
func (m *IdentityApiConnector) GetAuthenticationConfiguration()(*ApiAuthenticationConfigurationBase) {
    if m == nil {
        return nil
    } else {
        return m.authenticationConfiguration
    }
}
// Gets the displayName property value. The name of the API connector.
func (m *IdentityApiConnector) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the targetUrl property value. The URL of the API endpoint to call.
func (m *IdentityApiConnector) GetTargetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetUrl
    }
}
// The deserialization information for the current model
func (m *IdentityApiConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApiAuthenticationConfigurationBase() })
        if err != nil {
            return err
        }
        m.SetAuthenticationConfiguration(val.(*ApiAuthenticationConfigurationBase))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["targetUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetUrl(val)
        return nil
    }
    return res
}
func (m *IdentityApiConnector) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IdentityApiConnector) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authenticationConfiguration", m.GetAuthenticationConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetUrl", m.GetTargetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the authenticationConfiguration property value. The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client certificate are supported.
// Parameters:
//  - value : Value to set for the authenticationConfiguration property.
func (m *IdentityApiConnector) SetAuthenticationConfiguration(value *ApiAuthenticationConfigurationBase)() {
    m.authenticationConfiguration = value
}
// Sets the displayName property value. The name of the API connector.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *IdentityApiConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the targetUrl property value. The URL of the API endpoint to call.
// Parameters:
//  - value : Value to set for the targetUrl property.
func (m *IdentityApiConnector) SetTargetUrl(value *string)() {
    m.targetUrl = value
}
