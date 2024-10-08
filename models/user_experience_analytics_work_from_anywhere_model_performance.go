package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsWorkFromAnywhereModelPerformance the user experience analytics work from anywhere model performance.
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance struct {
    Entity
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreable interface {
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
// NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance instantiates a new UserExperienceAnalyticsWorkFromAnywhereModelPerformance and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance) {
    m := &UserExperienceAnalyticsWorkFromAnywhereModelPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance(), nil
}
// GetCloudIdentityScore gets the cloudIdentityScore property value. The cloud identity score of the device model. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudIdentityScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreable) {
    val, err := m.GetBackingStore().Get("cloudIdentityScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreable)
    }
    return nil
}
// GetCloudManagementScore gets the cloudManagementScore property value. The cloud management score of the device model. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudManagementScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreable) {
    val, err := m.GetBackingStore().Get("cloudManagementScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreable)
    }
    return nil
}
// GetCloudProvisioningScore gets the cloudProvisioningScore property value. The cloud provisioning score of the device model.  Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudProvisioningScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreable) {
    val, err := m.GetBackingStore().Get("cloudProvisioningScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudIdentityScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudIdentityScore(val.(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScore))
        }
        return nil
    }
    res["cloudManagementScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudManagementScore(val.(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScore))
        }
        return nil
    }
    res["cloudProvisioningScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudProvisioningScore(val.(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScore))
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
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelDeviceCount(val)
        }
        return nil
    }
    res["windowsScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsScore(val.(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScore))
        }
        return nil
    }
    res["workFromAnywhereScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFromAnywhereScore(val.(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScore))
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The manufacturer name of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetManufacturer()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModelDeviceCount gets the modelDeviceCount property value. The devices count for the model. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetModelDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("modelDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWindowsScore gets the windowsScore property value. The window score of the device model. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetWindowsScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreable) {
    val, err := m.GetBackingStore().Get("windowsScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreable)
    }
    return nil
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. The work from anywhere score of the device model. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetWorkFromAnywhereScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreable) {
    val, err := m.GetBackingStore().Get("workFromAnywhereScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("cloudIdentityScore", m.GetCloudIdentityScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cloudManagementScore", m.GetCloudManagementScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("cloudProvisioningScore", m.GetCloudProvisioningScore())
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
        err = writer.WriteInt32Value("modelDeviceCount", m.GetModelDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsScore", m.GetWindowsScore())
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
// SetCloudIdentityScore sets the cloudIdentityScore property value. The cloud identity score of the device model. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudIdentityScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreable)() {
    err := m.GetBackingStore().Set("cloudIdentityScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudManagementScore sets the cloudManagementScore property value. The cloud management score of the device model. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudManagementScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreable)() {
    err := m.GetBackingStore().Set("cloudManagementScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudProvisioningScore sets the cloudProvisioningScore property value. The cloud provisioning score of the device model.  Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudProvisioningScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreable)() {
    err := m.GetBackingStore().Set("cloudProvisioningScore", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The model name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetModelDeviceCount sets the modelDeviceCount property value. The devices count for the model. Supports: $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetModelDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("modelDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsScore sets the windowsScore property value. The window score of the device model. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetWindowsScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreable)() {
    err := m.GetBackingStore().Set("windowsScore", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. The work from anywhere score of the device model. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetWorkFromAnywhereScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreable)() {
    err := m.GetBackingStore().Set("workFromAnywhereScore", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudIdentityScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreable)
    GetCloudManagementScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreable)
    GetCloudProvisioningScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreable)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelDeviceCount()(*int32)
    GetWindowsScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreable)
    GetWorkFromAnywhereScore()(UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreable)
    SetCloudIdentityScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudIdentityScoreable)()
    SetCloudManagementScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudManagementScoreable)()
    SetCloudProvisioningScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_cloudProvisioningScoreable)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelDeviceCount(value *int32)()
    SetWindowsScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_windowsScoreable)()
    SetWorkFromAnywhereScore(value UserExperienceAnalyticsWorkFromAnywhereModelPerformance_UserExperienceAnalyticsWorkFromAnywhereModelPerformance_workFromAnywhereScoreable)()
}
