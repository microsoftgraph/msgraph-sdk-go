package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new imageInfo and sets the default values.
func NewImageInfo()(*ImageInfo) {
    m := &ImageInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the addImageQuery property value. Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
func (m *ImageInfo) GetAddImageQuery()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.addImageQuery
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImageInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the alternateText property value. Optional; alt-text accessible content for the image
func (m *ImageInfo) GetAlternateText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateText
    }
}
// Gets the alternativeText property value. 
func (m *ImageInfo) GetAlternativeText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternativeText
    }
}
// Gets the iconUrl property value. Optional; URI that points to an icon which represents the application used to generate the activity
func (m *ImageInfo) GetIconUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.iconUrl
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the addImageQuery property value. Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
// Parameters:
//  - value : Value to set for the addImageQuery property.
func (m *ImageInfo) SetAddImageQuery(value *bool)() {
    m.addImageQuery = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ImageInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the alternateText property value. Optional; alt-text accessible content for the image
// Parameters:
//  - value : Value to set for the alternateText property.
func (m *ImageInfo) SetAlternateText(value *string)() {
    m.alternateText = value
}
// Sets the alternativeText property value. 
// Parameters:
//  - value : Value to set for the alternativeText property.
func (m *ImageInfo) SetAlternativeText(value *string)() {
    m.alternativeText = value
}
// Sets the iconUrl property value. Optional; URI that points to an icon which represents the application used to generate the activity
// Parameters:
//  - value : Value to set for the iconUrl property.
func (m *ImageInfo) SetIconUrl(value *string)() {
    m.iconUrl = value
}
