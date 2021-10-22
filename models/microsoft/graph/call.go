package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Call struct {
    Entity
    callbackUri *string;
    callChainId *string;
    callOptions *CallOptions;
    callRoutes []CallRoute;
    chatInfo *ChatInfo;
    direction *CallDirection;
    incomingContext *IncomingContext;
    mediaConfig *MediaConfig;
    mediaState *CallMediaState;
    meetingInfo *MeetingInfo;
    myParticipantId *string;
    operations []CommsOperation;
    participants []Participant;
    requestedModalities []Modality;
    resultInfo *ResultInfo;
    source *ParticipantInfo;
    state *CallState;
    subject *string;
    targets []InvitationParticipantInfo;
    tenantId *string;
    toneInfo *ToneInfo;
    transcription *CallTranscriptionInfo;
}
func NewCall()(*Call) {
    m := &Call{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Call) GetCallbackUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callbackUri
    }
}
func (m *Call) GetCallChainId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.callChainId
    }
}
func (m *Call) GetCallOptions()(*CallOptions) {
    if m == nil {
        return nil
    } else {
        return m.callOptions
    }
}
func (m *Call) GetCallRoutes()([]CallRoute) {
    if m == nil {
        return nil
    } else {
        return m.callRoutes
    }
}
func (m *Call) GetChatInfo()(*ChatInfo) {
    if m == nil {
        return nil
    } else {
        return m.chatInfo
    }
}
func (m *Call) GetDirection()(*CallDirection) {
    if m == nil {
        return nil
    } else {
        return m.direction
    }
}
func (m *Call) GetIncomingContext()(*IncomingContext) {
    if m == nil {
        return nil
    } else {
        return m.incomingContext
    }
}
func (m *Call) GetMediaConfig()(*MediaConfig) {
    if m == nil {
        return nil
    } else {
        return m.mediaConfig
    }
}
func (m *Call) GetMediaState()(*CallMediaState) {
    if m == nil {
        return nil
    } else {
        return m.mediaState
    }
}
func (m *Call) GetMeetingInfo()(*MeetingInfo) {
    if m == nil {
        return nil
    } else {
        return m.meetingInfo
    }
}
func (m *Call) GetMyParticipantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.myParticipantId
    }
}
func (m *Call) GetOperations()([]CommsOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
func (m *Call) GetParticipants()([]Participant) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
func (m *Call) GetRequestedModalities()([]Modality) {
    if m == nil {
        return nil
    } else {
        return m.requestedModalities
    }
}
func (m *Call) GetResultInfo()(*ResultInfo) {
    if m == nil {
        return nil
    } else {
        return m.resultInfo
    }
}
func (m *Call) GetSource()(*ParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
func (m *Call) GetState()(*CallState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *Call) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
func (m *Call) GetTargets()([]InvitationParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.targets
    }
}
func (m *Call) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *Call) GetToneInfo()(*ToneInfo) {
    if m == nil {
        return nil
    } else {
        return m.toneInfo
    }
}
func (m *Call) GetTranscription()(*CallTranscriptionInfo) {
    if m == nil {
        return nil
    } else {
        return m.transcription
    }
}
func (m *Call) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["callbackUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCallbackUri(val)
        return nil
    }
    res["callChainId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCallChainId(val)
        return nil
    }
    res["callOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallOptions() })
        if err != nil {
            return err
        }
        m.SetCallOptions(val.(*CallOptions))
        return nil
    }
    res["callRoutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallRoute() })
        if err != nil {
            return err
        }
        res := make([]CallRoute, len(val))
        for i, v := range val {
            res[i] = *(v.(*CallRoute))
        }
        m.SetCallRoutes(res)
        return nil
    }
    res["chatInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatInfo() })
        if err != nil {
            return err
        }
        m.SetChatInfo(val.(*ChatInfo))
        return nil
    }
    res["direction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallDirection)
        if err != nil {
            return err
        }
        cast := val.(CallDirection)
        m.SetDirection(&cast)
        return nil
    }
    res["incomingContext"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIncomingContext() })
        if err != nil {
            return err
        }
        m.SetIncomingContext(val.(*IncomingContext))
        return nil
    }
    res["mediaConfig"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMediaConfig() })
        if err != nil {
            return err
        }
        m.SetMediaConfig(val.(*MediaConfig))
        return nil
    }
    res["mediaState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallMediaState() })
        if err != nil {
            return err
        }
        m.SetMediaState(val.(*CallMediaState))
        return nil
    }
    res["meetingInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingInfo() })
        if err != nil {
            return err
        }
        m.SetMeetingInfo(val.(*MeetingInfo))
        return nil
    }
    res["myParticipantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMyParticipantId(val)
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCommsOperation() })
        if err != nil {
            return err
        }
        res := make([]CommsOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*CommsOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParticipant() })
        if err != nil {
            return err
        }
        res := make([]Participant, len(val))
        for i, v := range val {
            res[i] = *(v.(*Participant))
        }
        m.SetParticipants(res)
        return nil
    }
    res["requestedModalities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseModality)
        if err != nil {
            return err
        }
        res := make([]Modality, len(val))
        for i, v := range val {
            res[i] = *(v.(*Modality))
        }
        m.SetRequestedModalities(res)
        return nil
    }
    res["resultInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResultInfo() })
        if err != nil {
            return err
        }
        m.SetResultInfo(val.(*ResultInfo))
        return nil
    }
    res["source"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParticipantInfo() })
        if err != nil {
            return err
        }
        m.SetSource(val.(*ParticipantInfo))
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallState)
        if err != nil {
            return err
        }
        cast := val.(CallState)
        m.SetState(&cast)
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
    res["targets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInvitationParticipantInfo() })
        if err != nil {
            return err
        }
        res := make([]InvitationParticipantInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*InvitationParticipantInfo))
        }
        m.SetTargets(res)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["toneInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewToneInfo() })
        if err != nil {
            return err
        }
        m.SetToneInfo(val.(*ToneInfo))
        return nil
    }
    res["transcription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCallTranscriptionInfo() })
        if err != nil {
            return err
        }
        m.SetTranscription(val.(*CallTranscriptionInfo))
        return nil
    }
    return res
}
func (m *Call) IsNil()(bool) {
    return m == nil
}
func (m *Call) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
func (m *Call) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
func (m *Call) SetCallChainId(value *string)() {
    m.callChainId = value
}
func (m *Call) SetCallOptions(value *CallOptions)() {
    m.callOptions = value
}
func (m *Call) SetCallRoutes(value []CallRoute)() {
    m.callRoutes = value
}
func (m *Call) SetChatInfo(value *ChatInfo)() {
    m.chatInfo = value
}
func (m *Call) SetDirection(value *CallDirection)() {
    m.direction = value
}
func (m *Call) SetIncomingContext(value *IncomingContext)() {
    m.incomingContext = value
}
func (m *Call) SetMediaConfig(value *MediaConfig)() {
    m.mediaConfig = value
}
func (m *Call) SetMediaState(value *CallMediaState)() {
    m.mediaState = value
}
func (m *Call) SetMeetingInfo(value *MeetingInfo)() {
    m.meetingInfo = value
}
func (m *Call) SetMyParticipantId(value *string)() {
    m.myParticipantId = value
}
func (m *Call) SetOperations(value []CommsOperation)() {
    m.operations = value
}
func (m *Call) SetParticipants(value []Participant)() {
    m.participants = value
}
func (m *Call) SetRequestedModalities(value []Modality)() {
    m.requestedModalities = value
}
func (m *Call) SetResultInfo(value *ResultInfo)() {
    m.resultInfo = value
}
func (m *Call) SetSource(value *ParticipantInfo)() {
    m.source = value
}
func (m *Call) SetState(value *CallState)() {
    m.state = value
}
func (m *Call) SetSubject(value *string)() {
    m.subject = value
}
func (m *Call) SetTargets(value []InvitationParticipantInfo)() {
    m.targets = value
}
func (m *Call) SetTenantId(value *string)() {
    m.tenantId = value
}
func (m *Call) SetToneInfo(value *ToneInfo)() {
    m.toneInfo = value
}
func (m *Call) SetTranscription(value *CallTranscriptionInfo)() {
    m.transcription = value
}
