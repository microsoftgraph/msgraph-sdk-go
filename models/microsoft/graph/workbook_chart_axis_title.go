package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxisTitle provides operations to manage the collection of drive entities.
type WorkbookChartAxisTitle struct {
    Entity
    // Represents the formatting of chart axis title. Read-only.
    format WorkbookChartAxisTitleFormatable;
    // Represents the axis title.
    text *string;
    // A boolean that specifies the visibility of an axis title.
    visible *bool;
}
// NewWorkbookChartAxisTitle instantiates a new workbookChartAxisTitle and sets the default values.
func NewWorkbookChartAxisTitle()(*WorkbookChartAxisTitle) {
    m := &WorkbookChartAxisTitle{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxisTitleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxisTitleFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartAxisTitle(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxisTitle) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartAxisTitleFormatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(WorkbookChartAxisTitleFormatable))
        }
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    res["visible"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisible(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Represents the formatting of chart axis title. Read-only.
func (m *WorkbookChartAxisTitle) GetFormat()(WorkbookChartAxisTitleFormatable) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetText gets the text property value. Represents the axis title.
func (m *WorkbookChartAxisTitle) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
// GetVisible gets the visible property value. A boolean that specifies the visibility of an axis title.
func (m *WorkbookChartAxisTitle) GetVisible()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.visible
    }
}
func (m *WorkbookChartAxisTitle) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartAxisTitle) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("visible", m.GetVisible())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Represents the formatting of chart axis title. Read-only.
func (m *WorkbookChartAxisTitle) SetFormat(value WorkbookChartAxisTitleFormatable)() {
    if m != nil {
        m.format = value
    }
}
// SetText sets the text property value. Represents the axis title.
func (m *WorkbookChartAxisTitle) SetText(value *string)() {
    if m != nil {
        m.text = value
    }
}
// SetVisible sets the visible property value. A boolean that specifies the visibility of an axis title.
func (m *WorkbookChartAxisTitle) SetVisible(value *bool)() {
    if m != nil {
        m.visible = value
    }
}
