package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetMemberObjectsResponse 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemGetMemberObjectsResponse struct {
    ItemGetMemberObjectsPostResponse
}
// NewItemGetMemberObjectsResponse instantiates a new ItemGetMemberObjectsResponse and sets the default values.
func NewItemGetMemberObjectsResponse()(*ItemGetMemberObjectsResponse) {
    m := &ItemGetMemberObjectsResponse{
        ItemGetMemberObjectsPostResponse: *NewItemGetMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemGetMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetMemberObjectsResponse(), nil
}
// ItemGetMemberObjectsResponseable 
// Deprecated: This class is obsolete. Use getMemberObjectsPostResponse instead.
type ItemGetMemberObjectsResponseable interface {
    ItemGetMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
