package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerBucket provides operations to manage the collection of drive entities.
type PlannerBucket struct {
    Entity
    // Name of the bucket.
    name *string;
    // Hint used to order items of this type in a list view. The format is defined as outlined here.
    orderHint *string;
    // Plan ID to which the bucket belongs.
    planId *string;
    // Read-only. Nullable. The collection of tasks in the bucket.
    tasks []PlannerTaskable;
}
// NewPlannerBucket instantiates a new plannerBucket and sets the default values.
func NewPlannerBucket()(*PlannerBucket) {
    m := &PlannerBucket{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerBucketFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerBucketFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerBucket(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerBucket) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["orderHint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderHint(val)
        }
        return nil
    }
    res["planId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanId(val)
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
// GetName gets the name property value. Name of the bucket.
func (m *PlannerBucket) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOrderHint gets the orderHint property value. Hint used to order items of this type in a list view. The format is defined as outlined here.
func (m *PlannerBucket) GetOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderHint
    }
}
// GetPlanId gets the planId property value. Plan ID to which the bucket belongs.
func (m *PlannerBucket) GetPlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.planId
    }
}
// GetTasks gets the tasks property value. Read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerBucket) GetTasks()([]PlannerTaskable) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
func (m *PlannerBucket) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerBucket) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("orderHint", m.GetOrderHint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("planId", m.GetPlanId())
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
// SetName sets the name property value. Name of the bucket.
func (m *PlannerBucket) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOrderHint sets the orderHint property value. Hint used to order items of this type in a list view. The format is defined as outlined here.
func (m *PlannerBucket) SetOrderHint(value *string)() {
    if m != nil {
        m.orderHint = value
    }
}
// SetPlanId sets the planId property value. Plan ID to which the bucket belongs.
func (m *PlannerBucket) SetPlanId(value *string)() {
    if m != nil {
        m.planId = value
    }
}
// SetTasks sets the tasks property value. Read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerBucket) SetTasks(value []PlannerTaskable)() {
    if m != nil {
        m.tasks = value
    }
}
