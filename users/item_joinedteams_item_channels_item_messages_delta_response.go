package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedteamsItemChannelsItemMessagesDeltaGetResponseable instead.
type ItemJoinedteamsItemChannelsItemMessagesDeltaResponse struct {
    ItemJoinedteamsItemChannelsItemMessagesDeltaGetResponse
}
// NewItemJoinedteamsItemChannelsItemMessagesDeltaResponse instantiates a new ItemJoinedteamsItemChannelsItemMessagesDeltaResponse and sets the default values.
func NewItemJoinedteamsItemChannelsItemMessagesDeltaResponse()(*ItemJoinedteamsItemChannelsItemMessagesDeltaResponse) {
    m := &ItemJoinedteamsItemChannelsItemMessagesDeltaResponse{
        ItemJoinedteamsItemChannelsItemMessagesDeltaGetResponse: *NewItemJoinedteamsItemChannelsItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemJoinedteamsItemChannelsItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedteamsItemChannelsItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedteamsItemChannelsItemMessagesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedteamsItemChannelsItemMessagesDeltaGetResponseable instead.
type ItemJoinedteamsItemChannelsItemMessagesDeltaResponseable interface {
    ItemJoinedteamsItemChannelsItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
