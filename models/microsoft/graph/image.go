package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Image provides operations to manage the drive singleton.
type Image struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Optional. Height of the image, in pixels. Read-only.
    height *int32;
    // Optional. Width of the image, in pixels. Read-only.
    width *int32;
}
// NewImage instantiates a new image and sets the default values.
func NewImage()(*Image) {
    m := &Image{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewImage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Image) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Image) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
// GetHeight gets the height property value. Optional. Height of the image, in pixels. Read-only.
func (m *Image) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// GetWidth gets the width property value. Optional. Width of the image, in pixels. Read-only.
func (m *Image) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
func (m *Image) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Image) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("width", m.GetWidth())
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
func (m *Image) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHeight sets the height property value. Optional. Height of the image, in pixels. Read-only.
func (m *Image) SetHeight(value *int32)() {
    if m != nil {
        m.height = value
    }
}
// SetWidth sets the width property value. Optional. Width of the image, in pixels. Read-only.
func (m *Image) SetWidth(value *int32)() {
    if m != nil {
        m.width = value
    }
}
