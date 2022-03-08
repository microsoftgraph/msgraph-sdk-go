package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisionedIdentityable 
type ProvisionedIdentityable interface {
    Identityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDetails()(DetailsInfoable)
    GetIdentityType()(*string)
    SetDetails(value DetailsInfoable)()
    SetIdentityType(value *string)()
}
