package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPermissionGrantsGetByIdsResponse 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemPermissionGrantsGetByIdsResponse struct {
    ItemPermissionGrantsGetByIdsPostResponse
}
// NewItemPermissionGrantsGetByIdsResponse instantiates a new ItemPermissionGrantsGetByIdsResponse and sets the default values.
func NewItemPermissionGrantsGetByIdsResponse()(*ItemPermissionGrantsGetByIdsResponse) {
    m := &ItemPermissionGrantsGetByIdsResponse{
        ItemPermissionGrantsGetByIdsPostResponse: *NewItemPermissionGrantsGetByIdsPostResponse(),
    }
    return m
}
// CreateItemPermissionGrantsGetByIdsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPermissionGrantsGetByIdsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPermissionGrantsGetByIdsResponse(), nil
}
// ItemPermissionGrantsGetByIdsResponseable 
// Deprecated: This class is obsolete. Use getByIdsPostResponse instead.
type ItemPermissionGrantsGetByIdsResponseable interface {
    ItemPermissionGrantsGetByIdsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
