package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemContactfoldersItemChildfoldersItemContactsDeltaGetResponseable instead.
type ItemContactfoldersItemChildfoldersItemContactsDeltaResponse struct {
    ItemContactfoldersItemChildfoldersItemContactsDeltaGetResponse
}
// NewItemContactfoldersItemChildfoldersItemContactsDeltaResponse instantiates a new ItemContactfoldersItemChildfoldersItemContactsDeltaResponse and sets the default values.
func NewItemContactfoldersItemChildfoldersItemContactsDeltaResponse()(*ItemContactfoldersItemChildfoldersItemContactsDeltaResponse) {
    m := &ItemContactfoldersItemChildfoldersItemContactsDeltaResponse{
        ItemContactfoldersItemChildfoldersItemContactsDeltaGetResponse: *NewItemContactfoldersItemChildfoldersItemContactsDeltaGetResponse(),
    }
    return m
}
// CreateItemContactfoldersItemChildfoldersItemContactsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemContactfoldersItemChildfoldersItemContactsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemContactfoldersItemChildfoldersItemContactsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemContactfoldersItemChildfoldersItemContactsDeltaGetResponseable instead.
type ItemContactfoldersItemChildfoldersItemContactsDeltaResponseable interface {
    ItemContactfoldersItemChildfoldersItemContactsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
