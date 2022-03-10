package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Reminderable 
type Reminderable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetChangeKey()(*string)
    GetEventEndTime()(DateTimeTimeZoneable)
    GetEventId()(*string)
    GetEventLocation()(Locationable)
    GetEventStartTime()(DateTimeTimeZoneable)
    GetEventSubject()(*string)
    GetEventWebLink()(*string)
    GetReminderFireTime()(DateTimeTimeZoneable)
    SetChangeKey(value *string)()
    SetEventEndTime(value DateTimeTimeZoneable)()
    SetEventId(value *string)()
    SetEventLocation(value Locationable)()
    SetEventStartTime(value DateTimeTimeZoneable)()
    SetEventSubject(value *string)()
    SetEventWebLink(value *string)()
    SetReminderFireTime(value DateTimeTimeZoneable)()
}
