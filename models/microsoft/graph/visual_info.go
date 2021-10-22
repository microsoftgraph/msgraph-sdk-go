package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type VisualInfo struct {
    additionalData map[string]interface{};
    attribution *ImageInfo;
    backgroundColor *string;
    content *Json;
    description *string;
    displayText *string;
}
func NewVisualInfo()(*VisualInfo) {
    m := &VisualInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *VisualInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *VisualInfo) GetAttribution()(*ImageInfo) {
    if m == nil {
        return nil
    } else {
        return m.attribution
    }
}
func (m *VisualInfo) GetBackgroundColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.backgroundColor
    }
}
func (m *VisualInfo) GetContent()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *VisualInfo) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *VisualInfo) GetDisplayText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayText
    }
}
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
func (m *VisualInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *VisualInfo) SetAttribution(value *ImageInfo)() {
    m.attribution = value
}
func (m *VisualInfo) SetBackgroundColor(value *string)() {
    m.backgroundColor = value
}
func (m *VisualInfo) SetContent(value *Json)() {
    m.content = value
}
func (m *VisualInfo) SetDescription(value *string)() {
    m.description = value
}
func (m *VisualInfo) SetDisplayText(value *string)() {
    m.displayText = value
}
