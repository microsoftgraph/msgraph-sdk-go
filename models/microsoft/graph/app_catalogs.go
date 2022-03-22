package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppCatalogs provides operations to manage the appCatalogs singleton.
type AppCatalogs struct {
    Entity
    // 
    teamsApps []TeamsAppable;
}
// NewAppCatalogs instantiates a new appCatalogs and sets the default values.
func NewAppCatalogs()(*AppCatalogs) {
    m := &AppCatalogs{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppCatalogsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppCatalogsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAppCatalogs(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppCatalogs) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["teamsApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamsAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsAppable, len(val))
            for i, v := range val {
                res[i] = v.(TeamsAppable)
            }
            m.SetTeamsApps(res)
        }
        return nil
    }
    return res
}
// GetTeamsApps gets the teamsApps property value. 
func (m *AppCatalogs) GetTeamsApps()([]TeamsAppable) {
    if m == nil {
        return nil
    } else {
        return m.teamsApps
    }
}
// Serialize serializes information the current object
func (m *AppCatalogs) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetTeamsApps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTeamsApps()))
        for i, v := range m.GetTeamsApps() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("teamsApps", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTeamsApps sets the teamsApps property value. 
func (m *AppCatalogs) SetTeamsApps(value []TeamsAppable)() {
    if m != nil {
        m.teamsApps = value
    }
}
