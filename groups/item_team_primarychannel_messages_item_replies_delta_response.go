package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponseable instead.
type ItemTeamPrimarychannelMessagesItemRepliesDeltaResponse struct {
    ItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponse
}
// NewItemTeamPrimarychannelMessagesItemRepliesDeltaResponse instantiates a new ItemTeamPrimarychannelMessagesItemRepliesDeltaResponse and sets the default values.
func NewItemTeamPrimarychannelMessagesItemRepliesDeltaResponse()(*ItemTeamPrimarychannelMessagesItemRepliesDeltaResponse) {
    m := &ItemTeamPrimarychannelMessagesItemRepliesDeltaResponse{
        ItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponse: *NewItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponse(),
    }
    return m
}
// CreateItemTeamPrimarychannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamPrimarychannelMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamPrimarychannelMessagesItemRepliesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponseable instead.
type ItemTeamPrimarychannelMessagesItemRepliesDeltaResponseable interface {
    ItemTeamPrimarychannelMessagesItemRepliesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
