package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceAnnouncement 
type ServiceAnnouncement struct {
    Entity
    // A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
    healthOverviews []ServiceHealth;
    // A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
    issues []ServiceHealthIssue;
    // A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
    messages []ServiceUpdateMessage;
}
// NewServiceAnnouncement instantiates a new serviceAnnouncement and sets the default values.
func NewServiceAnnouncement()(*ServiceAnnouncement) {
    m := &ServiceAnnouncement{
        Entity: *NewEntity(),
    }
    return m
}
// GetHealthOverviews gets the healthOverviews property value. A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetHealthOverviews()([]ServiceHealth) {
    if m == nil {
        return nil
    } else {
        return m.healthOverviews
    }
}
// GetIssues gets the issues property value. A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetIssues()([]ServiceHealthIssue) {
    if m == nil {
        return nil
    } else {
        return m.issues
    }
}
// GetMessages gets the messages property value. A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetMessages()([]ServiceUpdateMessage) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAnnouncement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthOverviews"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealth, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServiceHealth))
            }
            m.SetHealthOverviews(res)
        }
        return nil
    }
    res["issues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealthIssue() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealthIssue, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServiceHealthIssue))
            }
            m.SetIssues(res)
        }
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceUpdateMessage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceUpdateMessage, len(val))
            for i, v := range val {
                res[i] = *(v.(*ServiceUpdateMessage))
            }
            m.SetMessages(res)
        }
        return nil
    }
    return res
}
func (m *ServiceAnnouncement) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetHealthOverviews sets the healthOverviews property value. A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetHealthOverviews(value []ServiceHealth)() {
    if m != nil {
        m.healthOverviews = value
    }
}
// SetIssues sets the issues property value. A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetIssues(value []ServiceHealthIssue)() {
    if m != nil {
        m.issues = value
    }
}
// SetMessages sets the messages property value. A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetMessages(value []ServiceUpdateMessage)() {
    if m != nil {
        m.messages = value
    }
}
