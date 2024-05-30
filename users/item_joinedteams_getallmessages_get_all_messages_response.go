package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedteamsGetallmessagesGetAllMessagesGetResponseable instead.
type ItemJoinedteamsGetallmessagesGetAllMessagesResponse struct {
    ItemJoinedteamsGetallmessagesGetAllMessagesGetResponse
}
// NewItemJoinedteamsGetallmessagesGetAllMessagesResponse instantiates a new ItemJoinedteamsGetallmessagesGetAllMessagesResponse and sets the default values.
func NewItemJoinedteamsGetallmessagesGetAllMessagesResponse()(*ItemJoinedteamsGetallmessagesGetAllMessagesResponse) {
    m := &ItemJoinedteamsGetallmessagesGetAllMessagesResponse{
        ItemJoinedteamsGetallmessagesGetAllMessagesGetResponse: *NewItemJoinedteamsGetallmessagesGetAllMessagesGetResponse(),
    }
    return m
}
// CreateItemJoinedteamsGetallmessagesGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedteamsGetallmessagesGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedteamsGetallmessagesGetAllMessagesResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedteamsGetallmessagesGetAllMessagesGetResponseable instead.
type ItemJoinedteamsGetallmessagesGetAllMessagesResponseable interface {
    ItemJoinedteamsGetallmessagesGetAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
