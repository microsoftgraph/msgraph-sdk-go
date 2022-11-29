package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemDrivesItemItemsItemCheckinPostRequestBody provides operations to call the checkin method.
type UsersItemDrivesItemItemsItemCheckinPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The checkInAs property
    checkInAs *string
    // The comment property
    comment *string
}
// NewUsersItemDrivesItemItemsItemCheckinPostRequestBody instantiates a new UsersItemDrivesItemItemsItemCheckinPostRequestBody and sets the default values.
func NewUsersItemDrivesItemItemsItemCheckinPostRequestBody()(*UsersItemDrivesItemItemsItemCheckinPostRequestBody) {
    m := &UsersItemDrivesItemItemsItemCheckinPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemDrivesItemItemsItemCheckinPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemDrivesItemItemsItemCheckinPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemDrivesItemItemsItemCheckinPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemDrivesItemItemsItemCheckinPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCheckInAs gets the checkInAs property value. The checkInAs property
func (m *UsersItemDrivesItemItemsItemCheckinPostRequestBody) GetCheckInAs()(*string) {
    return m.checkInAs
}
// GetComment gets the comment property value. The comment property
func (m *UsersItemDrivesItemItemsItemCheckinPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemDrivesItemItemsItemCheckinPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checkInAs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCheckInAs)
    res["comment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetComment)
    return res
}
// Serialize serializes information the current object
func (m *UsersItemDrivesItemItemsItemCheckinPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("checkInAs", m.GetCheckInAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comment", m.GetComment())
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
func (m *UsersItemDrivesItemItemsItemCheckinPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCheckInAs sets the checkInAs property value. The checkInAs property
func (m *UsersItemDrivesItemItemsItemCheckinPostRequestBody) SetCheckInAs(value *string)() {
    m.checkInAs = value
}
// SetComment sets the comment property value. The comment property
func (m *UsersItemDrivesItemItemsItemCheckinPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
