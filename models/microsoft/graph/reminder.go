package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Reminder provides operations to call the reminderView method.
type Reminder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
    changeKey *string;
    // The date, time and time zone that the event ends.
    eventEndTime DateTimeTimeZoneable;
    // The unique ID of the event. Read only.
    eventId *string;
    // The location of the event.
    eventLocation Locationable;
    // The date, time, and time zone that the event starts.
    eventStartTime DateTimeTimeZoneable;
    // The text of the event's subject line.
    eventSubject *string;
    // The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
    eventWebLink *string;
    // The date, time, and time zone that the reminder is set to occur.
    reminderFireTime DateTimeTimeZoneable;
}
// NewReminder instantiates a new reminder and sets the default values.
func NewReminder()(*Reminder) {
    m := &Reminder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateReminderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReminderFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewReminder(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Reminder) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChangeKey gets the changeKey property value. Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
func (m *Reminder) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
// GetEventEndTime gets the eventEndTime property value. The date, time and time zone that the event ends.
func (m *Reminder) GetEventEndTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.eventEndTime
    }
}
// GetEventId gets the eventId property value. The unique ID of the event. Read only.
func (m *Reminder) GetEventId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventId
    }
}
// GetEventLocation gets the eventLocation property value. The location of the event.
func (m *Reminder) GetEventLocation()(Locationable) {
    if m == nil {
        return nil
    } else {
        return m.eventLocation
    }
}
// GetEventStartTime gets the eventStartTime property value. The date, time, and time zone that the event starts.
func (m *Reminder) GetEventStartTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.eventStartTime
    }
}
// GetEventSubject gets the eventSubject property value. The text of the event's subject line.
func (m *Reminder) GetEventSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventSubject
    }
}
// GetEventWebLink gets the eventWebLink property value. The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
func (m *Reminder) GetEventWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventWebLink
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Reminder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["changeKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["eventEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventEndTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["eventId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventId(val)
        }
        return nil
    }
    res["eventLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventLocation(val.(Locationable))
        }
        return nil
    }
    res["eventStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventStartTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["eventSubject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventSubject(val)
        }
        return nil
    }
    res["eventWebLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventWebLink(val)
        }
        return nil
    }
    res["reminderFireTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderFireTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetReminderFireTime gets the reminderFireTime property value. The date, time, and time zone that the reminder is set to occur.
func (m *Reminder) GetReminderFireTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.reminderFireTime
    }
}
func (m *Reminder) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Reminder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventEndTime", m.GetEventEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventId", m.GetEventId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventLocation", m.GetEventLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("eventStartTime", m.GetEventStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventSubject", m.GetEventSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventWebLink", m.GetEventWebLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reminderFireTime", m.GetReminderFireTime())
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
func (m *Reminder) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChangeKey sets the changeKey property value. Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
func (m *Reminder) SetChangeKey(value *string)() {
    if m != nil {
        m.changeKey = value
    }
}
// SetEventEndTime sets the eventEndTime property value. The date, time and time zone that the event ends.
func (m *Reminder) SetEventEndTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.eventEndTime = value
    }
}
// SetEventId sets the eventId property value. The unique ID of the event. Read only.
func (m *Reminder) SetEventId(value *string)() {
    if m != nil {
        m.eventId = value
    }
}
// SetEventLocation sets the eventLocation property value. The location of the event.
func (m *Reminder) SetEventLocation(value Locationable)() {
    if m != nil {
        m.eventLocation = value
    }
}
// SetEventStartTime sets the eventStartTime property value. The date, time, and time zone that the event starts.
func (m *Reminder) SetEventStartTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.eventStartTime = value
    }
}
// SetEventSubject sets the eventSubject property value. The text of the event's subject line.
func (m *Reminder) SetEventSubject(value *string)() {
    if m != nil {
        m.eventSubject = value
    }
}
// SetEventWebLink sets the eventWebLink property value. The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
func (m *Reminder) SetEventWebLink(value *string)() {
    if m != nil {
        m.eventWebLink = value
    }
}
// SetReminderFireTime sets the reminderFireTime property value. The date, time, and time zone that the reminder is set to occur.
func (m *Reminder) SetReminderFireTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.reminderFireTime = value
    }
}
