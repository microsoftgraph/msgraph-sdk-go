package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EmployeeOrgData struct {
    additionalData map[string]interface{};
    costCenter *string;
    division *string;
}
func NewEmployeeOrgData()(*EmployeeOrgData) {
    m := &EmployeeOrgData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EmployeeOrgData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EmployeeOrgData) GetCostCenter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.costCenter
    }
}
func (m *EmployeeOrgData) GetDivision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.division
    }
}
func (m *EmployeeOrgData) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["costCenter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCostCenter(val)
        return nil
    }
    res["division"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDivision(val)
        return nil
    }
    return res
}
func (m *EmployeeOrgData) IsNil()(bool) {
    return m == nil
}
func (m *EmployeeOrgData) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("costCenter", m.GetCostCenter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("division", m.GetDivision())
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
func (m *EmployeeOrgData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EmployeeOrgData) SetCostCenter(value *string)() {
    m.costCenter = value
}
func (m *EmployeeOrgData) SetDivision(value *string)() {
    m.division = value
}
