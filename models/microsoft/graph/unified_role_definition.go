package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UnifiedRoleDefinition struct {
    Entity
    // The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
    description *string;
    // The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq, in).
    displayName *string;
    // Read-only collection of role definitions that the given role definition inherits from. Only Azure AD built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
    inheritsPermissionsFrom []UnifiedRoleDefinition;
    // Flag indicating whether the role definition is part of the default set included in Azure Active Directory (Azure AD) or a custom definition. Read-only. Supports $filter (eq, in).
    isBuiltIn *bool;
    // Flag indicating whether the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
    isEnabled *bool;
    // List of the scopes or permissions the role definition applies to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
    resourceScopes []string;
    // List of permissions included in the role. Read-only when isBuiltIn is true. Required.
    rolePermissions []UnifiedRolePermission;
    // Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true. This identifier is typically used if one needs an identifier to be the same across different directories.
    templateId *string;
    // Indicates version of the role definition. Read-only when isBuiltIn is true.
    version *string;
}
// Instantiates a new unifiedRoleDefinition and sets the default values.
func NewUnifiedRoleDefinition()(*UnifiedRoleDefinition) {
    m := &UnifiedRoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq, in).
func (m *UnifiedRoleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the inheritsPermissionsFrom property value. Read-only collection of role definitions that the given role definition inherits from. Only Azure AD built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
func (m *UnifiedRoleDefinition) GetInheritsPermissionsFrom()([]UnifiedRoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.inheritsPermissionsFrom
    }
}
// Gets the isBuiltIn property value. Flag indicating whether the role definition is part of the default set included in Azure Active Directory (Azure AD) or a custom definition. Read-only. Supports $filter (eq, in).
func (m *UnifiedRoleDefinition) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// Gets the isEnabled property value. Flag indicating whether the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// Gets the resourceScopes property value. List of the scopes or permissions the role definition applies to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
func (m *UnifiedRoleDefinition) GetResourceScopes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.resourceScopes
    }
}
// Gets the rolePermissions property value. List of permissions included in the role. Read-only when isBuiltIn is true. Required.
func (m *UnifiedRoleDefinition) GetRolePermissions()([]UnifiedRolePermission) {
    if m == nil {
        return nil
    } else {
        return m.rolePermissions
    }
}
// Gets the templateId property value. Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true. This identifier is typically used if one needs an identifier to be the same across different directories.
func (m *UnifiedRoleDefinition) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// Gets the version property value. Indicates version of the role definition. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *UnifiedRoleDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["inheritsPermissionsFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRoleDefinition() })
        if err != nil {
            return err
        }
        res := make([]UnifiedRoleDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedRoleDefinition))
        }
        m.SetInheritsPermissionsFrom(res)
        return nil
    }
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsBuiltIn(val)
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabled(val)
        return nil
    }
    res["resourceScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetResourceScopes(res)
        return nil
    }
    res["rolePermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnifiedRolePermission() })
        if err != nil {
            return err
        }
        res := make([]UnifiedRolePermission, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnifiedRolePermission))
        }
        m.SetRolePermissions(res)
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTemplateId(val)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *UnifiedRoleDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UnifiedRoleDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInheritsPermissionsFrom()))
        for i, v := range m.GetInheritsPermissionsFrom() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("inheritsPermissionsFrom", cast)
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
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("resourceScopes", m.GetResourceScopes())
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
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
// Parameters:
//  - value : Value to set for the description property.
func (m *UnifiedRoleDefinition) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq, in).
// Parameters:
//  - value : Value to set for the displayName property.
func (m *UnifiedRoleDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the inheritsPermissionsFrom property value. Read-only collection of role definitions that the given role definition inherits from. Only Azure AD built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
// Parameters:
//  - value : Value to set for the inheritsPermissionsFrom property.
func (m *UnifiedRoleDefinition) SetInheritsPermissionsFrom(value []UnifiedRoleDefinition)() {
    m.inheritsPermissionsFrom = value
}
// Sets the isBuiltIn property value. Flag indicating whether the role definition is part of the default set included in Azure Active Directory (Azure AD) or a custom definition. Read-only. Supports $filter (eq, in).
// Parameters:
//  - value : Value to set for the isBuiltIn property.
func (m *UnifiedRoleDefinition) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// Sets the isEnabled property value. Flag indicating whether the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
// Parameters:
//  - value : Value to set for the isEnabled property.
func (m *UnifiedRoleDefinition) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// Sets the resourceScopes property value. List of the scopes or permissions the role definition applies to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
// Parameters:
//  - value : Value to set for the resourceScopes property.
func (m *UnifiedRoleDefinition) SetResourceScopes(value []string)() {
    m.resourceScopes = value
}
// Sets the rolePermissions property value. List of permissions included in the role. Read-only when isBuiltIn is true. Required.
// Parameters:
//  - value : Value to set for the rolePermissions property.
func (m *UnifiedRoleDefinition) SetRolePermissions(value []UnifiedRolePermission)() {
    m.rolePermissions = value
}
// Sets the templateId property value. Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true. This identifier is typically used if one needs an identifier to be the same across different directories.
// Parameters:
//  - value : Value to set for the templateId property.
func (m *UnifiedRoleDefinition) SetTemplateId(value *string)() {
    m.templateId = value
}
// Sets the version property value. Indicates version of the role definition. Read-only when isBuiltIn is true.
// Parameters:
//  - value : Value to set for the version property.
func (m *UnifiedRoleDefinition) SetVersion(value *string)() {
    m.version = value
}
