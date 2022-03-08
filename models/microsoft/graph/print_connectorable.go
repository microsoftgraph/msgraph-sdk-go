package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintConnectorable 
type PrintConnectorable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAppVersion()(*string)
    GetDisplayName()(*string)
    GetFullyQualifiedDomainName()(*string)
    GetLocation()(PrinterLocationable)
    GetOperatingSystem()(*string)
    GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAppVersion(value *string)()
    SetDisplayName(value *string)()
    SetFullyQualifiedDomainName(value *string)()
    SetLocation(value PrinterLocationable)()
    SetOperatingSystem(value *string)()
    SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
