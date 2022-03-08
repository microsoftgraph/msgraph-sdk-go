package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementExportJobable 
type DeviceManagementExportJobable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFilter()(*string)
    GetFormat()(*DeviceManagementReportFileFormat)
    GetLocalizationType()(*DeviceManagementExportJobLocalizationType)
    GetReportName()(*string)
    GetRequestDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSelect()([]string)
    GetSnapshotId()(*string)
    GetStatus()(*DeviceManagementReportStatus)
    GetUrl()(*string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFilter(value *string)()
    SetFormat(value *DeviceManagementReportFileFormat)()
    SetLocalizationType(value *DeviceManagementExportJobLocalizationType)()
    SetReportName(value *string)()
    SetRequestDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSelect(value []string)()
    SetSnapshotId(value *string)()
    SetStatus(value *DeviceManagementReportStatus)()
    SetUrl(value *string)()
}
