package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintServiceEndpoint 
type PrintServiceEndpoint struct {
    Entity
    // A human-readable display name for the endpoint.
    displayName *string;
    // The URI that can be used to access the service.
    uri *string;
}
// NewPrintServiceEndpoint instantiates a new printServiceEndpoint and sets the default values.
func NewPrintServiceEndpoint()(*PrintServiceEndpoint) {
    m := &PrintServiceEndpoint{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintServiceEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintServiceEndpointFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintServiceEndpoint(), nil
}
// GetDisplayName gets the displayName property value. A human-readable display name for the endpoint.
func (m *PrintServiceEndpoint) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintServiceEndpoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["uri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetUri gets the uri property value. The URI that can be used to access the service.
func (m *PrintServiceEndpoint) GetUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uri
    }
}
// Serialize serializes information the current object
func (m *PrintServiceEndpoint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uri", m.GetUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. A human-readable display name for the endpoint.
func (m *PrintServiceEndpoint) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetUri sets the uri property value. The URI that can be used to access the service.
func (m *PrintServiceEndpoint) SetUri(value *string)() {
    if m != nil {
        m.uri = value
    }
}
