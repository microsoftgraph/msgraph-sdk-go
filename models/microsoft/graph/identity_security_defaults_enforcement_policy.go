package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IdentitySecurityDefaultsEnforcementPolicy struct {
    PolicyBase
    // If set to true, Azure Active Directory security defaults is enabled for the tenant.
    isEnabled *bool;
}
// Instantiates a new identitySecurityDefaultsEnforcementPolicy and sets the default values.
func NewIdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicy) {
    m := &IdentitySecurityDefaultsEnforcementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// Gets the isEnabled property value. If set to true, Azure Active Directory security defaults is enabled for the tenant.
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// The deserialization information for the current model
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    return res
}
func (m *IdentitySecurityDefaultsEnforcementPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IdentitySecurityDefaultsEnforcementPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the isEnabled property value. If set to true, Azure Active Directory security defaults is enabled for the tenant.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
