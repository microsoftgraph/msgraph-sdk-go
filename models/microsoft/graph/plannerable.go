package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Plannerable 
type Plannerable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetBuckets()([]PlannerBucketable)
    GetPlans()([]PlannerPlanable)
    GetTasks()([]PlannerTaskable)
    SetBuckets(value []PlannerBucketable)()
    SetPlans(value []PlannerPlanable)()
    SetTasks(value []PlannerTaskable)()
}
