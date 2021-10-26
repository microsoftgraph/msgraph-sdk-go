package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ResourceVisualization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
    containerDisplayName *string;
    // Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
    containerType *string;
    // A path leading to the folder in which the item is stored.
    containerWebUrl *string;
    // The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Note that not all Media Mime Types are supported.
    mediaType *string;
    // A URL leading to the preview image for the item.
    previewImageUrl *string;
    // A preview text for the item.
    previewText *string;
    // The item's title text.
    title *string;
    // The item's media type. Can be used for filtering for a specific file based on a specific type. See below for supported types.
    type_escaped *string;
}
// Instantiates a new resourceVisualization and sets the default values.
func NewResourceVisualization()(*ResourceVisualization) {
    m := &ResourceVisualization{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceVisualization) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the containerDisplayName property value. A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
func (m *ResourceVisualization) GetContainerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerDisplayName
    }
}
// Gets the containerType property value. Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
func (m *ResourceVisualization) GetContainerType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerType
    }
}
// Gets the containerWebUrl property value. A path leading to the folder in which the item is stored.
func (m *ResourceVisualization) GetContainerWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerWebUrl
    }
}
// Gets the mediaType property value. The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Note that not all Media Mime Types are supported.
func (m *ResourceVisualization) GetMediaType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// Gets the previewImageUrl property value. A URL leading to the preview image for the item.
func (m *ResourceVisualization) GetPreviewImageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.previewImageUrl
    }
}
// Gets the previewText property value. A preview text for the item.
func (m *ResourceVisualization) GetPreviewText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.previewText
    }
}
// Gets the title property value. The item's title text.
func (m *ResourceVisualization) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Gets the type_escaped property value. The item's media type. Can be used for filtering for a specific file based on a specific type. See below for supported types.
func (m *ResourceVisualization) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *ResourceVisualization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["containerDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContainerDisplayName(val)
        return nil
    }
    res["containerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContainerType(val)
        return nil
    }
    res["containerWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContainerWebUrl(val)
        return nil
    }
    res["mediaType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMediaType(val)
        return nil
    }
    res["previewImageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreviewImageUrl(val)
        return nil
    }
    res["previewText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreviewText(val)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
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
    return res
}
func (m *ResourceVisualization) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ResourceVisualization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("containerDisplayName", m.GetContainerDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("containerType", m.GetContainerType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("containerWebUrl", m.GetContainerWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewImageUrl", m.GetPreviewImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *ResourceVisualization) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the containerDisplayName property value. A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
// Parameters:
//  - value : Value to set for the containerDisplayName property.
func (m *ResourceVisualization) SetContainerDisplayName(value *string)() {
    m.containerDisplayName = value
}
// Sets the containerType property value. Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
// Parameters:
//  - value : Value to set for the containerType property.
func (m *ResourceVisualization) SetContainerType(value *string)() {
    m.containerType = value
}
// Sets the containerWebUrl property value. A path leading to the folder in which the item is stored.
// Parameters:
//  - value : Value to set for the containerWebUrl property.
func (m *ResourceVisualization) SetContainerWebUrl(value *string)() {
    m.containerWebUrl = value
}
// Sets the mediaType property value. The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Note that not all Media Mime Types are supported.
// Parameters:
//  - value : Value to set for the mediaType property.
func (m *ResourceVisualization) SetMediaType(value *string)() {
    m.mediaType = value
}
// Sets the previewImageUrl property value. A URL leading to the preview image for the item.
// Parameters:
//  - value : Value to set for the previewImageUrl property.
func (m *ResourceVisualization) SetPreviewImageUrl(value *string)() {
    m.previewImageUrl = value
}
// Sets the previewText property value. A preview text for the item.
// Parameters:
//  - value : Value to set for the previewText property.
func (m *ResourceVisualization) SetPreviewText(value *string)() {
    m.previewText = value
}
// Sets the title property value. The item's title text.
// Parameters:
//  - value : Value to set for the title property.
func (m *ResourceVisualization) SetTitle(value *string)() {
    m.title = value
}
// Sets the type_escaped property value. The item's media type. Can be used for filtering for a specific file based on a specific type. See below for supported types.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *ResourceVisualization) SetType_escaped(value *string)() {
    m.type_escaped = value
}
