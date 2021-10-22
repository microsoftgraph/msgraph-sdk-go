package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignmentOrder struct {
    additionalData map[string]interface{};
    order []string;
}
func NewAssignmentOrder()(*AssignmentOrder) {
    m := &AssignmentOrder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssignmentOrder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssignmentOrder) GetOrder()([]string) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
func (m *AssignmentOrder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["order"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOrder(res)
        return nil
    }
    return res
}
func (m *AssignmentOrder) IsNil()(bool) {
    return m == nil
}
func (m *AssignmentOrder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("order", m.GetOrder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AssignmentOrder) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssignmentOrder) SetOrder(value []string)() {
    m.order = value
}
