package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamChannelsItemMessagesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamChannelsItemMessagesDeltaResponse struct {
    ItemTeamChannelsItemMessagesDeltaGetResponse
}
// NewItemTeamChannelsItemMessagesDeltaResponse instantiates a new ItemTeamChannelsItemMessagesDeltaResponse and sets the default values.
func NewItemTeamChannelsItemMessagesDeltaResponse()(*ItemTeamChannelsItemMessagesDeltaResponse) {
    m := &ItemTeamChannelsItemMessagesDeltaResponse{
        ItemTeamChannelsItemMessagesDeltaGetResponse: *NewItemTeamChannelsItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamChannelsItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemMessagesDeltaResponse(), nil
}
// ItemTeamChannelsItemMessagesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTeamChannelsItemMessagesDeltaResponseable interface {
    ItemTeamChannelsItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
