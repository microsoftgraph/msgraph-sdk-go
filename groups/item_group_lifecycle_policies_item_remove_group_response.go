package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGroupLifecyclePoliciesItemRemoveGroupResponse 
// Deprecated: This class is obsolete. Use removeGroupPostResponse instead.
type ItemGroupLifecyclePoliciesItemRemoveGroupResponse struct {
    ItemGroupLifecyclePoliciesItemRemoveGroupPostResponse
}
// NewItemGroupLifecyclePoliciesItemRemoveGroupResponse instantiates a new ItemGroupLifecyclePoliciesItemRemoveGroupResponse and sets the default values.
func NewItemGroupLifecyclePoliciesItemRemoveGroupResponse()(*ItemGroupLifecyclePoliciesItemRemoveGroupResponse) {
    m := &ItemGroupLifecyclePoliciesItemRemoveGroupResponse{
        ItemGroupLifecyclePoliciesItemRemoveGroupPostResponse: *NewItemGroupLifecyclePoliciesItemRemoveGroupPostResponse(),
    }
    return m
}
// CreateItemGroupLifecyclePoliciesItemRemoveGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGroupLifecyclePoliciesItemRemoveGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGroupLifecyclePoliciesItemRemoveGroupResponse(), nil
}
// ItemGroupLifecyclePoliciesItemRemoveGroupResponseable 
// Deprecated: This class is obsolete. Use removeGroupPostResponse instead.
type ItemGroupLifecyclePoliciesItemRemoveGroupResponseable interface {
    ItemGroupLifecyclePoliciesItemRemoveGroupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
