package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationsItemSetVerifiedPublisherPostRequestBody provides operations to call the setVerifiedPublisher method.
type ApplicationsItemSetVerifiedPublisherPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The verifiedPublisherId property
    verifiedPublisherId *string
}
// NewApplicationsItemSetVerifiedPublisherPostRequestBody instantiates a new ApplicationsItemSetVerifiedPublisherPostRequestBody and sets the default values.
func NewApplicationsItemSetVerifiedPublisherPostRequestBody()(*ApplicationsItemSetVerifiedPublisherPostRequestBody) {
    m := &ApplicationsItemSetVerifiedPublisherPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApplicationsItemSetVerifiedPublisherPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationsItemSetVerifiedPublisherPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicationsItemSetVerifiedPublisherPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplicationsItemSetVerifiedPublisherPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationsItemSetVerifiedPublisherPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["verifiedPublisherId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVerifiedPublisherId)
    return res
}
// GetVerifiedPublisherId gets the verifiedPublisherId property value. The verifiedPublisherId property
func (m *ApplicationsItemSetVerifiedPublisherPostRequestBody) GetVerifiedPublisherId()(*string) {
    return m.verifiedPublisherId
}
// Serialize serializes information the current object
func (m *ApplicationsItemSetVerifiedPublisherPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("verifiedPublisherId", m.GetVerifiedPublisherId())
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
func (m *ApplicationsItemSetVerifiedPublisherPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetVerifiedPublisherId sets the verifiedPublisherId property value. The verifiedPublisherId property
func (m *ApplicationsItemSetVerifiedPublisherPostRequestBody) SetVerifiedPublisherId(value *string)() {
    m.verifiedPublisherId = value
}
