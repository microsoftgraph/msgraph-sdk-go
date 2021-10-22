package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OfferShiftRequest struct {
    ScheduleChangeRequest
    recipientActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    recipientActionMessage *string;
    recipientUserId *string;
    senderShiftId *string;
}
func NewOfferShiftRequest()(*OfferShiftRequest) {
    m := &OfferShiftRequest{
        ScheduleChangeRequest: *NewScheduleChangeRequest(),
    }
    return m
}
func (m *OfferShiftRequest) GetRecipientActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.recipientActionDateTime
    }
}
func (m *OfferShiftRequest) GetRecipientActionMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientActionMessage
    }
}
func (m *OfferShiftRequest) GetRecipientUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientUserId
    }
}
func (m *OfferShiftRequest) GetSenderShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.senderShiftId
    }
}
func (m *OfferShiftRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ScheduleChangeRequest.GetFieldDeserializers()
    res["recipientActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRecipientActionDateTime(val)
        return nil
    }
    res["recipientActionMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecipientActionMessage(val)
        return nil
    }
    res["recipientUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecipientUserId(val)
        return nil
    }
    res["senderShiftId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSenderShiftId(val)
        return nil
    }
    return res
}
func (m *OfferShiftRequest) IsNil()(bool) {
    return m == nil
}
func (m *OfferShiftRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ScheduleChangeRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("recipientActionDateTime", m.GetRecipientActionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientActionMessage", m.GetRecipientActionMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientUserId", m.GetRecipientUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderShiftId", m.GetSenderShiftId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OfferShiftRequest) SetRecipientActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recipientActionDateTime = value
}
func (m *OfferShiftRequest) SetRecipientActionMessage(value *string)() {
    m.recipientActionMessage = value
}
func (m *OfferShiftRequest) SetRecipientUserId(value *string)() {
    m.recipientUserId = value
}
func (m *OfferShiftRequest) SetSenderShiftId(value *string)() {
    m.senderShiftId = value
}
