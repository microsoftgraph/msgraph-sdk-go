package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ShiftAvailability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the pattern for recurrence
    recurrence *PatternedRecurrence;
    // The time slot(s) preferred by the user.
    timeSlots []TimeRange;
    // Specifies the time zone for the indicated time.
    timeZone *string;
}
// Instantiates a new shiftAvailability and sets the default values.
func NewShiftAvailability()(*ShiftAvailability) {
    m := &ShiftAvailability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ShiftAvailability) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the recurrence property value. Specifies the pattern for recurrence
func (m *ShiftAvailability) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// Gets the timeSlots property value. The time slot(s) preferred by the user.
func (m *ShiftAvailability) GetTimeSlots()([]TimeRange) {
    if m == nil {
        return nil
    } else {
        return m.timeSlots
    }
}
// Gets the timeZone property value. Specifies the time zone for the indicated time.
func (m *ShiftAvailability) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// The deserialization information for the current model
func (m *ShiftAvailability) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        m.SetRecurrence(val.(*PatternedRecurrence))
        return nil
    }
    res["timeSlots"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeRange() })
        if err != nil {
            return err
        }
        res := make([]TimeRange, len(val))
        for i, v := range val {
            res[i] = *(v.(*TimeRange))
        }
        m.SetTimeSlots(res)
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTimeZone(val)
        return nil
    }
    return res
}
func (m *ShiftAvailability) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ShiftAvailability) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeSlots()))
        for i, v := range m.GetTimeSlots() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("timeSlots", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeZone", m.GetTimeZone())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ShiftAvailability) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the recurrence property value. Specifies the pattern for recurrence
// Parameters:
//  - value : Value to set for the recurrence property.
func (m *ShiftAvailability) SetRecurrence(value *PatternedRecurrence)() {
    m.recurrence = value
}
// Sets the timeSlots property value. The time slot(s) preferred by the user.
// Parameters:
//  - value : Value to set for the timeSlots property.
func (m *ShiftAvailability) SetTimeSlots(value []TimeRange)() {
    m.timeSlots = value
}
// Sets the timeZone property value. Specifies the time zone for the indicated time.
// Parameters:
//  - value : Value to set for the timeZone property.
func (m *ShiftAvailability) SetTimeZone(value *string)() {
    m.timeZone = value
}
