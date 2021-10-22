package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SignInFrequencySessionControl struct {
    ConditionalAccessSessionControl
    type_escpaped *SigninFrequencyType;
    value *int32;
}
func NewSignInFrequencySessionControl()(*SignInFrequencySessionControl) {
    m := &SignInFrequencySessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
func (m *SignInFrequencySessionControl) GetType_escpaped()(*SigninFrequencyType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *SignInFrequencySessionControl) GetValue()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *SignInFrequencySessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSigninFrequencyType)
        if err != nil {
            return err
        }
        cast := val.(SigninFrequencyType)
        m.SetType_escpaped(&cast)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *SignInFrequencySessionControl) IsNil()(bool) {
    return m == nil
}
func (m *SignInFrequencySessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err = writer.WriteStringValue("type_escpaped", &cast)
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
func (m *SignInFrequencySessionControl) SetType_escpaped(value *SigninFrequencyType)() {
    m.type_escpaped = value
}
func (m *SignInFrequencySessionControl) SetValue(value *int32)() {
    m.value = value
}
