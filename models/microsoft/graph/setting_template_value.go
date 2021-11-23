package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// settingTemplateValue 
type SettingTemplateValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Default value for the setting.
    defaultValue *string;
    // Description of the setting.
    description *string;
    // Name of the setting.
    name *string;
    // Type of the setting.
    type_escaped *string;
}
// NewSettingTemplateValue instantiates a new settingTemplateValue and sets the default values.
func NewSettingTemplateValue()(*SettingTemplateValue) {
    m := &SettingTemplateValue{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingTemplateValue) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefaultValue gets the defaultValue property value. Default value for the setting.
func (m *SettingTemplateValue) GetDefaultValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetDescription gets the description property value. Description of the setting.
func (m *SettingTemplateValue) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetName gets the name property value. Name of the setting.
func (m *SettingTemplateValue) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetType_escaped gets the type_escaped property value. Type of the setting.
func (m *SettingTemplateValue) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SettingTemplateValue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
        }
        return nil
    }
    return res
}
func (m *SettingTemplateValue) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SettingTemplateValue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingTemplateValue) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultValue sets the defaultValue property value. Default value for the setting.
func (m *SettingTemplateValue) SetDefaultValue(value *string)() {
    m.defaultValue = value
}
// SetDescription sets the description property value. Description of the setting.
func (m *SettingTemplateValue) SetDescription(value *string)() {
    m.description = value
}
// SetName sets the name property value. Name of the setting.
func (m *SettingTemplateValue) SetName(value *string)() {
    m.name = value
}
// SetType_escaped sets the type_escaped property value. Type of the setting.
func (m *SettingTemplateValue) SetType_escaped(value *string)() {
    m.type_escaped = value
}
