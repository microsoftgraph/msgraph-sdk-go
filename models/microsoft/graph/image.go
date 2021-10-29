package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Image struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Optional. Height of the image, in pixels. Read-only.
    height *int32;
    // Optional. Width of the image, in pixels. Read-only.
    width *int32;
}
// Instantiates a new image and sets the default values.
func NewImage()(*Image) {
    m := &Image{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Image) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the height property value. Optional. Height of the image, in pixels. Read-only.
func (m *Image) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// Gets the width property value. Optional. Width of the image, in pixels. Read-only.
func (m *Image) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
// The deserialization information for the current model
func (m *Image) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetHeight(val)
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWidth(val)
        return nil
    }
    return res
}
func (m *Image) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Image) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the height property value. Optional. Height of the image, in pixels. Read-only.
// Parameters:
//  - value : Value to set for the height property.
func (m *Image) SetHeight(value *int32)() {
    m.height = value
}
// Sets the width property value. Optional. Width of the image, in pixels. Read-only.
// Parameters:
//  - value : Value to set for the width property.
func (m *Image) SetWidth(value *int32)() {
    m.width = value
}
