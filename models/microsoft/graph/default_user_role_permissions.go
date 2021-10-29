package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DefaultUserRolePermissions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the default user role can create applications.
    allowedToCreateApps *bool;
    // Indicates whether the default user role can create security groups.
    allowedToCreateSecurityGroups *bool;
    // Indicates whether the default user role can read other users.
    allowedToReadOtherUsers *bool;
    // Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
    permissionGrantPoliciesAssigned []string;
}
// Instantiates a new defaultUserRolePermissions and sets the default values.
func NewDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    m := &DefaultUserRolePermissions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultUserRolePermissions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowedToCreateApps property value. Indicates whether the default user role can create applications.
func (m *DefaultUserRolePermissions) GetAllowedToCreateApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToCreateApps
    }
}
// Gets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups.
func (m *DefaultUserRolePermissions) GetAllowedToCreateSecurityGroups()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToCreateSecurityGroups
    }
}
// Gets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users.
func (m *DefaultUserRolePermissions) GetAllowedToReadOtherUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToReadOtherUsers
    }
}
// Gets the permissionGrantPoliciesAssigned property value. Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
func (m *DefaultUserRolePermissions) GetPermissionGrantPoliciesAssigned()([]string) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPoliciesAssigned
    }
}
// The deserialization information for the current model
func (m *DefaultUserRolePermissions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedToCreateApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowedToCreateApps(val)
        return nil
    }
    res["allowedToCreateSecurityGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowedToCreateSecurityGroups(val)
        return nil
    }
    res["allowedToReadOtherUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowedToReadOtherUsers(val)
        return nil
    }
    res["permissionGrantPoliciesAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetPermissionGrantPoliciesAssigned(res)
        return nil
    }
    return res
}
func (m *DefaultUserRolePermissions) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DefaultUserRolePermissions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowedToCreateApps", m.GetAllowedToCreateApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToCreateSecurityGroups", m.GetAllowedToCreateSecurityGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowedToReadOtherUsers", m.GetAllowedToReadOtherUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("permissionGrantPoliciesAssigned", m.GetPermissionGrantPoliciesAssigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DefaultUserRolePermissions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowedToCreateApps property value. Indicates whether the default user role can create applications.
// Parameters:
//  - value : Value to set for the allowedToCreateApps property.
func (m *DefaultUserRolePermissions) SetAllowedToCreateApps(value *bool)() {
    m.allowedToCreateApps = value
}
// Sets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups.
// Parameters:
//  - value : Value to set for the allowedToCreateSecurityGroups property.
func (m *DefaultUserRolePermissions) SetAllowedToCreateSecurityGroups(value *bool)() {
    m.allowedToCreateSecurityGroups = value
}
// Sets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users.
// Parameters:
//  - value : Value to set for the allowedToReadOtherUsers property.
func (m *DefaultUserRolePermissions) SetAllowedToReadOtherUsers(value *bool)() {
    m.allowedToReadOtherUsers = value
}
// Sets the permissionGrantPoliciesAssigned property value. Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
// Parameters:
//  - value : Value to set for the permissionGrantPoliciesAssigned property.
func (m *DefaultUserRolePermissions) SetPermissionGrantPoliciesAssigned(value []string)() {
    m.permissionGrantPoliciesAssigned = value
}
