package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ImageInfo struct {
    addImageQuery *bool;
    additionalData map[string]interface{};
    alternateText *string;
    alternativeText *string;
    iconUrl *string;
}
func NewImageInfo()(*ImageInfo) {
    m := &ImageInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ImageInfo) GetAddImageQuery()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.addImageQuery
    }
}
func (m *ImageInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ImageInfo) GetAlternateText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateText
    }
}
func (m *ImageInfo) GetAlternativeText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternativeText
    }
}
func (m *ImageInfo) GetIconUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.iconUrl
    }
}
func (m *ImageInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addImageQuery"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAddImageQuery(val)
        return nil
    }
    res["alternateText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlternateText(val)
        return nil
    }
    res["alternativeText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlternativeText(val)
        return nil
    }
    res["iconUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIconUrl(val)
        return nil
    }
    return res
}
func (m *ImageInfo) IsNil()(bool) {
    return m == nil
}
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
func (m *ImageInfo) SetAddImageQuery(value *bool)() {
    m.addImageQuery = value
}
func (m *ImageInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ImageInfo) SetAlternateText(value *string)() {
    m.alternateText = value
}
func (m *ImageInfo) SetAlternativeText(value *string)() {
    m.alternativeText = value
}
func (m *ImageInfo) SetIconUrl(value *string)() {
    m.iconUrl = value
}
