package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OAuth2PermissionGrant provides operations to manage the admin singleton.
type OAuth2PermissionGrant struct {
    Entity
    // The id of the client service principal for the application which is authorized to act on behalf of a signed-in user when accessing an API. Required. Supports $filter (eq only).
    clientId *string
    // Indicates whether authorization is granted for the client application to impersonate all users or only a specific user. AllPrincipals indicates authorization to impersonate all users. Principal indicates authorization to impersonate a specific user. Consent on behalf of all users can be granted by an administrator. Non-admin users may be authorized to consent on behalf of themselves in some cases, for some delegated permissions. Required. Supports $filter (eq only).
    consentType *string
    // The id of the user on behalf of whom the client is authorized to access the resource, when consentType is Principal. If consentType is AllPrincipals this value is null. Required when consentType is Principal. Supports $filter (eq only).
    principalId *string
    // The id of the resource service principal to which access is authorized. This identifies the API which the client is authorized to attempt to call on behalf of a signed-in user. Supports $filter (eq only).
    resourceId *string
    // A space-separated list of the claim values for delegated permissions which should be included in access tokens for the resource application (the API). For example, openid User.Read GroupMember.Read.All. Each claim value should match the value field of one of the delegated permissions defined by the API, listed in the publishedPermissionScopes property of the resource service principal.
    scope *string
}
// NewOAuth2PermissionGrant instantiates a new oAuth2PermissionGrant and sets the default values.
func NewOAuth2PermissionGrant()(*OAuth2PermissionGrant) {
    m := &OAuth2PermissionGrant{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOAuth2PermissionGrantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOAuth2PermissionGrantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOAuth2PermissionGrant(), nil
}
// GetClientId gets the clientId property value. The id of the client service principal for the application which is authorized to act on behalf of a signed-in user when accessing an API. Required. Supports $filter (eq only).
func (m *OAuth2PermissionGrant) GetClientId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientId
    }
}
// GetConsentType gets the consentType property value. Indicates whether authorization is granted for the client application to impersonate all users or only a specific user. AllPrincipals indicates authorization to impersonate all users. Principal indicates authorization to impersonate a specific user. Consent on behalf of all users can be granted by an administrator. Non-admin users may be authorized to consent on behalf of themselves in some cases, for some delegated permissions. Required. Supports $filter (eq only).
func (m *OAuth2PermissionGrant) GetConsentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.consentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OAuth2PermissionGrant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["consentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsentType(val)
        }
        return nil
    }
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    return res
}
// GetPrincipalId gets the principalId property value. The id of the user on behalf of whom the client is authorized to access the resource, when consentType is Principal. If consentType is AllPrincipals this value is null. Required when consentType is Principal. Supports $filter (eq only).
func (m *OAuth2PermissionGrant) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
// GetResourceId gets the resourceId property value. The id of the resource service principal to which access is authorized. This identifies the API which the client is authorized to attempt to call on behalf of a signed-in user. Supports $filter (eq only).
func (m *OAuth2PermissionGrant) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// GetScope gets the scope property value. A space-separated list of the claim values for delegated permissions which should be included in access tokens for the resource application (the API). For example, openid User.Read GroupMember.Read.All. Each claim value should match the value field of one of the delegated permissions defined by the API, listed in the publishedPermissionScopes property of the resource service principal.
func (m *OAuth2PermissionGrant) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Serialize serializes information the current object
func (m *OAuth2PermissionGrant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("consentType", m.GetConsentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientId sets the clientId property value. The id of the client service principal for the application which is authorized to act on behalf of a signed-in user when accessing an API. Required. Supports $filter (eq only).
func (m *OAuth2PermissionGrant) SetClientId(value *string)() {
    if m != nil {
        m.clientId = value
    }
}
// SetConsentType sets the consentType property value. Indicates whether authorization is granted for the client application to impersonate all users or only a specific user. AllPrincipals indicates authorization to impersonate all users. Principal indicates authorization to impersonate a specific user. Consent on behalf of all users can be granted by an administrator. Non-admin users may be authorized to consent on behalf of themselves in some cases, for some delegated permissions. Required. Supports $filter (eq only).
func (m *OAuth2PermissionGrant) SetConsentType(value *string)() {
    if m != nil {
        m.consentType = value
    }
}
// SetPrincipalId sets the principalId property value. The id of the user on behalf of whom the client is authorized to access the resource, when consentType is Principal. If consentType is AllPrincipals this value is null. Required when consentType is Principal. Supports $filter (eq only).
func (m *OAuth2PermissionGrant) SetPrincipalId(value *string)() {
    if m != nil {
        m.principalId = value
    }
}
// SetResourceId sets the resourceId property value. The id of the resource service principal to which access is authorized. This identifies the API which the client is authorized to attempt to call on behalf of a signed-in user. Supports $filter (eq only).
func (m *OAuth2PermissionGrant) SetResourceId(value *string)() {
    if m != nil {
        m.resourceId = value
    }
}
// SetScope sets the scope property value. A space-separated list of the claim values for delegated permissions which should be included in access tokens for the resource application (the API). For example, openid User.Read GroupMember.Read.All. Each claim value should match the value field of one of the delegated permissions defined by the API, listed in the publishedPermissionScopes property of the resource service principal.
func (m *OAuth2PermissionGrant) SetScope(value *string)() {
    if m != nil {
        m.scope = value
    }
}
