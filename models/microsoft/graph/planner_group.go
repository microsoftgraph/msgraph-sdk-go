package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerGroup 
type PlannerGroup struct {
    Entity
    // Read-only. Nullable. Returns the plannerPlans owned by the group.
    plans []PlannerPlan;
}
// NewPlannerGroup instantiates a new plannerGroup and sets the default values.
func NewPlannerGroup()(*PlannerGroup) {
    m := &PlannerGroup{
        Entity: *NewEntity(),
    }
    return m
}
// GetPlans gets the plans property value. Read-only. Nullable. Returns the plannerPlans owned by the group.
func (m *PlannerGroup) GetPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
func (m *PlannerGroup) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    return nil
}
// SetPlans sets the plans property value. Read-only. Nullable. Returns the plannerPlans owned by the group.
func (m *PlannerGroup) SetPlans(value []PlannerPlan)() {
    m.plans = value
}
