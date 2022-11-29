package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse provides operations to call the delta method.
type UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []Contactable
}
// NewUsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse instantiates a new UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse and sets the default values.
func NewUsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse()(*UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse) {
    m := &UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateUsersItemContactFoldersItemChildFoldersItemContactsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemContactFoldersItemChildFoldersItemContactsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateContactFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse) GetValue()([]Contactable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UsersItemContactFoldersItemChildFoldersItemContactsDeltaResponse) SetValue(value []Contactable)() {
    m.value = value
}
