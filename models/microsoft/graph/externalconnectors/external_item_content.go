package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExternalItemContent provides operations to manage the collection of externalConnection entities.
type ExternalItemContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of content in the value property. Possible values are: text, html, unknownFutureValue.
    type_escaped *ExternalItemContentType;
    // The content for the externalItem. Required.
    value *string;
}
// NewExternalItemContent instantiates a new externalItemContent and sets the default values.
func NewExternalItemContent()(*ExternalItemContent) {
    m := &ExternalItemContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExternalItemContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalItemContentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewExternalItemContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExternalItemContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalItemContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalItemContentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*ExternalItemContentType))
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type of content in the value property. Possible values are: text, html, unknownFutureValue.
func (m *ExternalItemContent) GetType()(*ExternalItemContentType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetValue gets the value property value. The content for the externalItem. Required.
func (m *ExternalItemContent) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *ExternalItemContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExternalItemContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *ExternalItemContent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetType sets the type property value. The type of content in the value property. Possible values are: text, html, unknownFutureValue.
func (m *ExternalItemContent) SetType(value *ExternalItemContentType)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetValue sets the value property value. The content for the externalItem. Required.
func (m *ExternalItemContent) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
