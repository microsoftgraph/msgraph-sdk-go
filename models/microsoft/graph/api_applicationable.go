package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApiApplicationable 
type ApiApplicationable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAcceptMappedClaims()(*bool)
    GetKnownClientApplications()([]string)
    GetOauth2PermissionScopes()([]PermissionScopeable)
    GetPreAuthorizedApplications()([]PreAuthorizedApplicationable)
    GetRequestedAccessTokenVersion()(*int32)
    SetAcceptMappedClaims(value *bool)()
    SetKnownClientApplications(value []string)()
    SetOauth2PermissionScopes(value []PermissionScopeable)()
    SetPreAuthorizedApplications(value []PreAuthorizedApplicationable)()
    SetRequestedAccessTokenVersion(value *int32)()
}
