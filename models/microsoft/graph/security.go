package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Security provides operations to manage the security singleton.
type Security struct {
    Entity
    // Read-only. Nullable.
    alerts []Alertable;
    // 
    secureScoreControlProfiles []SecureScoreControlProfileable;
    // 
    secureScores []SecureScoreable;
}
// NewSecurity instantiates a new security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecurityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSecurity(), nil
}
// GetAlerts gets the alerts property value. Read-only. Nullable.
func (m *Security) GetAlerts()([]Alertable) {
    if m == nil {
        return nil
    } else {
        return m.alerts
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alerts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAlertFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Alertable, len(val))
            for i, v := range val {
                res[i] = v.(Alertable)
            }
            m.SetAlerts(res)
        }
        return nil
    }
    res["secureScoreControlProfiles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreControlProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreControlProfileable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoreControlProfileable)
            }
            m.SetSecureScoreControlProfiles(res)
        }
        return nil
    }
    res["secureScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoreable)
            }
            m.SetSecureScores(res)
        }
        return nil
    }
    return res
}
// GetSecureScoreControlProfiles gets the secureScoreControlProfiles property value. 
func (m *Security) GetSecureScoreControlProfiles()([]SecureScoreControlProfileable) {
    if m == nil {
        return nil
    } else {
        return m.secureScoreControlProfiles
    }
}
// GetSecureScores gets the secureScores property value. 
func (m *Security) GetSecureScores()([]SecureScoreable) {
    if m == nil {
        return nil
    } else {
        return m.secureScores
    }
}
// Serialize serializes information the current object
func (m *Security) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlerts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlerts()))
        for i, v := range m.GetAlerts() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("alerts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScoreControlProfiles() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecureScoreControlProfiles()))
        for i, v := range m.GetSecureScoreControlProfiles() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("secureScoreControlProfiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecureScores() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSecureScores()))
        for i, v := range m.GetSecureScores() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("secureScores", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlerts sets the alerts property value. Read-only. Nullable.
func (m *Security) SetAlerts(value []Alertable)() {
    if m != nil {
        m.alerts = value
    }
}
// SetSecureScoreControlProfiles sets the secureScoreControlProfiles property value. 
func (m *Security) SetSecureScoreControlProfiles(value []SecureScoreControlProfileable)() {
    if m != nil {
        m.secureScoreControlProfiles = value
    }
}
// SetSecureScores sets the secureScores property value. 
func (m *Security) SetSecureScores(value []SecureScoreable)() {
    if m != nil {
        m.secureScores = value
    }
}
