package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponseable instead.
type DeleteditemsItemCheckmembergroupsCheckMemberGroupsResponse struct {
    DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponse
}
// NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsResponse instantiates a new DeleteditemsItemCheckmembergroupsCheckMemberGroupsResponse and sets the default values.
func NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsResponse()(*DeleteditemsItemCheckmembergroupsCheckMemberGroupsResponse) {
    m := &DeleteditemsItemCheckmembergroupsCheckMemberGroupsResponse{
        DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponse: *NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponse(),
    }
    return m
}
// CreateDeleteditemsItemCheckmembergroupsCheckMemberGroupsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeleteditemsItemCheckmembergroupsCheckMemberGroupsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeleteditemsItemCheckmembergroupsCheckMemberGroupsResponse(), nil
}
// Deprecated: This class is obsolete. Use DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponseable instead.
type DeleteditemsItemCheckmembergroupsCheckMemberGroupsResponseable interface {
    DeleteditemsItemCheckmembergroupsCheckMemberGroupsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
