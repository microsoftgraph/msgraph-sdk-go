package applycustomfilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApplyCustomFilterRequestBody struct {
    additionalData map[string]interface{};
    criteria1 *string;
    criteria2 *string;
    oper *string;
}
func NewApplyCustomFilterRequestBody()(*ApplyCustomFilterRequestBody) {
    m := &ApplyCustomFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApplyCustomFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApplyCustomFilterRequestBody) GetCriteria1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.criteria1
    }
}
func (m *ApplyCustomFilterRequestBody) GetCriteria2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.criteria2
    }
}
func (m *ApplyCustomFilterRequestBody) GetOper()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oper
    }
}
func (m *ApplyCustomFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["criteria1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCriteria1(val)
        return nil
    }
    res["criteria2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCriteria2(val)
        return nil
    }
    res["oper"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOper(val)
        return nil
    }
    return res
}
func (m *ApplyCustomFilterRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ApplyCustomFilterRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("criteria1", m.GetCriteria1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criteria2", m.GetCriteria2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oper", m.GetOper())
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
func (m *ApplyCustomFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApplyCustomFilterRequestBody) SetCriteria1(value *string)() {
    m.criteria1 = value
}
func (m *ApplyCustomFilterRequestBody) SetCriteria2(value *string)() {
    m.criteria2 = value
}
func (m *ApplyCustomFilterRequestBody) SetOper(value *string)() {
    m.oper = value
}
