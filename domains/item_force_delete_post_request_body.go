package domains

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemForceDeletePostRequestBody provides operations to call the forceDelete method.
type ItemForceDeletePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The disableUserAccounts property
    disableUserAccounts *bool
}
// NewItemForceDeletePostRequestBody instantiates a new ItemForceDeletePostRequestBody and sets the default values.
func NewItemForceDeletePostRequestBody()(*ItemForceDeletePostRequestBody) {
    m := &ItemForceDeletePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemForceDeletePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemForceDeletePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemForceDeletePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemForceDeletePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisableUserAccounts gets the disableUserAccounts property value. The disableUserAccounts property
func (m *ItemForceDeletePostRequestBody) GetDisableUserAccounts()(*bool) {
    return m.disableUserAccounts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemForceDeletePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["disableUserAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableUserAccounts(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemForceDeletePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("disableUserAccounts", m.GetDisableUserAccounts())
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
func (m *ItemForceDeletePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisableUserAccounts sets the disableUserAccounts property value. The disableUserAccounts property
func (m *ItemForceDeletePostRequestBody) SetDisableUserAccounts(value *bool)() {
    m.disableUserAccounts = value
}
