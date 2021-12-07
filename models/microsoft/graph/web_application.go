package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WebApplication 
type WebApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Home page or landing page of the application.
    homePageUrl *string;
    // Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
    implicitGrantSettings *ImplicitGrantSettings;
    // Specifies the URL that will be used by Microsoft's authorization service to logout an user using front-channel, back-channel or SAML logout protocols.
    logoutUrl *string;
    // Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
    redirectUris []string;
}
// NewWebApplication instantiates a new webApplication and sets the default values.
func NewWebApplication()(*WebApplication) {
    m := &WebApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHomePageUrl gets the homePageUrl property value. Home page or landing page of the application.
func (m *WebApplication) GetHomePageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homePageUrl
    }
}
// GetImplicitGrantSettings gets the implicitGrantSettings property value. Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
func (m *WebApplication) GetImplicitGrantSettings()(*ImplicitGrantSettings) {
    if m == nil {
        return nil
    } else {
        return m.implicitGrantSettings
    }
}
// GetLogoutUrl gets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using front-channel, back-channel or SAML logout protocols.
func (m *WebApplication) GetLogoutUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoutUrl
    }
}
// GetRedirectUris gets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *WebApplication) GetRedirectUris()([]string) {
    if m == nil {
        return nil
    } else {
        return m.redirectUris
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["homePageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomePageUrl(val)
        }
        return nil
    }
    res["implicitGrantSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImplicitGrantSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImplicitGrantSettings(val.(*ImplicitGrantSettings))
        }
        return nil
    }
    res["logoutUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoutUrl(val)
        }
        return nil
    }
    res["redirectUris"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRedirectUris(res)
        }
        return nil
    }
    return res
}
func (m *WebApplication) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WebApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("homePageUrl", m.GetHomePageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("implicitGrantSettings", m.GetImplicitGrantSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logoutUrl", m.GetLogoutUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("redirectUris", m.GetRedirectUris())
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
func (m *WebApplication) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHomePageUrl sets the homePageUrl property value. Home page or landing page of the application.
func (m *WebApplication) SetHomePageUrl(value *string)() {
    if m != nil {
        m.homePageUrl = value
    }
}
// SetImplicitGrantSettings sets the implicitGrantSettings property value. Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow.
func (m *WebApplication) SetImplicitGrantSettings(value *ImplicitGrantSettings)() {
    if m != nil {
        m.implicitGrantSettings = value
    }
}
// SetLogoutUrl sets the logoutUrl property value. Specifies the URL that will be used by Microsoft's authorization service to logout an user using front-channel, back-channel or SAML logout protocols.
func (m *WebApplication) SetLogoutUrl(value *string)() {
    if m != nil {
        m.logoutUrl = value
    }
}
// SetRedirectUris sets the redirectUris property value. Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent.
func (m *WebApplication) SetRedirectUris(value []string)() {
    if m != nil {
        m.redirectUris = value
    }
}
