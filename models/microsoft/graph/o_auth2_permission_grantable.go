package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OAuth2PermissionGrantable 
type OAuth2PermissionGrantable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetClientId()(*string)
    GetConsentType()(*string)
    GetPrincipalId()(*string)
    GetResourceId()(*string)
    GetScope()(*string)
    SetClientId(value *string)()
    SetConsentType(value *string)()
    SetPrincipalId(value *string)()
    SetResourceId(value *string)()
    SetScope(value *string)()
}
