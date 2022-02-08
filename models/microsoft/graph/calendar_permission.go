package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CalendarPermission 
type CalendarPermission struct {
    Entity
    // List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
    allowedRoles []CalendarRoleType;
    // Represents a sharee or delegate who has access to the calendar. For the 'My Organization' sharee, the address property is null. Read-only.
    emailAddress *EmailAddress;
    // True if the user in context (sharee or delegate) is inside the same organization as the calendar owner.
    isInsideOrganization *bool;
    // True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The 'My organization' user determines the permissions other people within your organization have to the given calendar. You cannot remove 'My organization' as a sharee to a calendar.
    isRemovable *bool;
    // Current permission level of the calendar sharee or delegate.
    role *CalendarRoleType;
}
// NewCalendarPermission instantiates a new calendarPermission and sets the default values.
func NewCalendarPermission()(*CalendarPermission) {
    m := &CalendarPermission{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllowedRoles gets the allowedRoles property value. List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
func (m *CalendarPermission) GetAllowedRoles()([]CalendarRoleType) {
    if m == nil {
        return nil
    } else {
        return m.allowedRoles
    }
}
// GetEmailAddress gets the emailAddress property value. Represents a sharee or delegate who has access to the calendar. For the 'My Organization' sharee, the address property is null. Read-only.
func (m *CalendarPermission) GetEmailAddress()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetIsInsideOrganization gets the isInsideOrganization property value. True if the user in context (sharee or delegate) is inside the same organization as the calendar owner.
func (m *CalendarPermission) GetIsInsideOrganization()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isInsideOrganization
    }
}
// GetIsRemovable gets the isRemovable property value. True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The 'My organization' user determines the permissions other people within your organization have to the given calendar. You cannot remove 'My organization' as a sharee to a calendar.
func (m *CalendarPermission) GetIsRemovable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRemovable
    }
}
// GetRole gets the role property value. Current permission level of the calendar sharee or delegate.
func (m *CalendarPermission) GetRole()(*CalendarRoleType) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalendarPermission) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseCalendarRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CalendarRoleType, len(val))
            for i, v := range val {
                res[i] = *(v.(*CalendarRoleType))
            }
            m.SetAllowedRoles(res)
        }
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val.(*EmailAddress))
        }
        return nil
    }
    res["isInsideOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInsideOrganization(val)
        }
        return nil
    }
    res["isRemovable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRemovable(val)
        }
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*CalendarRoleType))
        }
        return nil
    }
    return res
}
func (m *CalendarPermission) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CalendarPermission) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedRoles() != nil {
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
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedRoles sets the allowedRoles property value. List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
func (m *CalendarPermission) SetAllowedRoles(value []CalendarRoleType)() {
    if m != nil {
        m.allowedRoles = value
    }
}
// SetEmailAddress sets the emailAddress property value. Represents a sharee or delegate who has access to the calendar. For the 'My Organization' sharee, the address property is null. Read-only.
func (m *CalendarPermission) SetEmailAddress(value *EmailAddress)() {
    if m != nil {
        m.emailAddress = value
    }
}
// SetIsInsideOrganization sets the isInsideOrganization property value. True if the user in context (sharee or delegate) is inside the same organization as the calendar owner.
func (m *CalendarPermission) SetIsInsideOrganization(value *bool)() {
    if m != nil {
        m.isInsideOrganization = value
    }
}
// SetIsRemovable sets the isRemovable property value. True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The 'My organization' user determines the permissions other people within your organization have to the given calendar. You cannot remove 'My organization' as a sharee to a calendar.
func (m *CalendarPermission) SetIsRemovable(value *bool)() {
    if m != nil {
        m.isRemovable = value
    }
}
// SetRole sets the role property value. Current permission level of the calendar sharee or delegate.
func (m *CalendarPermission) SetRole(value *CalendarRoleType)() {
    if m != nil {
        m.role = value
    }
}
