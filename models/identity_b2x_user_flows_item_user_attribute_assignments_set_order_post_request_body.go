package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody provides operations to call the setOrder method.
type IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The newAssignmentOrder property
    newAssignmentOrder AssignmentOrderable
}
// NewIdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody instantiates a new IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody and sets the default values.
func NewIdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody()(*IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) {
    m := &IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newAssignmentOrder"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAssignmentOrderFromDiscriminatorValue , m.SetNewAssignmentOrder)
    return res
}
// GetNewAssignmentOrder gets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetNewAssignmentOrder()(AssignmentOrderable) {
    return m.newAssignmentOrder
}
// Serialize serializes information the current object
func (m *IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newAssignmentOrder", m.GetNewAssignmentOrder())
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
func (m *IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewAssignmentOrder sets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *IdentityB2xUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) SetNewAssignmentOrder(value AssignmentOrderable)() {
    m.newAssignmentOrder = value
}
