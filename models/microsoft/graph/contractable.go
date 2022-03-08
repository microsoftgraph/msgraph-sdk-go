package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Contractable 
type Contractable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    DirectoryObjectable
    GetContractType()(*string)
    GetCustomerId()(*string)
    GetDefaultDomainName()(*string)
    GetDisplayName()(*string)
    SetContractType(value *string)()
    SetCustomerId(value *string)()
    SetDefaultDomainName(value *string)()
    SetDisplayName(value *string)()
}
