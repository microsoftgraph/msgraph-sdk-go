package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type HyperlinkOrPictureColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether the display format used for URL columns is an image or a hyperlink.
    isPicture *bool;
}
// Instantiates a new hyperlinkOrPictureColumn and sets the default values.
func NewHyperlinkOrPictureColumn()(*HyperlinkOrPictureColumn) {
    m := &HyperlinkOrPictureColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HyperlinkOrPictureColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isPicture property value. Specifies whether the display format used for URL columns is an image or a hyperlink.
func (m *HyperlinkOrPictureColumn) GetIsPicture()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPicture
    }
}
// The deserialization information for the current model
func (m *HyperlinkOrPictureColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isPicture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPicture(val)
        }
        return nil
    }
    return res
}
func (m *HyperlinkOrPictureColumn) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *HyperlinkOrPictureColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isPicture", m.GetIsPicture())
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
func (m *HyperlinkOrPictureColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isPicture property value. Specifies whether the display format used for URL columns is an image or a hyperlink.
// Parameters:
//  - value : Value to set for the isPicture property.
func (m *HyperlinkOrPictureColumn) SetIsPicture(value *bool)() {
    m.isPicture = value
}
