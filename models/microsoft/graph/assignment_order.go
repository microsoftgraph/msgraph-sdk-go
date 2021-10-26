package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AssignmentOrder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A list of identityUserFlowAttribute object identifiers that determine the order in which attributes should be collected within a user flow.
    order []string;
}
// Instantiates a new assignmentOrder and sets the default values.
func NewAssignmentOrder()(*AssignmentOrder) {
    m := &AssignmentOrder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentOrder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the order property value. A list of identityUserFlowAttribute object identifiers that determine the order in which attributes should be collected within a user flow.
func (m *AssignmentOrder) GetOrder()([]string) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
// The deserialization information for the current model
func (m *AssignmentOrder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["order"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOrder(res)
        return nil
    }
    return res
}
func (m *AssignmentOrder) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AssignmentOrder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("order", m.GetOrder())
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
func (m *AssignmentOrder) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the order property value. A list of identityUserFlowAttribute object identifiers that determine the order in which attributes should be collected within a user flow.
// Parameters:
//  - value : Value to set for the order property.
func (m *AssignmentOrder) SetOrder(value []string)() {
    m.order = value
}
