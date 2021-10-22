package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CalendarGroup struct {
    Entity
    calendars []Calendar;
    changeKey *string;
    classId *string;
    name *string;
}
func NewCalendarGroup()(*CalendarGroup) {
    m := &CalendarGroup{
        Entity: *NewEntity(),
    }
    return m
}
func (m *CalendarGroup) GetCalendars()([]Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendars
    }
}
func (m *CalendarGroup) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
func (m *CalendarGroup) GetClassId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classId
    }
}
func (m *CalendarGroup) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *CalendarGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calendars"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendar() })
        if err != nil {
            return err
        }
        res := make([]Calendar, len(val))
        for i, v := range val {
            res[i] = *(v.(*Calendar))
        }
        m.SetCalendars(res)
        return nil
    }
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChangeKey(val)
        return nil
    }
    res["classId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClassId(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    return res
}
func (m *CalendarGroup) IsNil()(bool) {
    return m == nil
}
func (m *CalendarGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendars()))
        for i, v := range m.GetCalendars() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendars", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("classId", m.GetClassId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *CalendarGroup) SetCalendars(value []Calendar)() {
    m.calendars = value
}
func (m *CalendarGroup) SetChangeKey(value *string)() {
    m.changeKey = value
}
func (m *CalendarGroup) SetClassId(value *string)() {
    m.classId = value
}
func (m *CalendarGroup) SetName(value *string)() {
    m.name = value
}
