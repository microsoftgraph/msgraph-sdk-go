package grouplifecyclepolicies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemRemoveGroupResponse 
// Deprecated: This class is obsolete. Use removeGroupPostResponse instead.
type ItemRemoveGroupResponse struct {
    ItemRemoveGroupPostResponse
}
// NewItemRemoveGroupResponse instantiates a new ItemRemoveGroupResponse and sets the default values.
func NewItemRemoveGroupResponse()(*ItemRemoveGroupResponse) {
    m := &ItemRemoveGroupResponse{
        ItemRemoveGroupPostResponse: *NewItemRemoveGroupPostResponse(),
    }
    return m
}
// CreateItemRemoveGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemRemoveGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRemoveGroupResponse(), nil
}
// ItemRemoveGroupResponseable 
// Deprecated: This class is obsolete. Use removeGroupPostResponse instead.
type ItemRemoveGroupResponseable interface {
    ItemRemoveGroupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
