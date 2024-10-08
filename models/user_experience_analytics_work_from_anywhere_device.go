package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsWorkFromAnywhereDevice the user experience analytics device for work from anywhere report.
type UserExperienceAnalyticsWorkFromAnywhereDevice struct {
    Entity
}
// UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore()(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore()(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore()(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore()(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore instantiates a new UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore()(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore()
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreable interface {
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
type UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreable interface {
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
// NewUserExperienceAnalyticsWorkFromAnywhereDevice instantiates a new UserExperienceAnalyticsWorkFromAnywhereDevice and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevice()(*UserExperienceAnalyticsWorkFromAnywhereDevice) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereDevice(), nil
}
// GetAutoPilotProfileAssigned gets the autoPilotProfileAssigned property value. When TRUE, indicates the intune device's autopilot profile is assigned. When FALSE, indicates it's not Assigned. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAutoPilotProfileAssigned()(*bool) {
    val, err := m.GetBackingStore().Get("autoPilotProfileAssigned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAutoPilotRegistered gets the autoPilotRegistered property value. When TRUE, indicates the intune device's autopilot is registered. When FALSE, indicates it's not registered. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAutoPilotRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("autoPilotRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. The Azure Active Directory (Azure AD) device Id. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("azureAdDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureAdJoinType gets the azureAdJoinType property value. The work from anywhere device's Azure Active Directory (Azure AD) join type. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdJoinType()(*string) {
    val, err := m.GetBackingStore().Get("azureAdJoinType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureAdRegistered gets the azureAdRegistered property value. When TRUE, indicates the device's Azure Active Directory (Azure AD) is registered. When False, indicates it's not registered. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("azureAdRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCloudIdentityScore gets the cloudIdentityScore property value. Indicates per device cloud identity score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudIdentityScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreable) {
    val, err := m.GetBackingStore().Get("cloudIdentityScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreable)
    }
    return nil
}
// GetCloudManagementScore gets the cloudManagementScore property value. Indicates per device cloud management score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudManagementScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreable) {
    val, err := m.GetBackingStore().Get("cloudManagementScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreable)
    }
    return nil
}
// GetCloudProvisioningScore gets the cloudProvisioningScore property value. Indicates per device cloud provisioning score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudProvisioningScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreable) {
    val, err := m.GetBackingStore().Get("cloudProvisioningScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreable)
    }
    return nil
}
// GetCompliancePolicySetToIntune gets the compliancePolicySetToIntune property value. When TRUE, indicates the device's compliance policy is set to intune. When FALSE, indicates it's not set to intune. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCompliancePolicySetToIntune()(*bool) {
    val, err := m.GetBackingStore().Get("compliancePolicySetToIntune")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The name of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["autoPilotProfileAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoPilotProfileAssigned(val)
        }
        return nil
    }
    res["autoPilotRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoPilotRegistered(val)
        }
        return nil
    }
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    res["azureAdJoinType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdJoinType(val)
        }
        return nil
    }
    res["azureAdRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdRegistered(val)
        }
        return nil
    }
    res["cloudIdentityScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudIdentityScore(val.(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScore))
        }
        return nil
    }
    res["cloudManagementScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudManagementScore(val.(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScore))
        }
        return nil
    }
    res["cloudProvisioningScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudProvisioningScore(val.(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScore))
        }
        return nil
    }
    res["compliancePolicySetToIntune"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicySetToIntune(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["isCloudManagedGatewayEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCloudManagedGatewayEnabled(val)
        }
        return nil
    }
    res["managedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBy(val)
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
    res["osCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsCheckFailed(val)
        }
        return nil
    }
    res["osDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["otherWorkloadsSetToIntune"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherWorkloadsSetToIntune(val)
        }
        return nil
    }
    res["ownership"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnership(val)
        }
        return nil
    }
    res["processor64BitCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessor64BitCheckFailed(val)
        }
        return nil
    }
    res["processorCoreCountCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorCoreCountCheckFailed(val)
        }
        return nil
    }
    res["processorFamilyCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorFamilyCheckFailed(val)
        }
        return nil
    }
    res["processorSpeedCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorSpeedCheckFailed(val)
        }
        return nil
    }
    res["ramCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamCheckFailed(val)
        }
        return nil
    }
    res["secureBootCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBootCheckFailed(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["storageCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageCheckFailed(val)
        }
        return nil
    }
    res["tenantAttached"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAttached(val)
        }
        return nil
    }
    res["tpmCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmCheckFailed(val)
        }
        return nil
    }
    res["upgradeEligibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperatingSystemUpgradeEligibility)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeEligibility(val.(*OperatingSystemUpgradeEligibility))
        }
        return nil
    }
    res["windowsScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsScore(val.(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScore))
        }
        return nil
    }
    res["workFromAnywhereScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFromAnywhereScore(val.(*UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScore))
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
// returns a *UserExperienceAnalyticsHealthState when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    val, err := m.GetBackingStore().Get("healthStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsHealthState)
    }
    return nil
}
// GetIsCloudManagedGatewayEnabled gets the isCloudManagedGatewayEnabled property value. When TRUE, indicates the device's Cloud Management Gateway for Configuration Manager is enabled. When FALSE, indicates it's not enabled. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetIsCloudManagedGatewayEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isCloudManagedGatewayEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManagedBy gets the managedBy property value. The management agent of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetManagedBy()(*string) {
    val, err := m.GetBackingStore().Get("managedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The manufacturer name of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetManufacturer()(*string) {
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsCheckFailed gets the osCheckFailed property value. When TRUE, indicates OS check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("osCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOsDescription gets the osDescription property value. The OS description of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsDescription()(*string) {
    val, err := m.GetBackingStore().Get("osDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The OS version of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOtherWorkloadsSetToIntune gets the otherWorkloadsSetToIntune property value. When TRUE, indicates the device's other workloads is set to intune. When FALSE, indicates it's not set to intune. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOtherWorkloadsSetToIntune()(*bool) {
    val, err := m.GetBackingStore().Get("otherWorkloadsSetToIntune")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOwnership gets the ownership property value. Ownership of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOwnership()(*string) {
    val, err := m.GetBackingStore().Get("ownership")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProcessor64BitCheckFailed gets the processor64BitCheckFailed property value. When TRUE, indicates processor hardware 64-bit architecture check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessor64BitCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("processor64BitCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProcessorCoreCountCheckFailed gets the processorCoreCountCheckFailed property value. When TRUE, indicates processor hardware core count check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorCoreCountCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("processorCoreCountCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProcessorFamilyCheckFailed gets the processorFamilyCheckFailed property value. When TRUE, indicates processor hardware family check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorFamilyCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("processorFamilyCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProcessorSpeedCheckFailed gets the processorSpeedCheckFailed property value. When TRUE, indicates processor hardware speed check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorSpeedCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("processorSpeedCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRamCheckFailed gets the ramCheckFailed property value. When TRUE, indicates RAM hardware check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetRamCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("ramCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSecureBootCheckFailed gets the secureBootCheckFailed property value. When TRUE, indicates secure boot hardware check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetSecureBootCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("secureBootCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. The serial number of the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStorageCheckFailed gets the storageCheckFailed property value. When TRUE, indicates storage hardware check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetStorageCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("storageCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTenantAttached gets the tenantAttached property value. When TRUE, indicates the device is Tenant Attached. When FALSE, indicates it's not Tenant Attached. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetTenantAttached()(*bool) {
    val, err := m.GetBackingStore().Get("tenantAttached")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTpmCheckFailed gets the tpmCheckFailed property value. When TRUE, indicates Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade to windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
// returns a *bool when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetTpmCheckFailed()(*bool) {
    val, err := m.GetBackingStore().Get("tpmCheckFailed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUpgradeEligibility gets the upgradeEligibility property value. Work From Anywhere windows device upgrade eligibility status.
// returns a *OperatingSystemUpgradeEligibility when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetUpgradeEligibility()(*OperatingSystemUpgradeEligibility) {
    val, err := m.GetBackingStore().Get("upgradeEligibility")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OperatingSystemUpgradeEligibility)
    }
    return nil
}
// GetWindowsScore gets the windowsScore property value. Indicates per device windows score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetWindowsScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreable) {
    val, err := m.GetBackingStore().Get("windowsScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreable)
    }
    return nil
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. Indicates work from anywhere per device overall score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreable when successful
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetWorkFromAnywhereScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreable) {
    val, err := m.GetBackingStore().Get("workFromAnywhereScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("autoPilotProfileAssigned", m.GetAutoPilotProfileAssigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoPilotRegistered", m.GetAutoPilotRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureAdJoinType", m.GetAzureAdJoinType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("azureAdRegistered", m.GetAzureAdRegistered())
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteBoolValue("compliancePolicySetToIntune", m.GetCompliancePolicySetToIntune())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
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
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCloudManagedGatewayEnabled", m.GetIsCloudManagedGatewayEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedBy", m.GetManagedBy())
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
        err = writer.WriteBoolValue("osCheckFailed", m.GetOsCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("otherWorkloadsSetToIntune", m.GetOtherWorkloadsSetToIntune())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownership", m.GetOwnership())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processor64BitCheckFailed", m.GetProcessor64BitCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorCoreCountCheckFailed", m.GetProcessorCoreCountCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorFamilyCheckFailed", m.GetProcessorFamilyCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorSpeedCheckFailed", m.GetProcessorSpeedCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ramCheckFailed", m.GetRamCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("secureBootCheckFailed", m.GetSecureBootCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageCheckFailed", m.GetStorageCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("tenantAttached", m.GetTenantAttached())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("tpmCheckFailed", m.GetTpmCheckFailed())
        if err != nil {
            return err
        }
    }
    if m.GetUpgradeEligibility() != nil {
        cast := (*m.GetUpgradeEligibility()).String()
        err = writer.WriteStringValue("upgradeEligibility", &cast)
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
// SetAutoPilotProfileAssigned sets the autoPilotProfileAssigned property value. When TRUE, indicates the intune device's autopilot profile is assigned. When FALSE, indicates it's not Assigned. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAutoPilotProfileAssigned(value *bool)() {
    err := m.GetBackingStore().Set("autoPilotProfileAssigned", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoPilotRegistered sets the autoPilotRegistered property value. When TRUE, indicates the intune device's autopilot is registered. When FALSE, indicates it's not registered. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAutoPilotRegistered(value *bool)() {
    err := m.GetBackingStore().Set("autoPilotRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. The Azure Active Directory (Azure AD) device Id. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdDeviceId(value *string)() {
    err := m.GetBackingStore().Set("azureAdDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureAdJoinType sets the azureAdJoinType property value. The work from anywhere device's Azure Active Directory (Azure AD) join type. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdJoinType(value *string)() {
    err := m.GetBackingStore().Set("azureAdJoinType", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureAdRegistered sets the azureAdRegistered property value. When TRUE, indicates the device's Azure Active Directory (Azure AD) is registered. When False, indicates it's not registered. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdRegistered(value *bool)() {
    err := m.GetBackingStore().Set("azureAdRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudIdentityScore sets the cloudIdentityScore property value. Indicates per device cloud identity score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudIdentityScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreable)() {
    err := m.GetBackingStore().Set("cloudIdentityScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudManagementScore sets the cloudManagementScore property value. Indicates per device cloud management score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudManagementScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreable)() {
    err := m.GetBackingStore().Set("cloudManagementScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudProvisioningScore sets the cloudProvisioningScore property value. Indicates per device cloud provisioning score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudProvisioningScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreable)() {
    err := m.GetBackingStore().Set("cloudProvisioningScore", value)
    if err != nil {
        panic(err)
    }
}
// SetCompliancePolicySetToIntune sets the compliancePolicySetToIntune property value. When TRUE, indicates the device's compliance policy is set to intune. When FALSE, indicates it's not set to intune. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCompliancePolicySetToIntune(value *bool)() {
    err := m.GetBackingStore().Set("compliancePolicySetToIntune", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The Intune device id of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    err := m.GetBackingStore().Set("healthStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCloudManagedGatewayEnabled sets the isCloudManagedGatewayEnabled property value. When TRUE, indicates the device's Cloud Management Gateway for Configuration Manager is enabled. When FALSE, indicates it's not enabled. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetIsCloudManagedGatewayEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isCloudManagedGatewayEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedBy sets the managedBy property value. The management agent of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetManagedBy(value *string)() {
    err := m.GetBackingStore().Set("managedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The model name of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetOsCheckFailed sets the osCheckFailed property value. When TRUE, indicates OS check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("osCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetOsDescription sets the osDescription property value. The OS description of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsDescription(value *string)() {
    err := m.GetBackingStore().Set("osDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The OS version of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOtherWorkloadsSetToIntune sets the otherWorkloadsSetToIntune property value. When TRUE, indicates the device's other workloads is set to intune. When FALSE, indicates it's not set to intune. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOtherWorkloadsSetToIntune(value *bool)() {
    err := m.GetBackingStore().Set("otherWorkloadsSetToIntune", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnership sets the ownership property value. Ownership of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOwnership(value *string)() {
    err := m.GetBackingStore().Set("ownership", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessor64BitCheckFailed sets the processor64BitCheckFailed property value. When TRUE, indicates processor hardware 64-bit architecture check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessor64BitCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("processor64BitCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessorCoreCountCheckFailed sets the processorCoreCountCheckFailed property value. When TRUE, indicates processor hardware core count check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorCoreCountCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("processorCoreCountCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessorFamilyCheckFailed sets the processorFamilyCheckFailed property value. When TRUE, indicates processor hardware family check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorFamilyCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("processorFamilyCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessorSpeedCheckFailed sets the processorSpeedCheckFailed property value. When TRUE, indicates processor hardware speed check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorSpeedCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("processorSpeedCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetRamCheckFailed sets the ramCheckFailed property value. When TRUE, indicates RAM hardware check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetRamCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("ramCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetSecureBootCheckFailed sets the secureBootCheckFailed property value. When TRUE, indicates secure boot hardware check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetSecureBootCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("secureBootCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. The serial number of the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetStorageCheckFailed sets the storageCheckFailed property value. When TRUE, indicates storage hardware check failed for device to upgrade to the latest version of windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetStorageCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("storageCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantAttached sets the tenantAttached property value. When TRUE, indicates the device is Tenant Attached. When FALSE, indicates it's not Tenant Attached. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetTenantAttached(value *bool)() {
    err := m.GetBackingStore().Set("tenantAttached", value)
    if err != nil {
        panic(err)
    }
}
// SetTpmCheckFailed sets the tpmCheckFailed property value. When TRUE, indicates Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade to windows. When FALSE, indicates the check succeeded. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetTpmCheckFailed(value *bool)() {
    err := m.GetBackingStore().Set("tpmCheckFailed", value)
    if err != nil {
        panic(err)
    }
}
// SetUpgradeEligibility sets the upgradeEligibility property value. Work From Anywhere windows device upgrade eligibility status.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetUpgradeEligibility(value *OperatingSystemUpgradeEligibility)() {
    err := m.GetBackingStore().Set("upgradeEligibility", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsScore sets the windowsScore property value. Indicates per device windows score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetWindowsScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreable)() {
    err := m.GetBackingStore().Set("windowsScore", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. Indicates work from anywhere per device overall score. Valid values 0 to 100. Value -1 means associated score is unavailable. Supports: $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetWorkFromAnywhereScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreable)() {
    err := m.GetBackingStore().Set("workFromAnywhereScore", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsWorkFromAnywhereDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoPilotProfileAssigned()(*bool)
    GetAutoPilotRegistered()(*bool)
    GetAzureAdDeviceId()(*string)
    GetAzureAdJoinType()(*string)
    GetAzureAdRegistered()(*bool)
    GetCloudIdentityScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreable)
    GetCloudManagementScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreable)
    GetCloudProvisioningScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreable)
    GetCompliancePolicySetToIntune()(*bool)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetIsCloudManagedGatewayEnabled()(*bool)
    GetManagedBy()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetOsCheckFailed()(*bool)
    GetOsDescription()(*string)
    GetOsVersion()(*string)
    GetOtherWorkloadsSetToIntune()(*bool)
    GetOwnership()(*string)
    GetProcessor64BitCheckFailed()(*bool)
    GetProcessorCoreCountCheckFailed()(*bool)
    GetProcessorFamilyCheckFailed()(*bool)
    GetProcessorSpeedCheckFailed()(*bool)
    GetRamCheckFailed()(*bool)
    GetSecureBootCheckFailed()(*bool)
    GetSerialNumber()(*string)
    GetStorageCheckFailed()(*bool)
    GetTenantAttached()(*bool)
    GetTpmCheckFailed()(*bool)
    GetUpgradeEligibility()(*OperatingSystemUpgradeEligibility)
    GetWindowsScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreable)
    GetWorkFromAnywhereScore()(UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreable)
    SetAutoPilotProfileAssigned(value *bool)()
    SetAutoPilotRegistered(value *bool)()
    SetAzureAdDeviceId(value *string)()
    SetAzureAdJoinType(value *string)()
    SetAzureAdRegistered(value *bool)()
    SetCloudIdentityScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudIdentityScoreable)()
    SetCloudManagementScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudManagementScoreable)()
    SetCloudProvisioningScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_cloudProvisioningScoreable)()
    SetCompliancePolicySetToIntune(value *bool)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetIsCloudManagedGatewayEnabled(value *bool)()
    SetManagedBy(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetOsCheckFailed(value *bool)()
    SetOsDescription(value *string)()
    SetOsVersion(value *string)()
    SetOtherWorkloadsSetToIntune(value *bool)()
    SetOwnership(value *string)()
    SetProcessor64BitCheckFailed(value *bool)()
    SetProcessorCoreCountCheckFailed(value *bool)()
    SetProcessorFamilyCheckFailed(value *bool)()
    SetProcessorSpeedCheckFailed(value *bool)()
    SetRamCheckFailed(value *bool)()
    SetSecureBootCheckFailed(value *bool)()
    SetSerialNumber(value *string)()
    SetStorageCheckFailed(value *bool)()
    SetTenantAttached(value *bool)()
    SetTpmCheckFailed(value *bool)()
    SetUpgradeEligibility(value *OperatingSystemUpgradeEligibility)()
    SetWindowsScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_windowsScoreable)()
    SetWorkFromAnywhereScore(value UserExperienceAnalyticsWorkFromAnywhereDevice_UserExperienceAnalyticsWorkFromAnywhereDevice_workFromAnywhereScoreable)()
}
