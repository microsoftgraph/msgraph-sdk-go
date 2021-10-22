package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConditionalAccessSessionControl struct {
    additionalData map[string]interface{};
    isEnabled *bool;
}
func NewConditionalAccessSessionControl()(*ConditionalAccessSessionControl) {
    m := &ConditionalAccessSessionControl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConditionalAccessSessionControl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConditionalAccessSessionControl) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
func (m *ConditionalAccessSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
func (m *ConditionalAccessSessionControl) IsNil()(bool) {
    return m == nil
}
func (m *ConditionalAccessSessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ConditionalAccessSessionControl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConditionalAccessSessionControl) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
