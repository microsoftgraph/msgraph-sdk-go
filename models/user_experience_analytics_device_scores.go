package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsDeviceScores the user experience analytics device scores entity consolidates the various Endpoint Analytics scores.
type UserExperienceAnalyticsDeviceScores struct {
    Entity
}
// UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore instantiates a new UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore and sets the default values.
func NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore()(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) {
    m := &UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore()
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore instantiates a new UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore and sets the default values.
func NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore()(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) {
    m := &UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore()
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore instantiates a new UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore and sets the default values.
func NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore()(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) {
    m := &UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore()
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore instantiates a new UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore and sets the default values.
func NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore()(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) {
    m := &UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore()
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore instantiates a new UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore and sets the default values.
func NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore()(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) {
    m := &UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore()
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreable interface {
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
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreable interface {
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
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreable interface {
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
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreable interface {
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
type UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreable interface {
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
// NewUserExperienceAnalyticsDeviceScores instantiates a new UserExperienceAnalyticsDeviceScores and sets the default values.
func NewUserExperienceAnalyticsDeviceScores()(*UserExperienceAnalyticsDeviceScores) {
    m := &UserExperienceAnalyticsDeviceScores{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceScoresFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDeviceScoresFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceScores(), nil
}
// GetAppReliabilityScore gets the appReliabilityScore property value. Indicates a score calculated from application health data to indicate when a device is having problems running one or more applications. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreable when successful
func (m *UserExperienceAnalyticsDeviceScores) GetAppReliabilityScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreable) {
    val, err := m.GetBackingStore().Get("appReliabilityScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreable)
    }
    return nil
}
// GetBatteryHealthScore gets the batteryHealthScore property value. Indicates a calulated score indicating the health of the device's battery. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreable when successful
func (m *UserExperienceAnalyticsDeviceScores) GetBatteryHealthScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreable) {
    val, err := m.GetBackingStore().Get("batteryHealthScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreable)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The name of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceScores) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndpointAnalyticsScore gets the endpointAnalyticsScore property value. Indicates a weighted average of the various scores. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreable when successful
func (m *UserExperienceAnalyticsDeviceScores) GetEndpointAnalyticsScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreable) {
    val, err := m.GetBackingStore().Get("endpointAnalyticsScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsDeviceScores) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appReliabilityScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppReliabilityScore(val.(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScore))
        }
        return nil
    }
    res["batteryHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthScore(val.(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScore))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["endpointAnalyticsScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointAnalyticsScore(val.(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScore))
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*UserExperienceAnalyticsHealthState))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val.(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScore))
        }
        return nil
    }
    res["workFromAnywhereScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFromAnywhereScore(val.(*UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScore))
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsDeviceScores) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The manufacturer name of the device. Examples: Microsoft Corporation, HP, Lenovo. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceScores) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. The model name of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsDeviceScores) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. Indicates a weighted average of boot score and logon score used for measuring startup performance. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreable when successful
func (m *UserExperienceAnalyticsDeviceScores) GetStartupPerformanceScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreable) {
    val, err := m.GetBackingStore().Get("startupPerformanceScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreable)
    }
    return nil
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. Indicates a weighted score of the work from anywhere on a device level. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreable when successful
func (m *UserExperienceAnalyticsDeviceScores) GetWorkFromAnywhereScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreable) {
    val, err := m.GetBackingStore().Get("workFromAnywhereScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceScores) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appReliabilityScore", m.GetAppReliabilityScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("batteryHealthScore", m.GetBatteryHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("endpointAnalyticsScore", m.GetEndpointAnalyticsScore())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startupPerformanceScore", m.GetStartupPerformanceScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("workFromAnywhereScore", m.GetWorkFromAnywhereScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppReliabilityScore sets the appReliabilityScore property value. Indicates a score calculated from application health data to indicate when a device is having problems running one or more applications. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetAppReliabilityScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreable)() {
    err := m.GetBackingStore().Set("appReliabilityScore", value)
    if err != nil {
        panic(err)
    }
}
// SetBatteryHealthScore sets the batteryHealthScore property value. Indicates a calulated score indicating the health of the device's battery. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetBatteryHealthScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreable)() {
    err := m.GetBackingStore().Set("batteryHealthScore", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceScores) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetEndpointAnalyticsScore sets the endpointAnalyticsScore property value. Indicates a weighted average of the various scores. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetEndpointAnalyticsScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreable)() {
    err := m.GetBackingStore().Set("endpointAnalyticsScore", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsDeviceScores) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer name of the device. Examples: Microsoft Corporation, HP, Lenovo. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceScores) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The model name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsDeviceScores) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. Indicates a weighted average of boot score and logon score used for measuring startup performance. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetStartupPerformanceScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreable)() {
    err := m.GetBackingStore().Set("startupPerformanceScore", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. Indicates a weighted score of the work from anywhere on a device level. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetWorkFromAnywhereScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreable)() {
    err := m.GetBackingStore().Set("workFromAnywhereScore", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsDeviceScoresable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppReliabilityScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreable)
    GetBatteryHealthScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreable)
    GetDeviceName()(*string)
    GetEndpointAnalyticsScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreable)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetStartupPerformanceScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreable)
    GetWorkFromAnywhereScore()(UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreable)
    SetAppReliabilityScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_appReliabilityScoreable)()
    SetBatteryHealthScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_batteryHealthScoreable)()
    SetDeviceName(value *string)()
    SetEndpointAnalyticsScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_endpointAnalyticsScoreable)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetStartupPerformanceScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_startupPerformanceScoreable)()
    SetWorkFromAnywhereScore(value UserExperienceAnalyticsDeviceScores_UserExperienceAnalyticsDeviceScores_workFromAnywhereScoreable)()
}
