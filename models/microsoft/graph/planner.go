package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Planner provides operations to manage the planner singleton.
type Planner struct {
    Entity
    // Read-only. Nullable. Returns a collection of the specified buckets
    buckets []PlannerBucketable;
    // Read-only. Nullable. Returns a collection of the specified plans
    plans []PlannerPlanable;
    // Read-only. Nullable. Returns a collection of the specified tasks
    tasks []PlannerTaskable;
}
// NewPlanner instantiates a new planner and sets the default values.
func NewPlanner()(*Planner) {
    m := &Planner{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlanner(), nil
}
// GetBuckets gets the buckets property value. Read-only. Nullable. Returns a collection of the specified buckets
func (m *Planner) GetBuckets()([]PlannerBucketable) {
    if m == nil {
        return nil
    } else {
        return m.buckets
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Planner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buckets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerBucketFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerBucketable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerBucketable)
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["plans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlanable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerPlanable)
            }
            m.SetPlans(res)
        }
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerTaskable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetPlans gets the plans property value. Read-only. Nullable. Returns a collection of the specified plans
func (m *Planner) GetPlans()([]PlannerPlanable) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// GetTasks gets the tasks property value. Read-only. Nullable. Returns a collection of the specified tasks
func (m *Planner) GetTasks()([]PlannerTaskable) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// Serialize serializes information the current object
func (m *Planner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBuckets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBuckets()))
        for i, v := range m.GetBuckets() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlans() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPlans()))
        for i, v := range m.GetPlans() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("plans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuckets sets the buckets property value. Read-only. Nullable. Returns a collection of the specified buckets
func (m *Planner) SetBuckets(value []PlannerBucketable)() {
    if m != nil {
        m.buckets = value
    }
}
// SetPlans sets the plans property value. Read-only. Nullable. Returns a collection of the specified plans
func (m *Planner) SetPlans(value []PlannerPlanable)() {
    if m != nil {
        m.plans = value
    }
}
// SetTasks sets the tasks property value. Read-only. Nullable. Returns a collection of the specified tasks
func (m *Planner) SetTasks(value []PlannerTaskable)() {
    if m != nil {
        m.tasks = value
    }
}
