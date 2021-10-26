package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type WorkbookFilterCriteria struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    color *string;
    // 
    criterion1 *string;
    // 
    criterion2 *string;
    // 
    dynamicCriteria *string;
    // 
    filterOn *string;
    // 
    icon *WorkbookIcon;
    // 
    operator *string;
    // 
    values *Json;
}
// Instantiates a new workbookFilterCriteria and sets the default values.
func NewWorkbookFilterCriteria()(*WorkbookFilterCriteria) {
    m := &WorkbookFilterCriteria{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookFilterCriteria) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the color property value. 
func (m *WorkbookFilterCriteria) GetColor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.color
    }
}
// Gets the criterion1 property value. 
func (m *WorkbookFilterCriteria) GetCriterion1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.criterion1
    }
}
// Gets the criterion2 property value. 
func (m *WorkbookFilterCriteria) GetCriterion2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.criterion2
    }
}
// Gets the dynamicCriteria property value. 
func (m *WorkbookFilterCriteria) GetDynamicCriteria()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dynamicCriteria
    }
}
// Gets the filterOn property value. 
func (m *WorkbookFilterCriteria) GetFilterOn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.filterOn
    }
}
// Gets the icon property value. 
func (m *WorkbookFilterCriteria) GetIcon()(*WorkbookIcon) {
    if m == nil {
        return nil
    } else {
        return m.icon
    }
}
// Gets the operator property value. 
func (m *WorkbookFilterCriteria) GetOperator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operator
    }
}
// Gets the values property value. 
func (m *WorkbookFilterCriteria) GetValues()(*Json) {
    if m == nil {
        return nil
    } else {
        return m.values
    }
}
// The deserialization information for the current model
func (m *WorkbookFilterCriteria) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["color"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColor(val)
        return nil
    }
    res["criterion1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCriterion1(val)
        return nil
    }
    res["criterion2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCriterion2(val)
        return nil
    }
    res["dynamicCriteria"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDynamicCriteria(val)
        return nil
    }
    res["filterOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFilterOn(val)
        return nil
    }
    res["icon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookIcon() })
        if err != nil {
            return err
        }
        m.SetIcon(val.(*WorkbookIcon))
        return nil
    }
    res["operator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperator(val)
        return nil
    }
    res["values"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewJson() })
        if err != nil {
            return err
        }
        m.SetValues(val.(*Json))
        return nil
    }
    return res
}
func (m *WorkbookFilterCriteria) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkbookFilterCriteria) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criterion1", m.GetCriterion1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criterion2", m.GetCriterion2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dynamicCriteria", m.GetDynamicCriteria())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filterOn", m.GetFilterOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("icon", m.GetIcon())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operator", m.GetOperator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("values", m.GetValues())
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
func (m *WorkbookFilterCriteria) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the color property value. 
// Parameters:
//  - value : Value to set for the color property.
func (m *WorkbookFilterCriteria) SetColor(value *string)() {
    m.color = value
}
// Sets the criterion1 property value. 
// Parameters:
//  - value : Value to set for the criterion1 property.
func (m *WorkbookFilterCriteria) SetCriterion1(value *string)() {
    m.criterion1 = value
}
// Sets the criterion2 property value. 
// Parameters:
//  - value : Value to set for the criterion2 property.
func (m *WorkbookFilterCriteria) SetCriterion2(value *string)() {
    m.criterion2 = value
}
// Sets the dynamicCriteria property value. 
// Parameters:
//  - value : Value to set for the dynamicCriteria property.
func (m *WorkbookFilterCriteria) SetDynamicCriteria(value *string)() {
    m.dynamicCriteria = value
}
// Sets the filterOn property value. 
// Parameters:
//  - value : Value to set for the filterOn property.
func (m *WorkbookFilterCriteria) SetFilterOn(value *string)() {
    m.filterOn = value
}
// Sets the icon property value. 
// Parameters:
//  - value : Value to set for the icon property.
func (m *WorkbookFilterCriteria) SetIcon(value *WorkbookIcon)() {
    m.icon = value
}
// Sets the operator property value. 
// Parameters:
//  - value : Value to set for the operator property.
func (m *WorkbookFilterCriteria) SetOperator(value *string)() {
    m.operator = value
}
// Sets the values property value. 
// Parameters:
//  - value : Value to set for the values property.
func (m *WorkbookFilterCriteria) SetValues(value *Json)() {
    m.values = value
}
