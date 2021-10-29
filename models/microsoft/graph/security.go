package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Security struct {
    Entity
    // Read-only. Nullable.
    alerts []Alert;
    // 
    secureScoreControlProfiles []SecureScoreControlProfile;
    // 
    secureScores []SecureScore;
}
// Instantiates a new security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the alerts property value. Read-only. Nullable.
func (m *Security) GetAlerts()([]Alert) {
    if m == nil {
        return nil
    } else {
        return m.alerts
    }
}
// Gets the secureScoreControlProfiles property value. 
func (m *Security) GetSecureScoreControlProfiles()([]SecureScoreControlProfile) {
    if m == nil {
        return nil
    } else {
        return m.secureScoreControlProfiles
    }
}
// Gets the secureScores property value. 
func (m *Security) GetSecureScores()([]SecureScore) {
    if m == nil {
        return nil
    } else {
        return m.secureScores
    }
}
// The deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlert() })
        if err != nil {
            return err
        }
        res := make([]Alert, len(val))
        for i, v := range val {
            res[i] = *(v.(*Alert))
        }
        m.SetAlerts(res)
        return nil
    }
    res["secureScoreControlProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecureScoreControlProfile() })
        if err != nil {
            return err
        }
        res := make([]SecureScoreControlProfile, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecureScoreControlProfile))
        }
        m.SetSecureScoreControlProfiles(res)
        return nil
    }
    res["secureScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecureScore() })
        if err != nil {
            return err
        }
        res := make([]SecureScore, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecureScore))
        }
        m.SetSecureScores(res)
        return nil
    }
    return res
}
func (m *Security) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Security) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecureScoreControlProfiles()))
        for i, v := range m.GetSecureScoreControlProfiles() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("secureScoreControlProfiles", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecureScores()))
        for i, v := range m.GetSecureScores() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("secureScores", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the alerts property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the alerts property.
func (m *Security) SetAlerts(value []Alert)() {
    m.alerts = value
}
// Sets the secureScoreControlProfiles property value. 
// Parameters:
//  - value : Value to set for the secureScoreControlProfiles property.
func (m *Security) SetSecureScoreControlProfiles(value []SecureScoreControlProfile)() {
    m.secureScoreControlProfiles = value
}
// Sets the secureScores property value. 
// Parameters:
//  - value : Value to set for the secureScores property.
func (m *Security) SetSecureScores(value []SecureScore)() {
    m.secureScores = value
}
