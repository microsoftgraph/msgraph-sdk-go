package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSimulationDetails 
type UserSimulationDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignedTrainingsCount property
    assignedTrainingsCount *int32
    // The completedTrainingsCount property
    completedTrainingsCount *int32
    // The compromisedDateTime property
    compromisedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The inProgressTrainingsCount property
    inProgressTrainingsCount *int32
    // The isCompromised property
    isCompromised *bool
    // The OdataType property
    odataType *string
    // The reportedPhishDateTime property
    reportedPhishDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The simulationEvents property
    simulationEvents []UserSimulationEventInfoable
    // The simulationUser property
    simulationUser AttackSimulationUserable
    // The trainingEvents property
    trainingEvents []UserTrainingEventInfoable
}
// NewUserSimulationDetails instantiates a new userSimulationDetails and sets the default values.
func NewUserSimulationDetails()(*UserSimulationDetails) {
    m := &UserSimulationDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.userSimulationDetails";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserSimulationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSimulationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSimulationDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSimulationDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignedTrainingsCount gets the assignedTrainingsCount property value. The assignedTrainingsCount property
func (m *UserSimulationDetails) GetAssignedTrainingsCount()(*int32) {
    return m.assignedTrainingsCount
}
// GetCompletedTrainingsCount gets the completedTrainingsCount property value. The completedTrainingsCount property
func (m *UserSimulationDetails) GetCompletedTrainingsCount()(*int32) {
    return m.completedTrainingsCount
}
// GetCompromisedDateTime gets the compromisedDateTime property value. The compromisedDateTime property
func (m *UserSimulationDetails) GetCompromisedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.compromisedDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSimulationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedTrainingsCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetAssignedTrainingsCount)
    res["completedTrainingsCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCompletedTrainingsCount)
    res["compromisedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCompromisedDateTime)
    res["inProgressTrainingsCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetInProgressTrainingsCount)
    res["isCompromised"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCompromised)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["reportedPhishDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetReportedPhishDateTime)
    res["simulationEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserSimulationEventInfoFromDiscriminatorValue , m.SetSimulationEvents)
    res["simulationUser"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAttackSimulationUserFromDiscriminatorValue , m.SetSimulationUser)
    res["trainingEvents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserTrainingEventInfoFromDiscriminatorValue , m.SetTrainingEvents)
    return res
}
// GetInProgressTrainingsCount gets the inProgressTrainingsCount property value. The inProgressTrainingsCount property
func (m *UserSimulationDetails) GetInProgressTrainingsCount()(*int32) {
    return m.inProgressTrainingsCount
}
// GetIsCompromised gets the isCompromised property value. The isCompromised property
func (m *UserSimulationDetails) GetIsCompromised()(*bool) {
    return m.isCompromised
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserSimulationDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetReportedPhishDateTime gets the reportedPhishDateTime property value. The reportedPhishDateTime property
func (m *UserSimulationDetails) GetReportedPhishDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reportedPhishDateTime
}
// GetSimulationEvents gets the simulationEvents property value. The simulationEvents property
func (m *UserSimulationDetails) GetSimulationEvents()([]UserSimulationEventInfoable) {
    return m.simulationEvents
}
// GetSimulationUser gets the simulationUser property value. The simulationUser property
func (m *UserSimulationDetails) GetSimulationUser()(AttackSimulationUserable) {
    return m.simulationUser
}
// GetTrainingEvents gets the trainingEvents property value. The trainingEvents property
func (m *UserSimulationDetails) GetTrainingEvents()([]UserTrainingEventInfoable) {
    return m.trainingEvents
}
// Serialize serializes information the current object
func (m *UserSimulationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("assignedTrainingsCount", m.GetAssignedTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("completedTrainingsCount", m.GetCompletedTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("compromisedDateTime", m.GetCompromisedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inProgressTrainingsCount", m.GetInProgressTrainingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCompromised", m.GetIsCompromised())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("reportedPhishDateTime", m.GetReportedPhishDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSimulationEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSimulationEvents())
        err := writer.WriteCollectionOfObjectValues("simulationEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("simulationUser", m.GetSimulationUser())
        if err != nil {
            return err
        }
    }
    if m.GetTrainingEvents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTrainingEvents())
        err := writer.WriteCollectionOfObjectValues("trainingEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSimulationDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignedTrainingsCount sets the assignedTrainingsCount property value. The assignedTrainingsCount property
func (m *UserSimulationDetails) SetAssignedTrainingsCount(value *int32)() {
    m.assignedTrainingsCount = value
}
// SetCompletedTrainingsCount sets the completedTrainingsCount property value. The completedTrainingsCount property
func (m *UserSimulationDetails) SetCompletedTrainingsCount(value *int32)() {
    m.completedTrainingsCount = value
}
// SetCompromisedDateTime sets the compromisedDateTime property value. The compromisedDateTime property
func (m *UserSimulationDetails) SetCompromisedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.compromisedDateTime = value
}
// SetInProgressTrainingsCount sets the inProgressTrainingsCount property value. The inProgressTrainingsCount property
func (m *UserSimulationDetails) SetInProgressTrainingsCount(value *int32)() {
    m.inProgressTrainingsCount = value
}
// SetIsCompromised sets the isCompromised property value. The isCompromised property
func (m *UserSimulationDetails) SetIsCompromised(value *bool)() {
    m.isCompromised = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserSimulationDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReportedPhishDateTime sets the reportedPhishDateTime property value. The reportedPhishDateTime property
func (m *UserSimulationDetails) SetReportedPhishDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reportedPhishDateTime = value
}
// SetSimulationEvents sets the simulationEvents property value. The simulationEvents property
func (m *UserSimulationDetails) SetSimulationEvents(value []UserSimulationEventInfoable)() {
    m.simulationEvents = value
}
// SetSimulationUser sets the simulationUser property value. The simulationUser property
func (m *UserSimulationDetails) SetSimulationUser(value AttackSimulationUserable)() {
    m.simulationUser = value
}
// SetTrainingEvents sets the trainingEvents property value. The trainingEvents property
func (m *UserSimulationDetails) SetTrainingEvents(value []UserTrainingEventInfoable)() {
    m.trainingEvents = value
}
