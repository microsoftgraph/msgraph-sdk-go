package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SettingValue 
type SettingValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the setting (as defined by the groupSettingTemplate).
    name *string;
    // Value of the setting.
    value *string;
}
// NewSettingValue instantiates a new settingValue and sets the default values.
func NewSettingValue()(*SettingValue) {
    m := &SettingValue{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSettingValue(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingValue) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SettingValue) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetName gets the name property value. Name of the setting (as defined by the groupSettingTemplate).
func (m *SettingValue) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetValue gets the value property value. Value of the setting.
func (m *SettingValue) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *SettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *SettingValue) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Name of the setting (as defined by the groupSettingTemplate).
func (m *SettingValue) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetValue sets the value property value. Value of the setting.
func (m *SettingValue) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
