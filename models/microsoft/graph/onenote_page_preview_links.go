package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenotePagePreviewLinks provides operations to call the preview method.
type OnenotePagePreviewLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    previewImageUrl ExternalLinkable;
}
// NewOnenotePagePreviewLinks instantiates a new onenotePagePreviewLinks and sets the default values.
func NewOnenotePagePreviewLinks()(*OnenotePagePreviewLinks) {
    m := &OnenotePagePreviewLinks{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnenotePagePreviewLinksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenotePagePreviewLinksFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnenotePagePreviewLinks(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePagePreviewLinks) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenotePagePreviewLinks) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["previewImageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewImageUrl(val.(ExternalLinkable))
        }
        return nil
    }
    return res
}
// GetPreviewImageUrl gets the previewImageUrl property value. 
func (m *OnenotePagePreviewLinks) GetPreviewImageUrl()(ExternalLinkable) {
    if m == nil {
        return nil
    } else {
        return m.previewImageUrl
    }
}
func (m *OnenotePagePreviewLinks) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePagePreviewLinks) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPreviewImageUrl sets the previewImageUrl property value. 
func (m *OnenotePagePreviewLinks) SetPreviewImageUrl(value ExternalLinkable)() {
    if m != nil {
        m.previewImageUrl = value
    }
}
