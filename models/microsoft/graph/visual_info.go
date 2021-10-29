package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new visualInfo and sets the default values.
func NewVisualInfo()(*VisualInfo) {
    m := &VisualInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VisualInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attribution property value. Optional. JSON object used to represent an icon which represents the application used to generate the activity
func (m *VisualInfo) GetAttribution()(*ImageInfo) {
    if m == nil {
        return nil
    } else {
        return m.attribution
    }
}
// Gets the backgroundColor property value. Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
func (m *VisualInfo) GetBackgroundColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.backgroundColor
    }
}
// Gets the content property value. Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
func (m *VisualInfo) GetContent()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the description property value. Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
func (m *VisualInfo) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayText property value. Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
func (m *VisualInfo) GetDisplayText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayText
    }
}
// The deserialization information for the current model
func (m *VisualInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attribution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImageInfo() })
        if err != nil {
            return err
        }
        m.SetAttribution(val.(*ImageInfo))
        return nil
    }
    res["backgroundColor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBackgroundColor(val)
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetContent(val.(*Json))
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayText(val)
        return nil
    }
    return res
}
func (m *VisualInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *VisualInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attribution property value. Optional. JSON object used to represent an icon which represents the application used to generate the activity
// Parameters:
//  - value : Value to set for the attribution property.
func (m *VisualInfo) SetAttribution(value *ImageInfo)() {
    m.attribution = value
}
// Sets the backgroundColor property value. Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
// Parameters:
//  - value : Value to set for the backgroundColor property.
func (m *VisualInfo) SetBackgroundColor(value *string)() {
    m.backgroundColor = value
}
// Sets the content property value. Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
// Parameters:
//  - value : Value to set for the content property.
func (m *VisualInfo) SetContent(value *Json)() {
    m.content = value
}
// Sets the description property value. Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
// Parameters:
//  - value : Value to set for the description property.
func (m *VisualInfo) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayText property value. Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
// Parameters:
//  - value : Value to set for the displayText property.
func (m *VisualInfo) SetDisplayText(value *string)() {
    m.displayText = value
}
