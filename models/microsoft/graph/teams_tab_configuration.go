package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamsTabConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Url used for rendering tab contents in Teams. Required.
    contentUrl *string;
    // Identifier for the entity hosted by the tab provider.
    entityId *string;
    // Url called by Teams client when a Tab is removed using the Teams Client.
    removeUrl *string;
    // Url for showing tab contents outside of Teams.
    websiteUrl *string;
}
// Instantiates a new teamsTabConfiguration and sets the default values.
func NewTeamsTabConfiguration()(*TeamsTabConfiguration) {
    m := &TeamsTabConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamsTabConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the contentUrl property value. Url used for rendering tab contents in Teams. Required.
func (m *TeamsTabConfiguration) GetContentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentUrl
    }
}
// Gets the entityId property value. Identifier for the entity hosted by the tab provider.
func (m *TeamsTabConfiguration) GetEntityId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.entityId
    }
}
// Gets the removeUrl property value. Url called by Teams client when a Tab is removed using the Teams Client.
func (m *TeamsTabConfiguration) GetRemoveUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.removeUrl
    }
}
// Gets the websiteUrl property value. Url for showing tab contents outside of Teams.
func (m *TeamsTabConfiguration) GetWebsiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.websiteUrl
    }
}
// The deserialization information for the current model
func (m *TeamsTabConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
        }
        return nil
    }
    res["entityId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityId(val)
        }
        return nil
    }
    res["removeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveUrl(val)
        }
        return nil
    }
    res["websiteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebsiteUrl(val)
        }
        return nil
    }
    return res
}
func (m *TeamsTabConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamsTabConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityId", m.GetEntityId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("removeUrl", m.GetRemoveUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("websiteUrl", m.GetWebsiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TeamsTabConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the contentUrl property value. Url used for rendering tab contents in Teams. Required.
// Parameters:
//  - value : Value to set for the contentUrl property.
func (m *TeamsTabConfiguration) SetContentUrl(value *string)() {
    m.contentUrl = value
}
// Sets the entityId property value. Identifier for the entity hosted by the tab provider.
// Parameters:
//  - value : Value to set for the entityId property.
func (m *TeamsTabConfiguration) SetEntityId(value *string)() {
    m.entityId = value
}
// Sets the removeUrl property value. Url called by Teams client when a Tab is removed using the Teams Client.
// Parameters:
//  - value : Value to set for the removeUrl property.
func (m *TeamsTabConfiguration) SetRemoveUrl(value *string)() {
    m.removeUrl = value
}
// Sets the websiteUrl property value. Url for showing tab contents outside of Teams.
// Parameters:
//  - value : Value to set for the websiteUrl property.
func (m *TeamsTabConfiguration) SetWebsiteUrl(value *string)() {
    m.websiteUrl = value
}
