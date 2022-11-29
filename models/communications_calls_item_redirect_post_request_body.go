package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallsItemRedirectPostRequestBody provides operations to call the redirect method.
type CommunicationsCallsItemRedirectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The callbackUri property
    callbackUri *string
    // The targets property
    targets []InvitationParticipantInfoable
    // The timeout property
    timeout *int32
}
// NewCommunicationsCallsItemRedirectPostRequestBody instantiates a new CommunicationsCallsItemRedirectPostRequestBody and sets the default values.
func NewCommunicationsCallsItemRedirectPostRequestBody()(*CommunicationsCallsItemRedirectPostRequestBody) {
    m := &CommunicationsCallsItemRedirectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommunicationsCallsItemRedirectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsCallsItemRedirectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsCallsItemRedirectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsCallsItemRedirectPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *CommunicationsCallsItemRedirectPostRequestBody) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsCallsItemRedirectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callbackUri"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCallbackUri)
    res["targets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateInvitationParticipantInfoFromDiscriminatorValue , m.SetTargets)
    res["timeout"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTimeout)
    return res
}
// GetTargets gets the targets property value. The targets property
func (m *CommunicationsCallsItemRedirectPostRequestBody) GetTargets()([]InvitationParticipantInfoable) {
    return m.targets
}
// GetTimeout gets the timeout property value. The timeout property
func (m *CommunicationsCallsItemRedirectPostRequestBody) GetTimeout()(*int32) {
    return m.timeout
}
// Serialize serializes information the current object
func (m *CommunicationsCallsItemRedirectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    if m.GetTargets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetTargets())
        err := writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeout", m.GetTimeout())
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
func (m *CommunicationsCallsItemRedirectPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *CommunicationsCallsItemRedirectPostRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetTargets sets the targets property value. The targets property
func (m *CommunicationsCallsItemRedirectPostRequestBody) SetTargets(value []InvitationParticipantInfoable)() {
    m.targets = value
}
// SetTimeout sets the timeout property value. The timeout property
func (m *CommunicationsCallsItemRedirectPostRequestBody) SetTimeout(value *int32)() {
    m.timeout = value
}
