package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new linkedResource and sets the default values.
func NewLinkedResource()(*LinkedResource) {
    m := &LinkedResource{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applicationName property value. Field indicating the app name of the source that is sending the linkedResource.
func (m *LinkedResource) GetApplicationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationName
    }
}
// Gets the displayName property value. Field indicating the title of the linkedResource.
func (m *LinkedResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the externalId property value. Id of the object that is associated with this task on the third-party/partner system.
func (m *LinkedResource) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the webUrl property value. Deep link to the linkedResource.
func (m *LinkedResource) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *LinkedResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationName(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *LinkedResource) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the applicationName property value. Field indicating the app name of the source that is sending the linkedResource.
// Parameters:
//  - value : Value to set for the applicationName property.
func (m *LinkedResource) SetApplicationName(value *string)() {
    m.applicationName = value
}
// Sets the displayName property value. Field indicating the title of the linkedResource.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *LinkedResource) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the externalId property value. Id of the object that is associated with this task on the third-party/partner system.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *LinkedResource) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the webUrl property value. Deep link to the linkedResource.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *LinkedResource) SetWebUrl(value *string)() {
    m.webUrl = value
}
