package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemSitesRemovePostResponseable instead.
type ItemSitesItemSitesRemoveResponse struct {
    ItemSitesItemSitesRemovePostResponse
}
// NewItemSitesItemSitesRemoveResponse instantiates a new ItemSitesItemSitesRemoveResponse and sets the default values.
func NewItemSitesItemSitesRemoveResponse()(*ItemSitesItemSitesRemoveResponse) {
    m := &ItemSitesItemSitesRemoveResponse{
        ItemSitesItemSitesRemovePostResponse: *NewItemSitesItemSitesRemovePostResponse(),
    }
    return m
}
// CreateItemSitesItemSitesRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemSitesRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemSitesRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemSitesRemovePostResponseable instead.
type ItemSitesItemSitesRemoveResponseable interface {
    ItemSitesItemSitesRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
