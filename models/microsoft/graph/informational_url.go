package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationalUrl provides operations to manage the collection of application entities.
type InformationalUrl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // CDN URL to the application's logo, Read-only.
    logoUrl *string;
    // Link to the application's marketing page. For example, https://www.contoso.com/app/marketing
    marketingUrl *string;
    // Link to the application's privacy statement. For example, https://www.contoso.com/app/privacy
    privacyStatementUrl *string;
    // Link to the application's support page. For example, https://www.contoso.com/app/support
    supportUrl *string;
    // Link to the application's terms of service statement. For example, https://www.contoso.com/app/termsofservice
    termsOfServiceUrl *string;
}
// NewInformationalUrl instantiates a new informationalUrl and sets the default values.
func NewInformationalUrl()(*InformationalUrl) {
    m := &InformationalUrl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInformationalUrlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationalUrlFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInformationalUrl(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationalUrl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationalUrl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["logoUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoUrl(val)
        }
        return nil
    }
    res["marketingUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarketingUrl(val)
        }
        return nil
    }
    res["privacyStatementUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyStatementUrl(val)
        }
        return nil
    }
    res["supportUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportUrl(val)
        }
        return nil
    }
    res["termsOfServiceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsOfServiceUrl(val)
        }
        return nil
    }
    return res
}
// GetLogoUrl gets the logoUrl property value. CDN URL to the application's logo, Read-only.
func (m *InformationalUrl) GetLogoUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoUrl
    }
}
// GetMarketingUrl gets the marketingUrl property value. Link to the application's marketing page. For example, https://www.contoso.com/app/marketing
func (m *InformationalUrl) GetMarketingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.marketingUrl
    }
}
// GetPrivacyStatementUrl gets the privacyStatementUrl property value. Link to the application's privacy statement. For example, https://www.contoso.com/app/privacy
func (m *InformationalUrl) GetPrivacyStatementUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyStatementUrl
    }
}
// GetSupportUrl gets the supportUrl property value. Link to the application's support page. For example, https://www.contoso.com/app/support
func (m *InformationalUrl) GetSupportUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.supportUrl
    }
}
// GetTermsOfServiceUrl gets the termsOfServiceUrl property value. Link to the application's terms of service statement. For example, https://www.contoso.com/app/termsofservice
func (m *InformationalUrl) GetTermsOfServiceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.termsOfServiceUrl
    }
}
func (m *InformationalUrl) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationalUrl) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLogoUrl sets the logoUrl property value. CDN URL to the application's logo, Read-only.
func (m *InformationalUrl) SetLogoUrl(value *string)() {
    if m != nil {
        m.logoUrl = value
    }
}
// SetMarketingUrl sets the marketingUrl property value. Link to the application's marketing page. For example, https://www.contoso.com/app/marketing
func (m *InformationalUrl) SetMarketingUrl(value *string)() {
    if m != nil {
        m.marketingUrl = value
    }
}
// SetPrivacyStatementUrl sets the privacyStatementUrl property value. Link to the application's privacy statement. For example, https://www.contoso.com/app/privacy
func (m *InformationalUrl) SetPrivacyStatementUrl(value *string)() {
    if m != nil {
        m.privacyStatementUrl = value
    }
}
// SetSupportUrl sets the supportUrl property value. Link to the application's support page. For example, https://www.contoso.com/app/support
func (m *InformationalUrl) SetSupportUrl(value *string)() {
    if m != nil {
        m.supportUrl = value
    }
}
// SetTermsOfServiceUrl sets the termsOfServiceUrl property value. Link to the application's terms of service statement. For example, https://www.contoso.com/app/termsofservice
func (m *InformationalUrl) SetTermsOfServiceUrl(value *string)() {
    if m != nil {
        m.termsOfServiceUrl = value
    }
}
