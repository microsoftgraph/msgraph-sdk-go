package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupSetting 
type GroupSetting struct {
    Entity
    // Display name of this group of settings, which comes from the associated template.
    displayName *string;
    // Unique identifier for the template used to create this group of settings. Read-only.
    templateId *string;
    // Collection of name value pairs. Must contain and set all the settings defined in the template.
    values []SettingValue;
}
// NewGroupSetting instantiates a new groupSetting and sets the default values.
func NewGroupSetting()(*GroupSetting) {
    m := &GroupSetting{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Display name of this group of settings, which comes from the associated template.
func (m *GroupSetting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetTemplateId gets the templateId property value. Unique identifier for the template used to create this group of settings. Read-only.
func (m *GroupSetting) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// GetValues gets the values property value. Collection of name value pairs. Must contain and set all the settings defined in the template.
func (m *GroupSetting) GetValues()([]SettingValue) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupSetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingValue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingValue, len(val))
            for i, v := range val {
                res[i] = *(v.(*SettingValue))
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
func (m *GroupSetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetValues() != nil {
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
// SetDisplayName sets the displayName property value. Display name of this group of settings, which comes from the associated template.
func (m *GroupSetting) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetTemplateId sets the templateId property value. Unique identifier for the template used to create this group of settings. Read-only.
func (m *GroupSetting) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
// SetValues sets the values property value. Collection of name value pairs. Must contain and set all the settings defined in the template.
func (m *GroupSetting) SetValues(value []SettingValue)() {
    if m != nil {
        m.values = value
    }
}
