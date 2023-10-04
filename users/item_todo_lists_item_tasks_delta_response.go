package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTodoListsItemTasksDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTodoListsItemTasksDeltaResponse struct {
    ItemTodoListsItemTasksDeltaGetResponse
}
// NewItemTodoListsItemTasksDeltaResponse instantiates a new ItemTodoListsItemTasksDeltaResponse and sets the default values.
func NewItemTodoListsItemTasksDeltaResponse()(*ItemTodoListsItemTasksDeltaResponse) {
    m := &ItemTodoListsItemTasksDeltaResponse{
        ItemTodoListsItemTasksDeltaGetResponse: *NewItemTodoListsItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateItemTodoListsItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTodoListsItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTodoListsItemTasksDeltaResponse(), nil
}
// ItemTodoListsItemTasksDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemTodoListsItemTasksDeltaResponseable interface {
    ItemTodoListsItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
