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
    endTime *string;
    // The time of the day that the user starts working.
    startTime *string;
    // The time zone to which the working hours apply.
    timeZone *TimeZoneBase;
}
// NewWorkingHours instantiates a new workingHours and sets the default values.
func NewWorkingHours()(*WorkingHours) {
    m := &WorkingHours{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
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
func (m *WorkingHours) GetEndTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endTime
    }
}
// GetStartTime gets the startTime property value. The time of the day that the user starts working.
func (m *WorkingHours) GetStartTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
// GetTimeZone gets the timeZone property value. The time zone to which the working hours apply.
func (m *WorkingHours) GetTimeZone()(*TimeZoneBase) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
        }
        return nil
    }
    res["startTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeZoneBase() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val.(*TimeZoneBase))
        }
        return nil
    }
    return res
}
func (m *WorkingHours) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkingHours) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("daysOfWeek", SerializeDayOfWeek(m.GetDaysOfWeek()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("endTime", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("startTime", m.GetStartTime())
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
func (m *WorkingHours) SetEndTime(value *string)() {
    if m != nil {
        m.endTime = value
    }
}
// SetStartTime sets the startTime property value. The time of the day that the user starts working.
func (m *WorkingHours) SetStartTime(value *string)() {
    if m != nil {
        m.startTime = value
    }
}
// SetTimeZone sets the timeZone property value. The time zone to which the working hours apply.
func (m *WorkingHours) SetTimeZone(value *TimeZoneBase)() {
    if m != nil {
        m.timeZone = value
    }
}
