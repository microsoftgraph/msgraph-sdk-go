package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemDrivesItemItemsItemDeltaWithTokenResponse provides operations to call the delta method.
type UsersItemDrivesItemItemsItemDeltaWithTokenResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []DriveItemable
}
// NewUsersItemDrivesItemItemsItemDeltaWithTokenResponse instantiates a new UsersItemDrivesItemItemsItemDeltaWithTokenResponse and sets the default values.
func NewUsersItemDrivesItemItemsItemDeltaWithTokenResponse()(*UsersItemDrivesItemItemsItemDeltaWithTokenResponse) {
    m := &UsersItemDrivesItemItemsItemDeltaWithTokenResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateUsersItemDrivesItemItemsItemDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemDrivesItemItemsItemDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemDrivesItemItemsItemDeltaWithTokenResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemDrivesItemItemsItemDeltaWithTokenResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemDrivesItemItemsItemDeltaWithTokenResponse) GetValue()([]DriveItemable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemDrivesItemItemsItemDeltaWithTokenResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UsersItemDrivesItemItemsItemDeltaWithTokenResponse) SetValue(value []DriveItemable)() {
    m.value = value
}
