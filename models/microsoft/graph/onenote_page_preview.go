package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnenotePagePreview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    links *OnenotePagePreviewLinks;
    // 
    previewText *string;
}
// Instantiates a new onenotePagePreview and sets the default values.
func NewOnenotePagePreview()(*OnenotePagePreview) {
    m := &OnenotePagePreview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePagePreview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the links property value. 
func (m *OnenotePagePreview) GetLinks()(*OnenotePagePreviewLinks) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// Gets the previewText property value. 
func (m *OnenotePagePreview) GetPreviewText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.previewText
    }
}
// The deserialization information for the current model
func (m *OnenotePagePreview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenotePagePreviewLinks() })
        if err != nil {
            return err
        }
        m.SetLinks(val.(*OnenotePagePreviewLinks))
        return nil
    }
    res["previewText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreviewText(val)
        return nil
    }
    return res
}
func (m *OnenotePagePreview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnenotePagePreview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewText", m.GetPreviewText())
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
func (m *OnenotePagePreview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the links property value. 
// Parameters:
//  - value : Value to set for the links property.
func (m *OnenotePagePreview) SetLinks(value *OnenotePagePreviewLinks)() {
    m.links = value
}
// Sets the previewText property value. 
// Parameters:
//  - value : Value to set for the previewText property.
func (m *OnenotePagePreview) SetPreviewText(value *string)() {
    m.previewText = value
}
