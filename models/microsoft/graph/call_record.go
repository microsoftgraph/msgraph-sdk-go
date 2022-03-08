package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/callrecords"
)

// CallRecord provides operations to manage the cloudCommunications singleton.
type CallRecord struct {
    Entity
    // UTC time when the last user left the call. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Meeting URL associated to the call. May not be available for a peerToPeer call record type.
    joinWebUrl *string;
    // UTC time when the call record was created. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of all the modalities used in the call. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
    modalities []i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality;
    // The organizing party's identity.
    organizer IdentitySetable;
    // List of distinct identities involved in the call.
    participants []IdentitySetable;
    // List of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
    sessions []Sessionable;
    // UTC time when the first user joined the call. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates the type of the call. Possible values are: unknown, groupCall, peerToPeer, unknownFutureValue.
    type_escaped *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.CallType;
    // Monotonically increasing version of the call record. Higher version call records with the same id includes additional data compared to the lower version.
    version *int64;
}
// NewCallRecord instantiates a new callRecord and sets the default values.
func NewCallRecord()(*CallRecord) {
    m := &CallRecord{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCallRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallRecordFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCallRecord(), nil
}
// GetEndDateTime gets the endDateTime property value. UTC time when the last user left the call. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CallRecord) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallRecord) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["organizer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(IdentitySetable))
        }
        return nil
    }
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentitySetable, len(val))
            for i, v := range val {
                res[i] = v.(IdentitySetable)
            }
            m.SetParticipants(res)
        }
        return nil
    }
    res["sessions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSessionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Sessionable, len(val))
            for i, v := range val {
                res[i] = v.(Sessionable)
            }
            m.SetSessions(res)
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
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseCallType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.CallType))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetJoinWebUrl gets the joinWebUrl property value. Meeting URL associated to the call. May not be available for a peerToPeer call record type.
func (m *CallRecord) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. UTC time when the call record was created. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CallRecord) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetModalities gets the modalities property value. List of all the modalities used in the call. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
func (m *CallRecord) GetModalities()([]i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality) {
    if m == nil {
        return nil
    } else {
        return m.modalities
    }
}
// GetOrganizer gets the organizer property value. The organizing party's identity.
func (m *CallRecord) GetOrganizer()(IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.organizer
    }
}
// GetParticipants gets the participants property value. List of distinct identities involved in the call.
func (m *CallRecord) GetParticipants()([]IdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// GetSessions gets the sessions property value. List of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
func (m *CallRecord) GetSessions()([]Sessionable) {
    if m == nil {
        return nil
    } else {
        return m.sessions
    }
}
// GetStartDateTime gets the startDateTime property value. UTC time when the first user joined the call. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CallRecord) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetType gets the type property value. Indicates the type of the call. Possible values are: unknown, groupCall, peerToPeer, unknownFutureValue.
func (m *CallRecord) GetType()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.CallType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetVersion gets the version property value. Monotonically increasing version of the call record. Higher version call records with the same id includes additional data compared to the lower version.
func (m *CallRecord) GetVersion()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *CallRecord) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CallRecord) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
    {
        err = writer.WriteObjectValue("organizer", m.GetOrganizer())
        if err != nil {
            return err
        }
    }
    if m.GetParticipants() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("participants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSessions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndDateTime sets the endDateTime property value. UTC time when the last user left the call. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CallRecord) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetJoinWebUrl sets the joinWebUrl property value. Meeting URL associated to the call. May not be available for a peerToPeer call record type.
func (m *CallRecord) SetJoinWebUrl(value *string)() {
    if m != nil {
        m.joinWebUrl = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. UTC time when the call record was created. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *CallRecord) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetModalities sets the modalities property value. List of all the modalities used in the call. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
func (m *CallRecord) SetModalities(value []i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.Modality)() {
    if m != nil {
        m.modalities = value
    }
}
// SetOrganizer sets the organizer property value. The organizing party's identity.
func (m *CallRecord) SetOrganizer(value IdentitySetable)() {
    if m != nil {
        m.organizer = value
    }
}
// SetParticipants sets the participants property value. List of distinct identities involved in the call.
func (m *CallRecord) SetParticipants(value []IdentitySetable)() {
    if m != nil {
        m.participants = value
    }
}
// SetSessions sets the sessions property value. List of sessions involved in the call. Peer-to-peer calls typically only have one session, whereas group calls typically have at least one session per participant. Read-only. Nullable.
func (m *CallRecord) SetSessions(value []Sessionable)() {
    if m != nil {
        m.sessions = value
    }
}
// SetStartDateTime sets the startDateTime property value. UTC time when the first user joined the call. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CallRecord) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetType sets the type property value. Indicates the type of the call. Possible values are: unknown, groupCall, peerToPeer, unknownFutureValue.
func (m *CallRecord) SetType(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.CallType)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetVersion sets the version property value. Monotonically increasing version of the call record. Higher version call records with the same id includes additional data compared to the lower version.
func (m *CallRecord) SetVersion(value *int64)() {
    if m != nil {
        m.version = value
    }
}
