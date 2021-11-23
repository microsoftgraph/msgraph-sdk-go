package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApiApplication 
type ApiApplication struct {
    // When true, allows an application to use claims mapping without specifying a custom signing key.
    acceptMappedClaims *bool;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
    knownClientApplications []string;
    // The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
    oauth2PermissionScopes []PermissionScope;
    // Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
    preAuthorizedApplications []PreAuthorizedApplication;
    // Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
    requestedAccessTokenVersion *int32;
}
// NewApiApplication instantiates a new apiApplication and sets the default values.
func NewApiApplication()(*ApiApplication) {
    m := &ApiApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAcceptMappedClaims gets the acceptMappedClaims property value. When true, allows an application to use claims mapping without specifying a custom signing key.
func (m *ApiApplication) GetAcceptMappedClaims()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acceptMappedClaims
    }
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetKnownClientApplications gets the knownClientApplications property value. Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
func (m *ApiApplication) GetKnownClientApplications()([]string) {
    if m == nil {
        return nil
    } else {
        return m.knownClientApplications
    }
}
// GetOauth2PermissionScopes gets the oauth2PermissionScopes property value. The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
func (m *ApiApplication) GetOauth2PermissionScopes()([]PermissionScope) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionScopes
    }
}
// GetPreAuthorizedApplications gets the preAuthorizedApplications property value. Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
func (m *ApiApplication) GetPreAuthorizedApplications()([]PreAuthorizedApplication) {
    if m == nil {
        return nil
    } else {
        return m.preAuthorizedApplications
    }
}
// GetRequestedAccessTokenVersion gets the requestedAccessTokenVersion property value. Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
func (m *ApiApplication) GetRequestedAccessTokenVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.requestedAccessTokenVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApiApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["acceptMappedClaims"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcceptMappedClaims(val)
        }
        return nil
    }
    res["knownClientApplications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKnownClientApplications(res)
        }
        return nil
    }
    res["oauth2PermissionScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PermissionScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*PermissionScope))
            }
            m.SetOauth2PermissionScopes(res)
        }
        return nil
    }
    res["preAuthorizedApplications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPreAuthorizedApplication() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PreAuthorizedApplication, len(val))
            for i, v := range val {
                res[i] = *(v.(*PreAuthorizedApplication))
            }
            m.SetPreAuthorizedApplications(res)
        }
        return nil
    }
    res["requestedAccessTokenVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedAccessTokenVersion(val)
        }
        return nil
    }
    return res
}
func (m *ApiApplication) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApiApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acceptMappedClaims", m.GetAcceptMappedClaims())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("knownClientApplications", m.GetKnownClientApplications())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOauth2PermissionScopes()))
        for i, v := range m.GetOauth2PermissionScopes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("oauth2PermissionScopes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPreAuthorizedApplications()))
        for i, v := range m.GetPreAuthorizedApplications() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("preAuthorizedApplications", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("requestedAccessTokenVersion", m.GetRequestedAccessTokenVersion())
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
// SetAcceptMappedClaims sets the acceptMappedClaims property value. When true, allows an application to use claims mapping without specifying a custom signing key.
func (m *ApiApplication) SetAcceptMappedClaims(value *bool)() {
    m.acceptMappedClaims = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKnownClientApplications sets the knownClientApplications property value. Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app. If you set the appID of the client app to this value, the user only consents once to the client app. Azure AD knows that consenting to the client means implicitly consenting to the web API and automatically provisions service principals for both APIs at the same time. Both the client and the web API app must be registered in the same tenant.
func (m *ApiApplication) SetKnownClientApplications(value []string)() {
    m.knownClientApplications = value
}
// SetOauth2PermissionScopes sets the oauth2PermissionScopes property value. The definition of the delegated permissions exposed by the web API represented by this application registration. These delegated permissions may be requested by a client application, and may be granted by users or administrators during consent. Delegated permissions are sometimes referred to as OAuth 2.0 scopes.
func (m *ApiApplication) SetOauth2PermissionScopes(value []PermissionScope)() {
    m.oauth2PermissionScopes = value
}
// SetPreAuthorizedApplications sets the preAuthorizedApplications property value. Lists the client applications that are pre-authorized with the specified delegated permissions to access this application's APIs. Users are not required to consent to any pre-authorized application (for the permissions specified). However, any additional permissions not listed in preAuthorizedApplications (requested through incremental consent for example) will require user consent.
func (m *ApiApplication) SetPreAuthorizedApplications(value []PreAuthorizedApplication)() {
    m.preAuthorizedApplications = value
}
// SetRequestedAccessTokenVersion sets the requestedAccessTokenVersion property value. Specifies the access token version expected by this resource. This changes the version and format of the JWT produced independent of the endpoint or client used to request the access token.  The endpoint used, v1.0 or v2.0, is chosen by the client and only impacts the version of id_tokens. Resources need to explicitly configure requestedAccessTokenVersion to indicate the supported access token format.  Possible values for requestedAccessTokenVersion are 1, 2, or null. If the value is null, this defaults to 1, which corresponds to the v1.0 endpoint.  If signInAudience on the application is configured as AzureADandPersonalMicrosoftAccount, the value for this property must be 2
func (m *ApiApplication) SetRequestedAccessTokenVersion(value *int32)() {
    m.requestedAccessTokenVersion = value
}
