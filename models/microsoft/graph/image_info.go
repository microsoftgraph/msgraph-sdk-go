package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImageInfo 
type ImageInfo struct {
    // Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
    addImageQuery *bool;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Optional; alt-text accessible content for the image
    alternateText *string;
    // 
    alternativeText *string;
    // Optional; URI that points to an icon which represents the application used to generate the activity
    iconUrl *string;
}
// NewImageInfo instantiates a new imageInfo and sets the default values.
func NewImageInfo()(*ImageInfo) {
    m := &ImageInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAddImageQuery gets the addImageQuery property value. Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
func (m *ImageInfo) GetAddImageQuery()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.addImageQuery
    }
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImageInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAlternateText gets the alternateText property value. Optional; alt-text accessible content for the image
func (m *ImageInfo) GetAlternateText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateText
    }
}
// GetAlternativeText gets the alternativeText property value. 
func (m *ImageInfo) GetAlternativeText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternativeText
    }
}
// GetIconUrl gets the iconUrl property value. Optional; URI that points to an icon which represents the application used to generate the activity
func (m *ImageInfo) GetIconUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.iconUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImageInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addImageQuery"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddImageQuery(val)
        }
        return nil
    }
    res["alternateText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateText(val)
        }
        return nil
    }
    res["alternativeText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternativeText(val)
        }
        return nil
    }
    res["iconUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIconUrl(val)
        }
        return nil
    }
    return res
}
func (m *ImageInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ImageInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("addImageQuery", m.GetAddImageQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alternateText", m.GetAlternateText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alternativeText", m.GetAlternativeText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("iconUrl", m.GetIconUrl())
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
// SetAddImageQuery sets the addImageQuery property value. Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
func (m *ImageInfo) SetAddImageQuery(value *bool)() {
    if m != nil {
        m.addImageQuery = value
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImageInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAlternateText sets the alternateText property value. Optional; alt-text accessible content for the image
func (m *ImageInfo) SetAlternateText(value *string)() {
    if m != nil {
        m.alternateText = value
    }
}
// SetAlternativeText sets the alternativeText property value. 
func (m *ImageInfo) SetAlternativeText(value *string)() {
    if m != nil {
        m.alternativeText = value
    }
}
// SetIconUrl sets the iconUrl property value. Optional; URI that points to an icon which represents the application used to generate the activity
func (m *ImageInfo) SetIconUrl(value *string)() {
    if m != nil {
        m.iconUrl = value
    }
}
