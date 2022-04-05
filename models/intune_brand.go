package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntuneBrand intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
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
    darkBackgroundLogo MimeContentable;
    // Company/organization name that is displayed to end users.
    displayName *string;
    // Logo image displayed in Company Portal apps which have a light background behind the logo.
    lightBackgroundLogo MimeContentable;
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
    themeColor RgbColorable;
}
// NewIntuneBrand instantiates a new intuneBrand and sets the default values.
func NewIntuneBrand()(*IntuneBrand) {
    m := &IntuneBrand{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIntuneBrandFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntuneBrandFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntuneBrand(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntuneBrand) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContactITEmailAddress gets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITEmailAddress
    }
}
// GetContactITName gets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITName
    }
}
// GetContactITNotes gets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITNotes
    }
}
// GetContactITPhoneNumber gets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) GetContactITPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contactITPhoneNumber
    }
}
// GetDarkBackgroundLogo gets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) GetDarkBackgroundLogo()(MimeContentable) {
    if m == nil {
        return nil
    } else {
        return m.darkBackgroundLogo
    }
}
// GetDisplayName gets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntuneBrand) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contactITEmailAddress"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITEmailAddress(val)
        }
        return nil
    }
    res["contactITName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITName(val)
        }
        return nil
    }
    res["contactITNotes"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITNotes(val)
        }
        return nil
    }
    res["contactITPhoneNumber"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactITPhoneNumber(val)
        }
        return nil
    }
    res["darkBackgroundLogo"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDarkBackgroundLogo(val.(MimeContentable))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lightBackgroundLogo"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLightBackgroundLogo(val.(MimeContentable))
        }
        return nil
    }
    res["onlineSupportSiteName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteName(val)
        }
        return nil
    }
    res["onlineSupportSiteUrl"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineSupportSiteUrl(val)
        }
        return nil
    }
    res["privacyUrl"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyUrl(val)
        }
        return nil
    }
    res["showDisplayNameNextToLogo"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowDisplayNameNextToLogo(val)
        }
        return nil
    }
    res["showLogo"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowLogo(val)
        }
        return nil
    }
    res["showNameNextToLogo"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowNameNextToLogo(val)
        }
        return nil
    }
    res["themeColor"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRgbColorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThemeColor(val.(RgbColorable))
        }
        return nil
    }
    return res
}
// GetLightBackgroundLogo gets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) GetLightBackgroundLogo()(MimeContentable) {
    if m == nil {
        return nil
    } else {
        return m.lightBackgroundLogo
    }
}
// GetOnlineSupportSiteName gets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteName
    }
}
// GetOnlineSupportSiteUrl gets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) GetOnlineSupportSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineSupportSiteUrl
    }
}
// GetPrivacyUrl gets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) GetPrivacyUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyUrl
    }
}
// GetShowDisplayNameNextToLogo gets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowDisplayNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showDisplayNameNextToLogo
    }
}
// GetShowLogo gets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) GetShowLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showLogo
    }
}
// GetShowNameNextToLogo gets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) GetShowNameNextToLogo()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showNameNextToLogo
    }
}
// GetThemeColor gets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) GetThemeColor()(RgbColorable) {
    if m == nil {
        return nil
    } else {
        return m.themeColor
    }
}
// Serialize serializes information the current object
func (m *IntuneBrand) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntuneBrand) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContactITEmailAddress sets the contactITEmailAddress property value. Email address of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITEmailAddress(value *string)() {
    if m != nil {
        m.contactITEmailAddress = value
    }
}
// SetContactITName sets the contactITName property value. Name of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITName(value *string)() {
    if m != nil {
        m.contactITName = value
    }
}
// SetContactITNotes sets the contactITNotes property value. Text comments regarding the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITNotes(value *string)() {
    if m != nil {
        m.contactITNotes = value
    }
}
// SetContactITPhoneNumber sets the contactITPhoneNumber property value. Phone number of the person/organization responsible for IT support.
func (m *IntuneBrand) SetContactITPhoneNumber(value *string)() {
    if m != nil {
        m.contactITPhoneNumber = value
    }
}
// SetDarkBackgroundLogo sets the darkBackgroundLogo property value. Logo image displayed in Company Portal apps which have a dark background behind the logo.
func (m *IntuneBrand) SetDarkBackgroundLogo(value MimeContentable)() {
    if m != nil {
        m.darkBackgroundLogo = value
    }
}
// SetDisplayName sets the displayName property value. Company/organization name that is displayed to end users.
func (m *IntuneBrand) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLightBackgroundLogo sets the lightBackgroundLogo property value. Logo image displayed in Company Portal apps which have a light background behind the logo.
func (m *IntuneBrand) SetLightBackgroundLogo(value MimeContentable)() {
    if m != nil {
        m.lightBackgroundLogo = value
    }
}
// SetOnlineSupportSiteName sets the onlineSupportSiteName property value. Display name of the company/organization’s IT helpdesk site.
func (m *IntuneBrand) SetOnlineSupportSiteName(value *string)() {
    if m != nil {
        m.onlineSupportSiteName = value
    }
}
// SetOnlineSupportSiteUrl sets the onlineSupportSiteUrl property value. URL to the company/organization’s IT helpdesk site.
func (m *IntuneBrand) SetOnlineSupportSiteUrl(value *string)() {
    if m != nil {
        m.onlineSupportSiteUrl = value
    }
}
// SetPrivacyUrl sets the privacyUrl property value. URL to the company/organization’s privacy policy.
func (m *IntuneBrand) SetPrivacyUrl(value *string)() {
    if m != nil {
        m.privacyUrl = value
    }
}
// SetShowDisplayNameNextToLogo sets the showDisplayNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) SetShowDisplayNameNextToLogo(value *bool)() {
    if m != nil {
        m.showDisplayNameNextToLogo = value
    }
}
// SetShowLogo sets the showLogo property value. Boolean that represents whether the administrator-supplied logo images are shown or not shown.
func (m *IntuneBrand) SetShowLogo(value *bool)() {
    if m != nil {
        m.showLogo = value
    }
}
// SetShowNameNextToLogo sets the showNameNextToLogo property value. Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
func (m *IntuneBrand) SetShowNameNextToLogo(value *bool)() {
    if m != nil {
        m.showNameNextToLogo = value
    }
}
// SetThemeColor sets the themeColor property value. Primary theme color used in the Company Portal applications and web portal.
func (m *IntuneBrand) SetThemeColor(value RgbColorable)() {
    if m != nil {
        m.themeColor = value
    }
}
