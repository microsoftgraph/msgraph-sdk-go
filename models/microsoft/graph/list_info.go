package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new listInfo and sets the default values.
func NewListInfo()(*ListInfo) {
    m := &ListInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ListInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the contentTypesEnabled property value. If true, indicates that content types are enabled for this list.
func (m *ListInfo) GetContentTypesEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contentTypesEnabled
    }
}
// Gets the hidden property value. If true, indicates that the list is not normally visible in the SharePoint user experience.
func (m *ListInfo) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// Gets the template property value. An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
func (m *ListInfo) GetTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.template
    }
}
// The deserialization information for the current model
func (m *ListInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentTypesEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetContentTypesEnabled(val)
        return nil
    }
    res["hidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHidden(val)
        return nil
    }
    res["template"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTemplate(val)
        return nil
    }
    return res
}
func (m *ListInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ListInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the contentTypesEnabled property value. If true, indicates that content types are enabled for this list.
// Parameters:
//  - value : Value to set for the contentTypesEnabled property.
func (m *ListInfo) SetContentTypesEnabled(value *bool)() {
    m.contentTypesEnabled = value
}
// Sets the hidden property value. If true, indicates that the list is not normally visible in the SharePoint user experience.
// Parameters:
//  - value : Value to set for the hidden property.
func (m *ListInfo) SetHidden(value *bool)() {
    m.hidden = value
}
// Sets the template property value. An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more.
// Parameters:
//  - value : Value to set for the template property.
func (m *ListInfo) SetTemplate(value *string)() {
    m.template = value
}
