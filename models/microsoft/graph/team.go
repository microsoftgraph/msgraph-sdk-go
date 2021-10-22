package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Team struct {
    Entity
    channels []Channel;
    classification *string;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    description *string;
    displayName *string;
    funSettings *TeamFunSettings;
    group *Group;
    guestSettings *TeamGuestSettings;
    installedApps []TeamsAppInstallation;
    internalId *string;
    isArchived *bool;
    members []ConversationMember;
    memberSettings *TeamMemberSettings;
    messagingSettings *TeamMessagingSettings;
    operations []TeamsAsyncOperation;
    primaryChannel *Channel;
    schedule *Schedule;
    specialization *TeamSpecialization;
    template *TeamsTemplate;
    visibility *TeamVisibilityType;
    webUrl *string;
}
func NewTeam()(*Team) {
    m := &Team{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Team) GetChannels()([]Channel) {
    if m == nil {
        return nil
    } else {
        return m.channels
    }
}
func (m *Team) GetClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
func (m *Team) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *Team) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *Team) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Team) GetFunSettings()(*TeamFunSettings) {
    if m == nil {
        return nil
    } else {
        return m.funSettings
    }
}
func (m *Team) GetGroup()(*Group) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
func (m *Team) GetGuestSettings()(*TeamGuestSettings) {
    if m == nil {
        return nil
    } else {
        return m.guestSettings
    }
}
func (m *Team) GetInstalledApps()([]TeamsAppInstallation) {
    if m == nil {
        return nil
    } else {
        return m.installedApps
    }
}
func (m *Team) GetInternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalId
    }
}
func (m *Team) GetIsArchived()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isArchived
    }
}
func (m *Team) GetMembers()([]ConversationMember) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
func (m *Team) GetMemberSettings()(*TeamMemberSettings) {
    if m == nil {
        return nil
    } else {
        return m.memberSettings
    }
}
func (m *Team) GetMessagingSettings()(*TeamMessagingSettings) {
    if m == nil {
        return nil
    } else {
        return m.messagingSettings
    }
}
func (m *Team) GetOperations()([]TeamsAsyncOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
func (m *Team) GetPrimaryChannel()(*Channel) {
    if m == nil {
        return nil
    } else {
        return m.primaryChannel
    }
}
func (m *Team) GetSchedule()(*Schedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
func (m *Team) GetSpecialization()(*TeamSpecialization) {
    if m == nil {
        return nil
    } else {
        return m.specialization
    }
}
func (m *Team) GetTemplate()(*TeamsTemplate) {
    if m == nil {
        return nil
    } else {
        return m.template
    }
}
func (m *Team) GetVisibility()(*TeamVisibilityType) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
func (m *Team) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
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
func (m *Team) SetChannels(value []Channel)() {
    m.channels = value
}
func (m *Team) SetClassification(value *string)() {
    m.classification = value
}
func (m *Team) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *Team) SetDescription(value *string)() {
    m.description = value
}
func (m *Team) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Team) SetFunSettings(value *TeamFunSettings)() {
    m.funSettings = value
}
func (m *Team) SetGroup(value *Group)() {
    m.group = value
}
func (m *Team) SetGuestSettings(value *TeamGuestSettings)() {
    m.guestSettings = value
}
func (m *Team) SetInstalledApps(value []TeamsAppInstallation)() {
    m.installedApps = value
}
func (m *Team) SetInternalId(value *string)() {
    m.internalId = value
}
func (m *Team) SetIsArchived(value *bool)() {
    m.isArchived = value
}
func (m *Team) SetMembers(value []ConversationMember)() {
    m.members = value
}
func (m *Team) SetMemberSettings(value *TeamMemberSettings)() {
    m.memberSettings = value
}
func (m *Team) SetMessagingSettings(value *TeamMessagingSettings)() {
    m.messagingSettings = value
}
func (m *Team) SetOperations(value []TeamsAsyncOperation)() {
    m.operations = value
}
func (m *Team) SetPrimaryChannel(value *Channel)() {
    m.primaryChannel = value
}
func (m *Team) SetSchedule(value *Schedule)() {
    m.schedule = value
}
func (m *Team) SetSpecialization(value *TeamSpecialization)() {
    m.specialization = value
}
func (m *Team) SetTemplate(value *TeamsTemplate)() {
    m.template = value
}
func (m *Team) SetVisibility(value *TeamVisibilityType)() {
    m.visibility = value
}
func (m *Team) SetWebUrl(value *string)() {
    m.webUrl = value
}
