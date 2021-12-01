package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeSlot 
type TimeSlot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    end *DateTimeTimeZone;
    // 
    start *DateTimeTimeZone;
}
// NewTimeSlot instantiates a new timeSlot and sets the default values.
func NewTimeSlot()(*TimeSlot) {
    m := &TimeSlot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeSlot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnd gets the end property value. 
func (m *TimeSlot) GetEnd()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// GetStart gets the start property value. 
func (m *TimeSlot) GetStart()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeSlot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(*DateTimeTimeZone))
        }
        return nil
    }
    return res
}
func (m *TimeSlot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TimeSlot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("start", m.GetStart())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeSlot) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnd sets the end property value. 
func (m *TimeSlot) SetEnd(value *DateTimeTimeZone)() {
    if m != nil {
        m.end = value
    }
}
// SetStart sets the start property value. 
func (m *TimeSlot) SetStart(value *DateTimeTimeZone)() {
    if m != nil {
        m.start = value
    }
}
