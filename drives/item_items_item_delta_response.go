package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemDeltagetResponseable instead.
type ItemItemsItemDeltaResponse struct {
    ItemItemsItemDeltagetResponse
}
// NewItemItemsItemDeltaResponse instantiates a new ItemItemsItemDeltaResponse and sets the default values.
func NewItemItemsItemDeltaResponse()(*ItemItemsItemDeltaResponse) {
    m := &ItemItemsItemDeltaResponse{
        ItemItemsItemDeltagetResponse: *NewItemItemsItemDeltagetResponse(),
    }
    return m
}
// CreateItemItemsItemDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemDeltagetResponseable instead.
type ItemItemsItemDeltaResponseable interface {
    ItemItemsItemDeltagetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
