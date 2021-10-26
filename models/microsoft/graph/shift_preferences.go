package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ShiftPreferences struct {
    ChangeTrackedEntity
    // Availability of the user to be scheduled for work and its recurrence pattern.
    availability []ShiftAvailability;
}
// Instantiates a new shiftPreferences and sets the default values.
func NewShiftPreferences()(*ShiftPreferences) {
    m := &ShiftPreferences{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    return m
}
// Gets the availability property value. Availability of the user to be scheduled for work and its recurrence pattern.
func (m *ShiftPreferences) GetAvailability()([]ShiftAvailability) {
    if m == nil {
        return nil
    } else {
        return m.availability
    }
}
// The deserialization information for the current model
func (m *ShiftPreferences) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["availability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShiftAvailability() })
        if err != nil {
            return err
        }
        res := make([]ShiftAvailability, len(val))
        for i, v := range val {
            res[i] = *(v.(*ShiftAvailability))
        }
        m.SetAvailability(res)
        return nil
    }
    return res
}
func (m *ShiftPreferences) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ShiftPreferences) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAvailability()))
        for i, v := range m.GetAvailability() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("availability", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the availability property value. Availability of the user to be scheduled for work and its recurrence pattern.
// Parameters:
//  - value : Value to set for the availability property.
func (m *ShiftPreferences) SetAvailability(value []ShiftAvailability)() {
    m.availability = value
}
