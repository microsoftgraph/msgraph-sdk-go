package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponseable instead.
type ItemGrouplifecyclepoliciesItemAddgroupAddGroupResponse struct {
    ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponse
}
// NewItemGrouplifecyclepoliciesItemAddgroupAddGroupResponse instantiates a new ItemGrouplifecyclepoliciesItemAddgroupAddGroupResponse and sets the default values.
func NewItemGrouplifecyclepoliciesItemAddgroupAddGroupResponse()(*ItemGrouplifecyclepoliciesItemAddgroupAddGroupResponse) {
    m := &ItemGrouplifecyclepoliciesItemAddgroupAddGroupResponse{
        ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponse: *NewItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponse(),
    }
    return m
}
// CreateItemGrouplifecyclepoliciesItemAddgroupAddGroupResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGrouplifecyclepoliciesItemAddgroupAddGroupResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGrouplifecyclepoliciesItemAddgroupAddGroupResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponseable instead.
type ItemGrouplifecyclepoliciesItemAddgroupAddGroupResponseable interface {
    ItemGrouplifecyclepoliciesItemAddgroupAddGroupPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
