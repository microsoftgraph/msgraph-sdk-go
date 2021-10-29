package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerGroup struct {
    Entity
    // Read-only. Nullable. Returns the plannerPlans owned by the group.
    plans []PlannerPlan;
}
// Instantiates a new plannerGroup and sets the default values.
func NewPlannerGroup()(*PlannerGroup) {
    m := &PlannerGroup{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the plans property value. Read-only. Nullable. Returns the plannerPlans owned by the group.
func (m *PlannerGroup) GetPlans()([]PlannerPlan) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
}
// The deserialization information for the current model
func (m *PlannerGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    return res
}
func (m *PlannerGroup) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the plans property value. Read-only. Nullable. Returns the plannerPlans owned by the group.
// Parameters:
//  - value : Value to set for the plans property.
func (m *PlannerGroup) SetPlans(value []PlannerPlan)() {
    m.plans = value
}
