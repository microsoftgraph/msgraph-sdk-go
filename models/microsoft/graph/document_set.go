package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DocumentSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Content types allowed in document set.
    allowedContentTypes []ContentTypeInfo;
    // Default contents of document set.
    defaultContents []DocumentSetContent;
    // Specifies whether to push welcome page changes to inherited content types.
    propagateWelcomePageChanges *bool;
    // 
    sharedColumns []ColumnDefinition;
    // Add the name of the document set to each file name.
    shouldPrefixNameToFile *bool;
    // 
    welcomePageColumns []ColumnDefinition;
    // Welcome page absolute URL.
    welcomePageUrl *string;
}
// Instantiates a new documentSet and sets the default values.
func NewDocumentSet()(*DocumentSet) {
    m := &DocumentSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DocumentSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowedContentTypes property value. Content types allowed in document set.
func (m *DocumentSet) GetAllowedContentTypes()([]ContentTypeInfo) {
    if m == nil {
        return nil
    } else {
        return m.allowedContentTypes
    }
}
// Gets the defaultContents property value. Default contents of document set.
func (m *DocumentSet) GetDefaultContents()([]DocumentSetContent) {
    if m == nil {
        return nil
    } else {
        return m.defaultContents
    }
}
// Gets the propagateWelcomePageChanges property value. Specifies whether to push welcome page changes to inherited content types.
func (m *DocumentSet) GetPropagateWelcomePageChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateWelcomePageChanges
    }
}
// Gets the sharedColumns property value. 
func (m *DocumentSet) GetSharedColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.sharedColumns
    }
}
// Gets the shouldPrefixNameToFile property value. Add the name of the document set to each file name.
func (m *DocumentSet) GetShouldPrefixNameToFile()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shouldPrefixNameToFile
    }
}
// Gets the welcomePageColumns property value. 
func (m *DocumentSet) GetWelcomePageColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.welcomePageColumns
    }
}
// Gets the welcomePageUrl property value. Welcome page absolute URL.
func (m *DocumentSet) GetWelcomePageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.welcomePageUrl
    }
}
// The deserialization information for the current model
func (m *DocumentSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedContentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeInfo() })
        if err != nil {
            return err
        }
        res := make([]ContentTypeInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*ContentTypeInfo))
        }
        m.SetAllowedContentTypes(res)
        return nil
    }
    res["defaultContents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSetContent() })
        if err != nil {
            return err
        }
        res := make([]DocumentSetContent, len(val))
        for i, v := range val {
            res[i] = *(v.(*DocumentSetContent))
        }
        m.SetDefaultContents(res)
        return nil
    }
    res["propagateWelcomePageChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPropagateWelcomePageChanges(val)
        return nil
    }
    res["sharedColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetSharedColumns(res)
        return nil
    }
    res["shouldPrefixNameToFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShouldPrefixNameToFile(val)
        return nil
    }
    res["welcomePageColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetWelcomePageColumns(res)
        return nil
    }
    res["welcomePageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWelcomePageUrl(val)
        return nil
    }
    return res
}
func (m *DocumentSet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DocumentSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedContentTypes()))
        for i, v := range m.GetAllowedContentTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("allowedContentTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefaultContents()))
        for i, v := range m.GetDefaultContents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("defaultContents", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("propagateWelcomePageChanges", m.GetPropagateWelcomePageChanges())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSharedColumns()))
        for i, v := range m.GetSharedColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("sharedColumns", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("shouldPrefixNameToFile", m.GetShouldPrefixNameToFile())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWelcomePageColumns()))
        for i, v := range m.GetWelcomePageColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("welcomePageColumns", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("welcomePageUrl", m.GetWelcomePageUrl())
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
func (m *DocumentSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowedContentTypes property value. Content types allowed in document set.
// Parameters:
//  - value : Value to set for the allowedContentTypes property.
func (m *DocumentSet) SetAllowedContentTypes(value []ContentTypeInfo)() {
    m.allowedContentTypes = value
}
// Sets the defaultContents property value. Default contents of document set.
// Parameters:
//  - value : Value to set for the defaultContents property.
func (m *DocumentSet) SetDefaultContents(value []DocumentSetContent)() {
    m.defaultContents = value
}
// Sets the propagateWelcomePageChanges property value. Specifies whether to push welcome page changes to inherited content types.
// Parameters:
//  - value : Value to set for the propagateWelcomePageChanges property.
func (m *DocumentSet) SetPropagateWelcomePageChanges(value *bool)() {
    m.propagateWelcomePageChanges = value
}
// Sets the sharedColumns property value. 
// Parameters:
//  - value : Value to set for the sharedColumns property.
func (m *DocumentSet) SetSharedColumns(value []ColumnDefinition)() {
    m.sharedColumns = value
}
// Sets the shouldPrefixNameToFile property value. Add the name of the document set to each file name.
// Parameters:
//  - value : Value to set for the shouldPrefixNameToFile property.
func (m *DocumentSet) SetShouldPrefixNameToFile(value *bool)() {
    m.shouldPrefixNameToFile = value
}
// Sets the welcomePageColumns property value. 
// Parameters:
//  - value : Value to set for the welcomePageColumns property.
func (m *DocumentSet) SetWelcomePageColumns(value []ColumnDefinition)() {
    m.welcomePageColumns = value
}
// Sets the welcomePageUrl property value. Welcome page absolute URL.
// Parameters:
//  - value : Value to set for the welcomePageUrl property.
func (m *DocumentSet) SetWelcomePageUrl(value *string)() {
    m.welcomePageUrl = value
}
