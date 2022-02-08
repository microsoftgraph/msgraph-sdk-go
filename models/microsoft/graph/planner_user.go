package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerUser 
type PlannerUser struct {
    Entity
    // Read-only. Nullable. Returns the plannerTasks assigned to the user.
    plans []PlannerPlan;
    // Read-only. Nullable. Returns the plannerPlans shared with the user.
    tasks []PlannerTask;
}
// NewPlannerUser instantiates a new plannerUser and sets the default values.
func NewPlannerUser()(*PlannerUser) {
    m := &PlannerUser{
        Entity: *NewEntity(),
    }
    return m
}
// GetPlans gets the plans property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) GetPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// GetTasks gets the tasks property value. Read-only. Nullable. Returns the plannerPlans shared with the user.
func (m *PlannerUser) GetTasks()([]PlannerTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["plans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerPlan() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerPlan, len(val))
            for i, v := range val {
                res[i] = *(v.(*PlannerPlan))
            }
            m.SetPlans(res)
        }
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerTask() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerTask, len(val))
            for i, v := range val {
                res[i] = *(v.(*PlannerTask))
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
func (m *PlannerUser) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPlans() != nil {
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
    if m.GetTasks() != nil {
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
// SetPlans sets the plans property value. Read-only. Nullable. Returns the plannerTasks assigned to the user.
func (m *PlannerUser) SetPlans(value []PlannerPlan)() {
    if m != nil {
        m.plans = value
    }
}
// SetTasks sets the tasks property value. Read-only. Nullable. Returns the plannerPlans shared with the user.
func (m *PlannerUser) SetTasks(value []PlannerTask)() {
    if m != nil {
        m.tasks = value
    }
}
