package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnenotePagePreview struct {
    additionalData map[string]interface{};
    links *OnenotePagePreviewLinks;
    previewText *string;
}
func NewOnenotePagePreview()(*OnenotePagePreview) {
    m := &OnenotePagePreview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OnenotePagePreview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OnenotePagePreview) GetLinks()(*OnenotePagePreviewLinks) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
func (m *OnenotePagePreview) GetPreviewText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.previewText
    }
}
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
func (m *OnenotePagePreview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OnenotePagePreview) SetLinks(value *OnenotePagePreviewLinks)() {
    m.links = value
}
func (m *OnenotePagePreview) SetPreviewText(value *string)() {
    m.previewText = value
}
