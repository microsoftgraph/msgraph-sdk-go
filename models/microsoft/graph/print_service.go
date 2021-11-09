package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrintService struct {
    Entity
    // Endpoints that can be used to access the service. Read-only. Nullable.
    endpoints []PrintServiceEndpoint;
}
// Instantiates a new printService and sets the default values.
func NewPrintService()(*PrintService) {
    m := &PrintService{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the endpoints property value. Endpoints that can be used to access the service. Read-only. Nullable.
func (m *PrintService) GetEndpoints()([]PrintServiceEndpoint) {
    if m == nil {
        return nil
    } else {
        return m.endpoints
    }
}
// The deserialization information for the current model
func (m *PrintService) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["endpoints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintServiceEndpoint() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintServiceEndpoint, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintServiceEndpoint))
            }
            m.SetEndpoints(res)
        }
        return nil
    }
    return res
}
func (m *PrintService) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrintService) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEndpoints()))
        for i, v := range m.GetEndpoints() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("endpoints", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the endpoints property value. Endpoints that can be used to access the service. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the endpoints property.
func (m *PrintService) SetEndpoints(value []PrintServiceEndpoint)() {
    m.endpoints = value
}
