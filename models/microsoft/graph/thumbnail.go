package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Thumbnail struct {
    additionalData map[string]interface{};
    content []byte;
    height *int32;
    sourceItemId *string;
    url *string;
    width *int32;
}
func NewThumbnail()(*Thumbnail) {
    m := &Thumbnail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Thumbnail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Thumbnail) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *Thumbnail) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
func (m *Thumbnail) GetSourceItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceItemId
    }
}
func (m *Thumbnail) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
func (m *Thumbnail) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
func (m *Thumbnail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetHeight(val)
        return nil
    }
    res["sourceItemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceItemId(val)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWidth(val)
        return nil
    }
    return res
}
func (m *Thumbnail) IsNil()(bool) {
    return m == nil
}
func (m *Thumbnail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceItemId", m.GetSourceItemId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("width", m.GetWidth())
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
func (m *Thumbnail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Thumbnail) SetContent(value []byte)() {
    m.content = value
}
func (m *Thumbnail) SetHeight(value *int32)() {
    m.height = value
}
func (m *Thumbnail) SetSourceItemId(value *string)() {
    m.sourceItemId = value
}
func (m *Thumbnail) SetUrl(value *string)() {
    m.url = value
}
func (m *Thumbnail) SetWidth(value *int32)() {
    m.width = value
}
