package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WebApplication struct {
    additionalData map[string]interface{};
    homePageUrl *string;
    implicitGrantSettings *ImplicitGrantSettings;
    logoutUrl *string;
    redirectUris []string;
}
func NewWebApplication()(*WebApplication) {
    m := &WebApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WebApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WebApplication) GetHomePageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homePageUrl
    }
}
func (m *WebApplication) GetImplicitGrantSettings()(*ImplicitGrantSettings) {
    if m == nil {
        return nil
    } else {
        return m.implicitGrantSettings
    }
}
func (m *WebApplication) GetLogoutUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoutUrl
    }
}
func (m *WebApplication) GetRedirectUris()([]string) {
    if m == nil {
        return nil
    } else {
        return m.redirectUris
    }
}
func (m *WebApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["homePageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHomePageUrl(val)
        return nil
    }
    res["implicitGrantSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImplicitGrantSettings() })
        if err != nil {
            return err
        }
        m.SetImplicitGrantSettings(val.(*ImplicitGrantSettings))
        return nil
    }
    res["logoutUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLogoutUrl(val)
        return nil
    }
    res["redirectUris"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRedirectUris(res)
        return nil
    }
    return res
}
func (m *WebApplication) IsNil()(bool) {
    return m == nil
}
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
func (m *WebApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WebApplication) SetHomePageUrl(value *string)() {
    m.homePageUrl = value
}
func (m *WebApplication) SetImplicitGrantSettings(value *ImplicitGrantSettings)() {
    m.implicitGrantSettings = value
}
func (m *WebApplication) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
func (m *WebApplication) SetRedirectUris(value []string)() {
    m.redirectUris = value
}
