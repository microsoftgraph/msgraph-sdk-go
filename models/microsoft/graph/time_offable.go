package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeOffable 
type TimeOffable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    ChangeTrackedEntityable
    GetDraftTimeOff()(TimeOffItemable)
    GetSharedTimeOff()(TimeOffItemable)
    GetUserId()(*string)
    SetDraftTimeOff(value TimeOffItemable)()
    SetSharedTimeOff(value TimeOffItemable)()
    SetUserId(value *string)()
}
