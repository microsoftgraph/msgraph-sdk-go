package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PatternedRecurrence struct {
    additionalData map[string]interface{};
    pattern *RecurrencePattern;
    range_escpaped *RecurrenceRange;
}
func NewPatternedRecurrence()(*PatternedRecurrence) {
    m := &PatternedRecurrence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PatternedRecurrence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PatternedRecurrence) GetPattern()(*RecurrencePattern) {
    if m == nil {
        return nil
    } else {
        return m.pattern
    }
}
func (m *PatternedRecurrence) GetRange_escpaped()(*RecurrenceRange) {
    if m == nil {
        return nil
    } else {
        return m.range_escpaped
    }
}
func (m *PatternedRecurrence) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["pattern"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecurrencePattern() })
        if err != nil {
            return err
        }
        m.SetPattern(val.(*RecurrencePattern))
        return nil
    }
    res["range_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecurrenceRange() })
        if err != nil {
            return err
        }
        m.SetRange_escpaped(val.(*RecurrenceRange))
        return nil
    }
    return res
}
func (m *PatternedRecurrence) IsNil()(bool) {
    return m == nil
}
func (m *PatternedRecurrence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("pattern", m.GetPattern())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("range_escpaped", m.GetRange_escpaped())
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
func (m *PatternedRecurrence) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PatternedRecurrence) SetPattern(value *RecurrencePattern)() {
    m.pattern = value
}
func (m *PatternedRecurrence) SetRange_escpaped(value *RecurrenceRange)() {
    m.range_escpaped = value
}
