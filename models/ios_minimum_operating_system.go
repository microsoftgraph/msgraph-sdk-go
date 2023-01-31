package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosMinimumOperatingSystem contains properties of the minimum operating system required for an iOS mobile app.
type IosMinimumOperatingSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // 
    odataType *string
    // When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v10_0 *bool
    // When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v11_0 *bool
    // When TRUE, only Version 12.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v12_0 *bool
    // When TRUE, only Version 13.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v13_0 *bool
    // When TRUE, only Version 14.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v14_0 *bool
    // When TRUE, only Version 15.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v15_0 *bool
    // When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v8_0 *bool
    // When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v9_0 *bool
}
// NewIosMinimumOperatingSystem instantiates a new iosMinimumOperatingSystem and sets the default values.
func NewIosMinimumOperatingSystem()(*IosMinimumOperatingSystem) {
    m := &IosMinimumOperatingSystem{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateIosMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosMinimumOperatingSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosMinimumOperatingSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["v10_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV10_0(val)
        }
        return nil
    }
    res["v11_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV11_0(val)
        }
        return nil
    }
    res["v12_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV12_0(val)
        }
        return nil
    }
    res["v13_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV13_0(val)
        }
        return nil
    }
    res["v14_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV14_0(val)
        }
        return nil
    }
    res["v15_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV15_0(val)
        }
        return nil
    }
    res["v8_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV8_0(val)
        }
        return nil
    }
    res["v9_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV9_0(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. 
func (m *IosMinimumOperatingSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetV10_0 gets the v10_0 property value. When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) GetV10_0()(*bool) {
    return m.v10_0
}
// GetV11_0 gets the v11_0 property value. When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) GetV11_0()(*bool) {
    return m.v11_0
}
// GetV12_0 gets the v12_0 property value. When TRUE, only Version 12.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) GetV12_0()(*bool) {
    return m.v12_0
}
// GetV13_0 gets the v13_0 property value. When TRUE, only Version 13.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) GetV13_0()(*bool) {
    return m.v13_0
}
// GetV14_0 gets the v14_0 property value. When TRUE, only Version 14.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) GetV14_0()(*bool) {
    return m.v14_0
}
// GetV15_0 gets the v15_0 property value. When TRUE, only Version 15.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) GetV15_0()(*bool) {
    return m.v15_0
}
// GetV8_0 gets the v8_0 property value. When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) GetV8_0()(*bool) {
    return m.v8_0
}
// GetV9_0 gets the v9_0 property value. When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) GetV9_0()(*bool) {
    return m.v9_0
}
// Serialize serializes information the current object
func (m *IosMinimumOperatingSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_0", m.GetV10_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v11_0", m.GetV11_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v12_0", m.GetV12_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v13_0", m.GetV13_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v14_0", m.GetV14_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v15_0", m.GetV15_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_0", m.GetV8_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v9_0", m.GetV9_0())
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
func (m *IosMinimumOperatingSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. 
func (m *IosMinimumOperatingSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetV10_0 sets the v10_0 property value. When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) SetV10_0(value *bool)() {
    m.v10_0 = value
}
// SetV11_0 sets the v11_0 property value. When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) SetV11_0(value *bool)() {
    m.v11_0 = value
}
// SetV12_0 sets the v12_0 property value. When TRUE, only Version 12.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) SetV12_0(value *bool)() {
    m.v12_0 = value
}
// SetV13_0 sets the v13_0 property value. When TRUE, only Version 13.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) SetV13_0(value *bool)() {
    m.v13_0 = value
}
// SetV14_0 sets the v14_0 property value. When TRUE, only Version 14.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) SetV14_0(value *bool)() {
    m.v14_0 = value
}
// SetV15_0 sets the v15_0 property value. When TRUE, only Version 15.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) SetV15_0(value *bool)() {
    m.v15_0 = value
}
// SetV8_0 sets the v8_0 property value. When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) SetV8_0(value *bool)() {
    m.v8_0 = value
}
// SetV9_0 sets the v9_0 property value. When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *IosMinimumOperatingSystem) SetV9_0(value *bool)() {
    m.v9_0 = value
}
