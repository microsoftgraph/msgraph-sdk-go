package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsAsyncOperation provides operations to manage the educationRoot singleton.
type TeamsAsyncOperation struct {
    Entity
    // Number of times the operation was attempted before being marked successful or failed.
    attemptsCount *int32;
    // Time when the operation was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Any error that causes the async operation to fail.
    error OperationErrorable;
    // Time when the async operation was last updated.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Denotes which type of operation is being described.
    operationType *TeamsAsyncOperationType;
    // Operation status.
    status *TeamsAsyncOperationStatus;
    // The ID of the object that's created or modified as result of this async operation, typically a team.
    targetResourceId *string;
    // The location of the object that's created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths.
    targetResourceLocation *string;
}
// NewTeamsAsyncOperation instantiates a new teamsAsyncOperation and sets the default values.
func NewTeamsAsyncOperation()(*TeamsAsyncOperation) {
    m := &TeamsAsyncOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAsyncOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAsyncOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamsAsyncOperation(), nil
}
// GetAttemptsCount gets the attemptsCount property value. Number of times the operation was attempted before being marked successful or failed.
func (m *TeamsAsyncOperation) GetAttemptsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.attemptsCount
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Time when the operation was created.
func (m *TeamsAsyncOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetError gets the error property value. Any error that causes the async operation to fail.
func (m *TeamsAsyncOperation) GetError()(OperationErrorable) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAsyncOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attemptsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttemptsCount(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateOperationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(OperationErrorable))
        }
        return nil
    }
    res["lastActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["operationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAsyncOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*TeamsAsyncOperationType))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAsyncOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TeamsAsyncOperationStatus))
        }
        return nil
    }
    res["targetResourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetResourceId(val)
        }
        return nil
    }
    res["targetResourceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetResourceLocation(val)
        }
        return nil
    }
    return res
}
// GetLastActionDateTime gets the lastActionDateTime property value. Time when the async operation was last updated.
func (m *TeamsAsyncOperation) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// GetOperationType gets the operationType property value. Denotes which type of operation is being described.
func (m *TeamsAsyncOperation) GetOperationType()(*TeamsAsyncOperationType) {
    if m == nil {
        return nil
    } else {
        return m.operationType
    }
}
// GetStatus gets the status property value. Operation status.
func (m *TeamsAsyncOperation) GetStatus()(*TeamsAsyncOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTargetResourceId gets the targetResourceId property value. The ID of the object that's created or modified as result of this async operation, typically a team.
func (m *TeamsAsyncOperation) GetTargetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetResourceId
    }
}
// GetTargetResourceLocation gets the targetResourceLocation property value. The location of the object that's created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths.
func (m *TeamsAsyncOperation) GetTargetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetResourceLocation
    }
}
func (m *TeamsAsyncOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamsAsyncOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("attemptsCount", m.GetAttemptsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOperationType() != nil {
        cast := (*m.GetOperationType()).String()
        err = writer.WriteStringValue("operationType", &cast)
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
    {
        err = writer.WriteStringValue("targetResourceId", m.GetTargetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetResourceLocation", m.GetTargetResourceLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttemptsCount sets the attemptsCount property value. Number of times the operation was attempted before being marked successful or failed.
func (m *TeamsAsyncOperation) SetAttemptsCount(value *int32)() {
    if m != nil {
        m.attemptsCount = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Time when the operation was created.
func (m *TeamsAsyncOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetError sets the error property value. Any error that causes the async operation to fail.
func (m *TeamsAsyncOperation) SetError(value OperationErrorable)() {
    if m != nil {
        m.error = value
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. Time when the async operation was last updated.
func (m *TeamsAsyncOperation) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActionDateTime = value
    }
}
// SetOperationType sets the operationType property value. Denotes which type of operation is being described.
func (m *TeamsAsyncOperation) SetOperationType(value *TeamsAsyncOperationType)() {
    if m != nil {
        m.operationType = value
    }
}
// SetStatus sets the status property value. Operation status.
func (m *TeamsAsyncOperation) SetStatus(value *TeamsAsyncOperationStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTargetResourceId sets the targetResourceId property value. The ID of the object that's created or modified as result of this async operation, typically a team.
func (m *TeamsAsyncOperation) SetTargetResourceId(value *string)() {
    if m != nil {
        m.targetResourceId = value
    }
}
// SetTargetResourceLocation sets the targetResourceLocation property value. The location of the object that's created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths.
func (m *TeamsAsyncOperation) SetTargetResourceLocation(value *string)() {
    if m != nil {
        m.targetResourceLocation = value
    }
}
