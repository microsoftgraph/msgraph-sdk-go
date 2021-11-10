package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Teamwork struct {
    Entity
    // 
    workforceIntegrations []WorkforceIntegration;
}
// Instantiates a new teamwork and sets the default values.
func NewTeamwork()(*Teamwork) {
    m := &Teamwork{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the workforceIntegrations property value. 
func (m *Teamwork) GetWorkforceIntegrations()([]WorkforceIntegration) {
    if m == nil {
        return nil
    } else {
        return m.workforceIntegrations
    }
}
// The deserialization information for the current model
func (m *Teamwork) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["workforceIntegrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkforceIntegration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkforceIntegration, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkforceIntegration))
            }
            m.SetWorkforceIntegrations(res)
        }
        return nil
    }
    return res
}
func (m *Teamwork) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Teamwork) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkforceIntegrations()))
        for i, v := range m.GetWorkforceIntegrations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("workforceIntegrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the workforceIntegrations property value. 
// Parameters:
//  - value : Value to set for the workforceIntegrations property.
func (m *Teamwork) SetWorkforceIntegrations(value []WorkforceIntegration)() {
    m.workforceIntegrations = value
}
