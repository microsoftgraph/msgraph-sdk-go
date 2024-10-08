package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsModelScores the user experience analytics model scores entity consolidates the various Endpoint Analytics scores.
type UserExperienceAnalyticsModelScores struct {
    Entity
}
// UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore instantiates a new UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore and sets the default values.
func NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore()(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) {
    m := &UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore()
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore instantiates a new UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore and sets the default values.
func NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore()(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) {
    m := &UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore()
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore instantiates a new UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore and sets the default values.
func NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore()(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) {
    m := &UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore()
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore instantiates a new UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore and sets the default values.
func NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore()(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) {
    m := &UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore()
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore instantiates a new UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore and sets the default values.
func NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore()(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) {
    m := &UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore()
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreable interface {
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
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreable interface {
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
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreable interface {
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
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreable interface {
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
type UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreable interface {
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
// NewUserExperienceAnalyticsModelScores instantiates a new UserExperienceAnalyticsModelScores and sets the default values.
func NewUserExperienceAnalyticsModelScores()(*UserExperienceAnalyticsModelScores) {
    m := &UserExperienceAnalyticsModelScores{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsModelScores(), nil
}
// GetAppReliabilityScore gets the appReliabilityScore property value. Indicates a score calculated from application health data to indicate when a device is having problems running one or more applications. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreable when successful
func (m *UserExperienceAnalyticsModelScores) GetAppReliabilityScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreable) {
    val, err := m.GetBackingStore().Get("appReliabilityScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreable)
    }
    return nil
}
// GetBatteryHealthScore gets the batteryHealthScore property value. Indicates a calulated score indicating the health of the device's battery. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreable when successful
func (m *UserExperienceAnalyticsModelScores) GetBatteryHealthScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreable) {
    val, err := m.GetBackingStore().Get("batteryHealthScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreable)
    }
    return nil
}
// GetEndpointAnalyticsScore gets the endpointAnalyticsScore property value. Indicates a weighted average of the various scores. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreable when successful
func (m *UserExperienceAnalyticsModelScores) GetEndpointAnalyticsScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreable) {
    val, err := m.GetBackingStore().Get("endpointAnalyticsScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsModelScores) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appReliabilityScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppReliabilityScore(val.(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScore))
        }
        return nil
    }
    res["batteryHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthScore(val.(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScore))
        }
        return nil
    }
    res["endpointAnalyticsScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointAnalyticsScore(val.(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScore))
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
    res["modelDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelDeviceCount(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val.(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScore))
        }
        return nil
    }
    res["workFromAnywhereScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFromAnywhereScore(val.(*UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScore))
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsModelScores) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
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
func (m *UserExperienceAnalyticsModelScores) GetManufacturer()(*string) {
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
func (m *UserExperienceAnalyticsModelScores) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModelDeviceCount gets the modelDeviceCount property value. Indicates unique devices count of given model in a consolidated report. Supports: $select, $OrderBy. Read-only. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
// returns a *int64 when successful
func (m *UserExperienceAnalyticsModelScores) GetModelDeviceCount()(*int64) {
    val, err := m.GetBackingStore().Get("modelDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. Indicates a weighted average of boot score and logon score used for measuring startup performance. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreable when successful
func (m *UserExperienceAnalyticsModelScores) GetStartupPerformanceScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreable) {
    val, err := m.GetBackingStore().Get("startupPerformanceScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreable)
    }
    return nil
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. Indicates a weighted score of the work from anywhere on a device level. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreable when successful
func (m *UserExperienceAnalyticsModelScores) GetWorkFromAnywhereScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreable) {
    val, err := m.GetBackingStore().Get("workFromAnywhereScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsModelScores) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("modelDeviceCount", m.GetModelDeviceCount())
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
func (m *UserExperienceAnalyticsModelScores) SetAppReliabilityScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreable)() {
    err := m.GetBackingStore().Set("appReliabilityScore", value)
    if err != nil {
        panic(err)
    }
}
// SetBatteryHealthScore sets the batteryHealthScore property value. Indicates a calulated score indicating the health of the device's battery. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetBatteryHealthScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreable)() {
    err := m.GetBackingStore().Set("batteryHealthScore", value)
    if err != nil {
        panic(err)
    }
}
// SetEndpointAnalyticsScore sets the endpointAnalyticsScore property value. Indicates a weighted average of the various scores. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetEndpointAnalyticsScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreable)() {
    err := m.GetBackingStore().Set("endpointAnalyticsScore", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsModelScores) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer name of the device. Examples: Microsoft Corporation, HP, Lenovo. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsModelScores) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The model name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsModelScores) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetModelDeviceCount sets the modelDeviceCount property value. Indicates unique devices count of given model in a consolidated report. Supports: $select, $OrderBy. Read-only. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
func (m *UserExperienceAnalyticsModelScores) SetModelDeviceCount(value *int64)() {
    err := m.GetBackingStore().Set("modelDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. Indicates a weighted average of boot score and logon score used for measuring startup performance. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetStartupPerformanceScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreable)() {
    err := m.GetBackingStore().Set("startupPerformanceScore", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. Indicates a weighted score of the work from anywhere on a device level. Valid values range from 0-100. Value -1 means associated score is unavailable. A higher score indicates a healthier device. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetWorkFromAnywhereScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreable)() {
    err := m.GetBackingStore().Set("workFromAnywhereScore", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsModelScoresable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppReliabilityScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreable)
    GetBatteryHealthScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreable)
    GetEndpointAnalyticsScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreable)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelDeviceCount()(*int64)
    GetStartupPerformanceScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreable)
    GetWorkFromAnywhereScore()(UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreable)
    SetAppReliabilityScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_appReliabilityScoreable)()
    SetBatteryHealthScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_batteryHealthScoreable)()
    SetEndpointAnalyticsScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_endpointAnalyticsScoreable)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelDeviceCount(value *int64)()
    SetStartupPerformanceScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_startupPerformanceScoreable)()
    SetWorkFromAnywhereScore(value UserExperienceAnalyticsModelScores_UserExperienceAnalyticsModelScores_workFromAnywhereScoreable)()
}
