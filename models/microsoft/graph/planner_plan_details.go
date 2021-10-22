package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerPlanDetails struct {
    Entity
    categoryDescriptions *PlannerCategoryDescriptions;
    sharedWith *PlannerUserIds;
}
func NewPlannerPlanDetails()(*PlannerPlanDetails) {
    m := &PlannerPlanDetails{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PlannerPlanDetails) GetCategoryDescriptions()(*PlannerCategoryDescriptions) {
    if m == nil {
        return nil
    } else {
        return m.categoryDescriptions
    }
}
func (m *PlannerPlanDetails) GetSharedWith()(*PlannerUserIds) {
    if m == nil {
        return nil
    } else {
        return m.sharedWith
    }
}
func (m *PlannerPlanDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryDescriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerCategoryDescriptions() })
        if err != nil {
            return err
        }
        m.SetCategoryDescriptions(val.(*PlannerCategoryDescriptions))
        return nil
    }
    res["sharedWith"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerUserIds() })
        if err != nil {
            return err
        }
        m.SetSharedWith(val.(*PlannerUserIds))
        return nil
    }
    return res
}
func (m *PlannerPlanDetails) IsNil()(bool) {
    return m == nil
}
func (m *PlannerPlanDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("categoryDescriptions", m.GetCategoryDescriptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedWith", m.GetSharedWith())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PlannerPlanDetails) SetCategoryDescriptions(value *PlannerCategoryDescriptions)() {
    m.categoryDescriptions = value
}
func (m *PlannerPlanDetails) SetSharedWith(value *PlannerUserIds)() {
    m.sharedWith = value
}
