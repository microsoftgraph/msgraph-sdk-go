package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Shiftable 
type Shiftable interface {
    ChangeTrackedEntityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDraftShift()(ShiftItemable)
    GetSchedulingGroupId()(*string)
    GetSharedShift()(ShiftItemable)
    GetUserId()(*string)
    SetDraftShift(value ShiftItemable)()
    SetSchedulingGroupId(value *string)()
    SetSharedShift(value ShiftItemable)()
    SetUserId(value *string)()
}
