package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCheckmemberobjectsCheckMemberObjectsPostResponseable instead.
type ItemCheckmemberobjectsCheckMemberObjectsResponse struct {
    ItemCheckmemberobjectsCheckMemberObjectsPostResponse
}
// NewItemCheckmemberobjectsCheckMemberObjectsResponse instantiates a new ItemCheckmemberobjectsCheckMemberObjectsResponse and sets the default values.
func NewItemCheckmemberobjectsCheckMemberObjectsResponse()(*ItemCheckmemberobjectsCheckMemberObjectsResponse) {
    m := &ItemCheckmemberobjectsCheckMemberObjectsResponse{
        ItemCheckmemberobjectsCheckMemberObjectsPostResponse: *NewItemCheckmemberobjectsCheckMemberObjectsPostResponse(),
    }
    return m
}
// CreateItemCheckmemberobjectsCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCheckmemberobjectsCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCheckmemberobjectsCheckMemberObjectsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCheckmemberobjectsCheckMemberObjectsPostResponseable instead.
type ItemCheckmemberobjectsCheckMemberObjectsResponseable interface {
    ItemCheckmemberobjectsCheckMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
