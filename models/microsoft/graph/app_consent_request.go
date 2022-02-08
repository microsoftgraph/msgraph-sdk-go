package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppConsentRequest 
type AppConsentRequest struct {
    Entity
    // The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby.
    appDisplayName *string;
    // The identifier of the application. Required. Supports $filter (eq only) and $orderby.
    appId *string;
    // A list of pending scopes waiting for approval. Required.
    pendingScopes []AppConsentRequestScope;
    // A list of pending user consent requests.
    userConsentRequests []UserConsentRequest;
}
// NewAppConsentRequest instantiates a new appConsentRequest and sets the default values.
func NewAppConsentRequest()(*AppConsentRequest) {
    m := &AppConsentRequest{
        Entity: *NewEntity(),
    }
    return m
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
// GetPendingScopes gets the pendingScopes property value. A list of pending scopes waiting for approval. Required.
func (m *AppConsentRequest) GetPendingScopes()([]AppConsentRequestScope) {
    if m == nil {
        return nil
    } else {
        return m.pendingScopes
    }
}
// GetUserConsentRequests gets the userConsentRequests property value. A list of pending user consent requests.
func (m *AppConsentRequest) GetUserConsentRequests()([]UserConsentRequest) {
    if m == nil {
        return nil
    } else {
        return m.userConsentRequests
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppConsentRequestScope() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppConsentRequestScope, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppConsentRequestScope))
            }
            m.SetPendingScopes(res)
        }
        return nil
    }
    res["userConsentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserConsentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserConsentRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserConsentRequest))
            }
            m.SetUserConsentRequests(res)
        }
        return nil
    }
    return res
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
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("pendingScopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserConsentRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserConsentRequests()))
        for i, v := range m.GetUserConsentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *AppConsentRequest) SetPendingScopes(value []AppConsentRequestScope)() {
    if m != nil {
        m.pendingScopes = value
    }
}
// SetUserConsentRequests sets the userConsentRequests property value. A list of pending user consent requests.
func (m *AppConsentRequest) SetUserConsentRequests(value []UserConsentRequest)() {
    if m != nil {
        m.userConsentRequests = value
    }
}
