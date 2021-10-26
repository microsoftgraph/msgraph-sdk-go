package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AppRoleAssignment struct {
    DirectoryObject
    // The identifier (id) for the app role which is assigned to the principal. This app role must be exposed in the appRoles property on the resource application's service principal (resourceId). If the resource application has not declared any app roles, a default app role ID of 00000000-0000-0000-0000-000000000000 can be specified to signal that the principal is assigned to the resource app without any specific app roles. Required on create.
    appRoleId *string;
    // The time when the app role assignment was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The display name of the user, group, or service principal that was granted the app role assignment. Read-only. Supports $filter (eq and startswith).
    principalDisplayName *string;
    // The unique identifier (id) for the user, group or service principal being granted the app role. Required on create.
    principalId *string;
    // The type of the assigned principal. This can either be User, Group or ServicePrincipal. Read-only.
    principalType *string;
    // The display name of the resource app's service principal to which the assignment is made.
    resourceDisplayName *string;
    // The unique identifier (id) for the resource service principal for which the assignment is made. Required on create. Supports $filter (eq only).
    resourceId *string;
}
// Instantiates a new appRoleAssignment and sets the default values.
func NewAppRoleAssignment()(*AppRoleAssignment) {
    m := &AppRoleAssignment{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the appRoleId property value. The identifier (id) for the app role which is assigned to the principal. This app role must be exposed in the appRoles property on the resource application's service principal (resourceId). If the resource application has not declared any app roles, a default app role ID of 00000000-0000-0000-0000-000000000000 can be specified to signal that the principal is assigned to the resource app without any specific app roles. Required on create.
func (m *AppRoleAssignment) GetAppRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appRoleId
    }
}
// Gets the createdDateTime property value. The time when the app role assignment was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AppRoleAssignment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the principalDisplayName property value. The display name of the user, group, or service principal that was granted the app role assignment. Read-only. Supports $filter (eq and startswith).
func (m *AppRoleAssignment) GetPrincipalDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalDisplayName
    }
}
// Gets the principalId property value. The unique identifier (id) for the user, group or service principal being granted the app role. Required on create.
func (m *AppRoleAssignment) GetPrincipalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalId
    }
}
// Gets the principalType property value. The type of the assigned principal. This can either be User, Group or ServicePrincipal. Read-only.
func (m *AppRoleAssignment) GetPrincipalType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.principalType
    }
}
// Gets the resourceDisplayName property value. The display name of the resource app's service principal to which the assignment is made.
func (m *AppRoleAssignment) GetResourceDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceDisplayName
    }
}
// Gets the resourceId property value. The unique identifier (id) for the resource service principal for which the assignment is made. Required on create. Supports $filter (eq only).
func (m *AppRoleAssignment) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
// The deserialization information for the current model
func (m *AppRoleAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["appRoleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppRoleId(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["principalDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrincipalDisplayName(val)
        return nil
    }
    res["principalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrincipalId(val)
        return nil
    }
    res["principalType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrincipalType(val)
        return nil
    }
    res["resourceDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceDisplayName(val)
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
        return nil
    }
    return res
}
func (m *AppRoleAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AppRoleAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appRoleId", m.GetAppRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalDisplayName", m.GetPrincipalDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalType", m.GetPrincipalType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceDisplayName", m.GetResourceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appRoleId property value. The identifier (id) for the app role which is assigned to the principal. This app role must be exposed in the appRoles property on the resource application's service principal (resourceId). If the resource application has not declared any app roles, a default app role ID of 00000000-0000-0000-0000-000000000000 can be specified to signal that the principal is assigned to the resource app without any specific app roles. Required on create.
// Parameters:
//  - value : Value to set for the appRoleId property.
func (m *AppRoleAssignment) SetAppRoleId(value *string)() {
    m.appRoleId = value
}
// Sets the createdDateTime property value. The time when the app role assignment was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AppRoleAssignment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the principalDisplayName property value. The display name of the user, group, or service principal that was granted the app role assignment. Read-only. Supports $filter (eq and startswith).
// Parameters:
//  - value : Value to set for the principalDisplayName property.
func (m *AppRoleAssignment) SetPrincipalDisplayName(value *string)() {
    m.principalDisplayName = value
}
// Sets the principalId property value. The unique identifier (id) for the user, group or service principal being granted the app role. Required on create.
// Parameters:
//  - value : Value to set for the principalId property.
func (m *AppRoleAssignment) SetPrincipalId(value *string)() {
    m.principalId = value
}
// Sets the principalType property value. The type of the assigned principal. This can either be User, Group or ServicePrincipal. Read-only.
// Parameters:
//  - value : Value to set for the principalType property.
func (m *AppRoleAssignment) SetPrincipalType(value *string)() {
    m.principalType = value
}
// Sets the resourceDisplayName property value. The display name of the resource app's service principal to which the assignment is made.
// Parameters:
//  - value : Value to set for the resourceDisplayName property.
func (m *AppRoleAssignment) SetResourceDisplayName(value *string)() {
    m.resourceDisplayName = value
}
// Sets the resourceId property value. The unique identifier (id) for the resource service principal for which the assignment is made. Required on create. Supports $filter (eq only).
// Parameters:
//  - value : Value to set for the resourceId property.
func (m *AppRoleAssignment) SetResourceId(value *string)() {
    m.resourceId = value
}
