package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PatternedRecurrence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The frequency of an event. Do not specify for a one-time access review.
    pattern *RecurrencePattern;
    // The duration of an event.
    range_escaped *RecurrenceRange;
}
// Instantiates a new patternedRecurrence and sets the default values.
func NewPatternedRecurrence()(*PatternedRecurrence) {
    m := &PatternedRecurrence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PatternedRecurrence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the pattern property value. The frequency of an event. Do not specify for a one-time access review.
func (m *PatternedRecurrence) GetPattern()(*RecurrencePattern) {
    if m == nil {
        return nil
    } else {
        return m.pattern
    }
}
// Gets the range_escaped property value. The duration of an event.
func (m *PatternedRecurrence) GetRange_escaped()(*RecurrenceRange) {
    if m == nil {
        return nil
    } else {
        return m.range_escaped
    }
}
// The deserialization information for the current model
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
    res["range_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecurrenceRange() })
        if err != nil {
            return err
        }
        m.SetRange_escaped(val.(*RecurrenceRange))
        return nil
    }
    return res
}
func (m *PatternedRecurrence) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PatternedRecurrence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("pattern", m.GetPattern())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("range_escaped", m.GetRange_escaped())
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
func (m *PatternedRecurrence) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the pattern property value. The frequency of an event. Do not specify for a one-time access review.
// Parameters:
//  - value : Value to set for the pattern property.
func (m *PatternedRecurrence) SetPattern(value *RecurrencePattern)() {
    m.pattern = value
}
// Sets the range_escaped property value. The duration of an event.
// Parameters:
//  - value : Value to set for the range_escaped property.
func (m *PatternedRecurrence) SetRange_escaped(value *RecurrenceRange)() {
    m.range_escaped = value
}
