package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ServiceAnnouncement struct {
    Entity
    // A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
    healthOverviews []ServiceHealth;
    // A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
    issues []ServiceHealthIssue;
    // A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
    messages []ServiceUpdateMessage;
}
// Instantiates a new serviceAnnouncement and sets the default values.
func NewServiceAnnouncement()(*ServiceAnnouncement) {
    m := &ServiceAnnouncement{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the healthOverviews property value. A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetHealthOverviews()([]ServiceHealth) {
    if m == nil {
        return nil
    } else {
        return m.healthOverviews
    }
}
// Gets the issues property value. A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetIssues()([]ServiceHealthIssue) {
    if m == nil {
        return nil
    } else {
        return m.issues
    }
}
// Gets the messages property value. A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetMessages()([]ServiceUpdateMessage) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// The deserialization information for the current model
func (m *ServiceAnnouncement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthOverviews"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealth() })
        if err != nil {
            return err
        }
        res := make([]ServiceHealth, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServiceHealth))
        }
        m.SetHealthOverviews(res)
        return nil
    }
    res["issues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealthIssue() })
        if err != nil {
            return err
        }
        res := make([]ServiceHealthIssue, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServiceHealthIssue))
        }
        m.SetIssues(res)
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceUpdateMessage() })
        if err != nil {
            return err
        }
        res := make([]ServiceUpdateMessage, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServiceUpdateMessage))
        }
        m.SetMessages(res)
        return nil
    }
    return res
}
func (m *ServiceAnnouncement) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ServiceAnnouncement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHealthOverviews()))
        for i, v := range m.GetHealthOverviews() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("healthOverviews", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIssues()))
        for i, v := range m.GetIssues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("issues", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the healthOverviews property value. A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
// Parameters:
//  - value : Value to set for the healthOverviews property.
func (m *ServiceAnnouncement) SetHealthOverviews(value []ServiceHealth)() {
    m.healthOverviews = value
}
// Sets the issues property value. A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
// Parameters:
//  - value : Value to set for the issues property.
func (m *ServiceAnnouncement) SetIssues(value []ServiceHealthIssue)() {
    m.issues = value
}
// Sets the messages property value. A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
// Parameters:
//  - value : Value to set for the messages property.
func (m *ServiceAnnouncement) SetMessages(value []ServiceUpdateMessage)() {
    m.messages = value
}
