package domains

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPromoteResponse 
// Deprecated: This class is obsolete. Use promotePostResponse instead.
type ItemPromoteResponse struct {
    ItemPromotePostResponse
}
// NewItemPromoteResponse instantiates a new ItemPromoteResponse and sets the default values.
func NewItemPromoteResponse()(*ItemPromoteResponse) {
    m := &ItemPromoteResponse{
        ItemPromotePostResponse: *NewItemPromotePostResponse(),
    }
    return m
}
// CreateItemPromoteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPromoteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPromoteResponse(), nil
}
// ItemPromoteResponseable 
// Deprecated: This class is obsolete. Use promotePostResponse instead.
type ItemPromoteResponseable interface {
    ItemPromotePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
