package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemFollowedsitesRemovePostResponseable instead.
type ItemFollowedsitesRemoveResponse struct {
    ItemFollowedsitesRemovePostResponse
}
// NewItemFollowedsitesRemoveResponse instantiates a new ItemFollowedsitesRemoveResponse and sets the default values.
func NewItemFollowedsitesRemoveResponse()(*ItemFollowedsitesRemoveResponse) {
    m := &ItemFollowedsitesRemoveResponse{
        ItemFollowedsitesRemovePostResponse: *NewItemFollowedsitesRemovePostResponse(),
    }
    return m
}
// CreateItemFollowedsitesRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemFollowedsitesRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemFollowedsitesRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemFollowedsitesRemovePostResponseable instead.
type ItemFollowedsitesRemoveResponseable interface {
    ItemFollowedsitesRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
