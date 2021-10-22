package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamsAsyncOperation struct {
    Entity
    attemptsCount *int32;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    error *OperationError;
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    operationType *TeamsAsyncOperationType;
    status *TeamsAsyncOperationStatus;
    targetResourceId *string;
    targetResourceLocation *string;
}
func NewTeamsAsyncOperation()(*TeamsAsyncOperation) {
    m := &TeamsAsyncOperation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TeamsAsyncOperation) GetAttemptsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.attemptsCount
    }
}
func (m *TeamsAsyncOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *TeamsAsyncOperation) GetError()(*OperationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *TeamsAsyncOperation) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
func (m *TeamsAsyncOperation) GetOperationType()(*TeamsAsyncOperationType) {
    if m == nil {
        return nil
    } else {
        return m.operationType
    }
}
func (m *TeamsAsyncOperation) GetStatus()(*TeamsAsyncOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *TeamsAsyncOperation) GetTargetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetResourceId
    }
}
func (m *TeamsAsyncOperation) GetTargetResourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetResourceLocation
    }
}
func (m *TeamsAsyncOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attemptsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAttemptsCount(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOperationError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*OperationError))
        return nil
    }
    res["lastActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastActionDateTime(val)
        return nil
    }
    res["operationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAsyncOperationType)
        if err != nil {
            return err
        }
        cast := val.(TeamsAsyncOperationType)
        m.SetOperationType(&cast)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAsyncOperationStatus)
        if err != nil {
            return err
        }
        cast := val.(TeamsAsyncOperationStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["targetResourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetResourceId(val)
        return nil
    }
    res["targetResourceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetResourceLocation(val)
        return nil
    }
    return res
}
func (m *TeamsAsyncOperation) IsNil()(bool) {
    return m == nil
}
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
        cast := m.GetOperationType().String()
        err = writer.WriteStringValue("operationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
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
func (m *TeamsAsyncOperation) SetAttemptsCount(value *int32)() {
    m.attemptsCount = value
}
func (m *TeamsAsyncOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *TeamsAsyncOperation) SetError(value *OperationError)() {
    m.error = value
}
func (m *TeamsAsyncOperation) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
func (m *TeamsAsyncOperation) SetOperationType(value *TeamsAsyncOperationType)() {
    m.operationType = value
}
func (m *TeamsAsyncOperation) SetStatus(value *TeamsAsyncOperationStatus)() {
    m.status = value
}
func (m *TeamsAsyncOperation) SetTargetResourceId(value *string)() {
    m.targetResourceId = value
}
func (m *TeamsAsyncOperation) SetTargetResourceLocation(value *string)() {
    m.targetResourceLocation = value
}
