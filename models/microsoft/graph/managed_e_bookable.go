package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedEBookable 
type ManagedEBookable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAssignments()([]ManagedEBookAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceStates()([]DeviceInstallStateable)
    GetDisplayName()(*string)
    GetInformationUrl()(*string)
    GetInstallSummary()(EBookInstallSummaryable)
    GetLargeCover()(MimeContentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPrivacyInformationUrl()(*string)
    GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPublisher()(*string)
    GetUserStateSummary()([]UserInstallStateSummaryable)
    SetAssignments(value []ManagedEBookAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceStates(value []DeviceInstallStateable)()
    SetDisplayName(value *string)()
    SetInformationUrl(value *string)()
    SetInstallSummary(value EBookInstallSummaryable)()
    SetLargeCover(value MimeContentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPrivacyInformationUrl(value *string)()
    SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPublisher(value *string)()
    SetUserStateSummary(value []UserInstallStateSummaryable)()
}
