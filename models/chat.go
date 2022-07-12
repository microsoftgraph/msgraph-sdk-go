package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Chat provides operations to manage the admin singleton.
type Chat struct {
    Entity
    // The chatType property
    chatType *ChatType
    // Date and time at which the chat was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of all the apps in the chat. Nullable.
    installedApps []TeamsAppInstallationable
    // Date and time at which the chat was renamed or list of members were last changed. Read-only.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of all the members in the chat. Nullable.
    members []ConversationMemberable
    // A collection of all the messages in the chat. Nullable.
    messages []ChatMessageable
    // Represents details about an online meeting. If the chat isn't associated with an online meeting, the property is empty. Read-only.
    onlineMeetingInfo TeamworkOnlineMeetingInfoable
    // A collection of all the tabs in the chat. Nullable.
    tabs []TeamsTabable
    // The identifier of the tenant in which the chat was created. Read-only.
    tenantId *string
    // (Optional) Subject or topic for the chat. Only available for group chats.
    topic *string
    // The URL for the chat in Microsoft Teams. The URL should be treated as an opaque blob, and not parsed. Read-only.
    webUrl *string
}
// NewChat instantiates a new chat and sets the default values.
func NewChat()(*Chat) {
    m := &Chat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateChatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChat(), nil
}
// GetChatType gets the chatType property value. The chatType property
func (m *Chat) GetChatType()(*ChatType) {
    if m == nil {
        return nil
    } else {
        return m.chatType
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time at which the chat was created. Read-only.
func (m *Chat) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Chat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["chatType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatType(val.(*ChatType))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["installedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamsAppInstallationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAppInstallationable, len(val))
            for i, v := range val {
                res[i] = v.(TeamsAppInstallationable)
            }
            m.SetInstalledApps(res)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConversationMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationMemberable, len(val))
            for i, v := range val {
                res[i] = v.(ConversationMemberable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["messages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageable)
            }
            m.SetMessages(res)
        }
        return nil
    }
    res["onlineMeetingInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkOnlineMeetingInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineMeetingInfo(val.(TeamworkOnlineMeetingInfoable))
        }
        return nil
    }
    res["tabs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamsTabFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsTabable, len(val))
            for i, v := range val {
                res[i] = v.(TeamsTabable)
            }
            m.SetTabs(res)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["topic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetInstalledApps gets the installedApps property value. A collection of all the apps in the chat. Nullable.
func (m *Chat) GetInstalledApps()([]TeamsAppInstallationable) {
    if m == nil {
        return nil
    } else {
        return m.installedApps
    }
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Date and time at which the chat was renamed or list of members were last changed. Read-only.
func (m *Chat) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetMembers gets the members property value. A collection of all the members in the chat. Nullable.
func (m *Chat) GetMembers()([]ConversationMemberable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetMessages gets the messages property value. A collection of all the messages in the chat. Nullable.
func (m *Chat) GetMessages()([]ChatMessageable) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// GetOnlineMeetingInfo gets the onlineMeetingInfo property value. Represents details about an online meeting. If the chat isn't associated with an online meeting, the property is empty. Read-only.
func (m *Chat) GetOnlineMeetingInfo()(TeamworkOnlineMeetingInfoable) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetingInfo
    }
}
// GetTabs gets the tabs property value. A collection of all the tabs in the chat. Nullable.
func (m *Chat) GetTabs()([]TeamsTabable) {
    if m == nil {
        return nil
    } else {
        return m.tabs
    }
}
// GetTenantId gets the tenantId property value. The identifier of the tenant in which the chat was created. Read-only.
func (m *Chat) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetTopic gets the topic property value. (Optional) Subject or topic for the chat. Only available for group chats.
func (m *Chat) GetTopic()(*string) {
    if m == nil {
        return nil
    } else {
        return m.topic
    }
}
// GetWebUrl gets the webUrl property value. The URL for the chat in Microsoft Teams. The URL should be treated as an opaque blob, and not parsed. Read-only.
func (m *Chat) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// Serialize serializes information the current object
func (m *Chat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChatType() != nil {
        cast := (*m.GetChatType()).String()
        err = writer.WriteStringValue("chatType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetInstalledApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstalledApps()))
        for i, v := range m.GetInstalledApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("installedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMessages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onlineMeetingInfo", m.GetOnlineMeetingInfo())
        if err != nil {
            return err
        }
    }
    if m.GetTabs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTabs()))
        for i, v := range m.GetTabs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tabs", cast)
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
        err = writer.WriteStringValue("topic", m.GetTopic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChatType sets the chatType property value. The chatType property
func (m *Chat) SetChatType(value *ChatType)() {
    if m != nil {
        m.chatType = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time at which the chat was created. Read-only.
func (m *Chat) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetInstalledApps sets the installedApps property value. A collection of all the apps in the chat. Nullable.
func (m *Chat) SetInstalledApps(value []TeamsAppInstallationable)() {
    if m != nil {
        m.installedApps = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Date and time at which the chat was renamed or list of members were last changed. Read-only.
func (m *Chat) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetMembers sets the members property value. A collection of all the members in the chat. Nullable.
func (m *Chat) SetMembers(value []ConversationMemberable)() {
    if m != nil {
        m.members = value
    }
}
// SetMessages sets the messages property value. A collection of all the messages in the chat. Nullable.
func (m *Chat) SetMessages(value []ChatMessageable)() {
    if m != nil {
        m.messages = value
    }
}
// SetOnlineMeetingInfo sets the onlineMeetingInfo property value. Represents details about an online meeting. If the chat isn't associated with an online meeting, the property is empty. Read-only.
func (m *Chat) SetOnlineMeetingInfo(value TeamworkOnlineMeetingInfoable)() {
    if m != nil {
        m.onlineMeetingInfo = value
    }
}
// SetTabs sets the tabs property value. A collection of all the tabs in the chat. Nullable.
func (m *Chat) SetTabs(value []TeamsTabable)() {
    if m != nil {
        m.tabs = value
    }
}
// SetTenantId sets the tenantId property value. The identifier of the tenant in which the chat was created. Read-only.
func (m *Chat) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetTopic sets the topic property value. (Optional) Subject or topic for the chat. Only available for group chats.
func (m *Chat) SetTopic(value *string)() {
    if m != nil {
        m.topic = value
    }
}
// SetWebUrl sets the webUrl property value. The URL for the chat in Microsoft Teams. The URL should be treated as an opaque blob, and not parsed. Read-only.
func (m *Chat) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
