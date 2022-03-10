package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PermissionScopeable 
type PermissionScopeable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdminConsentDescription()(*string)
    GetAdminConsentDisplayName()(*string)
    GetId()(*string)
    GetIsEnabled()(*bool)
    GetOrigin()(*string)
    GetType()(*string)
    GetUserConsentDescription()(*string)
    GetUserConsentDisplayName()(*string)
    GetValue()(*string)
    SetAdminConsentDescription(value *string)()
    SetAdminConsentDisplayName(value *string)()
    SetId(value *string)()
    SetIsEnabled(value *bool)()
    SetOrigin(value *string)()
    SetType(value *string)()
    SetUserConsentDescription(value *string)()
    SetUserConsentDisplayName(value *string)()
    SetValue(value *string)()
}
