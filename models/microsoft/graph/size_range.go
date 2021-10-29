package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SizeRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
    maximumSize *int32;
    // The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
    minimumSize *int32;
}
// Instantiates a new sizeRange and sets the default values.
func NewSizeRange()(*SizeRange) {
    m := &SizeRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SizeRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) GetMaximumSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumSize
    }
}
// Gets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) GetMinimumSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumSize
    }
}
// The deserialization information for the current model
func (m *SizeRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maximumSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaximumSize(val)
        return nil
    }
    res["minimumSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMinimumSize(val)
        return nil
    }
    return res
}
func (m *SizeRange) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SizeRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maximumSize", m.GetMaximumSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minimumSize", m.GetMinimumSize())
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
func (m *SizeRange) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
// Parameters:
//  - value : Value to set for the maximumSize property.
func (m *SizeRange) SetMaximumSize(value *int32)() {
    m.maximumSize = value
}
// Sets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
// Parameters:
//  - value : Value to set for the minimumSize property.
func (m *SizeRange) SetMinimumSize(value *int32)() {
    m.minimumSize = value
}
