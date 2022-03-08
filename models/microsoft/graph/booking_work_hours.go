package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingWorkHours provides operations to manage the solutionsRoot singleton.
type BookingWorkHours struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The day of the week represented by this instance. Possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
    day *DayOfWeek;
    // A list of start/end times during a day.
    timeSlots []BookingWorkTimeSlotable;
}
// NewBookingWorkHours instantiates a new bookingWorkHours and sets the default values.
func NewBookingWorkHours()(*BookingWorkHours) {
    m := &BookingWorkHours{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBookingWorkHoursFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingWorkHoursFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBookingWorkHours(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingWorkHours) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDay gets the day property value. The day of the week represented by this instance. Possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
func (m *BookingWorkHours) GetDay()(*DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.day
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingWorkHours) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["day"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDay(val.(*DayOfWeek))
        }
        return nil
    }
    res["timeSlots"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingWorkTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkTimeSlotable, len(val))
            for i, v := range val {
                res[i] = v.(BookingWorkTimeSlotable)
            }
            m.SetTimeSlots(res)
        }
        return nil
    }
    return res
}
// GetTimeSlots gets the timeSlots property value. A list of start/end times during a day.
func (m *BookingWorkHours) GetTimeSlots()([]BookingWorkTimeSlotable) {
    if m == nil {
        return nil
    } else {
        return m.timeSlots
    }
}
func (m *BookingWorkHours) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingWorkHours) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDay() != nil {
        cast := (*m.GetDay()).String()
        err := writer.WriteStringValue("day", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingWorkHours) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDay sets the day property value. The day of the week represented by this instance. Possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
func (m *BookingWorkHours) SetDay(value *DayOfWeek)() {
    if m != nil {
        m.day = value
    }
}
// SetTimeSlots sets the timeSlots property value. A list of start/end times during a day.
func (m *BookingWorkHours) SetTimeSlots(value []BookingWorkTimeSlotable)() {
    if m != nil {
        m.timeSlots = value
    }
}
