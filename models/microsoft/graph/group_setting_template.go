package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupSettingTemplate 
type GroupSettingTemplate struct {
    DirectoryObject
    // Description of the template.
    description *string;
    // Display name of the template.
    displayName *string;
    // Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.
    values []SettingTemplateValue;
}
// NewGroupSettingTemplate instantiates a new groupSettingTemplate and sets the default values.
func NewGroupSettingTemplate()(*GroupSettingTemplate) {
    m := &GroupSettingTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetDescription gets the description property value. Description of the template.
func (m *GroupSettingTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name of the template.
func (m *GroupSettingTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetValues gets the values property value. Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.
func (m *GroupSettingTemplate) GetValues()([]SettingTemplateValue) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupSettingTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingTemplateValue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingTemplateValue, len(val))
            for i, v := range val {
                res[i] = *(v.(*SettingTemplateValue))
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
func (m *GroupSettingTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupSettingTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
// SetDescription sets the description property value. Description of the template.
func (m *GroupSettingTemplate) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the template.
func (m *GroupSettingTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetValues sets the values property value. Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.
func (m *GroupSettingTemplate) SetValues(value []SettingTemplateValue)() {
    if m != nil {
        m.values = value
    }
}
