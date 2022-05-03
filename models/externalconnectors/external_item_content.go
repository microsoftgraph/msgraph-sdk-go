package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalItemContent 
type ExternalItemContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type of content in the value property. Possible values are text and html. These are the content types that the indexer supports, and not the file extension types allowed. Required.
    type_escaped *ExternalItemContentType
    // The content for the externalItem. Required.
    value *string
}
// NewExternalItemContent instantiates a new externalItemContent and sets the default values.
func NewExternalItemContent()(*ExternalItemContent) {
    m := &ExternalItemContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExternalItemContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalItemContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
func (m *ExternalItemContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalItemContentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*ExternalItemContentType))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetType gets the type property value. The type of content in the value property. Possible values are text and html. These are the content types that the indexer supports, and not the file extension types allowed. Required.
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
// Serialize serializes information the current object
func (m *ExternalItemContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetType sets the type property value. The type of content in the value property. Possible values are text and html. These are the content types that the indexer supports, and not the file extension types allowed. Required.
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
