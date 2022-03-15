package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SiteCollection provides operations to manage the drive singleton.
type SiteCollection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The geographic region code for where this site collection resides. Read-only.
    dataLocationCode *string;
    // The hostname for the site collection. Read-only.
    hostname *string;
    // If present, indicates that this is a root site collection in SharePoint. Read-only.
    root Rootable;
}
// NewSiteCollection instantiates a new siteCollection and sets the default values.
func NewSiteCollection()(*SiteCollection) {
    m := &SiteCollection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSiteCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSiteCollectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSiteCollection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SiteCollection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDataLocationCode gets the dataLocationCode property value. The geographic region code for where this site collection resides. Read-only.
func (m *SiteCollection) GetDataLocationCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataLocationCode
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetObjectValue(CreateRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(Rootable))
        }
        return nil
    }
    return res
}
// GetHostname gets the hostname property value. The hostname for the site collection. Read-only.
func (m *SiteCollection) GetHostname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hostname
    }
}
// GetRoot gets the root property value. If present, indicates that this is a root site collection in SharePoint. Read-only.
func (m *SiteCollection) GetRoot()(Rootable) {
    if m == nil {
        return nil
    } else {
        return m.root
    }
}
func (m *SiteCollection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SiteCollection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDataLocationCode sets the dataLocationCode property value. The geographic region code for where this site collection resides. Read-only.
func (m *SiteCollection) SetDataLocationCode(value *string)() {
    if m != nil {
        m.dataLocationCode = value
    }
}
// SetHostname sets the hostname property value. The hostname for the site collection. Read-only.
func (m *SiteCollection) SetHostname(value *string)() {
    if m != nil {
        m.hostname = value
    }
}
// SetRoot sets the root property value. If present, indicates that this is a root site collection in SharePoint. Read-only.
func (m *SiteCollection) SetRoot(value Rootable)() {
    if m != nil {
        m.root = value
    }
}
