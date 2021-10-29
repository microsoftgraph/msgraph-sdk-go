package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnenotePagePreviewLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    previewImageUrl *ExternalLink;
}
// Instantiates a new onenotePagePreviewLinks and sets the default values.
func NewOnenotePagePreviewLinks()(*OnenotePagePreviewLinks) {
    m := &OnenotePagePreviewLinks{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePagePreviewLinks) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the previewImageUrl property value. 
func (m *OnenotePagePreviewLinks) GetPreviewImageUrl()(*ExternalLink) {
    if m == nil {
        return nil
    } else {
        return m.previewImageUrl
    }
}
// The deserialization information for the current model
func (m *OnenotePagePreviewLinks) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["previewImageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalLink() })
        if err != nil {
            return err
        }
        m.SetPreviewImageUrl(val.(*ExternalLink))
        return nil
    }
    return res
}
func (m *OnenotePagePreviewLinks) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnenotePagePreviewLinks) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("previewImageUrl", m.GetPreviewImageUrl())
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
func (m *OnenotePagePreviewLinks) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the previewImageUrl property value. 
// Parameters:
//  - value : Value to set for the previewImageUrl property.
func (m *OnenotePagePreviewLinks) SetPreviewImageUrl(value *ExternalLink)() {
    m.previewImageUrl = value
}
