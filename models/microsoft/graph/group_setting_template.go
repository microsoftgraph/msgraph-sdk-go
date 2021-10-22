package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GroupSettingTemplate struct {
    DirectoryObject
    description *string;
    displayName *string;
    values []SettingTemplateValue;
}
func NewGroupSettingTemplate()(*GroupSettingTemplate) {
    m := &GroupSettingTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
func (m *GroupSettingTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *GroupSettingTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GroupSettingTemplate) GetValues()([]SettingTemplateValue) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
func (m *GroupSettingTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingTemplateValue() })
        if err != nil {
            return err
        }
        res := make([]SettingTemplateValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*SettingTemplateValue))
        }
        m.SetValues(res)
        return nil
    }
    return res
}
func (m *GroupSettingTemplate) IsNil()(bool) {
    return m == nil
}
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
func (m *GroupSettingTemplate) SetDescription(value *string)() {
    m.description = value
}
func (m *GroupSettingTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GroupSettingTemplate) SetValues(value []SettingTemplateValue)() {
    m.values = value
}
