package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemGetmailtipsGetMailTipsPostResponseable instead.
type ItemGetmailtipsGetMailTipsResponse struct {
    ItemGetmailtipsGetMailTipsPostResponse
}
// NewItemGetmailtipsGetMailTipsResponse instantiates a new ItemGetmailtipsGetMailTipsResponse and sets the default values.
func NewItemGetmailtipsGetMailTipsResponse()(*ItemGetmailtipsGetMailTipsResponse) {
    m := &ItemGetmailtipsGetMailTipsResponse{
        ItemGetmailtipsGetMailTipsPostResponse: *NewItemGetmailtipsGetMailTipsPostResponse(),
    }
    return m
}
// CreateItemGetmailtipsGetMailTipsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemGetmailtipsGetMailTipsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetmailtipsGetMailTipsResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemGetmailtipsGetMailTipsPostResponseable instead.
type ItemGetmailtipsGetMailTipsResponseable interface {
    ItemGetmailtipsGetMailTipsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
