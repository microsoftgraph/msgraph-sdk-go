package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemRevokesigninsessionsRevokeSignInSessionsPostResponseable instead.
type ItemRevokesigninsessionsRevokeSignInSessionsResponse struct {
    ItemRevokesigninsessionsRevokeSignInSessionsPostResponse
}
// NewItemRevokesigninsessionsRevokeSignInSessionsResponse instantiates a new ItemRevokesigninsessionsRevokeSignInSessionsResponse and sets the default values.
func NewItemRevokesigninsessionsRevokeSignInSessionsResponse()(*ItemRevokesigninsessionsRevokeSignInSessionsResponse) {
    m := &ItemRevokesigninsessionsRevokeSignInSessionsResponse{
        ItemRevokesigninsessionsRevokeSignInSessionsPostResponse: *NewItemRevokesigninsessionsRevokeSignInSessionsPostResponse(),
    }
    return m
}
// CreateItemRevokesigninsessionsRevokeSignInSessionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemRevokesigninsessionsRevokeSignInSessionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRevokesigninsessionsRevokeSignInSessionsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemRevokesigninsessionsRevokeSignInSessionsPostResponseable instead.
type ItemRevokesigninsessionsRevokeSignInSessionsResponseable interface {
    ItemRevokesigninsessionsRevokeSignInSessionsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
