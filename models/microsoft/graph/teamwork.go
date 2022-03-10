package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Teamwork provides operations to manage the teamwork singleton.
type Teamwork struct {
    Entity
    // A workforce integration with shifts.
    workforceIntegrations []WorkforceIntegrationable;
}
// NewTeamwork instantiates a new teamwork and sets the default values.
func NewTeamwork()(*Teamwork) {
    m := &Teamwork{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamwork(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Teamwork) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["workforceIntegrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkforceIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkforceIntegrationable, len(val))
            for i, v := range val {
                res[i] = v.(WorkforceIntegrationable)
            }
            m.SetWorkforceIntegrations(res)
        }
        return nil
    }
    return res
}
// GetWorkforceIntegrations gets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) GetWorkforceIntegrations()([]WorkforceIntegrationable) {
    if m == nil {
        return nil
    } else {
        return m.workforceIntegrations
    }
}
func (m *Teamwork) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Teamwork) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetWorkforceIntegrations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkforceIntegrations()))
        for i, v := range m.GetWorkforceIntegrations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("workforceIntegrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWorkforceIntegrations sets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) SetWorkforceIntegrations(value []WorkforceIntegrationable)() {
    if m != nil {
        m.workforceIntegrations = value
    }
}
