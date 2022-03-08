package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecurrencePattern provides operations to manage the drive singleton.
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
    // The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required. For more information, see values of type property.
    type_escaped *RecurrencePatternType;
}
// NewRecurrencePattern instantiates a new recurrencePattern and sets the default values.
func NewRecurrencePattern()(*RecurrencePattern) {
    m := &RecurrencePattern{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRecurrencePatternFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecurrencePatternFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRecurrencePattern(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecurrencePattern) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDayOfMonth gets the dayOfMonth property value. The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
func (m *RecurrencePattern) GetDayOfMonth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dayOfMonth
    }
}
// GetDaysOfWeek gets the daysOfWeek property value. A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
func (m *RecurrencePattern) GetDaysOfWeek()([]DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.daysOfWeek
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecurrencePattern) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dayOfMonth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDayOfMonth(val)
        }
        return nil
    }
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
    res["firstDayOfWeek"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstDayOfWeek(val.(*DayOfWeek))
        }
        return nil
    }
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeekIndex)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val.(*WeekIndex))
        }
        return nil
    }
    res["interval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterval(val)
        }
        return nil
    }
    res["month"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonth(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecurrencePatternType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*RecurrencePatternType))
        }
        return nil
    }
    return res
}
// GetFirstDayOfWeek gets the firstDayOfWeek property value. The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
func (m *RecurrencePattern) GetFirstDayOfWeek()(*DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.firstDayOfWeek
    }
}
// GetIndex gets the index property value. Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
func (m *RecurrencePattern) GetIndex()(*WeekIndex) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// GetInterval gets the interval property value. The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
func (m *RecurrencePattern) GetInterval()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.interval
    }
}
// GetMonth gets the month property value. The month in which the event occurs.  This is a number from 1 to 12.
func (m *RecurrencePattern) GetMonth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.month
    }
}
// GetType gets the type property value. The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required. For more information, see values of type property.
func (m *RecurrencePattern) GetType()(*RecurrencePatternType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *RecurrencePattern) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RecurrencePattern) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("dayOfMonth", m.GetDayOfMonth())
        if err != nil {
            return err
        }
    }
    if m.GetDaysOfWeek() != nil {
        err := writer.WriteCollectionOfStringValues("daysOfWeek", SerializeDayOfWeek(m.GetDaysOfWeek()))
        if err != nil {
            return err
        }
    }
    if m.GetFirstDayOfWeek() != nil {
        cast := (*m.GetFirstDayOfWeek()).String()
        err := writer.WriteStringValue("firstDayOfWeek", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIndex() != nil {
        cast := (*m.GetIndex()).String()
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *RecurrencePattern) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDayOfMonth sets the dayOfMonth property value. The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
func (m *RecurrencePattern) SetDayOfMonth(value *int32)() {
    if m != nil {
        m.dayOfMonth = value
    }
}
// SetDaysOfWeek sets the daysOfWeek property value. A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
func (m *RecurrencePattern) SetDaysOfWeek(value []DayOfWeek)() {
    if m != nil {
        m.daysOfWeek = value
    }
}
// SetFirstDayOfWeek sets the firstDayOfWeek property value. The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
func (m *RecurrencePattern) SetFirstDayOfWeek(value *DayOfWeek)() {
    if m != nil {
        m.firstDayOfWeek = value
    }
}
// SetIndex sets the index property value. Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
func (m *RecurrencePattern) SetIndex(value *WeekIndex)() {
    if m != nil {
        m.index = value
    }
}
// SetInterval sets the interval property value. The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
func (m *RecurrencePattern) SetInterval(value *int32)() {
    if m != nil {
        m.interval = value
    }
}
// SetMonth sets the month property value. The month in which the event occurs.  This is a number from 1 to 12.
func (m *RecurrencePattern) SetMonth(value *int32)() {
    if m != nil {
        m.month = value
    }
}
// SetType sets the type property value. The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required. For more information, see values of type property.
func (m *RecurrencePattern) SetType(value *RecurrencePatternType)() {
    if m != nil {
        m.type_escaped = value
    }
}
