package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new applicationTemplate and sets the default values.
func NewApplicationTemplate()(*ApplicationTemplate) {
    m := &ApplicationTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the categories property value. The list of categories for the application. Supported values can be: Collaboration, Business Management, Consumer, Content management, CRM, Data services, Developer services, E-commerce, Education, ERP, Finance, Health, Human resources, IT infrastructure, Mail, Management, Marketing, Media, Productivity, Project management, Telecommunications, Tools, Travel, and Web design & hosting.
func (m *ApplicationTemplate) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the description property value. A description of the application.
func (m *ApplicationTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of the application.
func (m *ApplicationTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the homePageUrl property value. The home page URL of the application.
func (m *ApplicationTemplate) GetHomePageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homePageUrl
    }
}
// Gets the logoUrl property value. The URL to get the logo for this application.
func (m *ApplicationTemplate) GetLogoUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logoUrl
    }
}
// Gets the publisher property value. The name of the publisher for this application.
func (m *ApplicationTemplate) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the supportedProvisioningTypes property value. The list of provisioning modes supported by this application. The only valid value is sync.
func (m *ApplicationTemplate) GetSupportedProvisioningTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedProvisioningTypes
    }
}
// Gets the supportedSingleSignOnModes property value. The list of single sign-on modes supported by this application. The supported values are oidc, password, saml, and notSupported.
func (m *ApplicationTemplate) GetSupportedSingleSignOnModes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.supportedSingleSignOnModes
    }
}
// The deserialization information for the current model
func (m *ApplicationTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetCategories(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["homePageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHomePageUrl(val)
        return nil
    }
    res["logoUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLogoUrl(val)
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublisher(val)
        return nil
    }
    res["supportedProvisioningTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSupportedProvisioningTypes(res)
        return nil
    }
    res["supportedSingleSignOnModes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSupportedSingleSignOnModes(res)
        return nil
    }
    return res
}
func (m *ApplicationTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the categories property value. The list of categories for the application. Supported values can be: Collaboration, Business Management, Consumer, Content management, CRM, Data services, Developer services, E-commerce, Education, ERP, Finance, Health, Human resources, IT infrastructure, Mail, Management, Marketing, Media, Productivity, Project management, Telecommunications, Tools, Travel, and Web design & hosting.
// Parameters:
//  - value : Value to set for the categories property.
func (m *ApplicationTemplate) SetCategories(value []string)() {
    m.categories = value
}
// Sets the description property value. A description of the application.
// Parameters:
//  - value : Value to set for the description property.
func (m *ApplicationTemplate) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The name of the application.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ApplicationTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the homePageUrl property value. The home page URL of the application.
// Parameters:
//  - value : Value to set for the homePageUrl property.
func (m *ApplicationTemplate) SetHomePageUrl(value *string)() {
    m.homePageUrl = value
}
// Sets the logoUrl property value. The URL to get the logo for this application.
// Parameters:
//  - value : Value to set for the logoUrl property.
func (m *ApplicationTemplate) SetLogoUrl(value *string)() {
    m.logoUrl = value
}
// Sets the publisher property value. The name of the publisher for this application.
// Parameters:
//  - value : Value to set for the publisher property.
func (m *ApplicationTemplate) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the supportedProvisioningTypes property value. The list of provisioning modes supported by this application. The only valid value is sync.
// Parameters:
//  - value : Value to set for the supportedProvisioningTypes property.
func (m *ApplicationTemplate) SetSupportedProvisioningTypes(value []string)() {
    m.supportedProvisioningTypes = value
}
// Sets the supportedSingleSignOnModes property value. The list of single sign-on modes supported by this application. The supported values are oidc, password, saml, and notSupported.
// Parameters:
//  - value : Value to set for the supportedSingleSignOnModes property.
func (m *ApplicationTemplate) SetSupportedSingleSignOnModes(value []string)() {
    m.supportedSingleSignOnModes = value
}
