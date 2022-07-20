package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignment 
type AccessPackageAssignment struct {
    Entity
    // Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
    accessPackage AccessPackageable
    // Read-only. Supports $filter (eq) on the id property and $expand query parameters.
    assignmentPolicy AccessPackageAssignmentPolicyable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    expiredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // When the access assignment is to be in place. Read-only.
    schedule EntitlementManagementScheduleable
    // The state of the access package assignment. The possible values are: delivering, partiallyDelivered, delivered, expired, deliveryFailed, unknownFutureValue. Read-only. Supports $filter (eq).
    state *AccessPackageAssignmentState
    // More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
    status *string
    // The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
    target AccessPackageSubjectable
}
// NewAccessPackageAssignment instantiates a new accessPackageAssignment and sets the default values.
func NewAccessPackageAssignment()(*AccessPackageAssignment) {
    m := &AccessPackageAssignment{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageAssignment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignment(), nil
}
// GetAccessPackage gets the accessPackage property value. Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) GetAccessPackage()(AccessPackageable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// GetAssignmentPolicy gets the assignmentPolicy property value. Read-only. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) GetAssignmentPolicy()(AccessPackageAssignmentPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.assignmentPolicy
    }
}
// GetExpiredDateTime gets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignment) GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiredDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["assignmentPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentPolicy(val.(AccessPackageAssignmentPolicyable))
        }
        return nil
    }
    res["expiredDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiredDateTime(val)
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
        val, err := n.GetEnumValue(ParseAccessPackageAssignmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AccessPackageAssignmentState))
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
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(AccessPackageSubjectable))
        }
        return nil
    }
    return res
}
// GetSchedule gets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) GetSchedule()(EntitlementManagementScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetState gets the state property value. The state of the access package assignment. The possible values are: delivering, partiallyDelivered, delivered, expired, deliveryFailed, unknownFutureValue. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) GetState()(*AccessPackageAssignmentState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetStatus gets the status property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTarget gets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) GetTarget()(AccessPackageSubjectable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Serialize serializes information the current object
func (m *AccessPackageAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("assignmentPolicy", m.GetAssignmentPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expiredDateTime", m.GetExpiredDateTime())
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
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. Read-only. Nullable. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) SetAccessPackage(value AccessPackageable)() {
    if m != nil {
        m.accessPackage = value
    }
}
// SetAssignmentPolicy sets the assignmentPolicy property value. Read-only. Supports $filter (eq) on the id property and $expand query parameters.
func (m *AccessPackageAssignment) SetAssignmentPolicy(value AccessPackageAssignmentPolicyable)() {
    if m != nil {
        m.assignmentPolicy = value
    }
}
// SetExpiredDateTime sets the expiredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *AccessPackageAssignment) SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expiredDateTime = value
    }
}
// SetSchedule sets the schedule property value. When the access assignment is to be in place. Read-only.
func (m *AccessPackageAssignment) SetSchedule(value EntitlementManagementScheduleable)() {
    if m != nil {
        m.schedule = value
    }
}
// SetState sets the state property value. The state of the access package assignment. The possible values are: delivering, partiallyDelivered, delivered, expired, deliveryFailed, unknownFutureValue. Read-only. Supports $filter (eq).
func (m *AccessPackageAssignment) SetState(value *AccessPackageAssignmentState)() {
    if m != nil {
        m.state = value
    }
}
// SetStatus sets the status property value. More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only.
func (m *AccessPackageAssignment) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTarget sets the target property value. The subject of the access package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (m *AccessPackageAssignment) SetTarget(value AccessPackageSubjectable)() {
    if m != nil {
        m.target = value
    }
}
