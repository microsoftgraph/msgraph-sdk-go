package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnPremisesExtensionAttributes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // First customizable extension attribute.
    extensionAttribute1 *string;
    // Tenth customizable extension attribute.
    extensionAttribute10 *string;
    // Eleventh customizable extension attribute.
    extensionAttribute11 *string;
    // Twelfth customizable extension attribute.
    extensionAttribute12 *string;
    // Thirteenth customizable extension attribute.
    extensionAttribute13 *string;
    // Fourteenth customizable extension attribute.
    extensionAttribute14 *string;
    // Fifteenth customizable extension attribute.
    extensionAttribute15 *string;
    // Second customizable extension attribute.
    extensionAttribute2 *string;
    // Third customizable extension attribute.
    extensionAttribute3 *string;
    // Fourth customizable extension attribute.
    extensionAttribute4 *string;
    // Fifth customizable extension attribute.
    extensionAttribute5 *string;
    // Sixth customizable extension attribute.
    extensionAttribute6 *string;
    // Seventh customizable extension attribute.
    extensionAttribute7 *string;
    // Eighth customizable extension attribute.
    extensionAttribute8 *string;
    // Ninth customizable extension attribute.
    extensionAttribute9 *string;
}
// Instantiates a new onPremisesExtensionAttributes and sets the default values.
func NewOnPremisesExtensionAttributes()(*OnPremisesExtensionAttributes) {
    m := &OnPremisesExtensionAttributes{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesExtensionAttributes) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the extensionAttribute1 property value. First customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute1
    }
}
// Gets the extensionAttribute10 property value. Tenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute10()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute10
    }
}
// Gets the extensionAttribute11 property value. Eleventh customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute11()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute11
    }
}
// Gets the extensionAttribute12 property value. Twelfth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute12()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute12
    }
}
// Gets the extensionAttribute13 property value. Thirteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute13()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute13
    }
}
// Gets the extensionAttribute14 property value. Fourteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute14()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute14
    }
}
// Gets the extensionAttribute15 property value. Fifteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute15()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute15
    }
}
// Gets the extensionAttribute2 property value. Second customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute2
    }
}
// Gets the extensionAttribute3 property value. Third customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute3()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute3
    }
}
// Gets the extensionAttribute4 property value. Fourth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute4
    }
}
// Gets the extensionAttribute5 property value. Fifth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute5()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute5
    }
}
// Gets the extensionAttribute6 property value. Sixth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute6
    }
}
// Gets the extensionAttribute7 property value. Seventh customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute7()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute7
    }
}
// Gets the extensionAttribute8 property value. Eighth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute8()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute8
    }
}
// Gets the extensionAttribute9 property value. Ninth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute9()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute9
    }
}
// The deserialization information for the current model
func (m *OnPremisesExtensionAttributes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["extensionAttribute1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute1(val)
        return nil
    }
    res["extensionAttribute10"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute10(val)
        return nil
    }
    res["extensionAttribute11"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute11(val)
        return nil
    }
    res["extensionAttribute12"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute12(val)
        return nil
    }
    res["extensionAttribute13"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute13(val)
        return nil
    }
    res["extensionAttribute14"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute14(val)
        return nil
    }
    res["extensionAttribute15"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute15(val)
        return nil
    }
    res["extensionAttribute2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute2(val)
        return nil
    }
    res["extensionAttribute3"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute3(val)
        return nil
    }
    res["extensionAttribute4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute4(val)
        return nil
    }
    res["extensionAttribute5"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute5(val)
        return nil
    }
    res["extensionAttribute6"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute6(val)
        return nil
    }
    res["extensionAttribute7"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute7(val)
        return nil
    }
    res["extensionAttribute8"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute8(val)
        return nil
    }
    res["extensionAttribute9"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionAttribute9(val)
        return nil
    }
    return res
}
func (m *OnPremisesExtensionAttributes) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnPremisesExtensionAttributes) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("extensionAttribute1", m.GetExtensionAttribute1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute10", m.GetExtensionAttribute10())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute11", m.GetExtensionAttribute11())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute12", m.GetExtensionAttribute12())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute13", m.GetExtensionAttribute13())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute14", m.GetExtensionAttribute14())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute15", m.GetExtensionAttribute15())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute2", m.GetExtensionAttribute2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute3", m.GetExtensionAttribute3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute4", m.GetExtensionAttribute4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute5", m.GetExtensionAttribute5())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute6", m.GetExtensionAttribute6())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute7", m.GetExtensionAttribute7())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute8", m.GetExtensionAttribute8())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute9", m.GetExtensionAttribute9())
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
func (m *OnPremisesExtensionAttributes) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the extensionAttribute1 property value. First customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute1 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute1(value *string)() {
    m.extensionAttribute1 = value
}
// Sets the extensionAttribute10 property value. Tenth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute10 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute10(value *string)() {
    m.extensionAttribute10 = value
}
// Sets the extensionAttribute11 property value. Eleventh customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute11 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute11(value *string)() {
    m.extensionAttribute11 = value
}
// Sets the extensionAttribute12 property value. Twelfth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute12 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute12(value *string)() {
    m.extensionAttribute12 = value
}
// Sets the extensionAttribute13 property value. Thirteenth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute13 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute13(value *string)() {
    m.extensionAttribute13 = value
}
// Sets the extensionAttribute14 property value. Fourteenth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute14 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute14(value *string)() {
    m.extensionAttribute14 = value
}
// Sets the extensionAttribute15 property value. Fifteenth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute15 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute15(value *string)() {
    m.extensionAttribute15 = value
}
// Sets the extensionAttribute2 property value. Second customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute2 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute2(value *string)() {
    m.extensionAttribute2 = value
}
// Sets the extensionAttribute3 property value. Third customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute3 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute3(value *string)() {
    m.extensionAttribute3 = value
}
// Sets the extensionAttribute4 property value. Fourth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute4 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute4(value *string)() {
    m.extensionAttribute4 = value
}
// Sets the extensionAttribute5 property value. Fifth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute5 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute5(value *string)() {
    m.extensionAttribute5 = value
}
// Sets the extensionAttribute6 property value. Sixth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute6 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute6(value *string)() {
    m.extensionAttribute6 = value
}
// Sets the extensionAttribute7 property value. Seventh customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute7 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute7(value *string)() {
    m.extensionAttribute7 = value
}
// Sets the extensionAttribute8 property value. Eighth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute8 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute8(value *string)() {
    m.extensionAttribute8 = value
}
// Sets the extensionAttribute9 property value. Ninth customizable extension attribute.
// Parameters:
//  - value : Value to set for the extensionAttribute9 property.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute9(value *string)() {
    m.extensionAttribute9 = value
}
