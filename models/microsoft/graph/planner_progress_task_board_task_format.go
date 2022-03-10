package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerProgressTaskBoardTaskFormat provides operations to manage the collection of drive entities.
type PlannerProgressTaskBoardTaskFormat struct {
    Entity
    // Hint value used to order the task on the Progress view of the Task Board. The format is defined as outlined here.
    orderHint *string;
}
// NewPlannerProgressTaskBoardTaskFormat instantiates a new plannerProgressTaskBoardTaskFormat and sets the default values.
func NewPlannerProgressTaskBoardTaskFormat()(*PlannerProgressTaskBoardTaskFormat) {
    m := &PlannerProgressTaskBoardTaskFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerProgressTaskBoardTaskFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerProgressTaskBoardTaskFormatFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerProgressTaskBoardTaskFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerProgressTaskBoardTaskFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    return res
}
// GetOrderHint gets the orderHint property value. Hint value used to order the task on the Progress view of the Task Board. The format is defined as outlined here.
func (m *PlannerProgressTaskBoardTaskFormat) GetOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.orderHint
    }
}
func (m *PlannerProgressTaskBoardTaskFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlannerProgressTaskBoardTaskFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetOrderHint sets the orderHint property value. Hint value used to order the task on the Progress view of the Task Board. The format is defined as outlined here.
func (m *PlannerProgressTaskBoardTaskFormat) SetOrderHint(value *string)() {
    if m != nil {
        m.orderHint = value
    }
}
