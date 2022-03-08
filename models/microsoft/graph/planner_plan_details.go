package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerPlanDetails provides operations to manage the drive singleton.
type PlannerPlanDetails struct {
    Entity
    // An object that specifies the descriptions of the six categories that can be associated with tasks in the plan
    categoryDescriptions PlannerCategoryDescriptionsable;
    // Set of user ids that this plan is shared with. If you are leveraging Microsoft 365 groups, use the Groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection though it is not required for them to access the plan owned by the group.
    sharedWith PlannerUserIdsable;
}
// NewPlannerPlanDetails instantiates a new plannerPlanDetails and sets the default values.
func NewPlannerPlanDetails()(*PlannerPlanDetails) {
    m := &PlannerPlanDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerPlanDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanDetailsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerPlanDetails(), nil
}
// GetCategoryDescriptions gets the categoryDescriptions property value. An object that specifies the descriptions of the six categories that can be associated with tasks in the plan
func (m *PlannerPlanDetails) GetCategoryDescriptions()(PlannerCategoryDescriptionsable) {
    if m == nil {
        return nil
    } else {
        return m.categoryDescriptions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryDescriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerCategoryDescriptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryDescriptions(val.(PlannerCategoryDescriptionsable))
        }
        return nil
    }
    res["sharedWith"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerUserIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedWith(val.(PlannerUserIdsable))
        }
        return nil
    }
    return res
}
// GetSharedWith gets the sharedWith property value. Set of user ids that this plan is shared with. If you are leveraging Microsoft 365 groups, use the Groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection though it is not required for them to access the plan owned by the group.
func (m *PlannerPlanDetails) GetSharedWith()(PlannerUserIdsable) {
    if m == nil {
        return nil
    } else {
        return m.sharedWith
    }
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
// SetCategoryDescriptions sets the categoryDescriptions property value. An object that specifies the descriptions of the six categories that can be associated with tasks in the plan
func (m *PlannerPlanDetails) SetCategoryDescriptions(value PlannerCategoryDescriptionsable)() {
    if m != nil {
        m.categoryDescriptions = value
    }
}
// SetSharedWith sets the sharedWith property value. Set of user ids that this plan is shared with. If you are leveraging Microsoft 365 groups, use the Groups API to manage group membership to share the group's plan. You can also add existing members of the group to this collection though it is not required for them to access the plan owned by the group.
func (m *PlannerPlanDetails) SetSharedWith(value PlannerUserIdsable)() {
    if m != nil {
        m.sharedWith = value
    }
}
