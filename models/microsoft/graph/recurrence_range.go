package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RecurrenceRange struct {
    additionalData map[string]interface{};
    endDate *string;
    numberOfOccurrences *int32;
    recurrenceTimeZone *string;
    startDate *string;
    type_escpaped *RecurrenceRangeType;
}
func NewRecurrenceRange()(*RecurrenceRange) {
    m := &RecurrenceRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RecurrenceRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RecurrenceRange) GetEndDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
func (m *RecurrenceRange) GetNumberOfOccurrences()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfOccurrences
    }
}
func (m *RecurrenceRange) GetRecurrenceTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceTimeZone
    }
}
func (m *RecurrenceRange) GetStartDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
func (m *RecurrenceRange) GetType_escpaped()(*RecurrenceRangeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escpaped
    }
}
func (m *RecurrenceRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEndDate(val)
        return nil
    }
    res["numberOfOccurrences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetNumberOfOccurrences(val)
        return nil
    }
    res["recurrenceTimeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecurrenceTimeZone(val)
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStartDate(val)
        return nil
    }
    res["type_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecurrenceRangeType)
        if err != nil {
            return err
        }
        cast := val.(RecurrenceRangeType)
        m.SetType_escpaped(&cast)
        return nil
    }
    return res
}
func (m *RecurrenceRange) IsNil()(bool) {
    return m == nil
}
func (m *RecurrenceRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numberOfOccurrences", m.GetNumberOfOccurrences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceTimeZone", m.GetRecurrenceTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("startDate", m.GetStartDate())
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
func (m *RecurrenceRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RecurrenceRange) SetEndDate(value *string)() {
    m.endDate = value
}
func (m *RecurrenceRange) SetNumberOfOccurrences(value *int32)() {
    m.numberOfOccurrences = value
}
func (m *RecurrenceRange) SetRecurrenceTimeZone(value *string)() {
    m.recurrenceTimeZone = value
}
func (m *RecurrenceRange) SetStartDate(value *string)() {
    m.startDate = value
}
func (m *RecurrenceRange) SetType_escpaped(value *RecurrenceRangeType)() {
    m.type_escpaped = value
}
