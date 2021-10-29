package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RecurrencePattern struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
    dayOfMonth *int32;
    // A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
    daysOfWeek []DayOfWeek;
    // The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
    firstDayOfWeek *DayOfWeek;
    // Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
    index *WeekIndex;
    // The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
    interval *int32;
    // The month in which the event occurs.  This is a number from 1 to 12.
    month *int32;
    // The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required.
    type_escaped *RecurrencePatternType;
}
// Instantiates a new recurrencePattern and sets the default values.
func NewRecurrencePattern()(*RecurrencePattern) {
    m := &RecurrencePattern{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecurrencePattern) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the dayOfMonth property value. The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
func (m *RecurrencePattern) GetDayOfMonth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dayOfMonth
    }
}
// Gets the daysOfWeek property value. A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
func (m *RecurrencePattern) GetDaysOfWeek()([]DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.daysOfWeek
    }
}
// Gets the firstDayOfWeek property value. The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
func (m *RecurrencePattern) GetFirstDayOfWeek()(*DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.firstDayOfWeek
    }
}
// Gets the index property value. Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
func (m *RecurrencePattern) GetIndex()(*WeekIndex) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// Gets the interval property value. The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
func (m *RecurrencePattern) GetInterval()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.interval
    }
}
// Gets the month property value. The month in which the event occurs.  This is a number from 1 to 12.
func (m *RecurrencePattern) GetMonth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.month
    }
}
// Gets the type_escaped property value. The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required.
func (m *RecurrencePattern) GetType_escaped()(*RecurrencePatternType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *RecurrencePattern) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dayOfMonth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDayOfMonth(val)
        return nil
    }
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
    res["firstDayOfWeek"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        cast := val.(DayOfWeek)
        m.SetFirstDayOfWeek(&cast)
        return nil
    }
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeekIndex)
        if err != nil {
            return err
        }
        cast := val.(WeekIndex)
        m.SetIndex(&cast)
        return nil
    }
    res["interval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInterval(val)
        return nil
    }
    res["month"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMonth(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecurrencePatternType)
        if err != nil {
            return err
        }
        cast := val.(RecurrencePatternType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *RecurrencePattern) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RecurrencePattern) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("dayOfMonth", m.GetDayOfMonth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("daysOfWeek", SerializeDayOfWeek(m.GetDaysOfWeek()))
        if err != nil {
            return err
        }
    }
    if m.GetFirstDayOfWeek() != nil {
        cast := m.GetFirstDayOfWeek().String()
        err := writer.WriteStringValue("firstDayOfWeek", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIndex() != nil {
        cast := m.GetIndex().String()
        err := writer.WriteStringValue("index", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("interval", m.GetInterval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("month", m.GetMonth())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
func (m *RecurrencePattern) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the dayOfMonth property value. The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
// Parameters:
//  - value : Value to set for the dayOfMonth property.
func (m *RecurrencePattern) SetDayOfMonth(value *int32)() {
    m.dayOfMonth = value
}
// Sets the daysOfWeek property value. A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
// Parameters:
//  - value : Value to set for the daysOfWeek property.
func (m *RecurrencePattern) SetDaysOfWeek(value []DayOfWeek)() {
    m.daysOfWeek = value
}
// Sets the firstDayOfWeek property value. The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
// Parameters:
//  - value : Value to set for the firstDayOfWeek property.
func (m *RecurrencePattern) SetFirstDayOfWeek(value *DayOfWeek)() {
    m.firstDayOfWeek = value
}
// Sets the index property value. Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
// Parameters:
//  - value : Value to set for the index property.
func (m *RecurrencePattern) SetIndex(value *WeekIndex)() {
    m.index = value
}
// Sets the interval property value. The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
// Parameters:
//  - value : Value to set for the interval property.
func (m *RecurrencePattern) SetInterval(value *int32)() {
    m.interval = value
}
// Sets the month property value. The month in which the event occurs.  This is a number from 1 to 12.
// Parameters:
//  - value : Value to set for the month property.
func (m *RecurrencePattern) SetMonth(value *int32)() {
    m.month = value
}
// Sets the type_escaped property value. The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *RecurrencePattern) SetType_escaped(value *RecurrencePatternType)() {
    m.type_escaped = value
}
