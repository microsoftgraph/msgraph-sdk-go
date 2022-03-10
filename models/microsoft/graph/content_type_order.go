package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContentTypeOrder provides operations to manage the educationRoot singleton.
type ContentTypeOrder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Whether this is the default Content Type
    default_escaped *bool;
    // Specifies the position in which the Content Type appears in the selection UI.
    position *int32;
}
// NewContentTypeOrder instantiates a new contentTypeOrder and sets the default values.
func NewContentTypeOrder()(*ContentTypeOrder) {
    m := &ContentTypeOrder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateContentTypeOrderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentTypeOrderFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewContentTypeOrder(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentTypeOrder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDefault gets the default property value. Whether this is the default Content Type
func (m *ContentTypeOrder) GetDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.default_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentTypeOrder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["default"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefault(val)
        }
        return nil
    }
    res["position"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    return res
}
// GetPosition gets the position property value. Specifies the position in which the Content Type appears in the selection UI.
func (m *ContentTypeOrder) GetPosition()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.position
    }
}
func (m *ContentTypeOrder) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ContentTypeOrder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("default", m.GetDefault())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContentTypeOrder) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDefault sets the default property value. Whether this is the default Content Type
func (m *ContentTypeOrder) SetDefault(value *bool)() {
    if m != nil {
        m.default_escaped = value
    }
}
// SetPosition sets the position property value. Specifies the position in which the Content Type appears in the selection UI.
func (m *ContentTypeOrder) SetPosition(value *int32)() {
    if m != nil {
        m.position = value
    }
}
