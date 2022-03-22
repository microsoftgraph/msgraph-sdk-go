package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintMargin 
type PrintMargin struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The margin in microns from the bottom edge.
    bottom *int32;
    // The margin in microns from the left edge.
    left *int32;
    // The margin in microns from the right edge.
    right *int32;
    // The margin in microns from the top edge.
    top *int32;
}
// NewPrintMargin instantiates a new printMargin and sets the default values.
func NewPrintMargin()(*PrintMargin) {
    m := &PrintMargin{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrintMarginFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintMarginFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintMargin(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintMargin) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBottom gets the bottom property value. The margin in microns from the bottom edge.
func (m *PrintMargin) GetBottom()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.bottom
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintMargin) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bottom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBottom(val)
        }
        return nil
    }
    res["left"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLeft(val)
        }
        return nil
    }
    res["right"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRight(val)
        }
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    return res
}
// GetLeft gets the left property value. The margin in microns from the left edge.
func (m *PrintMargin) GetLeft()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.left
    }
}
// GetRight gets the right property value. The margin in microns from the right edge.
func (m *PrintMargin) GetRight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.right
    }
}
// GetTop gets the top property value. The margin in microns from the top edge.
func (m *PrintMargin) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintMargin) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBottom sets the bottom property value. The margin in microns from the bottom edge.
func (m *PrintMargin) SetBottom(value *int32)() {
    if m != nil {
        m.bottom = value
    }
}
// SetLeft sets the left property value. The margin in microns from the left edge.
func (m *PrintMargin) SetLeft(value *int32)() {
    if m != nil {
        m.left = value
    }
}
// SetRight sets the right property value. The margin in microns from the right edge.
func (m *PrintMargin) SetRight(value *int32)() {
    if m != nil {
        m.right = value
    }
}
// SetTop sets the top property value. The margin in microns from the top edge.
func (m *PrintMargin) SetTop(value *int32)() {
    if m != nil {
        m.top = value
    }
}
