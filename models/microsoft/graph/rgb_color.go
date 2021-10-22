package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RgbColor struct {
    additionalData map[string]interface{};
    b *int32;
    g *int32;
    r *int32;
}
func NewRgbColor()(*RgbColor) {
    m := &RgbColor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RgbColor) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RgbColor) GetB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.b
    }
}
func (m *RgbColor) GetG()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.g
    }
}
func (m *RgbColor) GetR()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.r
    }
}
func (m *RgbColor) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["b"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetB(val)
        return nil
    }
    res["g"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetG(val)
        return nil
    }
    res["r"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetR(val)
        return nil
    }
    return res
}
func (m *RgbColor) IsNil()(bool) {
    return m == nil
}
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
func (m *RgbColor) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RgbColor) SetB(value *int32)() {
    m.b = value
}
func (m *RgbColor) SetG(value *int32)() {
    m.g = value
}
func (m *RgbColor) SetR(value *int32)() {
    m.r = value
}
