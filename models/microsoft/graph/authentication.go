package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Authentication 
type Authentication struct {
    Entity
    // 
    fido2Methods []Fido2AuthenticationMethod;
    // 
    methods []AuthenticationMethod;
    // 
    microsoftAuthenticatorMethods []MicrosoftAuthenticatorAuthenticationMethod;
    // 
    windowsHelloForBusinessMethods []WindowsHelloForBusinessAuthenticationMethod;
}
// NewAuthentication instantiates a new authentication and sets the default values.
func NewAuthentication()(*Authentication) {
    m := &Authentication{
        Entity: *NewEntity(),
    }
    return m
}
// GetFido2Methods gets the fido2Methods property value. 
func (m *Authentication) GetFido2Methods()([]Fido2AuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.fido2Methods
    }
}
// GetMethods gets the methods property value. 
func (m *Authentication) GetMethods()([]AuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.methods
    }
}
// GetMicrosoftAuthenticatorMethods gets the microsoftAuthenticatorMethods property value. 
func (m *Authentication) GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.microsoftAuthenticatorMethods
    }
}
// GetWindowsHelloForBusinessMethods gets the windowsHelloForBusinessMethods property value. 
func (m *Authentication) GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethod) {
    if m == nil {
        return nil
    } else {
        return m.windowsHelloForBusinessMethods
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Authentication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fido2Methods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFido2AuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fido2AuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*Fido2AuthenticationMethod))
            }
            m.SetFido2Methods(res)
        }
        return nil
    }
    res["methods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*AuthenticationMethod))
            }
            m.SetMethods(res)
        }
        return nil
    }
    res["microsoftAuthenticatorMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftAuthenticatorAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftAuthenticatorAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*MicrosoftAuthenticatorAuthenticationMethod))
            }
            m.SetMicrosoftAuthenticatorMethods(res)
        }
        return nil
    }
    res["windowsHelloForBusinessMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsHelloForBusinessAuthenticationMethod() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsHelloForBusinessAuthenticationMethod, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsHelloForBusinessAuthenticationMethod))
            }
            m.SetWindowsHelloForBusinessMethods(res)
        }
        return nil
    }
    return res
}
func (m *Authentication) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Authentication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFido2Methods()))
        for i, v := range m.GetFido2Methods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("fido2Methods", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMethods()))
        for i, v := range m.GetMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("methods", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrosoftAuthenticatorMethods()))
        for i, v := range m.GetMicrosoftAuthenticatorMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftAuthenticatorMethods", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsHelloForBusinessMethods()))
        for i, v := range m.GetWindowsHelloForBusinessMethods() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsHelloForBusinessMethods", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFido2Methods sets the fido2Methods property value. 
func (m *Authentication) SetFido2Methods(value []Fido2AuthenticationMethod)() {
    m.fido2Methods = value
}
// SetMethods sets the methods property value. 
func (m *Authentication) SetMethods(value []AuthenticationMethod)() {
    m.methods = value
}
// SetMicrosoftAuthenticatorMethods sets the microsoftAuthenticatorMethods property value. 
func (m *Authentication) SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethod)() {
    m.microsoftAuthenticatorMethods = value
}
// SetWindowsHelloForBusinessMethods sets the windowsHelloForBusinessMethods property value. 
func (m *Authentication) SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethod)() {
    m.windowsHelloForBusinessMethods = value
}
