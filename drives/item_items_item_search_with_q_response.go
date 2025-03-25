package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemItemsItemSearchWithQgetResponseable instead.
type ItemItemsItemSearchWithQResponse struct {
    ItemItemsItemSearchWithQgetResponse
}
// NewItemItemsItemSearchWithQResponse instantiates a new ItemItemsItemSearchWithQResponse and sets the default values.
func NewItemItemsItemSearchWithQResponse()(*ItemItemsItemSearchWithQResponse) {
    m := &ItemItemsItemSearchWithQResponse{
        ItemItemsItemSearchWithQgetResponse: *NewItemItemsItemSearchWithQgetResponse(),
    }
    return m
}
// CreateItemItemsItemSearchWithQResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemSearchWithQResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemSearchWithQResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemItemsItemSearchWithQgetResponseable instead.
type ItemItemsItemSearchWithQResponseable interface {
    ItemItemsItemSearchWithQgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
