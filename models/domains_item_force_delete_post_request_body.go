package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainsItemForceDeletePostRequestBody provides operations to call the forceDelete method.
type DomainsItemForceDeletePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The disableUserAccounts property
    disableUserAccounts *bool
}
// NewDomainsItemForceDeletePostRequestBody instantiates a new DomainsItemForceDeletePostRequestBody and sets the default values.
func NewDomainsItemForceDeletePostRequestBody()(*DomainsItemForceDeletePostRequestBody) {
    m := &DomainsItemForceDeletePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDomainsItemForceDeletePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDomainsItemForceDeletePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDomainsItemForceDeletePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DomainsItemForceDeletePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisableUserAccounts gets the disableUserAccounts property value. The disableUserAccounts property
func (m *DomainsItemForceDeletePostRequestBody) GetDisableUserAccounts()(*bool) {
    return m.disableUserAccounts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DomainsItemForceDeletePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["disableUserAccounts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDisableUserAccounts)
    return res
}
// Serialize serializes information the current object
func (m *DomainsItemForceDeletePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DomainsItemForceDeletePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisableUserAccounts sets the disableUserAccounts property value. The disableUserAccounts property
func (m *DomainsItemForceDeletePostRequestBody) SetDisableUserAccounts(value *bool)() {
    m.disableUserAccounts = value
}
