package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SizeRange 
type SizeRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
    maximumSize *int32;
    // The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
    minimumSize *int32;
}
// NewSizeRange instantiates a new sizeRange and sets the default values.
func NewSizeRange()(*SizeRange) {
    m := &SizeRange{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SizeRange) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMaximumSize gets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) GetMaximumSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumSize
    }
}
// GetMinimumSize gets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) GetMinimumSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumSize
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SizeRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maximumSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumSize(val)
        }
        return nil
    }
    res["minimumSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSize(val)
        }
        return nil
    }
    return res
}
func (m *SizeRange) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SizeRange) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaximumSize sets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) SetMaximumSize(value *int32)() {
    if m != nil {
        m.maximumSize = value
    }
}
// SetMinimumSize sets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) SetMinimumSize(value *int32)() {
    if m != nil {
        m.minimumSize = value
    }
}
