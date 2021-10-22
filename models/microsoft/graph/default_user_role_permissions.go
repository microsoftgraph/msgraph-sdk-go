package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DefaultUserRolePermissions struct {
    additionalData map[string]interface{};
    allowedToCreateApps *bool;
    allowedToCreateSecurityGroups *bool;
    allowedToReadOtherUsers *bool;
    permissionGrantPoliciesAssigned []string;
}
func NewDefaultUserRolePermissions()(*DefaultUserRolePermissions) {
    m := &DefaultUserRolePermissions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DefaultUserRolePermissions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DefaultUserRolePermissions) GetAllowedToCreateApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToCreateApps
    }
}
func (m *DefaultUserRolePermissions) GetAllowedToCreateSecurityGroups()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToCreateSecurityGroups
    }
}
func (m *DefaultUserRolePermissions) GetAllowedToReadOtherUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowedToReadOtherUsers
    }
}
func (m *DefaultUserRolePermissions) GetPermissionGrantPoliciesAssigned()([]string) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrantPoliciesAssigned
    }
}
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
func (m *DefaultUserRolePermissions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DefaultUserRolePermissions) SetAllowedToCreateApps(value *bool)() {
    m.allowedToCreateApps = value
}
func (m *DefaultUserRolePermissions) SetAllowedToCreateSecurityGroups(value *bool)() {
    m.allowedToCreateSecurityGroups = value
}
func (m *DefaultUserRolePermissions) SetAllowedToReadOtherUsers(value *bool)() {
    m.allowedToReadOtherUsers = value
}
func (m *DefaultUserRolePermissions) SetPermissionGrantPoliciesAssigned(value []string)() {
    m.permissionGrantPoliciesAssigned = value
}
