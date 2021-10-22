package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignedLicense struct {
    additionalData map[string]interface{};
    disabledPlans []string;
    skuId *string;
}
func NewAssignedLicense()(*AssignedLicense) {
    m := &AssignedLicense{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssignedLicense) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssignedLicense) GetDisabledPlans()([]string) {
    if m == nil {
        return nil
    } else {
        return m.disabledPlans
    }
}
func (m *AssignedLicense) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
func (m *AssignedLicense) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["disabledPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDisabledPlans(res)
        return nil
    }
    res["skuId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSkuId(val)
        return nil
    }
    return res
}
func (m *AssignedLicense) IsNil()(bool) {
    return m == nil
}
func (m *AssignedLicense) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("disabledPlans", m.GetDisabledPlans())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("skuId", m.GetSkuId())
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
func (m *AssignedLicense) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssignedLicense) SetDisabledPlans(value []string)() {
    m.disabledPlans = value
}
func (m *AssignedLicense) SetSkuId(value *string)() {
    m.skuId = value
}
