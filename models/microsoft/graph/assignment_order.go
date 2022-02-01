package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentOrder 
type AssignmentOrder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A list of identityUserFlowAttribute IDs provided to determine the order in which attributes should be collected within a user flow.
    order []string;
}
// NewAssignmentOrder instantiates a new assignmentOrder and sets the default values.
func NewAssignmentOrder()(*AssignmentOrder) {
    m := &AssignmentOrder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentOrder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetOrder gets the order property value. A list of identityUserFlowAttribute IDs provided to determine the order in which attributes should be collected within a user flow.
func (m *AssignmentOrder) GetOrder()([]string) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentOrder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["order"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrder(res)
        }
        return nil
    }
    return res
}
func (m *AssignmentOrder) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssignmentOrder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetOrder() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentOrder) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOrder sets the order property value. A list of identityUserFlowAttribute IDs provided to determine the order in which attributes should be collected within a user flow.
func (m *AssignmentOrder) SetOrder(value []string)() {
    if m != nil {
        m.order = value
    }
}
