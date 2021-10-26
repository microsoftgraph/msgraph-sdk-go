package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ContentTypeOrder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Whether this is the default Content Type
    default_escaped *bool;
    // Specifies the position in which the Content Type appears in the selection UI.
    position *int32;
}
// Instantiates a new contentTypeOrder and sets the default values.
func NewContentTypeOrder()(*ContentTypeOrder) {
    m := &ContentTypeOrder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentTypeOrder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the default_escaped property value. Whether this is the default Content Type
func (m *ContentTypeOrder) GetDefault_escaped()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.default_escaped
    }
}
// Gets the position property value. Specifies the position in which the Content Type appears in the selection UI.
func (m *ContentTypeOrder) GetPosition()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
// The deserialization information for the current model
func (m *ContentTypeOrder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["default_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDefault_escaped(val)
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPosition(val)
        return nil
    }
    return res
}
func (m *ContentTypeOrder) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ContentTypeOrder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("default_escaped", m.GetDefault_escaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("position", m.GetPosition())
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
func (m *ContentTypeOrder) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the default_escaped property value. Whether this is the default Content Type
// Parameters:
//  - value : Value to set for the default_escaped property.
func (m *ContentTypeOrder) SetDefault_escaped(value *bool)() {
    m.default_escaped = value
}
// Sets the position property value. Specifies the position in which the Content Type appears in the selection UI.
// Parameters:
//  - value : Value to set for the position property.
func (m *ContentTypeOrder) SetPosition(value *int32)() {
    m.position = value
}
