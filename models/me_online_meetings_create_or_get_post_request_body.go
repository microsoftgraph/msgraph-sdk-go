package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeOnlineMeetingsCreateOrGetPostRequestBody provides operations to call the createOrGet method.
type MeOnlineMeetingsCreateOrGetPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The chatInfo property
    chatInfo ChatInfoable
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The externalId property
    externalId *string
    // The participants property
    participants MeetingParticipantsable
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject property
    subject *string
}
// NewMeOnlineMeetingsCreateOrGetPostRequestBody instantiates a new MeOnlineMeetingsCreateOrGetPostRequestBody and sets the default values.
func NewMeOnlineMeetingsCreateOrGetPostRequestBody()(*MeOnlineMeetingsCreateOrGetPostRequestBody) {
    m := &MeOnlineMeetingsCreateOrGetPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeOnlineMeetingsCreateOrGetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeOnlineMeetingsCreateOrGetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeOnlineMeetingsCreateOrGetPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetChatInfo gets the chatInfo property value. The chatInfo property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) GetChatInfo()(ChatInfoable) {
    return m.chatInfo
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetExternalId gets the externalId property value. The externalId property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["chatInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateChatInfoFromDiscriminatorValue , m.SetChatInfo)
    res["endDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEndDateTime)
    res["externalId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExternalId)
    res["participants"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMeetingParticipantsFromDiscriminatorValue , m.SetParticipants)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    res["subject"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubject)
    return res
}
// GetParticipants gets the participants property value. The participants property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) GetParticipants()(MeetingParticipantsable) {
    return m.participants
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetSubject gets the subject property value. The subject property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("chatInfo", m.GetChatInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subject", m.GetSubject())
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
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetChatInfo sets the chatInfo property value. The chatInfo property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) SetChatInfo(value ChatInfoable)() {
    m.chatInfo = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetExternalId sets the externalId property value. The externalId property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) SetExternalId(value *string)() {
    m.externalId = value
}
// SetParticipants sets the participants property value. The participants property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) SetParticipants(value MeetingParticipantsable)() {
    m.participants = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetSubject sets the subject property value. The subject property
func (m *MeOnlineMeetingsCreateOrGetPostRequestBody) SetSubject(value *string)() {
    m.subject = value
}
