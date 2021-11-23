package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// columnDefinition 
type ColumnDefinition struct {
    Entity
    // This column stores boolean values.
    boolean *BooleanColumn;
    // This column's data is calculated based on other columns.
    calculated *CalculatedColumn;
    // This column stores data from a list of choices.
    choice *ChoiceColumn;
    // For site columns, the name of the group this column belongs to. Helps organize related columns.
    columnGroup *string;
    // This column stores content approval status.
    contentApprovalStatus *ContentApprovalStatusColumn;
    // This column stores currency values.
    currency *CurrencyColumn;
    // This column stores DateTime values.
    dateTime *DateTimeColumn;
    // The default value for this column.
    defaultValue *DefaultColumnValue;
    // The user-facing description of the column.
    description *string;
    // The user-facing name of the column.
    displayName *string;
    // If true, no two list items may have the same value for this column.
    enforceUniqueValues *bool;
    // This column stores a geolocation.
    geolocation *GeolocationColumn;
    // Specifies whether the column is displayed in the user interface.
    hidden *bool;
    // This column stores hyperlink or picture values.
    hyperlinkOrPicture *HyperlinkOrPictureColumn;
    // Specifies whether the column values can be used for sorting and searching.
    indexed *bool;
    // Indicates whether this column can be deleted.
    isDeletable *bool;
    // Indicates whether values in the column can be reordered. Read-only.
    isReorderable *bool;
    // Specifies whether the column can be changed.
    isSealed *bool;
    // This column's data is looked up from another source in the site.
    lookup *LookupColumn;
    // The API-facing name of the column as it appears in the [fields][] on a [listItem][]. For the user-facing name, see displayName.
    name *string;
    // This column stores number values.
    number *NumberColumn;
    // This column stores Person or Group values.
    personOrGroup *PersonOrGroupColumn;
    // If 'true', changes to this column will be propagated to lists that implement the column.
    propagateChanges *bool;
    // Specifies whether the column values can be modified.
    readOnly *bool;
    // Specifies whether the column value isn't optional.
    required *bool;
    // The source column for the content type column.
    sourceColumn *ColumnDefinition;
    // This column stores taxonomy terms.
    term *TermColumn;
    // This column stores text values.
    text *TextColumn;
    // This column stores thumbnail values.
    thumbnail *ThumbnailColumn;
    // For site columns, the type of column. Read-only.
    type_escaped *ColumnTypes;
    // This column stores validation formula and message for the column.
    validation *ColumnValidation;
}
// NewColumnDefinition instantiates a new columnDefinition and sets the default values.
func NewColumnDefinition()(*ColumnDefinition) {
    m := &ColumnDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// GetBoolean gets the boolean property value. This column stores boolean values.
func (m *ColumnDefinition) GetBoolean()(*BooleanColumn) {
    if m == nil {
        return nil
    } else {
        return m.boolean
    }
}
// GetCalculated gets the calculated property value. This column's data is calculated based on other columns.
func (m *ColumnDefinition) GetCalculated()(*CalculatedColumn) {
    if m == nil {
        return nil
    } else {
        return m.calculated
    }
}
// GetChoice gets the choice property value. This column stores data from a list of choices.
func (m *ColumnDefinition) GetChoice()(*ChoiceColumn) {
    if m == nil {
        return nil
    } else {
        return m.choice
    }
}
// GetColumnGroup gets the columnGroup property value. For site columns, the name of the group this column belongs to. Helps organize related columns.
func (m *ColumnDefinition) GetColumnGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.columnGroup
    }
}
// GetContentApprovalStatus gets the contentApprovalStatus property value. This column stores content approval status.
func (m *ColumnDefinition) GetContentApprovalStatus()(*ContentApprovalStatusColumn) {
    if m == nil {
        return nil
    } else {
        return m.contentApprovalStatus
    }
}
// GetCurrency gets the currency property value. This column stores currency values.
func (m *ColumnDefinition) GetCurrency()(*CurrencyColumn) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// GetDateTime gets the dateTime property value. This column stores DateTime values.
func (m *ColumnDefinition) GetDateTime()(*DateTimeColumn) {
    if m == nil {
        return nil
    } else {
        return m.dateTime
    }
}
// GetDefaultValue gets the defaultValue property value. The default value for this column.
func (m *ColumnDefinition) GetDefaultValue()(*DefaultColumnValue) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// GetDescription gets the description property value. The user-facing description of the column.
func (m *ColumnDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The user-facing name of the column.
func (m *ColumnDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnforceUniqueValues gets the enforceUniqueValues property value. If true, no two list items may have the same value for this column.
func (m *ColumnDefinition) GetEnforceUniqueValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enforceUniqueValues
    }
}
// GetGeolocation gets the geolocation property value. This column stores a geolocation.
func (m *ColumnDefinition) GetGeolocation()(*GeolocationColumn) {
    if m == nil {
        return nil
    } else {
        return m.geolocation
    }
}
// GetHidden gets the hidden property value. Specifies whether the column is displayed in the user interface.
func (m *ColumnDefinition) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// GetHyperlinkOrPicture gets the hyperlinkOrPicture property value. This column stores hyperlink or picture values.
func (m *ColumnDefinition) GetHyperlinkOrPicture()(*HyperlinkOrPictureColumn) {
    if m == nil {
        return nil
    } else {
        return m.hyperlinkOrPicture
    }
}
// GetIndexed gets the indexed property value. Specifies whether the column values can be used for sorting and searching.
func (m *ColumnDefinition) GetIndexed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.indexed
    }
}
// GetIsDeletable gets the isDeletable property value. Indicates whether this column can be deleted.
func (m *ColumnDefinition) GetIsDeletable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeletable
    }
}
// GetIsReorderable gets the isReorderable property value. Indicates whether values in the column can be reordered. Read-only.
func (m *ColumnDefinition) GetIsReorderable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReorderable
    }
}
// GetIsSealed gets the isSealed property value. Specifies whether the column can be changed.
func (m *ColumnDefinition) GetIsSealed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSealed
    }
}
// GetLookup gets the lookup property value. This column's data is looked up from another source in the site.
func (m *ColumnDefinition) GetLookup()(*LookupColumn) {
    if m == nil {
        return nil
    } else {
        return m.lookup
    }
}
// GetName gets the name property value. The API-facing name of the column as it appears in the [fields][] on a [listItem][]. For the user-facing name, see displayName.
func (m *ColumnDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetNumber gets the number property value. This column stores number values.
func (m *ColumnDefinition) GetNumber()(*NumberColumn) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// GetPersonOrGroup gets the personOrGroup property value. This column stores Person or Group values.
func (m *ColumnDefinition) GetPersonOrGroup()(*PersonOrGroupColumn) {
    if m == nil {
        return nil
    } else {
        return m.personOrGroup
    }
}
// GetPropagateChanges gets the propagateChanges property value. If 'true', changes to this column will be propagated to lists that implement the column.
func (m *ColumnDefinition) GetPropagateChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateChanges
    }
}
// GetReadOnly gets the readOnly property value. Specifies whether the column values can be modified.
func (m *ColumnDefinition) GetReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.readOnly
    }
}
// GetRequired gets the required property value. Specifies whether the column value isn't optional.
func (m *ColumnDefinition) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// GetSourceColumn gets the sourceColumn property value. The source column for the content type column.
func (m *ColumnDefinition) GetSourceColumn()(*ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.sourceColumn
    }
}
// GetTerm gets the term property value. This column stores taxonomy terms.
func (m *ColumnDefinition) GetTerm()(*TermColumn) {
    if m == nil {
        return nil
    } else {
        return m.term
    }
}
// GetText gets the text property value. This column stores text values.
func (m *ColumnDefinition) GetText()(*TextColumn) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetThumbnail gets the thumbnail property value. This column stores thumbnail values.
func (m *ColumnDefinition) GetThumbnail()(*ThumbnailColumn) {
    if m == nil {
        return nil
    } else {
        return m.thumbnail
    }
}
// GetType_escaped gets the type_escaped property value. For site columns, the type of column. Read-only.
func (m *ColumnDefinition) GetType_escaped()(*ColumnTypes) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetValidation gets the validation property value. This column stores validation formula and message for the column.
func (m *ColumnDefinition) GetValidation()(*ColumnValidation) {
    if m == nil {
        return nil
    } else {
        return m.validation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["boolean"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBooleanColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBoolean(val.(*BooleanColumn))
        }
        return nil
    }
    res["calculated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalculatedColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalculated(val.(*CalculatedColumn))
        }
        return nil
    }
    res["choice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChoiceColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChoice(val.(*ChoiceColumn))
        }
        return nil
    }
    res["columnGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnGroup(val)
        }
        return nil
    }
    res["contentApprovalStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentApprovalStatusColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentApprovalStatus(val.(*ContentApprovalStatusColumn))
        }
        return nil
    }
    res["currency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrencyColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val.(*CurrencyColumn))
        }
        return nil
    }
    res["dateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val.(*DateTimeColumn))
        }
        return nil
    }
    res["defaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDefaultColumnValue() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(*DefaultColumnValue))
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enforceUniqueValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceUniqueValues(val)
        }
        return nil
    }
    res["geolocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeolocationColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeolocation(val.(*GeolocationColumn))
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
    res["hyperlinkOrPicture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHyperlinkOrPictureColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHyperlinkOrPicture(val.(*HyperlinkOrPictureColumn))
        }
        return nil
    }
    res["indexed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndexed(val)
        }
        return nil
    }
    res["isDeletable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeletable(val)
        }
        return nil
    }
    res["isReorderable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReorderable(val)
        }
        return nil
    }
    res["isSealed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSealed(val)
        }
        return nil
    }
    res["lookup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLookupColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLookup(val.(*LookupColumn))
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
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNumberColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val.(*NumberColumn))
        }
        return nil
    }
    res["personOrGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonOrGroupColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonOrGroup(val.(*PersonOrGroupColumn))
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
    res["required"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequired(val)
        }
        return nil
    }
    res["sourceColumn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceColumn(val.(*ColumnDefinition))
        }
        return nil
    }
    res["term"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTerm(val.(*TermColumn))
        }
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTextColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(*TextColumn))
        }
        return nil
    }
    res["thumbnail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnailColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnail(val.(*ThumbnailColumn))
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseColumnTypes)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ColumnTypes)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    res["validation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnValidation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidation(val.(*ColumnValidation))
        }
        return nil
    }
    return res
}
func (m *ColumnDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ColumnDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("boolean", m.GetBoolean())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calculated", m.GetCalculated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("choice", m.GetChoice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("columnGroup", m.GetColumnGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentApprovalStatus", m.GetContentApprovalStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dateTime", m.GetDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultValue", m.GetDefaultValue())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enforceUniqueValues", m.GetEnforceUniqueValues())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("geolocation", m.GetGeolocation())
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
        err = writer.WriteObjectValue("hyperlinkOrPicture", m.GetHyperlinkOrPicture())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("indexed", m.GetIndexed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeletable", m.GetIsDeletable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReorderable", m.GetIsReorderable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSealed", m.GetIsSealed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lookup", m.GetLookup())
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
        err = writer.WriteObjectValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("personOrGroup", m.GetPersonOrGroup())
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
        err = writer.WriteBoolValue("required", m.GetRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceColumn", m.GetSourceColumn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("term", m.GetTerm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("thumbnail", m.GetThumbnail())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("validation", m.GetValidation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBoolean sets the boolean property value. This column stores boolean values.
func (m *ColumnDefinition) SetBoolean(value *BooleanColumn)() {
    m.boolean = value
}
// SetCalculated sets the calculated property value. This column's data is calculated based on other columns.
func (m *ColumnDefinition) SetCalculated(value *CalculatedColumn)() {
    m.calculated = value
}
// SetChoice sets the choice property value. This column stores data from a list of choices.
func (m *ColumnDefinition) SetChoice(value *ChoiceColumn)() {
    m.choice = value
}
// SetColumnGroup sets the columnGroup property value. For site columns, the name of the group this column belongs to. Helps organize related columns.
func (m *ColumnDefinition) SetColumnGroup(value *string)() {
    m.columnGroup = value
}
// SetContentApprovalStatus sets the contentApprovalStatus property value. This column stores content approval status.
func (m *ColumnDefinition) SetContentApprovalStatus(value *ContentApprovalStatusColumn)() {
    m.contentApprovalStatus = value
}
// SetCurrency sets the currency property value. This column stores currency values.
func (m *ColumnDefinition) SetCurrency(value *CurrencyColumn)() {
    m.currency = value
}
// SetDateTime sets the dateTime property value. This column stores DateTime values.
func (m *ColumnDefinition) SetDateTime(value *DateTimeColumn)() {
    m.dateTime = value
}
// SetDefaultValue sets the defaultValue property value. The default value for this column.
func (m *ColumnDefinition) SetDefaultValue(value *DefaultColumnValue)() {
    m.defaultValue = value
}
// SetDescription sets the description property value. The user-facing description of the column.
func (m *ColumnDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The user-facing name of the column.
func (m *ColumnDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnforceUniqueValues sets the enforceUniqueValues property value. If true, no two list items may have the same value for this column.
func (m *ColumnDefinition) SetEnforceUniqueValues(value *bool)() {
    m.enforceUniqueValues = value
}
// SetGeolocation sets the geolocation property value. This column stores a geolocation.
func (m *ColumnDefinition) SetGeolocation(value *GeolocationColumn)() {
    m.geolocation = value
}
// SetHidden sets the hidden property value. Specifies whether the column is displayed in the user interface.
func (m *ColumnDefinition) SetHidden(value *bool)() {
    m.hidden = value
}
// SetHyperlinkOrPicture sets the hyperlinkOrPicture property value. This column stores hyperlink or picture values.
func (m *ColumnDefinition) SetHyperlinkOrPicture(value *HyperlinkOrPictureColumn)() {
    m.hyperlinkOrPicture = value
}
// SetIndexed sets the indexed property value. Specifies whether the column values can be used for sorting and searching.
func (m *ColumnDefinition) SetIndexed(value *bool)() {
    m.indexed = value
}
// SetIsDeletable sets the isDeletable property value. Indicates whether this column can be deleted.
func (m *ColumnDefinition) SetIsDeletable(value *bool)() {
    m.isDeletable = value
}
// SetIsReorderable sets the isReorderable property value. Indicates whether values in the column can be reordered. Read-only.
func (m *ColumnDefinition) SetIsReorderable(value *bool)() {
    m.isReorderable = value
}
// SetIsSealed sets the isSealed property value. Specifies whether the column can be changed.
func (m *ColumnDefinition) SetIsSealed(value *bool)() {
    m.isSealed = value
}
// SetLookup sets the lookup property value. This column's data is looked up from another source in the site.
func (m *ColumnDefinition) SetLookup(value *LookupColumn)() {
    m.lookup = value
}
// SetName sets the name property value. The API-facing name of the column as it appears in the [fields][] on a [listItem][]. For the user-facing name, see displayName.
func (m *ColumnDefinition) SetName(value *string)() {
    m.name = value
}
// SetNumber sets the number property value. This column stores number values.
func (m *ColumnDefinition) SetNumber(value *NumberColumn)() {
    m.number = value
}
// SetPersonOrGroup sets the personOrGroup property value. This column stores Person or Group values.
func (m *ColumnDefinition) SetPersonOrGroup(value *PersonOrGroupColumn)() {
    m.personOrGroup = value
}
// SetPropagateChanges sets the propagateChanges property value. If 'true', changes to this column will be propagated to lists that implement the column.
func (m *ColumnDefinition) SetPropagateChanges(value *bool)() {
    m.propagateChanges = value
}
// SetReadOnly sets the readOnly property value. Specifies whether the column values can be modified.
func (m *ColumnDefinition) SetReadOnly(value *bool)() {
    m.readOnly = value
}
// SetRequired sets the required property value. Specifies whether the column value isn't optional.
func (m *ColumnDefinition) SetRequired(value *bool)() {
    m.required = value
}
// SetSourceColumn sets the sourceColumn property value. The source column for the content type column.
func (m *ColumnDefinition) SetSourceColumn(value *ColumnDefinition)() {
    m.sourceColumn = value
}
// SetTerm sets the term property value. This column stores taxonomy terms.
func (m *ColumnDefinition) SetTerm(value *TermColumn)() {
    m.term = value
}
// SetText sets the text property value. This column stores text values.
func (m *ColumnDefinition) SetText(value *TextColumn)() {
    m.text = value
}
// SetThumbnail sets the thumbnail property value. This column stores thumbnail values.
func (m *ColumnDefinition) SetThumbnail(value *ThumbnailColumn)() {
    m.thumbnail = value
}
// SetType_escaped sets the type_escaped property value. For site columns, the type of column. Read-only.
func (m *ColumnDefinition) SetType_escaped(value *ColumnTypes)() {
    m.type_escaped = value
}
// SetValidation sets the validation property value. This column stores validation formula and message for the column.
func (m *ColumnDefinition) SetValidation(value *ColumnValidation)() {
    m.validation = value
}
