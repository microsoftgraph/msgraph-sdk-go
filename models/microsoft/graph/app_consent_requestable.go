package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppConsentRequestable 
type AppConsentRequestable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppDisplayName()(*string)
    GetAppId()(*string)
    GetPendingScopes()([]AppConsentRequestScopeable)
    GetUserConsentRequests()([]UserConsentRequestable)
    SetAppDisplayName(value *string)()
    SetAppId(value *string)()
    SetPendingScopes(value []AppConsentRequestScopeable)()
    SetUserConsentRequests(value []UserConsentRequestable)()
}
