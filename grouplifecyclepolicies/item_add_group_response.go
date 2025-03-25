package grouplifecyclepolicies

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemAddGrouppostResponseable instead.
type ItemAddGroupResponse struct {
    ItemAddGrouppostResponse
}
// NewItemAddGroupResponse instantiates a new ItemAddGroupResponse and sets the default values.
func NewItemAddGroupResponse()(*ItemAddGroupResponse) {
    m := &ItemAddGroupResponse{
        ItemAddGrouppostResponse: *NewItemAddGrouppostResponse(),
    }
    return m
}
// CreateItemAddGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAddGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAddGroupResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemAddGrouppostResponseable instead.
type ItemAddGroupResponseable interface {
    ItemAddGrouppostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
