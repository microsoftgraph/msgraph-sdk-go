package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsDevicePerformance the user experience analytics device performance entity contains device boot performance details.
type UserExperienceAnalyticsDevicePerformance struct {
    Entity
}
// UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens instantiates a new UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens and sets the default values.
func NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens()(*UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) {
    m := &UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens()
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts instantiates a new UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts and sets the default values.
func NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts()(*UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) {
    m := &UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts()
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore instantiates a new UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore and sets the default values.
func NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore()(*UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) {
    m := &UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore()
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore instantiates a new UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore and sets the default values.
func NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore()(*UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) {
    m := &UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore()
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensable interface {
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
type UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsable interface {
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
type UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreable interface {
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
type UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreable interface {
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
// NewUserExperienceAnalyticsDevicePerformance instantiates a new UserExperienceAnalyticsDevicePerformance and sets the default values.
func NewUserExperienceAnalyticsDevicePerformance()(*UserExperienceAnalyticsDevicePerformance) {
    m := &UserExperienceAnalyticsDevicePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsDevicePerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDevicePerformance(), nil
}
// GetAverageBlueScreens gets the averageBlueScreens property value. Average (mean) number of Blue Screens per device in the last 30 days. Valid values 0 to 9999999
// returns a UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensable when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetAverageBlueScreens()(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensable) {
    val, err := m.GetBackingStore().Get("averageBlueScreens")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensable)
    }
    return nil
}
// GetAverageRestarts gets the averageRestarts property value. Average (mean) number of Restarts per device in the last 30 days. Valid values 0 to 9999999
// returns a UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsable when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetAverageRestarts()(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsable) {
    val, err := m.GetBackingStore().Get("averageRestarts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsable)
    }
    return nil
}
// GetBlueScreenCount gets the blueScreenCount property value. Number of Blue Screens in the last 30 days. Valid values 0 to 9999999
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetBlueScreenCount()(*int32) {
    val, err := m.GetBackingStore().Get("blueScreenCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBootScore gets the bootScore property value. The user experience analytics device boot score.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetBootScore()(*int32) {
    val, err := m.GetBackingStore().Get("bootScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCoreBootTimeInMs gets the coreBootTimeInMs property value. The user experience analytics device core boot time in milliseconds.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetCoreBootTimeInMs()(*int32) {
    val, err := m.GetBackingStore().Get("coreBootTimeInMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetCoreLoginTimeInMs gets the coreLoginTimeInMs property value. The user experience analytics device core login time in milliseconds.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetCoreLoginTimeInMs()(*int32) {
    val, err := m.GetBackingStore().Get("coreLoginTimeInMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceCount gets the deviceCount property value. User experience analytics summarized device count.
// returns a *int64 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetDeviceCount()(*int64) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The user experience analytics device name.
// returns a *string when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiskType gets the diskType property value. The diskType property
// returns a *DiskType when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetDiskType()(*DiskType) {
    val, err := m.GetBackingStore().Get("diskType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DiskType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["averageBlueScreens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageBlueScreens(val.(*UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreens))
        }
        return nil
    }
    res["averageRestarts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageRestarts(val.(*UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestarts))
        }
        return nil
    }
    res["blueScreenCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlueScreenCount(val)
        }
        return nil
    }
    res["bootScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootScore(val)
        }
        return nil
    }
    res["coreBootTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreBootTimeInMs(val)
        }
        return nil
    }
    res["coreLoginTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreLoginTimeInMs(val)
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
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
    res["diskType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiskType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskType(val.(*DiskType))
        }
        return nil
    }
    res["groupPolicyBootTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyBootTimeInMs(val)
        }
        return nil
    }
    res["groupPolicyLoginTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyLoginTimeInMs(val)
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
    res["loginScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginScore(val)
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
    res["modelStartupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelStartupPerformanceScore(val.(*UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScore))
        }
        return nil
    }
    res["operatingSystemVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemVersion(val)
        }
        return nil
    }
    res["responsiveDesktopTimeInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponsiveDesktopTimeInMs(val)
        }
        return nil
    }
    res["restartCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartCount(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val.(*UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScore))
        }
        return nil
    }
    return res
}
// GetGroupPolicyBootTimeInMs gets the groupPolicyBootTimeInMs property value. The user experience analytics device group policy boot time in milliseconds.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetGroupPolicyBootTimeInMs()(*int32) {
    val, err := m.GetBackingStore().Get("groupPolicyBootTimeInMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetGroupPolicyLoginTimeInMs gets the groupPolicyLoginTimeInMs property value. The user experience analytics device group policy login time in milliseconds.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetGroupPolicyLoginTimeInMs()(*int32) {
    val, err := m.GetBackingStore().Get("groupPolicyLoginTimeInMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// GetLoginScore gets the loginScore property value. The user experience analytics device login score.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetLoginScore()(*int32) {
    val, err := m.GetBackingStore().Get("loginScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The user experience analytics device manufacturer.
// returns a *string when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. The user experience analytics device model.
// returns a *string when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModelStartupPerformanceScore gets the modelStartupPerformanceScore property value. The user experience analytics model level startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreable when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetModelStartupPerformanceScore()(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreable) {
    val, err := m.GetBackingStore().Get("modelStartupPerformanceScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreable)
    }
    return nil
}
// GetOperatingSystemVersion gets the operatingSystemVersion property value. The user experience analytics device Operating System version.
// returns a *string when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetOperatingSystemVersion()(*string) {
    val, err := m.GetBackingStore().Get("operatingSystemVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResponsiveDesktopTimeInMs gets the responsiveDesktopTimeInMs property value. The user experience analytics responsive desktop time in milliseconds.
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetResponsiveDesktopTimeInMs()(*int32) {
    val, err := m.GetBackingStore().Get("responsiveDesktopTimeInMs")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRestartCount gets the restartCount property value. Number of Restarts in the last 30 days. Valid values 0 to 9999999
// returns a *int32 when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetRestartCount()(*int32) {
    val, err := m.GetBackingStore().Get("restartCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreable when successful
func (m *UserExperienceAnalyticsDevicePerformance) GetStartupPerformanceScore()(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreable) {
    val, err := m.GetBackingStore().Get("startupPerformanceScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDevicePerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("averageBlueScreens", m.GetAverageBlueScreens())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("averageRestarts", m.GetAverageRestarts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("blueScreenCount", m.GetBlueScreenCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("bootScore", m.GetBootScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("coreBootTimeInMs", m.GetCoreBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("coreLoginTimeInMs", m.GetCoreLoginTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
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
    if m.GetDiskType() != nil {
        cast := (*m.GetDiskType()).String()
        err = writer.WriteStringValue("diskType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupPolicyBootTimeInMs", m.GetGroupPolicyBootTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupPolicyLoginTimeInMs", m.GetGroupPolicyLoginTimeInMs())
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
        err = writer.WriteInt32Value("loginScore", m.GetLoginScore())
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
        err = writer.WriteObjectValue("modelStartupPerformanceScore", m.GetModelStartupPerformanceScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("responsiveDesktopTimeInMs", m.GetResponsiveDesktopTimeInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("restartCount", m.GetRestartCount())
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
    return nil
}
// SetAverageBlueScreens sets the averageBlueScreens property value. Average (mean) number of Blue Screens per device in the last 30 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) SetAverageBlueScreens(value UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensable)() {
    err := m.GetBackingStore().Set("averageBlueScreens", value)
    if err != nil {
        panic(err)
    }
}
// SetAverageRestarts sets the averageRestarts property value. Average (mean) number of Restarts per device in the last 30 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) SetAverageRestarts(value UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsable)() {
    err := m.GetBackingStore().Set("averageRestarts", value)
    if err != nil {
        panic(err)
    }
}
// SetBlueScreenCount sets the blueScreenCount property value. Number of Blue Screens in the last 30 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) SetBlueScreenCount(value *int32)() {
    err := m.GetBackingStore().Set("blueScreenCount", value)
    if err != nil {
        panic(err)
    }
}
// SetBootScore sets the bootScore property value. The user experience analytics device boot score.
func (m *UserExperienceAnalyticsDevicePerformance) SetBootScore(value *int32)() {
    err := m.GetBackingStore().Set("bootScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCoreBootTimeInMs sets the coreBootTimeInMs property value. The user experience analytics device core boot time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetCoreBootTimeInMs(value *int32)() {
    err := m.GetBackingStore().Set("coreBootTimeInMs", value)
    if err != nil {
        panic(err)
    }
}
// SetCoreLoginTimeInMs sets the coreLoginTimeInMs property value. The user experience analytics device core login time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetCoreLoginTimeInMs(value *int32)() {
    err := m.GetBackingStore().Set("coreLoginTimeInMs", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCount sets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsDevicePerformance) SetDeviceCount(value *int64)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The user experience analytics device name.
func (m *UserExperienceAnalyticsDevicePerformance) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDiskType sets the diskType property value. The diskType property
func (m *UserExperienceAnalyticsDevicePerformance) SetDiskType(value *DiskType)() {
    err := m.GetBackingStore().Set("diskType", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyBootTimeInMs sets the groupPolicyBootTimeInMs property value. The user experience analytics device group policy boot time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetGroupPolicyBootTimeInMs(value *int32)() {
    err := m.GetBackingStore().Set("groupPolicyBootTimeInMs", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyLoginTimeInMs sets the groupPolicyLoginTimeInMs property value. The user experience analytics device group policy login time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetGroupPolicyLoginTimeInMs(value *int32)() {
    err := m.GetBackingStore().Set("groupPolicyLoginTimeInMs", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsDevicePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLoginScore sets the loginScore property value. The user experience analytics device login score.
func (m *UserExperienceAnalyticsDevicePerformance) SetLoginScore(value *int32)() {
    err := m.GetBackingStore().Set("loginScore", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsDevicePerformance) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsDevicePerformance) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetModelStartupPerformanceScore sets the modelStartupPerformanceScore property value. The user experience analytics model level startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDevicePerformance) SetModelStartupPerformanceScore(value UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreable)() {
    err := m.GetBackingStore().Set("modelStartupPerformanceScore", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystemVersion sets the operatingSystemVersion property value. The user experience analytics device Operating System version.
func (m *UserExperienceAnalyticsDevicePerformance) SetOperatingSystemVersion(value *string)() {
    err := m.GetBackingStore().Set("operatingSystemVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetResponsiveDesktopTimeInMs sets the responsiveDesktopTimeInMs property value. The user experience analytics responsive desktop time in milliseconds.
func (m *UserExperienceAnalyticsDevicePerformance) SetResponsiveDesktopTimeInMs(value *int32)() {
    err := m.GetBackingStore().Set("responsiveDesktopTimeInMs", value)
    if err != nil {
        panic(err)
    }
}
// SetRestartCount sets the restartCount property value. Number of Restarts in the last 30 days. Valid values 0 to 9999999
func (m *UserExperienceAnalyticsDevicePerformance) SetRestartCount(value *int32)() {
    err := m.GetBackingStore().Set("restartCount", value)
    if err != nil {
        panic(err)
    }
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDevicePerformance) SetStartupPerformanceScore(value UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreable)() {
    err := m.GetBackingStore().Set("startupPerformanceScore", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsDevicePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageBlueScreens()(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensable)
    GetAverageRestarts()(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsable)
    GetBlueScreenCount()(*int32)
    GetBootScore()(*int32)
    GetCoreBootTimeInMs()(*int32)
    GetCoreLoginTimeInMs()(*int32)
    GetDeviceCount()(*int64)
    GetDeviceName()(*string)
    GetDiskType()(*DiskType)
    GetGroupPolicyBootTimeInMs()(*int32)
    GetGroupPolicyLoginTimeInMs()(*int32)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetLoginScore()(*int32)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelStartupPerformanceScore()(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreable)
    GetOperatingSystemVersion()(*string)
    GetResponsiveDesktopTimeInMs()(*int32)
    GetRestartCount()(*int32)
    GetStartupPerformanceScore()(UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreable)
    SetAverageBlueScreens(value UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageBlueScreensable)()
    SetAverageRestarts(value UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_averageRestartsable)()
    SetBlueScreenCount(value *int32)()
    SetBootScore(value *int32)()
    SetCoreBootTimeInMs(value *int32)()
    SetCoreLoginTimeInMs(value *int32)()
    SetDeviceCount(value *int64)()
    SetDeviceName(value *string)()
    SetDiskType(value *DiskType)()
    SetGroupPolicyBootTimeInMs(value *int32)()
    SetGroupPolicyLoginTimeInMs(value *int32)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetLoginScore(value *int32)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelStartupPerformanceScore(value UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_modelStartupPerformanceScoreable)()
    SetOperatingSystemVersion(value *string)()
    SetResponsiveDesktopTimeInMs(value *int32)()
    SetRestartCount(value *int32)()
    SetStartupPerformanceScore(value UserExperienceAnalyticsDevicePerformance_UserExperienceAnalyticsDevicePerformance_startupPerformanceScoreable)()
}
