package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GroupSetting struct {
    Entity
    // Display name of this group of settings, which comes from the associated template.
    displayName *string;
    // Unique identifier for the template used to create this group of settings. Read-only.
    templateId *string;
    // Collection of name value pairs. Must contain and set all the settings defined in the template.
    values []SettingValue;
}
// Instantiates a new groupSetting and sets the default values.
func NewGroupSetting()(*GroupSetting) {
    m := &GroupSetting{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. Display name of this group of settings, which comes from the associated template.
func (m *GroupSetting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the templateId property value. Unique identifier for the template used to create this group of settings. Read-only.
func (m *GroupSetting) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// Gets the values property value. Collection of name value pairs. Must contain and set all the settings defined in the template.
func (m *GroupSetting) GetValues()([]SettingValue) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// The deserialization information for the current model
func (m *GroupSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTemplateId(val)
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingValue() })
        if err != nil {
            return err
        }
        res := make([]SettingValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*SettingValue))
        }
        m.SetValues(res)
        return nil
    }
    return res
}
func (m *GroupSetting) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GroupSetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("values", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. Display name of this group of settings, which comes from the associated template.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GroupSetting) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the templateId property value. Unique identifier for the template used to create this group of settings. Read-only.
// Parameters:
//  - value : Value to set for the templateId property.
func (m *GroupSetting) SetTemplateId(value *string)() {
    m.templateId = value
}
// Sets the values property value. Collection of name value pairs. Must contain and set all the settings defined in the template.
// Parameters:
//  - value : Value to set for the values property.
func (m *GroupSetting) SetValues(value []SettingValue)() {
    m.values = value
}
