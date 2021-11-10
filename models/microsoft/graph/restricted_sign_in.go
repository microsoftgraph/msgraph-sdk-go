package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RestrictedSignIn struct {
    SignIn
    // 
    targetTenantId *string;
}
// Instantiates a new restrictedSignIn and sets the default values.
func NewRestrictedSignIn()(*RestrictedSignIn) {
    m := &RestrictedSignIn{
        SignIn: *NewSignIn(),
    }
    return m
}
// Gets the targetTenantId property value. 
func (m *RestrictedSignIn) GetTargetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetTenantId
    }
}
// The deserialization information for the current model
func (m *RestrictedSignIn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.SignIn.GetFieldDeserializers()
    res["targetTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetTenantId(val)
        }
        return nil
    }
    return res
}
func (m *RestrictedSignIn) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RestrictedSignIn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.SignIn.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetTenantId", m.GetTargetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the targetTenantId property value. 
// Parameters:
//  - value : Value to set for the targetTenantId property.
func (m *RestrictedSignIn) SetTargetTenantId(value *string)() {
    m.targetTenantId = value
}
