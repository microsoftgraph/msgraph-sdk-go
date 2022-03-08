package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemActivityStatable 
type ItemActivityStatable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccess()(ItemActionStatable)
    GetActivities()([]ItemActivityable)
    GetCreate()(ItemActionStatable)
    GetDelete()(ItemActionStatable)
    GetEdit()(ItemActionStatable)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIncompleteData()(IncompleteDataable)
    GetIsTrending()(*bool)
    GetMove()(ItemActionStatable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccess(value ItemActionStatable)()
    SetActivities(value []ItemActivityable)()
    SetCreate(value ItemActionStatable)()
    SetDelete(value ItemActionStatable)()
    SetEdit(value ItemActionStatable)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIncompleteData(value IncompleteDataable)()
    SetIsTrending(value *bool)()
    SetMove(value ItemActionStatable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
