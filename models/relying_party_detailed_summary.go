package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type RelyingPartyDetailedSummary struct {
    Entity
}
// RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate composed type wrapper for classes float64, ReferenceNumeric, string
type RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewRelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate instantiates a new RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate and sets the default values.
func NewRelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate()(*RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) {
    m := &RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateRelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewRelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate()
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
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) GetDouble()(*float64) {
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
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) GetString()(*string) {
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
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateable interface {
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
// NewRelyingPartyDetailedSummary instantiates a new RelyingPartyDetailedSummary and sets the default values.
func NewRelyingPartyDetailedSummary()(*RelyingPartyDetailedSummary) {
    m := &RelyingPartyDetailedSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRelyingPartyDetailedSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelyingPartyDetailedSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelyingPartyDetailedSummary(), nil
}
// GetFailedSignInCount gets the failedSignInCount property value. Number of failed sign ins on AD FS in the period specified. Supports $orderby, $filter (eq).
// returns a *int64 when successful
func (m *RelyingPartyDetailedSummary) GetFailedSignInCount()(*int64) {
    val, err := m.GetBackingStore().Get("failedSignInCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RelyingPartyDetailedSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedSignInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedSignInCount(val)
        }
        return nil
    }
    res["migrationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMigrationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMigrationStatus(val.(*MigrationStatus))
        }
        return nil
    }
    res["migrationValidationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(KeyValuePairable)
                }
            }
            m.SetMigrationValidationDetails(res)
        }
        return nil
    }
    res["relyingPartyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelyingPartyId(val)
        }
        return nil
    }
    res["relyingPartyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelyingPartyName(val)
        }
        return nil
    }
    res["replyUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetReplyUrls(res)
        }
        return nil
    }
    res["serviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceId(val)
        }
        return nil
    }
    res["signInSuccessRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInSuccessRate(val.(*RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRate))
        }
        return nil
    }
    res["successfulSignInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulSignInCount(val)
        }
        return nil
    }
    res["totalSignInCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSignInCount(val)
        }
        return nil
    }
    res["uniqueUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueUserCount(val)
        }
        return nil
    }
    return res
}
// GetMigrationStatus gets the migrationStatus property value. The migrationStatus property
// returns a *MigrationStatus when successful
func (m *RelyingPartyDetailedSummary) GetMigrationStatus()(*MigrationStatus) {
    val, err := m.GetBackingStore().Get("migrationStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MigrationStatus)
    }
    return nil
}
// GetMigrationValidationDetails gets the migrationValidationDetails property value. Specifies all the validations checks done on applications config details.
// returns a []KeyValuePairable when successful
func (m *RelyingPartyDetailedSummary) GetMigrationValidationDetails()([]KeyValuePairable) {
    val, err := m.GetBackingStore().Get("migrationValidationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]KeyValuePairable)
    }
    return nil
}
// GetRelyingPartyId gets the relyingPartyId property value. Identifies the relying party to this federation service. It's used when issuing claims to the relying party. Supports $orderby, $filter (eq).
// returns a *string when successful
func (m *RelyingPartyDetailedSummary) GetRelyingPartyId()(*string) {
    val, err := m.GetBackingStore().Get("relyingPartyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRelyingPartyName gets the relyingPartyName property value. Name of the relying party's website or other entity on the Internet that uses an identity provider to authenticate a user who wants to log in. Supports $orderby, $filter (eq).
// returns a *string when successful
func (m *RelyingPartyDetailedSummary) GetRelyingPartyName()(*string) {
    val, err := m.GetBackingStore().Get("relyingPartyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReplyUrls gets the replyUrls property value. Specifies where the relying party expects to receive the token.
// returns a []string when successful
func (m *RelyingPartyDetailedSummary) GetReplyUrls()([]string) {
    val, err := m.GetBackingStore().Get("replyUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetServiceId gets the serviceId property value. Uniquely identifies the Active Directory forest. Supports $orderby, $filter (eq).
// returns a *string when successful
func (m *RelyingPartyDetailedSummary) GetServiceId()(*string) {
    val, err := m.GetBackingStore().Get("serviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSignInSuccessRate gets the signInSuccessRate property value. Calculated as Number of successful / (Number of successful + Number of failed sign ins) or successfulSignInCount / totalSignInCount on AD FS in the period specified. Supports $orderby, $filter (eq).
// returns a RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateable when successful
func (m *RelyingPartyDetailedSummary) GetSignInSuccessRate()(RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateable) {
    val, err := m.GetBackingStore().Get("signInSuccessRate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateable)
    }
    return nil
}
// GetSuccessfulSignInCount gets the successfulSignInCount property value. Number of successful sign ins on AD FS. Supports $orderby, $filter (eq).
// returns a *int64 when successful
func (m *RelyingPartyDetailedSummary) GetSuccessfulSignInCount()(*int64) {
    val, err := m.GetBackingStore().Get("successfulSignInCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTotalSignInCount gets the totalSignInCount property value. Number of successful + failed sign ins on AD FS in the period specified. Supports $orderby, $filter (eq).
// returns a *int64 when successful
func (m *RelyingPartyDetailedSummary) GetTotalSignInCount()(*int64) {
    val, err := m.GetBackingStore().Get("totalSignInCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetUniqueUserCount gets the uniqueUserCount property value. Number of unique users that signed into the application. Supports $orderby, $filter (eq).
// returns a *int64 when successful
func (m *RelyingPartyDetailedSummary) GetUniqueUserCount()(*int64) {
    val, err := m.GetBackingStore().Get("uniqueUserCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RelyingPartyDetailedSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("failedSignInCount", m.GetFailedSignInCount())
        if err != nil {
            return err
        }
    }
    if m.GetMigrationStatus() != nil {
        cast := (*m.GetMigrationStatus()).String()
        err = writer.WriteStringValue("migrationStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMigrationValidationDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMigrationValidationDetails()))
        for i, v := range m.GetMigrationValidationDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("migrationValidationDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("relyingPartyId", m.GetRelyingPartyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("relyingPartyName", m.GetRelyingPartyName())
        if err != nil {
            return err
        }
    }
    if m.GetReplyUrls() != nil {
        err = writer.WriteCollectionOfStringValues("replyUrls", m.GetReplyUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("signInSuccessRate", m.GetSignInSuccessRate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("successfulSignInCount", m.GetSuccessfulSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalSignInCount", m.GetTotalSignInCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("uniqueUserCount", m.GetUniqueUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFailedSignInCount sets the failedSignInCount property value. Number of failed sign ins on AD FS in the period specified. Supports $orderby, $filter (eq).
func (m *RelyingPartyDetailedSummary) SetFailedSignInCount(value *int64)() {
    err := m.GetBackingStore().Set("failedSignInCount", value)
    if err != nil {
        panic(err)
    }
}
// SetMigrationStatus sets the migrationStatus property value. The migrationStatus property
func (m *RelyingPartyDetailedSummary) SetMigrationStatus(value *MigrationStatus)() {
    err := m.GetBackingStore().Set("migrationStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetMigrationValidationDetails sets the migrationValidationDetails property value. Specifies all the validations checks done on applications config details.
func (m *RelyingPartyDetailedSummary) SetMigrationValidationDetails(value []KeyValuePairable)() {
    err := m.GetBackingStore().Set("migrationValidationDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetRelyingPartyId sets the relyingPartyId property value. Identifies the relying party to this federation service. It's used when issuing claims to the relying party. Supports $orderby, $filter (eq).
func (m *RelyingPartyDetailedSummary) SetRelyingPartyId(value *string)() {
    err := m.GetBackingStore().Set("relyingPartyId", value)
    if err != nil {
        panic(err)
    }
}
// SetRelyingPartyName sets the relyingPartyName property value. Name of the relying party's website or other entity on the Internet that uses an identity provider to authenticate a user who wants to log in. Supports $orderby, $filter (eq).
func (m *RelyingPartyDetailedSummary) SetRelyingPartyName(value *string)() {
    err := m.GetBackingStore().Set("relyingPartyName", value)
    if err != nil {
        panic(err)
    }
}
// SetReplyUrls sets the replyUrls property value. Specifies where the relying party expects to receive the token.
func (m *RelyingPartyDetailedSummary) SetReplyUrls(value []string)() {
    err := m.GetBackingStore().Set("replyUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceId sets the serviceId property value. Uniquely identifies the Active Directory forest. Supports $orderby, $filter (eq).
func (m *RelyingPartyDetailedSummary) SetServiceId(value *string)() {
    err := m.GetBackingStore().Set("serviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetSignInSuccessRate sets the signInSuccessRate property value. Calculated as Number of successful / (Number of successful + Number of failed sign ins) or successfulSignInCount / totalSignInCount on AD FS in the period specified. Supports $orderby, $filter (eq).
func (m *RelyingPartyDetailedSummary) SetSignInSuccessRate(value RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateable)() {
    err := m.GetBackingStore().Set("signInSuccessRate", value)
    if err != nil {
        panic(err)
    }
}
// SetSuccessfulSignInCount sets the successfulSignInCount property value. Number of successful sign ins on AD FS. Supports $orderby, $filter (eq).
func (m *RelyingPartyDetailedSummary) SetSuccessfulSignInCount(value *int64)() {
    err := m.GetBackingStore().Set("successfulSignInCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalSignInCount sets the totalSignInCount property value. Number of successful + failed sign ins on AD FS in the period specified. Supports $orderby, $filter (eq).
func (m *RelyingPartyDetailedSummary) SetTotalSignInCount(value *int64)() {
    err := m.GetBackingStore().Set("totalSignInCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUniqueUserCount sets the uniqueUserCount property value. Number of unique users that signed into the application. Supports $orderby, $filter (eq).
func (m *RelyingPartyDetailedSummary) SetUniqueUserCount(value *int64)() {
    err := m.GetBackingStore().Set("uniqueUserCount", value)
    if err != nil {
        panic(err)
    }
}
type RelyingPartyDetailedSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFailedSignInCount()(*int64)
    GetMigrationStatus()(*MigrationStatus)
    GetMigrationValidationDetails()([]KeyValuePairable)
    GetRelyingPartyId()(*string)
    GetRelyingPartyName()(*string)
    GetReplyUrls()([]string)
    GetServiceId()(*string)
    GetSignInSuccessRate()(RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateable)
    GetSuccessfulSignInCount()(*int64)
    GetTotalSignInCount()(*int64)
    GetUniqueUserCount()(*int64)
    SetFailedSignInCount(value *int64)()
    SetMigrationStatus(value *MigrationStatus)()
    SetMigrationValidationDetails(value []KeyValuePairable)()
    SetRelyingPartyId(value *string)()
    SetRelyingPartyName(value *string)()
    SetReplyUrls(value []string)()
    SetServiceId(value *string)()
    SetSignInSuccessRate(value RelyingPartyDetailedSummary_RelyingPartyDetailedSummary_signInSuccessRateable)()
    SetSuccessfulSignInCount(value *int64)()
    SetTotalSignInCount(value *int64)()
    SetUniqueUserCount(value *int64)()
}
