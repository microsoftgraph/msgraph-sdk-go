package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HyperlinkOrPictureColumn provides operations to manage the collection of group entities.
type HyperlinkOrPictureColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether the display format used for URL columns is an image or a hyperlink.
    isPicture *bool;
}
// NewHyperlinkOrPictureColumn instantiates a new hyperlinkOrPictureColumn and sets the default values.
func NewHyperlinkOrPictureColumn()(*HyperlinkOrPictureColumn) {
    m := &HyperlinkOrPictureColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateHyperlinkOrPictureColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHyperlinkOrPictureColumnFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewHyperlinkOrPictureColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HyperlinkOrPictureColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetIsPicture gets the isPicture property value. Specifies whether the display format used for URL columns is an image or a hyperlink.
func (m *HyperlinkOrPictureColumn) GetIsPicture()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPicture
    }
}
func (m *HyperlinkOrPictureColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HyperlinkOrPictureColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsPicture sets the isPicture property value. Specifies whether the display format used for URL columns is an image or a hyperlink.
func (m *HyperlinkOrPictureColumn) SetIsPicture(value *bool)() {
    if m != nil {
        m.isPicture = value
    }
}
