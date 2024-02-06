package billing

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RunningOperation 
type RunningOperation struct {
    Operation
}
// NewRunningOperation instantiates a new runningOperation and sets the default values.
func NewRunningOperation()(*RunningOperation) {
    m := &RunningOperation{
        Operation: *NewOperation(),
    }
    return m
}
// CreateRunningOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRunningOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunningOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RunningOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Operation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RunningOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Operation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// RunningOperationable 
type RunningOperationable interface {
    Operationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
