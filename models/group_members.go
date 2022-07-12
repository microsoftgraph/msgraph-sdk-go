package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupMembers 
type GroupMembers struct {
    SubjectSet
    // The name of the group in Azure AD. Read only.
    description *string
    // The ID of the group in Azure AD.
    groupId *string
}
// NewGroupMembers instantiates a new GroupMembers and sets the default values.
func NewGroupMembers()(*GroupMembers) {
    m := &GroupMembers{
        SubjectSet: *NewSubjectSet(),
    }
    return m
}
// CreateGroupMembersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupMembersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupMembers(), nil
}
// GetDescription gets the description property value. The name of the group in Azure AD. Read only.
func (m *GroupMembers) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupMembers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectSet.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The ID of the group in Azure AD.
func (m *GroupMembers) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
// Serialize serializes information the current object
func (m *GroupMembers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectSet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The name of the group in Azure AD. Read only.
func (m *GroupMembers) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetGroupId sets the groupId property value. The ID of the group in Azure AD.
func (m *GroupMembers) SetGroupId(value *string)() {
    if m != nil {
        m.groupId = value
    }
}
