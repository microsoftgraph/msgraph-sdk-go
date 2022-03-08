package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityProviderable 
type IdentityProviderable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetName()(*string)
    GetType()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetName(value *string)()
    SetType(value *string)()
}
