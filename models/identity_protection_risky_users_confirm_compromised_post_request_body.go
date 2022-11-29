package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody provides operations to call the confirmCompromised method.
type IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The userIds property
    userIds []string
}
// NewIdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody instantiates a new IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody and sets the default values.
func NewIdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody()(*IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody) {
    m := &IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityProtectionRiskyUsersConfirmCompromisedPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityProtectionRiskyUsersConfirmCompromisedPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["userIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetUserIds)
    return res
}
// GetUserIds gets the userIds property value. The userIds property
func (m *IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody) GetUserIds()([]string) {
    return m.userIds
}
// Serialize serializes information the current object
func (m *IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUserIds() != nil {
        err := writer.WriteCollectionOfStringValues("userIds", m.GetUserIds())
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
func (m *IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetUserIds sets the userIds property value. The userIds property
func (m *IdentityProtectionRiskyUsersConfirmCompromisedPostRequestBody) SetUserIds(value []string)() {
    m.userIds = value
}
