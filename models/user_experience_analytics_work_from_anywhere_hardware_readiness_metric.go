package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric the user experience analytics hardware readiness entity contains account level information about hardware blockers for windows upgrade.
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric struct {
    Entity
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageable interface {
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
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric instantiates a new UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["osCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentage))
        }
        return nil
    }
    res["processor64BitCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessor64BitCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentage))
        }
        return nil
    }
    res["processorCoreCountCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorCoreCountCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentage))
        }
        return nil
    }
    res["processorFamilyCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorFamilyCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentage))
        }
        return nil
    }
    res["processorSpeedCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorSpeedCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentage))
        }
        return nil
    }
    res["ramCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentage))
        }
        return nil
    }
    res["secureBootCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBootCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentage))
        }
        return nil
    }
    res["storageCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentage))
        }
        return nil
    }
    res["totalDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDeviceCount(val)
        }
        return nil
    }
    res["tpmCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmCheckFailedPercentage(val.(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentage))
        }
        return nil
    }
    res["upgradeEligibleDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeEligibleDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetOsCheckFailedPercentage gets the osCheckFailedPercentage property value. The percentage of devices for which OS check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetOsCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("osCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageable)
    }
    return nil
}
// GetProcessor64BitCheckFailedPercentage gets the processor64BitCheckFailedPercentage property value. The percentage of devices for which processor hardware 64-bit architecture check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessor64BitCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("processor64BitCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageable)
    }
    return nil
}
// GetProcessorCoreCountCheckFailedPercentage gets the processorCoreCountCheckFailedPercentage property value. The percentage of devices for which processor hardware core count check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorCoreCountCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("processorCoreCountCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageable)
    }
    return nil
}
// GetProcessorFamilyCheckFailedPercentage gets the processorFamilyCheckFailedPercentage property value. The percentage of devices for which processor hardware family check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorFamilyCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("processorFamilyCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageable)
    }
    return nil
}
// GetProcessorSpeedCheckFailedPercentage gets the processorSpeedCheckFailedPercentage property value. The percentage of devices for which processor hardware speed check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorSpeedCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("processorSpeedCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageable)
    }
    return nil
}
// GetRamCheckFailedPercentage gets the ramCheckFailedPercentage property value. The percentage of devices for which RAM hardware check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetRamCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("ramCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageable)
    }
    return nil
}
// GetSecureBootCheckFailedPercentage gets the secureBootCheckFailedPercentage property value. The percentage of devices for which secure boot hardware check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetSecureBootCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("secureBootCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageable)
    }
    return nil
}
// GetStorageCheckFailedPercentage gets the storageCheckFailedPercentage property value. The percentage of devices for which storage hardware check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetStorageCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("storageCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageable)
    }
    return nil
}
// GetTotalDeviceCount gets the totalDeviceCount property value. The count of total devices in an organization. Valid values 0 to 2147483647. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetTotalDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTpmCheckFailedPercentage gets the tpmCheckFailedPercentage property value. The percentage of devices for which Trusted Platform Module (TPM) hardware check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetTpmCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageable) {
    val, err := m.GetBackingStore().Get("tpmCheckFailedPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageable)
    }
    return nil
}
// GetUpgradeEligibleDeviceCount gets the upgradeEligibleDeviceCount property value. The count of devices in an organization eligible for windows upgrade. Valid values 0 to 2147483647. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetUpgradeEligibleDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("upgradeEligibleDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("osCheckFailedPercentage", m.GetOsCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("processor64BitCheckFailedPercentage", m.GetProcessor64BitCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("processorCoreCountCheckFailedPercentage", m.GetProcessorCoreCountCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("processorFamilyCheckFailedPercentage", m.GetProcessorFamilyCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("processorSpeedCheckFailedPercentage", m.GetProcessorSpeedCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ramCheckFailedPercentage", m.GetRamCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("secureBootCheckFailedPercentage", m.GetSecureBootCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("storageCheckFailedPercentage", m.GetStorageCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalDeviceCount", m.GetTotalDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tpmCheckFailedPercentage", m.GetTpmCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("upgradeEligibleDeviceCount", m.GetUpgradeEligibleDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOsCheckFailedPercentage sets the osCheckFailedPercentage property value. The percentage of devices for which OS check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetOsCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("osCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessor64BitCheckFailedPercentage sets the processor64BitCheckFailedPercentage property value. The percentage of devices for which processor hardware 64-bit architecture check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessor64BitCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("processor64BitCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessorCoreCountCheckFailedPercentage sets the processorCoreCountCheckFailedPercentage property value. The percentage of devices for which processor hardware core count check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorCoreCountCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("processorCoreCountCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessorFamilyCheckFailedPercentage sets the processorFamilyCheckFailedPercentage property value. The percentage of devices for which processor hardware family check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorFamilyCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("processorFamilyCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessorSpeedCheckFailedPercentage sets the processorSpeedCheckFailedPercentage property value. The percentage of devices for which processor hardware speed check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorSpeedCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("processorSpeedCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetRamCheckFailedPercentage sets the ramCheckFailedPercentage property value. The percentage of devices for which RAM hardware check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetRamCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("ramCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetSecureBootCheckFailedPercentage sets the secureBootCheckFailedPercentage property value. The percentage of devices for which secure boot hardware check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetSecureBootCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("secureBootCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageCheckFailedPercentage sets the storageCheckFailedPercentage property value. The percentage of devices for which storage hardware check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetStorageCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("storageCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalDeviceCount sets the totalDeviceCount property value. The count of total devices in an organization. Valid values 0 to 2147483647. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetTotalDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("totalDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTpmCheckFailedPercentage sets the tpmCheckFailedPercentage property value. The percentage of devices for which Trusted Platform Module (TPM) hardware check has failed. Valid values 0 to 100. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetTpmCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageable)() {
    err := m.GetBackingStore().Set("tpmCheckFailedPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetUpgradeEligibleDeviceCount sets the upgradeEligibleDeviceCount property value. The count of devices in an organization eligible for windows upgrade. Valid values 0 to 2147483647. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetUpgradeEligibleDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("upgradeEligibleDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOsCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageable)
    GetProcessor64BitCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageable)
    GetProcessorCoreCountCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageable)
    GetProcessorFamilyCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageable)
    GetProcessorSpeedCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageable)
    GetRamCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageable)
    GetSecureBootCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageable)
    GetStorageCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageable)
    GetTotalDeviceCount()(*int32)
    GetTpmCheckFailedPercentage()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageable)
    GetUpgradeEligibleDeviceCount()(*int32)
    SetOsCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_osCheckFailedPercentageable)()
    SetProcessor64BitCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processor64BitCheckFailedPercentageable)()
    SetProcessorCoreCountCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorCoreCountCheckFailedPercentageable)()
    SetProcessorFamilyCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorFamilyCheckFailedPercentageable)()
    SetProcessorSpeedCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_processorSpeedCheckFailedPercentageable)()
    SetRamCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_ramCheckFailedPercentageable)()
    SetSecureBootCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_secureBootCheckFailedPercentageable)()
    SetStorageCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_storageCheckFailedPercentageable)()
    SetTotalDeviceCount(value *int32)()
    SetTpmCheckFailedPercentage(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric_tpmCheckFailedPercentageable)()
    SetUpgradeEligibleDeviceCount(value *int32)()
}
