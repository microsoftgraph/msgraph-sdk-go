package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetMailTipsResponse 
// Deprecated: This class is obsolete. Use getMailTipsPostResponse instead.
type ItemGetMailTipsResponse struct {
    ItemGetMailTipsPostResponse
}
// NewItemGetMailTipsResponse instantiates a new ItemGetMailTipsResponse and sets the default values.
func NewItemGetMailTipsResponse()(*ItemGetMailTipsResponse) {
    m := &ItemGetMailTipsResponse{
        ItemGetMailTipsPostResponse: *NewItemGetMailTipsPostResponse(),
    }
    return m
}
// CreateItemGetMailTipsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetMailTipsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetMailTipsResponse(), nil
}
// ItemGetMailTipsResponseable 
// Deprecated: This class is obsolete. Use getMailTipsPostResponse instead.
type ItemGetMailTipsResponseable interface {
    ItemGetMailTipsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
