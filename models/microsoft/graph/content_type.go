package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContentType 
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
// NewContentType instantiates a new contentType and sets the default values.
func NewContentType()(*ContentType) {
    m := &ContentType{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssociatedHubsUrls gets the associatedHubsUrls property value. List of canonical URLs for hub sites with which this content type is associated to. This will contain all hub sites where this content type is queued to be enforced or is already enforced. Enforcing a content type means that the content type will be applied to the lists in the enforced sites.
func (m *ContentType) GetAssociatedHubsUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.associatedHubsUrls
    }
}
// GetBase gets the base property value. Parent contentType from which this content type is derived.
func (m *ContentType) GetBase()(*ContentType) {
    if m == nil {
        return nil
    } else {
        return m.base
    }
}
// GetBaseTypes gets the baseTypes property value. The collection of content types that are ancestors of this content type.
func (m *ContentType) GetBaseTypes()([]ContentType) {
    if m == nil {
        return nil
    } else {
        return m.baseTypes
    }
}
// GetColumnLinks gets the columnLinks property value. The collection of columns that are required by this content type.
func (m *ContentType) GetColumnLinks()([]ColumnLink) {
    if m == nil {
        return nil
    } else {
        return m.columnLinks
    }
}
// GetColumnPositions gets the columnPositions property value. Column order information in a content type.
func (m *ContentType) GetColumnPositions()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.columnPositions
    }
}
// GetColumns gets the columns property value. The collection of column definitions for this contentType.
func (m *ContentType) GetColumns()([]ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
// GetDescription gets the description property value. The descriptive text for the item.
func (m *ContentType) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDocumentSet gets the documentSet property value. Document Set metadata.
func (m *ContentType) GetDocumentSet()(*DocumentSet) {
    if m == nil {
        return nil
    } else {
        return m.documentSet
    }
}
// GetDocumentTemplate gets the documentTemplate property value. Document template metadata. To make sure that documents have consistent content across a site and its subsites, you can associate a Word, Excel, or PowerPoint template with a site content type.
func (m *ContentType) GetDocumentTemplate()(*DocumentSetContent) {
    if m == nil {
        return nil
    } else {
        return m.documentTemplate
    }
}
// GetGroup gets the group property value. The name of the group this content type belongs to. Helps organize related content types.
func (m *ContentType) GetGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.group
    }
}
// GetHidden gets the hidden property value. Indicates whether the content type is hidden in the list's 'New' menu.
func (m *ContentType) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// GetInheritedFrom gets the inheritedFrom property value. If this content type is inherited from another scope (like a site), provides a reference to the item where the content type is defined.
func (m *ContentType) GetInheritedFrom()(*ItemReference) {
    if m == nil {
        return nil
    } else {
        return m.inheritedFrom
    }
}
// GetIsBuiltIn gets the isBuiltIn property value. Specifies if a content type is a built-in content type.
func (m *ContentType) GetIsBuiltIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBuiltIn
    }
}
// GetName gets the name property value. The name of the content type.
func (m *ContentType) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetOrder gets the order property value. Specifies the order in which the content type appears in the selection UI.
func (m *ContentType) GetOrder()(*ContentTypeOrder) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
// GetParentId gets the parentId property value. The unique identifier of the content type.
func (m *ContentType) GetParentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentId
    }
}
// GetPropagateChanges gets the propagateChanges property value. If true, any changes made to the content type will be pushed to inherited content types and lists that implement the content type.
func (m *ContentType) GetPropagateChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateChanges
    }
}
// GetReadOnly gets the readOnly property value. If true, the content type can't be modified unless this value is first set to false.
func (m *ContentType) GetReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.readOnly
    }
}
// GetSealed gets the sealed property value. If true, the content type can't be modified by users or through push-down operations. Only site collection administrators can seal or unseal content types.
func (m *ContentType) GetSealed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sealed
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["associatedHubsUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssociatedHubsUrls(res)
        }
        return nil
    }
    res["base"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBase(val.(*ContentType))
        }
        return nil
    }
    res["baseTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentType() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContentType, len(val))
            for i, v := range val {
                res[i] = *(v.(*ContentType))
            }
            m.SetBaseTypes(res)
        }
        return nil
    }
    res["columnLinks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnLink() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnLink, len(val))
            for i, v := range val {
                res[i] = *(v.(*ColumnLink))
            }
            m.SetColumnLinks(res)
        }
        return nil
    }
    res["columnPositions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*ColumnDefinition))
            }
            m.SetColumnPositions(res)
        }
        return nil
    }
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*ColumnDefinition))
            }
            m.SetColumns(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["documentSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentSet(val.(*DocumentSet))
        }
        return nil
    }
    res["documentTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDocumentSetContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentTemplate(val.(*DocumentSetContent))
        }
        return nil
    }
    res["group"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val)
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
    res["inheritedFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemReference() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInheritedFrom(val.(*ItemReference))
        }
        return nil
    }
    res["isBuiltIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["order"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeOrder() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrder(val.(*ContentTypeOrder))
        }
        return nil
    }
    res["parentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentId(val)
        }
        return nil
    }
    res["propagateChanges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropagateChanges(val)
        }
        return nil
    }
    res["readOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOnly(val)
        }
        return nil
    }
    res["sealed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSealed(val)
        }
        return nil
    }
    return res
}
func (m *ContentType) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAssociatedHubsUrls sets the associatedHubsUrls property value. List of canonical URLs for hub sites with which this content type is associated to. This will contain all hub sites where this content type is queued to be enforced or is already enforced. Enforcing a content type means that the content type will be applied to the lists in the enforced sites.
func (m *ContentType) SetAssociatedHubsUrls(value []string)() {
    if m != nil {
        m.associatedHubsUrls = value
    }
}
// SetBase sets the base property value. Parent contentType from which this content type is derived.
func (m *ContentType) SetBase(value *ContentType)() {
    if m != nil {
        m.base = value
    }
}
// SetBaseTypes sets the baseTypes property value. The collection of content types that are ancestors of this content type.
func (m *ContentType) SetBaseTypes(value []ContentType)() {
    if m != nil {
        m.baseTypes = value
    }
}
// SetColumnLinks sets the columnLinks property value. The collection of columns that are required by this content type.
func (m *ContentType) SetColumnLinks(value []ColumnLink)() {
    if m != nil {
        m.columnLinks = value
    }
}
// SetColumnPositions sets the columnPositions property value. Column order information in a content type.
func (m *ContentType) SetColumnPositions(value []ColumnDefinition)() {
    if m != nil {
        m.columnPositions = value
    }
}
// SetColumns sets the columns property value. The collection of column definitions for this contentType.
func (m *ContentType) SetColumns(value []ColumnDefinition)() {
    if m != nil {
        m.columns = value
    }
}
// SetDescription sets the description property value. The descriptive text for the item.
func (m *ContentType) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDocumentSet sets the documentSet property value. Document Set metadata.
func (m *ContentType) SetDocumentSet(value *DocumentSet)() {
    if m != nil {
        m.documentSet = value
    }
}
// SetDocumentTemplate sets the documentTemplate property value. Document template metadata. To make sure that documents have consistent content across a site and its subsites, you can associate a Word, Excel, or PowerPoint template with a site content type.
func (m *ContentType) SetDocumentTemplate(value *DocumentSetContent)() {
    if m != nil {
        m.documentTemplate = value
    }
}
// SetGroup sets the group property value. The name of the group this content type belongs to. Helps organize related content types.
func (m *ContentType) SetGroup(value *string)() {
    if m != nil {
        m.group = value
    }
}
// SetHidden sets the hidden property value. Indicates whether the content type is hidden in the list's 'New' menu.
func (m *ContentType) SetHidden(value *bool)() {
    if m != nil {
        m.hidden = value
    }
}
// SetInheritedFrom sets the inheritedFrom property value. If this content type is inherited from another scope (like a site), provides a reference to the item where the content type is defined.
func (m *ContentType) SetInheritedFrom(value *ItemReference)() {
    if m != nil {
        m.inheritedFrom = value
    }
}
// SetIsBuiltIn sets the isBuiltIn property value. Specifies if a content type is a built-in content type.
func (m *ContentType) SetIsBuiltIn(value *bool)() {
    if m != nil {
        m.isBuiltIn = value
    }
}
// SetName sets the name property value. The name of the content type.
func (m *ContentType) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetOrder sets the order property value. Specifies the order in which the content type appears in the selection UI.
func (m *ContentType) SetOrder(value *ContentTypeOrder)() {
    if m != nil {
        m.order = value
    }
}
// SetParentId sets the parentId property value. The unique identifier of the content type.
func (m *ContentType) SetParentId(value *string)() {
    if m != nil {
        m.parentId = value
    }
}
// SetPropagateChanges sets the propagateChanges property value. If true, any changes made to the content type will be pushed to inherited content types and lists that implement the content type.
func (m *ContentType) SetPropagateChanges(value *bool)() {
    if m != nil {
        m.propagateChanges = value
    }
}
// SetReadOnly sets the readOnly property value. If true, the content type can't be modified unless this value is first set to false.
func (m *ContentType) SetReadOnly(value *bool)() {
    if m != nil {
        m.readOnly = value
    }
}
// SetSealed sets the sealed property value. If true, the content type can't be modified by users or through push-down operations. Only site collection administrators can seal or unseal content types.
func (m *ContentType) SetSealed(value *bool)() {
    if m != nil {
        m.sealed = value
    }
}
