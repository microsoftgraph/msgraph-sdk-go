package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new thumbnail and sets the default values.
func NewThumbnail()(*Thumbnail) {
    m := &Thumbnail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Thumbnail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the content property value. The content stream for the thumbnail.
func (m *Thumbnail) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the height property value. The height of the thumbnail, in pixels.
func (m *Thumbnail) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// Gets the sourceItemId property value. The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
func (m *Thumbnail) GetSourceItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceItemId
    }
}
// Gets the url property value. The URL used to fetch the thumbnail content.
func (m *Thumbnail) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
// Gets the width property value. The width of the thumbnail, in pixels.
func (m *Thumbnail) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Thumbnail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the content property value. The content stream for the thumbnail.
// Parameters:
//  - value : Value to set for the content property.
func (m *Thumbnail) SetContent(value []byte)() {
    m.content = value
}
// Sets the height property value. The height of the thumbnail, in pixels.
// Parameters:
//  - value : Value to set for the height property.
func (m *Thumbnail) SetHeight(value *int32)() {
    m.height = value
}
// Sets the sourceItemId property value. The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested.
// Parameters:
//  - value : Value to set for the sourceItemId property.
func (m *Thumbnail) SetSourceItemId(value *string)() {
    m.sourceItemId = value
}
// Sets the url property value. The URL used to fetch the thumbnail content.
// Parameters:
//  - value : Value to set for the url property.
func (m *Thumbnail) SetUrl(value *string)() {
    m.url = value
}
// Sets the width property value. The width of the thumbnail, in pixels.
// Parameters:
//  - value : Value to set for the width property.
func (m *Thumbnail) SetWidth(value *int32)() {
    m.width = value
}
