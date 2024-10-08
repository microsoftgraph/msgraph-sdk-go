package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceLogCollectionResponse windows Log Collection request entity.
type DeviceLogCollectionResponse struct {
    Entity
}
// DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB composed type wrapper for classes float64, ReferenceNumeric, string
type DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB instantiates a new DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB and sets the default values.
func NewDeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB()(*DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) {
    m := &DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateDeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewDeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB()
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
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) GetDouble()(*float64) {
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
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) GetReferenceNumeric()(*ReferenceNumeric) {
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
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) GetString()(*string) {
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
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBable interface {
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
// NewDeviceLogCollectionResponse instantiates a new DeviceLogCollectionResponse and sets the default values.
func NewDeviceLogCollectionResponse()(*DeviceLogCollectionResponse) {
    m := &DeviceLogCollectionResponse{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceLogCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceLogCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceLogCollectionResponse(), nil
}
// GetEnrolledByUser gets the enrolledByUser property value. The User Principal Name (UPN) of the user that enrolled the device.
// returns a *string when successful
func (m *DeviceLogCollectionResponse) GetEnrolledByUser()(*string) {
    val, err := m.GetBackingStore().Get("enrolledByUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExpirationDateTimeUTC gets the expirationDateTimeUTC property value. The DateTime of the expiration of the logs.
// returns a *Time when successful
func (m *DeviceLogCollectionResponse) GetExpirationDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTimeUTC")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceLogCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enrolledByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledByUser(val)
        }
        return nil
    }
    res["expirationDateTimeUTC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTimeUTC(val)
        }
        return nil
    }
    res["initiatedByUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedByUserPrincipalName(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["receivedDateTimeUTC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTimeUTC(val)
        }
        return nil
    }
    res["requestedDateTimeUTC"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedDateTimeUTC(val)
        }
        return nil
    }
    res["sizeInKB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInKB(val.(*DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKB))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLogUploadState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AppLogUploadState))
        }
        return nil
    }
    return res
}
// GetInitiatedByUserPrincipalName gets the initiatedByUserPrincipalName property value. The UPN for who initiated the request.
// returns a *string when successful
func (m *DeviceLogCollectionResponse) GetInitiatedByUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("initiatedByUserPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManagedDeviceId gets the managedDeviceId property value. Indicates Intune device unique identifier.
// returns a *UUID when successful
func (m *DeviceLogCollectionResponse) GetManagedDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("managedDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetReceivedDateTimeUTC gets the receivedDateTimeUTC property value. The DateTime the request was received.
// returns a *Time when successful
func (m *DeviceLogCollectionResponse) GetReceivedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("receivedDateTimeUTC")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRequestedDateTimeUTC gets the requestedDateTimeUTC property value. The DateTime of the request.
// returns a *Time when successful
func (m *DeviceLogCollectionResponse) GetRequestedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestedDateTimeUTC")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSizeInKB gets the sizeInKB property value. The size of the logs in KB. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// returns a DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBable when successful
func (m *DeviceLogCollectionResponse) GetSizeInKB()(DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBable) {
    val, err := m.GetBackingStore().Get("sizeInKB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBable)
    }
    return nil
}
// GetStatus gets the status property value. AppLogUploadStatus
// returns a *AppLogUploadState when successful
func (m *DeviceLogCollectionResponse) GetStatus()(*AppLogUploadState) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AppLogUploadState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceLogCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("enrolledByUser", m.GetEnrolledByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTimeUTC", m.GetExpirationDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initiatedByUserPrincipalName", m.GetInitiatedByUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTimeUTC", m.GetReceivedDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestedDateTimeUTC", m.GetRequestedDateTimeUTC())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sizeInKB", m.GetSizeInKB())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnrolledByUser sets the enrolledByUser property value. The User Principal Name (UPN) of the user that enrolled the device.
func (m *DeviceLogCollectionResponse) SetEnrolledByUser(value *string)() {
    err := m.GetBackingStore().Set("enrolledByUser", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTimeUTC sets the expirationDateTimeUTC property value. The DateTime of the expiration of the logs.
func (m *DeviceLogCollectionResponse) SetExpirationDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTimeUTC", value)
    if err != nil {
        panic(err)
    }
}
// SetInitiatedByUserPrincipalName sets the initiatedByUserPrincipalName property value. The UPN for who initiated the request.
func (m *DeviceLogCollectionResponse) SetInitiatedByUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("initiatedByUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Indicates Intune device unique identifier.
func (m *DeviceLogCollectionResponse) SetManagedDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("managedDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivedDateTimeUTC sets the receivedDateTimeUTC property value. The DateTime the request was received.
func (m *DeviceLogCollectionResponse) SetReceivedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("receivedDateTimeUTC", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedDateTimeUTC sets the requestedDateTimeUTC property value. The DateTime of the request.
func (m *DeviceLogCollectionResponse) SetRequestedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestedDateTimeUTC", value)
    if err != nil {
        panic(err)
    }
}
// SetSizeInKB sets the sizeInKB property value. The size of the logs in KB. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *DeviceLogCollectionResponse) SetSizeInKB(value DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBable)() {
    err := m.GetBackingStore().Set("sizeInKB", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. AppLogUploadStatus
func (m *DeviceLogCollectionResponse) SetStatus(value *AppLogUploadState)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
type DeviceLogCollectionResponseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnrolledByUser()(*string)
    GetExpirationDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInitiatedByUserPrincipalName()(*string)
    GetManagedDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetReceivedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSizeInKB()(DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBable)
    GetStatus()(*AppLogUploadState)
    SetEnrolledByUser(value *string)()
    SetExpirationDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInitiatedByUserPrincipalName(value *string)()
    SetManagedDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetReceivedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSizeInKB(value DeviceLogCollectionResponse_DeviceLogCollectionResponse_sizeInKBable)()
    SetStatus(value *AppLogUploadState)()
}
