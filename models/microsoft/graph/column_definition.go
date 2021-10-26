package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new columnDefinition and sets the default values.
func NewColumnDefinition()(*ColumnDefinition) {
    m := &ColumnDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the boolean property value. This column stores boolean values.
func (m *ColumnDefinition) GetBoolean()(*BooleanColumn) {
    if m == nil {
        return nil
    } else {
        return m.boolean
    }
}
// Gets the calculated property value. This column's data is calculated based on other columns.
func (m *ColumnDefinition) GetCalculated()(*CalculatedColumn) {
    if m == nil {
        return nil
    } else {
        return m.calculated
    }
}
// Gets the choice property value. This column stores data from a list of choices.
func (m *ColumnDefinition) GetChoice()(*ChoiceColumn) {
    if m == nil {
        return nil
    } else {
        return m.choice
    }
}
// Gets the columnGroup property value. For site columns, the name of the group this column belongs to. Helps organize related columns.
func (m *ColumnDefinition) GetColumnGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.columnGroup
    }
}
// Gets the contentApprovalStatus property value. This column stores content approval status.
func (m *ColumnDefinition) GetContentApprovalStatus()(*ContentApprovalStatusColumn) {
    if m == nil {
        return nil
    } else {
        return m.contentApprovalStatus
    }
}
// Gets the currency property value. This column stores currency values.
func (m *ColumnDefinition) GetCurrency()(*CurrencyColumn) {
    if m == nil {
        return nil
    } else {
        return m.currency
    }
}
// Gets the dateTime property value. This column stores DateTime values.
func (m *ColumnDefinition) GetDateTime()(*DateTimeColumn) {
    if m == nil {
        return nil
    } else {
        return m.dateTime
    }
}
// Gets the defaultValue property value. The default value for this column.
func (m *ColumnDefinition) GetDefaultValue()(*DefaultColumnValue) {
    if m == nil {
        return nil
    } else {
        return m.defaultValue
    }
}
// Gets the description property value. The user-facing description of the column.
func (m *ColumnDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The user-facing name of the column.
func (m *ColumnDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enforceUniqueValues property value. If true, no two list items may have the same value for this column.
func (m *ColumnDefinition) GetEnforceUniqueValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enforceUniqueValues
    }
}
// Gets the geolocation property value. This column stores a geolocation.
func (m *ColumnDefinition) GetGeolocation()(*GeolocationColumn) {
    if m == nil {
        return nil
    } else {
        return m.geolocation
    }
}
// Gets the hidden property value. Specifies whether the column is displayed in the user interface.
func (m *ColumnDefinition) GetHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hidden
    }
}
// Gets the hyperlinkOrPicture property value. This column stores hyperlink or picture values.
func (m *ColumnDefinition) GetHyperlinkOrPicture()(*HyperlinkOrPictureColumn) {
    if m == nil {
        return nil
    } else {
        return m.hyperlinkOrPicture
    }
}
// Gets the indexed property value. Specifies whether the column values can be used for sorting and searching.
func (m *ColumnDefinition) GetIndexed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.indexed
    }
}
// Gets the isDeletable property value. Indicates whether this column can be deleted.
func (m *ColumnDefinition) GetIsDeletable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeletable
    }
}
// Gets the isReorderable property value. Indicates whether values in the column can be reordered. Read-only.
func (m *ColumnDefinition) GetIsReorderable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReorderable
    }
}
// Gets the isSealed property value. Specifies whether the column can be changed.
func (m *ColumnDefinition) GetIsSealed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSealed
    }
}
// Gets the lookup property value. This column's data is looked up from another source in the site.
func (m *ColumnDefinition) GetLookup()(*LookupColumn) {
    if m == nil {
        return nil
    } else {
        return m.lookup
    }
}
// Gets the name property value. The API-facing name of the column as it appears in the [fields][] on a [listItem][]. For the user-facing name, see displayName.
func (m *ColumnDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the number property value. This column stores number values.
func (m *ColumnDefinition) GetNumber()(*NumberColumn) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the personOrGroup property value. This column stores Person or Group values.
func (m *ColumnDefinition) GetPersonOrGroup()(*PersonOrGroupColumn) {
    if m == nil {
        return nil
    } else {
        return m.personOrGroup
    }
}
// Gets the propagateChanges property value. If 'true', changes to this column will be propagated to lists that implement the column.
func (m *ColumnDefinition) GetPropagateChanges()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateChanges
    }
}
// Gets the readOnly property value. Specifies whether the column values can be modified.
func (m *ColumnDefinition) GetReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.readOnly
    }
}
// Gets the required property value. Specifies whether the column value isn't optional.
func (m *ColumnDefinition) GetRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.required
    }
}
// Gets the sourceColumn property value. The source column for the content type column.
func (m *ColumnDefinition) GetSourceColumn()(*ColumnDefinition) {
    if m == nil {
        return nil
    } else {
        return m.sourceColumn
    }
}
// Gets the term property value. This column stores taxonomy terms.
func (m *ColumnDefinition) GetTerm()(*TermColumn) {
    if m == nil {
        return nil
    } else {
        return m.term
    }
}
// Gets the text property value. This column stores text values.
func (m *ColumnDefinition) GetText()(*TextColumn) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// Gets the thumbnail property value. This column stores thumbnail values.
func (m *ColumnDefinition) GetThumbnail()(*ThumbnailColumn) {
    if m == nil {
        return nil
    } else {
        return m.thumbnail
    }
}
// Gets the type_escaped property value. For site columns, the type of column. Read-only.
func (m *ColumnDefinition) GetType_escaped()(*ColumnTypes) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the validation property value. This column stores validation formula and message for the column.
func (m *ColumnDefinition) GetValidation()(*ColumnValidation) {
    if m == nil {
        return nil
    } else {
        return m.validation
    }
}
// The deserialization information for the current model
func (m *ColumnDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["boolean"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBooleanColumn() })
        if err != nil {
            return err
        }
        m.SetBoolean(val.(*BooleanColumn))
        return nil
    }
    res["calculated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalculatedColumn() })
        if err != nil {
            return err
        }
        m.SetCalculated(val.(*CalculatedColumn))
        return nil
    }
    res["choice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChoiceColumn() })
        if err != nil {
            return err
        }
        m.SetChoice(val.(*ChoiceColumn))
        return nil
    }
    res["columnGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetColumnGroup(val)
        return nil
    }
    res["contentApprovalStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentApprovalStatusColumn() })
        if err != nil {
            return err
        }
        m.SetContentApprovalStatus(val.(*ContentApprovalStatusColumn))
        return nil
    }
    res["currency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCurrencyColumn() })
        if err != nil {
            return err
        }
        m.SetCurrency(val.(*CurrencyColumn))
        return nil
    }
    res["dateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeColumn() })
        if err != nil {
            return err
        }
        m.SetDateTime(val.(*DateTimeColumn))
        return nil
    }
    res["defaultValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDefaultColumnValue() })
        if err != nil {
            return err
        }
        m.SetDefaultValue(val.(*DefaultColumnValue))
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
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["enforceUniqueValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnforceUniqueValues(val)
        return nil
    }
    res["geolocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGeolocationColumn() })
        if err != nil {
            return err
        }
        m.SetGeolocation(val.(*GeolocationColumn))
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
    res["hyperlinkOrPicture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHyperlinkOrPictureColumn() })
        if err != nil {
            return err
        }
        m.SetHyperlinkOrPicture(val.(*HyperlinkOrPictureColumn))
        return nil
    }
    res["indexed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIndexed(val)
        return nil
    }
    res["isDeletable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeletable(val)
        return nil
    }
    res["isReorderable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsReorderable(val)
        return nil
    }
    res["isSealed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSealed(val)
        return nil
    }
    res["lookup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLookupColumn() })
        if err != nil {
            return err
        }
        m.SetLookup(val.(*LookupColumn))
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
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNumberColumn() })
        if err != nil {
            return err
        }
        m.SetNumber(val.(*NumberColumn))
        return nil
    }
    res["personOrGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonOrGroupColumn() })
        if err != nil {
            return err
        }
        m.SetPersonOrGroup(val.(*PersonOrGroupColumn))
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
    res["required"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequired(val)
        return nil
    }
    res["sourceColumn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnDefinition() })
        if err != nil {
            return err
        }
        m.SetSourceColumn(val.(*ColumnDefinition))
        return nil
    }
    res["term"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermColumn() })
        if err != nil {
            return err
        }
        m.SetTerm(val.(*TermColumn))
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTextColumn() })
        if err != nil {
            return err
        }
        m.SetText(val.(*TextColumn))
        return nil
    }
    res["thumbnail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewThumbnailColumn() })
        if err != nil {
            return err
        }
        m.SetThumbnail(val.(*ThumbnailColumn))
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseColumnTypes)
        if err != nil {
            return err
        }
        cast := val.(ColumnTypes)
        m.SetType_escaped(&cast)
        return nil
    }
    res["validation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewColumnValidation() })
        if err != nil {
            return err
        }
        m.SetValidation(val.(*ColumnValidation))
        return nil
    }
    return res
}
func (m *ColumnDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the boolean property value. This column stores boolean values.
// Parameters:
//  - value : Value to set for the boolean property.
func (m *ColumnDefinition) SetBoolean(value *BooleanColumn)() {
    m.boolean = value
}
// Sets the calculated property value. This column's data is calculated based on other columns.
// Parameters:
//  - value : Value to set for the calculated property.
func (m *ColumnDefinition) SetCalculated(value *CalculatedColumn)() {
    m.calculated = value
}
// Sets the choice property value. This column stores data from a list of choices.
// Parameters:
//  - value : Value to set for the choice property.
func (m *ColumnDefinition) SetChoice(value *ChoiceColumn)() {
    m.choice = value
}
// Sets the columnGroup property value. For site columns, the name of the group this column belongs to. Helps organize related columns.
// Parameters:
//  - value : Value to set for the columnGroup property.
func (m *ColumnDefinition) SetColumnGroup(value *string)() {
    m.columnGroup = value
}
// Sets the contentApprovalStatus property value. This column stores content approval status.
// Parameters:
//  - value : Value to set for the contentApprovalStatus property.
func (m *ColumnDefinition) SetContentApprovalStatus(value *ContentApprovalStatusColumn)() {
    m.contentApprovalStatus = value
}
// Sets the currency property value. This column stores currency values.
// Parameters:
//  - value : Value to set for the currency property.
func (m *ColumnDefinition) SetCurrency(value *CurrencyColumn)() {
    m.currency = value
}
// Sets the dateTime property value. This column stores DateTime values.
// Parameters:
//  - value : Value to set for the dateTime property.
func (m *ColumnDefinition) SetDateTime(value *DateTimeColumn)() {
    m.dateTime = value
}
// Sets the defaultValue property value. The default value for this column.
// Parameters:
//  - value : Value to set for the defaultValue property.
func (m *ColumnDefinition) SetDefaultValue(value *DefaultColumnValue)() {
    m.defaultValue = value
}
// Sets the description property value. The user-facing description of the column.
// Parameters:
//  - value : Value to set for the description property.
func (m *ColumnDefinition) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The user-facing name of the column.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ColumnDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enforceUniqueValues property value. If true, no two list items may have the same value for this column.
// Parameters:
//  - value : Value to set for the enforceUniqueValues property.
func (m *ColumnDefinition) SetEnforceUniqueValues(value *bool)() {
    m.enforceUniqueValues = value
}
// Sets the geolocation property value. This column stores a geolocation.
// Parameters:
//  - value : Value to set for the geolocation property.
func (m *ColumnDefinition) SetGeolocation(value *GeolocationColumn)() {
    m.geolocation = value
}
// Sets the hidden property value. Specifies whether the column is displayed in the user interface.
// Parameters:
//  - value : Value to set for the hidden property.
func (m *ColumnDefinition) SetHidden(value *bool)() {
    m.hidden = value
}
// Sets the hyperlinkOrPicture property value. This column stores hyperlink or picture values.
// Parameters:
//  - value : Value to set for the hyperlinkOrPicture property.
func (m *ColumnDefinition) SetHyperlinkOrPicture(value *HyperlinkOrPictureColumn)() {
    m.hyperlinkOrPicture = value
}
// Sets the indexed property value. Specifies whether the column values can be used for sorting and searching.
// Parameters:
//  - value : Value to set for the indexed property.
func (m *ColumnDefinition) SetIndexed(value *bool)() {
    m.indexed = value
}
// Sets the isDeletable property value. Indicates whether this column can be deleted.
// Parameters:
//  - value : Value to set for the isDeletable property.
func (m *ColumnDefinition) SetIsDeletable(value *bool)() {
    m.isDeletable = value
}
// Sets the isReorderable property value. Indicates whether values in the column can be reordered. Read-only.
// Parameters:
//  - value : Value to set for the isReorderable property.
func (m *ColumnDefinition) SetIsReorderable(value *bool)() {
    m.isReorderable = value
}
// Sets the isSealed property value. Specifies whether the column can be changed.
// Parameters:
//  - value : Value to set for the isSealed property.
func (m *ColumnDefinition) SetIsSealed(value *bool)() {
    m.isSealed = value
}
// Sets the lookup property value. This column's data is looked up from another source in the site.
// Parameters:
//  - value : Value to set for the lookup property.
func (m *ColumnDefinition) SetLookup(value *LookupColumn)() {
    m.lookup = value
}
// Sets the name property value. The API-facing name of the column as it appears in the [fields][] on a [listItem][]. For the user-facing name, see displayName.
// Parameters:
//  - value : Value to set for the name property.
func (m *ColumnDefinition) SetName(value *string)() {
    m.name = value
}
// Sets the number property value. This column stores number values.
// Parameters:
//  - value : Value to set for the number property.
func (m *ColumnDefinition) SetNumber(value *NumberColumn)() {
    m.number = value
}
// Sets the personOrGroup property value. This column stores Person or Group values.
// Parameters:
//  - value : Value to set for the personOrGroup property.
func (m *ColumnDefinition) SetPersonOrGroup(value *PersonOrGroupColumn)() {
    m.personOrGroup = value
}
// Sets the propagateChanges property value. If 'true', changes to this column will be propagated to lists that implement the column.
// Parameters:
//  - value : Value to set for the propagateChanges property.
func (m *ColumnDefinition) SetPropagateChanges(value *bool)() {
    m.propagateChanges = value
}
// Sets the readOnly property value. Specifies whether the column values can be modified.
// Parameters:
//  - value : Value to set for the readOnly property.
func (m *ColumnDefinition) SetReadOnly(value *bool)() {
    m.readOnly = value
}
// Sets the required property value. Specifies whether the column value isn't optional.
// Parameters:
//  - value : Value to set for the required property.
func (m *ColumnDefinition) SetRequired(value *bool)() {
    m.required = value
}
// Sets the sourceColumn property value. The source column for the content type column.
// Parameters:
//  - value : Value to set for the sourceColumn property.
func (m *ColumnDefinition) SetSourceColumn(value *ColumnDefinition)() {
    m.sourceColumn = value
}
// Sets the term property value. This column stores taxonomy terms.
// Parameters:
//  - value : Value to set for the term property.
func (m *ColumnDefinition) SetTerm(value *TermColumn)() {
    m.term = value
}
// Sets the text property value. This column stores text values.
// Parameters:
//  - value : Value to set for the text property.
func (m *ColumnDefinition) SetText(value *TextColumn)() {
    m.text = value
}
// Sets the thumbnail property value. This column stores thumbnail values.
// Parameters:
//  - value : Value to set for the thumbnail property.
func (m *ColumnDefinition) SetThumbnail(value *ThumbnailColumn)() {
    m.thumbnail = value
}
// Sets the type_escaped property value. For site columns, the type of column. Read-only.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *ColumnDefinition) SetType_escaped(value *ColumnTypes)() {
    m.type_escaped = value
}
// Sets the validation property value. This column stores validation formula and message for the column.
// Parameters:
//  - value : Value to set for the validation property.
func (m *ColumnDefinition) SetValidation(value *ColumnValidation)() {
    m.validation = value
}
