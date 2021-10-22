package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkingHours struct {
    additionalData map[string]interface{};
    daysOfWeek []DayOfWeek;
    endTime *string;
    startTime *string;
    timeZone *TimeZoneBase;
}
func NewWorkingHours()(*WorkingHours) {
    m := &WorkingHours{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *WorkingHours) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *WorkingHours) GetDaysOfWeek()([]DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.daysOfWeek
    }
}
func (m *WorkingHours) GetEndTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endTime
    }
}
func (m *WorkingHours) GetStartTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startTime
    }
}
func (m *WorkingHours) GetTimeZone()(*TimeZoneBase) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
func (m *WorkingHours) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["daysOfWeek"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDayOfWeek)
        if err != nil {
            return err
        }
        res := make([]DayOfWeek, len(val))
        for i, v := range val {
            res[i] = *(v.(*DayOfWeek))
        }
        m.SetDaysOfWeek(res)
        return nil
    }
    res["endTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEndTime(val)
        return nil
    }
    res["startTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStartTime(val)
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeZoneBase() })
        if err != nil {
            return err
        }
        m.SetTimeZone(val.(*TimeZoneBase))
        return nil
    }
    return res
}
func (m *WorkingHours) IsNil()(bool) {
    return m == nil
}
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
func (m *WorkingHours) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *WorkingHours) SetDaysOfWeek(value []DayOfWeek)() {
    m.daysOfWeek = value
}
func (m *WorkingHours) SetEndTime(value *string)() {
    m.endTime = value
}
func (m *WorkingHours) SetStartTime(value *string)() {
    m.startTime = value
}
func (m *WorkingHours) SetTimeZone(value *TimeZoneBase)() {
    m.timeZone = value
}
