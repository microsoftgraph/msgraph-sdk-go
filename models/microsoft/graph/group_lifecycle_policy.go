package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GroupLifecyclePolicy struct {
    Entity
    alternateNotificationEmails *string;
    groupLifetimeInDays *int32;
    managedGroupTypes *string;
}
func NewGroupLifecyclePolicy()(*GroupLifecyclePolicy) {
    m := &GroupLifecyclePolicy{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GroupLifecyclePolicy) GetAlternateNotificationEmails()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alternateNotificationEmails
    }
}
func (m *GroupLifecyclePolicy) GetGroupLifetimeInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.groupLifetimeInDays
    }
}
func (m *GroupLifecyclePolicy) GetManagedGroupTypes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedGroupTypes
    }
}
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
func (m *GroupLifecyclePolicy) SetAlternateNotificationEmails(value *string)() {
    m.alternateNotificationEmails = value
}
func (m *GroupLifecyclePolicy) SetGroupLifetimeInDays(value *int32)() {
    m.groupLifetimeInDays = value
}
func (m *GroupLifecyclePolicy) SetManagedGroupTypes(value *string)() {
    m.managedGroupTypes = value
}
