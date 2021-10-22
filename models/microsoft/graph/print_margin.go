package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintMargin struct {
    additionalData map[string]interface{};
    bottom *int32;
    left *int32;
    right *int32;
    top *int32;
}
func NewPrintMargin()(*PrintMargin) {
    m := &PrintMargin{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PrintMargin) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PrintMargin) GetBottom()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bottom
    }
}
func (m *PrintMargin) GetLeft()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.left
    }
}
func (m *PrintMargin) GetRight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.right
    }
}
func (m *PrintMargin) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
func (m *PrintMargin) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bottom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBottom(val)
        return nil
    }
    res["left"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLeft(val)
        return nil
    }
    res["right"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRight(val)
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTop(val)
        return nil
    }
    return res
}
func (m *PrintMargin) IsNil()(bool) {
    return m == nil
}
func (m *PrintMargin) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("bottom", m.GetBottom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("left", m.GetLeft())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("right", m.GetRight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
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
func (m *PrintMargin) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PrintMargin) SetBottom(value *int32)() {
    m.bottom = value
}
func (m *PrintMargin) SetLeft(value *int32)() {
    m.left = value
}
func (m *PrintMargin) SetRight(value *int32)() {
    m.right = value
}
func (m *PrintMargin) SetTop(value *int32)() {
    m.top = value
}
