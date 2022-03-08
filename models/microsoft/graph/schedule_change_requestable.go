package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ScheduleChangeRequestable 
type ScheduleChangeRequestable interface {
    ChangeTrackedEntityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignedTo()(*ScheduleChangeRequestActor)
    GetManagerActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagerActionMessage()(*string)
    GetManagerUserId()(*string)
    GetSenderDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSenderMessage()(*string)
    GetSenderUserId()(*string)
    GetState()(*ScheduleChangeState)
    SetAssignedTo(value *ScheduleChangeRequestActor)()
    SetManagerActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagerActionMessage(value *string)()
    SetManagerUserId(value *string)()
    SetSenderDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSenderMessage(value *string)()
    SetSenderUserId(value *string)()
    SetState(value *ScheduleChangeState)()
}
