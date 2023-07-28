package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalBrandingProperties 
type OrganizationalBrandingProperties struct {
    Entity
}
// NewOrganizationalBrandingProperties instantiates a new organizationalBrandingProperties and sets the default values.
func NewOrganizationalBrandingProperties()(*OrganizationalBrandingProperties) {
    m := &OrganizationalBrandingProperties{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOrganizationalBrandingPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalBrandingPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.organizationalBranding":
                        return NewOrganizationalBranding(), nil
                    case "#microsoft.graph.organizationalBrandingLocalization":
                        return NewOrganizationalBrandingLocalization(), nil
                }
            }
        }
    }
    return NewOrganizationalBrandingProperties(), nil
}
// GetBackgroundColor gets the backgroundColor property value. Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
func (m *OrganizationalBrandingProperties) GetBackgroundColor()(*string) {
    val, err := m.GetBackingStore().Get("backgroundColor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackgroundImage gets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
func (m *OrganizationalBrandingProperties) GetBackgroundImage()([]byte) {
    val, err := m.GetBackingStore().Get("backgroundImage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetBackgroundImageRelativeUrl gets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBackgroundImageRelativeUrl()(*string) {
    val, err := m.GetBackingStore().Get("backgroundImageRelativeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBannerLogo gets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetBannerLogo()([]byte) {
    val, err := m.GetBackingStore().Get("bannerLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetBannerLogoRelativeUrl gets the bannerLogoRelativeUrl property value. A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetBannerLogoRelativeUrl()(*string) {
    val, err := m.GetBackingStore().Get("bannerLogoRelativeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCdnList gets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
func (m *OrganizationalBrandingProperties) GetCdnList()([]string) {
    val, err := m.GetBackingStore().Get("cdnList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetCustomAccountResetCredentialsUrl gets the customAccountResetCredentialsUrl property value. The customAccountResetCredentialsUrl property
func (m *OrganizationalBrandingProperties) GetCustomAccountResetCredentialsUrl()(*string) {
    val, err := m.GetBackingStore().Get("customAccountResetCredentialsUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomCannotAccessYourAccountText gets the customCannotAccessYourAccountText property value. The customCannotAccessYourAccountText property
func (m *OrganizationalBrandingProperties) GetCustomCannotAccessYourAccountText()(*string) {
    val, err := m.GetBackingStore().Get("customCannotAccessYourAccountText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomCannotAccessYourAccountUrl gets the customCannotAccessYourAccountUrl property value. The customCannotAccessYourAccountUrl property
func (m *OrganizationalBrandingProperties) GetCustomCannotAccessYourAccountUrl()(*string) {
    val, err := m.GetBackingStore().Get("customCannotAccessYourAccountUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomCSS gets the customCSS property value. The customCSS property
func (m *OrganizationalBrandingProperties) GetCustomCSS()([]byte) {
    val, err := m.GetBackingStore().Get("customCSS")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetCustomCSSRelativeUrl gets the customCSSRelativeUrl property value. The customCSSRelativeUrl property
func (m *OrganizationalBrandingProperties) GetCustomCSSRelativeUrl()(*string) {
    val, err := m.GetBackingStore().Get("customCSSRelativeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomForgotMyPasswordText gets the customForgotMyPasswordText property value. The customForgotMyPasswordText property
func (m *OrganizationalBrandingProperties) GetCustomForgotMyPasswordText()(*string) {
    val, err := m.GetBackingStore().Get("customForgotMyPasswordText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomPrivacyAndCookiesText gets the customPrivacyAndCookiesText property value. The customPrivacyAndCookiesText property
func (m *OrganizationalBrandingProperties) GetCustomPrivacyAndCookiesText()(*string) {
    val, err := m.GetBackingStore().Get("customPrivacyAndCookiesText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomPrivacyAndCookiesUrl gets the customPrivacyAndCookiesUrl property value. The customPrivacyAndCookiesUrl property
func (m *OrganizationalBrandingProperties) GetCustomPrivacyAndCookiesUrl()(*string) {
    val, err := m.GetBackingStore().Get("customPrivacyAndCookiesUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomResetItNowText gets the customResetItNowText property value. The customResetItNowText property
func (m *OrganizationalBrandingProperties) GetCustomResetItNowText()(*string) {
    val, err := m.GetBackingStore().Get("customResetItNowText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomTermsOfUseText gets the customTermsOfUseText property value. The customTermsOfUseText property
func (m *OrganizationalBrandingProperties) GetCustomTermsOfUseText()(*string) {
    val, err := m.GetBackingStore().Get("customTermsOfUseText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomTermsOfUseUrl gets the customTermsOfUseUrl property value. The customTermsOfUseUrl property
func (m *OrganizationalBrandingProperties) GetCustomTermsOfUseUrl()(*string) {
    val, err := m.GetBackingStore().Get("customTermsOfUseUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFavicon gets the favicon property value. The favicon property
func (m *OrganizationalBrandingProperties) GetFavicon()([]byte) {
    val, err := m.GetBackingStore().Get("favicon")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetFaviconRelativeUrl gets the faviconRelativeUrl property value. The faviconRelativeUrl property
func (m *OrganizationalBrandingProperties) GetFaviconRelativeUrl()(*string) {
    val, err := m.GetBackingStore().Get("faviconRelativeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalBrandingProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["backgroundColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundColor(val)
        }
        return nil
    }
    res["backgroundImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundImage(val)
        }
        return nil
    }
    res["backgroundImageRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundImageRelativeUrl(val)
        }
        return nil
    }
    res["bannerLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerLogo(val)
        }
        return nil
    }
    res["bannerLogoRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBannerLogoRelativeUrl(val)
        }
        return nil
    }
    res["cdnList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCdnList(res)
        }
        return nil
    }
    res["customAccountResetCredentialsUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomAccountResetCredentialsUrl(val)
        }
        return nil
    }
    res["customCannotAccessYourAccountText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCannotAccessYourAccountText(val)
        }
        return nil
    }
    res["customCannotAccessYourAccountUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCannotAccessYourAccountUrl(val)
        }
        return nil
    }
    res["customCSS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCSS(val)
        }
        return nil
    }
    res["customCSSRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomCSSRelativeUrl(val)
        }
        return nil
    }
    res["customForgotMyPasswordText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomForgotMyPasswordText(val)
        }
        return nil
    }
    res["customPrivacyAndCookiesText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrivacyAndCookiesText(val)
        }
        return nil
    }
    res["customPrivacyAndCookiesUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomPrivacyAndCookiesUrl(val)
        }
        return nil
    }
    res["customResetItNowText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomResetItNowText(val)
        }
        return nil
    }
    res["customTermsOfUseText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomTermsOfUseText(val)
        }
        return nil
    }
    res["customTermsOfUseUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomTermsOfUseUrl(val)
        }
        return nil
    }
    res["favicon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFavicon(val)
        }
        return nil
    }
    res["faviconRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaviconRelativeUrl(val)
        }
        return nil
    }
    res["headerBackgroundColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderBackgroundColor(val)
        }
        return nil
    }
    res["headerLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderLogo(val)
        }
        return nil
    }
    res["headerLogoRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderLogoRelativeUrl(val)
        }
        return nil
    }
    res["loginPageLayoutConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLoginPageLayoutConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginPageLayoutConfiguration(val.(LoginPageLayoutConfigurationable))
        }
        return nil
    }
    res["loginPageTextVisibilitySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLoginPageTextVisibilitySettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginPageTextVisibilitySettings(val.(LoginPageTextVisibilitySettingsable))
        }
        return nil
    }
    res["signInPageText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInPageText(val)
        }
        return nil
    }
    res["squareLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogo(val)
        }
        return nil
    }
    res["squareLogoDark"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoDark(val)
        }
        return nil
    }
    res["squareLogoDarkRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoDarkRelativeUrl(val)
        }
        return nil
    }
    res["squareLogoRelativeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSquareLogoRelativeUrl(val)
        }
        return nil
    }
    res["usernameHintText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameHintText(val)
        }
        return nil
    }
    return res
}
// GetHeaderBackgroundColor gets the headerBackgroundColor property value. The headerBackgroundColor property
func (m *OrganizationalBrandingProperties) GetHeaderBackgroundColor()(*string) {
    val, err := m.GetBackingStore().Get("headerBackgroundColor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHeaderLogo gets the headerLogo property value. The headerLogo property
func (m *OrganizationalBrandingProperties) GetHeaderLogo()([]byte) {
    val, err := m.GetBackingStore().Get("headerLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetHeaderLogoRelativeUrl gets the headerLogoRelativeUrl property value. The headerLogoRelativeUrl property
func (m *OrganizationalBrandingProperties) GetHeaderLogoRelativeUrl()(*string) {
    val, err := m.GetBackingStore().Get("headerLogoRelativeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLoginPageLayoutConfiguration gets the loginPageLayoutConfiguration property value. The loginPageLayoutConfiguration property
func (m *OrganizationalBrandingProperties) GetLoginPageLayoutConfiguration()(LoginPageLayoutConfigurationable) {
    val, err := m.GetBackingStore().Get("loginPageLayoutConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LoginPageLayoutConfigurationable)
    }
    return nil
}
// GetLoginPageTextVisibilitySettings gets the loginPageTextVisibilitySettings property value. The loginPageTextVisibilitySettings property
func (m *OrganizationalBrandingProperties) GetLoginPageTextVisibilitySettings()(LoginPageTextVisibilitySettingsable) {
    val, err := m.GetBackingStore().Get("loginPageTextVisibilitySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LoginPageTextVisibilitySettingsable)
    }
    return nil
}
// GetSignInPageText gets the signInPageText property value. Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
func (m *OrganizationalBrandingProperties) GetSignInPageText()(*string) {
    val, err := m.GetBackingStore().Get("signInPageText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSquareLogo gets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) GetSquareLogo()([]byte) {
    val, err := m.GetBackingStore().Get("squareLogo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetSquareLogoDark gets the squareLogoDark property value. The squareLogoDark property
func (m *OrganizationalBrandingProperties) GetSquareLogoDark()([]byte) {
    val, err := m.GetBackingStore().Get("squareLogoDark")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetSquareLogoDarkRelativeUrl gets the squareLogoDarkRelativeUrl property value. The squareLogoDarkRelativeUrl property
func (m *OrganizationalBrandingProperties) GetSquareLogoDarkRelativeUrl()(*string) {
    val, err := m.GetBackingStore().Get("squareLogoDarkRelativeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSquareLogoRelativeUrl gets the squareLogoRelativeUrl property value. A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) GetSquareLogoRelativeUrl()(*string) {
    val, err := m.GetBackingStore().Get("squareLogoRelativeUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUsernameHintText gets the usernameHintText property value. String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
func (m *OrganizationalBrandingProperties) GetUsernameHintText()(*string) {
    val, err := m.GetBackingStore().Get("usernameHintText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OrganizationalBrandingProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("backgroundColor", m.GetBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("backgroundImage", m.GetBackgroundImage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("backgroundImageRelativeUrl", m.GetBackgroundImageRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("bannerLogo", m.GetBannerLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bannerLogoRelativeUrl", m.GetBannerLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    if m.GetCdnList() != nil {
        err = writer.WriteCollectionOfStringValues("cdnList", m.GetCdnList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customAccountResetCredentialsUrl", m.GetCustomAccountResetCredentialsUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCannotAccessYourAccountText", m.GetCustomCannotAccessYourAccountText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCannotAccessYourAccountUrl", m.GetCustomCannotAccessYourAccountUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("customCSS", m.GetCustomCSS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customCSSRelativeUrl", m.GetCustomCSSRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customForgotMyPasswordText", m.GetCustomForgotMyPasswordText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customPrivacyAndCookiesText", m.GetCustomPrivacyAndCookiesText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customPrivacyAndCookiesUrl", m.GetCustomPrivacyAndCookiesUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customResetItNowText", m.GetCustomResetItNowText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customTermsOfUseText", m.GetCustomTermsOfUseText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customTermsOfUseUrl", m.GetCustomTermsOfUseUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("favicon", m.GetFavicon())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("faviconRelativeUrl", m.GetFaviconRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerBackgroundColor", m.GetHeaderBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("headerLogo", m.GetHeaderLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerLogoRelativeUrl", m.GetHeaderLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("loginPageLayoutConfiguration", m.GetLoginPageLayoutConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("loginPageTextVisibilitySettings", m.GetLoginPageTextVisibilitySettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInPageText", m.GetSignInPageText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("squareLogo", m.GetSquareLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("squareLogoDark", m.GetSquareLogoDark())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoDarkRelativeUrl", m.GetSquareLogoDarkRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("squareLogoRelativeUrl", m.GetSquareLogoRelativeUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usernameHintText", m.GetUsernameHintText())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackgroundColor sets the backgroundColor property value. Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF.
func (m *OrganizationalBrandingProperties) SetBackgroundColor(value *string)() {
    err := m.GetBackingStore().Set("backgroundColor", value)
    if err != nil {
        panic(err)
    }
}
// SetBackgroundImage sets the backgroundImage property value. Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster.
func (m *OrganizationalBrandingProperties) SetBackgroundImage(value []byte)() {
    err := m.GetBackingStore().Set("backgroundImage", value)
    if err != nil {
        panic(err)
    }
}
// SetBackgroundImageRelativeUrl sets the backgroundImageRelativeUrl property value. A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetBackgroundImageRelativeUrl(value *string)() {
    err := m.GetBackingStore().Set("backgroundImageRelativeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetBannerLogo sets the bannerLogo property value. A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) SetBannerLogo(value []byte)() {
    err := m.GetBackingStore().Set("bannerLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetBannerLogoRelativeUrl sets the bannerLogoRelativeUrl property value. A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetBannerLogoRelativeUrl(value *string)() {
    err := m.GetBackingStore().Set("bannerLogoRelativeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCdnList sets the cdnList property value. A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only.
func (m *OrganizationalBrandingProperties) SetCdnList(value []string)() {
    err := m.GetBackingStore().Set("cdnList", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomAccountResetCredentialsUrl sets the customAccountResetCredentialsUrl property value. The customAccountResetCredentialsUrl property
func (m *OrganizationalBrandingProperties) SetCustomAccountResetCredentialsUrl(value *string)() {
    err := m.GetBackingStore().Set("customAccountResetCredentialsUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomCannotAccessYourAccountText sets the customCannotAccessYourAccountText property value. The customCannotAccessYourAccountText property
func (m *OrganizationalBrandingProperties) SetCustomCannotAccessYourAccountText(value *string)() {
    err := m.GetBackingStore().Set("customCannotAccessYourAccountText", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomCannotAccessYourAccountUrl sets the customCannotAccessYourAccountUrl property value. The customCannotAccessYourAccountUrl property
func (m *OrganizationalBrandingProperties) SetCustomCannotAccessYourAccountUrl(value *string)() {
    err := m.GetBackingStore().Set("customCannotAccessYourAccountUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomCSS sets the customCSS property value. The customCSS property
func (m *OrganizationalBrandingProperties) SetCustomCSS(value []byte)() {
    err := m.GetBackingStore().Set("customCSS", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomCSSRelativeUrl sets the customCSSRelativeUrl property value. The customCSSRelativeUrl property
func (m *OrganizationalBrandingProperties) SetCustomCSSRelativeUrl(value *string)() {
    err := m.GetBackingStore().Set("customCSSRelativeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomForgotMyPasswordText sets the customForgotMyPasswordText property value. The customForgotMyPasswordText property
func (m *OrganizationalBrandingProperties) SetCustomForgotMyPasswordText(value *string)() {
    err := m.GetBackingStore().Set("customForgotMyPasswordText", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomPrivacyAndCookiesText sets the customPrivacyAndCookiesText property value. The customPrivacyAndCookiesText property
func (m *OrganizationalBrandingProperties) SetCustomPrivacyAndCookiesText(value *string)() {
    err := m.GetBackingStore().Set("customPrivacyAndCookiesText", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomPrivacyAndCookiesUrl sets the customPrivacyAndCookiesUrl property value. The customPrivacyAndCookiesUrl property
func (m *OrganizationalBrandingProperties) SetCustomPrivacyAndCookiesUrl(value *string)() {
    err := m.GetBackingStore().Set("customPrivacyAndCookiesUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomResetItNowText sets the customResetItNowText property value. The customResetItNowText property
func (m *OrganizationalBrandingProperties) SetCustomResetItNowText(value *string)() {
    err := m.GetBackingStore().Set("customResetItNowText", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomTermsOfUseText sets the customTermsOfUseText property value. The customTermsOfUseText property
func (m *OrganizationalBrandingProperties) SetCustomTermsOfUseText(value *string)() {
    err := m.GetBackingStore().Set("customTermsOfUseText", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomTermsOfUseUrl sets the customTermsOfUseUrl property value. The customTermsOfUseUrl property
func (m *OrganizationalBrandingProperties) SetCustomTermsOfUseUrl(value *string)() {
    err := m.GetBackingStore().Set("customTermsOfUseUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetFavicon sets the favicon property value. The favicon property
func (m *OrganizationalBrandingProperties) SetFavicon(value []byte)() {
    err := m.GetBackingStore().Set("favicon", value)
    if err != nil {
        panic(err)
    }
}
// SetFaviconRelativeUrl sets the faviconRelativeUrl property value. The faviconRelativeUrl property
func (m *OrganizationalBrandingProperties) SetFaviconRelativeUrl(value *string)() {
    err := m.GetBackingStore().Set("faviconRelativeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetHeaderBackgroundColor sets the headerBackgroundColor property value. The headerBackgroundColor property
func (m *OrganizationalBrandingProperties) SetHeaderBackgroundColor(value *string)() {
    err := m.GetBackingStore().Set("headerBackgroundColor", value)
    if err != nil {
        panic(err)
    }
}
// SetHeaderLogo sets the headerLogo property value. The headerLogo property
func (m *OrganizationalBrandingProperties) SetHeaderLogo(value []byte)() {
    err := m.GetBackingStore().Set("headerLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetHeaderLogoRelativeUrl sets the headerLogoRelativeUrl property value. The headerLogoRelativeUrl property
func (m *OrganizationalBrandingProperties) SetHeaderLogoRelativeUrl(value *string)() {
    err := m.GetBackingStore().Set("headerLogoRelativeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetLoginPageLayoutConfiguration sets the loginPageLayoutConfiguration property value. The loginPageLayoutConfiguration property
func (m *OrganizationalBrandingProperties) SetLoginPageLayoutConfiguration(value LoginPageLayoutConfigurationable)() {
    err := m.GetBackingStore().Set("loginPageLayoutConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetLoginPageTextVisibilitySettings sets the loginPageTextVisibilitySettings property value. The loginPageTextVisibilitySettings property
func (m *OrganizationalBrandingProperties) SetLoginPageTextVisibilitySettings(value LoginPageTextVisibilitySettingsable)() {
    err := m.GetBackingStore().Set("loginPageTextVisibilitySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInPageText sets the signInPageText property value. Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters.
func (m *OrganizationalBrandingProperties) SetSignInPageText(value *string)() {
    err := m.GetBackingStore().Set("signInPageText", value)
    if err != nil {
        panic(err)
    }
}
// SetSquareLogo sets the squareLogo property value. A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo.
func (m *OrganizationalBrandingProperties) SetSquareLogo(value []byte)() {
    err := m.GetBackingStore().Set("squareLogo", value)
    if err != nil {
        panic(err)
    }
}
// SetSquareLogoDark sets the squareLogoDark property value. The squareLogoDark property
func (m *OrganizationalBrandingProperties) SetSquareLogoDark(value []byte)() {
    err := m.GetBackingStore().Set("squareLogoDark", value)
    if err != nil {
        panic(err)
    }
}
// SetSquareLogoDarkRelativeUrl sets the squareLogoDarkRelativeUrl property value. The squareLogoDarkRelativeUrl property
func (m *OrganizationalBrandingProperties) SetSquareLogoDarkRelativeUrl(value *string)() {
    err := m.GetBackingStore().Set("squareLogoDarkRelativeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetSquareLogoRelativeUrl sets the squareLogoRelativeUrl property value. A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only.
func (m *OrganizationalBrandingProperties) SetSquareLogoRelativeUrl(value *string)() {
    err := m.GetBackingStore().Set("squareLogoRelativeUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetUsernameHintText sets the usernameHintText property value. String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can't exceed 64 characters.
func (m *OrganizationalBrandingProperties) SetUsernameHintText(value *string)() {
    err := m.GetBackingStore().Set("usernameHintText", value)
    if err != nil {
        panic(err)
    }
}
// OrganizationalBrandingPropertiesable 
type OrganizationalBrandingPropertiesable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackgroundColor()(*string)
    GetBackgroundImage()([]byte)
    GetBackgroundImageRelativeUrl()(*string)
    GetBannerLogo()([]byte)
    GetBannerLogoRelativeUrl()(*string)
    GetCdnList()([]string)
    GetCustomAccountResetCredentialsUrl()(*string)
    GetCustomCannotAccessYourAccountText()(*string)
    GetCustomCannotAccessYourAccountUrl()(*string)
    GetCustomCSS()([]byte)
    GetCustomCSSRelativeUrl()(*string)
    GetCustomForgotMyPasswordText()(*string)
    GetCustomPrivacyAndCookiesText()(*string)
    GetCustomPrivacyAndCookiesUrl()(*string)
    GetCustomResetItNowText()(*string)
    GetCustomTermsOfUseText()(*string)
    GetCustomTermsOfUseUrl()(*string)
    GetFavicon()([]byte)
    GetFaviconRelativeUrl()(*string)
    GetHeaderBackgroundColor()(*string)
    GetHeaderLogo()([]byte)
    GetHeaderLogoRelativeUrl()(*string)
    GetLoginPageLayoutConfiguration()(LoginPageLayoutConfigurationable)
    GetLoginPageTextVisibilitySettings()(LoginPageTextVisibilitySettingsable)
    GetSignInPageText()(*string)
    GetSquareLogo()([]byte)
    GetSquareLogoDark()([]byte)
    GetSquareLogoDarkRelativeUrl()(*string)
    GetSquareLogoRelativeUrl()(*string)
    GetUsernameHintText()(*string)
    SetBackgroundColor(value *string)()
    SetBackgroundImage(value []byte)()
    SetBackgroundImageRelativeUrl(value *string)()
    SetBannerLogo(value []byte)()
    SetBannerLogoRelativeUrl(value *string)()
    SetCdnList(value []string)()
    SetCustomAccountResetCredentialsUrl(value *string)()
    SetCustomCannotAccessYourAccountText(value *string)()
    SetCustomCannotAccessYourAccountUrl(value *string)()
    SetCustomCSS(value []byte)()
    SetCustomCSSRelativeUrl(value *string)()
    SetCustomForgotMyPasswordText(value *string)()
    SetCustomPrivacyAndCookiesText(value *string)()
    SetCustomPrivacyAndCookiesUrl(value *string)()
    SetCustomResetItNowText(value *string)()
    SetCustomTermsOfUseText(value *string)()
    SetCustomTermsOfUseUrl(value *string)()
    SetFavicon(value []byte)()
    SetFaviconRelativeUrl(value *string)()
    SetHeaderBackgroundColor(value *string)()
    SetHeaderLogo(value []byte)()
    SetHeaderLogoRelativeUrl(value *string)()
    SetLoginPageLayoutConfiguration(value LoginPageLayoutConfigurationable)()
    SetLoginPageTextVisibilitySettings(value LoginPageTextVisibilitySettingsable)()
    SetSignInPageText(value *string)()
    SetSquareLogo(value []byte)()
    SetSquareLogoDark(value []byte)()
    SetSquareLogoDarkRelativeUrl(value *string)()
    SetSquareLogoRelativeUrl(value *string)()
    SetUsernameHintText(value *string)()
}
