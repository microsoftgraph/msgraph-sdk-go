package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ScheduleItemable 
type ScheduleItemable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetEnd()(DateTimeTimeZoneable)
    GetIsPrivate()(*bool)
    GetLocation()(*string)
    GetStart()(DateTimeTimeZoneable)
    GetStatus()(*FreeBusyStatus)
    GetSubject()(*string)
    SetEnd(value DateTimeTimeZoneable)()
    SetIsPrivate(value *bool)()
    SetLocation(value *string)()
    SetStart(value DateTimeTimeZoneable)()
    SetStatus(value *FreeBusyStatus)()
    SetSubject(value *string)()
}
