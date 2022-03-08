package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LinkedResource provides operations to manage the drive singleton.
type LinkedResource struct {
    Entity
    // Field indicating the app name of the source that is sending the linkedResource.
    applicationName *string;
    // Field indicating the title of the linkedResource.
    displayName *string;
    // Id of the object that is associated with this task on the third-party/partner system.
    externalId *string;
    // Deep link to the linkedResource.
    webUrl *string;
}
// NewLinkedResource instantiates a new linkedResource and sets the default values.
func NewLinkedResource()(*LinkedResource) {
    m := &LinkedResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLinkedResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLinkedResourceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewLinkedResource(), nil
}
// GetApplicationName gets the applicationName property value. Field indicating the app name of the source that is sending the linkedResource.
func (m *LinkedResource) GetApplicationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationName
    }
}
// GetDisplayName gets the displayName property value. Field indicating the title of the linkedResource.
func (m *LinkedResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExternalId gets the externalId property value. Id of the object that is associated with this task on the third-party/partner system.
func (m *LinkedResource) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LinkedResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationName(val)
        }
        return nil
    }
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
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetWebUrl gets the webUrl property value. Deep link to the linkedResource.
func (m *LinkedResource) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *LinkedResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *LinkedResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationName", m.GetApplicationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationName sets the applicationName property value. Field indicating the app name of the source that is sending the linkedResource.
func (m *LinkedResource) SetApplicationName(value *string)() {
    if m != nil {
        m.applicationName = value
    }
}
// SetDisplayName sets the displayName property value. Field indicating the title of the linkedResource.
func (m *LinkedResource) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExternalId sets the externalId property value. Id of the object that is associated with this task on the third-party/partner system.
func (m *LinkedResource) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetWebUrl sets the webUrl property value. Deep link to the linkedResource.
func (m *LinkedResource) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
