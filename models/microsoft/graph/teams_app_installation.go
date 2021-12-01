package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsAppInstallation 
type TeamsAppInstallation struct {
    Entity
    // The app that is installed.
    teamsApp *TeamsApp;
    // The details of this version of the app.
    teamsAppDefinition *TeamsAppDefinition;
}
// NewTeamsAppInstallation instantiates a new teamsAppInstallation and sets the default values.
func NewTeamsAppInstallation()(*TeamsAppInstallation) {
    m := &TeamsAppInstallation{
        Entity: *NewEntity(),
    }
    return m
}
// GetTeamsApp gets the teamsApp property value. The app that is installed.
func (m *TeamsAppInstallation) GetTeamsApp()(*TeamsApp) {
    if m == nil {
        return nil
    } else {
        return m.teamsApp
    }
}
// GetTeamsAppDefinition gets the teamsAppDefinition property value. The details of this version of the app.
func (m *TeamsAppInstallation) GetTeamsAppDefinition()(*TeamsAppDefinition) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppDefinition
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppInstallation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["teamsAppDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppDefinition(val.(*TeamsAppDefinition))
        }
        return nil
    }
    return res
}
func (m *TeamsAppInstallation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamsAppInstallation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("teamsApp", m.GetTeamsApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsAppDefinition", m.GetTeamsAppDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTeamsApp sets the teamsApp property value. The app that is installed.
func (m *TeamsAppInstallation) SetTeamsApp(value *TeamsApp)() {
    if m != nil {
        m.teamsApp = value
    }
}
// SetTeamsAppDefinition sets the teamsAppDefinition property value. The details of this version of the app.
func (m *TeamsAppInstallation) SetTeamsAppDefinition(value *TeamsAppDefinition)() {
    if m != nil {
        m.teamsAppDefinition = value
    }
}
