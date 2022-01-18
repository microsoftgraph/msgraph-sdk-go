package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerPlanDetails 
type PlannerPlanDetails struct {
    Entity
    // An object that specifies the descriptions of the 25 categories that can be associated with tasks in the plan
    categoryDescriptions *PlannerCategoryDescriptions;
    // The set of user IDs that this plan is shared with. If you are using Microsoft 365 groups, use the groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection, although it is not required in order for them to access the plan owned by the group.
    sharedWith *PlannerUserIds;
}
// NewPlannerPlanDetails instantiates a new plannerPlanDetails and sets the default values.
func NewPlannerPlanDetails()(*PlannerPlanDetails) {
    m := &PlannerPlanDetails{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategoryDescriptions gets the categoryDescriptions property value. An object that specifies the descriptions of the 25 categories that can be associated with tasks in the plan
func (m *PlannerPlanDetails) GetCategoryDescriptions()(*PlannerCategoryDescriptions) {
    if m == nil {
        return nil
    } else {
        return m.categoryDescriptions
    }
}
// GetSharedWith gets the sharedWith property value. The set of user IDs that this plan is shared with. If you are using Microsoft 365 groups, use the groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection, although it is not required in order for them to access the plan owned by the group.
func (m *PlannerPlanDetails) GetSharedWith()(*PlannerUserIds) {
    if m == nil {
        return nil
    } else {
        return m.sharedWith
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryDescriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerCategoryDescriptions() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryDescriptions(val.(*PlannerCategoryDescriptions))
        }
        return nil
    }
    res["sharedWith"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerUserIds() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedWith(val.(*PlannerUserIds))
        }
        return nil
    }
    return res
}
func (m *PlannerPlanDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCategoryDescriptions sets the categoryDescriptions property value. An object that specifies the descriptions of the 25 categories that can be associated with tasks in the plan
func (m *PlannerPlanDetails) SetCategoryDescriptions(value *PlannerCategoryDescriptions)() {
    if m != nil {
        m.categoryDescriptions = value
    }
}
// SetSharedWith sets the sharedWith property value. The set of user IDs that this plan is shared with. If you are using Microsoft 365 groups, use the groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection, although it is not required in order for them to access the plan owned by the group.
func (m *PlannerPlanDetails) SetSharedWith(value *PlannerUserIds)() {
    if m != nil {
        m.sharedWith = value
    }
}
