package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// identitySecurityDefaultsEnforcementPolicy 
type IdentitySecurityDefaultsEnforcementPolicy struct {
    PolicyBase
    // If set to true, Azure Active Directory security defaults is enabled for the tenant.
    isEnabled *bool;
}
// NewIdentitySecurityDefaultsEnforcementPolicy instantiates a new identitySecurityDefaultsEnforcementPolicy and sets the default values.
func NewIdentitySecurityDefaultsEnforcementPolicy()(*IdentitySecurityDefaultsEnforcementPolicy) {
    m := &IdentitySecurityDefaultsEnforcementPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    return m
}
// GetIsEnabled gets the isEnabled property value. If set to true, Azure Active Directory security defaults is enabled for the tenant.
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentitySecurityDefaultsEnforcementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    return res
}
func (m *IdentitySecurityDefaultsEnforcementPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetIsEnabled sets the isEnabled property value. If set to true, Azure Active Directory security defaults is enabled for the tenant.
func (m *IdentitySecurityDefaultsEnforcementPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
