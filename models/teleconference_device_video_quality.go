package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TeleconferenceDeviceVideoQuality struct {
    TeleconferenceDeviceMediaQuality
}
// TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate composed type wrapper for classes float64, ReferenceNumeric, string
type TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate instantiates a new TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate and sets the default values.
func NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate()(*TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) {
    m := &TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate()
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) GetDouble()(*float64) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) GetString()(*string) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate composed type wrapper for classes float64, ReferenceNumeric, string
type TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate instantiates a new TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate and sets the default values.
func NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate()(*TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) {
    m := &TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate()
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) GetDouble()(*float64) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) GetString()(*string) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate composed type wrapper for classes float64, ReferenceNumeric, string
type TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate instantiates a new TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate and sets the default values.
func NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate()(*TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) {
    m := &TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate()
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) GetDouble()(*float64) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) GetString()(*string) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate composed type wrapper for classes float64, ReferenceNumeric, string
type TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate instantiates a new TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate and sets the default values.
func NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate()(*TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) {
    m := &TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate()
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) GetDouble()(*float64) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) GetString()(*string) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateable interface {
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
type TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateable interface {
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
type TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateable interface {
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
type TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateable interface {
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
// NewTeleconferenceDeviceVideoQuality instantiates a new TeleconferenceDeviceVideoQuality and sets the default values.
func NewTeleconferenceDeviceVideoQuality()(*TeleconferenceDeviceVideoQuality) {
    m := &TeleconferenceDeviceVideoQuality{
        TeleconferenceDeviceMediaQuality: *NewTeleconferenceDeviceMediaQuality(),
    }
    odataTypeValue := "#microsoft.graph.teleconferenceDeviceVideoQuality"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeleconferenceDeviceVideoQualityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeleconferenceDeviceVideoQualityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.teleconferenceDeviceScreenSharingQuality":
                        return NewTeleconferenceDeviceScreenSharingQuality(), nil
                }
            }
        }
    }
    return NewTeleconferenceDeviceVideoQuality(), nil
}
// GetAverageInboundBitRate gets the averageInboundBitRate property value. The average inbound stream video bit rate per second.
// returns a TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateable when successful
func (m *TeleconferenceDeviceVideoQuality) GetAverageInboundBitRate()(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateable) {
    val, err := m.GetBackingStore().Get("averageInboundBitRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateable)
    }
    return nil
}
// GetAverageInboundFrameRate gets the averageInboundFrameRate property value. The average inbound stream video frame rate per second.
// returns a TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateable when successful
func (m *TeleconferenceDeviceVideoQuality) GetAverageInboundFrameRate()(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateable) {
    val, err := m.GetBackingStore().Get("averageInboundFrameRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateable)
    }
    return nil
}
// GetAverageOutboundBitRate gets the averageOutboundBitRate property value. The average outbound stream video bit rate per second.
// returns a TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateable when successful
func (m *TeleconferenceDeviceVideoQuality) GetAverageOutboundBitRate()(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateable) {
    val, err := m.GetBackingStore().Get("averageOutboundBitRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateable)
    }
    return nil
}
// GetAverageOutboundFrameRate gets the averageOutboundFrameRate property value. The average outbound stream video frame rate per second.
// returns a TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateable when successful
func (m *TeleconferenceDeviceVideoQuality) GetAverageOutboundFrameRate()(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateable) {
    val, err := m.GetBackingStore().Get("averageOutboundFrameRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeleconferenceDeviceVideoQuality) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeleconferenceDeviceMediaQuality.GetFieldDeserializers()
    res["averageInboundBitRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundBitRate(val.(*TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRate))
        }
        return nil
    }
    res["averageInboundFrameRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundFrameRate(val.(*TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRate))
        }
        return nil
    }
    res["averageOutboundBitRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundBitRate(val.(*TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRate))
        }
        return nil
    }
    res["averageOutboundFrameRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundFrameRate(val.(*TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRate))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeleconferenceDeviceVideoQuality) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeleconferenceDeviceMediaQuality.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("averageInboundBitRate", m.GetAverageInboundBitRate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("averageInboundFrameRate", m.GetAverageInboundFrameRate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("averageOutboundBitRate", m.GetAverageOutboundBitRate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("averageOutboundFrameRate", m.GetAverageOutboundFrameRate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAverageInboundBitRate sets the averageInboundBitRate property value. The average inbound stream video bit rate per second.
func (m *TeleconferenceDeviceVideoQuality) SetAverageInboundBitRate(value TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateable)() {
    err := m.GetBackingStore().Set("averageInboundBitRate", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageInboundFrameRate sets the averageInboundFrameRate property value. The average inbound stream video frame rate per second.
func (m *TeleconferenceDeviceVideoQuality) SetAverageInboundFrameRate(value TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateable)() {
    err := m.GetBackingStore().Set("averageInboundFrameRate", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageOutboundBitRate sets the averageOutboundBitRate property value. The average outbound stream video bit rate per second.
func (m *TeleconferenceDeviceVideoQuality) SetAverageOutboundBitRate(value TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateable)() {
    err := m.GetBackingStore().Set("averageOutboundBitRate", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageOutboundFrameRate sets the averageOutboundFrameRate property value. The average outbound stream video frame rate per second.
func (m *TeleconferenceDeviceVideoQuality) SetAverageOutboundFrameRate(value TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateable)() {
    err := m.GetBackingStore().Set("averageOutboundFrameRate", value)
    if err != nil {
        panic(err)
    }
}
type TeleconferenceDeviceVideoQualityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeleconferenceDeviceMediaQualityable
    GetAverageInboundBitRate()(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateable)
    GetAverageInboundFrameRate()(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateable)
    GetAverageOutboundBitRate()(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateable)
    GetAverageOutboundFrameRate()(TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateable)
    SetAverageInboundBitRate(value TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundBitRateable)()
    SetAverageInboundFrameRate(value TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageInboundFrameRateable)()
    SetAverageOutboundBitRate(value TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundBitRateable)()
    SetAverageOutboundFrameRate(value TeleconferenceDeviceVideoQuality_TeleconferenceDeviceVideoQuality_averageOutboundFrameRateable)()
}
