package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContentTypeable 
type ContentTypeable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssociatedHubsUrls()([]string)
    GetBase()(ContentTypeable)
    GetBaseTypes()([]ContentTypeable)
    GetColumnLinks()([]ColumnLinkable)
    GetColumnPositions()([]ColumnDefinitionable)
    GetColumns()([]ColumnDefinitionable)
    GetDescription()(*string)
    GetDocumentSet()(DocumentSetable)
    GetDocumentTemplate()(DocumentSetContentable)
    GetGroup()(*string)
    GetHidden()(*bool)
    GetInheritedFrom()(ItemReferenceable)
    GetIsBuiltIn()(*bool)
    GetName()(*string)
    GetOrder()(ContentTypeOrderable)
    GetParentId()(*string)
    GetPropagateChanges()(*bool)
    GetReadOnly()(*bool)
    GetSealed()(*bool)
    SetAssociatedHubsUrls(value []string)()
    SetBase(value ContentTypeable)()
    SetBaseTypes(value []ContentTypeable)()
    SetColumnLinks(value []ColumnLinkable)()
    SetColumnPositions(value []ColumnDefinitionable)()
    SetColumns(value []ColumnDefinitionable)()
    SetDescription(value *string)()
    SetDocumentSet(value DocumentSetable)()
    SetDocumentTemplate(value DocumentSetContentable)()
    SetGroup(value *string)()
    SetHidden(value *bool)()
    SetInheritedFrom(value ItemReferenceable)()
    SetIsBuiltIn(value *bool)()
    SetName(value *string)()
    SetOrder(value ContentTypeOrderable)()
    SetParentId(value *string)()
    SetPropagateChanges(value *bool)()
    SetReadOnly(value *bool)()
    SetSealed(value *bool)()
}
