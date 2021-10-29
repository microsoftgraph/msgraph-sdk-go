package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SharingLink struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The app the link is associated with.
    application *Identity;
    // If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
    preventsDownload *bool;
    // The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
    scope *string;
    // The type of the link created.
    type_escaped *string;
    // For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
    webHtml *string;
    // A URL that opens the item in the browser on the OneDrive website.
    webUrl *string;
}
// Instantiates a new sharingLink and sets the default values.
func NewSharingLink()(*SharingLink) {
    m := &SharingLink{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharingLink) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the application property value. The app the link is associated with.
func (m *SharingLink) GetApplication()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// Gets the preventsDownload property value. If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
func (m *SharingLink) GetPreventsDownload()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.preventsDownload
    }
}
// Gets the scope property value. The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
func (m *SharingLink) GetScope()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scope
    }
}
// Gets the type_escaped property value. The type of the link created.
func (m *SharingLink) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the webHtml property value. For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
func (m *SharingLink) GetWebHtml()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webHtml
    }
}
// Gets the webUrl property value. A URL that opens the item in the browser on the OneDrive website.
func (m *SharingLink) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *SharingLink) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        m.SetApplication(val.(*Identity))
        return nil
    }
    res["preventsDownload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPreventsDownload(val)
        return nil
    }
    res["scope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScope(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    res["webHtml"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebHtml(val)
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
func (m *SharingLink) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SharingLink) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("preventsDownload", m.GetPreventsDownload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webHtml", m.GetWebHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *SharingLink) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the application property value. The app the link is associated with.
// Parameters:
//  - value : Value to set for the application property.
func (m *SharingLink) SetApplication(value *Identity)() {
    m.application = value
}
// Sets the preventsDownload property value. If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint.
// Parameters:
//  - value : Value to set for the preventsDownload property.
func (m *SharingLink) SetPreventsDownload(value *bool)() {
    m.preventsDownload = value
}
// Sets the scope property value. The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant.
// Parameters:
//  - value : Value to set for the scope property.
func (m *SharingLink) SetScope(value *string)() {
    m.scope = value
}
// Sets the type_escaped property value. The type of the link created.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *SharingLink) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the webHtml property value. For embed links, this property contains the HTML code for an <iframe> element that will embed the item in a webpage.
// Parameters:
//  - value : Value to set for the webHtml property.
func (m *SharingLink) SetWebHtml(value *string)() {
    m.webHtml = value
}
// Sets the webUrl property value. A URL that opens the item in the browser on the OneDrive website.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *SharingLink) SetWebUrl(value *string)() {
    m.webUrl = value
}
