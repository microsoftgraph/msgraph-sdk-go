package shares

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListItemsDeltaWithTokengetResponseable instead.
type ItemListItemsDeltaWithTokenResponse struct {
    ItemListItemsDeltaWithTokengetResponse
}
// NewItemListItemsDeltaWithTokenResponse instantiates a new ItemListItemsDeltaWithTokenResponse and sets the default values.
func NewItemListItemsDeltaWithTokenResponse()(*ItemListItemsDeltaWithTokenResponse) {
    m := &ItemListItemsDeltaWithTokenResponse{
        ItemListItemsDeltaWithTokengetResponse: *NewItemListItemsDeltaWithTokengetResponse(),
    }
    return m
}
// CreateItemListItemsDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListItemsDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListItemsDeltaWithTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListItemsDeltaWithTokengetResponseable instead.
type ItemListItemsDeltaWithTokenResponseable interface {
    ItemListItemsDeltaWithTokengetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
