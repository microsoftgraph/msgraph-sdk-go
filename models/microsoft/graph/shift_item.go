package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ShiftItem 
type ShiftItem struct {
    ScheduleEntity
    // An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
    activities []ShiftActivity;
    // The shift label of the shiftItem.
    displayName *string;
    // The shift notes for the shiftItem.
    notes *string;
}
// NewShiftItem instantiates a new shiftItem and sets the default values.
func NewShiftItem()(*ShiftItem) {
    m := &ShiftItem{
        ScheduleEntity: *NewScheduleEntity(),
    }
    return m
}
// GetActivities gets the activities property value. An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
func (m *ShiftItem) GetActivities()([]ShiftActivity) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
// GetDisplayName gets the displayName property value. The shift label of the shiftItem.
func (m *ShiftItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetNotes gets the notes property value. The shift notes for the shiftItem.
func (m *ShiftItem) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ShiftItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ScheduleEntity.GetFieldDeserializers()
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShiftActivity() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ShiftActivity, len(val))
            for i, v := range val {
                res[i] = *(v.(*ShiftActivity))
            }
            m.SetActivities(res)
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
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    return res
}
func (m *ShiftItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ShiftItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ScheduleEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("activities", cast)
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
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivities sets the activities property value. An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
func (m *ShiftItem) SetActivities(value []ShiftActivity)() {
    if m != nil {
        m.activities = value
    }
}
// SetDisplayName sets the displayName property value. The shift label of the shiftItem.
func (m *ShiftItem) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetNotes sets the notes property value. The shift notes for the shiftItem.
func (m *ShiftItem) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
