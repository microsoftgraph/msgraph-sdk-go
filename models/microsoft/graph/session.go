package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/callrecords"
)

// Session 
type Session struct {
    Entity
    // Endpoint that answered the session.
    callee *Endpoint;
    // Endpoint that initiated the session.
    caller *Endpoint;
    // UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Failure information associated with the session if the session failed.
    failureInfo *FailureInfo;
    // List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
    modalities []i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality;
    // The list of segments involved in the session. Read-only. Nullable.
    segments []Segment;
    // UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewSession instantiates a new session and sets the default values.
func NewSession()(*Session) {
    m := &Session{
        Entity: *NewEntity(),
    }
    return m
}
// GetCallee gets the callee property value. Endpoint that answered the session.
func (m *Session) GetCallee()(*Endpoint) {
    if m == nil {
        return nil
    } else {
        return m.callee
    }
}
// GetCaller gets the caller property value. Endpoint that initiated the session.
func (m *Session) GetCaller()(*Endpoint) {
    if m == nil {
        return nil
    } else {
        return m.caller
    }
}
// GetEndDateTime gets the endDateTime property value. UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Session) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetFailureInfo gets the failureInfo property value. Failure information associated with the session if the session failed.
func (m *Session) GetFailureInfo()(*FailureInfo) {
    if m == nil {
        return nil
    } else {
        return m.failureInfo
    }
}
// GetModalities gets the modalities property value. List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
func (m *Session) GetModalities()([]i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality) {
    if m == nil {
        return nil
    } else {
        return m.modalities
    }
}
// GetSegments gets the segments property value. The list of segments involved in the session. Read-only. Nullable.
func (m *Session) GetSegments()([]Segment) {
    if m == nil {
        return nil
    } else {
        return m.segments
    }
}
// GetStartDateTime gets the startDateTime property value. UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Session) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Session) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["callee"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEndpoint() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallee(val.(*Endpoint))
        }
        return nil
    }
    res["caller"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEndpoint() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaller(val.(*Endpoint))
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["failureInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFailureInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureInfo(val.(*FailureInfo))
        }
        return nil
    }
    res["modalities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality, len(val))
            for i, v := range val {
                res[i] = *(v.(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality))
            }
            m.SetModalities(res)
        }
        return nil
    }
    res["segments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSegment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Segment, len(val))
            for i, v := range val {
                res[i] = *(v.(*Segment))
            }
            m.SetSegments(res)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
func (m *Session) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Session) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetModalities() != nil {
        err = writer.WriteCollectionOfStringValues("modalities", i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.SerializeModality(m.GetModalities()))
        if err != nil {
            return err
        }
    }
    if m.GetSegments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSegments()))
        for i, v := range m.GetSegments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("segments", cast)
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
// SetCallee sets the callee property value. Endpoint that answered the session.
func (m *Session) SetCallee(value *Endpoint)() {
    if m != nil {
        m.callee = value
    }
}
// SetCaller sets the caller property value. Endpoint that initiated the session.
func (m *Session) SetCaller(value *Endpoint)() {
    if m != nil {
        m.caller = value
    }
}
// SetEndDateTime sets the endDateTime property value. UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Session) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetFailureInfo sets the failureInfo property value. Failure information associated with the session if the session failed.
func (m *Session) SetFailureInfo(value *FailureInfo)() {
    if m != nil {
        m.failureInfo = value
    }
}
// SetModalities sets the modalities property value. List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
func (m *Session) SetModalities(value []i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality)() {
    if m != nil {
        m.modalities = value
    }
}
// SetSegments sets the segments property value. The list of segments involved in the session. Read-only. Nullable.
func (m *Session) SetSegments(value []Segment)() {
    if m != nil {
        m.segments = value
    }
}
// SetStartDateTime sets the startDateTime property value. UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Session) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
