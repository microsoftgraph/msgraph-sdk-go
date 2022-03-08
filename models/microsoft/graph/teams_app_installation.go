package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsAppInstallation provides operations to manage the collection of chat entities.
type TeamsAppInstallation struct {
    Entity
    // The app that is installed.
    teamsApp TeamsAppable;
    // The details of this version of the app.
    teamsAppDefinition TeamsAppDefinitionable;
}
// NewTeamsAppInstallation instantiates a new teamsAppInstallation and sets the default values.
func NewTeamsAppInstallation()(*TeamsAppInstallation) {
    m := &TeamsAppInstallation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppInstallationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppInstallationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamsAppInstallation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppInstallation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["teamsApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsApp(val.(TeamsAppable))
        }
        return nil
    }
    res["teamsAppDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppDefinition(val.(TeamsAppDefinitionable))
        }
        return nil
    }
    return res
}
// GetTeamsApp gets the teamsApp property value. The app that is installed.
func (m *TeamsAppInstallation) GetTeamsApp()(TeamsAppable) {
    if m == nil {
        return nil
    } else {
        return m.teamsApp
    }
}
// GetTeamsAppDefinition gets the teamsAppDefinition property value. The details of this version of the app.
func (m *TeamsAppInstallation) GetTeamsAppDefinition()(TeamsAppDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppDefinition
    }
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
func (m *TeamsAppInstallation) SetTeamsApp(value TeamsAppable)() {
    if m != nil {
        m.teamsApp = value
    }
}
// SetTeamsAppDefinition sets the teamsAppDefinition property value. The details of this version of the app.
func (m *TeamsAppInstallation) SetTeamsAppDefinition(value TeamsAppDefinitionable)() {
    if m != nil {
        m.teamsAppDefinition = value
    }
}
