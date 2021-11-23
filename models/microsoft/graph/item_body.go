package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemBody 
type ItemBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The content of the item.
    content *string;
    // The type of the content. Possible values are text and html.
    contentType *BodyType;
}
// NewItemBody instantiates a new itemBody and sets the default values.
func NewItemBody()(*ItemBody) {
    m := &ItemBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContent gets the content property value. The content of the item.
func (m *ItemBody) GetContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetContentType gets the contentType property value. The type of the content. Possible values are text and html.
func (m *ItemBody) GetContentType()(*BodyType) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBodyType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(BodyType)
            m.SetContentType(&cast)
        }
        return nil
    }
    return res
}
func (m *ItemBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ItemBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    if m.GetContentType() != nil {
        cast := m.GetContentType().String()
        err := writer.WriteStringValue("contentType", &cast)
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
func (m *ItemBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContent sets the content property value. The content of the item.
func (m *ItemBody) SetContent(value *string)() {
    m.content = value
}
// SetContentType sets the contentType property value. The type of the content. Possible values are text and html.
func (m *ItemBody) SetContentType(value *BodyType)() {
    m.contentType = value
}
