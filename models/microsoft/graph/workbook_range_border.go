package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookRangeBorder struct {
    Entity
    color *string;
    sideIndex *string;
    style *string;
    weight *string;
}
func NewWorkbookRangeBorder()(*WorkbookRangeBorder) {
    m := &WorkbookRangeBorder{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookRangeBorder) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
func (m *WorkbookRangeBorder) GetSideIndex()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sideIndex
    }
}
func (m *WorkbookRangeBorder) GetStyle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.style
    }
}
func (m *WorkbookRangeBorder) GetWeight()(*string) {
    if m == nil {
        return nil
    } else {
        return m.weight
    }
}
func (m *WorkbookRangeBorder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColor(val)
        return nil
    }
    res["sideIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSideIndex(val)
        return nil
    }
    res["style"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStyle(val)
        return nil
    }
    res["weight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWeight(val)
        return nil
    }
    return res
}
func (m *WorkbookRangeBorder) IsNil()(bool) {
    return m == nil
}
func (m *WorkbookRangeBorder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sideIndex", m.GetSideIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("style", m.GetStyle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("weight", m.GetWeight())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *WorkbookRangeBorder) SetColor(value *string)() {
    m.color = value
}
func (m *WorkbookRangeBorder) SetSideIndex(value *string)() {
    m.sideIndex = value
}
func (m *WorkbookRangeBorder) SetStyle(value *string)() {
    m.style = value
}
func (m *WorkbookRangeBorder) SetWeight(value *string)() {
    m.weight = value
}
