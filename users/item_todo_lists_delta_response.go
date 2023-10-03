package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTodoListsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTodoListsDeltaResponse struct {
    ItemTodoListsDeltaGetResponse
}
// NewItemTodoListsDeltaResponse instantiates a new ItemTodoListsDeltaResponse and sets the default values.
func NewItemTodoListsDeltaResponse()(*ItemTodoListsDeltaResponse) {
    m := &ItemTodoListsDeltaResponse{
        ItemTodoListsDeltaGetResponse: *NewItemTodoListsDeltaGetResponse(),
    }
    return m
}
// CreateItemTodoListsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTodoListsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTodoListsDeltaResponse(), nil
}
// ItemTodoListsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTodoListsDeltaResponseable interface {
    ItemTodoListsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
