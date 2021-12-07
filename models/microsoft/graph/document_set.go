package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DocumentSet 
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
// NewDocumentSet instantiates a new documentSet and sets the default values.
func NewDocumentSet()(*DocumentSet) {
    m := &DocumentSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DocumentSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedContentTypes gets the allowedContentTypes property value. Content types allowed in document set.
func (m *DocumentSet) GetAllowedContentTypes()([]ContentTypeInfo) {
    if m == nil {
        return nil
    } else {
        return m.allowedContentTypes
    }
}
// GetDefaultContents gets the defaultContents property value. Default contents of document set.
func (m *DocumentSet) GetDefaultContents()([]DocumentSetContent) {
    if m == nil {
        return nil
    } else {
        return m.defaultContents
    }
}
// GetPropagateWelcomePageChanges gets the propagateWelcomePageChanges property value. Specifies whether to push welcome page changes to inherited content types.
func (m *DocumentSet) GetPropagateWelcomePageChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateWelcomePageChanges
    }
}
// GetSharedColumns gets the sharedColumns property value. 
func (m *DocumentSet) GetSharedColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.sharedColumns
    }
}
// GetShouldPrefixNameToFile gets the shouldPrefixNameToFile property value. Add the name of the document set to each file name.
func (m *DocumentSet) GetShouldPrefixNameToFile()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shouldPrefixNameToFile
    }
}
// GetWelcomePageColumns gets the welcomePageColumns property value. 
func (m *DocumentSet) GetWelcomePageColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.welcomePageColumns
    }
}
// GetWelcomePageUrl gets the welcomePageUrl property value. Welcome page absolute URL.
func (m *DocumentSet) GetWelcomePageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.welcomePageUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedContentTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContentTypeInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*ContentTypeInfo))
            }
            m.SetAllowedContentTypes(res)
        }
        return nil
    }
    res["defaultContents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSetContent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DocumentSetContent, len(val))
            for i, v := range val {
                res[i] = *(v.(*DocumentSetContent))
            }
            m.SetDefaultContents(res)
        }
        return nil
    }
    res["propagateWelcomePageChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropagateWelcomePageChanges(val)
        }
        return nil
    }
    res["sharedColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*ColumnDefinition))
            }
            m.SetSharedColumns(res)
        }
        return nil
    }
    res["shouldPrefixNameToFile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldPrefixNameToFile(val)
        }
        return nil
    }
    res["welcomePageColumns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*ColumnDefinition))
            }
            m.SetWelcomePageColumns(res)
        }
        return nil
    }
    res["welcomePageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWelcomePageUrl(val)
        }
        return nil
    }
    return res
}
func (m *DocumentSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DocumentSet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowedContentTypes sets the allowedContentTypes property value. Content types allowed in document set.
func (m *DocumentSet) SetAllowedContentTypes(value []ContentTypeInfo)() {
    if m != nil {
        m.allowedContentTypes = value
    }
}
// SetDefaultContents sets the defaultContents property value. Default contents of document set.
func (m *DocumentSet) SetDefaultContents(value []DocumentSetContent)() {
    if m != nil {
        m.defaultContents = value
    }
}
// SetPropagateWelcomePageChanges sets the propagateWelcomePageChanges property value. Specifies whether to push welcome page changes to inherited content types.
func (m *DocumentSet) SetPropagateWelcomePageChanges(value *bool)() {
    if m != nil {
        m.propagateWelcomePageChanges = value
    }
}
// SetSharedColumns sets the sharedColumns property value. 
func (m *DocumentSet) SetSharedColumns(value []ColumnDefinition)() {
    if m != nil {
        m.sharedColumns = value
    }
}
// SetShouldPrefixNameToFile sets the shouldPrefixNameToFile property value. Add the name of the document set to each file name.
func (m *DocumentSet) SetShouldPrefixNameToFile(value *bool)() {
    if m != nil {
        m.shouldPrefixNameToFile = value
    }
}
// SetWelcomePageColumns sets the welcomePageColumns property value. 
func (m *DocumentSet) SetWelcomePageColumns(value []ColumnDefinition)() {
    if m != nil {
        m.welcomePageColumns = value
    }
}
// SetWelcomePageUrl sets the welcomePageUrl property value. Welcome page absolute URL.
func (m *DocumentSet) SetWelcomePageUrl(value *string)() {
    if m != nil {
        m.welcomePageUrl = value
    }
}
