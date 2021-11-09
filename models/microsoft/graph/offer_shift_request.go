package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OfferShiftRequest struct {
    ScheduleChangeRequest
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    recipientActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Custom message sent by recipient of the offer shift request.
    recipientActionMessage *string;
    // User ID of the recipient of the offer shift request.
    recipientUserId *string;
    // User ID of the sender of the offer shift request.
    senderShiftId *string;
}
// Instantiates a new offerShiftRequest and sets the default values.
func NewOfferShiftRequest()(*OfferShiftRequest) {
    m := &OfferShiftRequest{
        ScheduleChangeRequest: *NewScheduleChangeRequest(),
    }
    return m
}
// Gets the recipientActionDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *OfferShiftRequest) GetRecipientActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.recipientActionDateTime
    }
}
// Gets the recipientActionMessage property value. Custom message sent by recipient of the offer shift request.
func (m *OfferShiftRequest) GetRecipientActionMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientActionMessage
    }
}
// Gets the recipientUserId property value. User ID of the recipient of the offer shift request.
func (m *OfferShiftRequest) GetRecipientUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientUserId
    }
}
// Gets the senderShiftId property value. User ID of the sender of the offer shift request.
func (m *OfferShiftRequest) GetSenderShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.senderShiftId
    }
}
// The deserialization information for the current model
func (m *OfferShiftRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ScheduleChangeRequest.GetFieldDeserializers()
    res["recipientActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientActionDateTime(val)
        }
        return nil
    }
    res["recipientActionMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientActionMessage(val)
        }
        return nil
    }
    res["recipientUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientUserId(val)
        }
        return nil
    }
    res["senderShiftId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderShiftId(val)
        }
        return nil
    }
    return res
}
func (m *OfferShiftRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the recipientActionDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the recipientActionDateTime property.
func (m *OfferShiftRequest) SetRecipientActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recipientActionDateTime = value
}
// Sets the recipientActionMessage property value. Custom message sent by recipient of the offer shift request.
// Parameters:
//  - value : Value to set for the recipientActionMessage property.
func (m *OfferShiftRequest) SetRecipientActionMessage(value *string)() {
    m.recipientActionMessage = value
}
// Sets the recipientUserId property value. User ID of the recipient of the offer shift request.
// Parameters:
//  - value : Value to set for the recipientUserId property.
func (m *OfferShiftRequest) SetRecipientUserId(value *string)() {
    m.recipientUserId = value
}
// Sets the senderShiftId property value. User ID of the sender of the offer shift request.
// Parameters:
//  - value : Value to set for the senderShiftId property.
func (m *OfferShiftRequest) SetSenderShiftId(value *string)() {
    m.senderShiftId = value
}
