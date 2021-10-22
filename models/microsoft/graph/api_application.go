package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApiApplication struct {
    acceptMappedClaims *bool;
    additionalData map[string]interface{};
    knownClientApplications []string;
    oauth2PermissionScopes []PermissionScope;
    preAuthorizedApplications []PreAuthorizedApplication;
    requestedAccessTokenVersion *int32;
}
func NewApiApplication()(*ApiApplication) {
    m := &ApiApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApiApplication) GetAcceptMappedClaims()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acceptMappedClaims
    }
}
func (m *ApiApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApiApplication) GetKnownClientApplications()([]string) {
    if m == nil {
        return nil
    } else {
        return m.knownClientApplications
    }
}
func (m *ApiApplication) GetOauth2PermissionScopes()([]PermissionScope) {
    if m == nil {
        return nil
    } else {
        return m.oauth2PermissionScopes
    }
}
func (m *ApiApplication) GetPreAuthorizedApplications()([]PreAuthorizedApplication) {
    if m == nil {
        return nil
    } else {
        return m.preAuthorizedApplications
    }
}
func (m *ApiApplication) GetRequestedAccessTokenVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.requestedAccessTokenVersion
    }
}
func (m *ApiApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["acceptMappedClaims"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAcceptMappedClaims(val)
        return nil
    }
    res["knownClientApplications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetKnownClientApplications(res)
        return nil
    }
    res["oauth2PermissionScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPermissionScope() })
        if err != nil {
            return err
        }
        res := make([]PermissionScope, len(val))
        for i, v := range val {
            res[i] = *(v.(*PermissionScope))
        }
        m.SetOauth2PermissionScopes(res)
        return nil
    }
    res["preAuthorizedApplications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPreAuthorizedApplication() })
        if err != nil {
            return err
        }
        res := make([]PreAuthorizedApplication, len(val))
        for i, v := range val {
            res[i] = *(v.(*PreAuthorizedApplication))
        }
        m.SetPreAuthorizedApplications(res)
        return nil
    }
    res["requestedAccessTokenVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRequestedAccessTokenVersion(val)
        return nil
    }
    return res
}
func (m *ApiApplication) IsNil()(bool) {
    return m == nil
}
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
func (m *ApiApplication) SetAcceptMappedClaims(value *bool)() {
    m.acceptMappedClaims = value
}
func (m *ApiApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApiApplication) SetKnownClientApplications(value []string)() {
    m.knownClientApplications = value
}
func (m *ApiApplication) SetOauth2PermissionScopes(value []PermissionScope)() {
    m.oauth2PermissionScopes = value
}
func (m *ApiApplication) SetPreAuthorizedApplications(value []PreAuthorizedApplication)() {
    m.preAuthorizedApplications = value
}
func (m *ApiApplication) SetRequestedAccessTokenVersion(value *int32)() {
    m.requestedAccessTokenVersion = value
}
