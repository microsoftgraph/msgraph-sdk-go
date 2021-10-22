package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InformationalUrl struct {
    additionalData map[string]interface{};
    logoUrl *string;
    marketingUrl *string;
    privacyStatementUrl *string;
    supportUrl *string;
    termsOfServiceUrl *string;
}
func NewInformationalUrl()(*InformationalUrl) {
    m := &InformationalUrl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *InformationalUrl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *InformationalUrl) GetLogoUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoUrl
    }
}
func (m *InformationalUrl) GetMarketingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.marketingUrl
    }
}
func (m *InformationalUrl) GetPrivacyStatementUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyStatementUrl
    }
}
func (m *InformationalUrl) GetSupportUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportUrl
    }
}
func (m *InformationalUrl) GetTermsOfServiceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.termsOfServiceUrl
    }
}
func (m *InformationalUrl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["logoUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLogoUrl(val)
        return nil
    }
    res["marketingUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMarketingUrl(val)
        return nil
    }
    res["privacyStatementUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrivacyStatementUrl(val)
        return nil
    }
    res["supportUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSupportUrl(val)
        return nil
    }
    res["termsOfServiceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTermsOfServiceUrl(val)
        return nil
    }
    return res
}
func (m *InformationalUrl) IsNil()(bool) {
    return m == nil
}
func (m *InformationalUrl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("logoUrl", m.GetLogoUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("marketingUrl", m.GetMarketingUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("privacyStatementUrl", m.GetPrivacyStatementUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("supportUrl", m.GetSupportUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("termsOfServiceUrl", m.GetTermsOfServiceUrl())
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
func (m *InformationalUrl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *InformationalUrl) SetLogoUrl(value *string)() {
    m.logoUrl = value
}
func (m *InformationalUrl) SetMarketingUrl(value *string)() {
    m.marketingUrl = value
}
func (m *InformationalUrl) SetPrivacyStatementUrl(value *string)() {
    m.privacyStatementUrl = value
}
func (m *InformationalUrl) SetSupportUrl(value *string)() {
    m.supportUrl = value
}
func (m *InformationalUrl) SetTermsOfServiceUrl(value *string)() {
    m.termsOfServiceUrl = value
}
