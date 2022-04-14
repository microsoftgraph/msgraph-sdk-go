package setorder

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// SetOrderRequestBody provides operations to call the setOrder method.
type SetOrderRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The newAssignmentOrder property
    newAssignmentOrder iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignmentOrderable
}
// NewSetOrderRequestBody instantiates a new setOrderRequestBody and sets the default values.
func NewSetOrderRequestBody()(*SetOrderRequestBody) {
    m := &SetOrderRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSetOrderRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetOrderRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSetOrderRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetOrderRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetOrderRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newAssignmentOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAssignmentOrderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewAssignmentOrder(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignmentOrderable))
        }
        return nil
    }
    return res
}
// GetNewAssignmentOrder gets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *SetOrderRequestBody) GetNewAssignmentOrder()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignmentOrderable) {
    if m == nil {
        return nil
    } else {
        return m.newAssignmentOrder
    }
}
// Serialize serializes information the current object
func (m *SetOrderRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SetOrderRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNewAssignmentOrder sets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *SetOrderRequestBody) SetNewAssignmentOrder(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignmentOrderable)() {
    if m != nil {
        m.newAssignmentOrder = value
    }
}
