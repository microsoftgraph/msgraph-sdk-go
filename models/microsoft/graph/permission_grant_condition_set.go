package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PermissionGrantConditionSet struct {
    Entity
    // A list of appId values for the client applications to match with, or a list with the single value all to match any client application. Default is the single value all.
    clientApplicationIds []string;
    // A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the single value all to match with client apps from any publisher. Default is the single value all.
    clientApplicationPublisherIds []string;
    // Set to true to only match on client applications with a verified publisher. Set to false to match on any client app, even if it does not have a verified publisher. Default is false.
    clientApplicationsFromVerifiedPublisherOnly *bool;
    // A list of Azure Active Directory tenant IDs in which the client application is registered, or a list with the single value all to match with client apps registered in any tenant. Default is the single value all.
    clientApplicationTenantIds []string;
    // The permission classification for the permission being granted, or all to match with any permission classification (including permissions which are not classified). Default is all.
    permissionClassification *string;
    // The list of id values for the specific permissions to match with, or a list with the single value all to match with any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API's **servicePrincipal** object. The id of application permissions can be found in the appRoles property of the API's **servicePrincipal** object. The id of resource-specific application permissions can be found in the resourceSpecificApplicationPermissions property of the API's **servicePrincipal** object. Default is the single value all.
    permissions []string;
    // The permission type of the permission being granted. Possible values: application for application permissions (e.g. app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated permissions which have not been configured by the API publisher to require admin consent—this value may be used in built-in permission grant policies, but cannot be used in custom permission grant policies. Required.
    permissionType *PermissionType;
    // The appId of the resource application (e.g. the API) for which a permission is being granted, or any to match with any resource application or API. Default is any.
    resourceApplication *string;
}
// Instantiates a new permissionGrantConditionSet and sets the default values.
func NewPermissionGrantConditionSet()(*PermissionGrantConditionSet) {
    m := &PermissionGrantConditionSet{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the clientApplicationIds property value. A list of appId values for the client applications to match with, or a list with the single value all to match any client application. Default is the single value all.
func (m *PermissionGrantConditionSet) GetClientApplicationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.clientApplicationIds
    }
}
// Gets the clientApplicationPublisherIds property value. A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the single value all to match with client apps from any publisher. Default is the single value all.
func (m *PermissionGrantConditionSet) GetClientApplicationPublisherIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.clientApplicationPublisherIds
    }
}
// Gets the clientApplicationsFromVerifiedPublisherOnly property value. Set to true to only match on client applications with a verified publisher. Set to false to match on any client app, even if it does not have a verified publisher. Default is false.
func (m *PermissionGrantConditionSet) GetClientApplicationsFromVerifiedPublisherOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.clientApplicationsFromVerifiedPublisherOnly
    }
}
// Gets the clientApplicationTenantIds property value. A list of Azure Active Directory tenant IDs in which the client application is registered, or a list with the single value all to match with client apps registered in any tenant. Default is the single value all.
func (m *PermissionGrantConditionSet) GetClientApplicationTenantIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.clientApplicationTenantIds
    }
}
// Gets the permissionClassification property value. The permission classification for the permission being granted, or all to match with any permission classification (including permissions which are not classified). Default is all.
func (m *PermissionGrantConditionSet) GetPermissionClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionClassification
    }
}
// Gets the permissions property value. The list of id values for the specific permissions to match with, or a list with the single value all to match with any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API's **servicePrincipal** object. The id of application permissions can be found in the appRoles property of the API's **servicePrincipal** object. The id of resource-specific application permissions can be found in the resourceSpecificApplicationPermissions property of the API's **servicePrincipal** object. Default is the single value all.
func (m *PermissionGrantConditionSet) GetPermissions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.permissions
    }
}
// Gets the permissionType property value. The permission type of the permission being granted. Possible values: application for application permissions (e.g. app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated permissions which have not been configured by the API publisher to require admin consent—this value may be used in built-in permission grant policies, but cannot be used in custom permission grant policies. Required.
func (m *PermissionGrantConditionSet) GetPermissionType()(*PermissionType) {
    if m == nil {
        return nil
    } else {
        return m.permissionType
    }
}
// Gets the resourceApplication property value. The appId of the resource application (e.g. the API) for which a permission is being granted, or any to match with any resource application or API. Default is any.
func (m *PermissionGrantConditionSet) GetResourceApplication()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceApplication
    }
}
// The deserialization information for the current model
func (m *PermissionGrantConditionSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientApplicationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetClientApplicationIds(res)
        return nil
    }
    res["clientApplicationPublisherIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetClientApplicationPublisherIds(res)
        return nil
    }
    res["clientApplicationsFromVerifiedPublisherOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetClientApplicationsFromVerifiedPublisherOnly(val)
        return nil
    }
    res["clientApplicationTenantIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetClientApplicationTenantIds(res)
        return nil
    }
    res["permissionClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPermissionClassification(val)
        return nil
    }
    res["permissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetPermissions(res)
        return nil
    }
    res["permissionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionType)
        if err != nil {
            return err
        }
        cast := val.(PermissionType)
        m.SetPermissionType(&cast)
        return nil
    }
    res["resourceApplication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceApplication(val)
        return nil
    }
    return res
}
func (m *PermissionGrantConditionSet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PermissionGrantConditionSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("clientApplicationIds", m.GetClientApplicationIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("clientApplicationPublisherIds", m.GetClientApplicationPublisherIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("clientApplicationsFromVerifiedPublisherOnly", m.GetClientApplicationsFromVerifiedPublisherOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("clientApplicationTenantIds", m.GetClientApplicationTenantIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionClassification", m.GetPermissionClassification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionType() != nil {
        cast := m.GetPermissionType().String()
        err = writer.WriteStringValue("permissionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceApplication", m.GetResourceApplication())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the clientApplicationIds property value. A list of appId values for the client applications to match with, or a list with the single value all to match any client application. Default is the single value all.
// Parameters:
//  - value : Value to set for the clientApplicationIds property.
func (m *PermissionGrantConditionSet) SetClientApplicationIds(value []string)() {
    m.clientApplicationIds = value
}
// Sets the clientApplicationPublisherIds property value. A list of Microsoft Partner Network (MPN) IDs for verified publishers of the client application, or a list with the single value all to match with client apps from any publisher. Default is the single value all.
// Parameters:
//  - value : Value to set for the clientApplicationPublisherIds property.
func (m *PermissionGrantConditionSet) SetClientApplicationPublisherIds(value []string)() {
    m.clientApplicationPublisherIds = value
}
// Sets the clientApplicationsFromVerifiedPublisherOnly property value. Set to true to only match on client applications with a verified publisher. Set to false to match on any client app, even if it does not have a verified publisher. Default is false.
// Parameters:
//  - value : Value to set for the clientApplicationsFromVerifiedPublisherOnly property.
func (m *PermissionGrantConditionSet) SetClientApplicationsFromVerifiedPublisherOnly(value *bool)() {
    m.clientApplicationsFromVerifiedPublisherOnly = value
}
// Sets the clientApplicationTenantIds property value. A list of Azure Active Directory tenant IDs in which the client application is registered, or a list with the single value all to match with client apps registered in any tenant. Default is the single value all.
// Parameters:
//  - value : Value to set for the clientApplicationTenantIds property.
func (m *PermissionGrantConditionSet) SetClientApplicationTenantIds(value []string)() {
    m.clientApplicationTenantIds = value
}
// Sets the permissionClassification property value. The permission classification for the permission being granted, or all to match with any permission classification (including permissions which are not classified). Default is all.
// Parameters:
//  - value : Value to set for the permissionClassification property.
func (m *PermissionGrantConditionSet) SetPermissionClassification(value *string)() {
    m.permissionClassification = value
}
// Sets the permissions property value. The list of id values for the specific permissions to match with, or a list with the single value all to match with any permission. The id of delegated permissions can be found in the oauth2PermissionScopes property of the API's **servicePrincipal** object. The id of application permissions can be found in the appRoles property of the API's **servicePrincipal** object. The id of resource-specific application permissions can be found in the resourceSpecificApplicationPermissions property of the API's **servicePrincipal** object. Default is the single value all.
// Parameters:
//  - value : Value to set for the permissions property.
func (m *PermissionGrantConditionSet) SetPermissions(value []string)() {
    m.permissions = value
}
// Sets the permissionType property value. The permission type of the permission being granted. Possible values: application for application permissions (e.g. app roles), or delegated for delegated permissions. The value delegatedUserConsentable indicates delegated permissions which have not been configured by the API publisher to require admin consent—this value may be used in built-in permission grant policies, but cannot be used in custom permission grant policies. Required.
// Parameters:
//  - value : Value to set for the permissionType property.
func (m *PermissionGrantConditionSet) SetPermissionType(value *PermissionType)() {
    m.permissionType = value
}
// Sets the resourceApplication property value. The appId of the resource application (e.g. the API) for which a permission is being granted, or any to match with any resource application or API. Default is any.
// Parameters:
//  - value : Value to set for the resourceApplication property.
func (m *PermissionGrantConditionSet) SetResourceApplication(value *string)() {
    m.resourceApplication = value
}
