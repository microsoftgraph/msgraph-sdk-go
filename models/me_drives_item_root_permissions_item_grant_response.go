package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeDrivesItemRootPermissionsItemGrantResponse provides operations to call the grant method.
type MeDrivesItemRootPermissionsItemGrantResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Permissionable
}
// NewMeDrivesItemRootPermissionsItemGrantResponse instantiates a new MeDrivesItemRootPermissionsItemGrantResponse and sets the default values.
func NewMeDrivesItemRootPermissionsItemGrantResponse()(*MeDrivesItemRootPermissionsItemGrantResponse) {
    m := &MeDrivesItemRootPermissionsItemGrantResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMeDrivesItemRootPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeDrivesItemRootPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeDrivesItemRootPermissionsItemGrantResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeDrivesItemRootPermissionsItemGrantResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePermissionFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *MeDrivesItemRootPermissionsItemGrantResponse) GetValue()([]Permissionable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeDrivesItemRootPermissionsItemGrantResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
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
func (m *MeDrivesItemRootPermissionsItemGrantResponse) SetValue(value []Permissionable)() {
    m.value = value
}
