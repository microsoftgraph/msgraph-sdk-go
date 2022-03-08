package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PermissionGrantConditionSetable 
type PermissionGrantConditionSetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetClientApplicationIds()([]string)
    GetClientApplicationPublisherIds()([]string)
    GetClientApplicationsFromVerifiedPublisherOnly()(*bool)
    GetClientApplicationTenantIds()([]string)
    GetPermissionClassification()(*string)
    GetPermissions()([]string)
    GetPermissionType()(*PermissionType)
    GetResourceApplication()(*string)
    SetClientApplicationIds(value []string)()
    SetClientApplicationPublisherIds(value []string)()
    SetClientApplicationsFromVerifiedPublisherOnly(value *bool)()
    SetClientApplicationTenantIds(value []string)()
    SetPermissionClassification(value *string)()
    SetPermissions(value []string)()
    SetPermissionType(value *PermissionType)()
    SetResourceApplication(value *string)()
}
