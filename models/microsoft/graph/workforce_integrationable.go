package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkforceIntegrationable 
type WorkforceIntegrationable interface {
    ChangeTrackedEntityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApiVersion()(*int32)
    GetDisplayName()(*string)
    GetEncryption()(WorkforceIntegrationEncryptionable)
    GetIsActive()(*bool)
    GetSupportedEntities()(*WorkforceIntegrationSupportedEntities)
    GetUrl()(*string)
    SetApiVersion(value *int32)()
    SetDisplayName(value *string)()
    SetEncryption(value WorkforceIntegrationEncryptionable)()
    SetIsActive(value *bool)()
    SetSupportedEntities(value *WorkforceIntegrationSupportedEntities)()
    SetUrl(value *string)()
}
