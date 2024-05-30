package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable instead.
type ItemListsItemItemsDeltawithtokenDeltaWithTokenResponse struct {
    ItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponse
}
// NewItemListsItemItemsDeltawithtokenDeltaWithTokenResponse instantiates a new ItemListsItemItemsDeltawithtokenDeltaWithTokenResponse and sets the default values.
func NewItemListsItemItemsDeltawithtokenDeltaWithTokenResponse()(*ItemListsItemItemsDeltawithtokenDeltaWithTokenResponse) {
    m := &ItemListsItemItemsDeltawithtokenDeltaWithTokenResponse{
        ItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponse: *NewItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponse(),
    }
    return m
}
// CreateItemListsItemItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemListsItemItemsDeltawithtokenDeltaWithTokenResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemItemsDeltawithtokenDeltaWithTokenResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable instead.
type ItemListsItemItemsDeltawithtokenDeltaWithTokenResponseable interface {
    ItemListsItemItemsDeltawithtokenDeltaWithTokenGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
