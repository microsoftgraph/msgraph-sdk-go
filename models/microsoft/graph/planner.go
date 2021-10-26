package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Planner struct {
    Entity
    // Read-only. Nullable. Returns a collection of the specified buckets
    buckets []PlannerBucket;
    // Read-only. Nullable. Returns a collection of the specified plans
    plans []PlannerPlan;
    // Read-only. Nullable. Returns a collection of the specified tasks
    tasks []PlannerTask;
}
// Instantiates a new planner and sets the default values.
func NewPlanner()(*Planner) {
    m := &Planner{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the buckets property value. Read-only. Nullable. Returns a collection of the specified buckets
func (m *Planner) GetBuckets()([]PlannerBucket) {
    if m == nil {
        return nil
    } else {
        return m.buckets
    }
}
// Gets the plans property value. Read-only. Nullable. Returns a collection of the specified plans
func (m *Planner) GetPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// Gets the tasks property value. Read-only. Nullable. Returns a collection of the specified tasks
func (m *Planner) GetTasks()([]PlannerTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// The deserialization information for the current model
func (m *Planner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buckets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerBucket() })
        if err != nil {
            return err
        }
        res := make([]PlannerBucket, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerBucket))
        }
        m.SetBuckets(res)
        return nil
    }
    res["plans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        res := make([]PlannerPlan, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerPlan))
        }
        m.SetPlans(res)
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerTask() })
        if err != nil {
            return err
        }
        res := make([]PlannerTask, len(val))
        for i, v := range val {
            res[i] = *(v.(*PlannerTask))
        }
        m.SetTasks(res)
        return nil
    }
    return res
}
func (m *Planner) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Planner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBuckets()))
        for i, v := range m.GetBuckets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPlans()))
        for i, v := range m.GetPlans() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("plans", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the buckets property value. Read-only. Nullable. Returns a collection of the specified buckets
// Parameters:
//  - value : Value to set for the buckets property.
func (m *Planner) SetBuckets(value []PlannerBucket)() {
    m.buckets = value
}
// Sets the plans property value. Read-only. Nullable. Returns a collection of the specified plans
// Parameters:
//  - value : Value to set for the plans property.
func (m *Planner) SetPlans(value []PlannerPlan)() {
    m.plans = value
}
// Sets the tasks property value. Read-only. Nullable. Returns a collection of the specified tasks
// Parameters:
//  - value : Value to set for the tasks property.
func (m *Planner) SetTasks(value []PlannerTask)() {
    m.tasks = value
}
