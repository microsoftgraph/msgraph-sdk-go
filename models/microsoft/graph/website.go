package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new website and sets the default values.
func NewWebsite()(*Website) {
    m := &Website{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Website) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the address property value. The URL of the website.
func (m *Website) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the displayName property value. The display name of the web site.
func (m *Website) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the type_escaped property value. The possible values are: other, home, work, blog, profile.
func (m *Website) GetType_escaped()(*WebsiteType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *Website) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAddress(val)
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseWebsiteType)
        if err != nil {
            return err
        }
        cast := val.(WebsiteType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *Website) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
func (m *Website) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the address property value. The URL of the website.
// Parameters:
//  - value : Value to set for the address property.
func (m *Website) SetAddress(value *string)() {
    m.address = value
}
// Sets the displayName property value. The display name of the web site.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Website) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the type_escaped property value. The possible values are: other, home, work, blog, profile.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *Website) SetType_escaped(value *WebsiteType)() {
    m.type_escaped = value
}
