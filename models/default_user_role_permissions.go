package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultUserRolePermissions 
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
// NewDefaultUserRolePermissions instantiates a new defaultUserRolePermissions and sets the default values.
func NewDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    m := &DefaultUserRolePermissions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDefaultUserRolePermissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultUserRolePermissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultUserRolePermissions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultUserRolePermissions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedToCreateApps gets the allowedToCreateApps property value. Indicates whether the default user role can create applications.
func (m *DefaultUserRolePermissions) GetAllowedToCreateApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToCreateApps
    }
}
// GetAllowedToCreateSecurityGroups gets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups.
func (m *DefaultUserRolePermissions) GetAllowedToCreateSecurityGroups()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToCreateSecurityGroups
    }
}
// GetAllowedToReadOtherUsers gets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users.
func (m *DefaultUserRolePermissions) GetAllowedToReadOtherUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToReadOtherUsers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultUserRolePermissions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedToCreateApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateApps(val)
        }
        return nil
    }
    res["allowedToCreateSecurityGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToCreateSecurityGroups(val)
        }
        return nil
    }
    res["allowedToReadOtherUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedToReadOtherUsers(val)
        }
        return nil
    }
    res["permissionGrantPoliciesAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPermissionGrantPoliciesAssigned(res)
        }
        return nil
    }
    return res
}
// GetPermissionGrantPoliciesAssigned gets the permissionGrantPoliciesAssigned property value. Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
func (m *DefaultUserRolePermissions) GetPermissionGrantPoliciesAssigned()([]string) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPoliciesAssigned
    }
}
// Serialize serializes information the current object
func (m *DefaultUserRolePermissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetPermissionGrantPoliciesAssigned() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultUserRolePermissions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowedToCreateApps sets the allowedToCreateApps property value. Indicates whether the default user role can create applications.
func (m *DefaultUserRolePermissions) SetAllowedToCreateApps(value *bool)() {
    if m != nil {
        m.allowedToCreateApps = value
    }
}
// SetAllowedToCreateSecurityGroups sets the allowedToCreateSecurityGroups property value. Indicates whether the default user role can create security groups.
func (m *DefaultUserRolePermissions) SetAllowedToCreateSecurityGroups(value *bool)() {
    if m != nil {
        m.allowedToCreateSecurityGroups = value
    }
}
// SetAllowedToReadOtherUsers sets the allowedToReadOtherUsers property value. Indicates whether the default user role can read other users.
func (m *DefaultUserRolePermissions) SetAllowedToReadOtherUsers(value *bool)() {
    if m != nil {
        m.allowedToReadOtherUsers = value
    }
}
// SetPermissionGrantPoliciesAssigned sets the permissionGrantPoliciesAssigned property value. Indicates if user consent to apps is allowed, and if it is, which permission to grant consent and which app consent policy (permissionGrantPolicy) govern the permission for users to grant consent. Value should be in the format managePermissionGrantsForSelf.{id}, where {id} is the id of a built-in or custom app consent policy. An empty list indicates user consent to apps is disabled.
func (m *DefaultUserRolePermissions) SetPermissionGrantPoliciesAssigned(value []string)() {
    if m != nil {
        m.permissionGrantPoliciesAssigned = value
    }
}
