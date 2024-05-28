package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemSearchwithqSearchWithQGetResponseable instead.
type ItemItemsItemSearchwithqSearchWithQResponse struct {
    ItemItemsItemSearchwithqSearchWithQGetResponse
}
// NewItemItemsItemSearchwithqSearchWithQResponse instantiates a new ItemItemsItemSearchwithqSearchWithQResponse and sets the default values.
func NewItemItemsItemSearchwithqSearchWithQResponse()(*ItemItemsItemSearchwithqSearchWithQResponse) {
    m := &ItemItemsItemSearchwithqSearchWithQResponse{
        ItemItemsItemSearchwithqSearchWithQGetResponse: *NewItemItemsItemSearchwithqSearchWithQGetResponse(),
    }
    return m
}
// CreateItemItemsItemSearchwithqSearchWithQResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemSearchwithqSearchWithQResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemSearchwithqSearchWithQResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemSearchwithqSearchWithQGetResponseable instead.
type ItemItemsItemSearchwithqSearchWithQResponseable interface {
    ItemItemsItemSearchwithqSearchWithQGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
