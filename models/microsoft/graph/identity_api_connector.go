package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IdentityApiConnector struct {
    Entity
    authenticationConfiguration *ApiAuthenticationConfigurationBase;
    displayName *string;
    targetUrl *string;
}
func NewIdentityApiConnector()(*IdentityApiConnector) {
    m := &IdentityApiConnector{
        Entity: *NewEntity(),
    }
    return m
}
func (m *IdentityApiConnector) GetAuthenticationConfiguration()(*ApiAuthenticationConfigurationBase) {
    if m == nil {
        return nil
    } else {
        return m.authenticationConfiguration
    }
}
func (m *IdentityApiConnector) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *IdentityApiConnector) GetTargetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetUrl
    }
}
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
func (m *IdentityApiConnector) SetAuthenticationConfiguration(value *ApiAuthenticationConfigurationBase)() {
    m.authenticationConfiguration = value
}
func (m *IdentityApiConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *IdentityApiConnector) SetTargetUrl(value *string)() {
    m.targetUrl = value
}
