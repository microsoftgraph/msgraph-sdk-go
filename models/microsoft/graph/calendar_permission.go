package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CalendarPermission struct {
    Entity
    allowedRoles []CalendarRoleType;
    emailAddress *EmailAddress;
    isInsideOrganization *bool;
    isRemovable *bool;
    role *CalendarRoleType;
}
func NewCalendarPermission()(*CalendarPermission) {
    m := &CalendarPermission{
        Entity: *NewEntity(),
    }
    return m
}
func (m *CalendarPermission) GetAllowedRoles()([]CalendarRoleType) {
    if m == nil {
        return nil
    } else {
        return m.allowedRoles
    }
}
func (m *CalendarPermission) GetEmailAddress()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
func (m *CalendarPermission) GetIsInsideOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInsideOrganization
    }
}
func (m *CalendarPermission) GetIsRemovable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemovable
    }
}
func (m *CalendarPermission) GetRole()(*CalendarRoleType) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
func (m *CalendarPermission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseCalendarRoleType)
        if err != nil {
            return err
        }
        res := make([]CalendarRoleType, len(val))
        for i, v := range val {
            res[i] = *(v.(*CalendarRoleType))
        }
        m.SetAllowedRoles(res)
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        m.SetEmailAddress(val.(*EmailAddress))
        return nil
    }
    res["isInsideOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsInsideOrganization(val)
        return nil
    }
    res["isRemovable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRemovable(val)
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarRoleType)
        if err != nil {
            return err
        }
        cast := val.(CalendarRoleType)
        m.SetRole(&cast)
        return nil
    }
    return res
}
func (m *CalendarPermission) IsNil()(bool) {
    return m == nil
}
func (m *CalendarPermission) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("allowedRoles", SerializeCalendarRoleType(m.GetAllowedRoles()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInsideOrganization", m.GetIsInsideOrganization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRemovable", m.GetIsRemovable())
        if err != nil {
            return err
        }
    }
    if m.GetRole() != nil {
        cast := m.GetRole().String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *CalendarPermission) SetAllowedRoles(value []CalendarRoleType)() {
    m.allowedRoles = value
}
func (m *CalendarPermission) SetEmailAddress(value *EmailAddress)() {
    m.emailAddress = value
}
func (m *CalendarPermission) SetIsInsideOrganization(value *bool)() {
    m.isInsideOrganization = value
}
func (m *CalendarPermission) SetIsRemovable(value *bool)() {
    m.isRemovable = value
}
func (m *CalendarPermission) SetRole(value *CalendarRoleType)() {
    m.role = value
}
