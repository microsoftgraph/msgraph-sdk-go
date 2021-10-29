package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerPlanDetails struct {
    Entity
    // An object that specifies the descriptions of the six categories that can be associated with tasks in the plan
    categoryDescriptions *PlannerCategoryDescriptions;
    // Set of user ids that this plan is shared with. If you are leveraging Microsoft 365 groups, use the Groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection though it is not required for them to access the plan owned by the group.
    sharedWith *PlannerUserIds;
}
// Instantiates a new plannerPlanDetails and sets the default values.
func NewPlannerPlanDetails()(*PlannerPlanDetails) {
    m := &PlannerPlanDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the categoryDescriptions property value. An object that specifies the descriptions of the six categories that can be associated with tasks in the plan
func (m *PlannerPlanDetails) GetCategoryDescriptions()(*PlannerCategoryDescriptions) {
    if m == nil {
        return nil
    } else {
        return m.categoryDescriptions
    }
}
// Gets the sharedWith property value. Set of user ids that this plan is shared with. If you are leveraging Microsoft 365 groups, use the Groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection though it is not required for them to access the plan owned by the group.
func (m *PlannerPlanDetails) GetSharedWith()(*PlannerUserIds) {
    if m == nil {
        return nil
    } else {
        return m.sharedWith
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the categoryDescriptions property value. An object that specifies the descriptions of the six categories that can be associated with tasks in the plan
// Parameters:
//  - value : Value to set for the categoryDescriptions property.
func (m *PlannerPlanDetails) SetCategoryDescriptions(value *PlannerCategoryDescriptions)() {
    m.categoryDescriptions = value
}
// Sets the sharedWith property value. Set of user ids that this plan is shared with. If you are leveraging Microsoft 365 groups, use the Groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection though it is not required for them to access the plan owned by the group.
// Parameters:
//  - value : Value to set for the sharedWith property.
func (m *PlannerPlanDetails) SetSharedWith(value *PlannerUserIds)() {
    m.sharedWith = value
}
