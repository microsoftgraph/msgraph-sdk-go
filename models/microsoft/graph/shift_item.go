package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ShiftItem struct {
    ScheduleEntity
    // An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
    activities []ShiftActivity;
    // The shift label of the shiftItem.
    displayName *string;
    // The shift notes for the shiftItem.
    notes *string;
}
// Instantiates a new shiftItem and sets the default values.
func NewShiftItem()(*ShiftItem) {
    m := &ShiftItem{
        ScheduleEntity: *NewScheduleEntity(),
    }
    return m
}
// Gets the activities property value. An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
func (m *ShiftItem) GetActivities()([]ShiftActivity) {
    if m == nil {
        return nil
    } else {
        return m.activities
    }
}
// Gets the displayName property value. The shift label of the shiftItem.
func (m *ShiftItem) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the notes property value. The shift notes for the shiftItem.
func (m *ShiftItem) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// The deserialization information for the current model
func (m *ShiftItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ScheduleEntity.GetFieldDeserializers()
    res["activities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShiftActivity() })
        if err != nil {
            return err
        }
        res := make([]ShiftActivity, len(val))
        for i, v := range val {
            res[i] = *(v.(*ShiftActivity))
        }
        m.SetActivities(res)
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
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotes(val)
        return nil
    }
    return res
}
func (m *ShiftItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activities property value. An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
// Parameters:
//  - value : Value to set for the activities property.
func (m *ShiftItem) SetActivities(value []ShiftActivity)() {
    m.activities = value
}
// Sets the displayName property value. The shift label of the shiftItem.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ShiftItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the notes property value. The shift notes for the shiftItem.
// Parameters:
//  - value : Value to set for the notes property.
func (m *ShiftItem) SetNotes(value *string)() {
    m.notes = value
}
