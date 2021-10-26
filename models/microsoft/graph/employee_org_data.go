package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EmployeeOrgData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The cost center associated with the user. Returned only on $select. Supports $filter.
    costCenter *string;
    // The name of the division in which the user works. Returned only on $select. Supports $filter.
    division *string;
}
// Instantiates a new employeeOrgData and sets the default values.
func NewEmployeeOrgData()(*EmployeeOrgData) {
    m := &EmployeeOrgData{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmployeeOrgData) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the costCenter property value. The cost center associated with the user. Returned only on $select. Supports $filter.
func (m *EmployeeOrgData) GetCostCenter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.costCenter
    }
}
// Gets the division property value. The name of the division in which the user works. Returned only on $select. Supports $filter.
func (m *EmployeeOrgData) GetDivision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.division
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EmployeeOrgData) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the costCenter property value. The cost center associated with the user. Returned only on $select. Supports $filter.
// Parameters:
//  - value : Value to set for the costCenter property.
func (m *EmployeeOrgData) SetCostCenter(value *string)() {
    m.costCenter = value
}
// Sets the division property value. The name of the division in which the user works. Returned only on $select. Supports $filter.
// Parameters:
//  - value : Value to set for the division property.
func (m *EmployeeOrgData) SetDivision(value *string)() {
    m.division = value
}
