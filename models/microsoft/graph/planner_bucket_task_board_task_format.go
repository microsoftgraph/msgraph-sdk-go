package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerBucketTaskBoardTaskFormat struct {
    Entity
    orderHint *string;
}
func NewPlannerBucketTaskBoardTaskFormat()(*PlannerBucketTaskBoardTaskFormat) {
    m := &PlannerBucketTaskBoardTaskFormat{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PlannerBucketTaskBoardTaskFormat) GetOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderHint
    }
}
func (m *PlannerBucketTaskBoardTaskFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["orderHint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrderHint(val)
        return nil
    }
    return res
}
func (m *PlannerBucketTaskBoardTaskFormat) IsNil()(bool) {
    return m == nil
}
func (m *PlannerBucketTaskBoardTaskFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("orderHint", m.GetOrderHint())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PlannerBucketTaskBoardTaskFormat) SetOrderHint(value *string)() {
    m.orderHint = value
}
