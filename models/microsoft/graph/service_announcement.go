package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ServiceAnnouncement provides operations to manage the admin singleton.
type ServiceAnnouncement struct {
    Entity
    // A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
    healthOverviews []ServiceHealthable;
    // A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
    issues []ServiceHealthIssueable;
    // A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
    messages []ServiceUpdateMessageable;
}
// NewServiceAnnouncement instantiates a new serviceAnnouncement and sets the default values.
func NewServiceAnnouncement()(*ServiceAnnouncement) {
    m := &ServiceAnnouncement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceAnnouncementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewServiceAnnouncement(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAnnouncement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthOverviews"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealthable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceHealthable)
            }
            m.SetHealthOverviews(res)
        }
        return nil
    }
    res["issues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceHealthIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealthIssueable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceHealthIssueable)
            }
            m.SetIssues(res)
        }
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceUpdateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceUpdateMessageable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceUpdateMessageable)
            }
            m.SetMessages(res)
        }
        return nil
    }
    return res
}
// GetHealthOverviews gets the healthOverviews property value. A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetHealthOverviews()([]ServiceHealthable) {
    if m == nil {
        return nil
    } else {
        return m.healthOverviews
    }
}
// GetIssues gets the issues property value. A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetIssues()([]ServiceHealthIssueable) {
    if m == nil {
        return nil
    } else {
        return m.issues
    }
}
// GetMessages gets the messages property value. A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetMessages()([]ServiceUpdateMessageable) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
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
    if m.GetHealthOverviews() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHealthOverviews()))
        for i, v := range m.GetHealthOverviews() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("healthOverviews", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIssues() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIssues()))
        for i, v := range m.GetIssues() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("issues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMessages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHealthOverviews sets the healthOverviews property value. A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetHealthOverviews(value []ServiceHealthable)() {
    if m != nil {
        m.healthOverviews = value
    }
}
// SetIssues sets the issues property value. A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetIssues(value []ServiceHealthIssueable)() {
    if m != nil {
        m.issues = value
    }
}
// SetMessages sets the messages property value. A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetMessages(value []ServiceUpdateMessageable)() {
    if m != nil {
        m.messages = value
    }
}
