package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetMemberGroupsResponse 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemGetMemberGroupsResponse struct {
    ItemGetMemberGroupsPostResponse
}
// NewItemGetMemberGroupsResponse instantiates a new ItemGetMemberGroupsResponse and sets the default values.
func NewItemGetMemberGroupsResponse()(*ItemGetMemberGroupsResponse) {
    m := &ItemGetMemberGroupsResponse{
        ItemGetMemberGroupsPostResponse: *NewItemGetMemberGroupsPostResponse(),
    }
    return m
}
// CreateItemGetMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetMemberGroupsResponse(), nil
}
// ItemGetMemberGroupsResponseable 
// Deprecated: This class is obsolete. Use getMemberGroupsPostResponse instead.
type ItemGetMemberGroupsResponseable interface {
    ItemGetMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
