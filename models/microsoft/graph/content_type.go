package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ContentType struct {
    Entity
    // List of canonical URLs for hub sites with which this content type is associated to. This will contain all hub sites where this content type is queued to be enforced or is already enforced. Enforcing a content type means that the content type will be applied to the lists in the enforced sites.
    associatedHubsUrls []string;
    // Parent contentType from which this content type is derived.
    base *ContentType;
    // The collection of content types that are ancestors of this content type.
    baseTypes []ContentType;
    // The collection of columns that are required by this content type.
    columnLinks []ColumnLink;
    // Column order information in a content type.
    columnPositions []ColumnDefinition;
    // The collection of column definitions for this contentType.
    columns []ColumnDefinition;
    // The descriptive text for the item.
    description *string;
    // Document Set metadata.
    documentSet *DocumentSet;
    // Document template metadata. To make sure that documents have consistent content across a site and its subsites, you can associate a Word, Excel, or PowerPoint template with a site content type.
    documentTemplate *DocumentSetContent;
    // The name of the group this content type belongs to. Helps organize related content types.
    group *string;
    // Indicates whether the content type is hidden in the list's 'New' menu.
    hidden *bool;
    // If this content type is inherited from another scope (like a site), provides a reference to the item where the content type is defined.
    inheritedFrom *ItemReference;
    // Specifies if a content type is a built-in content type.
    isBuiltIn *bool;
    // The name of the content type.
    name *string;
    // Specifies the order in which the content type appears in the selection UI.
    order *ContentTypeOrder;
    // The unique identifier of the content type.
    parentId *string;
    // If true, any changes made to the content type will be pushed to inherited content types and lists that implement the content type.
    propagateChanges *bool;
    // If true, the content type can't be modified unless this value is first set to false.
    readOnly *bool;
    // If true, the content type can't be modified by users or through push-down operations. Only site collection administrators can seal or unseal content types.
    sealed *bool;
}
// Instantiates a new contentType and sets the default values.
func NewContentType()(*ContentType) {
    m := &ContentType{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the associatedHubsUrls property value. List of canonical URLs for hub sites with which this content type is associated to. This will contain all hub sites where this content type is queued to be enforced or is already enforced. Enforcing a content type means that the content type will be applied to the lists in the enforced sites.
func (m *ContentType) GetAssociatedHubsUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.associatedHubsUrls
    }
}
// Gets the base property value. Parent contentType from which this content type is derived.
func (m *ContentType) GetBase()(*ContentType) {
    if m == nil {
        return nil
    } else {
        return m.base
    }
}
// Gets the baseTypes property value. The collection of content types that are ancestors of this content type.
func (m *ContentType) GetBaseTypes()([]ContentType) {
    if m == nil {
        return nil
    } else {
        return m.baseTypes
    }
}
// Gets the columnLinks property value. The collection of columns that are required by this content type.
func (m *ContentType) GetColumnLinks()([]ColumnLink) {
    if m == nil {
        return nil
    } else {
        return m.columnLinks
    }
}
// Gets the columnPositions property value. Column order information in a content type.
func (m *ContentType) GetColumnPositions()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.columnPositions
    }
}
// Gets the columns property value. The collection of column definitions for this contentType.
func (m *ContentType) GetColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
// Gets the description property value. The descriptive text for the item.
func (m *ContentType) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the documentSet property value. Document Set metadata.
func (m *ContentType) GetDocumentSet()(*DocumentSet) {
    if m == nil {
        return nil
    } else {
        return m.documentSet
    }
}
// Gets the documentTemplate property value. Document template metadata. To make sure that documents have consistent content across a site and its subsites, you can associate a Word, Excel, or PowerPoint template with a site content type.
func (m *ContentType) GetDocumentTemplate()(*DocumentSetContent) {
    if m == nil {
        return nil
    } else {
        return m.documentTemplate
    }
}
// Gets the group property value. The name of the group this content type belongs to. Helps organize related content types.
func (m *ContentType) GetGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// Gets the hidden property value. Indicates whether the content type is hidden in the list's 'New' menu.
func (m *ContentType) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// Gets the inheritedFrom property value. If this content type is inherited from another scope (like a site), provides a reference to the item where the content type is defined.
func (m *ContentType) GetInheritedFrom()(*ItemReference) {
    if m == nil {
        return nil
    } else {
        return m.inheritedFrom
    }
}
// Gets the isBuiltIn property value. Specifies if a content type is a built-in content type.
func (m *ContentType) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// Gets the name property value. The name of the content type.
func (m *ContentType) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the order property value. Specifies the order in which the content type appears in the selection UI.
func (m *ContentType) GetOrder()(*ContentTypeOrder) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
// Gets the parentId property value. The unique identifier of the content type.
func (m *ContentType) GetParentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentId
    }
}
// Gets the propagateChanges property value. If true, any changes made to the content type will be pushed to inherited content types and lists that implement the content type.
func (m *ContentType) GetPropagateChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateChanges
    }
}
// Gets the readOnly property value. If true, the content type can't be modified unless this value is first set to false.
func (m *ContentType) GetReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.readOnly
    }
}
// Gets the sealed property value. If true, the content type can't be modified by users or through push-down operations. Only site collection administrators can seal or unseal content types.
func (m *ContentType) GetSealed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sealed
    }
}
// The deserialization information for the current model
func (m *ContentType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["associatedHubsUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssociatedHubsUrls(res)
        return nil
    }
    res["base"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentType() })
        if err != nil {
            return err
        }
        m.SetBase(val.(*ContentType))
        return nil
    }
    res["baseTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentType() })
        if err != nil {
            return err
        }
        res := make([]ContentType, len(val))
        for i, v := range val {
            res[i] = *(v.(*ContentType))
        }
        m.SetBaseTypes(res)
        return nil
    }
    res["columnLinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnLink() })
        if err != nil {
            return err
        }
        res := make([]ColumnLink, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnLink))
        }
        m.SetColumnLinks(res)
        return nil
    }
    res["columnPositions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetColumnPositions(res)
        return nil
    }
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        res := make([]ColumnDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ColumnDefinition))
        }
        m.SetColumns(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["documentSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSet() })
        if err != nil {
            return err
        }
        m.SetDocumentSet(val.(*DocumentSet))
        return nil
    }
    res["documentTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSetContent() })
        if err != nil {
            return err
        }
        m.SetDocumentTemplate(val.(*DocumentSetContent))
        return nil
    }
    res["group"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroup(val)
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
    res["inheritedFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemReference() })
        if err != nil {
            return err
        }
        m.SetInheritedFrom(val.(*ItemReference))
        return nil
    }
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsBuiltIn(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["order"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeOrder() })
        if err != nil {
            return err
        }
        m.SetOrder(val.(*ContentTypeOrder))
        return nil
    }
    res["parentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentId(val)
        return nil
    }
    res["propagateChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPropagateChanges(val)
        return nil
    }
    res["readOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetReadOnly(val)
        return nil
    }
    res["sealed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSealed(val)
        return nil
    }
    return res
}
func (m *ContentType) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ContentType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("associatedHubsUrls", m.GetAssociatedHubsUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("base", m.GetBase())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBaseTypes()))
        for i, v := range m.GetBaseTypes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("baseTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumnLinks()))
        for i, v := range m.GetColumnLinks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columnLinks", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumnPositions()))
        for i, v := range m.GetColumnPositions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columnPositions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("documentSet", m.GetDocumentSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("documentTemplate", m.GetDocumentTemplate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hidden", m.GetHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inheritedFrom", m.GetInheritedFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("order", m.GetOrder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentId", m.GetParentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("propagateChanges", m.GetPropagateChanges())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("readOnly", m.GetReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sealed", m.GetSealed())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the associatedHubsUrls property value. List of canonical URLs for hub sites with which this content type is associated to. This will contain all hub sites where this content type is queued to be enforced or is already enforced. Enforcing a content type means that the content type will be applied to the lists in the enforced sites.
// Parameters:
//  - value : Value to set for the associatedHubsUrls property.
func (m *ContentType) SetAssociatedHubsUrls(value []string)() {
    m.associatedHubsUrls = value
}
// Sets the base property value. Parent contentType from which this content type is derived.
// Parameters:
//  - value : Value to set for the base property.
func (m *ContentType) SetBase(value *ContentType)() {
    m.base = value
}
// Sets the baseTypes property value. The collection of content types that are ancestors of this content type.
// Parameters:
//  - value : Value to set for the baseTypes property.
func (m *ContentType) SetBaseTypes(value []ContentType)() {
    m.baseTypes = value
}
// Sets the columnLinks property value. The collection of columns that are required by this content type.
// Parameters:
//  - value : Value to set for the columnLinks property.
func (m *ContentType) SetColumnLinks(value []ColumnLink)() {
    m.columnLinks = value
}
// Sets the columnPositions property value. Column order information in a content type.
// Parameters:
//  - value : Value to set for the columnPositions property.
func (m *ContentType) SetColumnPositions(value []ColumnDefinition)() {
    m.columnPositions = value
}
// Sets the columns property value. The collection of column definitions for this contentType.
// Parameters:
//  - value : Value to set for the columns property.
func (m *ContentType) SetColumns(value []ColumnDefinition)() {
    m.columns = value
}
// Sets the description property value. The descriptive text for the item.
// Parameters:
//  - value : Value to set for the description property.
func (m *ContentType) SetDescription(value *string)() {
    m.description = value
}
// Sets the documentSet property value. Document Set metadata.
// Parameters:
//  - value : Value to set for the documentSet property.
func (m *ContentType) SetDocumentSet(value *DocumentSet)() {
    m.documentSet = value
}
// Sets the documentTemplate property value. Document template metadata. To make sure that documents have consistent content across a site and its subsites, you can associate a Word, Excel, or PowerPoint template with a site content type.
// Parameters:
//  - value : Value to set for the documentTemplate property.
func (m *ContentType) SetDocumentTemplate(value *DocumentSetContent)() {
    m.documentTemplate = value
}
// Sets the group property value. The name of the group this content type belongs to. Helps organize related content types.
// Parameters:
//  - value : Value to set for the group property.
func (m *ContentType) SetGroup(value *string)() {
    m.group = value
}
// Sets the hidden property value. Indicates whether the content type is hidden in the list's 'New' menu.
// Parameters:
//  - value : Value to set for the hidden property.
func (m *ContentType) SetHidden(value *bool)() {
    m.hidden = value
}
// Sets the inheritedFrom property value. If this content type is inherited from another scope (like a site), provides a reference to the item where the content type is defined.
// Parameters:
//  - value : Value to set for the inheritedFrom property.
func (m *ContentType) SetInheritedFrom(value *ItemReference)() {
    m.inheritedFrom = value
}
// Sets the isBuiltIn property value. Specifies if a content type is a built-in content type.
// Parameters:
//  - value : Value to set for the isBuiltIn property.
func (m *ContentType) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// Sets the name property value. The name of the content type.
// Parameters:
//  - value : Value to set for the name property.
func (m *ContentType) SetName(value *string)() {
    m.name = value
}
// Sets the order property value. Specifies the order in which the content type appears in the selection UI.
// Parameters:
//  - value : Value to set for the order property.
func (m *ContentType) SetOrder(value *ContentTypeOrder)() {
    m.order = value
}
// Sets the parentId property value. The unique identifier of the content type.
// Parameters:
//  - value : Value to set for the parentId property.
func (m *ContentType) SetParentId(value *string)() {
    m.parentId = value
}
// Sets the propagateChanges property value. If true, any changes made to the content type will be pushed to inherited content types and lists that implement the content type.
// Parameters:
//  - value : Value to set for the propagateChanges property.
func (m *ContentType) SetPropagateChanges(value *bool)() {
    m.propagateChanges = value
}
// Sets the readOnly property value. If true, the content type can't be modified unless this value is first set to false.
// Parameters:
//  - value : Value to set for the readOnly property.
func (m *ContentType) SetReadOnly(value *bool)() {
    m.readOnly = value
}
// Sets the sealed property value. If true, the content type can't be modified by users or through push-down operations. Only site collection administrators can seal or unseal content types.
// Parameters:
//  - value : Value to set for the sealed property.
func (m *ContentType) SetSealed(value *bool)() {
    m.sealed = value
}
