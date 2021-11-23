package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RgbColor 
type RgbColor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Blue value
    b *int32;
    // Green value
    g *int32;
    // Red value
    r *int32;
}
// NewRgbColor instantiates a new rgbColor and sets the default values.
func NewRgbColor()(*RgbColor) {
    m := &RgbColor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RgbColor) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetB gets the b property value. Blue value
func (m *RgbColor) GetB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.b
    }
}
// GetG gets the g property value. Green value
func (m *RgbColor) GetG()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.g
    }
}
// GetR gets the r property value. Red value
func (m *RgbColor) GetR()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.r
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RgbColor) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["b"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetB(val)
        }
        return nil
    }
    res["g"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetG(val)
        }
        return nil
    }
    res["r"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetR(val)
        }
        return nil
    }
    return res
}
func (m *RgbColor) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RgbColor) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("b", m.GetB())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("g", m.GetG())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("r", m.GetR())
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
func (m *RgbColor) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetB sets the b property value. Blue value
func (m *RgbColor) SetB(value *int32)() {
    m.b = value
}
// SetG sets the g property value. Green value
func (m *RgbColor) SetG(value *int32)() {
    m.g = value
}
// SetR sets the r property value. Red value
func (m *RgbColor) SetR(value *int32)() {
    m.r = value
}
