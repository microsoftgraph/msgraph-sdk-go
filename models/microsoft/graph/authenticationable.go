package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Authenticationable 
type Authenticationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetFido2Methods()([]Fido2AuthenticationMethodable)
    GetMethods()([]AuthenticationMethodable)
    GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethodable)
    GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethodable)
    SetFido2Methods(value []Fido2AuthenticationMethodable)()
    SetMethods(value []AuthenticationMethodable)()
    SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethodable)()
    SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethodable)()
}
