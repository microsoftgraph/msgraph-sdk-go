package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type NotebookLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Opens the notebook in the OneNote native client if it's installed.
    oneNoteClientUrl *ExternalLink;
    // Opens the notebook in OneNote on the web.
    oneNoteWebUrl *ExternalLink;
}
// Instantiates a new notebookLinks and sets the default values.
func NewNotebookLinks()(*NotebookLinks) {
    m := &NotebookLinks{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NotebookLinks) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the oneNoteClientUrl property value. Opens the notebook in the OneNote native client if it's installed.
func (m *NotebookLinks) GetOneNoteClientUrl()(*ExternalLink) {
    if m == nil {
        return nil
    } else {
        return m.oneNoteClientUrl
    }
}
// Gets the oneNoteWebUrl property value. Opens the notebook in OneNote on the web.
func (m *NotebookLinks) GetOneNoteWebUrl()(*ExternalLink) {
    if m == nil {
        return nil
    } else {
        return m.oneNoteWebUrl
    }
}
// The deserialization information for the current model
func (m *NotebookLinks) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["oneNoteClientUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalLink() })
        if err != nil {
            return err
        }
        m.SetOneNoteClientUrl(val.(*ExternalLink))
        return nil
    }
    res["oneNoteWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExternalLink() })
        if err != nil {
            return err
        }
        m.SetOneNoteWebUrl(val.(*ExternalLink))
        return nil
    }
    return res
}
func (m *NotebookLinks) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *NotebookLinks) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *NotebookLinks) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the oneNoteClientUrl property value. Opens the notebook in the OneNote native client if it's installed.
// Parameters:
//  - value : Value to set for the oneNoteClientUrl property.
func (m *NotebookLinks) SetOneNoteClientUrl(value *ExternalLink)() {
    m.oneNoteClientUrl = value
}
// Sets the oneNoteWebUrl property value. Opens the notebook in OneNote on the web.
// Parameters:
//  - value : Value to set for the oneNoteWebUrl property.
func (m *NotebookLinks) SetOneNoteWebUrl(value *ExternalLink)() {
    m.oneNoteWebUrl = value
}
