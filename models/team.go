package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Team provides operations to manage the collection of agreement entities.
type Team struct {
    Entity
    // List of channels either hosted in or shared with the team (incoming channels).
    allChannels []Channelable
    // The collection of channels and messages associated with the team.
    channels []Channelable
    // An optional label. Typically describes the data or business sensitivity of the team. Must match one of a pre-configured set in the tenant's directory.
    classification *string
    // Timestamp at which the team was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // An optional description for the team. Maximum length: 1024 characters.
    description *string
    // The name of the team.
    displayName *string
    // Settings to configure use of Giphy, memes, and stickers in the team.
    funSettings TeamFunSettingsable
    // The group property
    group Groupable
    // Settings to configure whether guests can create, update, or delete channels in the team.
    guestSettings TeamGuestSettingsable
    // List of channels shared with the team.
    incomingChannels []Channelable
    // The apps installed in this team.
    installedApps []TeamsAppInstallationable
    // A unique ID for the team that has been used in a few places such as the audit log/Office 365 Management Activity API.
    internalId *string
    // Whether this team is in read-only mode.
    isArchived *bool
    // Members and owners of the team.
    members []ConversationMemberable
    // Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
    memberSettings TeamMemberSettingsable
    // Settings to configure messaging and mentions in the team.
    messagingSettings TeamMessagingSettingsable
    // The async operations that ran or are running on this team.
    operations []TeamsAsyncOperationable
    // The team photo.
    photo ProfilePhotoable
    // The general channel for the team.
    primaryChannel Channelable
    // The schedule of shifts for this team.
    schedule Scheduleable
    // Optional. Indicates whether the team is intended for a particular use case.  Each team specialization has access to unique behaviors and experiences targeted to its use case.
    specialization *TeamSpecialization
    // Contains summary information about the team, including number of owners, members, and guests.
    summary TeamSummaryable
    // The template this team was created from. See available templates.
    template TeamsTemplateable
    // The ID of the Azure Active Directory tenant.
    tenantId *string
    // The visibility of the group and team. Defaults to Public.
    visibility *TeamVisibilityType
    // A hyperlink that will go to the team in the Microsoft Teams client. This is the URL that you get when you right-click a team in the Microsoft Teams client and select Get link to team. This URL should be treated as an opaque blob, and not parsed.
    webUrl *string
}
// NewTeam instantiates a new team and sets the default values.
func NewTeam()(*Team) {
    m := &Team{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.team";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTeamFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeam(), nil
}
// GetAllChannels gets the allChannels property value. List of channels either hosted in or shared with the team (incoming channels).
func (m *Team) GetAllChannels()([]Channelable) {
    if m == nil {
        return nil
    } else {
        return m.allChannels
    }
}
// GetChannels gets the channels property value. The collection of channels and messages associated with the team.
func (m *Team) GetChannels()([]Channelable) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *Team) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChannelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Channelable, len(val))
            for i, v := range val {
                res[i] = v.(Channelable)
            }
            m.SetAllChannels(res)
        }
        return nil
    }
    res["channels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChannelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Channelable, len(val))
            for i, v := range val {
                res[i] = v.(Channelable)
            }
            m.SetChannels(res)
        }
        return nil
    }
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["funSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFunSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFunSettings(val.(TeamFunSettingsable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Groupable))
        }
        return nil
    }
    res["guestSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamGuestSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuestSettings(val.(TeamGuestSettingsable))
        }
        return nil
    }
    res["incomingChannels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChannelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Channelable, len(val))
            for i, v := range val {
                res[i] = v.(Channelable)
            }
            m.SetIncomingChannels(res)
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
    res["internalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalId(val)
        }
        return nil
    }
    res["isArchived"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsArchived(val)
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
    res["memberSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamMemberSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberSettings(val.(TeamMemberSettingsable))
        }
        return nil
    }
    res["messagingSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamMessagingSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessagingSettings(val.(TeamMessagingSettingsable))
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamsAsyncOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAsyncOperationable, len(val))
            for i, v := range val {
                res[i] = v.(TeamsAsyncOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["photo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfilePhotoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(ProfilePhotoable))
        }
        return nil
    }
    res["primaryChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChannelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryChannel(val.(Channelable))
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(Scheduleable))
        }
        return nil
    }
    res["specialization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamSpecialization)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecialization(val.(*TeamSpecialization))
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val.(TeamSummaryable))
        }
        return nil
    }
    res["template"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplate(val.(TeamsTemplateable))
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
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamVisibilityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val.(*TeamVisibilityType))
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
// GetFunSettings gets the funSettings property value. Settings to configure use of Giphy, memes, and stickers in the team.
func (m *Team) GetFunSettings()(TeamFunSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.funSettings
    }
}
// GetGroup gets the group property value. The group property
func (m *Team) GetGroup()(Groupable) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// GetGuestSettings gets the guestSettings property value. Settings to configure whether guests can create, update, or delete channels in the team.
func (m *Team) GetGuestSettings()(TeamGuestSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.guestSettings
    }
}
// GetIncomingChannels gets the incomingChannels property value. List of channels shared with the team.
func (m *Team) GetIncomingChannels()([]Channelable) {
    if m == nil {
        return nil
    } else {
        return m.incomingChannels
    }
}
// GetInstalledApps gets the installedApps property value. The apps installed in this team.
func (m *Team) GetInstalledApps()([]TeamsAppInstallationable) {
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
func (m *Team) GetMembers()([]ConversationMemberable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetMemberSettings gets the memberSettings property value. Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
func (m *Team) GetMemberSettings()(TeamMemberSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.memberSettings
    }
}
// GetMessagingSettings gets the messagingSettings property value. Settings to configure messaging and mentions in the team.
func (m *Team) GetMessagingSettings()(TeamMessagingSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.messagingSettings
    }
}
// GetOperations gets the operations property value. The async operations that ran or are running on this team.
func (m *Team) GetOperations()([]TeamsAsyncOperationable) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPhoto gets the photo property value. The team photo.
func (m *Team) GetPhoto()(ProfilePhotoable) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// GetPrimaryChannel gets the primaryChannel property value. The general channel for the team.
func (m *Team) GetPrimaryChannel()(Channelable) {
    if m == nil {
        return nil
    } else {
        return m.primaryChannel
    }
}
// GetSchedule gets the schedule property value. The schedule of shifts for this team.
func (m *Team) GetSchedule()(Scheduleable) {
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
// GetSummary gets the summary property value. Contains summary information about the team, including number of owners, members, and guests.
func (m *Team) GetSummary()(TeamSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
// GetTemplate gets the template property value. The template this team was created from. See available templates.
func (m *Team) GetTemplate()(TeamsTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.template
    }
}
// GetTenantId gets the tenantId property value. The ID of the Azure Active Directory tenant.
func (m *Team) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
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
// Serialize serializes information the current object
func (m *Team) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllChannels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllChannels()))
        for i, v := range m.GetAllChannels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("allChannels", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChannels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChannels()))
        for i, v := range m.GetChannels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetIncomingChannels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncomingChannels()))
        for i, v := range m.GetIncomingChannels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("incomingChannels", cast)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
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
        cast := (*m.GetSpecialization()).String()
        err = writer.WriteStringValue("specialization", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("summary", m.GetSummary())
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
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetVisibility() != nil {
        cast := (*m.GetVisibility()).String()
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
// SetAllChannels sets the allChannels property value. List of channels either hosted in or shared with the team (incoming channels).
func (m *Team) SetAllChannels(value []Channelable)() {
    if m != nil {
        m.allChannels = value
    }
}
// SetChannels sets the channels property value. The collection of channels and messages associated with the team.
func (m *Team) SetChannels(value []Channelable)() {
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
func (m *Team) SetFunSettings(value TeamFunSettingsable)() {
    if m != nil {
        m.funSettings = value
    }
}
// SetGroup sets the group property value. The group property
func (m *Team) SetGroup(value Groupable)() {
    if m != nil {
        m.group = value
    }
}
// SetGuestSettings sets the guestSettings property value. Settings to configure whether guests can create, update, or delete channels in the team.
func (m *Team) SetGuestSettings(value TeamGuestSettingsable)() {
    if m != nil {
        m.guestSettings = value
    }
}
// SetIncomingChannels sets the incomingChannels property value. List of channels shared with the team.
func (m *Team) SetIncomingChannels(value []Channelable)() {
    if m != nil {
        m.incomingChannels = value
    }
}
// SetInstalledApps sets the installedApps property value. The apps installed in this team.
func (m *Team) SetInstalledApps(value []TeamsAppInstallationable)() {
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
func (m *Team) SetMembers(value []ConversationMemberable)() {
    if m != nil {
        m.members = value
    }
}
// SetMemberSettings sets the memberSettings property value. Settings to configure whether members can perform certain actions, for example, create channels and add bots, in the team.
func (m *Team) SetMemberSettings(value TeamMemberSettingsable)() {
    if m != nil {
        m.memberSettings = value
    }
}
// SetMessagingSettings sets the messagingSettings property value. Settings to configure messaging and mentions in the team.
func (m *Team) SetMessagingSettings(value TeamMessagingSettingsable)() {
    if m != nil {
        m.messagingSettings = value
    }
}
// SetOperations sets the operations property value. The async operations that ran or are running on this team.
func (m *Team) SetOperations(value []TeamsAsyncOperationable)() {
    if m != nil {
        m.operations = value
    }
}
// SetPhoto sets the photo property value. The team photo.
func (m *Team) SetPhoto(value ProfilePhotoable)() {
    if m != nil {
        m.photo = value
    }
}
// SetPrimaryChannel sets the primaryChannel property value. The general channel for the team.
func (m *Team) SetPrimaryChannel(value Channelable)() {
    if m != nil {
        m.primaryChannel = value
    }
}
// SetSchedule sets the schedule property value. The schedule of shifts for this team.
func (m *Team) SetSchedule(value Scheduleable)() {
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
// SetSummary sets the summary property value. Contains summary information about the team, including number of owners, members, and guests.
func (m *Team) SetSummary(value TeamSummaryable)() {
    if m != nil {
        m.summary = value
    }
}
// SetTemplate sets the template property value. The template this team was created from. See available templates.
func (m *Team) SetTemplate(value TeamsTemplateable)() {
    if m != nil {
        m.template = value
    }
}
// SetTenantId sets the tenantId property value. The ID of the Azure Active Directory tenant.
func (m *Team) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
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
