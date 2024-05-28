package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimarychannelMessagesDeltaGetResponseable instead.
type ItemPrimarychannelMessagesDeltaResponse struct {
    ItemPrimarychannelMessagesDeltaGetResponse
}
// NewItemPrimarychannelMessagesDeltaResponse instantiates a new ItemPrimarychannelMessagesDeltaResponse and sets the default values.
func NewItemPrimarychannelMessagesDeltaResponse()(*ItemPrimarychannelMessagesDeltaResponse) {
    m := &ItemPrimarychannelMessagesDeltaResponse{
        ItemPrimarychannelMessagesDeltaGetResponse: *NewItemPrimarychannelMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemPrimarychannelMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimarychannelMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimarychannelMessagesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimarychannelMessagesDeltaGetResponseable instead.
type ItemPrimarychannelMessagesDeltaResponseable interface {
    ItemPrimarychannelMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
