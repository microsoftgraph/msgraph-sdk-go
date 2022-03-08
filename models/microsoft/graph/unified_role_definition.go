package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnifiedRoleDefinition provides operations to manage the roleManagement singleton.
type UnifiedRoleDefinition struct {
    Entity
    // The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
    description *string;
    // The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq, in).
    displayName *string;
    // Read-only collection of role definitions that the given role definition inherits from. Only Azure AD built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
    inheritsPermissionsFrom []UnifiedRoleDefinitionable;
    // Flag indicating whether the role definition is part of the default set included in Azure Active Directory (Azure AD) or a custom definition. Read-only. Supports $filter (eq, in).
    isBuiltIn *bool;
    // Flag indicating whether the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
    isEnabled *bool;
    // List of the scopes or permissions the role definition applies to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
    resourceScopes []string;
    // List of permissions included in the role. Read-only when isBuiltIn is true. Required.
    rolePermissions []UnifiedRolePermissionable;
    // Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true. This identifier is typically used if one needs an identifier to be the same across different directories.
    templateId *string;
    // Indicates version of the role definition. Read-only when isBuiltIn is true.
    version *string;
}
// NewUnifiedRoleDefinition instantiates a new unifiedRoleDefinition and sets the default values.
func NewUnifiedRoleDefinition()(*UnifiedRoleDefinition) {
    m := &UnifiedRoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleDefinitionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUnifiedRoleDefinition(), nil
}
// GetDescription gets the description property value. The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq, in).
func (m *UnifiedRoleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["inheritsPermissionsFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleDefinitionable)
            }
            m.SetInheritsPermissionsFrom(res)
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
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["resourceScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResourceScopes(res)
        }
        return nil
    }
    res["rolePermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRolePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRolePermissionable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRolePermissionable)
            }
            m.SetRolePermissions(res)
        }
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetInheritsPermissionsFrom gets the inheritsPermissionsFrom property value. Read-only collection of role definitions that the given role definition inherits from. Only Azure AD built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
func (m *UnifiedRoleDefinition) GetInheritsPermissionsFrom()([]UnifiedRoleDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.inheritsPermissionsFrom
    }
}
// GetIsBuiltIn gets the isBuiltIn property value. Flag indicating whether the role definition is part of the default set included in Azure Active Directory (Azure AD) or a custom definition. Read-only. Supports $filter (eq, in).
func (m *UnifiedRoleDefinition) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// GetIsEnabled gets the isEnabled property value. Flag indicating whether the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetResourceScopes gets the resourceScopes property value. List of the scopes or permissions the role definition applies to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
func (m *UnifiedRoleDefinition) GetResourceScopes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.resourceScopes
    }
}
// GetRolePermissions gets the rolePermissions property value. List of permissions included in the role. Read-only when isBuiltIn is true. Required.
func (m *UnifiedRoleDefinition) GetRolePermissions()([]UnifiedRolePermissionable) {
    if m == nil {
        return nil
    } else {
        return m.rolePermissions
    }
}
// GetTemplateId gets the templateId property value. Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true. This identifier is typically used if one needs an identifier to be the same across different directories.
func (m *UnifiedRoleDefinition) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
// GetVersion gets the version property value. Indicates version of the role definition. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *UnifiedRoleDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetInheritsPermissionsFrom() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInheritsPermissionsFrom()))
        for i, v := range m.GetInheritsPermissionsFrom() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetResourceScopes() != nil {
        err = writer.WriteCollectionOfStringValues("resourceScopes", m.GetResourceScopes())
        if err != nil {
            return err
        }
    }
    if m.GetRolePermissions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRolePermissions()))
        for i, v := range m.GetRolePermissions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetDescription sets the description property value. The description for the unifiedRoleDefinition. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the unifiedRoleDefinition. Read-only when isBuiltIn is true. Required.  Supports $filter (eq, in).
func (m *UnifiedRoleDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetInheritsPermissionsFrom sets the inheritsPermissionsFrom property value. Read-only collection of role definitions that the given role definition inherits from. Only Azure AD built-in roles (isBuiltIn is true) support this attribute. Supports $expand.
func (m *UnifiedRoleDefinition) SetInheritsPermissionsFrom(value []UnifiedRoleDefinitionable)() {
    if m != nil {
        m.inheritsPermissionsFrom = value
    }
}
// SetIsBuiltIn sets the isBuiltIn property value. Flag indicating whether the role definition is part of the default set included in Azure Active Directory (Azure AD) or a custom definition. Read-only. Supports $filter (eq, in).
func (m *UnifiedRoleDefinition) SetIsBuiltIn(value *bool)() {
    if m != nil {
        m.isBuiltIn = value
    }
}
// SetIsEnabled sets the isEnabled property value. Flag indicating whether the role is enabled for assignment. If false the role is not available for assignment. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetResourceScopes sets the resourceScopes property value. List of the scopes or permissions the role definition applies to. Currently only / is supported. Read-only when isBuiltIn is true. DO NOT USE. This will be deprecated soon. Attach scope to role assignment.
func (m *UnifiedRoleDefinition) SetResourceScopes(value []string)() {
    if m != nil {
        m.resourceScopes = value
    }
}
// SetRolePermissions sets the rolePermissions property value. List of permissions included in the role. Read-only when isBuiltIn is true. Required.
func (m *UnifiedRoleDefinition) SetRolePermissions(value []UnifiedRolePermissionable)() {
    if m != nil {
        m.rolePermissions = value
    }
}
// SetTemplateId sets the templateId property value. Custom template identifier that can be set when isBuiltIn is false but is read-only when isBuiltIn is true. This identifier is typically used if one needs an identifier to be the same across different directories.
func (m *UnifiedRoleDefinition) SetTemplateId(value *string)() {
    if m != nil {
        m.templateId = value
    }
}
// SetVersion sets the version property value. Indicates version of the role definition. Read-only when isBuiltIn is true.
func (m *UnifiedRoleDefinition) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
