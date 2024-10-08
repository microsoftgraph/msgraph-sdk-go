package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TeleconferenceDeviceMediaQuality struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage instantiates a new TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage and sets the default values.
func NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage()(*TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) {
    m := &TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage()
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
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) GetString()(*string) {
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage instantiates a new TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage and sets the default values.
func NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage()(*TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) {
    m := &TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage()
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
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) GetString()(*string) {
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage instantiates a new TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage and sets the default values.
func NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage()(*TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) {
    m := &TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage()
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
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) GetString()(*string) {
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage instantiates a new TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage and sets the default values.
func NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage()(*TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) {
    m := &TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage()
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
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) GetString()(*string) {
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
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
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
type TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewTeleconferenceDeviceMediaQuality instantiates a new TeleconferenceDeviceMediaQuality and sets the default values.
func NewTeleconferenceDeviceMediaQuality()(*TeleconferenceDeviceMediaQuality) {
    m := &TeleconferenceDeviceMediaQuality{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeleconferenceDeviceMediaQualityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceMediaQualityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.teleconferenceDeviceAudioQuality":
                        return NewTeleconferenceDeviceAudioQuality(), nil
                    case "#microsoft.graph.teleconferenceDeviceScreenSharingQuality":
                        return NewTeleconferenceDeviceScreenSharingQuality(), nil
                    case "#microsoft.graph.teleconferenceDeviceVideoQuality":
                        return NewTeleconferenceDeviceVideoQuality(), nil
                }
            }
        }
    }
    return NewTeleconferenceDeviceMediaQuality(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeleconferenceDeviceMediaQuality) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAverageInboundJitter gets the averageInboundJitter property value. The average inbound stream network jitter.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("averageInboundJitter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetAverageInboundPacketLossRateInPercentage gets the averageInboundPacketLossRateInPercentage property value. The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
// returns a TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageable when successful
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundPacketLossRateInPercentage()(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageable) {
    val, err := m.GetBackingStore().Get("averageInboundPacketLossRateInPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageable)
    }
    return nil
}
// GetAverageInboundRoundTripDelay gets the averageInboundRoundTripDelay property value. The average inbound stream network round trip delay.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetAverageInboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("averageInboundRoundTripDelay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetAverageOutboundJitter gets the averageOutboundJitter property value. The average outbound stream network jitter.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("averageOutboundJitter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetAverageOutboundPacketLossRateInPercentage gets the averageOutboundPacketLossRateInPercentage property value. The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
// returns a TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageable when successful
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundPacketLossRateInPercentage()(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageable) {
    val, err := m.GetBackingStore().Get("averageOutboundPacketLossRateInPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageable)
    }
    return nil
}
// GetAverageOutboundRoundTripDelay gets the averageOutboundRoundTripDelay property value. The average outbound stream network round trip delay.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetAverageOutboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("averageOutboundRoundTripDelay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *TeleconferenceDeviceMediaQuality) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetChannelIndex gets the channelIndex property value. The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3.
// returns a *int32 when successful
func (m *TeleconferenceDeviceMediaQuality) GetChannelIndex()(*int32) {
    val, err := m.GetBackingStore().Get("channelIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeleconferenceDeviceMediaQuality) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["averageInboundJitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundJitter(val)
        }
        return nil
    }
    res["averageInboundPacketLossRateInPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundPacketLossRateInPercentage(val.(*TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentage))
        }
        return nil
    }
    res["averageInboundRoundTripDelay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundRoundTripDelay(val)
        }
        return nil
    }
    res["averageOutboundJitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundJitter(val)
        }
        return nil
    }
    res["averageOutboundPacketLossRateInPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundPacketLossRateInPercentage(val.(*TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentage))
        }
        return nil
    }
    res["averageOutboundRoundTripDelay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundRoundTripDelay(val)
        }
        return nil
    }
    res["channelIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelIndex(val)
        }
        return nil
    }
    res["inboundPackets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInboundPackets(val)
        }
        return nil
    }
    res["localIPAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalIPAddress(val)
        }
        return nil
    }
    res["localPort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalPort(val)
        }
        return nil
    }
    res["maximumInboundJitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumInboundJitter(val)
        }
        return nil
    }
    res["maximumInboundPacketLossRateInPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumInboundPacketLossRateInPercentage(val.(*TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentage))
        }
        return nil
    }
    res["maximumInboundRoundTripDelay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumInboundRoundTripDelay(val)
        }
        return nil
    }
    res["maximumOutboundJitter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOutboundJitter(val)
        }
        return nil
    }
    res["maximumOutboundPacketLossRateInPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOutboundPacketLossRateInPercentage(val.(*TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentage))
        }
        return nil
    }
    res["maximumOutboundRoundTripDelay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumOutboundRoundTripDelay(val)
        }
        return nil
    }
    res["mediaDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaDuration(val)
        }
        return nil
    }
    res["networkLinkSpeedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkLinkSpeedInBytes(val)
        }
        return nil
    }
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
    res["outboundPackets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutboundPackets(val)
        }
        return nil
    }
    res["remoteIPAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteIPAddress(val)
        }
        return nil
    }
    res["remotePort"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemotePort(val)
        }
        return nil
    }
    return res
}
// GetInboundPackets gets the inboundPackets property value. The total number of the inbound packets.
// returns a *int64 when successful
func (m *TeleconferenceDeviceMediaQuality) GetInboundPackets()(*int64) {
    val, err := m.GetBackingStore().Get("inboundPackets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetLocalIPAddress gets the localIPAddress property value. the local IP address for the media session.
// returns a *string when successful
func (m *TeleconferenceDeviceMediaQuality) GetLocalIPAddress()(*string) {
    val, err := m.GetBackingStore().Get("localIPAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLocalPort gets the localPort property value. The local media port.
// returns a *int32 when successful
func (m *TeleconferenceDeviceMediaQuality) GetLocalPort()(*int32) {
    val, err := m.GetBackingStore().Get("localPort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMaximumInboundJitter gets the maximumInboundJitter property value. The maximum inbound stream network jitter.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("maximumInboundJitter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMaximumInboundPacketLossRateInPercentage gets the maximumInboundPacketLossRateInPercentage property value. The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
// returns a TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageable when successful
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundPacketLossRateInPercentage()(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageable) {
    val, err := m.GetBackingStore().Get("maximumInboundPacketLossRateInPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageable)
    }
    return nil
}
// GetMaximumInboundRoundTripDelay gets the maximumInboundRoundTripDelay property value. The maximum inbound stream network round trip delay.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetMaximumInboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("maximumInboundRoundTripDelay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMaximumOutboundJitter gets the maximumOutboundJitter property value. The maximum outbound stream network jitter.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("maximumOutboundJitter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMaximumOutboundPacketLossRateInPercentage gets the maximumOutboundPacketLossRateInPercentage property value. The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
// returns a TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageable when successful
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundPacketLossRateInPercentage()(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageable) {
    val, err := m.GetBackingStore().Get("maximumOutboundPacketLossRateInPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageable)
    }
    return nil
}
// GetMaximumOutboundRoundTripDelay gets the maximumOutboundRoundTripDelay property value. The maximum outbound stream network round trip delay.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetMaximumOutboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("maximumOutboundRoundTripDelay")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMediaDuration gets the mediaDuration property value. The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations.
// returns a *ISODuration when successful
func (m *TeleconferenceDeviceMediaQuality) GetMediaDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("mediaDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetNetworkLinkSpeedInBytes gets the networkLinkSpeedInBytes property value. The network link speed in bytes
// returns a *int64 when successful
func (m *TeleconferenceDeviceMediaQuality) GetNetworkLinkSpeedInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("networkLinkSpeedInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TeleconferenceDeviceMediaQuality) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOutboundPackets gets the outboundPackets property value. The total number of the outbound packets.
// returns a *int64 when successful
func (m *TeleconferenceDeviceMediaQuality) GetOutboundPackets()(*int64) {
    val, err := m.GetBackingStore().Get("outboundPackets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRemoteIPAddress gets the remoteIPAddress property value. The remote IP address for the media session.
// returns a *string when successful
func (m *TeleconferenceDeviceMediaQuality) GetRemoteIPAddress()(*string) {
    val, err := m.GetBackingStore().Get("remoteIPAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemotePort gets the remotePort property value. The remote media port.
// returns a *int32 when successful
func (m *TeleconferenceDeviceMediaQuality) GetRemotePort()(*int32) {
    val, err := m.GetBackingStore().Get("remotePort")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeleconferenceDeviceMediaQuality) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("averageInboundJitter", m.GetAverageInboundJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averageInboundPacketLossRateInPercentage", m.GetAverageInboundPacketLossRateInPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("averageInboundRoundTripDelay", m.GetAverageInboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("averageOutboundJitter", m.GetAverageOutboundJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("averageOutboundPacketLossRateInPercentage", m.GetAverageOutboundPacketLossRateInPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("averageOutboundRoundTripDelay", m.GetAverageOutboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("channelIndex", m.GetChannelIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("inboundPackets", m.GetInboundPackets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("localIPAddress", m.GetLocalIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("localPort", m.GetLocalPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maximumInboundJitter", m.GetMaximumInboundJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("maximumInboundPacketLossRateInPercentage", m.GetMaximumInboundPacketLossRateInPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maximumInboundRoundTripDelay", m.GetMaximumInboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maximumOutboundJitter", m.GetMaximumOutboundJitter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("maximumOutboundPacketLossRateInPercentage", m.GetMaximumOutboundPacketLossRateInPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("maximumOutboundRoundTripDelay", m.GetMaximumOutboundRoundTripDelay())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("mediaDuration", m.GetMediaDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("networkLinkSpeedInBytes", m.GetNetworkLinkSpeedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("outboundPackets", m.GetOutboundPackets())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remoteIPAddress", m.GetRemoteIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("remotePort", m.GetRemotePort())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeleconferenceDeviceMediaQuality) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageInboundJitter sets the averageInboundJitter property value. The average inbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("averageInboundJitter", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageInboundPacketLossRateInPercentage sets the averageInboundPacketLossRateInPercentage property value. The average inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundPacketLossRateInPercentage(value TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageable)() {
    err := m.GetBackingStore().Set("averageInboundPacketLossRateInPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageInboundRoundTripDelay sets the averageInboundRoundTripDelay property value. The average inbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) SetAverageInboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("averageInboundRoundTripDelay", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageOutboundJitter sets the averageOutboundJitter property value. The average outbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("averageOutboundJitter", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageOutboundPacketLossRateInPercentage sets the averageOutboundPacketLossRateInPercentage property value. The average outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundPacketLossRateInPercentage(value TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageable)() {
    err := m.GetBackingStore().Set("averageOutboundPacketLossRateInPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageOutboundRoundTripDelay sets the averageOutboundRoundTripDelay property value. The average outbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) SetAverageOutboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("averageOutboundRoundTripDelay", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeleconferenceDeviceMediaQuality) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetChannelIndex sets the channelIndex property value. The channel index of media. Indexing begins with 1.  If a media session contains 3 video modalities, channel indexes will be 1, 2, and 3.
func (m *TeleconferenceDeviceMediaQuality) SetChannelIndex(value *int32)() {
    err := m.GetBackingStore().Set("channelIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetInboundPackets sets the inboundPackets property value. The total number of the inbound packets.
func (m *TeleconferenceDeviceMediaQuality) SetInboundPackets(value *int64)() {
    err := m.GetBackingStore().Set("inboundPackets", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalIPAddress sets the localIPAddress property value. the local IP address for the media session.
func (m *TeleconferenceDeviceMediaQuality) SetLocalIPAddress(value *string)() {
    err := m.GetBackingStore().Set("localIPAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalPort sets the localPort property value. The local media port.
func (m *TeleconferenceDeviceMediaQuality) SetLocalPort(value *int32)() {
    err := m.GetBackingStore().Set("localPort", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumInboundJitter sets the maximumInboundJitter property value. The maximum inbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("maximumInboundJitter", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumInboundPacketLossRateInPercentage sets the maximumInboundPacketLossRateInPercentage property value. The maximum inbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundPacketLossRateInPercentage(value TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageable)() {
    err := m.GetBackingStore().Set("maximumInboundPacketLossRateInPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumInboundRoundTripDelay sets the maximumInboundRoundTripDelay property value. The maximum inbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumInboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("maximumInboundRoundTripDelay", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumOutboundJitter sets the maximumOutboundJitter property value. The maximum outbound stream network jitter.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("maximumOutboundJitter", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumOutboundPacketLossRateInPercentage sets the maximumOutboundPacketLossRateInPercentage property value. The maximum outbound stream packet loss rate in percentage (0-100). For example, 0.01 means 0.01%.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundPacketLossRateInPercentage(value TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageable)() {
    err := m.GetBackingStore().Set("maximumOutboundPacketLossRateInPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetMaximumOutboundRoundTripDelay sets the maximumOutboundRoundTripDelay property value. The maximum outbound stream network round trip delay.
func (m *TeleconferenceDeviceMediaQuality) SetMaximumOutboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("maximumOutboundRoundTripDelay", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaDuration sets the mediaDuration property value. The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of all of the durations.
func (m *TeleconferenceDeviceMediaQuality) SetMediaDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("mediaDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetNetworkLinkSpeedInBytes sets the networkLinkSpeedInBytes property value. The network link speed in bytes
func (m *TeleconferenceDeviceMediaQuality) SetNetworkLinkSpeedInBytes(value *int64)() {
    err := m.GetBackingStore().Set("networkLinkSpeedInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeleconferenceDeviceMediaQuality) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOutboundPackets sets the outboundPackets property value. The total number of the outbound packets.
func (m *TeleconferenceDeviceMediaQuality) SetOutboundPackets(value *int64)() {
    err := m.GetBackingStore().Set("outboundPackets", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteIPAddress sets the remoteIPAddress property value. The remote IP address for the media session.
func (m *TeleconferenceDeviceMediaQuality) SetRemoteIPAddress(value *string)() {
    err := m.GetBackingStore().Set("remoteIPAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetRemotePort sets the remotePort property value. The remote media port.
func (m *TeleconferenceDeviceMediaQuality) SetRemotePort(value *int32)() {
    err := m.GetBackingStore().Set("remotePort", value)
    if err != nil {
        panic(err)
    }
}
type TeleconferenceDeviceMediaQualityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageInboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageInboundPacketLossRateInPercentage()(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageable)
    GetAverageInboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageOutboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageOutboundPacketLossRateInPercentage()(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageable)
    GetAverageOutboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetChannelIndex()(*int32)
    GetInboundPackets()(*int64)
    GetLocalIPAddress()(*string)
    GetLocalPort()(*int32)
    GetMaximumInboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMaximumInboundPacketLossRateInPercentage()(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageable)
    GetMaximumInboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMaximumOutboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMaximumOutboundPacketLossRateInPercentage()(TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageable)
    GetMaximumOutboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMediaDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetNetworkLinkSpeedInBytes()(*int64)
    GetOdataType()(*string)
    GetOutboundPackets()(*int64)
    GetRemoteIPAddress()(*string)
    GetRemotePort()(*int32)
    SetAverageInboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageInboundPacketLossRateInPercentage(value TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageInboundPacketLossRateInPercentageable)()
    SetAverageInboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageOutboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageOutboundPacketLossRateInPercentage(value TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_averageOutboundPacketLossRateInPercentageable)()
    SetAverageOutboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetChannelIndex(value *int32)()
    SetInboundPackets(value *int64)()
    SetLocalIPAddress(value *string)()
    SetLocalPort(value *int32)()
    SetMaximumInboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMaximumInboundPacketLossRateInPercentage(value TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumInboundPacketLossRateInPercentageable)()
    SetMaximumInboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMaximumOutboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMaximumOutboundPacketLossRateInPercentage(value TeleconferenceDeviceMediaQuality_TeleconferenceDeviceMediaQuality_maximumOutboundPacketLossRateInPercentageable)()
    SetMaximumOutboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMediaDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetNetworkLinkSpeedInBytes(value *int64)()
    SetOdataType(value *string)()
    SetOutboundPackets(value *int64)()
    SetRemoteIPAddress(value *string)()
    SetRemotePort(value *int32)()
}
