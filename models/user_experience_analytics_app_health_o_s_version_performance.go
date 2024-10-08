package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsAppHealthOSVersionPerformance the user experience analytics device OS version performance entity contains OS version performance details.
type UserExperienceAnalyticsAppHealthOSVersionPerformance struct {
    Entity
}
// UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore instantiates a new UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore and sets the default values.
func NewUserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore()(*UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) {
    m := &UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore()
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
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) GetDouble()(*float64) {
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
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) GetString()(*string) {
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
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreable interface {
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
// NewUserExperienceAnalyticsAppHealthOSVersionPerformance instantiates a new UserExperienceAnalyticsAppHealthOSVersionPerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthOSVersionPerformance()(*UserExperienceAnalyticsAppHealthOSVersionPerformance) {
    m := &UserExperienceAnalyticsAppHealthOSVersionPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthOSVersionPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsAppHealthOSVersionPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthOSVersionPerformance(), nil
}
// GetActiveDeviceCount gets the activeDeviceCount property value. The number of active devices for the OS version. Valid values 0 to 2147483647. Supports: $filter, $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetActiveDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("activeDeviceCount")
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
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDeviceCount(val)
        }
        return nil
    }
    res["meanTimeToFailureInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeanTimeToFailureInMinutes(val)
        }
        return nil
    }
    res["osBuildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
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
    res["osVersionAppHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersionAppHealthScore(val.(*UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScore))
        }
        return nil
    }
    return res
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The mean time to failure for the application in minutes. Valid values 0 to 2147483647. Supports: $filter, $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
// returns a *int32 when successful
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    val, err := m.GetBackingStore().Get("meanTimeToFailureInMinutes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOsBuildNumber gets the osBuildNumber property value. The OS build number installed on the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsBuildNumber()(*string) {
    val, err := m.GetBackingStore().Get("osBuildNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The OS version installed on the device. Supports: $select, $OrderBy. Read-only.
// returns a *string when successful
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersionAppHealthScore gets the osVersionAppHealthScore property value. The application health score of the OS version. Valid values 0 to 100. Supports: $filter, $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreable when successful
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsVersionAppHealthScore()(UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreable) {
    val, err := m.GetBackingStore().Get("osVersionAppHealthScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDeviceCount", m.GetActiveDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("meanTimeToFailureInMinutes", m.GetMeanTimeToFailureInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
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
        err = writer.WriteObjectValue("osVersionAppHealthScore", m.GetOsVersionAppHealthScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDeviceCount sets the activeDeviceCount property value. The number of active devices for the OS version. Valid values 0 to 2147483647. Supports: $filter, $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetActiveDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("activeDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the application in minutes. Valid values 0 to 2147483647. Supports: $filter, $select, $OrderBy. Read-only. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    err := m.GetBackingStore().Set("meanTimeToFailureInMinutes", value)
    if err != nil {
        panic(err)
    }
}
// SetOsBuildNumber sets the osBuildNumber property value. The OS build number installed on the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsBuildNumber(value *string)() {
    err := m.GetBackingStore().Set("osBuildNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The OS version installed on the device. Supports: $select, $OrderBy. Read-only.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersionAppHealthScore sets the osVersionAppHealthScore property value. The application health score of the OS version. Valid values 0 to 100. Supports: $filter, $select, $OrderBy. Read-only. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsVersionAppHealthScore(value UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreable)() {
    err := m.GetBackingStore().Set("osVersionAppHealthScore", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsAppHealthOSVersionPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDeviceCount()(*int32)
    GetMeanTimeToFailureInMinutes()(*int32)
    GetOsBuildNumber()(*string)
    GetOsVersion()(*string)
    GetOsVersionAppHealthScore()(UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreable)
    SetActiveDeviceCount(value *int32)()
    SetMeanTimeToFailureInMinutes(value *int32)()
    SetOsBuildNumber(value *string)()
    SetOsVersion(value *string)()
    SetOsVersionAppHealthScore(value UserExperienceAnalyticsAppHealthOSVersionPerformance_UserExperienceAnalyticsAppHealthOSVersionPerformance_osVersionAppHealthScoreable)()
}
