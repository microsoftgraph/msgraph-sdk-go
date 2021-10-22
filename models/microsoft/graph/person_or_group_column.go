package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PersonOrGroupColumn struct {
    additionalData map[string]interface{};
    allowMultipleSelection *bool;
    chooseFromType *string;
    displayAs *string;
}
func NewPersonOrGroupColumn()(*PersonOrGroupColumn) {
    m := &PersonOrGroupColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PersonOrGroupColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PersonOrGroupColumn) GetAllowMultipleSelection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleSelection
    }
}
func (m *PersonOrGroupColumn) GetChooseFromType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.chooseFromType
    }
}
func (m *PersonOrGroupColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
func (m *PersonOrGroupColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleSelection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowMultipleSelection(val)
        return nil
    }
    res["chooseFromType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChooseFromType(val)
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
func (m *PersonOrGroupColumn) IsNil()(bool) {
    return m == nil
}
func (m *PersonOrGroupColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleSelection", m.GetAllowMultipleSelection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("chooseFromType", m.GetChooseFromType())
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
func (m *PersonOrGroupColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PersonOrGroupColumn) SetAllowMultipleSelection(value *bool)() {
    m.allowMultipleSelection = value
}
func (m *PersonOrGroupColumn) SetChooseFromType(value *string)() {
    m.chooseFromType = value
}
func (m *PersonOrGroupColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
