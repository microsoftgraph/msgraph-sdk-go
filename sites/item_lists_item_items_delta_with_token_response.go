package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListsItemItemsDeltaWithTokengetResponseable instead.
type ItemListsItemItemsDeltaWithTokenResponse struct {
    ItemListsItemItemsDeltaWithTokengetResponse
}
// NewItemListsItemItemsDeltaWithTokenResponse instantiates a new ItemListsItemItemsDeltaWithTokenResponse and sets the default values.
func NewItemListsItemItemsDeltaWithTokenResponse()(*ItemListsItemItemsDeltaWithTokenResponse) {
    m := &ItemListsItemItemsDeltaWithTokenResponse{
        ItemListsItemItemsDeltaWithTokengetResponse: *NewItemListsItemItemsDeltaWithTokengetResponse(),
    }
    return m
}
// CreateItemListsItemItemsDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListsItemItemsDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemItemsDeltaWithTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListsItemItemsDeltaWithTokengetResponseable instead.
type ItemListsItemItemsDeltaWithTokenResponseable interface {
    ItemListsItemItemsDeltaWithTokengetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
