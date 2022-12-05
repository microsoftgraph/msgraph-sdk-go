package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeAssignLicensePostRequestBody provides operations to call the assignLicense method.
type MeAssignLicensePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The addLicenses property
    addLicenses []AssignedLicenseable
    // The removeLicenses property
    removeLicenses []string
}
// NewMeAssignLicensePostRequestBody instantiates a new MeAssignLicensePostRequestBody and sets the default values.
func NewMeAssignLicensePostRequestBody()(*MeAssignLicensePostRequestBody) {
    m := &MeAssignLicensePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeAssignLicensePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeAssignLicensePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeAssignLicensePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeAssignLicensePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAddLicenses gets the addLicenses property value. The addLicenses property
func (m *MeAssignLicensePostRequestBody) GetAddLicenses()([]AssignedLicenseable) {
    return m.addLicenses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeAssignLicensePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLicenseable, len(val))
            for i, v := range val {
                res[i] = v.(AssignedLicenseable)
            }
            m.SetAddLicenses(res)
        }
        return nil
    }
    res["removeLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveLicenses(res)
        }
        return nil
    }
    return res
}
// GetRemoveLicenses gets the removeLicenses property value. The removeLicenses property
func (m *MeAssignLicensePostRequestBody) GetRemoveLicenses()([]string) {
    return m.removeLicenses
}
// Serialize serializes information the current object
func (m *MeAssignLicensePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteCollectionOfStringValues("removeLicenses", m.GetRemoveLicenses())
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
func (m *MeAssignLicensePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddLicenses sets the addLicenses property value. The addLicenses property
func (m *MeAssignLicensePostRequestBody) SetAddLicenses(value []AssignedLicenseable)() {
    m.addLicenses = value
}
// SetRemoveLicenses sets the removeLicenses property value. The removeLicenses property
func (m *MeAssignLicensePostRequestBody) SetRemoveLicenses(value []string)() {
    m.removeLicenses = value
}
