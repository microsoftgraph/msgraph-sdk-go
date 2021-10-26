package delta

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// 
type Delta struct {
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OutlookItem
    // true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
    allowNewTimeProposals *bool;
    // The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
    attachments []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attachment;
    // The collection of attendees for the event.
    attendees []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attendee;
    // The body of the message associated with the event. It can be in HTML or text format.
    body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBody;
    // The preview of the message associated with the event. It is in text format.
    bodyPreview *string;
    // The calendar that contains the event. Navigation property. Read-only.
    calendar *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar;
    // The date, time, and time zone that the event ends. By default, the end time is in UTC.
    end *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone;
    // The collection of open extensions defined for the event. Nullable.
    extensions []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension;
    // Set to true if the event has attachments.
    hasAttachments *bool;
    // When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
    hideAttendees *bool;
    // A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. Read-only.
    iCalUId *string;
    // 
    importance *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Importance;
    // The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
    instances []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event;
    // 
    isAllDay *bool;
    // 
    isCancelled *bool;
    // 
    isDraft *bool;
    // 
    isOnlineMeeting *bool;
    // 
    isOrganizer *bool;
    // 
    isReminderOn *bool;
    // 
    location *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location;
    // 
    locations []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location;
    // The collection of multi-value extended properties defined for the event. Read-only. Nullable.
    multiValueExtendedProperties []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty;
    // 
    onlineMeeting *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingInfo;
    // 
    onlineMeetingProvider *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingProviderType;
    // 
    onlineMeetingUrl *string;
    // 
    organizer *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient;
    // 
    originalEndTimeZone *string;
    // 
    originalStart *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    originalStartTimeZone *string;
    // 
    recurrence *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PatternedRecurrence;
    // 
    reminderMinutesBeforeStart *int32;
    // 
    responseRequested *bool;
    // 
    responseStatus *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResponseStatus;
    // 
    sensitivity *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Sensitivity;
    // 
    seriesMasterId *string;
    // 
    showAs *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.FreeBusyStatus;
    // The collection of single-value extended properties defined for the event. Read-only. Nullable.
    singleValueExtendedProperties []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty;
    // 
    start *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone;
    // 
    subject *string;
    // 
    transactionId *string;
    // 
    type_escaped *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EventType;
    // 
    webLink *string;
}
// Instantiates a new delta and sets the default values.
func NewDelta()(*Delta) {
    m := &Delta{
        OutlookItem: *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOutlookItem(),
    }
    return m
}
// Gets the allowNewTimeProposals property value. true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
func (m *Delta) GetAllowNewTimeProposals()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowNewTimeProposals
    }
}
// Gets the attachments property value. The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (m *Delta) GetAttachments()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// Gets the attendees property value. The collection of attendees for the event.
func (m *Delta) GetAttendees()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attendee) {
    if m == nil {
        return nil
    } else {
        return m.attendees
    }
}
// Gets the body property value. The body of the message associated with the event. It can be in HTML or text format.
func (m *Delta) GetBody()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// Gets the bodyPreview property value. The preview of the message associated with the event. It is in text format.
func (m *Delta) GetBodyPreview()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bodyPreview
    }
}
// Gets the calendar property value. The calendar that contains the event. Navigation property. Read-only.
func (m *Delta) GetCalendar()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// Gets the end property value. The date, time, and time zone that the event ends. By default, the end time is in UTC.
func (m *Delta) GetEnd()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// Gets the extensions property value. The collection of open extensions defined for the event. Nullable.
func (m *Delta) GetExtensions()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the hasAttachments property value. Set to true if the event has attachments.
func (m *Delta) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// Gets the hideAttendees property value. When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
func (m *Delta) GetHideAttendees()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideAttendees
    }
}
// Gets the iCalUId property value. A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. Read-only.
func (m *Delta) GetICalUId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.iCalUId
    }
}
// Gets the importance property value. 
func (m *Delta) GetImportance()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// Gets the instances property value. The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *Delta) GetInstances()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event) {
    if m == nil {
        return nil
    } else {
        return m.instances
    }
}
// Gets the isAllDay property value. 
func (m *Delta) GetIsAllDay()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAllDay
    }
}
// Gets the isCancelled property value. 
func (m *Delta) GetIsCancelled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCancelled
    }
}
// Gets the isDraft property value. 
func (m *Delta) GetIsDraft()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDraft
    }
}
// Gets the isOnlineMeeting property value. 
func (m *Delta) GetIsOnlineMeeting()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOnlineMeeting
    }
}
// Gets the isOrganizer property value. 
func (m *Delta) GetIsOrganizer()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizer
    }
}
// Gets the isReminderOn property value. 
func (m *Delta) GetIsReminderOn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReminderOn
    }
}
// Gets the location property value. 
func (m *Delta) GetLocation()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// Gets the locations property value. 
func (m *Delta) GetLocations()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
// Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the event. Read-only. Nullable.
func (m *Delta) GetMultiValueExtendedProperties()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// Gets the onlineMeeting property value. 
func (m *Delta) GetOnlineMeeting()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingInfo) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeeting
    }
}
// Gets the onlineMeetingProvider property value. 
func (m *Delta) GetOnlineMeetingProvider()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingProviderType) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetingProvider
    }
}
// Gets the onlineMeetingUrl property value. 
func (m *Delta) GetOnlineMeetingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetingUrl
    }
}
// Gets the organizer property value. 
func (m *Delta) GetOrganizer()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient) {
    if m == nil {
        return nil
    } else {
        return m.organizer
    }
}
// Gets the originalEndTimeZone property value. 
func (m *Delta) GetOriginalEndTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalEndTimeZone
    }
}
// Gets the originalStart property value. 
func (m *Delta) GetOriginalStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.originalStart
    }
}
// Gets the originalStartTimeZone property value. 
func (m *Delta) GetOriginalStartTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalStartTimeZone
    }
}
// Gets the recurrence property value. 
func (m *Delta) GetRecurrence()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// Gets the reminderMinutesBeforeStart property value. 
func (m *Delta) GetReminderMinutesBeforeStart()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.reminderMinutesBeforeStart
    }
}
// Gets the responseRequested property value. 
func (m *Delta) GetResponseRequested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.responseRequested
    }
}
// Gets the responseStatus property value. 
func (m *Delta) GetResponseStatus()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResponseStatus) {
    if m == nil {
        return nil
    } else {
        return m.responseStatus
    }
}
// Gets the sensitivity property value. 
func (m *Delta) GetSensitivity()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Sensitivity) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// Gets the seriesMasterId property value. 
func (m *Delta) GetSeriesMasterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.seriesMasterId
    }
}
// Gets the showAs property value. 
func (m *Delta) GetShowAs()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.FreeBusyStatus) {
    if m == nil {
        return nil
    } else {
        return m.showAs
    }
}
// Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the event. Read-only. Nullable.
func (m *Delta) GetSingleValueExtendedProperties()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// Gets the start property value. 
func (m *Delta) GetStart()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// Gets the subject property value. 
func (m *Delta) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Gets the transactionId property value. 
func (m *Delta) GetTransactionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.transactionId
    }
}
// Gets the type_escaped property value. 
func (m *Delta) GetType_escaped()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EventType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the webLink property value. 
func (m *Delta) GetWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webLink
    }
}
// The deserialization information for the current model
func (m *Delta) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["allowNewTimeProposals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowNewTimeProposals(val)
        return nil
    }
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAttachment() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attachment, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attachment))
        }
        m.SetAttachments(res)
        return nil
    }
    res["attendees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewAttendee() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attendee, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attendee))
        }
        m.SetAttendees(res)
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewItemBody() })
        if err != nil {
            return err
        }
        m.SetBody(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBody))
        return nil
    }
    res["bodyPreview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBodyPreview(val)
        return nil
    }
    res["calendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewCalendar() })
        if err != nil {
            return err
        }
        m.SetCalendar(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar))
        return nil
    }
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetEnd(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone))
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewExtension() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension))
        }
        m.SetExtensions(res)
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasAttachments(val)
        return nil
    }
    res["hideAttendees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHideAttendees(val)
        return nil
    }
    res["iCalUId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetICalUId(val)
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseImportance)
        if err != nil {
            return err
        }
        cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Importance)
        m.SetImportance(&cast)
        return nil
    }
    res["instances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEvent() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event))
        }
        m.SetInstances(res)
        return nil
    }
    res["isAllDay"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAllDay(val)
        return nil
    }
    res["isCancelled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCancelled(val)
        return nil
    }
    res["isDraft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDraft(val)
        return nil
    }
    res["isOnlineMeeting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOnlineMeeting(val)
        return nil
    }
    res["isOrganizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOrganizer(val)
        return nil
    }
    res["isReminderOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsReminderOn(val)
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewLocation() })
        if err != nil {
            return err
        }
        m.SetLocation(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location))
        return nil
    }
    res["locations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewLocation() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location))
        }
        m.SetLocations(res)
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty))
        }
        m.SetMultiValueExtendedProperties(res)
        return nil
    }
    res["onlineMeeting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewOnlineMeetingInfo() })
        if err != nil {
            return err
        }
        m.SetOnlineMeeting(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingInfo))
        return nil
    }
    res["onlineMeetingProvider"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseOnlineMeetingProviderType)
        if err != nil {
            return err
        }
        cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingProviderType)
        m.SetOnlineMeetingProvider(&cast)
        return nil
    }
    res["onlineMeetingUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOnlineMeetingUrl(val)
        return nil
    }
    res["organizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewRecipient() })
        if err != nil {
            return err
        }
        m.SetOrganizer(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient))
        return nil
    }
    res["originalEndTimeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginalEndTimeZone(val)
        return nil
    }
    res["originalStart"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetOriginalStart(val)
        return nil
    }
    res["originalStartTimeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginalStartTimeZone(val)
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        m.SetRecurrence(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PatternedRecurrence))
        return nil
    }
    res["reminderMinutesBeforeStart"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetReminderMinutesBeforeStart(val)
        return nil
    }
    res["responseRequested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetResponseRequested(val)
        return nil
    }
    res["responseStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewResponseStatus() })
        if err != nil {
            return err
        }
        m.SetResponseStatus(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResponseStatus))
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseSensitivity)
        if err != nil {
            return err
        }
        cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Sensitivity)
        m.SetSensitivity(&cast)
        return nil
    }
    res["seriesMasterId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSeriesMasterId(val)
        return nil
    }
    res["showAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.FreeBusyStatus)
        m.SetShowAs(&cast)
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty))
        }
        m.SetSingleValueExtendedProperties(res)
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        m.SetStart(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone))
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubject(val)
        return nil
    }
    res["transactionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTransactionId(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ParseEventType)
        if err != nil {
            return err
        }
        cast := val.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EventType)
        m.SetType_escaped(&cast)
        return nil
    }
    res["webLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebLink(val)
        return nil
    }
    return res
}
func (m *Delta) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Delta) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowNewTimeProposals", m.GetAllowNewTimeProposals())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bodyPreview", m.GetBodyPreview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideAttendees", m.GetHideAttendees())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("iCalUId", m.GetICalUId())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := m.GetImportance().String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAllDay", m.GetIsAllDay())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCancelled", m.GetIsCancelled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDraft", m.GetIsDraft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOnlineMeeting", m.GetIsOnlineMeeting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizer", m.GetIsOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReminderOn", m.GetIsReminderOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocations()))
        for i, v := range m.GetLocations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("locations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onlineMeeting", m.GetOnlineMeeting())
        if err != nil {
            return err
        }
    }
    if m.GetOnlineMeetingProvider() != nil {
        cast := m.GetOnlineMeetingProvider().String()
        err = writer.WriteStringValue("onlineMeetingProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineMeetingUrl", m.GetOnlineMeetingUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizer", m.GetOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalEndTimeZone", m.GetOriginalEndTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("originalStart", m.GetOriginalStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalStartTimeZone", m.GetOriginalStartTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("reminderMinutesBeforeStart", m.GetReminderMinutesBeforeStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("responseRequested", m.GetResponseRequested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("responseStatus", m.GetResponseStatus())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivity() != nil {
        cast := m.GetSensitivity().String()
        err = writer.WriteStringValue("sensitivity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("seriesMasterId", m.GetSeriesMasterId())
        if err != nil {
            return err
        }
    }
    if m.GetShowAs() != nil {
        cast := m.GetShowAs().String()
        err = writer.WriteStringValue("showAs", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("transactionId", m.GetTransactionId())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webLink", m.GetWebLink())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowNewTimeProposals property value. true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
// Parameters:
//  - value : Value to set for the allowNewTimeProposals property.
func (m *Delta) SetAllowNewTimeProposals(value *bool)() {
    m.allowNewTimeProposals = value
}
// Sets the attachments property value. The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the attachments property.
func (m *Delta) SetAttachments(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attachment)() {
    m.attachments = value
}
// Sets the attendees property value. The collection of attendees for the event.
// Parameters:
//  - value : Value to set for the attendees property.
func (m *Delta) SetAttendees(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Attendee)() {
    m.attendees = value
}
// Sets the body property value. The body of the message associated with the event. It can be in HTML or text format.
// Parameters:
//  - value : Value to set for the body property.
func (m *Delta) SetBody(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ItemBody)() {
    m.body = value
}
// Sets the bodyPreview property value. The preview of the message associated with the event. It is in text format.
// Parameters:
//  - value : Value to set for the bodyPreview property.
func (m *Delta) SetBodyPreview(value *string)() {
    m.bodyPreview = value
}
// Sets the calendar property value. The calendar that contains the event. Navigation property. Read-only.
// Parameters:
//  - value : Value to set for the calendar property.
func (m *Delta) SetCalendar(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar)() {
    m.calendar = value
}
// Sets the end property value. The date, time, and time zone that the event ends. By default, the end time is in UTC.
// Parameters:
//  - value : Value to set for the end property.
func (m *Delta) SetEnd(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone)() {
    m.end = value
}
// Sets the extensions property value. The collection of open extensions defined for the event. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *Delta) SetExtensions(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Extension)() {
    m.extensions = value
}
// Sets the hasAttachments property value. Set to true if the event has attachments.
// Parameters:
//  - value : Value to set for the hasAttachments property.
func (m *Delta) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// Sets the hideAttendees property value. When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
// Parameters:
//  - value : Value to set for the hideAttendees property.
func (m *Delta) SetHideAttendees(value *bool)() {
    m.hideAttendees = value
}
// Sets the iCalUId property value. A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. Read-only.
// Parameters:
//  - value : Value to set for the iCalUId property.
func (m *Delta) SetICalUId(value *string)() {
    m.iCalUId = value
}
// Sets the importance property value. 
// Parameters:
//  - value : Value to set for the importance property.
func (m *Delta) SetImportance(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Importance)() {
    m.importance = value
}
// Sets the instances property value. The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the instances property.
func (m *Delta) SetInstances(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Event)() {
    m.instances = value
}
// Sets the isAllDay property value. 
// Parameters:
//  - value : Value to set for the isAllDay property.
func (m *Delta) SetIsAllDay(value *bool)() {
    m.isAllDay = value
}
// Sets the isCancelled property value. 
// Parameters:
//  - value : Value to set for the isCancelled property.
func (m *Delta) SetIsCancelled(value *bool)() {
    m.isCancelled = value
}
// Sets the isDraft property value. 
// Parameters:
//  - value : Value to set for the isDraft property.
func (m *Delta) SetIsDraft(value *bool)() {
    m.isDraft = value
}
// Sets the isOnlineMeeting property value. 
// Parameters:
//  - value : Value to set for the isOnlineMeeting property.
func (m *Delta) SetIsOnlineMeeting(value *bool)() {
    m.isOnlineMeeting = value
}
// Sets the isOrganizer property value. 
// Parameters:
//  - value : Value to set for the isOrganizer property.
func (m *Delta) SetIsOrganizer(value *bool)() {
    m.isOrganizer = value
}
// Sets the isReminderOn property value. 
// Parameters:
//  - value : Value to set for the isReminderOn property.
func (m *Delta) SetIsReminderOn(value *bool)() {
    m.isReminderOn = value
}
// Sets the location property value. 
// Parameters:
//  - value : Value to set for the location property.
func (m *Delta) SetLocation(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location)() {
    m.location = value
}
// Sets the locations property value. 
// Parameters:
//  - value : Value to set for the locations property.
func (m *Delta) SetLocations(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Location)() {
    m.locations = value
}
// Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the event. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the multiValueExtendedProperties property.
func (m *Delta) SetMultiValueExtendedProperties(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
// Sets the onlineMeeting property value. 
// Parameters:
//  - value : Value to set for the onlineMeeting property.
func (m *Delta) SetOnlineMeeting(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingInfo)() {
    m.onlineMeeting = value
}
// Sets the onlineMeetingProvider property value. 
// Parameters:
//  - value : Value to set for the onlineMeetingProvider property.
func (m *Delta) SetOnlineMeetingProvider(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.OnlineMeetingProviderType)() {
    m.onlineMeetingProvider = value
}
// Sets the onlineMeetingUrl property value. 
// Parameters:
//  - value : Value to set for the onlineMeetingUrl property.
func (m *Delta) SetOnlineMeetingUrl(value *string)() {
    m.onlineMeetingUrl = value
}
// Sets the organizer property value. 
// Parameters:
//  - value : Value to set for the organizer property.
func (m *Delta) SetOrganizer(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Recipient)() {
    m.organizer = value
}
// Sets the originalEndTimeZone property value. 
// Parameters:
//  - value : Value to set for the originalEndTimeZone property.
func (m *Delta) SetOriginalEndTimeZone(value *string)() {
    m.originalEndTimeZone = value
}
// Sets the originalStart property value. 
// Parameters:
//  - value : Value to set for the originalStart property.
func (m *Delta) SetOriginalStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.originalStart = value
}
// Sets the originalStartTimeZone property value. 
// Parameters:
//  - value : Value to set for the originalStartTimeZone property.
func (m *Delta) SetOriginalStartTimeZone(value *string)() {
    m.originalStartTimeZone = value
}
// Sets the recurrence property value. 
// Parameters:
//  - value : Value to set for the recurrence property.
func (m *Delta) SetRecurrence(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PatternedRecurrence)() {
    m.recurrence = value
}
// Sets the reminderMinutesBeforeStart property value. 
// Parameters:
//  - value : Value to set for the reminderMinutesBeforeStart property.
func (m *Delta) SetReminderMinutesBeforeStart(value *int32)() {
    m.reminderMinutesBeforeStart = value
}
// Sets the responseRequested property value. 
// Parameters:
//  - value : Value to set for the responseRequested property.
func (m *Delta) SetResponseRequested(value *bool)() {
    m.responseRequested = value
}
// Sets the responseStatus property value. 
// Parameters:
//  - value : Value to set for the responseStatus property.
func (m *Delta) SetResponseStatus(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResponseStatus)() {
    m.responseStatus = value
}
// Sets the sensitivity property value. 
// Parameters:
//  - value : Value to set for the sensitivity property.
func (m *Delta) SetSensitivity(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Sensitivity)() {
    m.sensitivity = value
}
// Sets the seriesMasterId property value. 
// Parameters:
//  - value : Value to set for the seriesMasterId property.
func (m *Delta) SetSeriesMasterId(value *string)() {
    m.seriesMasterId = value
}
// Sets the showAs property value. 
// Parameters:
//  - value : Value to set for the showAs property.
func (m *Delta) SetShowAs(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.FreeBusyStatus)() {
    m.showAs = value
}
// Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the event. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the singleValueExtendedProperties property.
func (m *Delta) SetSingleValueExtendedProperties(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
// Sets the start property value. 
// Parameters:
//  - value : Value to set for the start property.
func (m *Delta) SetStart(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DateTimeTimeZone)() {
    m.start = value
}
// Sets the subject property value. 
// Parameters:
//  - value : Value to set for the subject property.
func (m *Delta) SetSubject(value *string)() {
    m.subject = value
}
// Sets the transactionId property value. 
// Parameters:
//  - value : Value to set for the transactionId property.
func (m *Delta) SetTransactionId(value *string)() {
    m.transactionId = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *Delta) SetType_escaped(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EventType)() {
    m.type_escaped = value
}
// Sets the webLink property value. 
// Parameters:
//  - value : Value to set for the webLink property.
func (m *Delta) SetWebLink(value *string)() {
    m.webLink = value
}
