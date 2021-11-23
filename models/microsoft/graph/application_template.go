package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplicationTemplate 
type ApplicationTemplate struct {
    Entity
    // The list of categories for the application. Supported values can be: Collaboration, Business Management, Consumer, Content management, CRM, Data services, Developer services, E-commerce, Education, ERP, Finance, Health, Human resources, IT infrastructure, Mail, Management, Marketing, Media, Productivity, Project management, Telecommunications, Tools, Travel, and Web design & hosting.
    categories []string;
    // A description of the application.
    description *string;
    // The name of the application.
    displayName *string;
    // The home page URL of the application.
    homePageUrl *string;
    // The URL to get the logo for this application.
    logoUrl *string;
    // The name of the publisher for this application.
    publisher *string;
    // The list of provisioning modes supported by this application. The only valid value is sync.
    supportedProvisioningTypes []string;
    // The list of single sign-on modes supported by this application. The supported values are oidc, password, saml, and notSupported.
    supportedSingleSignOnModes []string;
}
// NewApplicationTemplate instantiates a new applicationTemplate and sets the default values.
func NewApplicationTemplate()(*ApplicationTemplate) {
    m := &ApplicationTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategories gets the categories property value. The list of categories for the application. Supported values can be: Collaboration, Business Management, Consumer, Content management, CRM, Data services, Developer services, E-commerce, Education, ERP, Finance, Health, Human resources, IT infrastructure, Mail, Management, Marketing, Media, Productivity, Project management, Telecommunications, Tools, Travel, and Web design & hosting.
func (m *ApplicationTemplate) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetDescription gets the description property value. A description of the application.
func (m *ApplicationTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name of the application.
func (m *ApplicationTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetHomePageUrl gets the homePageUrl property value. The home page URL of the application.
func (m *ApplicationTemplate) GetHomePageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homePageUrl
    }
}
// GetLogoUrl gets the logoUrl property value. The URL to get the logo for this application.
func (m *ApplicationTemplate) GetLogoUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoUrl
    }
}
// GetPublisher gets the publisher property value. The name of the publisher for this application.
func (m *ApplicationTemplate) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetSupportedProvisioningTypes gets the supportedProvisioningTypes property value. The list of provisioning modes supported by this application. The only valid value is sync.
func (m *ApplicationTemplate) GetSupportedProvisioningTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedProvisioningTypes
    }
}
// GetSupportedSingleSignOnModes gets the supportedSingleSignOnModes property value. The list of single sign-on modes supported by this application. The supported values are oidc, password, saml, and notSupported.
func (m *ApplicationTemplate) GetSupportedSingleSignOnModes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedSingleSignOnModes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["homePageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomePageUrl(val)
        }
        return nil
    }
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
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["supportedProvisioningTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedProvisioningTypes(res)
        }
        return nil
    }
    res["supportedSingleSignOnModes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedSingleSignOnModes(res)
        }
        return nil
    }
    return res
}
func (m *ApplicationTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApplicationTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteStringValue("homePageUrl", m.GetHomePageUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("logoUrl", m.GetLogoUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("supportedProvisioningTypes", m.GetSupportedProvisioningTypes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("supportedSingleSignOnModes", m.GetSupportedSingleSignOnModes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategories sets the categories property value. The list of categories for the application. Supported values can be: Collaboration, Business Management, Consumer, Content management, CRM, Data services, Developer services, E-commerce, Education, ERP, Finance, Health, Human resources, IT infrastructure, Mail, Management, Marketing, Media, Productivity, Project management, Telecommunications, Tools, Travel, and Web design & hosting.
func (m *ApplicationTemplate) SetCategories(value []string)() {
    m.categories = value
}
// SetDescription sets the description property value. A description of the application.
func (m *ApplicationTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of the application.
func (m *ApplicationTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHomePageUrl sets the homePageUrl property value. The home page URL of the application.
func (m *ApplicationTemplate) SetHomePageUrl(value *string)() {
    m.homePageUrl = value
}
// SetLogoUrl sets the logoUrl property value. The URL to get the logo for this application.
func (m *ApplicationTemplate) SetLogoUrl(value *string)() {
    m.logoUrl = value
}
// SetPublisher sets the publisher property value. The name of the publisher for this application.
func (m *ApplicationTemplate) SetPublisher(value *string)() {
    m.publisher = value
}
// SetSupportedProvisioningTypes sets the supportedProvisioningTypes property value. The list of provisioning modes supported by this application. The only valid value is sync.
func (m *ApplicationTemplate) SetSupportedProvisioningTypes(value []string)() {
    m.supportedProvisioningTypes = value
}
// SetSupportedSingleSignOnModes sets the supportedSingleSignOnModes property value. The list of single sign-on modes supported by this application. The supported values are oidc, password, saml, and notSupported.
func (m *ApplicationTemplate) SetSupportedSingleSignOnModes(value []string)() {
    m.supportedSingleSignOnModes = value
}
