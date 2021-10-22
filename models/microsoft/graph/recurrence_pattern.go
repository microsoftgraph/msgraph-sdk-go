package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RecurrencePattern struct {
    additionalData map[string]interface{};
    dayOfMonth *int32;
    daysOfWeek []DayOfWeek;
    firstDayOfWeek *DayOfWeek;
    index *WeekIndex;
    interval *int32;
    month *int32;
    type_escpaped *RecurrencePatternType;
}
func NewRecurrencePattern()(*RecurrencePattern) {
    m := &RecurrencePattern{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RecurrencePattern) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RecurrencePattern) GetDayOfMonth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dayOfMonth
    }
}
func (m *RecurrencePattern) GetDaysOfWeek()([]DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.daysOfWeek
    }
}
func (m *RecurrencePattern) GetFirstDayOfWeek()(*DayOfWeek) {
    if m == nil {
        return nil
    } else {
        return m.firstDayOfWeek
    }
}
func (m *RecurrencePattern) GetIndex()(*WeekIndex) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
func (m *RecurrencePattern) GetInterval()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.interval
    }
}
func (m *RecurrencePattern) GetMonth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.month
    }
}
func (m *RecurrencePattern) GetType_escpaped()(*RecurrencePatternType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
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
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecurrencePatternType)
        if err != nil {
            return err
        }
        cast := val.(RecurrencePatternType)
        m.SetType_escpaped(&cast)
        return nil
    }
    return res
}
func (m *RecurrencePattern) IsNil()(bool) {
    return m == nil
}
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
    if m.GetType_escpaped() != nil {
        cast := m.GetType_escpaped().String()
        err := writer.WriteStringValue("type_escpaped", &cast)
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
func (m *RecurrencePattern) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RecurrencePattern) SetDayOfMonth(value *int32)() {
    m.dayOfMonth = value
}
func (m *RecurrencePattern) SetDaysOfWeek(value []DayOfWeek)() {
    m.daysOfWeek = value
}
func (m *RecurrencePattern) SetFirstDayOfWeek(value *DayOfWeek)() {
    m.firstDayOfWeek = value
}
func (m *RecurrencePattern) SetIndex(value *WeekIndex)() {
    m.index = value
}
func (m *RecurrencePattern) SetInterval(value *int32)() {
    m.interval = value
}
func (m *RecurrencePattern) SetMonth(value *int32)() {
    m.month = value
}
func (m *RecurrencePattern) SetType_escpaped(value *RecurrencePatternType)() {
    m.type_escpaped = value
}
