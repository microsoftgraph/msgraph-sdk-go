package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemChatsItemPermissionGrantsGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemChatsItemPermissionGrantsGetByIdsResponse struct {
    ItemChatsItemPermissionGrantsGetByIdsPostResponse
}
// NewItemChatsItemPermissionGrantsGetByIdsResponse instantiates a new ItemChatsItemPermissionGrantsGetByIdsResponse and sets the default values.
func NewItemChatsItemPermissionGrantsGetByIdsResponse()(*ItemChatsItemPermissionGrantsGetByIdsResponse) {
    m := &ItemChatsItemPermissionGrantsGetByIdsResponse{
        ItemChatsItemPermissionGrantsGetByIdsPostResponse: *NewItemChatsItemPermissionGrantsGetByIdsPostResponse(),
    }
    return m
}
// CreateItemChatsItemPermissionGrantsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemChatsItemPermissionGrantsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemChatsItemPermissionGrantsGetByIdsResponse(), nil
}
// ItemChatsItemPermissionGrantsGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemChatsItemPermissionGrantsGetByIdsResponseable interface {
    ItemChatsItemPermissionGrantsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
