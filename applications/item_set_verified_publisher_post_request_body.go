package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSetVerifiedPublisherPostRequestBody provides operations to call the setVerifiedPublisher method.
type ItemSetVerifiedPublisherPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The verifiedPublisherId property
    verifiedPublisherId *string
}
// NewItemSetVerifiedPublisherPostRequestBody instantiates a new ItemSetVerifiedPublisherPostRequestBody and sets the default values.
func NewItemSetVerifiedPublisherPostRequestBody()(*ItemSetVerifiedPublisherPostRequestBody) {
    m := &ItemSetVerifiedPublisherPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSetVerifiedPublisherPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSetVerifiedPublisherPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSetVerifiedPublisherPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSetVerifiedPublisherPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSetVerifiedPublisherPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["verifiedPublisherId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisherId(val)
        }
        return nil
    }
    return res
}
// GetVerifiedPublisherId gets the verifiedPublisherId property value. The verifiedPublisherId property
func (m *ItemSetVerifiedPublisherPostRequestBody) GetVerifiedPublisherId()(*string) {
    return m.verifiedPublisherId
}
// Serialize serializes information the current object
func (m *ItemSetVerifiedPublisherPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemSetVerifiedPublisherPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetVerifiedPublisherId sets the verifiedPublisherId property value. The verifiedPublisherId property
func (m *ItemSetVerifiedPublisherPostRequestBody) SetVerifiedPublisherId(value *string)() {
    m.verifiedPublisherId = value
}
