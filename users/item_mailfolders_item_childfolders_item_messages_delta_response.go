package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemMailfoldersItemChildfoldersItemMessagesDeltaGetResponseable instead.
type ItemMailfoldersItemChildfoldersItemMessagesDeltaResponse struct {
    ItemMailfoldersItemChildfoldersItemMessagesDeltaGetResponse
}
// NewItemMailfoldersItemChildfoldersItemMessagesDeltaResponse instantiates a new ItemMailfoldersItemChildfoldersItemMessagesDeltaResponse and sets the default values.
func NewItemMailfoldersItemChildfoldersItemMessagesDeltaResponse()(*ItemMailfoldersItemChildfoldersItemMessagesDeltaResponse) {
    m := &ItemMailfoldersItemChildfoldersItemMessagesDeltaResponse{
        ItemMailfoldersItemChildfoldersItemMessagesDeltaGetResponse: *NewItemMailfoldersItemChildfoldersItemMessagesDeltaGetResponse(),
    }
    return m
}
// CreateItemMailfoldersItemChildfoldersItemMessagesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemMailfoldersItemChildfoldersItemMessagesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailfoldersItemChildfoldersItemMessagesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemMailfoldersItemChildfoldersItemMessagesDeltaGetResponseable instead.
type ItemMailfoldersItemChildfoldersItemMessagesDeltaResponseable interface {
    ItemMailfoldersItemChildfoldersItemMessagesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
