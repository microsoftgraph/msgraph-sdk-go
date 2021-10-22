package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamsAppInstallation struct {
    Entity
    teamsApp *TeamsApp;
    teamsAppDefinition *TeamsAppDefinition;
}
func NewTeamsAppInstallation()(*TeamsAppInstallation) {
    m := &TeamsAppInstallation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TeamsAppInstallation) GetTeamsApp()(*TeamsApp) {
    if m == nil {
        return nil
    } else {
        return m.teamsApp
    }
}
func (m *TeamsAppInstallation) GetTeamsAppDefinition()(*TeamsAppDefinition) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppDefinition
    }
}
func (m *TeamsAppInstallation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["teamsApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsApp() })
        if err != nil {
            return err
        }
        m.SetTeamsApp(val.(*TeamsApp))
        return nil
    }
    res["teamsAppDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppDefinition() })
        if err != nil {
            return err
        }
        m.SetTeamsAppDefinition(val.(*TeamsAppDefinition))
        return nil
    }
    return res
}
func (m *TeamsAppInstallation) IsNil()(bool) {
    return m == nil
}
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
func (m *TeamsAppInstallation) SetTeamsApp(value *TeamsApp)() {
    m.teamsApp = value
}
func (m *TeamsAppInstallation) SetTeamsAppDefinition(value *TeamsAppDefinition)() {
    m.teamsAppDefinition = value
}
