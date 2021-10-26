package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserTeamwork struct {
    Entity
    // The apps installed in the personal scope of this user.
    installedApps []UserScopeTeamsAppInstallation;
}
// Instantiates a new userTeamwork and sets the default values.
func NewUserTeamwork()(*UserTeamwork) {
    m := &UserTeamwork{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the installedApps property value. The apps installed in the personal scope of this user.
func (m *UserTeamwork) GetInstalledApps()([]UserScopeTeamsAppInstallation) {
    if m == nil {
        return nil
    } else {
        return m.installedApps
    }
}
// The deserialization information for the current model
func (m *UserTeamwork) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["installedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserScopeTeamsAppInstallation() })
        if err != nil {
            return err
        }
        res := make([]UserScopeTeamsAppInstallation, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserScopeTeamsAppInstallation))
        }
        m.SetInstalledApps(res)
        return nil
    }
    return res
}
func (m *UserTeamwork) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserTeamwork) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    return nil
}
// Sets the installedApps property value. The apps installed in the personal scope of this user.
// Parameters:
//  - value : Value to set for the installedApps property.
func (m *UserTeamwork) SetInstalledApps(value []UserScopeTeamsAppInstallation)() {
    m.installedApps = value
}
