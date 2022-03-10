package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SectionLinks provides operations to manage the collection of drive entities.
type SectionLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Opens the section in the OneNote native client if it's installed.
    oneNoteClientUrl ExternalLinkable;
    // Opens the section in OneNote on the web.
    oneNoteWebUrl ExternalLinkable;
}
// NewSectionLinks instantiates a new sectionLinks and sets the default values.
func NewSectionLinks()(*SectionLinks) {
    m := &SectionLinks{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSectionLinksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSectionLinksFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSectionLinks(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SectionLinks) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SectionLinks) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["oneNoteClientUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNoteClientUrl(val.(ExternalLinkable))
        }
        return nil
    }
    res["oneNoteWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNoteWebUrl(val.(ExternalLinkable))
        }
        return nil
    }
    return res
}
// GetOneNoteClientUrl gets the oneNoteClientUrl property value. Opens the section in the OneNote native client if it's installed.
func (m *SectionLinks) GetOneNoteClientUrl()(ExternalLinkable) {
    if m == nil {
        return nil
    } else {
        return m.oneNoteClientUrl
    }
}
// GetOneNoteWebUrl gets the oneNoteWebUrl property value. Opens the section in OneNote on the web.
func (m *SectionLinks) GetOneNoteWebUrl()(ExternalLinkable) {
    if m == nil {
        return nil
    } else {
        return m.oneNoteWebUrl
    }
}
func (m *SectionLinks) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SectionLinks) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("oneNoteClientUrl", m.GetOneNoteClientUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("oneNoteWebUrl", m.GetOneNoteWebUrl())
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
func (m *SectionLinks) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOneNoteClientUrl sets the oneNoteClientUrl property value. Opens the section in the OneNote native client if it's installed.
func (m *SectionLinks) SetOneNoteClientUrl(value ExternalLinkable)() {
    if m != nil {
        m.oneNoteClientUrl = value
    }
}
// SetOneNoteWebUrl sets the oneNoteWebUrl property value. Opens the section in OneNote on the web.
func (m *SectionLinks) SetOneNoteWebUrl(value ExternalLinkable)() {
    if m != nil {
        m.oneNoteWebUrl = value
    }
}
