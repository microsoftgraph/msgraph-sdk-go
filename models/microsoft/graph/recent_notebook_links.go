package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecentNotebookLinks provides operations to call the getRecentNotebooks method.
type RecentNotebookLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Opens the notebook in the OneNote native client if it's installed.
    oneNoteClientUrl ExternalLinkable;
    // Opens the notebook in OneNote on the web.
    oneNoteWebUrl ExternalLinkable;
}
// NewRecentNotebookLinks instantiates a new recentNotebookLinks and sets the default values.
func NewRecentNotebookLinks()(*RecentNotebookLinks) {
    m := &RecentNotebookLinks{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRecentNotebookLinksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecentNotebookLinksFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewRecentNotebookLinks(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecentNotebookLinks) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecentNotebookLinks) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
// GetOneNoteClientUrl gets the oneNoteClientUrl property value. Opens the notebook in the OneNote native client if it's installed.
func (m *RecentNotebookLinks) GetOneNoteClientUrl()(ExternalLinkable) {
    if m == nil {
        return nil
    } else {
        return m.oneNoteClientUrl
    }
}
// GetOneNoteWebUrl gets the oneNoteWebUrl property value. Opens the notebook in OneNote on the web.
func (m *RecentNotebookLinks) GetOneNoteWebUrl()(ExternalLinkable) {
    if m == nil {
        return nil
    } else {
        return m.oneNoteWebUrl
    }
}
func (m *RecentNotebookLinks) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RecentNotebookLinks) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *RecentNotebookLinks) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOneNoteClientUrl sets the oneNoteClientUrl property value. Opens the notebook in the OneNote native client if it's installed.
func (m *RecentNotebookLinks) SetOneNoteClientUrl(value ExternalLinkable)() {
    if m != nil {
        m.oneNoteClientUrl = value
    }
}
// SetOneNoteWebUrl sets the oneNoteWebUrl property value. Opens the notebook in OneNote on the web.
func (m *RecentNotebookLinks) SetOneNoteWebUrl(value ExternalLinkable)() {
    if m != nil {
        m.oneNoteWebUrl = value
    }
}
