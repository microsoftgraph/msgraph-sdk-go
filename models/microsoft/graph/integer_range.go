package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IntegerRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The inclusive upper bound of the integer range.
    end *int64;
    // The inclusive lower bound of the integer range.
    start *int64;
}
// Instantiates a new integerRange and sets the default values.
func NewIntegerRange()(*IntegerRange) {
    m := &IntegerRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntegerRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the end property value. The inclusive upper bound of the integer range.
func (m *IntegerRange) GetEnd()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// Gets the start property value. The inclusive lower bound of the integer range.
func (m *IntegerRange) GetStart()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// The deserialization information for the current model
func (m *IntegerRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val)
        }
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
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
func (m *IntegerRange) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IntegerRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("start", m.GetStart())
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
func (m *IntegerRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the end property value. The inclusive upper bound of the integer range.
// Parameters:
//  - value : Value to set for the end property.
func (m *IntegerRange) SetEnd(value *int64)() {
    m.end = value
}
// Sets the start property value. The inclusive lower bound of the integer range.
// Parameters:
//  - value : Value to set for the start property.
func (m *IntegerRange) SetStart(value *int64)() {
    m.start = value
}
