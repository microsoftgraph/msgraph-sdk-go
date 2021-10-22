package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserFlowLanguageConfiguration struct {
    Entity
    defaultPages []UserFlowLanguagePage;
    displayName *string;
    isEnabled *bool;
    overridesPages []UserFlowLanguagePage;
}
func NewUserFlowLanguageConfiguration()(*UserFlowLanguageConfiguration) {
    m := &UserFlowLanguageConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserFlowLanguageConfiguration) GetDefaultPages()([]UserFlowLanguagePage) {
    if m == nil {
        return nil
    } else {
        return m.defaultPages
    }
}
func (m *UserFlowLanguageConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *UserFlowLanguageConfiguration) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
func (m *UserFlowLanguageConfiguration) GetOverridesPages()([]UserFlowLanguagePage) {
    if m == nil {
        return nil
    } else {
        return m.overridesPages
    }
}
func (m *UserFlowLanguageConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultPages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowLanguagePage() })
        if err != nil {
            return err
        }
        res := make([]UserFlowLanguagePage, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserFlowLanguagePage))
        }
        m.SetDefaultPages(res)
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
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["overridesPages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowLanguagePage() })
        if err != nil {
            return err
        }
        res := make([]UserFlowLanguagePage, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserFlowLanguagePage))
        }
        m.SetOverridesPages(res)
        return nil
    }
    return res
}
func (m *UserFlowLanguageConfiguration) IsNil()(bool) {
    return m == nil
}
func (m *UserFlowLanguageConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefaultPages()))
        for i, v := range m.GetDefaultPages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("defaultPages", cast)
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
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOverridesPages()))
        for i, v := range m.GetOverridesPages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("overridesPages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserFlowLanguageConfiguration) SetDefaultPages(value []UserFlowLanguagePage)() {
    m.defaultPages = value
}
func (m *UserFlowLanguageConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *UserFlowLanguageConfiguration) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
func (m *UserFlowLanguageConfiguration) SetOverridesPages(value []UserFlowLanguagePage)() {
    m.overridesPages = value
}
