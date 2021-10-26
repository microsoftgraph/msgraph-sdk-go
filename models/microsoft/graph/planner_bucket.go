package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerBucket struct {
    Entity
    // Name of the bucket.
    name *string;
    // Hint used to order items of this type in a list view. The format is defined as outlined here.
    orderHint *string;
    // Plan ID to which the bucket belongs.
    planId *string;
    // Read-only. Nullable. The collection of tasks in the bucket.
    tasks []PlannerTask;
}
// Instantiates a new plannerBucket and sets the default values.
func NewPlannerBucket()(*PlannerBucket) {
    m := &PlannerBucket{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the name property value. Name of the bucket.
func (m *PlannerBucket) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the orderHint property value. Hint used to order items of this type in a list view. The format is defined as outlined here.
func (m *PlannerBucket) GetOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderHint
    }
}
// Gets the planId property value. Plan ID to which the bucket belongs.
func (m *PlannerBucket) GetPlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.planId
    }
}
// Gets the tasks property value. Read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerBucket) GetTasks()([]PlannerTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// The deserialization information for the current model
func (m *PlannerBucket) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["orderHint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrderHint(val)
        return nil
    }
    res["planId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPlanId(val)
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
func (m *PlannerBucket) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the name property value. Name of the bucket.
// Parameters:
//  - value : Value to set for the name property.
func (m *PlannerBucket) SetName(value *string)() {
    m.name = value
}
// Sets the orderHint property value. Hint used to order items of this type in a list view. The format is defined as outlined here.
// Parameters:
//  - value : Value to set for the orderHint property.
func (m *PlannerBucket) SetOrderHint(value *string)() {
    m.orderHint = value
}
// Sets the planId property value. Plan ID to which the bucket belongs.
// Parameters:
//  - value : Value to set for the planId property.
func (m *PlannerBucket) SetPlanId(value *string)() {
    m.planId = value
}
// Sets the tasks property value. Read-only. Nullable. The collection of tasks in the bucket.
// Parameters:
//  - value : Value to set for the tasks property.
func (m *PlannerBucket) SetTasks(value []PlannerTask)() {
    m.tasks = value
}
