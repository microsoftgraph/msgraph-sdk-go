package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemJoinedTeamsGetAllMessagesResponse 
// Deprecated: This class is obsolete. Use getAllMessagesGetResponse instead.
type ItemJoinedTeamsGetAllMessagesResponse struct {
    ItemJoinedTeamsGetAllMessagesGetResponse
}
// NewItemJoinedTeamsGetAllMessagesResponse instantiates a new ItemJoinedTeamsGetAllMessagesResponse and sets the default values.
func NewItemJoinedTeamsGetAllMessagesResponse()(*ItemJoinedTeamsGetAllMessagesResponse) {
    m := &ItemJoinedTeamsGetAllMessagesResponse{
        ItemJoinedTeamsGetAllMessagesGetResponse: *NewItemJoinedTeamsGetAllMessagesGetResponse(),
    }
    return m
}
// CreateItemJoinedTeamsGetAllMessagesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemJoinedTeamsGetAllMessagesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsGetAllMessagesResponse(), nil
}
// ItemJoinedTeamsGetAllMessagesResponseable 
// Deprecated: This class is obsolete. Use getAllMessagesGetResponse instead.
type ItemJoinedTeamsGetAllMessagesResponseable interface {
    ItemJoinedTeamsGetAllMessagesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
