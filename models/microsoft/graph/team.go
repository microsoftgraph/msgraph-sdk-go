package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new team and sets the default values.
func NewTeam()(*Team) {
    m := &Team{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the channels property value. The collection of channels & messages associated with the team.
func (m *Team) GetChannels()([]Channel) {
    if m == nil {
        return nil
    } else {
        return m.channels
    }
}
// Gets the classification property value. An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.
func (m *Team) GetClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// Gets the createdDateTime property value. Timestamp at which the team was created.
func (m *Team) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. An optional description for the team. Maximum length: 1024 characters.
func (m *Team) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of the team.
func (m *Team) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the funSettings property value. Settings to configure use of Giphy, memes, and stickers in the team.
func (m *Team) GetFunSettings()(*TeamFunSettings) {
    if m == nil {
        return nil
    } else {
        return m.funSettings
    }
}
// Gets the group property value. 
func (m *Team) GetGroup()(*Group) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// Gets the guestSettings property value. Settings to configure whether guests can create, update, or delete channels in the team.
func (m *Team) GetGuestSettings()(*TeamGuestSettings) {
    if m == nil {
        return nil
    } else {
        return m.guestSettings
    }
}
// Gets the installedApps property value. The apps installed in this team.
func (m *Team) GetInstalledApps()([]TeamsAppInstallation) {
    if m == nil {
        return nil
    } else {
        return m.installedApps
    }
}
// Gets the internalId property value. A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.
func (m *Team) GetInternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalId
    }
}
// Gets the isArchived property value. Whether this team is in read-only mode.
func (m *Team) GetIsArchived()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isArchived
    }
}
// Gets the members property value. Members and owners of the team.
func (m *Team) GetMembers()([]ConversationMember) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the memberSettings property value. Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
func (m *Team) GetMemberSettings()(*TeamMemberSettings) {
    if m == nil {
        return nil
    } else {
        return m.memberSettings
    }
}
// Gets the messagingSettings property value. Settings to configure messaging and mentions in the team.
func (m *Team) GetMessagingSettings()(*TeamMessagingSettings) {
    if m == nil {
        return nil
    } else {
        return m.messagingSettings
    }
}
// Gets the operations property value. The async operations that ran or are running on this team.
func (m *Team) GetOperations()([]TeamsAsyncOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// Gets the primaryChannel property value. The general channel for the team.
func (m *Team) GetPrimaryChannel()(*Channel) {
    if m == nil {
        return nil
    } else {
        return m.primaryChannel
    }
}
// Gets the schedule property value. The schedule of shifts for this team.
func (m *Team) GetSchedule()(*Schedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// Gets the specialization property value. Optional. Indicates whether the team is intended for a particular use case.  Each team specialization has access to unique behaviors and experiences targeted to its use case.
func (m *Team) GetSpecialization()(*TeamSpecialization) {
    if m == nil {
        return nil
    } else {
        return m.specialization
    }
}
// Gets the template property value. The template this team was created from. See available templates.
func (m *Team) GetTemplate()(*TeamsTemplate) {
    if m == nil {
        return nil
    } else {
        return m.template
    }
}
// Gets the visibility property value. The visibility of the group and team. Defaults to Public.
func (m *Team) GetVisibility()(*TeamVisibilityType) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// Gets the webUrl property value. A hyperlink that will go to the team in the Microsoft Teams client. This is the URL that you get when you right-click a team in the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed.
func (m *Team) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *Team) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["channels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChannel() })
        if err != nil {
            return err
        }
        res := make([]Channel, len(val))
        for i, v := range val {
            res[i] = *(v.(*Channel))
        }
        m.SetChannels(res)
        return nil
    }
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClassification(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["funSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamFunSettings() })
        if err != nil {
            return err
        }
        m.SetFunSettings(val.(*TeamFunSettings))
        return nil
    }
    res["group"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        m.SetGroup(val.(*Group))
        return nil
    }
    res["guestSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamGuestSettings() })
        if err != nil {
            return err
        }
        m.SetGuestSettings(val.(*TeamGuestSettings))
        return nil
    }
    res["installedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppInstallation() })
        if err != nil {
            return err
        }
        res := make([]TeamsAppInstallation, len(val))
        for i, v := range val {
            res[i] = *(v.(*TeamsAppInstallation))
        }
        m.SetInstalledApps(res)
        return nil
    }
    res["internalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInternalId(val)
        return nil
    }
    res["isArchived"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsArchived(val)
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversationMember() })
        if err != nil {
            return err
        }
        res := make([]ConversationMember, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConversationMember))
        }
        m.SetMembers(res)
        return nil
    }
    res["memberSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamMemberSettings() })
        if err != nil {
            return err
        }
        m.SetMemberSettings(val.(*TeamMemberSettings))
        return nil
    }
    res["messagingSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamMessagingSettings() })
        if err != nil {
            return err
        }
        m.SetMessagingSettings(val.(*TeamMessagingSettings))
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAsyncOperation() })
        if err != nil {
            return err
        }
        res := make([]TeamsAsyncOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*TeamsAsyncOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["primaryChannel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChannel() })
        if err != nil {
            return err
        }
        m.SetPrimaryChannel(val.(*Channel))
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSchedule() })
        if err != nil {
            return err
        }
        m.SetSchedule(val.(*Schedule))
        return nil
    }
    res["specialization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamSpecialization)
        if err != nil {
            return err
        }
        cast := val.(TeamSpecialization)
        m.SetSpecialization(&cast)
        return nil
    }
    res["template"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsTemplate() })
        if err != nil {
            return err
        }
        m.SetTemplate(val.(*TeamsTemplate))
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamVisibilityType)
        if err != nil {
            return err
        }
        cast := val.(TeamVisibilityType)
        m.SetVisibility(&cast)
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *Team) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Team) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
// Sets the channels property value. The collection of channels & messages associated with the team.
// Parameters:
//  - value : Value to set for the channels property.
func (m *Team) SetChannels(value []Channel)() {
    m.channels = value
}
// Sets the classification property value. An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.
// Parameters:
//  - value : Value to set for the classification property.
func (m *Team) SetClassification(value *string)() {
    m.classification = value
}
// Sets the createdDateTime property value. Timestamp at which the team was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Team) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. An optional description for the team. Maximum length: 1024 characters.
// Parameters:
//  - value : Value to set for the description property.
func (m *Team) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The name of the team.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Team) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the funSettings property value. Settings to configure use of Giphy, memes, and stickers in the team.
// Parameters:
//  - value : Value to set for the funSettings property.
func (m *Team) SetFunSettings(value *TeamFunSettings)() {
    m.funSettings = value
}
// Sets the group property value. 
// Parameters:
//  - value : Value to set for the group property.
func (m *Team) SetGroup(value *Group)() {
    m.group = value
}
// Sets the guestSettings property value. Settings to configure whether guests can create, update, or delete channels in the team.
// Parameters:
//  - value : Value to set for the guestSettings property.
func (m *Team) SetGuestSettings(value *TeamGuestSettings)() {
    m.guestSettings = value
}
// Sets the installedApps property value. The apps installed in this team.
// Parameters:
//  - value : Value to set for the installedApps property.
func (m *Team) SetInstalledApps(value []TeamsAppInstallation)() {
    m.installedApps = value
}
// Sets the internalId property value. A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.
// Parameters:
//  - value : Value to set for the internalId property.
func (m *Team) SetInternalId(value *string)() {
    m.internalId = value
}
// Sets the isArchived property value. Whether this team is in read-only mode.
// Parameters:
//  - value : Value to set for the isArchived property.
func (m *Team) SetIsArchived(value *bool)() {
    m.isArchived = value
}
// Sets the members property value. Members and owners of the team.
// Parameters:
//  - value : Value to set for the members property.
func (m *Team) SetMembers(value []ConversationMember)() {
    m.members = value
}
// Sets the memberSettings property value. Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
// Parameters:
//  - value : Value to set for the memberSettings property.
func (m *Team) SetMemberSettings(value *TeamMemberSettings)() {
    m.memberSettings = value
}
// Sets the messagingSettings property value. Settings to configure messaging and mentions in the team.
// Parameters:
//  - value : Value to set for the messagingSettings property.
func (m *Team) SetMessagingSettings(value *TeamMessagingSettings)() {
    m.messagingSettings = value
}
// Sets the operations property value. The async operations that ran or are running on this team.
// Parameters:
//  - value : Value to set for the operations property.
func (m *Team) SetOperations(value []TeamsAsyncOperation)() {
    m.operations = value
}
// Sets the primaryChannel property value. The general channel for the team.
// Parameters:
//  - value : Value to set for the primaryChannel property.
func (m *Team) SetPrimaryChannel(value *Channel)() {
    m.primaryChannel = value
}
// Sets the schedule property value. The schedule of shifts for this team.
// Parameters:
//  - value : Value to set for the schedule property.
func (m *Team) SetSchedule(value *Schedule)() {
    m.schedule = value
}
// Sets the specialization property value. Optional. Indicates whether the team is intended for a particular use case.  Each team specialization has access to unique behaviors and experiences targeted to its use case.
// Parameters:
//  - value : Value to set for the specialization property.
func (m *Team) SetSpecialization(value *TeamSpecialization)() {
    m.specialization = value
}
// Sets the template property value. The template this team was created from. See available templates.
// Parameters:
//  - value : Value to set for the template property.
func (m *Team) SetTemplate(value *TeamsTemplate)() {
    m.template = value
}
// Sets the visibility property value. The visibility of the group and team. Defaults to Public.
// Parameters:
//  - value : Value to set for the visibility property.
func (m *Team) SetVisibility(value *TeamVisibilityType)() {
    m.visibility = value
}
// Sets the webUrl property value. A hyperlink that will go to the team in the Microsoft Teams client. This is the URL that you get when you right-click a team in the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *Team) SetWebUrl(value *string)() {
    m.webUrl = value
}
