package me

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// AssignLicensePostRequestBody 
type AssignLicensePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The addLicenses property
    addLicenses []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable
    // The removeLicenses property
    removeLicenses []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewAssignLicensePostRequestBody instantiates a new AssignLicensePostRequestBody and sets the default values.
func NewAssignLicensePostRequestBody()(*AssignLicensePostRequestBody) {
    m := &AssignLicensePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignLicensePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignLicensePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignLicensePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignLicensePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAddLicenses gets the addLicenses property value. The addLicenses property
func (m *AssignLicensePostRequestBody) GetAddLicenses()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable) {
    return m.addLicenses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignLicensePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable)
            }
            m.SetAddLicenses(res)
        }
        return nil
    }
    res["removeLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID, len(val))
            for i, v := range val {
                res[i] = *(v.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID))
            }
            m.SetRemoveLicenses(res)
        }
        return nil
    }
    return res
}
// GetRemoveLicenses gets the removeLicenses property value. The removeLicenses property
func (m *AssignLicensePostRequestBody) GetRemoveLicenses()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.removeLicenses
}
// Serialize serializes information the current object
func (m *AssignLicensePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddLicenses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddLicenses()))
        for i, v := range m.GetAddLicenses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("addLicenses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoveLicenses() != nil {
        err := writer.WriteCollectionOfUUIDValues("removeLicenses", m.GetRemoveLicenses())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignLicensePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddLicenses sets the addLicenses property value. The addLicenses property
func (m *AssignLicensePostRequestBody) SetAddLicenses(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable)() {
    m.addLicenses = value
}
// SetRemoveLicenses sets the removeLicenses property value. The removeLicenses property
func (m *AssignLicensePostRequestBody) SetRemoveLicenses(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.removeLicenses = value
}
