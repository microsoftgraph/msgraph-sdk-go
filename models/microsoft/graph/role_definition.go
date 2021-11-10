package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new roleDefinition and sets the default values.
func NewRoleDefinition()(*RoleDefinition) {
    m := &RoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. Description of the Role definition.
func (m *RoleDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Display Name of the Role definition.
func (m *RoleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isBuiltIn property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
func (m *RoleDefinition) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// Gets the roleAssignments property value. List of Role assignments for this role definition.
func (m *RoleDefinition) GetRoleAssignments()([]RoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// Gets the rolePermissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
func (m *RoleDefinition) GetRolePermissions()([]RolePermission) {
    if m == nil {
        return nil
    } else {
        return m.rolePermissions
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the description property value. Description of the Role definition.
// Parameters:
//  - value : Value to set for the description property.
func (m *RoleDefinition) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Display Name of the Role definition.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *RoleDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isBuiltIn property value. Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
// Parameters:
//  - value : Value to set for the isBuiltIn property.
func (m *RoleDefinition) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// Sets the roleAssignments property value. List of Role assignments for this role definition.
// Parameters:
//  - value : Value to set for the roleAssignments property.
func (m *RoleDefinition) SetRoleAssignments(value []RoleAssignment)() {
    m.roleAssignments = value
}
// Sets the rolePermissions property value. List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
// Parameters:
//  - value : Value to set for the rolePermissions property.
func (m *RoleDefinition) SetRolePermissions(value []RolePermission)() {
    m.rolePermissions = value
}
