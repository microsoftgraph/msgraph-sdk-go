package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemCheckMemberObjectspostResponseable instead.
type ItemCheckMemberObjectsResponse struct {
    ItemCheckMemberObjectspostResponse
}
// NewItemCheckMemberObjectsResponse instantiates a new ItemCheckMemberObjectsResponse and sets the default values.
func NewItemCheckMemberObjectsResponse()(*ItemCheckMemberObjectsResponse) {
    m := &ItemCheckMemberObjectsResponse{
        ItemCheckMemberObjectspostResponse: *NewItemCheckMemberObjectspostResponse(),
    }
    return m
}
// CreateItemCheckMemberObjectsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCheckMemberObjectsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCheckMemberObjectsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemCheckMemberObjectspostResponseable instead.
type ItemCheckMemberObjectsResponseable interface {
    ItemCheckMemberObjectspostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
