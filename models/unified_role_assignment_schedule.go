package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignmentSchedule provides operations to manage the roleManagement singleton.
type UnifiedRoleAssignmentSchedule struct {
    UnifiedRoleScheduleBase
    // If the roleAssignmentSchedule is activated by a roleEligibilitySchedule, this is the link to that schedule.
    activatedUsing UnifiedRoleEligibilityScheduleable
    // Type of the assignment. It can either be Assigned or Activated.
    assignmentType *string
    // Membership type of the assignment. It can either be Inherited, Direct, or Group.
    memberType *string
    // The schedule object of the role assignment request.
    scheduleInfo RequestScheduleable
}
// NewUnifiedRoleAssignmentSchedule instantiates a new unifiedRoleAssignmentSchedule and sets the default values.
func NewUnifiedRoleAssignmentSchedule()(*UnifiedRoleAssignmentSchedule) {
    m := &UnifiedRoleAssignmentSchedule{
        UnifiedRoleScheduleBase: *NewUnifiedRoleScheduleBase(),
    }
    return m
}
// CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleAssignmentSchedule(), nil
}
// GetActivatedUsing gets the activatedUsing property value. If the roleAssignmentSchedule is activated by a roleEligibilitySchedule, this is the link to that schedule.
func (m *UnifiedRoleAssignmentSchedule) GetActivatedUsing()(UnifiedRoleEligibilityScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.activatedUsing
    }
}
// GetAssignmentType gets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
func (m *UnifiedRoleAssignmentSchedule) GetAssignmentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleAssignmentSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleScheduleBase.GetFieldDeserializers()
    res["activatedUsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(UnifiedRoleEligibilityScheduleable))
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
    res["scheduleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleInfo(val.(RequestScheduleable))
        }
        return nil
    }
    return res
}
// GetMemberType gets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleAssignmentSchedule) GetMemberType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberType
    }
}
// GetScheduleInfo gets the scheduleInfo property value. The schedule object of the role assignment request.
func (m *UnifiedRoleAssignmentSchedule) GetScheduleInfo()(RequestScheduleable) {
    if m == nil {
        return nil
    } else {
        return m.scheduleInfo
    }
}
// Serialize serializes information the current object
func (m *UnifiedRoleAssignmentSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleScheduleBase.Serialize(writer)
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
        err = writer.WriteStringValue("memberType", m.GetMemberType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivatedUsing sets the activatedUsing property value. If the roleAssignmentSchedule is activated by a roleEligibilitySchedule, this is the link to that schedule.
func (m *UnifiedRoleAssignmentSchedule) SetActivatedUsing(value UnifiedRoleEligibilityScheduleable)() {
    if m != nil {
        m.activatedUsing = value
    }
}
// SetAssignmentType sets the assignmentType property value. Type of the assignment. It can either be Assigned or Activated.
func (m *UnifiedRoleAssignmentSchedule) SetAssignmentType(value *string)() {
    if m != nil {
        m.assignmentType = value
    }
}
// SetMemberType sets the memberType property value. Membership type of the assignment. It can either be Inherited, Direct, or Group.
func (m *UnifiedRoleAssignmentSchedule) SetMemberType(value *string)() {
    if m != nil {
        m.memberType = value
    }
}
// SetScheduleInfo sets the scheduleInfo property value. The schedule object of the role assignment request.
func (m *UnifiedRoleAssignmentSchedule) SetScheduleInfo(value RequestScheduleable)() {
    if m != nil {
        m.scheduleInfo = value
    }
}
