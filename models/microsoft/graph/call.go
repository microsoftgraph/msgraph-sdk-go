package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Call 
type Call struct {
    Entity
    // Read-only. Nullable.
    audioRoutingGroups []AudioRoutingGroup;
    // The callback URL on which callbacks will be delivered. Must be https.
    callbackUri *string;
    // A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in a P2P call.  This needs to be copied over from Microsoft.Graph.Call.CallChainId.
    callChainId *string;
    // 
    callOptions *CallOptions;
    // The routing information on how the call was retargeted. Read-only.
    callRoutes []CallRoute;
    // The chat information. Required information for joining a meeting.
    chatInfo *ChatInfo;
    // The direction of the call. The possible value are incoming or outgoing. Read-only.
    direction *CallDirection;
    // The context associated with an incoming call. Read-only. Server generated.
    incomingContext *IncomingContext;
    // The media configuration. Required.
    mediaConfig *MediaConfig;
    // Read-only. The call media state.
    mediaState *CallMediaState;
    // The meeting information that's required for joining a meeting.
    meetingInfo *MeetingInfo;
    // 
    myParticipantId *string;
    // Read-only. Nullable.
    operations []CommsOperation;
    // Read-only. Nullable.
    participants []Participant;
    // 
    requestedModalities []Modality;
    // 
    resultInfo *ResultInfo;
    // 
    source *ParticipantInfo;
    // 
    state *CallState;
    // 
    subject *string;
    // 
    targets []InvitationParticipantInfo;
    // 
    tenantId *string;
    // 
    toneInfo *ToneInfo;
    // The transcription information for the call. Read-only.
    transcription *CallTranscriptionInfo;
}
// NewCall instantiates a new call and sets the default values.
func NewCall()(*Call) {
    m := &Call{
        Entity: *NewEntity(),
    }
    return m
}
// GetAudioRoutingGroups gets the audioRoutingGroups property value. Read-only. Nullable.
func (m *Call) GetAudioRoutingGroups()([]AudioRoutingGroup) {
    if m == nil {
        return nil
    } else {
        return m.audioRoutingGroups
    }
}
// GetCallbackUri gets the callbackUri property value. The callback URL on which callbacks will be delivered. Must be https.
func (m *Call) GetCallbackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callbackUri
    }
}
// GetCallChainId gets the callChainId property value. A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in a P2P call.  This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *Call) GetCallChainId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callChainId
    }
}
// GetCallOptions gets the callOptions property value. 
func (m *Call) GetCallOptions()(*CallOptions) {
    if m == nil {
        return nil
    } else {
        return m.callOptions
    }
}
// GetCallRoutes gets the callRoutes property value. The routing information on how the call was retargeted. Read-only.
func (m *Call) GetCallRoutes()([]CallRoute) {
    if m == nil {
        return nil
    } else {
        return m.callRoutes
    }
}
// GetChatInfo gets the chatInfo property value. The chat information. Required information for joining a meeting.
func (m *Call) GetChatInfo()(*ChatInfo) {
    if m == nil {
        return nil
    } else {
        return m.chatInfo
    }
}
// GetDirection gets the direction property value. The direction of the call. The possible value are incoming or outgoing. Read-only.
func (m *Call) GetDirection()(*CallDirection) {
    if m == nil {
        return nil
    } else {
        return m.direction
    }
}
// GetIncomingContext gets the incomingContext property value. The context associated with an incoming call. Read-only. Server generated.
func (m *Call) GetIncomingContext()(*IncomingContext) {
    if m == nil {
        return nil
    } else {
        return m.incomingContext
    }
}
// GetMediaConfig gets the mediaConfig property value. The media configuration. Required.
func (m *Call) GetMediaConfig()(*MediaConfig) {
    if m == nil {
        return nil
    } else {
        return m.mediaConfig
    }
}
// GetMediaState gets the mediaState property value. Read-only. The call media state.
func (m *Call) GetMediaState()(*CallMediaState) {
    if m == nil {
        return nil
    } else {
        return m.mediaState
    }
}
// GetMeetingInfo gets the meetingInfo property value. The meeting information that's required for joining a meeting.
func (m *Call) GetMeetingInfo()(*MeetingInfo) {
    if m == nil {
        return nil
    } else {
        return m.meetingInfo
    }
}
// GetMyParticipantId gets the myParticipantId property value. 
func (m *Call) GetMyParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.myParticipantId
    }
}
// GetOperations gets the operations property value. Read-only. Nullable.
func (m *Call) GetOperations()([]CommsOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetParticipants gets the participants property value. Read-only. Nullable.
func (m *Call) GetParticipants()([]Participant) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// GetRequestedModalities gets the requestedModalities property value. 
func (m *Call) GetRequestedModalities()([]Modality) {
    if m == nil {
        return nil
    } else {
        return m.requestedModalities
    }
}
// GetResultInfo gets the resultInfo property value. 
func (m *Call) GetResultInfo()(*ResultInfo) {
    if m == nil {
        return nil
    } else {
        return m.resultInfo
    }
}
// GetSource gets the source property value. 
func (m *Call) GetSource()(*ParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetState gets the state property value. 
func (m *Call) GetState()(*CallState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetSubject gets the subject property value. 
func (m *Call) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetTargets gets the targets property value. 
func (m *Call) GetTargets()([]InvitationParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.targets
    }
}
// GetTenantId gets the tenantId property value. 
func (m *Call) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetToneInfo gets the toneInfo property value. 
func (m *Call) GetToneInfo()(*ToneInfo) {
    if m == nil {
        return nil
    } else {
        return m.toneInfo
    }
}
// GetTranscription gets the transcription property value. The transcription information for the call. Read-only.
func (m *Call) GetTranscription()(*CallTranscriptionInfo) {
    if m == nil {
        return nil
    } else {
        return m.transcription
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Call) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audioRoutingGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAudioRoutingGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AudioRoutingGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*AudioRoutingGroup))
            }
            m.SetAudioRoutingGroups(res)
        }
        return nil
    }
    res["callbackUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["callChainId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallChainId(val)
        }
        return nil
    }
    res["callOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallOptions() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallOptions(val.(*CallOptions))
        }
        return nil
    }
    res["callRoutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallRoute() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CallRoute, len(val))
            for i, v := range val {
                res[i] = *(v.(*CallRoute))
            }
            m.SetCallRoutes(res)
        }
        return nil
    }
    res["chatInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(*ChatInfo))
        }
        return nil
    }
    res["direction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallDirection)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CallDirection)
            m.SetDirection(&cast)
        }
        return nil
    }
    res["incomingContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncomingContext() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncomingContext(val.(*IncomingContext))
        }
        return nil
    }
    res["mediaConfig"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMediaConfig() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaConfig(val.(*MediaConfig))
        }
        return nil
    }
    res["mediaState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallMediaState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaState(val.(*CallMediaState))
        }
        return nil
    }
    res["meetingInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingInfo(val.(*MeetingInfo))
        }
        return nil
    }
    res["myParticipantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMyParticipantId(val)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCommsOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CommsOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*CommsOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParticipant() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Participant, len(val))
            for i, v := range val {
                res[i] = *(v.(*Participant))
            }
            m.SetParticipants(res)
        }
        return nil
    }
    res["requestedModalities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Modality, len(val))
            for i, v := range val {
                res[i] = *(v.(*Modality))
            }
            m.SetRequestedModalities(res)
        }
        return nil
    }
    res["resultInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResultInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultInfo(val.(*ResultInfo))
        }
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParticipantInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*ParticipantInfo))
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CallState)
            m.SetState(&cast)
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
    res["targets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInvitationParticipantInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InvitationParticipantInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*InvitationParticipantInfo))
            }
            m.SetTargets(res)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["toneInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewToneInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToneInfo(val.(*ToneInfo))
        }
        return nil
    }
    res["transcription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallTranscriptionInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranscription(val.(*CallTranscriptionInfo))
        }
        return nil
    }
    return res
}
func (m *Call) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Call) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAudioRoutingGroups()))
        for i, v := range m.GetAudioRoutingGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("audioRoutingGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callChainId", m.GetCallChainId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("callOptions", m.GetCallOptions())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCallRoutes()))
        for i, v := range m.GetCallRoutes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("callRoutes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("chatInfo", m.GetChatInfo())
        if err != nil {
            return err
        }
    }
    if m.GetDirection() != nil {
        cast := m.GetDirection().String()
        err = writer.WriteStringValue("direction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("incomingContext", m.GetIncomingContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaConfig", m.GetMediaConfig())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaState", m.GetMediaState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("meetingInfo", m.GetMeetingInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("myParticipantId", m.GetMyParticipantId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("participants", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("requestedModalities", SerializeModality(m.GetRequestedModalities()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resultInfo", m.GetResultInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("toneInfo", m.GetToneInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("transcription", m.GetTranscription())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudioRoutingGroups sets the audioRoutingGroups property value. Read-only. Nullable.
func (m *Call) SetAudioRoutingGroups(value []AudioRoutingGroup)() {
    if m != nil {
        m.audioRoutingGroups = value
    }
}
// SetCallbackUri sets the callbackUri property value. The callback URL on which callbacks will be delivered. Must be https.
func (m *Call) SetCallbackUri(value *string)() {
    if m != nil {
        m.callbackUri = value
    }
}
// SetCallChainId sets the callChainId property value. A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in a P2P call.  This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *Call) SetCallChainId(value *string)() {
    if m != nil {
        m.callChainId = value
    }
}
// SetCallOptions sets the callOptions property value. 
func (m *Call) SetCallOptions(value *CallOptions)() {
    if m != nil {
        m.callOptions = value
    }
}
// SetCallRoutes sets the callRoutes property value. The routing information on how the call was retargeted. Read-only.
func (m *Call) SetCallRoutes(value []CallRoute)() {
    if m != nil {
        m.callRoutes = value
    }
}
// SetChatInfo sets the chatInfo property value. The chat information. Required information for joining a meeting.
func (m *Call) SetChatInfo(value *ChatInfo)() {
    if m != nil {
        m.chatInfo = value
    }
}
// SetDirection sets the direction property value. The direction of the call. The possible value are incoming or outgoing. Read-only.
func (m *Call) SetDirection(value *CallDirection)() {
    if m != nil {
        m.direction = value
    }
}
// SetIncomingContext sets the incomingContext property value. The context associated with an incoming call. Read-only. Server generated.
func (m *Call) SetIncomingContext(value *IncomingContext)() {
    if m != nil {
        m.incomingContext = value
    }
}
// SetMediaConfig sets the mediaConfig property value. The media configuration. Required.
func (m *Call) SetMediaConfig(value *MediaConfig)() {
    if m != nil {
        m.mediaConfig = value
    }
}
// SetMediaState sets the mediaState property value. Read-only. The call media state.
func (m *Call) SetMediaState(value *CallMediaState)() {
    if m != nil {
        m.mediaState = value
    }
}
// SetMeetingInfo sets the meetingInfo property value. The meeting information that's required for joining a meeting.
func (m *Call) SetMeetingInfo(value *MeetingInfo)() {
    if m != nil {
        m.meetingInfo = value
    }
}
// SetMyParticipantId sets the myParticipantId property value. 
func (m *Call) SetMyParticipantId(value *string)() {
    if m != nil {
        m.myParticipantId = value
    }
}
// SetOperations sets the operations property value. Read-only. Nullable.
func (m *Call) SetOperations(value []CommsOperation)() {
    if m != nil {
        m.operations = value
    }
}
// SetParticipants sets the participants property value. Read-only. Nullable.
func (m *Call) SetParticipants(value []Participant)() {
    if m != nil {
        m.participants = value
    }
}
// SetRequestedModalities sets the requestedModalities property value. 
func (m *Call) SetRequestedModalities(value []Modality)() {
    if m != nil {
        m.requestedModalities = value
    }
}
// SetResultInfo sets the resultInfo property value. 
func (m *Call) SetResultInfo(value *ResultInfo)() {
    if m != nil {
        m.resultInfo = value
    }
}
// SetSource sets the source property value. 
func (m *Call) SetSource(value *ParticipantInfo)() {
    if m != nil {
        m.source = value
    }
}
// SetState sets the state property value. 
func (m *Call) SetState(value *CallState)() {
    if m != nil {
        m.state = value
    }
}
// SetSubject sets the subject property value. 
func (m *Call) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
// SetTargets sets the targets property value. 
func (m *Call) SetTargets(value []InvitationParticipantInfo)() {
    if m != nil {
        m.targets = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *Call) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetToneInfo sets the toneInfo property value. 
func (m *Call) SetToneInfo(value *ToneInfo)() {
    if m != nil {
        m.toneInfo = value
    }
}
// SetTranscription sets the transcription property value. The transcription information for the call. Read-only.
func (m *Call) SetTranscription(value *CallTranscriptionInfo)() {
    if m != nil {
        m.transcription = value
    }
}
