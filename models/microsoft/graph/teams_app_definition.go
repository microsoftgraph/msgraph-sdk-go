package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamsAppDefinition struct {
    Entity
    // The details of the bot specified in the Teams app manifest.
    bot *TeamworkBot;
    // 
    createdBy *IdentitySet;
    // Verbose description of the application.
    description *string;
    // The name of the app provided by the app developer.
    displayName *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
    publishingState *TeamsAppPublishingState;
    // Short description of the application.
    shortDescription *string;
    // The ID from the Teams app manifest.
    teamsAppId *string;
    // The version number of the application.
    version *string;
}
// Instantiates a new teamsAppDefinition and sets the default values.
func NewTeamsAppDefinition()(*TeamsAppDefinition) {
    m := &TeamsAppDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the bot property value. The details of the bot specified in the Teams app manifest.
func (m *TeamsAppDefinition) GetBot()(*TeamworkBot) {
    if m == nil {
        return nil
    } else {
        return m.bot
    }
}
// Gets the createdBy property value. 
func (m *TeamsAppDefinition) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the description property value. Verbose description of the application.
func (m *TeamsAppDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of the app provided by the app developer.
func (m *TeamsAppDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *TeamsAppDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
func (m *TeamsAppDefinition) GetPublishingState()(*TeamsAppPublishingState) {
    if m == nil {
        return nil
    } else {
        return m.publishingState
    }
}
// Gets the shortDescription property value. Short description of the application.
func (m *TeamsAppDefinition) GetShortDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortDescription
    }
}
// Gets the teamsAppId property value. The ID from the Teams app manifest.
func (m *TeamsAppDefinition) GetTeamsAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppId
    }
}
// Gets the version property value. The version number of the application.
func (m *TeamsAppDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *TeamsAppDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["bot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkBot() })
        if err != nil {
            return err
        }
        m.SetBot(val.(*TeamworkBot))
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*IdentitySet))
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["publishingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppPublishingState)
        if err != nil {
            return err
        }
        cast := val.(TeamsAppPublishingState)
        m.SetPublishingState(&cast)
        return nil
    }
    res["shortDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetShortDescription(val)
        return nil
    }
    res["teamsAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTeamsAppId(val)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *TeamsAppDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamsAppDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("bot", m.GetBot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPublishingState() != nil {
        cast := m.GetPublishingState().String()
        err = writer.WriteStringValue("publishingState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shortDescription", m.GetShortDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the bot property value. The details of the bot specified in the Teams app manifest.
// Parameters:
//  - value : Value to set for the bot property.
func (m *TeamsAppDefinition) SetBot(value *TeamworkBot)() {
    m.bot = value
}
// Sets the createdBy property value. 
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *TeamsAppDefinition) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the description property value. Verbose description of the application.
// Parameters:
//  - value : Value to set for the description property.
func (m *TeamsAppDefinition) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The name of the app provided by the app developer.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TeamsAppDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *TeamsAppDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
// Parameters:
//  - value : Value to set for the publishingState property.
func (m *TeamsAppDefinition) SetPublishingState(value *TeamsAppPublishingState)() {
    m.publishingState = value
}
// Sets the shortDescription property value. Short description of the application.
// Parameters:
//  - value : Value to set for the shortDescription property.
func (m *TeamsAppDefinition) SetShortDescription(value *string)() {
    m.shortDescription = value
}
// Sets the teamsAppId property value. The ID from the Teams app manifest.
// Parameters:
//  - value : Value to set for the teamsAppId property.
func (m *TeamsAppDefinition) SetTeamsAppId(value *string)() {
    m.teamsAppId = value
}
// Sets the version property value. The version number of the application.
// Parameters:
//  - value : Value to set for the version property.
func (m *TeamsAppDefinition) SetVersion(value *string)() {
    m.version = value
}
