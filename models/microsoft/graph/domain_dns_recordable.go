package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DomainDnsRecordable 
type DomainDnsRecordable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetIsOptional()(*bool)
    GetLabel()(*string)
    GetRecordType()(*string)
    GetSupportedService()(*string)
    GetTtl()(*int32)
    SetIsOptional(value *bool)()
    SetLabel(value *string)()
    SetRecordType(value *string)()
    SetSupportedService(value *string)()
    SetTtl(value *int32)()
}
