package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IntegerRange provides operations to manage the print singleton.
type IntegerRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The inclusive upper bound of the integer range.
    end *int32;
    // The inclusive lower bound of the integer range.
    start *int32;
}
// NewIntegerRange instantiates a new integerRange and sets the default values.
func NewIntegerRange()(*IntegerRange) {
    m := &IntegerRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIntegerRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntegerRangeFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewIntegerRange(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntegerRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnd gets the end property value. The inclusive upper bound of the integer range.
func (m *IntegerRange) GetEnd()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntegerRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val)
        }
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    return res
}
// GetStart gets the start property value. The inclusive lower bound of the integer range.
func (m *IntegerRange) GetStart()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
func (m *IntegerRange) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IntegerRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("start", m.GetStart())
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
func (m *IntegerRange) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnd sets the end property value. The inclusive upper bound of the integer range.
func (m *IntegerRange) SetEnd(value *int32)() {
    if m != nil {
        m.end = value
    }
}
// SetStart sets the start property value. The inclusive lower bound of the integer range.
func (m *IntegerRange) SetStart(value *int32)() {
    if m != nil {
        m.start = value
    }
}
