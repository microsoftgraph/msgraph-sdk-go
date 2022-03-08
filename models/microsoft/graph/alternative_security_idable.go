package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AlternativeSecurityIdable 
type AlternativeSecurityIdable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIdentityProvider()(*string)
    GetKey()([]byte)
    GetType()(*int32)
    SetIdentityProvider(value *string)()
    SetKey(value []byte)()
    SetType(value *int32)()
}
