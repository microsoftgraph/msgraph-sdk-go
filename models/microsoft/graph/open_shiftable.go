package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OpenShiftable 
type OpenShiftable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    ChangeTrackedEntityable
    GetDraftOpenShift()(OpenShiftItemable)
    GetSchedulingGroupId()(*string)
    GetSharedOpenShift()(OpenShiftItemable)
    SetDraftOpenShift(value OpenShiftItemable)()
    SetSchedulingGroupId(value *string)()
    SetSharedOpenShift(value OpenShiftItemable)()
}
