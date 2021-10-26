package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Segment struct {
    Entity
    // Endpoint that answered this segment.
    callee *Endpoint;
    // Endpoint that initiated this segment.
    caller *Endpoint;
    // UTC time when the segment ended. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Failure information associated with the segment if it failed.
    failureInfo *FailureInfo;
    // Media associated with this segment.
    media []Media;
    // UTC time when the segment started. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new segment and sets the default values.
func NewSegment()(*Segment) {
    m := &Segment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the callee property value. Endpoint that answered this segment.
func (m *Segment) GetCallee()(*Endpoint) {
    if m == nil {
        return nil
    } else {
        return m.callee
    }
}
// Gets the caller property value. Endpoint that initiated this segment.
func (m *Segment) GetCaller()(*Endpoint) {
    if m == nil {
        return nil
    } else {
        return m.caller
    }
}
// Gets the endDateTime property value. UTC time when the segment ended. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Segment) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the failureInfo property value. Failure information associated with the segment if it failed.
func (m *Segment) GetFailureInfo()(*FailureInfo) {
    if m == nil {
        return nil
    } else {
        return m.failureInfo
    }
}
// Gets the media property value. Media associated with this segment.
func (m *Segment) GetMedia()([]Media) {
    if m == nil {
        return nil
    } else {
        return m.media
    }
}
// Gets the startDateTime property value. UTC time when the segment started. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Segment) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// The deserialization information for the current model
func (m *Segment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["callee"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEndpoint() })
        if err != nil {
            return err
        }
        m.SetCallee(val.(*Endpoint))
        return nil
    }
    res["caller"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEndpoint() })
        if err != nil {
            return err
        }
        m.SetCaller(val.(*Endpoint))
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["failureInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFailureInfo() })
        if err != nil {
            return err
        }
        m.SetFailureInfo(val.(*FailureInfo))
        return nil
    }
    res["media"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMedia() })
        if err != nil {
            return err
        }
        res := make([]Media, len(val))
        for i, v := range val {
            res[i] = *(v.(*Media))
        }
        m.SetMedia(res)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    return res
}
func (m *Segment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Segment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("callee", m.GetCallee())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("caller", m.GetCaller())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("failureInfo", m.GetFailureInfo())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMedia()))
        for i, v := range m.GetMedia() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("media", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the callee property value. Endpoint that answered this segment.
// Parameters:
//  - value : Value to set for the callee property.
func (m *Segment) SetCallee(value *Endpoint)() {
    m.callee = value
}
// Sets the caller property value. Endpoint that initiated this segment.
// Parameters:
//  - value : Value to set for the caller property.
func (m *Segment) SetCaller(value *Endpoint)() {
    m.caller = value
}
// Sets the endDateTime property value. UTC time when the segment ended. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *Segment) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the failureInfo property value. Failure information associated with the segment if it failed.
// Parameters:
//  - value : Value to set for the failureInfo property.
func (m *Segment) SetFailureInfo(value *FailureInfo)() {
    m.failureInfo = value
}
// Sets the media property value. Media associated with this segment.
// Parameters:
//  - value : Value to set for the media property.
func (m *Segment) SetMedia(value []Media)() {
    m.media = value
}
// Sets the startDateTime property value. UTC time when the segment started. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *Segment) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
