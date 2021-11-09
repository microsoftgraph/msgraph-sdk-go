package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new workingHours and sets the default values.
func NewWorkingHours()(*WorkingHours) {
    m := &WorkingHours{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkingHours) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the daysOfWeek property value. The days of the week on which the user works.
func (m *WorkingHours) GetDaysOfWeek()([]DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.daysOfWeek
    }
}
// Gets the endTime property value. The time of the day that the user stops working.
func (m *WorkingHours) GetEndTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endTime
    }
}
// Gets the startTime property value. The time of the day that the user starts working.
func (m *WorkingHours) GetStartTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
// Gets the timeZone property value. The time zone to which the working hours apply.
func (m *WorkingHours) GetTimeZone()(*TimeZoneBase) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WorkingHours) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the daysOfWeek property value. The days of the week on which the user works.
// Parameters:
//  - value : Value to set for the daysOfWeek property.
func (m *WorkingHours) SetDaysOfWeek(value []DayOfWeek)() {
    m.daysOfWeek = value
}
// Sets the endTime property value. The time of the day that the user stops working.
// Parameters:
//  - value : Value to set for the endTime property.
func (m *WorkingHours) SetEndTime(value *string)() {
    m.endTime = value
}
// Sets the startTime property value. The time of the day that the user starts working.
// Parameters:
//  - value : Value to set for the startTime property.
func (m *WorkingHours) SetStartTime(value *string)() {
    m.startTime = value
}
// Sets the timeZone property value. The time zone to which the working hours apply.
// Parameters:
//  - value : Value to set for the timeZone property.
func (m *WorkingHours) SetTimeZone(value *TimeZoneBase)() {
    m.timeZone = value
}
