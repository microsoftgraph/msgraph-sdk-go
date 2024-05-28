package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemMailfoldersItemMessagesDeltaGetResponseable instead.
type ItemMailfoldersItemMessagesDeltaResponse struct {
    ItemMailfoldersItemMessagesDeltaGetResponse
}
// NewItemMailfoldersItemMessagesDeltaResponse instantiates a new ItemMailfoldersItemMessagesDeltaResponse and sets the default values.
func NewItemMailfoldersItemMessagesDeltaResponse()(*ItemMailfoldersItemMessagesDeltaResponse) {
    m := &ItemMailfoldersItemMessagesDeltaResponse{
        ItemMailfoldersItemMessagesDeltaGetResponse: *NewItemMailfoldersItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemMailfoldersItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemMailfoldersItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailfoldersItemMessagesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemMailfoldersItemMessagesDeltaGetResponseable instead.
type ItemMailfoldersItemMessagesDeltaResponseable interface {
    ItemMailfoldersItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
