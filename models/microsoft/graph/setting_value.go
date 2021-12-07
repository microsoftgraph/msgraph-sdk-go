package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SettingValue 
type SettingValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the setting (as defined by the groupSettingTemplate).
    name *string;
    // Value of the setting.
    value *string;
}
// NewSettingValue instantiates a new settingValue and sets the default values.
func NewSettingValue()(*SettingValue) {
    m := &SettingValue{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingValue) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetName gets the name property value. Name of the setting (as defined by the groupSettingTemplate).
func (m *SettingValue) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetValue gets the value property value. Value of the setting.
func (m *SettingValue) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SettingValue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
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
func (m *SettingValue) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SettingValue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingValue) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Name of the setting (as defined by the groupSettingTemplate).
func (m *SettingValue) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetValue sets the value property value. Value of the setting.
func (m *SettingValue) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
