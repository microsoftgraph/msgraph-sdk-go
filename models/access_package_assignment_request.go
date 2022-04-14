package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequest 
type AccessPackageAssignmentRequest struct {
    Entity
    // The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
    accessPackage AccessPackageable
    // For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
    assignment AccessPackageAssignmentable
    // The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
    requestor AccessPackageSubjectable
    // The type of the request. The possible values are: notSpecified, userAdd, userUpdate, userRemove, adminAdd, adminUpdate, adminRemove, systemAdd, systemUpdate, systemRemove, onBehalfAdd, unknownFutureValue. A request from the user themselves would have requestType of UserAdd or UserRemove. This property cannot be changed once set.
    requestType *AccessPackageRequestType
    // The range of dates that access is to be assigned to the requestor. This property cannot be changed once set.
    schedule EntitlementManagementScheduleable
    // The state of the request. The possible values are: submitted, pendingApproval, delivering, delivered, deliveryFailed, denied, scheduled, canceled, partiallyDelivered, unknownFutureValue. Read-only.
    state *AccessPackageRequestState
    // More information on the request processing status. Read-only.
    status *string
}
// NewAccessPackageAssignmentRequest instantiates a new accessPackageAssignmentRequest and sets the default values.
func NewAccessPackageAssignmentRequest()(*AccessPackageAssignmentRequest) {
    m := &AccessPackageAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequest(), nil
}
// GetAccessPackage gets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackage()(AccessPackageable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// GetAssignment gets the assignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAssignment()(AccessPackageAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignment
    }
}
// GetCompletedDateTime gets the completedDateTime property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
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
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(AccessPackageable))
        }
        return nil
    }
    res["assignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignment(val.(AccessPackageAssignmentable))
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["requestor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestor(val.(AccessPackageSubjectable))
        }
        return nil
    }
    res["requestType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageRequestType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestType(val.(*AccessPackageRequestType))
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntitlementManagementScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(EntitlementManagementScheduleable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageRequestState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AccessPackageRequestState))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetRequestor gets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) GetRequestor()(AccessPackageSubjectable) {
    if m == nil {
        return nil
    } else {
        return m.requestor
    }
}
// GetRequestType gets the requestType property value. The type of the request. The possible values are: notSpecified, userAdd, userUpdate, userRemove, adminAdd, adminUpdate, adminRemove, systemAdd, systemUpdate, systemRemove, onBehalfAdd, unknownFutureValue. A request from the user themselves would have requestType of UserAdd or UserRemove. This property cannot be changed once set.
func (m *AccessPackageAssignmentRequest) GetRequestType()(*AccessPackageRequestType) {
    if m == nil {
        return nil
    } else {
        return m.requestType
    }
}
// GetSchedule gets the schedule property value. The range of dates that access is to be assigned to the requestor. This property cannot be changed once set.
func (m *AccessPackageAssignmentRequest) GetSchedule()(EntitlementManagementScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetState gets the state property value. The state of the request. The possible values are: submitted, pendingApproval, delivering, delivered, deliveryFailed, denied, scheduled, canceled, partiallyDelivered, unknownFutureValue. Read-only.
func (m *AccessPackageAssignmentRequest) GetState()(*AccessPackageRequestState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStatus gets the status property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetRequestType()).String()
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
        cast := (*m.GetState()).String()
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
func (m *AccessPackageAssignmentRequest) SetAccessPackage(value AccessPackageable)() {
    if m != nil {
        m.accessPackage = value
    }
}
// SetAssignment sets the assignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequest) SetAssignment(value AccessPackageAssignmentable)() {
    if m != nil {
        m.assignment = value
    }
}
// SetCompletedDateTime sets the completedDateTime property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
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
func (m *AccessPackageAssignmentRequest) SetRequestor(value AccessPackageSubjectable)() {
    if m != nil {
        m.requestor = value
    }
}
// SetRequestType sets the requestType property value. The type of the request. The possible values are: notSpecified, userAdd, userUpdate, userRemove, adminAdd, adminUpdate, adminRemove, systemAdd, systemUpdate, systemRemove, onBehalfAdd, unknownFutureValue. A request from the user themselves would have requestType of UserAdd or UserRemove. This property cannot be changed once set.
func (m *AccessPackageAssignmentRequest) SetRequestType(value *AccessPackageRequestType)() {
    if m != nil {
        m.requestType = value
    }
}
// SetSchedule sets the schedule property value. The range of dates that access is to be assigned to the requestor. This property cannot be changed once set.
func (m *AccessPackageAssignmentRequest) SetSchedule(value EntitlementManagementScheduleable)() {
    if m != nil {
        m.schedule = value
    }
}
// SetState sets the state property value. The state of the request. The possible values are: submitted, pendingApproval, delivering, delivered, deliveryFailed, denied, scheduled, canceled, partiallyDelivered, unknownFutureValue. Read-only.
func (m *AccessPackageAssignmentRequest) SetState(value *AccessPackageRequestState)() {
    if m != nil {
        m.state = value
    }
}
// SetStatus sets the status property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
