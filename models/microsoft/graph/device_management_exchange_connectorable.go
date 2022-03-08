package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementExchangeConnectorable 
type DeviceManagementExchangeConnectorable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConnectorServerName()(*string)
    GetExchangeAlias()(*string)
    GetExchangeConnectorType()(*DeviceManagementExchangeConnectorType)
    GetExchangeOrganization()(*string)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPrimarySmtpAddress()(*string)
    GetServerName()(*string)
    GetStatus()(*DeviceManagementExchangeConnectorStatus)
    GetVersion()(*string)
    SetConnectorServerName(value *string)()
    SetExchangeAlias(value *string)()
    SetExchangeConnectorType(value *DeviceManagementExchangeConnectorType)()
    SetExchangeOrganization(value *string)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPrimarySmtpAddress(value *string)()
    SetServerName(value *string)()
    SetStatus(value *DeviceManagementExchangeConnectorStatus)()
    SetVersion(value *string)()
}
