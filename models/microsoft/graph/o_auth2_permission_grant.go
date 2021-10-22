package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OAuth2PermissionGrant struct {
    Entity
    clientId *string;
    consentType *string;
    principalId *string;
    resourceId *string;
    scope *string;
}
func NewOAuth2PermissionGrant()(*OAuth2PermissionGrant) {
    m := &OAuth2PermissionGrant{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OAuth2PermissionGrant) GetClientId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientId
    }
}
func (m *OAuth2PermissionGrant) GetConsentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.consentType
    }
}
func (m *OAuth2PermissionGrant) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
func (m *OAuth2PermissionGrant) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
func (m *OAuth2PermissionGrant) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
func (m *OAuth2PermissionGrant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientId(val)
        return nil
    }
    res["consentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConsentType(val)
        return nil
    }
    res["principalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrincipalId(val)
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScope(val)
        return nil
    }
    return res
}
func (m *OAuth2PermissionGrant) IsNil()(bool) {
    return m == nil
}
func (m *OAuth2PermissionGrant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *OAuth2PermissionGrant) SetClientId(value *string)() {
    m.clientId = value
}
func (m *OAuth2PermissionGrant) SetConsentType(value *string)() {
    m.consentType = value
}
func (m *OAuth2PermissionGrant) SetPrincipalId(value *string)() {
    m.principalId = value
}
func (m *OAuth2PermissionGrant) SetResourceId(value *string)() {
    m.resourceId = value
}
func (m *OAuth2PermissionGrant) SetScope(value *string)() {
    m.scope = value
}
