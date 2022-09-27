package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimulationAutomationRun provides operations to manage the collection of agreementAcceptance entities.
type SimulationAutomationRun struct {
    Entity
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The simulationId property
    simulationId *string
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *SimulationAutomationRunStatus
}
// NewSimulationAutomationRun instantiates a new simulationAutomationRun and sets the default values.
func NewSimulationAutomationRun()(*SimulationAutomationRun) {
    m := &SimulationAutomationRun{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.simulationAutomationRun";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSimulationAutomationRunFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimulationAutomationRunFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulationAutomationRun(), nil
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *SimulationAutomationRun) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationAutomationRun) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["endDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEndDateTime)
    res["simulationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSimulationId)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSimulationAutomationRunStatus , m.SetStatus)
    return res
}
// GetSimulationId gets the simulationId property value. The simulationId property
func (m *SimulationAutomationRun) GetSimulationId()(*string) {
    return m.simulationId
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *SimulationAutomationRun) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. The status property
func (m *SimulationAutomationRun) GetStatus()(*SimulationAutomationRunStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *SimulationAutomationRun) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("simulationId", m.GetSimulationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *SimulationAutomationRun) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetSimulationId sets the simulationId property value. The simulationId property
func (m *SimulationAutomationRun) SetSimulationId(value *string)() {
    m.simulationId = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *SimulationAutomationRun) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *SimulationAutomationRun) SetStatus(value *SimulationAutomationRunStatus)() {
    m.status = value
}
