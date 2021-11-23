package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsTab 
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
// NewTeamsTab instantiates a new teamsTab and sets the default values.
func NewTeamsTab()(*TeamsTab) {
    m := &TeamsTab{
        Entity: *NewEntity(),
    }
    return m
}
// GetConfiguration gets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) GetConfiguration()(*TeamsTabConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// GetDisplayName gets the displayName property value. Name of the tab.
func (m *TeamsTab) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetTeamsApp gets the teamsApp property value. The application that is linked to the tab. This cannot be changed after tab creation.
func (m *TeamsTab) GetTeamsApp()(*TeamsApp) {
    if m == nil {
        return nil
    } else {
        return m.teamsApp
    }
}
// GetWebUrl gets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsTab) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsTabConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(*TeamsTabConfiguration))
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
    res["teamsApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsApp() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsApp(val.(*TeamsApp))
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
func (m *TeamsTab) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetConfiguration sets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) SetConfiguration(value *TeamsTabConfiguration)() {
    m.configuration = value
}
// SetDisplayName sets the displayName property value. Name of the tab.
func (m *TeamsTab) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTeamsApp sets the teamsApp property value. The application that is linked to the tab. This cannot be changed after tab creation.
func (m *TeamsTab) SetTeamsApp(value *TeamsApp)() {
    m.teamsApp = value
}
// SetWebUrl sets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) SetWebUrl(value *string)() {
    m.webUrl = value
}
