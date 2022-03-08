package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppConsentRequest provides operations to manage the identityGovernance singleton.
type AppConsentRequest struct {
    Entity
    // The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
    appDisplayName *string;
    // The identifier of the application. Required. Supports $filter (eq only) and $orderby.
    appId *string;
    // A list of pending scopes waiting for approval. Required.
    pendingScopes []AppConsentRequestScopeable;
    // A list of pending user consent requests.
    userConsentRequests []UserConsentRequestable;
}
// NewAppConsentRequest instantiates a new appConsentRequest and sets the default values.
func NewAppConsentRequest()(*AppConsentRequest) {
    m := &AppConsentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppConsentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppConsentRequestFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAppConsentRequest(), nil
}
// GetAppDisplayName gets the appDisplayName property value. The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
func (m *AppConsentRequest) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
// GetAppId gets the appId property value. The identifier of the application. Required. Supports $filter (eq only) and $orderby.
func (m *AppConsentRequest) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppConsentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["pendingScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppConsentRequestScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppConsentRequestScopeable, len(val))
            for i, v := range val {
                res[i] = v.(AppConsentRequestScopeable)
            }
            m.SetPendingScopes(res)
        }
        return nil
    }
    res["userConsentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserConsentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserConsentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(UserConsentRequestable)
            }
            m.SetUserConsentRequests(res)
        }
        return nil
    }
    return res
}
// GetPendingScopes gets the pendingScopes property value. A list of pending scopes waiting for approval. Required.
func (m *AppConsentRequest) GetPendingScopes()([]AppConsentRequestScopeable) {
    if m == nil {
        return nil
    } else {
        return m.pendingScopes
    }
}
// GetUserConsentRequests gets the userConsentRequests property value. A list of pending user consent requests.
func (m *AppConsentRequest) GetUserConsentRequests()([]UserConsentRequestable) {
    if m == nil {
        return nil
    } else {
        return m.userConsentRequests
    }
}
func (m *AppConsentRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AppConsentRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetPendingScopes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPendingScopes()))
        for i, v := range m.GetPendingScopes() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("pendingScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserConsentRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserConsentRequests()))
        for i, v := range m.GetUserConsentRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userConsentRequests", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
func (m *AppConsentRequest) SetAppDisplayName(value *string)() {
    if m != nil {
        m.appDisplayName = value
    }
}
// SetAppId sets the appId property value. The identifier of the application. Required. Supports $filter (eq only) and $orderby.
func (m *AppConsentRequest) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetPendingScopes sets the pendingScopes property value. A list of pending scopes waiting for approval. Required.
func (m *AppConsentRequest) SetPendingScopes(value []AppConsentRequestScopeable)() {
    if m != nil {
        m.pendingScopes = value
    }
}
// SetUserConsentRequests sets the userConsentRequests property value. A list of pending user consent requests.
func (m *AppConsentRequest) SetUserConsentRequests(value []UserConsentRequestable)() {
    if m != nil {
        m.userConsentRequests = value
    }
}
