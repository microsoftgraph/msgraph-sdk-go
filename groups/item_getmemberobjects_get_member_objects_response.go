package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemGetmemberobjectsGetMemberObjectsPostResponseable instead.
type ItemGetmemberobjectsGetMemberObjectsResponse struct {
    ItemGetmemberobjectsGetMemberObjectsPostResponse
}
// NewItemGetmemberobjectsGetMemberObjectsResponse instantiates a new ItemGetmemberobjectsGetMemberObjectsResponse and sets the default values.
func NewItemGetmemberobjectsGetMemberObjectsResponse()(*ItemGetmemberobjectsGetMemberObjectsResponse) {
    m := &ItemGetmemberobjectsGetMemberObjectsResponse{
        ItemGetmemberobjectsGetMemberObjectsPostResponse: *NewItemGetmemberobjectsGetMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemGetmemberobjectsGetMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGetmemberobjectsGetMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetmemberobjectsGetMemberObjectsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemGetmemberobjectsGetMemberObjectsPostResponseable instead.
type ItemGetmemberobjectsGetMemberObjectsResponseable interface {
    ItemGetmemberobjectsGetMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
