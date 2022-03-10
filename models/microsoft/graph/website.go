package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Website provides operations to manage the collection of drive entities.
type Website struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The URL of the website.
    address *string;
    // The display name of the web site.
    displayName *string;
    // The possible values are: other, home, work, blog, profile.
    type_escaped *WebsiteType;
}
// NewWebsite instantiates a new website and sets the default values.
func NewWebsite()(*Website) {
    m := &Website{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWebsiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebsiteFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWebsite(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Website) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddress gets the address property value. The URL of the website.
func (m *Website) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetDisplayName gets the displayName property value. The display name of the web site.
func (m *Website) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Website) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
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
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWebsiteType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*WebsiteType))
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The possible values are: other, home, work, blog, profile.
func (m *Website) GetType()(*WebsiteType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *Website) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Website) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Website) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddress sets the address property value. The URL of the website.
func (m *Website) SetAddress(value *string)() {
    if m != nil {
        m.address = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the web site.
func (m *Website) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetType sets the type property value. The possible values are: other, home, work, blog, profile.
func (m *Website) SetType(value *WebsiteType)() {
    if m != nil {
        m.type_escaped = value
    }
}
