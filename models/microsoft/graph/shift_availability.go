package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ShiftAvailability provides operations to manage the collection of drive entities.
type ShiftAvailability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the pattern for recurrence
    recurrence PatternedRecurrenceable;
    // The time slot(s) preferred by the user.
    timeSlots []TimeRangeable;
    // Specifies the time zone for the indicated time.
    timeZone *string;
}
// NewShiftAvailability instantiates a new shiftAvailability and sets the default values.
func NewShiftAvailability()(*ShiftAvailability) {
    m := &ShiftAvailability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateShiftAvailabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateShiftAvailabilityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewShiftAvailability(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ShiftAvailability) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ShiftAvailability) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["timeSlots"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeRangeable, len(val))
            for i, v := range val {
                res[i] = v.(TimeRangeable)
            }
            m.SetTimeSlots(res)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    return res
}
// GetRecurrence gets the recurrence property value. Specifies the pattern for recurrence
func (m *ShiftAvailability) GetRecurrence()(PatternedRecurrenceable) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetTimeSlots gets the timeSlots property value. The time slot(s) preferred by the user.
func (m *ShiftAvailability) GetTimeSlots()([]TimeRangeable) {
    if m == nil {
        return nil
    } else {
        return m.timeSlots
    }
}
// GetTimeZone gets the timeZone property value. Specifies the time zone for the indicated time.
func (m *ShiftAvailability) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
func (m *ShiftAvailability) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ShiftAvailability) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    if m.GetTimeSlots() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeSlots()))
        for i, v := range m.GetTimeSlots() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ShiftAvailability) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRecurrence sets the recurrence property value. Specifies the pattern for recurrence
func (m *ShiftAvailability) SetRecurrence(value PatternedRecurrenceable)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetTimeSlots sets the timeSlots property value. The time slot(s) preferred by the user.
func (m *ShiftAvailability) SetTimeSlots(value []TimeRangeable)() {
    if m != nil {
        m.timeSlots = value
    }
}
// SetTimeZone sets the timeZone property value. Specifies the time zone for the indicated time.
func (m *ShiftAvailability) SetTimeZone(value *string)() {
    if m != nil {
        m.timeZone = value
    }
}
