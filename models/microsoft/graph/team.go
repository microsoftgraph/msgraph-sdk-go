package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Team 
type Team struct {
    Entity
    // The collection of channels & messages associated with the team.
    channels []Channel;
    // An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.
    classification *string;
    // Timestamp at which the team was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // An optional description for the team. Maximum length: 1024 characters.
    description *string;
    // The name of the team.
    displayName *string;
    // Settings to configure use of Giphy, memes, and stickers in the team.
    funSettings *TeamFunSettings;
    // 
    group *Group;
    // Settings to configure whether guests can create, update, or delete channels in the team.
    guestSettings *TeamGuestSettings;
    // The apps installed in this team.
    installedApps []TeamsAppInstallation;
    // A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.
    internalId *string;
    // Whether this team is in read-only mode.
    isArchived *bool;
    // Members and owners of the team.
    members []ConversationMember;
    // Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
    memberSettings *TeamMemberSettings;
    // Settings to configure messaging and mentions in the team.
    messagingSettings *TeamMessagingSettings;
    // The async operations that ran or are running on this team.
    operations []TeamsAsyncOperation;
    // The general channel for the team.
    primaryChannel *Channel;
    // The schedule of shifts for this team.
    schedule *Schedule;
    // Optional. Indicates whether the team is intended for a particular use case.  Each team specialization has access to unique behaviors and experiences targeted to its use case.
    specialization *TeamSpecialization;
    // The template this team was created from. See available templates.
    template *TeamsTemplate;
    // The visibility of the group and team. Defaults to Public.
    visibility *TeamVisibilityType;
    // A hyperlink that will go to the team in the Microsoft Teams client. This is the URL that you get when you right-click a team in the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed.
    webUrl *string;
}
// NewTeam instantiates a new team and sets the default values.
func NewTeam()(*Team) {
    m := &Team{
        Entity: *NewEntity(),
    }
    return m
}
// GetChannels gets the channels property value. The collection of channels & messages associated with the team.
func (m *Team) GetChannels()([]Channel) {
    if m == nil {
        return nil
    } else {
        return m.channels
    }
}
// GetClassification gets the classification property value. An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.
func (m *Team) GetClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp at which the team was created.
func (m *Team) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. An optional description for the team. Maximum length: 1024 characters.
func (m *Team) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name of the team.
func (m *Team) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFunSettings gets the funSettings property value. Settings to configure use of Giphy, memes, and stickers in the team.
func (m *Team) GetFunSettings()(*TeamFunSettings) {
    if m == nil {
        return nil
    } else {
        return m.funSettings
    }
}
// GetGroup gets the group property value. 
func (m *Team) GetGroup()(*Group) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// GetGuestSettings gets the guestSettings property value. Settings to configure whether guests can create, update, or delete channels in the team.
func (m *Team) GetGuestSettings()(*TeamGuestSettings) {
    if m == nil {
        return nil
    } else {
        return m.guestSettings
    }
}
// GetInstalledApps gets the installedApps property value. The apps installed in this team.
func (m *Team) GetInstalledApps()([]TeamsAppInstallation) {
    if m == nil {
        return nil
    } else {
        return m.installedApps
    }
}
// GetInternalId gets the internalId property value. A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.
func (m *Team) GetInternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalId
    }
}
// GetIsArchived gets the isArchived property value. Whether this team is in read-only mode.
func (m *Team) GetIsArchived()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isArchived
    }
}
// GetMembers gets the members property value. Members and owners of the team.
func (m *Team) GetMembers()([]ConversationMember) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetMemberSettings gets the memberSettings property value. Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
func (m *Team) GetMemberSettings()(*TeamMemberSettings) {
    if m == nil {
        return nil
    } else {
        return m.memberSettings
    }
}
// GetMessagingSettings gets the messagingSettings property value. Settings to configure messaging and mentions in the team.
func (m *Team) GetMessagingSettings()(*TeamMessagingSettings) {
    if m == nil {
        return nil
    } else {
        return m.messagingSettings
    }
}
// GetOperations gets the operations property value. The async operations that ran or are running on this team.
func (m *Team) GetOperations()([]TeamsAsyncOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPrimaryChannel gets the primaryChannel property value. The general channel for the team.
func (m *Team) GetPrimaryChannel()(*Channel) {
    if m == nil {
        return nil
    } else {
        return m.primaryChannel
    }
}
// GetSchedule gets the schedule property value. The schedule of shifts for this team.
func (m *Team) GetSchedule()(*Schedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetSpecialization gets the specialization property value. Optional. Indicates whether the team is intended for a particular use case.  Each team specialization has access to unique behaviors and experiences targeted to its use case.
func (m *Team) GetSpecialization()(*TeamSpecialization) {
    if m == nil {
        return nil
    } else {
        return m.specialization
    }
}
// GetTemplate gets the template property value. The template this team was created from. See available templates.
func (m *Team) GetTemplate()(*TeamsTemplate) {
    if m == nil {
        return nil
    } else {
        return m.template
    }
}
// GetVisibility gets the visibility property value. The visibility of the group and team. Defaults to Public.
func (m *Team) GetVisibility()(*TeamVisibilityType) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// GetWebUrl gets the webUrl property value. A hyperlink that will go to the team in the Microsoft Teams client. This is the URL that you get when you right-click a team in the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed.
func (m *Team) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Team) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["channels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChannel() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Channel, len(val))
            for i, v := range val {
                res[i] = *(v.(*Channel))
            }
            m.SetChannels(res)
        }
        return nil
    }
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["funSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamFunSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunSettings(val.(*TeamFunSettings))
        }
        return nil
    }
    res["group"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(*Group))
        }
        return nil
    }
    res["guestSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamGuestSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestSettings(val.(*TeamGuestSettings))
        }
        return nil
    }
    res["installedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppInstallation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAppInstallation, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamsAppInstallation))
            }
            m.SetInstalledApps(res)
        }
        return nil
    }
    res["internalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalId(val)
        }
        return nil
    }
    res["isArchived"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsArchived(val)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversationMember() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationMember, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConversationMember))
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["memberSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamMemberSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberSettings(val.(*TeamMemberSettings))
        }
        return nil
    }
    res["messagingSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamMessagingSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessagingSettings(val.(*TeamMessagingSettings))
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAsyncOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAsyncOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamsAsyncOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["primaryChannel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChannel() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryChannel(val.(*Channel))
        }
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(*Schedule))
        }
        return nil
    }
    res["specialization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamSpecialization)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TeamSpecialization)
            m.SetSpecialization(&cast)
        }
        return nil
    }
    res["template"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsTemplate() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplate(val.(*TeamsTemplate))
        }
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamVisibilityType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TeamVisibilityType)
            m.SetVisibility(&cast)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *Team) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Team) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChannels() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChannels()))
        for i, v := range m.GetChannels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("channels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("classification", m.GetClassification())
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
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("funSettings", m.GetFunSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("guestSettings", m.GetGuestSettings())
        if err != nil {
            return err
        }
    }
    if m.GetInstalledApps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInstalledApps()))
        for i, v := range m.GetInstalledApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("installedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internalId", m.GetInternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("memberSettings", m.GetMemberSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("messagingSettings", m.GetMessagingSettings())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
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
        err = writer.WriteObjectValue("primaryChannel", m.GetPrimaryChannel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    if m.GetSpecialization() != nil {
        cast := m.GetSpecialization().String()
        err = writer.WriteStringValue("specialization", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("template", m.GetTemplate())
        if err != nil {
            return err
        }
    }
    if m.GetVisibility() != nil {
        cast := m.GetVisibility().String()
        err = writer.WriteStringValue("visibility", &cast)
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
// SetChannels sets the channels property value. The collection of channels & messages associated with the team.
func (m *Team) SetChannels(value []Channel)() {
    if m != nil {
        m.channels = value
    }
}
// SetClassification sets the classification property value. An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.
func (m *Team) SetClassification(value *string)() {
    if m != nil {
        m.classification = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp at which the team was created.
func (m *Team) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. An optional description for the team. Maximum length: 1024 characters.
func (m *Team) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The name of the team.
func (m *Team) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFunSettings sets the funSettings property value. Settings to configure use of Giphy, memes, and stickers in the team.
func (m *Team) SetFunSettings(value *TeamFunSettings)() {
    if m != nil {
        m.funSettings = value
    }
}
// SetGroup sets the group property value. 
func (m *Team) SetGroup(value *Group)() {
    if m != nil {
        m.group = value
    }
}
// SetGuestSettings sets the guestSettings property value. Settings to configure whether guests can create, update, or delete channels in the team.
func (m *Team) SetGuestSettings(value *TeamGuestSettings)() {
    if m != nil {
        m.guestSettings = value
    }
}
// SetInstalledApps sets the installedApps property value. The apps installed in this team.
func (m *Team) SetInstalledApps(value []TeamsAppInstallation)() {
    if m != nil {
        m.installedApps = value
    }
}
// SetInternalId sets the internalId property value. A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.
func (m *Team) SetInternalId(value *string)() {
    if m != nil {
        m.internalId = value
    }
}
// SetIsArchived sets the isArchived property value. Whether this team is in read-only mode.
func (m *Team) SetIsArchived(value *bool)() {
    if m != nil {
        m.isArchived = value
    }
}
// SetMembers sets the members property value. Members and owners of the team.
func (m *Team) SetMembers(value []ConversationMember)() {
    if m != nil {
        m.members = value
    }
}
// SetMemberSettings sets the memberSettings property value. Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
func (m *Team) SetMemberSettings(value *TeamMemberSettings)() {
    if m != nil {
        m.memberSettings = value
    }
}
// SetMessagingSettings sets the messagingSettings property value. Settings to configure messaging and mentions in the team.
func (m *Team) SetMessagingSettings(value *TeamMessagingSettings)() {
    if m != nil {
        m.messagingSettings = value
    }
}
// SetOperations sets the operations property value. The async operations that ran or are running on this team.
func (m *Team) SetOperations(value []TeamsAsyncOperation)() {
    if m != nil {
        m.operations = value
    }
}
// SetPrimaryChannel sets the primaryChannel property value. The general channel for the team.
func (m *Team) SetPrimaryChannel(value *Channel)() {
    if m != nil {
        m.primaryChannel = value
    }
}
// SetSchedule sets the schedule property value. The schedule of shifts for this team.
func (m *Team) SetSchedule(value *Schedule)() {
    if m != nil {
        m.schedule = value
    }
}
// SetSpecialization sets the specialization property value. Optional. Indicates whether the team is intended for a particular use case.  Each team specialization has access to unique behaviors and experiences targeted to its use case.
func (m *Team) SetSpecialization(value *TeamSpecialization)() {
    if m != nil {
        m.specialization = value
    }
}
// SetTemplate sets the template property value. The template this team was created from. See available templates.
func (m *Team) SetTemplate(value *TeamsTemplate)() {
    if m != nil {
        m.template = value
    }
}
// SetVisibility sets the visibility property value. The visibility of the group and team. Defaults to Public.
func (m *Team) SetVisibility(value *TeamVisibilityType)() {
    if m != nil {
        m.visibility = value
    }
}
// SetWebUrl sets the webUrl property value. A hyperlink that will go to the team in the Microsoft Teams client. This is the URL that you get when you right-click a team in the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed.
func (m *Team) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
