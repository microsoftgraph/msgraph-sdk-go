package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ShiftItemable 
type ShiftItemable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    ScheduleEntityable
    GetActivities()([]ShiftActivityable)
    GetDisplayName()(*string)
    GetNotes()(*string)
    SetActivities(value []ShiftActivityable)()
    SetDisplayName(value *string)()
    SetNotes(value *string)()
}
