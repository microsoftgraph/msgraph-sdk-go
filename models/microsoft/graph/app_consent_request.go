package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AppConsentRequest struct {
    Entity
    appDisplayName *string;
    appId *string;
    pendingScopes []AppConsentRequestScope;
    userConsentRequests []UserConsentRequest;
}
func NewAppConsentRequest()(*AppConsentRequest) {
    m := &AppConsentRequest{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AppConsentRequest) GetAppDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appDisplayName
    }
}
func (m *AppConsentRequest) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
func (m *AppConsentRequest) GetPendingScopes()([]AppConsentRequestScope) {
    if m == nil {
        return nil
    } else {
        return m.pendingScopes
    }
}
func (m *AppConsentRequest) GetUserConsentRequests()([]UserConsentRequest) {
    if m == nil {
        return nil
    } else {
        return m.userConsentRequests
    }
}
func (m *AppConsentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppDisplayName(val)
        return nil
    }
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["pendingScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppConsentRequestScope() })
        if err != nil {
            return err
        }
        res := make([]AppConsentRequestScope, len(val))
        for i, v := range val {
            res[i] = *(v.(*AppConsentRequestScope))
        }
        m.SetPendingScopes(res)
        return nil
    }
    res["userConsentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserConsentRequest() })
        if err != nil {
            return err
        }
        res := make([]UserConsentRequest, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserConsentRequest))
        }
        m.SetUserConsentRequests(res)
        return nil
    }
    return res
}
func (m *AppConsentRequest) IsNil()(bool) {
    return m == nil
}
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
    {
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
    {
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
func (m *AppConsentRequest) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
func (m *AppConsentRequest) SetAppId(value *string)() {
    m.appId = value
}
func (m *AppConsentRequest) SetPendingScopes(value []AppConsentRequestScope)() {
    m.pendingScopes = value
}
func (m *AppConsentRequest) SetUserConsentRequests(value []UserConsentRequest)() {
    m.userConsentRequests = value
}
