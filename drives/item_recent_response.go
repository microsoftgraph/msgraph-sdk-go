package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemRecentgetResponseable instead.
type ItemRecentResponse struct {
    ItemRecentgetResponse
}
// NewItemRecentResponse instantiates a new ItemRecentResponse and sets the default values.
func NewItemRecentResponse()(*ItemRecentResponse) {
    m := &ItemRecentResponse{
        ItemRecentgetResponse: *NewItemRecentgetResponse(),
    }
    return m
}
// CreateItemRecentResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemRecentResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRecentResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemRecentgetResponseable instead.
type ItemRecentResponseable interface {
    ItemRecentgetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
