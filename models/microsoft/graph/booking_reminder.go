package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingReminder 
type BookingReminder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The message in the reminder.
    message *string;
    // The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
    offset *string;
    // The persons who should receive the reminder. Possible values are: allAttendees, staff, customer, unknownFutureValue.
    recipients *BookingReminderRecipients;
}
// NewBookingReminder instantiates a new bookingReminder and sets the default values.
func NewBookingReminder()(*BookingReminder) {
    m := &BookingReminder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingReminder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMessage gets the message property value. The message in the reminder.
func (m *BookingReminder) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// GetOffset gets the offset property value. The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
func (m *BookingReminder) GetOffset()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
// GetRecipients gets the recipients property value. The persons who should receive the reminder. Possible values are: allAttendees, staff, customer, unknownFutureValue.
func (m *BookingReminder) GetRecipients()(*BookingReminderRecipients) {
    if m == nil {
        return nil
    } else {
        return m.recipients
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingReminder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["offset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffset(val)
        }
        return nil
    }
    res["recipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingReminderRecipients)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(BookingReminderRecipients)
            m.SetRecipients(&cast)
        }
        return nil
    }
    return res
}
func (m *BookingReminder) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingReminder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("offset", m.GetOffset())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := m.GetRecipients().String()
        err := writer.WriteStringValue("recipients", &cast)
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
func (m *BookingReminder) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMessage sets the message property value. The message in the reminder.
func (m *BookingReminder) SetMessage(value *string)() {
    if m != nil {
        m.message = value
    }
}
// SetOffset sets the offset property value. The amount of time before the start of an appointment that the reminder should be sent. It's denoted in ISO 8601 format.
func (m *BookingReminder) SetOffset(value *string)() {
    if m != nil {
        m.offset = value
    }
}
// SetRecipients sets the recipients property value. The persons who should receive the reminder. Possible values are: allAttendees, staff, customer, unknownFutureValue.
func (m *BookingReminder) SetRecipients(value *BookingReminderRecipients)() {
    if m != nil {
        m.recipients = value
    }
}
