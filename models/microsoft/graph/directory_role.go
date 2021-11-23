package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DirectoryRole 
type DirectoryRole struct {
    DirectoryObject
    // The description for the directory role. Read-only.
    description *string;
    // The display name for the directory role. Read-only.
    displayName *string;
    // Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable.
    members []DirectoryObject;
    // The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only.
    roleTemplateId *string;
    // Members of this directory role that are scoped to administrative units. Read-only. Nullable.
    scopedMembers []ScopedRoleMembership;
}
// NewDirectoryRole instantiates a new directoryRole and sets the default values.
func NewDirectoryRole()(*DirectoryRole) {
    m := &DirectoryRole{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetDescription gets the description property value. The description for the directory role. Read-only.
func (m *DirectoryRole) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the directory role. Read-only.
func (m *DirectoryRole) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetMembers gets the members property value. Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable.
func (m *DirectoryRole) GetMembers()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetRoleTemplateId gets the roleTemplateId property value. The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only.
func (m *DirectoryRole) GetRoleTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleTemplateId
    }
}
// GetScopedMembers gets the scopedMembers property value. Members of this directory role that are scoped to administrative units. Read-only. Nullable.
func (m *DirectoryRole) GetScopedMembers()([]ScopedRoleMembership) {
    if m == nil {
        return nil
    } else {
        return m.scopedMembers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryRole) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["roleTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleTemplateId(val)
        }
        return nil
    }
    res["scopedMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewScopedRoleMembership() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScopedRoleMembership, len(val))
            for i, v := range val {
                res[i] = *(v.(*ScopedRoleMembership))
            }
            m.SetScopedMembers(res)
        }
        return nil
    }
    return res
}
func (m *DirectoryRole) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DirectoryRole) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleTemplateId", m.GetRoleTemplateId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScopedMembers()))
        for i, v := range m.GetScopedMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("scopedMembers", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description for the directory role. Read-only.
func (m *DirectoryRole) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the directory role. Read-only.
func (m *DirectoryRole) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMembers sets the members property value. Users that are members of this directory role. HTTP Methods: GET, POST, DELETE. Read-only. Nullable.
func (m *DirectoryRole) SetMembers(value []DirectoryObject)() {
    m.members = value
}
// SetRoleTemplateId sets the roleTemplateId property value. The id of the directoryRoleTemplate that this role is based on. The property must be specified when activating a directory role in a tenant with a POST operation. After the directory role has been activated, the property is read only.
func (m *DirectoryRole) SetRoleTemplateId(value *string)() {
    m.roleTemplateId = value
}
// SetScopedMembers sets the scopedMembers property value. Members of this directory role that are scoped to administrative units. Read-only. Nullable.
func (m *DirectoryRole) SetScopedMembers(value []ScopedRoleMembership)() {
    m.scopedMembers = value
}
