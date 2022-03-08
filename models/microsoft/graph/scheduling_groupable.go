package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SchedulingGroupable 
type SchedulingGroupable interface {
    ChangeTrackedEntityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetIsActive()(*bool)
    GetUserIds()([]string)
    SetDisplayName(value *string)()
    SetIsActive(value *bool)()
    SetUserIds(value []string)()
}
