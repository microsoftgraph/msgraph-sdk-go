package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ResourceSpecificPermissionGrant struct {
    DirectoryObject
    // ID of the service principal of the Azure AD app that has been granted access. Read-only.
    clientAppId *string;
    // ID of the Azure AD app that has been granted access. Read-only.
    clientId *string;
    // The name of the resource-specific permission. Read-only.
    permission *string;
    // The type of permission. Possible values are: Application, Delegated. Read-only.
    permissionType *string;
    // ID of the Azure AD app that is hosting the resource. Read-only.
    resourceAppId *string;
}
// Instantiates a new resourceSpecificPermissionGrant and sets the default values.
func NewResourceSpecificPermissionGrant()(*ResourceSpecificPermissionGrant) {
    m := &ResourceSpecificPermissionGrant{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the clientAppId property value. ID of the service principal of the Azure AD app that has been granted access. Read-only.
func (m *ResourceSpecificPermissionGrant) GetClientAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientAppId
    }
}
// Gets the clientId property value. ID of the Azure AD app that has been granted access. Read-only.
func (m *ResourceSpecificPermissionGrant) GetClientId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientId
    }
}
// Gets the permission property value. The name of the resource-specific permission. Read-only.
func (m *ResourceSpecificPermissionGrant) GetPermission()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permission
    }
}
// Gets the permissionType property value. The type of permission. Possible values are: Application, Delegated. Read-only.
func (m *ResourceSpecificPermissionGrant) GetPermissionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionType
    }
}
// Gets the resourceAppId property value. ID of the Azure AD app that is hosting the resource. Read-only.
func (m *ResourceSpecificPermissionGrant) GetResourceAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceAppId
    }
}
// The deserialization information for the current model
func (m *ResourceSpecificPermissionGrant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["clientAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAppId(val)
        }
        return nil
    }
    res["clientId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["permission"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val)
        }
        return nil
    }
    res["permissionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionType(val)
        }
        return nil
    }
    res["resourceAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAppId(val)
        }
        return nil
    }
    return res
}
func (m *ResourceSpecificPermissionGrant) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the clientAppId property value. ID of the service principal of the Azure AD app that has been granted access. Read-only.
// Parameters:
//  - value : Value to set for the clientAppId property.
func (m *ResourceSpecificPermissionGrant) SetClientAppId(value *string)() {
    m.clientAppId = value
}
// Sets the clientId property value. ID of the Azure AD app that has been granted access. Read-only.
// Parameters:
//  - value : Value to set for the clientId property.
func (m *ResourceSpecificPermissionGrant) SetClientId(value *string)() {
    m.clientId = value
}
// Sets the permission property value. The name of the resource-specific permission. Read-only.
// Parameters:
//  - value : Value to set for the permission property.
func (m *ResourceSpecificPermissionGrant) SetPermission(value *string)() {
    m.permission = value
}
// Sets the permissionType property value. The type of permission. Possible values are: Application, Delegated. Read-only.
// Parameters:
//  - value : Value to set for the permissionType property.
func (m *ResourceSpecificPermissionGrant) SetPermissionType(value *string)() {
    m.permissionType = value
}
// Sets the resourceAppId property value. ID of the Azure AD app that is hosting the resource. Read-only.
// Parameters:
//  - value : Value to set for the resourceAppId property.
func (m *ResourceSpecificPermissionGrant) SetResourceAppId(value *string)() {
    m.resourceAppId = value
}
