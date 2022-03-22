package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ListInfo 
type ListInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If true, indicates that content types are enabled for this list.
    contentTypesEnabled *bool;
    // If true, indicates that the list is not normally visible in the SharePoint user experience.
    hidden *bool;
    // An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
    template *string;
}
// NewListInfo instantiates a new listInfo and sets the default values.
func NewListInfo()(*ListInfo) {
    m := &ListInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateListInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewListInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ListInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentTypesEnabled gets the contentTypesEnabled property value. If true, indicates that content types are enabled for this list.
func (m *ListInfo) GetContentTypesEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contentTypesEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentTypesEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentTypesEnabled(val)
        }
        return nil
    }
    res["hidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidden(val)
        }
        return nil
    }
    res["template"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplate(val)
        }
        return nil
    }
    return res
}
// GetHidden gets the hidden property value. If true, indicates that the list is not normally visible in the SharePoint user experience.
func (m *ListInfo) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// GetTemplate gets the template property value. An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
func (m *ListInfo) GetTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.template
    }
}
// Serialize serializes information the current object
func (m *ListInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("contentTypesEnabled", m.GetContentTypesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hidden", m.GetHidden())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("template", m.GetTemplate())
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
func (m *ListInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentTypesEnabled sets the contentTypesEnabled property value. If true, indicates that content types are enabled for this list.
func (m *ListInfo) SetContentTypesEnabled(value *bool)() {
    if m != nil {
        m.contentTypesEnabled = value
    }
}
// SetHidden sets the hidden property value. If true, indicates that the list is not normally visible in the SharePoint user experience.
func (m *ListInfo) SetHidden(value *bool)() {
    if m != nil {
        m.hidden = value
    }
}
// SetTemplate sets the template property value. An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
func (m *ListInfo) SetTemplate(value *string)() {
    if m != nil {
        m.template = value
    }
}
