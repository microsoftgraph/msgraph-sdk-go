package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SignInFrequencySessionControl 
type SignInFrequencySessionControl struct {
    ConditionalAccessSessionControl
    // Possible values are: days, hours.
    type_escaped *SigninFrequencyType;
    // The number of days or hours.
    value *int32;
}
// NewSignInFrequencySessionControl instantiates a new signInFrequencySessionControl and sets the default values.
func NewSignInFrequencySessionControl()(*SignInFrequencySessionControl) {
    m := &SignInFrequencySessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
// GetType_escaped gets the type property value. Possible values are: days, hours.
func (m *SignInFrequencySessionControl) GetType_escaped()(*SigninFrequencyType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetValue gets the value property value. The number of days or hours.
func (m *SignInFrequencySessionControl) GetValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SignInFrequencySessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSigninFrequencyType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SigninFrequencyType)
            m.SetType(&cast)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
func (m *SignInFrequencySessionControl) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SignInFrequencySessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetType() != nil {
        cast := m.GetType().String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetType_escaped sets the type property value. Possible values are: days, hours.
func (m *SignInFrequencySessionControl) SetType_escaped(value *SigninFrequencyType)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetValue sets the value property value. The number of days or hours.
func (m *SignInFrequencySessionControl) SetValue(value *int32)() {
    if m != nil {
        m.value = value
    }
}
