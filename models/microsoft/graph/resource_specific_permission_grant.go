package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ResourceSpecificPermissionGrant struct {
    DirectoryObject
    clientAppId *string;
    clientId *string;
    permission *string;
    permissionType *string;
    resourceAppId *string;
}
func NewResourceSpecificPermissionGrant()(*ResourceSpecificPermissionGrant) {
    m := &ResourceSpecificPermissionGrant{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
func (m *ResourceSpecificPermissionGrant) GetClientAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientAppId
    }
}
func (m *ResourceSpecificPermissionGrant) GetClientId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientId
    }
}
func (m *ResourceSpecificPermissionGrant) GetPermission()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permission
    }
}
func (m *ResourceSpecificPermissionGrant) GetPermissionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionType
    }
}
func (m *ResourceSpecificPermissionGrant) GetResourceAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceAppId
    }
}
func (m *ResourceSpecificPermissionGrant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["clientAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientAppId(val)
        return nil
    }
    res["clientId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientId(val)
        return nil
    }
    res["permission"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPermission(val)
        return nil
    }
    res["permissionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPermissionType(val)
        return nil
    }
    res["resourceAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceAppId(val)
        return nil
    }
    return res
}
func (m *ResourceSpecificPermissionGrant) IsNil()(bool) {
    return m == nil
}
func (m *ResourceSpecificPermissionGrant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientAppId", m.GetClientAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permission", m.GetPermission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionType", m.GetPermissionType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceAppId", m.GetResourceAppId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ResourceSpecificPermissionGrant) SetClientAppId(value *string)() {
    m.clientAppId = value
}
func (m *ResourceSpecificPermissionGrant) SetClientId(value *string)() {
    m.clientId = value
}
func (m *ResourceSpecificPermissionGrant) SetPermission(value *string)() {
    m.permission = value
}
func (m *ResourceSpecificPermissionGrant) SetPermissionType(value *string)() {
    m.permissionType = value
}
func (m *ResourceSpecificPermissionGrant) SetResourceAppId(value *string)() {
    m.resourceAppId = value
}
