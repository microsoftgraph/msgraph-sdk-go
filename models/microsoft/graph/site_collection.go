package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SiteCollection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The geographic region code for where this site collection resides. Read-only.
    dataLocationCode *string;
    // The hostname for the site collection. Read-only.
    hostname *string;
    // If present, indicates that this is a root site collection in SharePoint. Read-only.
    root *Root;
}
// Instantiates a new siteCollection and sets the default values.
func NewSiteCollection()(*SiteCollection) {
    m := &SiteCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SiteCollection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the dataLocationCode property value. The geographic region code for where this site collection resides. Read-only.
func (m *SiteCollection) GetDataLocationCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataLocationCode
    }
}
// Gets the hostname property value. The hostname for the site collection. Read-only.
func (m *SiteCollection) GetHostname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hostname
    }
}
// Gets the root property value. If present, indicates that this is a root site collection in SharePoint. Read-only.
func (m *SiteCollection) GetRoot()(*Root) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
// The deserialization information for the current model
func (m *SiteCollection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataLocationCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataLocationCode(val)
        }
        return nil
    }
    res["hostname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val)
        }
        return nil
    }
    res["root"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoot() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(*Root))
        }
        return nil
    }
    return res
}
func (m *SiteCollection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SiteCollection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("dataLocationCode", m.GetDataLocationCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hostname", m.GetHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("root", m.GetRoot())
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
func (m *SiteCollection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the dataLocationCode property value. The geographic region code for where this site collection resides. Read-only.
// Parameters:
//  - value : Value to set for the dataLocationCode property.
func (m *SiteCollection) SetDataLocationCode(value *string)() {
    m.dataLocationCode = value
}
// Sets the hostname property value. The hostname for the site collection. Read-only.
// Parameters:
//  - value : Value to set for the hostname property.
func (m *SiteCollection) SetHostname(value *string)() {
    m.hostname = value
}
// Sets the root property value. If present, indicates that this is a root site collection in SharePoint. Read-only.
// Parameters:
//  - value : Value to set for the root property.
func (m *SiteCollection) SetRoot(value *Root)() {
    m.root = value
}
