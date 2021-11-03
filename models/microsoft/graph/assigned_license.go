package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AssignedLicense struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A collection of the unique identifiers for plans that have been disabled.
    disabledPlans []string;
    // The unique identifier for the SKU.
    skuId *string;
}
// Instantiates a new assignedLicense and sets the default values.
func NewAssignedLicense()(*AssignedLicense) {
    m := &AssignedLicense{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedLicense) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the disabledPlans property value. A collection of the unique identifiers for plans that have been disabled.
func (m *AssignedLicense) GetDisabledPlans()([]string) {
    if m == nil {
        return nil
    } else {
        return m.disabledPlans
    }
}
// Gets the skuId property value. The unique identifier for the SKU.
func (m *AssignedLicense) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// The deserialization information for the current model
func (m *AssignedLicense) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["disabledPlans"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AssignedLicense) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the disabledPlans property value. A collection of the unique identifiers for plans that have been disabled.
// Parameters:
//  - value : Value to set for the disabledPlans property.
func (m *AssignedLicense) SetDisabledPlans(value []string)() {
    m.disabledPlans = value
}
// Sets the skuId property value. The unique identifier for the SKU.
// Parameters:
//  - value : Value to set for the skuId property.
func (m *AssignedLicense) SetSkuId(value *string)() {
    m.skuId = value
}
