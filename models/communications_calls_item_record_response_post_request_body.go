package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallsItemRecordResponsePostRequestBody provides operations to call the recordResponse method.
type CommunicationsCallsItemRecordResponsePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The bargeInAllowed property
    bargeInAllowed *bool
    // The clientContext property
    clientContext *string
    // The initialSilenceTimeoutInSeconds property
    initialSilenceTimeoutInSeconds *int32
    // The maxRecordDurationInSeconds property
    maxRecordDurationInSeconds *int32
    // The maxSilenceTimeoutInSeconds property
    maxSilenceTimeoutInSeconds *int32
    // The playBeep property
    playBeep *bool
    // The prompts property
    prompts []Promptable
    // The stopTones property
    stopTones []string
}
// NewCommunicationsCallsItemRecordResponsePostRequestBody instantiates a new CommunicationsCallsItemRecordResponsePostRequestBody and sets the default values.
func NewCommunicationsCallsItemRecordResponsePostRequestBody()(*CommunicationsCallsItemRecordResponsePostRequestBody) {
    m := &CommunicationsCallsItemRecordResponsePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCommunicationsCallsItemRecordResponsePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsCallsItemRecordResponsePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsCallsItemRecordResponsePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBargeInAllowed gets the bargeInAllowed property value. The bargeInAllowed property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetBargeInAllowed()(*bool) {
    return m.bargeInAllowed
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bargeInAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetBargeInAllowed)
    res["clientContext"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetClientContext)
    res["initialSilenceTimeoutInSeconds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetInitialSilenceTimeoutInSeconds)
    res["maxRecordDurationInSeconds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaxRecordDurationInSeconds)
    res["maxSilenceTimeoutInSeconds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetMaxSilenceTimeoutInSeconds)
    res["playBeep"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPlayBeep)
    res["prompts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePromptFromDiscriminatorValue , m.SetPrompts)
    res["stopTones"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetStopTones)
    return res
}
// GetInitialSilenceTimeoutInSeconds gets the initialSilenceTimeoutInSeconds property value. The initialSilenceTimeoutInSeconds property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetInitialSilenceTimeoutInSeconds()(*int32) {
    return m.initialSilenceTimeoutInSeconds
}
// GetMaxRecordDurationInSeconds gets the maxRecordDurationInSeconds property value. The maxRecordDurationInSeconds property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetMaxRecordDurationInSeconds()(*int32) {
    return m.maxRecordDurationInSeconds
}
// GetMaxSilenceTimeoutInSeconds gets the maxSilenceTimeoutInSeconds property value. The maxSilenceTimeoutInSeconds property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetMaxSilenceTimeoutInSeconds()(*int32) {
    return m.maxSilenceTimeoutInSeconds
}
// GetPlayBeep gets the playBeep property value. The playBeep property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetPlayBeep()(*bool) {
    return m.playBeep
}
// GetPrompts gets the prompts property value. The prompts property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetPrompts()([]Promptable) {
    return m.prompts
}
// GetStopTones gets the stopTones property value. The stopTones property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) GetStopTones()([]string) {
    return m.stopTones
}
// Serialize serializes information the current object
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("bargeInAllowed", m.GetBargeInAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("initialSilenceTimeoutInSeconds", m.GetInitialSilenceTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxRecordDurationInSeconds", m.GetMaxRecordDurationInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxSilenceTimeoutInSeconds", m.GetMaxSilenceTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("playBeep", m.GetPlayBeep())
        if err != nil {
            return err
        }
    }
    if m.GetPrompts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPrompts())
        err := writer.WriteCollectionOfObjectValues("prompts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStopTones() != nil {
        err := writer.WriteCollectionOfStringValues("stopTones", m.GetStopTones())
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
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBargeInAllowed sets the bargeInAllowed property value. The bargeInAllowed property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetBargeInAllowed(value *bool)() {
    m.bargeInAllowed = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
// SetInitialSilenceTimeoutInSeconds sets the initialSilenceTimeoutInSeconds property value. The initialSilenceTimeoutInSeconds property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetInitialSilenceTimeoutInSeconds(value *int32)() {
    m.initialSilenceTimeoutInSeconds = value
}
// SetMaxRecordDurationInSeconds sets the maxRecordDurationInSeconds property value. The maxRecordDurationInSeconds property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetMaxRecordDurationInSeconds(value *int32)() {
    m.maxRecordDurationInSeconds = value
}
// SetMaxSilenceTimeoutInSeconds sets the maxSilenceTimeoutInSeconds property value. The maxSilenceTimeoutInSeconds property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetMaxSilenceTimeoutInSeconds(value *int32)() {
    m.maxSilenceTimeoutInSeconds = value
}
// SetPlayBeep sets the playBeep property value. The playBeep property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetPlayBeep(value *bool)() {
    m.playBeep = value
}
// SetPrompts sets the prompts property value. The prompts property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetPrompts(value []Promptable)() {
    m.prompts = value
}
// SetStopTones sets the stopTones property value. The stopTones property
func (m *CommunicationsCallsItemRecordResponsePostRequestBody) SetStopTones(value []string)() {
    m.stopTones = value
}
