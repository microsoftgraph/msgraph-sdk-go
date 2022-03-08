package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserIdentityable 
type UserIdentityable interface {
    Identityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIpAddress()(*string)
    GetUserPrincipalName()(*string)
    SetIpAddress(value *string)()
    SetUserPrincipalName(value *string)()
}
