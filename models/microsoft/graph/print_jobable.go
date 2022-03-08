package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintJobable 
type PrintJobable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetConfiguration()(PrintJobConfigurationable)
    GetCreatedBy()(UserIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDocuments()([]PrintDocumentable)
    GetIsFetchable()(*bool)
    GetRedirectedFrom()(*string)
    GetRedirectedTo()(*string)
    GetStatus()(PrintJobStatusable)
    GetTasks()([]PrintTaskable)
    SetConfiguration(value PrintJobConfigurationable)()
    SetCreatedBy(value UserIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDocuments(value []PrintDocumentable)()
    SetIsFetchable(value *bool)()
    SetRedirectedFrom(value *string)()
    SetRedirectedTo(value *string)()
    SetStatus(value PrintJobStatusable)()
    SetTasks(value []PrintTaskable)()
}
