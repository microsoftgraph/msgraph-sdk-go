package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnPremisesExtensionAttributes struct {
    additionalData map[string]interface{};
    extensionAttribute1 *string;
    extensionAttribute10 *string;
    extensionAttribute11 *string;
    extensionAttribute12 *string;
    extensionAttribute13 *string;
    extensionAttribute14 *string;
    extensionAttribute15 *string;
    extensionAttribute2 *string;
    extensionAttribute3 *string;
    extensionAttribute4 *string;
    extensionAttribute5 *string;
    extensionAttribute6 *string;
    extensionAttribute7 *string;
    extensionAttribute8 *string;
    extensionAttribute9 *string;
}
func NewOnPremisesExtensionAttributes()(*OnPremisesExtensionAttributes) {
    m := &OnPremisesExtensionAttributes{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OnPremisesExtensionAttributes) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute1
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute10()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute10
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute11()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute11
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute12()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute12
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute13()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute13
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute14()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute14
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute15()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute15
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute2
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute3()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute3
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute4
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute5()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute5
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute6()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute6
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute7()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute7
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute8()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute8
    }
}
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute9()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttribute9
    }
}
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
func (m *OnPremisesExtensionAttributes) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute1(value *string)() {
    m.extensionAttribute1 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute10(value *string)() {
    m.extensionAttribute10 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute11(value *string)() {
    m.extensionAttribute11 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute12(value *string)() {
    m.extensionAttribute12 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute13(value *string)() {
    m.extensionAttribute13 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute14(value *string)() {
    m.extensionAttribute14 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute15(value *string)() {
    m.extensionAttribute15 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute2(value *string)() {
    m.extensionAttribute2 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute3(value *string)() {
    m.extensionAttribute3 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute4(value *string)() {
    m.extensionAttribute4 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute5(value *string)() {
    m.extensionAttribute5 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute6(value *string)() {
    m.extensionAttribute6 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute7(value *string)() {
    m.extensionAttribute7 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute8(value *string)() {
    m.extensionAttribute8 = value
}
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute9(value *string)() {
    m.extensionAttribute9 = value
}
