package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VisualInfo 
type VisualInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Optional. JSON object used to represent an icon which represents the application used to generate the activity
    attribution *ImageInfo;
    // Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
    backgroundColor *string;
    // Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
    content *Json;
    // Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
    description *string;
    // Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
    displayText *string;
}
// NewVisualInfo instantiates a new visualInfo and sets the default values.
func NewVisualInfo()(*VisualInfo) {
    m := &VisualInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VisualInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttribution gets the attribution property value. Optional. JSON object used to represent an icon which represents the application used to generate the activity
func (m *VisualInfo) GetAttribution()(*ImageInfo) {
    if m == nil {
        return nil
    } else {
        return m.attribution
    }
}
// GetBackgroundColor gets the backgroundColor property value. Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
func (m *VisualInfo) GetBackgroundColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.backgroundColor
    }
}
// GetContent gets the content property value. Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
func (m *VisualInfo) GetContent()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetDescription gets the description property value. Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
func (m *VisualInfo) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayText gets the displayText property value. Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
func (m *VisualInfo) GetDisplayText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayText
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VisualInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attribution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImageInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttribution(val.(*ImageInfo))
        }
        return nil
    }
    res["backgroundColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundColor(val)
        }
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(*Json))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayText(val)
        }
        return nil
    }
    return res
}
func (m *VisualInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *VisualInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attribution", m.GetAttribution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("backgroundColor", m.GetBackgroundColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayText", m.GetDisplayText())
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
func (m *VisualInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttribution sets the attribution property value. Optional. JSON object used to represent an icon which represents the application used to generate the activity
func (m *VisualInfo) SetAttribution(value *ImageInfo)() {
    if m != nil {
        m.attribution = value
    }
}
// SetBackgroundColor sets the backgroundColor property value. Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
func (m *VisualInfo) SetBackgroundColor(value *string)() {
    if m != nil {
        m.backgroundColor = value
    }
}
// SetContent sets the content property value. Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
func (m *VisualInfo) SetContent(value *Json)() {
    if m != nil {
        m.content = value
    }
}
// SetDescription sets the description property value. Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
func (m *VisualInfo) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayText sets the displayText property value. Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
func (m *VisualInfo) SetDisplayText(value *string)() {
    if m != nil {
        m.displayText = value
    }
}
