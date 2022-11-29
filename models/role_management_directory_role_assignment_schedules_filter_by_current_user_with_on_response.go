package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse provides operations to call the filterByCurrentUser method.
type RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []UnifiedRoleAssignmentScheduleable
}
// NewRoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse instantiates a new RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse and sets the default values.
func NewRoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse()(*RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse) {
    m := &RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateRoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse) GetValue()([]UnifiedRoleAssignmentScheduleable) {
    return m.value
}
// Serialize serializes information the current object
func (m *RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RoleManagementDirectoryRoleAssignmentSchedulesFilterByCurrentUserWithOnResponse) SetValue(value []UnifiedRoleAssignmentScheduleable)() {
    m.value = value
}
