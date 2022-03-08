package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Attendeeable 
type Attendeeable interface {
    AttendeeBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetProposedNewTime()(TimeSlotable)
    GetStatus()(ResponseStatusable)
    SetProposedNewTime(value TimeSlotable)()
    SetStatus(value ResponseStatusable)()
}
