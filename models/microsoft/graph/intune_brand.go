package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IntuneBrand struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Email address of the person/organization responsible for IT support.
    contactITEmailAddress *string;
    // Name of the person/organization responsible for IT support.
    contactITName *string;
    // Text comments regarding the person/organization responsible for IT support.
    contactITNotes *string;
    // Phone number of the person/organization responsible for IT support.
    contactITPhoneNumber *string;
    // Logo image displayed in Company Portal apps which have a dark background behind the logo.
    darkBackgroundLogo *MimeContent;
    // Company/organization name that is displayed to end users.
    displayName *string;
    // Logo image displayed in Company Portal apps which have a light background behind the logo.
    lightBackgroundLogo *MimeContent;
    // Display name of the company/organization’s IT helpdesk site.
    onlineSupportSiteName *string;
    // URL to the company/organization’s IT helpdesk site.
    onlineSupportSiteUrl *string;
    // URL to the company/organization’s privacy policy.
    privacyUrl *string;
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
    showDisplayNameNextToLogo *bool;
    // Boolean that represents whether the administrator-supplied logo images are shown or not shown.
    showLogo *bool;
    // Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
    showNameNextToLogo *bool;
    // Primary theme color used in the Company Portal applications and web portal.
    themeColor *RgbColor;
}
// Instantiates a new intuneBrand and sets the default values.
func NewIntuneBrand()(*IntuneBrand) {
    m := &IntuneBrand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntuneBrand) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITEmailAddress
    }
}
// Gets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITName
    }
}
// Gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITNotes
    }
}
// Gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITPhoneNumber
    }
}
// Gets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) GetDarkBackgroundLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.darkBackgroundLogo
    }
}
// Gets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) GetLightBackgroundLogo()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.lightBackgroundLogo
    }
}
// Gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteName
    }
}
// Gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteUrl
    }
}
// Gets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) GetPrivacyUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyUrl
    }
}
// Gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowDisplayNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showDisplayNameNextToLogo
    }
}
// Gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) GetShowLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLogo
    }
}
// Gets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showNameNextToLogo
    }
}
// Gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) GetThemeColor()(*RgbColor) {
    if m == nil {
        return nil
    } else {
        return m.themeColor
    }
}
// The deserialization information for the current model
func (m *IntuneBrand) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contactITEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITEmailAddress(val)
        }
        return nil
    }
    res["contactITName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITName(val)
        }
        return nil
    }
    res["contactITNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITNotes(val)
        }
        return nil
    }
    res["contactITPhoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITPhoneNumber(val)
        }
        return nil
    }
    res["darkBackgroundLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDarkBackgroundLogo(val.(*MimeContent))
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
    res["lightBackgroundLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLightBackgroundLogo(val.(*MimeContent))
        }
        return nil
    }
    res["onlineSupportSiteName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteName(val)
        }
        return nil
    }
    res["onlineSupportSiteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteUrl(val)
        }
        return nil
    }
    res["privacyUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyUrl(val)
        }
        return nil
    }
    res["showDisplayNameNextToLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowDisplayNameNextToLogo(val)
        }
        return nil
    }
    res["showLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowLogo(val)
        }
        return nil
    }
    res["showNameNextToLogo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowNameNextToLogo(val)
        }
        return nil
    }
    res["themeColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRgbColor() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThemeColor(val.(*RgbColor))
        }
        return nil
    }
    return res
}
func (m *IntuneBrand) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IntuneBrand) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contactITEmailAddress", m.GetContactITEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITName", m.GetContactITName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITNotes", m.GetContactITNotes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contactITPhoneNumber", m.GetContactITPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("darkBackgroundLogo", m.GetDarkBackgroundLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lightBackgroundLogo", m.GetLightBackgroundLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onlineSupportSiteName", m.GetOnlineSupportSiteName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("onlineSupportSiteUrl", m.GetOnlineSupportSiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("privacyUrl", m.GetPrivacyUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showDisplayNameNextToLogo", m.GetShowDisplayNameNextToLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showLogo", m.GetShowLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showNameNextToLogo", m.GetShowNameNextToLogo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("themeColor", m.GetThemeColor())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *IntuneBrand) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
// Parameters:
//  - value : Value to set for the contactITEmailAddress property.
func (m *IntuneBrand) SetContactITEmailAddress(value *string)() {
    m.contactITEmailAddress = value
}
// Sets the contactITName property value. Name of the person/organization responsible for IT support.
// Parameters:
//  - value : Value to set for the contactITName property.
func (m *IntuneBrand) SetContactITName(value *string)() {
    m.contactITName = value
}
// Sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
// Parameters:
//  - value : Value to set for the contactITNotes property.
func (m *IntuneBrand) SetContactITNotes(value *string)() {
    m.contactITNotes = value
}
// Sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
// Parameters:
//  - value : Value to set for the contactITPhoneNumber property.
func (m *IntuneBrand) SetContactITPhoneNumber(value *string)() {
    m.contactITPhoneNumber = value
}
// Sets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
// Parameters:
//  - value : Value to set for the darkBackgroundLogo property.
func (m *IntuneBrand) SetDarkBackgroundLogo(value *MimeContent)() {
    m.darkBackgroundLogo = value
}
// Sets the displayName property value. Company/organization name that is displayed to end users.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *IntuneBrand) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
// Parameters:
//  - value : Value to set for the lightBackgroundLogo property.
func (m *IntuneBrand) SetLightBackgroundLogo(value *MimeContent)() {
    m.lightBackgroundLogo = value
}
// Sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
// Parameters:
//  - value : Value to set for the onlineSupportSiteName property.
func (m *IntuneBrand) SetOnlineSupportSiteName(value *string)() {
    m.onlineSupportSiteName = value
}
// Sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
// Parameters:
//  - value : Value to set for the onlineSupportSiteUrl property.
func (m *IntuneBrand) SetOnlineSupportSiteUrl(value *string)() {
    m.onlineSupportSiteUrl = value
}
// Sets the privacyUrl property value. URL to the company/organization’s privacy policy.
// Parameters:
//  - value : Value to set for the privacyUrl property.
func (m *IntuneBrand) SetPrivacyUrl(value *string)() {
    m.privacyUrl = value
}
// Sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
// Parameters:
//  - value : Value to set for the showDisplayNameNextToLogo property.
func (m *IntuneBrand) SetShowDisplayNameNextToLogo(value *bool)() {
    m.showDisplayNameNextToLogo = value
}
// Sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
// Parameters:
//  - value : Value to set for the showLogo property.
func (m *IntuneBrand) SetShowLogo(value *bool)() {
    m.showLogo = value
}
// Sets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
// Parameters:
//  - value : Value to set for the showNameNextToLogo property.
func (m *IntuneBrand) SetShowNameNextToLogo(value *bool)() {
    m.showNameNextToLogo = value
}
// Sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
// Parameters:
//  - value : Value to set for the themeColor property.
func (m *IntuneBrand) SetThemeColor(value *RgbColor)() {
    m.themeColor = value
}
