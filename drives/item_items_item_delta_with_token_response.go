package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemDeltaWithTokenResponse 
// Deprecated: This class is obsolete. Use deltaWithTokenGetResponse instead.
type ItemItemsItemDeltaWithTokenResponse struct {
    ItemItemsItemDeltaWithTokenGetResponse
}
// NewItemItemsItemDeltaWithTokenResponse instantiates a new ItemItemsItemDeltaWithTokenResponse and sets the default values.
func NewItemItemsItemDeltaWithTokenResponse()(*ItemItemsItemDeltaWithTokenResponse) {
    m := &ItemItemsItemDeltaWithTokenResponse{
        ItemItemsItemDeltaWithTokenGetResponse: *NewItemItemsItemDeltaWithTokenGetResponse(),
    }
    return m
}
// CreateItemItemsItemDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemDeltaWithTokenResponse(), nil
}
// ItemItemsItemDeltaWithTokenResponseable 
// Deprecated: This class is obsolete. Use deltaWithTokenGetResponse instead.
type ItemItemsItemDeltaWithTokenResponseable interface {
    ItemItemsItemDeltaWithTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
