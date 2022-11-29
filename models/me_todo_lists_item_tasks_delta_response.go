package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeTodoListsItemTasksDeltaResponse provides operations to call the delta method.
type MeTodoListsItemTasksDeltaResponse struct {
    BaseDeltaFunctionResponse
    // The value property
    value []TodoTaskable
}
// NewMeTodoListsItemTasksDeltaResponse instantiates a new MeTodoListsItemTasksDeltaResponse and sets the default values.
func NewMeTodoListsItemTasksDeltaResponse()(*MeTodoListsItemTasksDeltaResponse) {
    m := &MeTodoListsItemTasksDeltaResponse{
        BaseDeltaFunctionResponse: *NewBaseDeltaFunctionResponse(),
    }
    return m
}
// CreateMeTodoListsItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeTodoListsItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeTodoListsItemTasksDeltaResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeTodoListsItemTasksDeltaResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseDeltaFunctionResponse.GetFieldDeserializers()
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTodoTaskFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *MeTodoListsItemTasksDeltaResponse) GetValue()([]TodoTaskable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MeTodoListsItemTasksDeltaResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseDeltaFunctionResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MeTodoListsItemTasksDeltaResponse) SetValue(value []TodoTaskable)() {
    m.value = value
}
