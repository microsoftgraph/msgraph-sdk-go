package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintService provides operations to manage the print singleton.
type PrintService struct {
    Entity
    // Endpoints that can be used to access the service. Read-only. Nullable.
    endpoints []PrintServiceEndpointable;
}
// NewPrintService instantiates a new printService and sets the default values.
func NewPrintService()(*PrintService) {
    m := &PrintService{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintServiceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintService(), nil
}
// GetEndpoints gets the endpoints property value. Endpoints that can be used to access the service. Read-only. Nullable.
func (m *PrintService) GetEndpoints()([]PrintServiceEndpointable) {
    if m == nil {
        return nil
    } else {
        return m.endpoints
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintService) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["endpoints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintServiceEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintServiceEndpointable, len(val))
            for i, v := range val {
                res[i] = v.(PrintServiceEndpointable)
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
// Serialize serializes information the current object
func (m *PrintService) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEndpoints() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEndpoints()))
        for i, v := range m.GetEndpoints() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("endpoints", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndpoints sets the endpoints property value. Endpoints that can be used to access the service. Read-only. Nullable.
func (m *PrintService) SetEndpoints(value []PrintServiceEndpointable)() {
    if m != nil {
        m.endpoints = value
    }
}
