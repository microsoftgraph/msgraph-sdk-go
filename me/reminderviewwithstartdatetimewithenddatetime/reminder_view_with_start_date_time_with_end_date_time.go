package reminderviewwithstartdatetimewithenddatetime

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// 
type ReminderViewWithStartDateTimeWithEndDateTime struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
    changeKey *string;
    // The date, time and time zone that the event ends.
    eventEndTime *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone;
    // The unique ID of the event. Read only.
    eventId *string;
    // The location of the event.
    eventLocation *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location;
    // The date, time, and time zone that the event starts.
    eventStartTime *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone;
    // The text of the event's subject line.
    eventSubject *string;
    // The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
    eventWebLink *string;
    // The date, time, and time zone that the reminder is set to occur.
    reminderFireTime *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone;
}
// Instantiates a new reminderViewWithStartDateTimeWithEndDateTime and sets the default values.
func NewReminderViewWithStartDateTimeWithEndDateTime()(*ReminderViewWithStartDateTimeWithEndDateTime) {
    m := &ReminderViewWithStartDateTimeWithEndDateTime{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the changeKey property value. Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetChangeKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeKey
    }
}
// Gets the eventEndTime property value. The date, time and time zone that the event ends.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventEndTime()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.eventEndTime
    }
}
// Gets the eventId property value. The unique ID of the event. Read only.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventId
    }
}
// Gets the eventLocation property value. The location of the event.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventLocation()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location) {
    if m == nil {
        return nil
    } else {
        return m.eventLocation
    }
}
// Gets the eventStartTime property value. The date, time, and time zone that the event starts.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventStartTime()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.eventStartTime
    }
}
// Gets the eventSubject property value. The text of the event's subject line.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventSubject
    }
}
// Gets the eventWebLink property value. The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetEventWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventWebLink
    }
}
// Gets the reminderFireTime property value. The date, time, and time zone that the reminder is set to occur.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetReminderFireTime()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.reminderFireTime
    }
}
// The deserialization information for the current model
func (m *ReminderViewWithStartDateTimeWithEndDateTime) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventEndTime(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventLocation(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location))
        }
        return nil
    }
    res["eventStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventStartTime(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderFireTime(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone))
        }
        return nil
    }
    return res
}
func (m *ReminderViewWithStartDateTimeWithEndDateTime) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ReminderViewWithStartDateTimeWithEndDateTime) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the changeKey property value. Identifies the version of the reminder. Every time the reminder is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object.
// Parameters:
//  - value : Value to set for the changeKey property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetChangeKey(value *string)() {
    m.changeKey = value
}
// Sets the eventEndTime property value. The date, time and time zone that the event ends.
// Parameters:
//  - value : Value to set for the eventEndTime property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventEndTime(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone)() {
    m.eventEndTime = value
}
// Sets the eventId property value. The unique ID of the event. Read only.
// Parameters:
//  - value : Value to set for the eventId property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventId(value *string)() {
    m.eventId = value
}
// Sets the eventLocation property value. The location of the event.
// Parameters:
//  - value : Value to set for the eventLocation property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventLocation(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location)() {
    m.eventLocation = value
}
// Sets the eventStartTime property value. The date, time, and time zone that the event starts.
// Parameters:
//  - value : Value to set for the eventStartTime property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventStartTime(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone)() {
    m.eventStartTime = value
}
// Sets the eventSubject property value. The text of the event's subject line.
// Parameters:
//  - value : Value to set for the eventSubject property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventSubject(value *string)() {
    m.eventSubject = value
}
// Sets the eventWebLink property value. The URL to open the event in Outlook on the web.The event will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
// Parameters:
//  - value : Value to set for the eventWebLink property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetEventWebLink(value *string)() {
    m.eventWebLink = value
}
// Sets the reminderFireTime property value. The date, time, and time zone that the reminder is set to occur.
// Parameters:
//  - value : Value to set for the reminderFireTime property.
func (m *ReminderViewWithStartDateTimeWithEndDateTime) SetReminderFireTime(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone)() {
    m.reminderFireTime = value
}
