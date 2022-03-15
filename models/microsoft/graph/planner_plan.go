package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerPlan provides operations to manage the drive singleton.
type PlannerPlan struct {
    Entity
    // Read-only. Nullable. Collection of buckets in the plan.
    buckets []PlannerBucketable;
    // Read-only. The user who created the plan.
    createdBy IdentitySetable;
    // Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. Nullable. Additional details about the plan.
    details PlannerPlanDetailsable;
    // ID of the Group that owns the plan. A valid group must exist before this field can be set. After it is set, this property can’t be updated.
    owner *string;
    // Read-only. Nullable. Collection of tasks in the plan.
    tasks []PlannerTaskable;
    // Required. Title of the plan.
    title *string;
}
// NewPlannerPlan instantiates a new plannerPlan and sets the default values.
func NewPlannerPlan()(*PlannerPlan) {
    m := &PlannerPlan{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerPlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerPlan(), nil
}
// GetBuckets gets the buckets property value. Read-only. Nullable. Collection of buckets in the plan.
func (m *PlannerPlan) GetBuckets()([]PlannerBucketable) {
    if m == nil {
        return nil
    } else {
        return m.buckets
    }
}
// GetCreatedBy gets the createdBy property value. Read-only. The user who created the plan.
func (m *PlannerPlan) GetCreatedBy()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerPlan) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDetails gets the details property value. Read-only. Nullable. Additional details about the plan.
func (m *PlannerPlan) GetDetails()(PlannerPlanDetailsable) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlan) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buckets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerBucketFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerBucketable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerBucketable)
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["details"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerPlanDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(PlannerPlanDetailsable))
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
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
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. ID of the Group that owns the plan. A valid group must exist before this field can be set. After it is set, this property can’t be updated.
func (m *PlannerPlan) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetTasks gets the tasks property value. Read-only. Nullable. Collection of tasks in the plan.
func (m *PlannerPlan) GetTasks()([]PlannerTaskable) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// GetTitle gets the title property value. Required. Title of the plan.
func (m *PlannerPlan) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
func (m *PlannerPlan) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerPlan) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBuckets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBuckets()))
        for i, v := range m.GetBuckets() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
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
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuckets sets the buckets property value. Read-only. Nullable. Collection of buckets in the plan.
func (m *PlannerPlan) SetBuckets(value []PlannerBucketable)() {
    if m != nil {
        m.buckets = value
    }
}
// SetCreatedBy sets the createdBy property value. Read-only. The user who created the plan.
func (m *PlannerPlan) SetCreatedBy(value IdentitySetable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PlannerPlan) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDetails sets the details property value. Read-only. Nullable. Additional details about the plan.
func (m *PlannerPlan) SetDetails(value PlannerPlanDetailsable)() {
    if m != nil {
        m.details = value
    }
}
// SetOwner sets the owner property value. ID of the Group that owns the plan. A valid group must exist before this field can be set. After it is set, this property can’t be updated.
func (m *PlannerPlan) SetOwner(value *string)() {
    if m != nil {
        m.owner = value
    }
}
// SetTasks sets the tasks property value. Read-only. Nullable. Collection of tasks in the plan.
func (m *PlannerPlan) SetTasks(value []PlannerTaskable)() {
    if m != nil {
        m.tasks = value
    }
}
// SetTitle sets the title property value. Required. Title of the plan.
func (m *PlannerPlan) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
