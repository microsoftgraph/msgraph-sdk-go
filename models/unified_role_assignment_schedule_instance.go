package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignmentScheduleInstance 
type UnifiedRoleAssignmentScheduleInstance struct {
    UnifiedRoleScheduleInstanceBase
    // If the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
    activatedUsing UnifiedRoleEligibilityScheduleInstanceable;
    // Type of the assignment. It can either be Assigned or Activated.
    assignmentType *string;
    // Time that the roleAssignmentInstance will expire
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Membership type of the assignment. It can either be Inherited, Direct, or Group.
    memberType *string;
    // ID of the roleAssignment in the directory
    roleAssignmentOriginId *string;
    // ID of the parent roleAssignmentSchedule for this instance
    roleAssignmentScheduleId *string;
    // Time that the roleAssignmentInstance will start
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewUnifiedRoleAssignmentScheduleInstance instantiates a new unifiedRoleAssignmentScheduleInstance and sets the default values.
func NewUnifiedRoleAssignmentScheduleInstance()(*UnifiedRoleAssignmentScheduleInstance) {
    m := &UnifiedRoleAssignmentScheduleInstance{
        UnifiedRoleScheduleInstanceBase: *NewUnifiedRoleScheduleInstanceBase(),
    }
    return m
}
// CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleAssignmentScheduleInstance(), nil
}
// GetActivatedUsing gets the activatedUsing property value. If the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
func (m *UnifiedRoleAssignmentScheduleInstance) GetActivatedUsing()(UnifiedRoleEligibilityScheduleInstanceable) {
    if m == nil {
        return nil
    } else {
        return m.activatedUsing
    }
}
// GetAssignmentType gets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
func (m *UnifiedRoleAssignmentScheduleInstance) GetAssignmentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentType
    }
}
// GetEndDateTime gets the endDateTime property value. Time that the roleAssignmentInstance will expire
func (m *UnifiedRoleAssignmentScheduleInstance) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleAssignmentScheduleInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleScheduleInstanceBase.GetFieldDeserializers()
    res["activatedUsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(UnifiedRoleEligibilityScheduleInstanceable))
        }
        return nil
    }
    res["assignmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentType(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["memberType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberType(val)
        }
        return nil
    }
    res["roleAssignmentOriginId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleAssignmentOriginId(val)
        }
        return nil
    }
    res["roleAssignmentScheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleAssignmentScheduleId(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetMemberType gets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleAssignmentScheduleInstance) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// GetRoleAssignmentOriginId gets the roleAssignmentOriginId property value. ID of the roleAssignment in the directory
func (m *UnifiedRoleAssignmentScheduleInstance) GetRoleAssignmentOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentOriginId
    }
}
// GetRoleAssignmentScheduleId gets the roleAssignmentScheduleId property value. ID of the parent roleAssignmentSchedule for this instance
func (m *UnifiedRoleAssignmentScheduleInstance) GetRoleAssignmentScheduleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignmentScheduleId
    }
}
// GetStartDateTime gets the startDateTime property value. Time that the roleAssignmentInstance will start
func (m *UnifiedRoleAssignmentScheduleInstance) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleAssignmentScheduleInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleScheduleInstanceBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activatedUsing", m.GetActivatedUsing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("assignmentType", m.GetAssignmentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("memberType", m.GetMemberType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleAssignmentOriginId", m.GetRoleAssignmentOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleAssignmentScheduleId", m.GetRoleAssignmentScheduleId())
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
    return nil
}
// SetActivatedUsing sets the activatedUsing property value. If the roleAssignmentScheduleInstance is activated by a roleEligibilityScheduleRequest, this is the link to the related schedule instance.
func (m *UnifiedRoleAssignmentScheduleInstance) SetActivatedUsing(value UnifiedRoleEligibilityScheduleInstanceable)() {
    if m != nil {
        m.activatedUsing = value
    }
}
// SetAssignmentType sets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
func (m *UnifiedRoleAssignmentScheduleInstance) SetAssignmentType(value *string)() {
    if m != nil {
        m.assignmentType = value
    }
}
// SetEndDateTime sets the endDateTime property value. Time that the roleAssignmentInstance will expire
func (m *UnifiedRoleAssignmentScheduleInstance) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetMemberType sets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleAssignmentScheduleInstance) SetMemberType(value *string)() {
    if m != nil {
        m.memberType = value
    }
}
// SetRoleAssignmentOriginId sets the roleAssignmentOriginId property value. ID of the roleAssignment in the directory
func (m *UnifiedRoleAssignmentScheduleInstance) SetRoleAssignmentOriginId(value *string)() {
    if m != nil {
        m.roleAssignmentOriginId = value
    }
}
// SetRoleAssignmentScheduleId sets the roleAssignmentScheduleId property value. ID of the parent roleAssignmentSchedule for this instance
func (m *UnifiedRoleAssignmentScheduleInstance) SetRoleAssignmentScheduleId(value *string)() {
    if m != nil {
        m.roleAssignmentScheduleId = value
    }
}
// SetStartDateTime sets the startDateTime property value. Time that the roleAssignmentInstance will start
func (m *UnifiedRoleAssignmentScheduleInstance) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
