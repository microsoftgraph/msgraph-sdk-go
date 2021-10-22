package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerAssignedToTaskBoardTaskFormat struct {
    Entity
    orderHintsByAssignee *PlannerOrderHintsByAssignee;
    unassignedOrderHint *string;
}
func NewPlannerAssignedToTaskBoardTaskFormat()(*PlannerAssignedToTaskBoardTaskFormat) {
    m := &PlannerAssignedToTaskBoardTaskFormat{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssignee()(*PlannerOrderHintsByAssignee) {
    if m == nil {
        return nil
    } else {
        return m.orderHintsByAssignee
    }
}
func (m *PlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unassignedOrderHint
    }
}
func (m *PlannerAssignedToTaskBoardTaskFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["orderHintsByAssignee"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerOrderHintsByAssignee() })
        if err != nil {
            return err
        }
        m.SetOrderHintsByAssignee(val.(*PlannerOrderHintsByAssignee))
        return nil
    }
    res["unassignedOrderHint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUnassignedOrderHint(val)
        return nil
    }
    return res
}
func (m *PlannerAssignedToTaskBoardTaskFormat) IsNil()(bool) {
    return m == nil
}
func (m *PlannerAssignedToTaskBoardTaskFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("orderHintsByAssignee", m.GetOrderHintsByAssignee())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("unassignedOrderHint", m.GetUnassignedOrderHint())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssignee(value *PlannerOrderHintsByAssignee)() {
    m.orderHintsByAssignee = value
}
func (m *PlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHint(value *string)() {
    m.unassignedOrderHint = value
}
