package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use VirtualEventsTownhallsGetByUserRoleWithRolegetResponseable instead.
type VirtualEventsTownhallsGetByUserRoleWithRoleResponse struct {
    VirtualEventsTownhallsGetByUserRoleWithRolegetResponse
}
// NewVirtualEventsTownhallsGetByUserRoleWithRoleResponse instantiates a new VirtualEventsTownhallsGetByUserRoleWithRoleResponse and sets the default values.
func NewVirtualEventsTownhallsGetByUserRoleWithRoleResponse()(*VirtualEventsTownhallsGetByUserRoleWithRoleResponse) {
    m := &VirtualEventsTownhallsGetByUserRoleWithRoleResponse{
        VirtualEventsTownhallsGetByUserRoleWithRolegetResponse: *NewVirtualEventsTownhallsGetByUserRoleWithRolegetResponse(),
    }
    return m
}
// CreateVirtualEventsTownhallsGetByUserRoleWithRoleResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventsTownhallsGetByUserRoleWithRoleResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventsTownhallsGetByUserRoleWithRoleResponse(), nil
}
// Deprecated: This class is obsolete. Use VirtualEventsTownhallsGetByUserRoleWithRolegetResponseable instead.
type VirtualEventsTownhallsGetByUserRoleWithRoleResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VirtualEventsTownhallsGetByUserRoleWithRolegetResponseable
}
