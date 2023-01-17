package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidMinimumOperatingSystem contains properties for the minimum operating system required for an Android mobile app.
type AndroidMinimumOperatingSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v10_0 *bool
    // When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v11_0 *bool
    // When TRUE, only Version 4.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_0 *bool
    // When TRUE, only Version 4.0.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_0_3 *bool
    // When TRUE, only Version 4.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_1 *bool
    // When TRUE, only Version 4.2 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_2 *bool
    // When TRUE, only Version 4.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_3 *bool
    // When TRUE, only Version 4.4 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_4 *bool
    // When TRUE, only Version 5.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v5_0 *bool
    // When TRUE, only Version 5.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v5_1 *bool
    // When TRUE, only Version 6.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v6_0 *bool
    // When TRUE, only Version 7.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v7_0 *bool
    // When TRUE, only Version 7.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v7_1 *bool
    // When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v8_0 *bool
    // When TRUE, only Version 8.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v8_1 *bool
    // When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v9_0 *bool
}
// NewAndroidMinimumOperatingSystem instantiates a new androidMinimumOperatingSystem and sets the default values.
func NewAndroidMinimumOperatingSystem()(*AndroidMinimumOperatingSystem) {
    m := &AndroidMinimumOperatingSystem{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateAndroidMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidMinimumOperatingSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidMinimumOperatingSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["v4_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV4_0(val)
        }
        return nil
    }
    res["v4_0_3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV4_0_3(val)
        }
        return nil
    }
    res["v4_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV4_1(val)
        }
        return nil
    }
    res["v4_2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV4_2(val)
        }
        return nil
    }
    res["v4_3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV4_3(val)
        }
        return nil
    }
    res["v4_4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV4_4(val)
        }
        return nil
    }
    res["v5_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV5_0(val)
        }
        return nil
    }
    res["v5_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV5_1(val)
        }
        return nil
    }
    res["v6_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV6_0(val)
        }
        return nil
    }
    res["v7_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV7_0(val)
        }
        return nil
    }
    res["v7_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV7_1(val)
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
    res["v8_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV8_1(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidMinimumOperatingSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetV10_0 gets the v10_0 property value. When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV10_0()(*bool) {
    return m.v10_0
}
// GetV11_0 gets the v11_0 property value. When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV11_0()(*bool) {
    return m.v11_0
}
// GetV4_0 gets the v4_0 property value. When TRUE, only Version 4.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV4_0()(*bool) {
    return m.v4_0
}
// GetV4_0_3 gets the v4_0_3 property value. When TRUE, only Version 4.0.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV4_0_3()(*bool) {
    return m.v4_0_3
}
// GetV4_1 gets the v4_1 property value. When TRUE, only Version 4.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV4_1()(*bool) {
    return m.v4_1
}
// GetV4_2 gets the v4_2 property value. When TRUE, only Version 4.2 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV4_2()(*bool) {
    return m.v4_2
}
// GetV4_3 gets the v4_3 property value. When TRUE, only Version 4.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV4_3()(*bool) {
    return m.v4_3
}
// GetV4_4 gets the v4_4 property value. When TRUE, only Version 4.4 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV4_4()(*bool) {
    return m.v4_4
}
// GetV5_0 gets the v5_0 property value. When TRUE, only Version 5.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV5_0()(*bool) {
    return m.v5_0
}
// GetV5_1 gets the v5_1 property value. When TRUE, only Version 5.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV5_1()(*bool) {
    return m.v5_1
}
// GetV6_0 gets the v6_0 property value. When TRUE, only Version 6.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV6_0()(*bool) {
    return m.v6_0
}
// GetV7_0 gets the v7_0 property value. When TRUE, only Version 7.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV7_0()(*bool) {
    return m.v7_0
}
// GetV7_1 gets the v7_1 property value. When TRUE, only Version 7.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV7_1()(*bool) {
    return m.v7_1
}
// GetV8_0 gets the v8_0 property value. When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV8_0()(*bool) {
    return m.v8_0
}
// GetV8_1 gets the v8_1 property value. When TRUE, only Version 8.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV8_1()(*bool) {
    return m.v8_1
}
// GetV9_0 gets the v9_0 property value. When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV9_0()(*bool) {
    return m.v9_0
}
// Serialize serializes information the current object
func (m *AndroidMinimumOperatingSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteBoolValue("v4_0", m.GetV4_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_0_3", m.GetV4_0_3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_1", m.GetV4_1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_2", m.GetV4_2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_3", m.GetV4_3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_4", m.GetV4_4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v5_0", m.GetV5_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v5_1", m.GetV5_1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v6_0", m.GetV6_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v7_0", m.GetV7_0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v7_1", m.GetV7_1())
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
        err := writer.WriteBoolValue("v8_1", m.GetV8_1())
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
func (m *AndroidMinimumOperatingSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidMinimumOperatingSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetV10_0 sets the v10_0 property value. When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV10_0(value *bool)() {
    m.v10_0 = value
}
// SetV11_0 sets the v11_0 property value. When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV11_0(value *bool)() {
    m.v11_0 = value
}
// SetV4_0 sets the v4_0 property value. When TRUE, only Version 4.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV4_0(value *bool)() {
    m.v4_0 = value
}
// SetV4_0_3 sets the v4_0_3 property value. When TRUE, only Version 4.0.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV4_0_3(value *bool)() {
    m.v4_0_3 = value
}
// SetV4_1 sets the v4_1 property value. When TRUE, only Version 4.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV4_1(value *bool)() {
    m.v4_1 = value
}
// SetV4_2 sets the v4_2 property value. When TRUE, only Version 4.2 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV4_2(value *bool)() {
    m.v4_2 = value
}
// SetV4_3 sets the v4_3 property value. When TRUE, only Version 4.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV4_3(value *bool)() {
    m.v4_3 = value
}
// SetV4_4 sets the v4_4 property value. When TRUE, only Version 4.4 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV4_4(value *bool)() {
    m.v4_4 = value
}
// SetV5_0 sets the v5_0 property value. When TRUE, only Version 5.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV5_0(value *bool)() {
    m.v5_0 = value
}
// SetV5_1 sets the v5_1 property value. When TRUE, only Version 5.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV5_1(value *bool)() {
    m.v5_1 = value
}
// SetV6_0 sets the v6_0 property value. When TRUE, only Version 6.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV6_0(value *bool)() {
    m.v6_0 = value
}
// SetV7_0 sets the v7_0 property value. When TRUE, only Version 7.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV7_0(value *bool)() {
    m.v7_0 = value
}
// SetV7_1 sets the v7_1 property value. When TRUE, only Version 7.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV7_1(value *bool)() {
    m.v7_1 = value
}
// SetV8_0 sets the v8_0 property value. When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV8_0(value *bool)() {
    m.v8_0 = value
}
// SetV8_1 sets the v8_1 property value. When TRUE, only Version 8.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV8_1(value *bool)() {
    m.v8_1 = value
}
// SetV9_0 sets the v9_0 property value. When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV9_0(value *bool)() {
    m.v9_0 = value
}
