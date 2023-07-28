package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnDemandExecutionOnly 
type OnDemandExecutionOnly struct {
    WorkflowExecutionConditions
}
// NewOnDemandExecutionOnly instantiates a new onDemandExecutionOnly and sets the default values.
func NewOnDemandExecutionOnly()(*OnDemandExecutionOnly) {
    m := &OnDemandExecutionOnly{
        WorkflowExecutionConditions: *NewWorkflowExecutionConditions(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.onDemandExecutionOnly"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnDemandExecutionOnlyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnDemandExecutionOnlyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnDemandExecutionOnly(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnDemandExecutionOnly) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowExecutionConditions.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OnDemandExecutionOnly) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowExecutionConditions.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// OnDemandExecutionOnlyable 
type OnDemandExecutionOnlyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WorkflowExecutionConditionsable
}
