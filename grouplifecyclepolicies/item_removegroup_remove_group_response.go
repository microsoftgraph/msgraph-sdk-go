package grouplifecyclepolicies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemRemovegroupRemoveGroupPostResponseable instead.
type ItemRemovegroupRemoveGroupResponse struct {
    ItemRemovegroupRemoveGroupPostResponse
}
// NewItemRemovegroupRemoveGroupResponse instantiates a new ItemRemovegroupRemoveGroupResponse and sets the default values.
func NewItemRemovegroupRemoveGroupResponse()(*ItemRemovegroupRemoveGroupResponse) {
    m := &ItemRemovegroupRemoveGroupResponse{
        ItemRemovegroupRemoveGroupPostResponse: *NewItemRemovegroupRemoveGroupPostResponse(),
    }
    return m
}
// CreateItemRemovegroupRemoveGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemRemovegroupRemoveGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRemovegroupRemoveGroupResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemRemovegroupRemoveGroupPostResponseable instead.
type ItemRemovegroupRemoveGroupResponseable interface {
    ItemRemovegroupRemoveGroupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
