package grouplifecyclepolicies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemAddGroupResponse 
// Deprecated: This class is obsolete. Use addGroupPostResponse instead.
type ItemAddGroupResponse struct {
    ItemAddGroupPostResponse
}
// NewItemAddGroupResponse instantiates a new ItemAddGroupResponse and sets the default values.
func NewItemAddGroupResponse()(*ItemAddGroupResponse) {
    m := &ItemAddGroupResponse{
        ItemAddGroupPostResponse: *NewItemAddGroupPostResponse(),
    }
    return m
}
// CreateItemAddGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAddGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAddGroupResponse(), nil
}
// ItemAddGroupResponseable 
// Deprecated: This class is obsolete. Use addGroupPostResponse instead.
type ItemAddGroupResponseable interface {
    ItemAddGroupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
