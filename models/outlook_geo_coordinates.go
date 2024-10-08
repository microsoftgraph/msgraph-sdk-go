package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type OutlookGeoCoordinates struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy composed type wrapper for classes float64, ReferenceNumeric, string
type OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOutlookGeoCoordinates_OutlookGeoCoordinates_accuracy instantiates a new OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy and sets the default values.
func NewOutlookGeoCoordinates_OutlookGeoCoordinates_accuracy()(*OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) {
    m := &OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateOutlookGeoCoordinates_OutlookGeoCoordinates_accuracyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOutlookGeoCoordinates_OutlookGeoCoordinates_accuracyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOutlookGeoCoordinates_OutlookGeoCoordinates_accuracy()
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) GetDouble()(*float64) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) GetString()(*string) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// OutlookGeoCoordinates_OutlookGeoCoordinates_altitude composed type wrapper for classes float64, ReferenceNumeric, string
type OutlookGeoCoordinates_OutlookGeoCoordinates_altitude struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOutlookGeoCoordinates_OutlookGeoCoordinates_altitude instantiates a new OutlookGeoCoordinates_OutlookGeoCoordinates_altitude and sets the default values.
func NewOutlookGeoCoordinates_OutlookGeoCoordinates_altitude()(*OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) {
    m := &OutlookGeoCoordinates_OutlookGeoCoordinates_altitude{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOutlookGeoCoordinates_OutlookGeoCoordinates_altitude()
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) GetDouble()(*float64) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) GetString()(*string) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitude) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy composed type wrapper for classes float64, ReferenceNumeric, string
type OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy instantiates a new OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy and sets the default values.
func NewOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy()(*OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) {
    m := &OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy()
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) GetDouble()(*float64) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) GetString()(*string) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// OutlookGeoCoordinates_OutlookGeoCoordinates_latitude composed type wrapper for classes float64, ReferenceNumeric, string
type OutlookGeoCoordinates_OutlookGeoCoordinates_latitude struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOutlookGeoCoordinates_OutlookGeoCoordinates_latitude instantiates a new OutlookGeoCoordinates_OutlookGeoCoordinates_latitude and sets the default values.
func NewOutlookGeoCoordinates_OutlookGeoCoordinates_latitude()(*OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) {
    m := &OutlookGeoCoordinates_OutlookGeoCoordinates_latitude{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateOutlookGeoCoordinates_OutlookGeoCoordinates_latitudeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOutlookGeoCoordinates_OutlookGeoCoordinates_latitudeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOutlookGeoCoordinates_OutlookGeoCoordinates_latitude()
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) GetDouble()(*float64) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) GetString()(*string) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_latitude) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// OutlookGeoCoordinates_OutlookGeoCoordinates_longitude composed type wrapper for classes float64, ReferenceNumeric, string
type OutlookGeoCoordinates_OutlookGeoCoordinates_longitude struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOutlookGeoCoordinates_OutlookGeoCoordinates_longitude instantiates a new OutlookGeoCoordinates_OutlookGeoCoordinates_longitude and sets the default values.
func NewOutlookGeoCoordinates_OutlookGeoCoordinates_longitude()(*OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) {
    m := &OutlookGeoCoordinates_OutlookGeoCoordinates_longitude{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateOutlookGeoCoordinates_OutlookGeoCoordinates_longitudeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOutlookGeoCoordinates_OutlookGeoCoordinates_longitudeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewOutlookGeoCoordinates_OutlookGeoCoordinates_longitude()
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) GetDouble()(*float64) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) GetString()(*string) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *OutlookGeoCoordinates_OutlookGeoCoordinates_longitude) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type OutlookGeoCoordinates_OutlookGeoCoordinates_accuracyable interface {
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
type OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeable interface {
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
type OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyable interface {
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
type OutlookGeoCoordinates_OutlookGeoCoordinates_latitudeable interface {
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
type OutlookGeoCoordinates_OutlookGeoCoordinates_longitudeable interface {
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
// NewOutlookGeoCoordinates instantiates a new OutlookGeoCoordinates and sets the default values.
func NewOutlookGeoCoordinates()(*OutlookGeoCoordinates) {
    m := &OutlookGeoCoordinates{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOutlookGeoCoordinatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOutlookGeoCoordinatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookGeoCoordinates(), nil
}
// GetAccuracy gets the accuracy property value. The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters.
// returns a OutlookGeoCoordinates_OutlookGeoCoordinates_accuracyable when successful
func (m *OutlookGeoCoordinates) GetAccuracy()(OutlookGeoCoordinates_OutlookGeoCoordinates_accuracyable) {
    val, err := m.GetBackingStore().Get("accuracy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OutlookGeoCoordinates_OutlookGeoCoordinates_accuracyable)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OutlookGeoCoordinates) GetAdditionalData()(map[string]any) {
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
// GetAltitude gets the altitude property value. The altitude of the location.
// returns a OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeable when successful
func (m *OutlookGeoCoordinates) GetAltitude()(OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeable) {
    val, err := m.GetBackingStore().Get("altitude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeable)
    }
    return nil
}
// GetAltitudeAccuracy gets the altitudeAccuracy property value. The accuracy of the altitude.
// returns a OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyable when successful
func (m *OutlookGeoCoordinates) GetAltitudeAccuracy()(OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyable) {
    val, err := m.GetBackingStore().Get("altitudeAccuracy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *OutlookGeoCoordinates) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OutlookGeoCoordinates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accuracy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookGeoCoordinates_OutlookGeoCoordinates_accuracyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccuracy(val.(*OutlookGeoCoordinates_OutlookGeoCoordinates_accuracy))
        }
        return nil
    }
    res["altitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltitude(val.(*OutlookGeoCoordinates_OutlookGeoCoordinates_altitude))
        }
        return nil
    }
    res["altitudeAccuracy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltitudeAccuracy(val.(*OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracy))
        }
        return nil
    }
    res["latitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookGeoCoordinates_OutlookGeoCoordinates_latitudeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatitude(val.(*OutlookGeoCoordinates_OutlookGeoCoordinates_latitude))
        }
        return nil
    }
    res["longitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookGeoCoordinates_OutlookGeoCoordinates_longitudeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongitude(val.(*OutlookGeoCoordinates_OutlookGeoCoordinates_longitude))
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
    return res
}
// GetLatitude gets the latitude property value. The latitude of the location.
// returns a OutlookGeoCoordinates_OutlookGeoCoordinates_latitudeable when successful
func (m *OutlookGeoCoordinates) GetLatitude()(OutlookGeoCoordinates_OutlookGeoCoordinates_latitudeable) {
    val, err := m.GetBackingStore().Get("latitude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OutlookGeoCoordinates_OutlookGeoCoordinates_latitudeable)
    }
    return nil
}
// GetLongitude gets the longitude property value. The longitude of the location.
// returns a OutlookGeoCoordinates_OutlookGeoCoordinates_longitudeable when successful
func (m *OutlookGeoCoordinates) GetLongitude()(OutlookGeoCoordinates_OutlookGeoCoordinates_longitudeable) {
    val, err := m.GetBackingStore().Get("longitude")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OutlookGeoCoordinates_OutlookGeoCoordinates_longitudeable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *OutlookGeoCoordinates) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OutlookGeoCoordinates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accuracy", m.GetAccuracy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("altitude", m.GetAltitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("altitudeAccuracy", m.GetAltitudeAccuracy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("longitude", m.GetLongitude())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccuracy sets the accuracy property value. The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters.
func (m *OutlookGeoCoordinates) SetAccuracy(value OutlookGeoCoordinates_OutlookGeoCoordinates_accuracyable)() {
    err := m.GetBackingStore().Set("accuracy", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutlookGeoCoordinates) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAltitude sets the altitude property value. The altitude of the location.
func (m *OutlookGeoCoordinates) SetAltitude(value OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeable)() {
    err := m.GetBackingStore().Set("altitude", value)
    if err != nil {
        panic(err)
    }
}
// SetAltitudeAccuracy sets the altitudeAccuracy property value. The accuracy of the altitude.
func (m *OutlookGeoCoordinates) SetAltitudeAccuracy(value OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyable)() {
    err := m.GetBackingStore().Set("altitudeAccuracy", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OutlookGeoCoordinates) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetLatitude sets the latitude property value. The latitude of the location.
func (m *OutlookGeoCoordinates) SetLatitude(value OutlookGeoCoordinates_OutlookGeoCoordinates_latitudeable)() {
    err := m.GetBackingStore().Set("latitude", value)
    if err != nil {
        panic(err)
    }
}
// SetLongitude sets the longitude property value. The longitude of the location.
func (m *OutlookGeoCoordinates) SetLongitude(value OutlookGeoCoordinates_OutlookGeoCoordinates_longitudeable)() {
    err := m.GetBackingStore().Set("longitude", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OutlookGeoCoordinates) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type OutlookGeoCoordinatesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccuracy()(OutlookGeoCoordinates_OutlookGeoCoordinates_accuracyable)
    GetAltitude()(OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeable)
    GetAltitudeAccuracy()(OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetLatitude()(OutlookGeoCoordinates_OutlookGeoCoordinates_latitudeable)
    GetLongitude()(OutlookGeoCoordinates_OutlookGeoCoordinates_longitudeable)
    GetOdataType()(*string)
    SetAccuracy(value OutlookGeoCoordinates_OutlookGeoCoordinates_accuracyable)()
    SetAltitude(value OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeable)()
    SetAltitudeAccuracy(value OutlookGeoCoordinates_OutlookGeoCoordinates_altitudeAccuracyable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetLatitude(value OutlookGeoCoordinates_OutlookGeoCoordinates_latitudeable)()
    SetLongitude(value OutlookGeoCoordinates_OutlookGeoCoordinates_longitudeable)()
    SetOdataType(value *string)()
}
