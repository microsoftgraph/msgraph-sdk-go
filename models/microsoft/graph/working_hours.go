package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkingHours 
type WorkingHours struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The days of the week on which the user works.
    daysOfWeek []DayOfWeek;
    // The time of the day that the user stops working.
    endTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // The time of the day that the user starts working.
    startTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // The time zone to which the working hours apply.
    timeZone TimeZoneBaseable;
}
// NewWorkingHours instantiates a new workingHours and sets the default values.
func NewWorkingHours()(*WorkingHours) {
    m := &WorkingHours{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkingHoursFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkingHoursFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkingHours(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkingHours) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDaysOfWeek gets the daysOfWeek property value. The days of the week on which the user works.
func (m *WorkingHours) GetDaysOfWeek()([]DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.daysOfWeek
    }
}
// GetEndTime gets the endTime property value. The time of the day that the user stops working.
func (m *WorkingHours) GetEndTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.endTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkingHours) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["daysOfWeek"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DayOfWeek, len(val))
            for i, v := range val {
                res[i] = *(v.(*DayOfWeek))
            }
            m.SetDaysOfWeek(res)
        }
        return nil
    }
    res["endTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
        }
        return nil
    }
    res["startTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeZoneBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val.(TimeZoneBaseable))
        }
        return nil
    }
    return res
}
// GetStartTime gets the startTime property value. The time of the day that the user starts working.
func (m *WorkingHours) GetStartTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
// GetTimeZone gets the timeZone property value. The time zone to which the working hours apply.
func (m *WorkingHours) GetTimeZone()(TimeZoneBaseable) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// Serialize serializes information the current object
func (m *WorkingHours) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDaysOfWeek() != nil {
        err := writer.WriteCollectionOfStringValues("daysOfWeek", SerializeDayOfWeek(m.GetDaysOfWeek()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("endTime", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("startTime", m.GetStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timeZone", m.GetTimeZone())
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
func (m *WorkingHours) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDaysOfWeek sets the daysOfWeek property value. The days of the week on which the user works.
func (m *WorkingHours) SetDaysOfWeek(value []DayOfWeek)() {
    if m != nil {
        m.daysOfWeek = value
    }
}
// SetEndTime sets the endTime property value. The time of the day that the user stops working.
func (m *WorkingHours) SetEndTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.endTime = value
    }
}
// SetStartTime sets the startTime property value. The time of the day that the user starts working.
func (m *WorkingHours) SetStartTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.startTime = value
    }
}
// SetTimeZone sets the timeZone property value. The time zone to which the working hours apply.
func (m *WorkingHours) SetTimeZone(value TimeZoneBaseable)() {
    if m != nil {
        m.timeZone = value
    }
}
