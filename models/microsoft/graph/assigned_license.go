package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignedLicense 
type AssignedLicense struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A collection of the unique identifiers for plans that have been disabled.
    disabledPlans []string;
    // The unique identifier for the SKU.
    skuId *string;
}
// NewAssignedLicense instantiates a new assignedLicense and sets the default values.
func NewAssignedLicense()(*AssignedLicense) {
    m := &AssignedLicense{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedLicense) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisabledPlans gets the disabledPlans property value. A collection of the unique identifiers for plans that have been disabled.
func (m *AssignedLicense) GetDisabledPlans()([]string) {
    if m == nil {
        return nil
    } else {
        return m.disabledPlans
    }
}
// GetSkuId gets the skuId property value. The unique identifier for the SKU.
func (m *AssignedLicense) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignedLicense) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["disabledPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDisabledPlans(res)
        }
        return nil
    }
    res["skuId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    return res
}
func (m *AssignedLicense) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedLicense) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisabledPlans sets the disabledPlans property value. A collection of the unique identifiers for plans that have been disabled.
func (m *AssignedLicense) SetDisabledPlans(value []string)() {
    if m != nil {
        m.disabledPlans = value
    }
}
// SetSkuId sets the skuId property value. The unique identifier for the SKU.
func (m *AssignedLicense) SetSkuId(value *string)() {
    if m != nil {
        m.skuId = value
    }
}
