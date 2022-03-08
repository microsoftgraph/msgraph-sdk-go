package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ScheduleItem provides operations to call the getSchedule method.
type ScheduleItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date, time, and time zone that the corresponding event ends.
    end DateTimeTimeZoneable;
    // The sensitivity of the corresponding event. True if the event is marked private, false otherwise. Optional.
    isPrivate *bool;
    // The location where the corresponding event is held or attended from. Optional.
    location *string;
    // The date, time, and time zone that the corresponding event starts.
    start DateTimeTimeZoneable;
    // The availability status of the user or resource during the corresponding event. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
    status *FreeBusyStatus;
    // The corresponding event's subject line. Optional.
    subject *string;
}
// NewScheduleItem instantiates a new scheduleItem and sets the default values.
func NewScheduleItem()(*ScheduleItem) {
    m := &ScheduleItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateScheduleItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScheduleItemFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewScheduleItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScheduleItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnd gets the end property value. The date, time, and time zone that the corresponding event ends.
func (m *ScheduleItem) GetEnd()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScheduleItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["isPrivate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivate(val)
        }
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
        }
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*FreeBusyStatus))
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetIsPrivate gets the isPrivate property value. The sensitivity of the corresponding event. True if the event is marked private, false otherwise. Optional.
func (m *ScheduleItem) GetIsPrivate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPrivate
    }
}
// GetLocation gets the location property value. The location where the corresponding event is held or attended from. Optional.
func (m *ScheduleItem) GetLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetStart gets the start property value. The date, time, and time zone that the corresponding event starts.
func (m *ScheduleItem) GetStart()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
// GetStatus gets the status property value. The availability status of the user or resource during the corresponding event. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *ScheduleItem) GetStatus()(*FreeBusyStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetSubject gets the subject property value. The corresponding event's subject line. Optional.
func (m *ScheduleItem) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
func (m *ScheduleItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ScheduleItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPrivate", m.GetIsPrivate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *ScheduleItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnd sets the end property value. The date, time, and time zone that the corresponding event ends.
func (m *ScheduleItem) SetEnd(value DateTimeTimeZoneable)() {
    if m != nil {
        m.end = value
    }
}
// SetIsPrivate sets the isPrivate property value. The sensitivity of the corresponding event. True if the event is marked private, false otherwise. Optional.
func (m *ScheduleItem) SetIsPrivate(value *bool)() {
    if m != nil {
        m.isPrivate = value
    }
}
// SetLocation sets the location property value. The location where the corresponding event is held or attended from. Optional.
func (m *ScheduleItem) SetLocation(value *string)() {
    if m != nil {
        m.location = value
    }
}
// SetStart sets the start property value. The date, time, and time zone that the corresponding event starts.
func (m *ScheduleItem) SetStart(value DateTimeTimeZoneable)() {
    if m != nil {
        m.start = value
    }
}
// SetStatus sets the status property value. The availability status of the user or resource during the corresponding event. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *ScheduleItem) SetStatus(value *FreeBusyStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetSubject sets the subject property value. The corresponding event's subject line. Optional.
func (m *ScheduleItem) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
