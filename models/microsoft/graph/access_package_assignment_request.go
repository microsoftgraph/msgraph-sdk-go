package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentRequest 
type AccessPackageAssignmentRequest struct {
    Entity
    // The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
    accessPackage *AccessPackage;
    // 
    assignment *AccessPackageAssignment;
    // 
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
    requestor *AccessPackageSubject;
    // One of UserAdd, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd or UserRemove. Read-only.
    requestType *AccessPackageRequestType;
    // The range of dates that access is to be assigned to the requestor. Read-only.
    schedule *EntitlementManagementSchedule;
    // 
    state *AccessPackageRequestState;
    // 
    status *string;
}
// NewAccessPackageAssignmentRequest instantiates a new accessPackageAssignmentRequest and sets the default values.
func NewAccessPackageAssignmentRequest()(*AccessPackageAssignmentRequest) {
    m := &AccessPackageAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackage gets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackage()(*AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// GetAssignment gets the assignment property value. 
func (m *AccessPackageAssignmentRequest) GetAssignment()(*AccessPackageAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignment
    }
}
// GetCompletedDateTime gets the completedDateTime property value. 
func (m *AccessPackageAssignmentRequest) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetRequestor gets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) GetRequestor()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.requestor
    }
}
// GetRequestType gets the requestType property value. One of UserAdd, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd or UserRemove. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestType()(*AccessPackageRequestType) {
    if m == nil {
        return nil
    } else {
        return m.requestType
    }
}
// GetSchedule gets the schedule property value. The range of dates that access is to be assigned to the requestor. Read-only.
func (m *AccessPackageAssignmentRequest) GetSchedule()(*EntitlementManagementSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetState gets the state property value. 
func (m *AccessPackageAssignmentRequest) GetState()(*AccessPackageRequestState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStatus gets the status property value. 
func (m *AccessPackageAssignmentRequest) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(*AccessPackage))
        }
        return nil
    }
    res["assignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignment(val.(*AccessPackageAssignment))
        }
        return nil
    }
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
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
    res["requestor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestor(val.(*AccessPackageSubject))
        }
        return nil
    }
    res["requestType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageRequestType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AccessPackageRequestType)
            m.SetRequestType(&cast)
        }
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntitlementManagementSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(*EntitlementManagementSchedule))
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageRequestState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AccessPackageRequestState)
            m.SetState(&cast)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
func (m *AccessPackageAssignmentRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackage", m.GetAccessPackage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignment", m.GetAssignment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
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
        err = writer.WriteObjectValue("requestor", m.GetRequestor())
        if err != nil {
            return err
        }
    }
    if m.GetRequestType() != nil {
        cast := m.GetRequestType().String()
        err = writer.WriteStringValue("requestType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
func (m *AccessPackageAssignmentRequest) SetAccessPackage(value *AccessPackage)() {
    if m != nil {
        m.accessPackage = value
    }
}
// SetAssignment sets the assignment property value. 
func (m *AccessPackageAssignmentRequest) SetAssignment(value *AccessPackageAssignment)() {
    if m != nil {
        m.assignment = value
    }
}
// SetCompletedDateTime sets the completedDateTime property value. 
func (m *AccessPackageAssignmentRequest) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.completedDateTime = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetRequestor sets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) SetRequestor(value *AccessPackageSubject)() {
    if m != nil {
        m.requestor = value
    }
}
// SetRequestType sets the requestType property value. One of UserAdd, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd or UserRemove. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestType(value *AccessPackageRequestType)() {
    if m != nil {
        m.requestType = value
    }
}
// SetSchedule sets the schedule property value. The range of dates that access is to be assigned to the requestor. Read-only.
func (m *AccessPackageAssignmentRequest) SetSchedule(value *EntitlementManagementSchedule)() {
    if m != nil {
        m.schedule = value
    }
}
// SetState sets the state property value. 
func (m *AccessPackageAssignmentRequest) SetState(value *AccessPackageRequestState)() {
    if m != nil {
        m.state = value
    }
}
// SetStatus sets the status property value. 
func (m *AccessPackageAssignmentRequest) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
