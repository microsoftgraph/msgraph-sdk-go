package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChoiceColumn struct {
    additionalData map[string]interface{};
    allowTextEntry *bool;
    choices []string;
    displayAs *string;
}
func NewChoiceColumn()(*ChoiceColumn) {
    m := &ChoiceColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChoiceColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChoiceColumn) GetAllowTextEntry()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowTextEntry
    }
}
func (m *ChoiceColumn) GetChoices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.choices
    }
}
func (m *ChoiceColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
func (m *ChoiceColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowTextEntry"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowTextEntry(val)
        return nil
    }
    res["choices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetChoices(res)
        return nil
    }
    res["displayAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayAs(val)
        return nil
    }
    return res
}
func (m *ChoiceColumn) IsNil()(bool) {
    return m == nil
}
func (m *ChoiceColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowTextEntry", m.GetAllowTextEntry())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("choices", m.GetChoices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
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
func (m *ChoiceColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChoiceColumn) SetAllowTextEntry(value *bool)() {
    m.allowTextEntry = value
}
func (m *ChoiceColumn) SetChoices(value []string)() {
    m.choices = value
}
func (m *ChoiceColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
