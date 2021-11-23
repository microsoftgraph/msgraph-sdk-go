package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserFlowLanguageConfiguration 
type UserFlowLanguageConfiguration struct {
    Entity
    // Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
    defaultPages []UserFlowLanguagePage;
    // The language name to display. This property is read-only.
    displayName *string;
    // Indicates whether the language is enabled within the user flow.
    isEnabled *bool;
    // Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
    overridesPages []UserFlowLanguagePage;
}
// NewUserFlowLanguageConfiguration instantiates a new userFlowLanguageConfiguration and sets the default values.
func NewUserFlowLanguageConfiguration()(*UserFlowLanguageConfiguration) {
    m := &UserFlowLanguageConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// GetDefaultPages gets the defaultPages property value. Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
func (m *UserFlowLanguageConfiguration) GetDefaultPages()([]UserFlowLanguagePage) {
    if m == nil {
        return nil
    } else {
        return m.defaultPages
    }
}
// GetDisplayName gets the displayName property value. The language name to display. This property is read-only.
func (m *UserFlowLanguageConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the language is enabled within the user flow.
func (m *UserFlowLanguageConfiguration) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetOverridesPages gets the overridesPages property value. Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
func (m *UserFlowLanguageConfiguration) GetOverridesPages()([]UserFlowLanguagePage) {
    if m == nil {
        return nil
    } else {
        return m.overridesPages
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserFlowLanguageConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultPages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowLanguagePage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguagePage, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserFlowLanguagePage))
            }
            m.SetDefaultPages(res)
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
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["overridesPages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserFlowLanguagePage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguagePage, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserFlowLanguagePage))
            }
            m.SetOverridesPages(res)
        }
        return nil
    }
    return res
}
func (m *UserFlowLanguageConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetDefaultPages sets the defaultPages property value. Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
func (m *UserFlowLanguageConfiguration) SetDefaultPages(value []UserFlowLanguagePage)() {
    m.defaultPages = value
}
// SetDisplayName sets the displayName property value. The language name to display. This property is read-only.
func (m *UserFlowLanguageConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the language is enabled within the user flow.
func (m *UserFlowLanguageConfiguration) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetOverridesPages sets the overridesPages property value. Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
func (m *UserFlowLanguageConfiguration) SetOverridesPages(value []UserFlowLanguagePage)() {
    m.overridesPages = value
}
