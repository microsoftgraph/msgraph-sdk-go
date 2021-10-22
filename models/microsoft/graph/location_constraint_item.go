package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LocationConstraintItem struct {
    Location
    resolveAvailability *bool;
}
func NewLocationConstraintItem()(*LocationConstraintItem) {
    m := &LocationConstraintItem{
        Location: *NewLocation(),
    }
    return m
}
func (m *LocationConstraintItem) GetResolveAvailability()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.resolveAvailability
    }
}
func (m *LocationConstraintItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Location.GetFieldDeserializers()
    res["resolveAvailability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetResolveAvailability(val)
        return nil
    }
    return res
}
func (m *LocationConstraintItem) IsNil()(bool) {
    return m == nil
}
func (m *LocationConstraintItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Location.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("resolveAvailability", m.GetResolveAvailability())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *LocationConstraintItem) SetResolveAvailability(value *bool)() {
    m.resolveAvailability = value
}
