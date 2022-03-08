package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OrganizationalBrandingPropertiesable 
type OrganizationalBrandingPropertiesable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetBackgroundColor()(*string)
    GetBackgroundImage()([]byte)
    GetBackgroundImageRelativeUrl()(*string)
    GetBannerLogo()([]byte)
    GetBannerLogoRelativeUrl()(*string)
    GetCdnList()([]string)
    GetSignInPageText()(*string)
    GetSquareLogo()([]byte)
    GetSquareLogoRelativeUrl()(*string)
    GetUsernameHintText()(*string)
    SetBackgroundColor(value *string)()
    SetBackgroundImage(value []byte)()
    SetBackgroundImageRelativeUrl(value *string)()
    SetBannerLogo(value []byte)()
    SetBannerLogoRelativeUrl(value *string)()
    SetCdnList(value []string)()
    SetSignInPageText(value *string)()
    SetSquareLogo(value []byte)()
    SetSquareLogoRelativeUrl(value *string)()
    SetUsernameHintText(value *string)()
}
