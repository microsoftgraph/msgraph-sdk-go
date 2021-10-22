package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OutlookCategory struct {
    Entity
    color *CategoryColor;
    displayName *string;
}
func NewOutlookCategory()(*OutlookCategory) {
    m := &OutlookCategory{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OutlookCategory) GetColor()(*CategoryColor) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
func (m *OutlookCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *OutlookCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCategoryColor)
        if err != nil {
            return err
        }
        cast := val.(CategoryColor)
        m.SetColor(&cast)
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
    return res
}
func (m *OutlookCategory) IsNil()(bool) {
    return m == nil
}
func (m *OutlookCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColor() != nil {
        cast := m.GetColor().String()
        err = writer.WriteStringValue("color", &cast)
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
    return nil
}
func (m *OutlookCategory) SetColor(value *CategoryColor)() {
    m.color = value
}
func (m *OutlookCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
