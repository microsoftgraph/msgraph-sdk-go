package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GroupLifecyclePolicy struct {
    Entity
    // List of email address to send notifications for groups without owners. Multiple email address can be defined by separating email address with a semicolon.
    alternateNotificationEmails *string;
    // Number of days before a group expires and needs to be renewed. Once renewed, the group expiration is extended by the number of days defined.
    groupLifetimeInDays *int32;
    // The group type for which the expiration policy applies. Possible values are All, Selected or None.
    managedGroupTypes *string;
}
// Instantiates a new groupLifecyclePolicy and sets the default values.
func NewGroupLifecyclePolicy()(*GroupLifecyclePolicy) {
    m := &GroupLifecyclePolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the alternateNotificationEmails property value. List of email address to send notifications for groups without owners. Multiple email address can be defined by separating email address with a semicolon.
func (m *GroupLifecyclePolicy) GetAlternateNotificationEmails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateNotificationEmails
    }
}
// Gets the groupLifetimeInDays property value. Number of days before a group expires and needs to be renewed. Once renewed, the group expiration is extended by the number of days defined.
func (m *GroupLifecyclePolicy) GetGroupLifetimeInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupLifetimeInDays
    }
}
// Gets the managedGroupTypes property value. The group type for which the expiration policy applies. Possible values are All, Selected or None.
func (m *GroupLifecyclePolicy) GetManagedGroupTypes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedGroupTypes
    }
}
// The deserialization information for the current model
func (m *GroupLifecyclePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alternateNotificationEmails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlternateNotificationEmails(val)
        return nil
    }
    res["groupLifetimeInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetGroupLifetimeInDays(val)
        return nil
    }
    res["managedGroupTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedGroupTypes(val)
        return nil
    }
    return res
}
func (m *GroupLifecyclePolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GroupLifecyclePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("alternateNotificationEmails", m.GetAlternateNotificationEmails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("groupLifetimeInDays", m.GetGroupLifetimeInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedGroupTypes", m.GetManagedGroupTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the alternateNotificationEmails property value. List of email address to send notifications for groups without owners. Multiple email address can be defined by separating email address with a semicolon.
// Parameters:
//  - value : Value to set for the alternateNotificationEmails property.
func (m *GroupLifecyclePolicy) SetAlternateNotificationEmails(value *string)() {
    m.alternateNotificationEmails = value
}
// Sets the groupLifetimeInDays property value. Number of days before a group expires and needs to be renewed. Once renewed, the group expiration is extended by the number of days defined.
// Parameters:
//  - value : Value to set for the groupLifetimeInDays property.
func (m *GroupLifecyclePolicy) SetGroupLifetimeInDays(value *int32)() {
    m.groupLifetimeInDays = value
}
// Sets the managedGroupTypes property value. The group type for which the expiration policy applies. Possible values are All, Selected or None.
// Parameters:
//  - value : Value to set for the managedGroupTypes property.
func (m *GroupLifecyclePolicy) SetManagedGroupTypes(value *string)() {
    m.managedGroupTypes = value
}
