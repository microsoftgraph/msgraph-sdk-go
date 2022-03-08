package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// NetworkConnectionable 
type NetworkConnectionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetApplicationName()(*string)
    GetDestinationAddress()(*string)
    GetDestinationDomain()(*string)
    GetDestinationLocation()(*string)
    GetDestinationPort()(*string)
    GetDestinationUrl()(*string)
    GetDirection()(*ConnectionDirection)
    GetDomainRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocalDnsName()(*string)
    GetNatDestinationAddress()(*string)
    GetNatDestinationPort()(*string)
    GetNatSourceAddress()(*string)
    GetNatSourcePort()(*string)
    GetProtocol()(*SecurityNetworkProtocol)
    GetRiskScore()(*string)
    GetSourceAddress()(*string)
    GetSourceLocation()(*string)
    GetSourcePort()(*string)
    GetStatus()(*ConnectionStatus)
    GetUrlParameters()(*string)
    SetApplicationName(value *string)()
    SetDestinationAddress(value *string)()
    SetDestinationDomain(value *string)()
    SetDestinationLocation(value *string)()
    SetDestinationPort(value *string)()
    SetDestinationUrl(value *string)()
    SetDirection(value *ConnectionDirection)()
    SetDomainRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocalDnsName(value *string)()
    SetNatDestinationAddress(value *string)()
    SetNatDestinationPort(value *string)()
    SetNatSourceAddress(value *string)()
    SetNatSourcePort(value *string)()
    SetProtocol(value *SecurityNetworkProtocol)()
    SetRiskScore(value *string)()
    SetSourceAddress(value *string)()
    SetSourceLocation(value *string)()
    SetSourcePort(value *string)()
    SetStatus(value *ConnectionStatus)()
    SetUrlParameters(value *string)()
}
