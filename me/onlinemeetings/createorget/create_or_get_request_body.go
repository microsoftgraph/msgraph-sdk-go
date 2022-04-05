package createorget

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CreateOrGetRequestBody provides operations to call the createOrGet method.
type CreateOrGetRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The chatInfo property
    chatInfo iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatInfoable;
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The externalId property
    externalId *string;
    // The participants property
    participants iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingParticipantsable;
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The subject property
    subject *string;
}
// NewCreateOrGetRequestBody instantiates a new createOrGetRequestBody and sets the default values.
func NewCreateOrGetRequestBody()(*CreateOrGetRequestBody) {
    m := &CreateOrGetRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCreateOrGetRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreateOrGetRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateOrGetRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateOrGetRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChatInfo gets the chatInfo property value. The chatInfo property
func (m *CreateOrGetRequestBody) GetChatInfo()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatInfoable) {
    if m == nil {
        return nil
    } else {
        return m.chatInfo
    }
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *CreateOrGetRequestBody) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetExternalId gets the externalId property value. The externalId property
func (m *CreateOrGetRequestBody) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateOrGetRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["chatInfo"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatInfoable))
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["participants"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMeetingParticipantsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipants(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingParticipantsable))
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetParticipants gets the participants property value. The participants property
func (m *CreateOrGetRequestBody) GetParticipants()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingParticipantsable) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *CreateOrGetRequestBody) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetSubject gets the subject property value. The subject property
func (m *CreateOrGetRequestBody) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Serialize serializes information the current object
func (m *CreateOrGetRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CreateOrGetRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChatInfo sets the chatInfo property value. The chatInfo property
func (m *CreateOrGetRequestBody) SetChatInfo(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatInfoable)() {
    if m != nil {
        m.chatInfo = value
    }
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *CreateOrGetRequestBody) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetExternalId sets the externalId property value. The externalId property
func (m *CreateOrGetRequestBody) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetParticipants sets the participants property value. The participants property
func (m *CreateOrGetRequestBody) SetParticipants(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingParticipantsable)() {
    if m != nil {
        m.participants = value
    }
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *CreateOrGetRequestBody) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetSubject sets the subject property value. The subject property
func (m *CreateOrGetRequestBody) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
