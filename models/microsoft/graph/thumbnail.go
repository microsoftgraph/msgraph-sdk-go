package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Thumbnail 
type Thumbnail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The content stream for the thumbnail.
    content []byte;
    // The height of the thumbnail, in pixels.
    height *int32;
    // The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
    sourceItemId *string;
    // The URL used to fetch the thumbnail content.
    url *string;
    // The width of the thumbnail, in pixels.
    width *int32;
}
// NewThumbnail instantiates a new thumbnail and sets the default values.
func NewThumbnail()(*Thumbnail) {
    m := &Thumbnail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Thumbnail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContent gets the content property value. The content stream for the thumbnail.
func (m *Thumbnail) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetHeight gets the height property value. The height of the thumbnail, in pixels.
func (m *Thumbnail) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// GetSourceItemId gets the sourceItemId property value. The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
func (m *Thumbnail) GetSourceItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceItemId
    }
}
// GetUrl gets the url property value. The URL used to fetch the thumbnail content.
func (m *Thumbnail) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// GetWidth gets the width property value. The width of the thumbnail, in pixels.
func (m *Thumbnail) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Thumbnail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["sourceItemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceItemId(val)
        }
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
func (m *Thumbnail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Thumbnail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContent sets the content property value. The content stream for the thumbnail.
func (m *Thumbnail) SetContent(value []byte)() {
    m.content = value
}
// SetHeight sets the height property value. The height of the thumbnail, in pixels.
func (m *Thumbnail) SetHeight(value *int32)() {
    m.height = value
}
// SetSourceItemId sets the sourceItemId property value. The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
func (m *Thumbnail) SetSourceItemId(value *string)() {
    m.sourceItemId = value
}
// SetUrl sets the url property value. The URL used to fetch the thumbnail content.
func (m *Thumbnail) SetUrl(value *string)() {
    m.url = value
}
// SetWidth sets the width property value. The width of the thumbnail, in pixels.
func (m *Thumbnail) SetWidth(value *int32)() {
    m.width = value
}
