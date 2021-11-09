package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new rgbColor and sets the default values.
func NewRgbColor()(*RgbColor) {
    m := &RgbColor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RgbColor) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the b property value. Blue value
func (m *RgbColor) GetB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.b
    }
}
// Gets the g property value. Green value
func (m *RgbColor) GetG()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.g
    }
}
// Gets the r property value. Red value
func (m *RgbColor) GetR()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.r
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RgbColor) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the b property value. Blue value
// Parameters:
//  - value : Value to set for the b property.
func (m *RgbColor) SetB(value *int32)() {
    m.b = value
}
// Sets the g property value. Green value
// Parameters:
//  - value : Value to set for the g property.
func (m *RgbColor) SetG(value *int32)() {
    m.g = value
}
// Sets the r property value. Red value
// Parameters:
//  - value : Value to set for the r property.
func (m *RgbColor) SetR(value *int32)() {
    m.r = value
}
