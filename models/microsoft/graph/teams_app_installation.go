package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamsAppInstallation struct {
    Entity
    // The app that is installed.
    teamsApp *TeamsApp;
    // The details of this version of the app.
    teamsAppDefinition *TeamsAppDefinition;
}
// Instantiates a new teamsAppInstallation and sets the default values.
func NewTeamsAppInstallation()(*TeamsAppInstallation) {
    m := &TeamsAppInstallation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the teamsApp property value. The app that is installed.
func (m *TeamsAppInstallation) GetTeamsApp()(*TeamsApp) {
    if m == nil {
        return nil
    } else {
        return m.teamsApp
    }
}
// Gets the teamsAppDefinition property value. The details of this version of the app.
func (m *TeamsAppInstallation) GetTeamsAppDefinition()(*TeamsAppDefinition) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppDefinition
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the teamsApp property value. The app that is installed.
// Parameters:
//  - value : Value to set for the teamsApp property.
func (m *TeamsAppInstallation) SetTeamsApp(value *TeamsApp)() {
    m.teamsApp = value
}
// Sets the teamsAppDefinition property value. The details of this version of the app.
// Parameters:
//  - value : Value to set for the teamsAppDefinition property.
func (m *TeamsAppInstallation) SetTeamsAppDefinition(value *TeamsAppDefinition)() {
    m.teamsAppDefinition = value
}
