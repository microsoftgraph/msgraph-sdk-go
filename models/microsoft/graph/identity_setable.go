package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentitySetable 
type IdentitySetable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplication()(Identityable)
    GetDevice()(Identityable)
    GetUser()(Identityable)
    SetApplication(value Identityable)()
    SetDevice(value Identityable)()
    SetUser(value Identityable)()
}
