package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponseable instead.
type DeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponse struct {
    DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponse
}
// NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponse instantiates a new DeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponse and sets the default values.
func NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponse()(*DeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponse) {
    m := &DeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponse{
        DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponse: *NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponse(),
    }
    return m
}
// CreateDeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponse(), nil
}
// Deprecated: This class is obsolete. Use DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponseable instead.
type DeleteditemsItemCheckmemberobjectsCheckMemberObjectsResponseable interface {
    DeleteditemsItemCheckmemberobjectsCheckMemberObjectsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
