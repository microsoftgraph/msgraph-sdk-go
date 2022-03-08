package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HostSecurityStateable 
type HostSecurityStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetFqdn()(*string)
    GetIsAzureAdJoined()(*bool)
    GetIsAzureAdRegistered()(*bool)
    GetIsHybridAzureDomainJoined()(*bool)
    GetNetBiosName()(*string)
    GetOs()(*string)
    GetPrivateIpAddress()(*string)
    GetPublicIpAddress()(*string)
    GetRiskScore()(*string)
    SetFqdn(value *string)()
    SetIsAzureAdJoined(value *bool)()
    SetIsAzureAdRegistered(value *bool)()
    SetIsHybridAzureDomainJoined(value *bool)()
    SetNetBiosName(value *string)()
    SetOs(value *string)()
    SetPrivateIpAddress(value *string)()
    SetPublicIpAddress(value *string)()
    SetRiskScore(value *string)()
}
