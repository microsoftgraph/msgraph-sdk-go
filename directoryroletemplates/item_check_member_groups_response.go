package directoryroletemplates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCheckMemberGroupspostResponseable instead.
type ItemCheckMemberGroupsResponse struct {
    ItemCheckMemberGroupspostResponse
}
// NewItemCheckMemberGroupsResponse instantiates a new ItemCheckMemberGroupsResponse and sets the default values.
func NewItemCheckMemberGroupsResponse()(*ItemCheckMemberGroupsResponse) {
    m := &ItemCheckMemberGroupsResponse{
        ItemCheckMemberGroupspostResponse: *NewItemCheckMemberGroupspostResponse(),
    }
    return m
}
// CreateItemCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCheckMemberGroupsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCheckMemberGroupspostResponseable instead.
type ItemCheckMemberGroupsResponseable interface {
    ItemCheckMemberGroupspostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
