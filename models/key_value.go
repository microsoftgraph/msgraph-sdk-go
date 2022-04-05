package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyValue 
type KeyValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Key for the key-value pair.
    key *string;
    // Value for the key-value pair.
    value *string;
}
// NewKeyValue instantiates a new keyValue and sets the default values.
func NewKeyValue()(*KeyValue) {
    m := &KeyValue{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateKeyValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyValue(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeyValue) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyValue) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["key"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetKey gets the key property value. Key for the key-value pair.
func (m *KeyValue) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetValue gets the value property value. Value for the key-value pair.
func (m *KeyValue) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *KeyValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
func (m *KeyValue) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKey sets the key property value. Key for the key-value pair.
func (m *KeyValue) SetKey(value *string)() {
    if m != nil {
        m.key = value
    }
}
// SetValue sets the value property value. Value for the key-value pair.
func (m *KeyValue) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
