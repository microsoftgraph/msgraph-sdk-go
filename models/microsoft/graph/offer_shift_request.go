package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OfferShiftRequest provides operations to manage the collection of group entities.
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
// NewOfferShiftRequest instantiates a new offerShiftRequest and sets the default values.
func NewOfferShiftRequest()(*OfferShiftRequest) {
    m := &OfferShiftRequest{
        ScheduleChangeRequest: *NewScheduleChangeRequest(),
    }
    return m
}
// CreateOfferShiftRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfferShiftRequestFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOfferShiftRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetRecipientActionDateTime gets the recipientActionDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *OfferShiftRequest) GetRecipientActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.recipientActionDateTime
    }
}
// GetRecipientActionMessage gets the recipientActionMessage property value. Custom message sent by recipient of the offer shift request.
func (m *OfferShiftRequest) GetRecipientActionMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientActionMessage
    }
}
// GetRecipientUserId gets the recipientUserId property value. User ID of the recipient of the offer shift request.
func (m *OfferShiftRequest) GetRecipientUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recipientUserId
    }
}
// GetSenderShiftId gets the senderShiftId property value. User ID of the sender of the offer shift request.
func (m *OfferShiftRequest) GetSenderShiftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.senderShiftId
    }
}
func (m *OfferShiftRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetRecipientActionDateTime sets the recipientActionDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *OfferShiftRequest) SetRecipientActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.recipientActionDateTime = value
    }
}
// SetRecipientActionMessage sets the recipientActionMessage property value. Custom message sent by recipient of the offer shift request.
func (m *OfferShiftRequest) SetRecipientActionMessage(value *string)() {
    if m != nil {
        m.recipientActionMessage = value
    }
}
// SetRecipientUserId sets the recipientUserId property value. User ID of the recipient of the offer shift request.
func (m *OfferShiftRequest) SetRecipientUserId(value *string)() {
    if m != nil {
        m.recipientUserId = value
    }
}
// SetSenderShiftId sets the senderShiftId property value. User ID of the sender of the offer shift request.
func (m *OfferShiftRequest) SetSenderShiftId(value *string)() {
    if m != nil {
        m.senderShiftId = value
    }
}
