package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemFollowedsitesAddPostResponseable instead.
type ItemFollowedsitesAddResponse struct {
    ItemFollowedsitesAddPostResponse
}
// NewItemFollowedsitesAddResponse instantiates a new ItemFollowedsitesAddResponse and sets the default values.
func NewItemFollowedsitesAddResponse()(*ItemFollowedsitesAddResponse) {
    m := &ItemFollowedsitesAddResponse{
        ItemFollowedsitesAddPostResponse: *NewItemFollowedsitesAddPostResponse(),
    }
    return m
}
// CreateItemFollowedsitesAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemFollowedsitesAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemFollowedsitesAddResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemFollowedsitesAddPostResponseable instead.
type ItemFollowedsitesAddResponseable interface {
    ItemFollowedsitesAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
