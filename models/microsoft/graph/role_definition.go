package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleDefinition 
type RoleDefinition struct {
    Entity
    // Description of the Role definition.
    description *string;
    // Display Name of the Role definition.
    displayName *string;
    // Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
    isBuiltIn *bool;
    // List of Role assignments for this role definition.
    roleAssignments []RoleAssignment;
    // List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
    rolePermissions []RolePermission;
}
// NewRoleDefinition instantiates a new roleDefinition and sets the default values.
func NewRoleDefinition()(*RoleDefinition) {
    m := &RoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Description of the Role definition.
func (m *RoleDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display Name of the Role definition.
func (m *RoleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIsBuiltIn gets the isBuiltIn property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// GetRoleAssignments gets the roleAssignments property value. List of Role assignments for this role definition.
func (m *RoleDefinition) GetRoleAssignments()([]RoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// GetRolePermissions gets the rolePermissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) GetRolePermissions()([]RolePermission) {
    if m == nil {
        return nil
    } else {
        return m.rolePermissions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*RoleAssignment))
            }
            m.SetRoleAssignments(res)
        }
        return nil
    }
    res["rolePermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRolePermission() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RolePermission, len(val))
            for i, v := range val {
                res[i] = *(v.(*RolePermission))
            }
            m.SetRolePermissions(res)
        }
        return nil
    }
    return res
}
func (m *RoleDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RoleDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRolePermissions()))
        for i, v := range m.GetRolePermissions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("rolePermissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the Role definition.
func (m *RoleDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display Name of the Role definition.
func (m *RoleDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsBuiltIn sets the isBuiltIn property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// SetRoleAssignments sets the roleAssignments property value. List of Role assignments for this role definition.
func (m *RoleDefinition) SetRoleAssignments(value []RoleAssignment)() {
    m.roleAssignments = value
}
// SetRolePermissions sets the rolePermissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) SetRolePermissions(value []RolePermission)() {
    m.rolePermissions = value
}
