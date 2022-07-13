package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttributeRuleMembers 
type AttributeRuleMembers struct {
    SubjectSet
    // The description property
    description *string
    // The membershipRule property
    membershipRule *string
}
// NewAttributeRuleMembers instantiates a new AttributeRuleMembers and sets the default values.
func NewAttributeRuleMembers()(*AttributeRuleMembers) {
    m := &AttributeRuleMembers{
        SubjectSet: *NewSubjectSet(),
    }
    return m
}
// CreateAttributeRuleMembersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttributeRuleMembersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttributeRuleMembers(), nil
}
// GetDescription gets the description property value. The description property
func (m *AttributeRuleMembers) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttributeRuleMembers) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["membershipRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRule(val)
        }
        return nil
    }
    return res
}
// GetMembershipRule gets the membershipRule property value. The membershipRule property
func (m *AttributeRuleMembers) GetMembershipRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.membershipRule
    }
}
// Serialize serializes information the current object
func (m *AttributeRuleMembers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("membershipRule", m.GetMembershipRule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description property
func (m *AttributeRuleMembers) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetMembershipRule sets the membershipRule property value. The membershipRule property
func (m *AttributeRuleMembers) SetMembershipRule(value *string)() {
    if m != nil {
        m.membershipRule = value
    }
}
