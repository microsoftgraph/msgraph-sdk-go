package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkOnlineMeetingInfo 
type TeamworkOnlineMeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identifier of the calendar event associated with the meeting.
    calendarEventId *string;
    // The URL that users click to join or uniquely identify the meeting.
    joinWebUrl *string;
    // The organizer of the meeting.
    organizer *TeamworkUserIdentity;
}
// NewTeamworkOnlineMeetingInfo instantiates a new teamworkOnlineMeetingInfo and sets the default values.
func NewTeamworkOnlineMeetingInfo()(*TeamworkOnlineMeetingInfo) {
    m := &TeamworkOnlineMeetingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkOnlineMeetingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCalendarEventId gets the calendarEventId property value. The identifier of the calendar event associated with the meeting.
func (m *TeamworkOnlineMeetingInfo) GetCalendarEventId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.calendarEventId
    }
}
// GetJoinWebUrl gets the joinWebUrl property value. The URL that users click to join or uniquely identify the meeting.
func (m *TeamworkOnlineMeetingInfo) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// GetOrganizer gets the organizer property value. The organizer of the meeting.
func (m *TeamworkOnlineMeetingInfo) GetOrganizer()(*TeamworkUserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.organizer
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkOnlineMeetingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calendarEventId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendarEventId(val)
        }
        return nil
    }
    res["joinWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
        }
        return nil
    }
    res["organizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkUserIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(*TeamworkUserIdentity))
        }
        return nil
    }
    return res
}
func (m *TeamworkOnlineMeetingInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkOnlineMeetingInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("calendarEventId", m.GetCalendarEventId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("organizer", m.GetOrganizer())
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
func (m *TeamworkOnlineMeetingInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCalendarEventId sets the calendarEventId property value. The identifier of the calendar event associated with the meeting.
func (m *TeamworkOnlineMeetingInfo) SetCalendarEventId(value *string)() {
    if m != nil {
        m.calendarEventId = value
    }
}
// SetJoinWebUrl sets the joinWebUrl property value. The URL that users click to join or uniquely identify the meeting.
func (m *TeamworkOnlineMeetingInfo) SetJoinWebUrl(value *string)() {
    if m != nil {
        m.joinWebUrl = value
    }
}
// SetOrganizer sets the organizer property value. The organizer of the meeting.
func (m *TeamworkOnlineMeetingInfo) SetOrganizer(value *TeamworkUserIdentity)() {
    if m != nil {
        m.organizer = value
    }
}
