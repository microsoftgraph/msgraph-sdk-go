package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbooksItemRestorePostRequestBody provides operations to call the restore method.
type WorkbooksItemRestorePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name property
    name *string
    // The parentReference property
    parentReference ItemReferenceable
}
// NewWorkbooksItemRestorePostRequestBody instantiates a new WorkbooksItemRestorePostRequestBody and sets the default values.
func NewWorkbooksItemRestorePostRequestBody()(*WorkbooksItemRestorePostRequestBody) {
    m := &WorkbooksItemRestorePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbooksItemRestorePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbooksItemRestorePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbooksItemRestorePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbooksItemRestorePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbooksItemRestorePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["parentReference"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateItemReferenceFromDiscriminatorValue , m.SetParentReference)
    return res
}
// GetName gets the name property value. The name property
func (m *WorkbooksItemRestorePostRequestBody) GetName()(*string) {
    return m.name
}
// GetParentReference gets the parentReference property value. The parentReference property
func (m *WorkbooksItemRestorePostRequestBody) GetParentReference()(ItemReferenceable) {
    return m.parentReference
}
// Serialize serializes information the current object
func (m *WorkbooksItemRestorePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parentReference", m.GetParentReference())
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
func (m *WorkbooksItemRestorePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *WorkbooksItemRestorePostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetParentReference sets the parentReference property value. The parentReference property
func (m *WorkbooksItemRestorePostRequestBody) SetParentReference(value ItemReferenceable)() {
    m.parentReference = value
}
