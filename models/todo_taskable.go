package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TodoTaskable 
type TodoTaskable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
