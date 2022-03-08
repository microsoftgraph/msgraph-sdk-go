package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecurrenceRange provides operations to manage the drive singleton.
type RecurrenceRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate.
    endDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The number of times to repeat the event. Required and must be positive if type is numbered.
    numberOfOccurrences *int32;
    // Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
    recurrenceTimeZone *string;
    // The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required.
    startDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The recurrence range. The possible values are: endDate, noEnd, numbered. Required.
    type_escaped *RecurrenceRangeType;
}
// NewRecurrenceRange instantiates a new recurrenceRange and sets the default values.
func NewRecurrenceRange()(*RecurrenceRange) {
    m := &RecurrenceRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRecurrenceRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecurrenceRangeFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRecurrenceRange(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecurrenceRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEndDate gets the endDate property value. The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate.
func (m *RecurrenceRange) GetEndDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecurrenceRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["numberOfOccurrences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfOccurrences(val)
        }
        return nil
    }
    res["recurrenceTimeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceTimeZone(val)
        }
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecurrenceRangeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*RecurrenceRangeType))
        }
        return nil
    }
    return res
}
// GetNumberOfOccurrences gets the numberOfOccurrences property value. The number of times to repeat the event. Required and must be positive if type is numbered.
func (m *RecurrenceRange) GetNumberOfOccurrences()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfOccurrences
    }
}
// GetRecurrenceTimeZone gets the recurrenceTimeZone property value. Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
func (m *RecurrenceRange) GetRecurrenceTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceTimeZone
    }
}
// GetStartDate gets the startDate property value. The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required.
func (m *RecurrenceRange) GetStartDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// GetType gets the type property value. The recurrence range. The possible values are: endDate, noEnd, numbered. Required.
func (m *RecurrenceRange) GetType()(*RecurrenceRangeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *RecurrenceRange) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RecurrenceRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteDateOnlyValue("endDate", m.GetEndDate())
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
        err := writer.WriteDateOnlyValue("startDate", m.GetStartDate())
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
func (m *RecurrenceRange) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEndDate sets the endDate property value. The date to stop applying the recurrence pattern. Depending on the recurrence pattern of the event, the last occurrence of the meeting may not be this date. Required if type is endDate.
func (m *RecurrenceRange) SetEndDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.endDate = value
    }
}
// SetNumberOfOccurrences sets the numberOfOccurrences property value. The number of times to repeat the event. Required and must be positive if type is numbered.
func (m *RecurrenceRange) SetNumberOfOccurrences(value *int32)() {
    if m != nil {
        m.numberOfOccurrences = value
    }
}
// SetRecurrenceTimeZone sets the recurrenceTimeZone property value. Time zone for the startDate and endDate properties. Optional. If not specified, the time zone of the event is used.
func (m *RecurrenceRange) SetRecurrenceTimeZone(value *string)() {
    if m != nil {
        m.recurrenceTimeZone = value
    }
}
// SetStartDate sets the startDate property value. The date to start applying the recurrence pattern. The first occurrence of the meeting may be this date or later, depending on the recurrence pattern of the event. Must be the same value as the start property of the recurring event. Required.
func (m *RecurrenceRange) SetStartDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.startDate = value
    }
}
// SetType sets the type property value. The recurrence range. The possible values are: endDate, noEnd, numbered. Required.
func (m *RecurrenceRange) SetType(value *RecurrenceRangeType)() {
    if m != nil {
        m.type_escaped = value
    }
}
