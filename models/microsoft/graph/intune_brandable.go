package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IntuneBrandable 
type IntuneBrandable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContactITEmailAddress()(*string)
    GetContactITName()(*string)
    GetContactITNotes()(*string)
    GetContactITPhoneNumber()(*string)
    GetDarkBackgroundLogo()(MimeContentable)
    GetDisplayName()(*string)
    GetLightBackgroundLogo()(MimeContentable)
    GetOnlineSupportSiteName()(*string)
    GetOnlineSupportSiteUrl()(*string)
    GetPrivacyUrl()(*string)
    GetShowDisplayNameNextToLogo()(*bool)
    GetShowLogo()(*bool)
    GetShowNameNextToLogo()(*bool)
    GetThemeColor()(RgbColorable)
    SetContactITEmailAddress(value *string)()
    SetContactITName(value *string)()
    SetContactITNotes(value *string)()
    SetContactITPhoneNumber(value *string)()
    SetDarkBackgroundLogo(value MimeContentable)()
    SetDisplayName(value *string)()
    SetLightBackgroundLogo(value MimeContentable)()
    SetOnlineSupportSiteName(value *string)()
    SetOnlineSupportSiteUrl(value *string)()
    SetPrivacyUrl(value *string)()
    SetShowDisplayNameNextToLogo(value *bool)()
    SetShowLogo(value *bool)()
    SetShowNameNextToLogo(value *bool)()
    SetThemeColor(value RgbColorable)()
}
