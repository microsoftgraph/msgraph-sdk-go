package applycustomfilter

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplyCustomFilterRequestBody 
type ApplyCustomFilterRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    criteria1 *string;
    // 
    criteria2 *string;
    // 
    oper *string;
}
// NewApplyCustomFilterRequestBody instantiates a new applyCustomFilterRequestBody and sets the default values.
func NewApplyCustomFilterRequestBody()(*ApplyCustomFilterRequestBody) {
    m := &ApplyCustomFilterRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyCustomFilterRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCriteria1 gets the criteria1 property value. 
func (m *ApplyCustomFilterRequestBody) GetCriteria1()(*string) {
    if m == nil {
        return nil
    } else {
        return m.criteria1
    }
}
// GetCriteria2 gets the criteria2 property value. 
func (m *ApplyCustomFilterRequestBody) GetCriteria2()(*string) {
    if m == nil {
        return nil
    } else {
        return m.criteria2
    }
}
// GetOper gets the oper property value. 
func (m *ApplyCustomFilterRequestBody) GetOper()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oper
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyCustomFilterRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["criteria1"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria1(val)
        }
        return nil
    }
    res["criteria2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriteria2(val)
        }
        return nil
    }
    res["oper"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOper(val)
        }
        return nil
    }
    return res
}
func (m *ApplyCustomFilterRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyCustomFilterRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCriteria1 sets the criteria1 property value. 
func (m *ApplyCustomFilterRequestBody) SetCriteria1(value *string)() {
    if m != nil {
        m.criteria1 = value
    }
}
// SetCriteria2 sets the criteria2 property value. 
func (m *ApplyCustomFilterRequestBody) SetCriteria2(value *string)() {
    if m != nil {
        m.criteria2 = value
    }
}
// SetOper sets the oper property value. 
func (m *ApplyCustomFilterRequestBody) SetOper(value *string)() {
    if m != nil {
        m.oper = value
    }
}
