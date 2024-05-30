package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedteamsItemChannelsItemMessagesItemRepliesDeltaGetResponseable instead.
type DeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponse struct {
    DeletedteamsItemChannelsItemMessagesItemRepliesDeltaGetResponse
}
// NewDeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponse instantiates a new DeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponse and sets the default values.
func NewDeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponse()(*DeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponse) {
    m := &DeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponse{
        DeletedteamsItemChannelsItemMessagesItemRepliesDeltaGetResponse: *NewDeletedteamsItemChannelsItemMessagesItemRepliesDeltaGetResponse(),
    }
    return m
}
// CreateDeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedteamsItemChannelsItemMessagesItemRepliesDeltaGetResponseable instead.
type DeletedteamsItemChannelsItemMessagesItemRepliesDeltaResponseable interface {
    DeletedteamsItemChannelsItemMessagesItemRepliesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
