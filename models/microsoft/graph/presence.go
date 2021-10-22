package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Presence struct {
    Entity
    activity *string;
    availability *string;
}
func NewPresence()(*Presence) {
    m := &Presence{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Presence) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
func (m *Presence) GetAvailability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availability
    }
}
func (m *Presence) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivity(val)
        return nil
    }
    res["availability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAvailability(val)
        return nil
    }
    return res
}
func (m *Presence) IsNil()(bool) {
    return m == nil
}
func (m *Presence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("availability", m.GetAvailability())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Presence) SetActivity(value *string)() {
    m.activity = value
}
func (m *Presence) SetAvailability(value *string)() {
    m.availability = value
}
