package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamPrimarychannelMessagesDeltaGetResponseable instead.
type ItemTeamPrimarychannelMessagesDeltaResponse struct {
    ItemTeamPrimarychannelMessagesDeltaGetResponse
}
// NewItemTeamPrimarychannelMessagesDeltaResponse instantiates a new ItemTeamPrimarychannelMessagesDeltaResponse and sets the default values.
func NewItemTeamPrimarychannelMessagesDeltaResponse()(*ItemTeamPrimarychannelMessagesDeltaResponse) {
    m := &ItemTeamPrimarychannelMessagesDeltaResponse{
        ItemTeamPrimarychannelMessagesDeltaGetResponse: *NewItemTeamPrimarychannelMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamPrimarychannelMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamPrimarychannelMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimarychannelMessagesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamPrimarychannelMessagesDeltaGetResponseable instead.
type ItemTeamPrimarychannelMessagesDeltaResponseable interface {
    ItemTeamPrimarychannelMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
