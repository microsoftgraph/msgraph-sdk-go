package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Authentication 
type Authentication struct {
    Entity
    // Represents the FIDO2 security keys registered to a user for authentication.
    fido2Methods []Fido2AuthenticationMethodable;
    // Represents all authentication methods registered to a user.
    methods []AuthenticationMethodable;
    // The details of the Microsoft Authenticator app registered to a user for authentication.
    microsoftAuthenticatorMethods []MicrosoftAuthenticatorAuthenticationMethodable;
    // Represents the Windows Hello for Business authentication method registered to a user for authentication.
    windowsHelloForBusinessMethods []WindowsHelloForBusinessAuthenticationMethodable;
}
// NewAuthentication instantiates a new authentication and sets the default values.
func NewAuthentication()(*Authentication) {
    m := &Authentication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthentication(), nil
}
// GetFido2Methods gets the fido2Methods property value. Represents the FIDO2 security keys registered to a user for authentication.
func (m *Authentication) GetFido2Methods()([]Fido2AuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.fido2Methods
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Authentication) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fido2Methods"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFido2AuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fido2AuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(Fido2AuthenticationMethodable)
            }
            m.SetFido2Methods(res)
        }
        return nil
    }
    res["methods"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodable)
            }
            m.SetMethods(res)
        }
        return nil
    }
    res["microsoftAuthenticatorMethods"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftAuthenticatorAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftAuthenticatorAuthenticationMethodable)
            }
            m.SetMicrosoftAuthenticatorMethods(res)
        }
        return nil
    }
    res["windowsHelloForBusinessMethods"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsHelloForBusinessAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsHelloForBusinessAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsHelloForBusinessAuthenticationMethodable)
            }
            m.SetWindowsHelloForBusinessMethods(res)
        }
        return nil
    }
    return res
}
// GetMethods gets the methods property value. Represents all authentication methods registered to a user.
func (m *Authentication) GetMethods()([]AuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.methods
    }
}
// GetMicrosoftAuthenticatorMethods gets the microsoftAuthenticatorMethods property value. The details of the Microsoft Authenticator app registered to a user for authentication.
func (m *Authentication) GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.microsoftAuthenticatorMethods
    }
}
// GetWindowsHelloForBusinessMethods gets the windowsHelloForBusinessMethods property value. Represents the Windows Hello for Business authentication method registered to a user for authentication.
func (m *Authentication) GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethodable) {
    if m == nil {
        return nil
    } else {
        return m.windowsHelloForBusinessMethods
    }
}
// Serialize serializes information the current object
func (m *Authentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFido2Methods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFido2Methods()))
        for i, v := range m.GetFido2Methods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("fido2Methods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMethods()))
        for i, v := range m.GetMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("methods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftAuthenticatorMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftAuthenticatorMethods()))
        for i, v := range m.GetMicrosoftAuthenticatorMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftAuthenticatorMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsHelloForBusinessMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsHelloForBusinessMethods()))
        for i, v := range m.GetWindowsHelloForBusinessMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsHelloForBusinessMethods", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFido2Methods sets the fido2Methods property value. Represents the FIDO2 security keys registered to a user for authentication.
func (m *Authentication) SetFido2Methods(value []Fido2AuthenticationMethodable)() {
    if m != nil {
        m.fido2Methods = value
    }
}
// SetMethods sets the methods property value. Represents all authentication methods registered to a user.
func (m *Authentication) SetMethods(value []AuthenticationMethodable)() {
    if m != nil {
        m.methods = value
    }
}
// SetMicrosoftAuthenticatorMethods sets the microsoftAuthenticatorMethods property value. The details of the Microsoft Authenticator app registered to a user for authentication.
func (m *Authentication) SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethodable)() {
    if m != nil {
        m.microsoftAuthenticatorMethods = value
    }
}
// SetWindowsHelloForBusinessMethods sets the windowsHelloForBusinessMethods property value. Represents the Windows Hello for Business authentication method registered to a user for authentication.
func (m *Authentication) SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethodable)() {
    if m != nil {
        m.windowsHelloForBusinessMethods = value
    }
}
