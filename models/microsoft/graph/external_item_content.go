package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/externalconnectors"
)

// ExternalItemContent 
type ExternalItemContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of content in the value property. Possible values are text and html. Required.
    type_escaped *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ExternalItemContentType;
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
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExternalItemContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetType gets the type property value. The type of content in the value property. Possible values are text and html. Required.
func (m *ExternalItemContent) GetType()(*i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ExternalItemContentType) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalItemContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ParseExternalItemContentType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ExternalItemContentType)
            m.SetType(&cast)
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
func (m *ExternalItemContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExternalItemContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetType() != nil {
        cast := m.GetType().String()
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
// SetType sets the type property value. The type of content in the value property. Possible values are text and html. Required.
func (m *ExternalItemContent) SetType(value *i611946aca48221be342488e87b2af0987834716d9bc5792c53f59b5e10e9f8f8.ExternalItemContentType)() {
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
