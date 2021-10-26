package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamsTab struct {
    Entity
    // Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
    configuration *TeamsTabConfiguration;
    // Name of the tab.
    displayName *string;
    // The application that is linked to the tab. This cannot be changed after tab creation.
    teamsApp *TeamsApp;
    // Deep link URL of the tab instance. Read only.
    webUrl *string;
}
// Instantiates a new teamsTab and sets the default values.
func NewTeamsTab()(*TeamsTab) {
    m := &TeamsTab{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) GetConfiguration()(*TeamsTabConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// Gets the displayName property value. Name of the tab.
func (m *TeamsTab) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the teamsApp property value. The application that is linked to the tab. This cannot be changed after tab creation.
func (m *TeamsTab) GetTeamsApp()(*TeamsApp) {
    if m == nil {
        return nil
    } else {
        return m.teamsApp
    }
}
// Gets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *TeamsTab) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsTabConfiguration() })
        if err != nil {
            return err
        }
        m.SetConfiguration(val.(*TeamsTabConfiguration))
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
    res["teamsApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsApp() })
        if err != nil {
            return err
        }
        m.SetTeamsApp(val.(*TeamsApp))
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
func (m *TeamsTab) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamsTab) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
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
        err = writer.WriteObjectValue("teamsApp", m.GetTeamsApp())
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
// Sets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
// Parameters:
//  - value : Value to set for the configuration property.
func (m *TeamsTab) SetConfiguration(value *TeamsTabConfiguration)() {
    m.configuration = value
}
// Sets the displayName property value. Name of the tab.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TeamsTab) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the teamsApp property value. The application that is linked to the tab. This cannot be changed after tab creation.
// Parameters:
//  - value : Value to set for the teamsApp property.
func (m *TeamsTab) SetTeamsApp(value *TeamsApp)() {
    m.teamsApp = value
}
// Sets the webUrl property value. Deep link URL of the tab instance. Read only.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *TeamsTab) SetWebUrl(value *string)() {
    m.webUrl = value
}
