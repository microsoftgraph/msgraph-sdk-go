package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemMailfoldersItemChildfoldersDeltaGetResponseable instead.
type ItemMailfoldersItemChildfoldersDeltaResponse struct {
    ItemMailfoldersItemChildfoldersDeltaGetResponse
}
// NewItemMailfoldersItemChildfoldersDeltaResponse instantiates a new ItemMailfoldersItemChildfoldersDeltaResponse and sets the default values.
func NewItemMailfoldersItemChildfoldersDeltaResponse()(*ItemMailfoldersItemChildfoldersDeltaResponse) {
    m := &ItemMailfoldersItemChildfoldersDeltaResponse{
        ItemMailfoldersItemChildfoldersDeltaGetResponse: *NewItemMailfoldersItemChildfoldersDeltaGetResponse(),
    }
    return m
}
// CreateItemMailfoldersItemChildfoldersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemMailfoldersItemChildfoldersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailfoldersItemChildfoldersDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemMailfoldersItemChildfoldersDeltaGetResponseable instead.
type ItemMailfoldersItemChildfoldersDeltaResponseable interface {
    ItemMailfoldersItemChildfoldersDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
