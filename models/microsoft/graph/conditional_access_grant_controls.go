package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConditionalAccessGrantControls struct {
    additionalData map[string]interface{};
    builtInControls []ConditionalAccessGrantControl;
    customAuthenticationFactors []string;
    operator *string;
    termsOfUse []string;
}
func NewConditionalAccessGrantControls()(*ConditionalAccessGrantControls) {
    m := &ConditionalAccessGrantControls{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConditionalAccessGrantControls) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConditionalAccessGrantControls) GetBuiltInControls()([]ConditionalAccessGrantControl) {
    if m == nil {
        return nil
    } else {
        return m.builtInControls
    }
}
func (m *ConditionalAccessGrantControls) GetCustomAuthenticationFactors()([]string) {
    if m == nil {
        return nil
    } else {
        return m.customAuthenticationFactors
    }
}
func (m *ConditionalAccessGrantControls) GetOperator()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operator
    }
}
func (m *ConditionalAccessGrantControls) GetTermsOfUse()([]string) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUse
    }
}
func (m *ConditionalAccessGrantControls) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["builtInControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseConditionalAccessGrantControl)
        if err != nil {
            return err
        }
        res := make([]ConditionalAccessGrantControl, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConditionalAccessGrantControl))
        }
        m.SetBuiltInControls(res)
        return nil
    }
    res["customAuthenticationFactors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetCustomAuthenticationFactors(res)
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
    res["termsOfUse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTermsOfUse(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessGrantControls) IsNil()(bool) {
    return m == nil
}
func (m *ConditionalAccessGrantControls) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("builtInControls", SerializeConditionalAccessGrantControl(m.GetBuiltInControls()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("customAuthenticationFactors", m.GetCustomAuthenticationFactors())
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
        err := writer.WriteCollectionOfStringValues("termsOfUse", m.GetTermsOfUse())
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
func (m *ConditionalAccessGrantControls) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConditionalAccessGrantControls) SetBuiltInControls(value []ConditionalAccessGrantControl)() {
    m.builtInControls = value
}
func (m *ConditionalAccessGrantControls) SetCustomAuthenticationFactors(value []string)() {
    m.customAuthenticationFactors = value
}
func (m *ConditionalAccessGrantControls) SetOperator(value *string)() {
    m.operator = value
}
func (m *ConditionalAccessGrantControls) SetTermsOfUse(value []string)() {
    m.termsOfUse = value
}
