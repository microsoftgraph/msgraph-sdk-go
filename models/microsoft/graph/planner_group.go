package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerGroup provides operations to manage the collection of group entities.
type PlannerGroup struct {
    Entity
    // Read-only. Nullable. Returns the plannerPlans owned by the group.
    plans []PlannerPlanable;
}
// NewPlannerGroup instantiates a new plannerGroup and sets the default values.
func NewPlannerGroup()(*PlannerGroup) {
    m := &PlannerGroup{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerGroupFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    return res
}
// GetPlans gets the plans property value. Read-only. Nullable. Returns the plannerPlans owned by the group.
func (m *PlannerGroup) GetPlans()([]PlannerPlanable) {
    if m == nil {
        return nil
    } else {
        return m.plans
    }
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
    return nil
}
// SetPlans sets the plans property value. Read-only. Nullable. Returns the plannerPlans owned by the group.
func (m *PlannerGroup) SetPlans(value []PlannerPlanable)() {
    if m != nil {
        m.plans = value
    }
}
