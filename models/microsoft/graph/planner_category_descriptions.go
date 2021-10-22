package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PlannerCategoryDescriptions struct {
    additionalData map[string]interface{};
    category1 *string;
    category2 *string;
    category3 *string;
    category4 *string;
    category5 *string;
    category6 *string;
}
func NewPlannerCategoryDescriptions()(*PlannerCategoryDescriptions) {
    m := &PlannerCategoryDescriptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PlannerCategoryDescriptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PlannerCategoryDescriptions) GetCategory1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category1
    }
}
func (m *PlannerCategoryDescriptions) GetCategory2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category2
    }
}
func (m *PlannerCategoryDescriptions) GetCategory3()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category3
    }
}
func (m *PlannerCategoryDescriptions) GetCategory4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category4
    }
}
func (m *PlannerCategoryDescriptions) GetCategory5()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category5
    }
}
func (m *PlannerCategoryDescriptions) GetCategory6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category6
    }
}
func (m *PlannerCategoryDescriptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["category1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory1(val)
        return nil
    }
    res["category2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory2(val)
        return nil
    }
    res["category3"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory3(val)
        return nil
    }
    res["category4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory4(val)
        return nil
    }
    res["category5"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory5(val)
        return nil
    }
    res["category6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory6(val)
        return nil
    }
    return res
}
func (m *PlannerCategoryDescriptions) IsNil()(bool) {
    return m == nil
}
func (m *PlannerCategoryDescriptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("category1", m.GetCategory1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category2", m.GetCategory2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category3", m.GetCategory3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category4", m.GetCategory4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category5", m.GetCategory5())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category6", m.GetCategory6())
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
func (m *PlannerCategoryDescriptions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PlannerCategoryDescriptions) SetCategory1(value *string)() {
    m.category1 = value
}
func (m *PlannerCategoryDescriptions) SetCategory2(value *string)() {
    m.category2 = value
}
func (m *PlannerCategoryDescriptions) SetCategory3(value *string)() {
    m.category3 = value
}
func (m *PlannerCategoryDescriptions) SetCategory4(value *string)() {
    m.category4 = value
}
func (m *PlannerCategoryDescriptions) SetCategory5(value *string)() {
    m.category5 = value
}
func (m *PlannerCategoryDescriptions) SetCategory6(value *string)() {
    m.category6 = value
}
