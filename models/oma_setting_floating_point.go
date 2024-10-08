package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OmaSettingFloatingPoint oMA Settings Floating Point definition.
type OmaSettingFloatingPoint struct {
    OmaSetting
}
// OmaSettingFloatingPoint_OmaSettingFloatingPoint_value composed type wrapper for classes float32, ReferenceNumeric, string
type OmaSettingFloatingPoint_OmaSettingFloatingPoint_value struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOmaSettingFloatingPoint_OmaSettingFloatingPoint_value instantiates a new OmaSettingFloatingPoint_OmaSettingFloatingPoint_value and sets the default values.
func NewOmaSettingFloatingPoint_OmaSettingFloatingPoint_value()(*OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) {
    m := &OmaSettingFloatingPoint_OmaSettingFloatingPoint_value{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateOmaSettingFloatingPoint_OmaSettingFloatingPoint_valueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOmaSettingFloatingPoint_OmaSettingFloatingPoint_valueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOmaSettingFloatingPoint_OmaSettingFloatingPoint_value()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat32Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetFloat(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetFloat gets the float property value. Composed type representation for type float32
// returns a *float32 when successful
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) GetFloat()(*float32) {
    val, err := m.GetBackingStore().Get("float")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetFloat() != nil {
        err := writer.WriteFloat32Value("", m.GetFloat())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFloat sets the float property value. Composed type representation for type float32
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) SetFloat(value *float32)() {
    err := m.GetBackingStore().Set("float", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *OmaSettingFloatingPoint_OmaSettingFloatingPoint_value) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type OmaSettingFloatingPoint_OmaSettingFloatingPoint_valueable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFloat()(*float32)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFloat(value *float32)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewOmaSettingFloatingPoint instantiates a new OmaSettingFloatingPoint and sets the default values.
func NewOmaSettingFloatingPoint()(*OmaSettingFloatingPoint) {
    m := &OmaSettingFloatingPoint{
        OmaSetting: *NewOmaSetting(),
    }
    odataTypeValue := "#microsoft.graph.omaSettingFloatingPoint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOmaSettingFloatingPointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOmaSettingFloatingPointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingFloatingPoint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OmaSettingFloatingPoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OmaSetting.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOmaSettingFloatingPoint_OmaSettingFloatingPoint_valueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(*OmaSettingFloatingPoint_OmaSettingFloatingPoint_value))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. Value.
// returns a OmaSettingFloatingPoint_OmaSettingFloatingPoint_valueable when successful
func (m *OmaSettingFloatingPoint) GetValue()(OmaSettingFloatingPoint_OmaSettingFloatingPoint_valueable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OmaSettingFloatingPoint_OmaSettingFloatingPoint_valueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OmaSettingFloatingPoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OmaSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. Value.
func (m *OmaSettingFloatingPoint) SetValue(value OmaSettingFloatingPoint_OmaSettingFloatingPoint_valueable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type OmaSettingFloatingPointable interface {
    OmaSettingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()(OmaSettingFloatingPoint_OmaSettingFloatingPoint_valueable)
    SetValue(value OmaSettingFloatingPoint_OmaSettingFloatingPoint_valueable)()
}
