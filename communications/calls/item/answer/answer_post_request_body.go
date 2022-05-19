package answer

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// AnswerPostRequestBody provides operations to call the answer method.
type AnswerPostRequestBody struct {
    // The acceptedModalities property
    acceptedModalities []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The callbackUri property
    callbackUri *string
    // The callOptions property
    callOptions iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable
    // The mediaConfig property
    mediaConfig iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable
    // The participantCapacity property
    participantCapacity *int32
}
// NewAnswerPostRequestBody instantiates a new answerPostRequestBody and sets the default values.
func NewAnswerPostRequestBody()(*AnswerPostRequestBody) {
    m := &AnswerPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAnswerPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAnswerPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnswerPostRequestBody(), nil
}
// GetAcceptedModalities gets the acceptedModalities property value. The acceptedModalities property
func (m *AnswerPostRequestBody) GetAcceptedModalities()([]string) {
    if m == nil {
        return nil
    } else {
        return m.acceptedModalities
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *AnswerPostRequestBody) GetCallbackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callbackUri
    }
}
// GetCallOptions gets the callOptions property value. The callOptions property
func (m *AnswerPostRequestBody) GetCallOptions()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable) {
    if m == nil {
        return nil
    } else {
        return m.callOptions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AnswerPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptedModalities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAcceptedModalities(res)
        }
        return nil
    }
    res["callbackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["callOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIncomingCallOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallOptions(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable))
        }
        return nil
    }
    res["mediaConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMediaConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaConfig(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable))
        }
        return nil
    }
    res["participantCapacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantCapacity(val)
        }
        return nil
    }
    return res
}
// GetMediaConfig gets the mediaConfig property value. The mediaConfig property
func (m *AnswerPostRequestBody) GetMediaConfig()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable) {
    if m == nil {
        return nil
    } else {
        return m.mediaConfig
    }
}
// GetParticipantCapacity gets the participantCapacity property value. The participantCapacity property
func (m *AnswerPostRequestBody) GetParticipantCapacity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.participantCapacity
    }
}
// Serialize serializes information the current object
func (m *AnswerPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAcceptedModalities() != nil {
        err := writer.WriteCollectionOfStringValues("acceptedModalities", m.GetAcceptedModalities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("callOptions", m.GetCallOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mediaConfig", m.GetMediaConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("participantCapacity", m.GetParticipantCapacity())
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
// SetAcceptedModalities sets the acceptedModalities property value. The acceptedModalities property
func (m *AnswerPostRequestBody) SetAcceptedModalities(value []string)() {
    if m != nil {
        m.acceptedModalities = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *AnswerPostRequestBody) SetCallbackUri(value *string)() {
    if m != nil {
        m.callbackUri = value
    }
}
// SetCallOptions sets the callOptions property value. The callOptions property
func (m *AnswerPostRequestBody) SetCallOptions(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable)() {
    if m != nil {
        m.callOptions = value
    }
}
// SetMediaConfig sets the mediaConfig property value. The mediaConfig property
func (m *AnswerPostRequestBody) SetMediaConfig(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable)() {
    if m != nil {
        m.mediaConfig = value
    }
}
// SetParticipantCapacity sets the participantCapacity property value. The participantCapacity property
func (m *AnswerPostRequestBody) SetParticipantCapacity(value *int32)() {
    if m != nil {
        m.participantCapacity = value
    }
}
