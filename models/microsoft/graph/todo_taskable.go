package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TodoTaskable 
type TodoTaskable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetBody()(ItemBodyable)
    GetBodyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCompletedDateTime()(DateTimeTimeZoneable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDueDateTime()(DateTimeTimeZoneable)
    GetExtensions()([]Extensionable)
    GetImportance()(*Importance)
    GetIsReminderOn()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLinkedResources()([]LinkedResourceable)
    GetRecurrence()(PatternedRecurrenceable)
    GetReminderDateTime()(DateTimeTimeZoneable)
    GetStatus()(*TaskStatus)
    GetTitle()(*string)
    SetBody(value ItemBodyable)()
    SetBodyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCompletedDateTime(value DateTimeTimeZoneable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDueDateTime(value DateTimeTimeZoneable)()
    SetExtensions(value []Extensionable)()
    SetImportance(value *Importance)()
    SetIsReminderOn(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLinkedResources(value []LinkedResourceable)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetReminderDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *TaskStatus)()
    SetTitle(value *string)()
}
