package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CalendarGroup 
type CalendarGroup struct {
    Entity
    // The calendars in the calendar group. Navigation property. Read-only. Nullable.
    calendars []Calendar;
    // Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
    changeKey *string;
    // The class identifier. Read-only.
    classId *string;
    // The group name.
    name *string;
}
// NewCalendarGroup instantiates a new calendarGroup and sets the default values.
func NewCalendarGroup()(*CalendarGroup) {
    m := &CalendarGroup{
        Entity: *NewEntity(),
    }
    return m
}
// GetCalendars gets the calendars property value. The calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarGroup) GetCalendars()([]Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendars
    }
}
// GetChangeKey gets the changeKey property value. Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
func (m *CalendarGroup) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
// GetClassId gets the classId property value. The class identifier. Read-only.
func (m *CalendarGroup) GetClassId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classId
    }
}
// GetName gets the name property value. The group name.
func (m *CalendarGroup) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalendarGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calendars"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendar() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Calendar, len(val))
            for i, v := range val {
                res[i] = *(v.(*Calendar))
            }
            m.SetCalendars(res)
        }
        return nil
    }
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["classId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassId(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
func (m *CalendarGroup) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CalendarGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCalendars() != nil {
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
// SetCalendars sets the calendars property value. The calendars in the calendar group. Navigation property. Read-only. Nullable.
func (m *CalendarGroup) SetCalendars(value []Calendar)() {
    if m != nil {
        m.calendars = value
    }
}
// SetChangeKey sets the changeKey property value. Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
func (m *CalendarGroup) SetChangeKey(value *string)() {
    if m != nil {
        m.changeKey = value
    }
}
// SetClassId sets the classId property value. The class identifier. Read-only.
func (m *CalendarGroup) SetClassId(value *string)() {
    if m != nil {
        m.classId = value
    }
}
// SetName sets the name property value. The group name.
func (m *CalendarGroup) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
