package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfferShiftRequestable 
type OfferShiftRequestable interface {
    ScheduleChangeRequestable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetRecipientActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipientActionMessage()(*string)
    GetRecipientUserId()(*string)
    GetSenderShiftId()(*string)
    SetRecipientActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipientActionMessage(value *string)()
    SetRecipientUserId(value *string)()
    SetSenderShiftId(value *string)()
}
