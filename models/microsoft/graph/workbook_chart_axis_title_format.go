package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartAxisTitleFormat provides operations to manage the collection of drive entities.
type WorkbookChartAxisTitleFormat struct {
    Entity
    // Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
    font WorkbookChartFontable;
}
// NewWorkbookChartAxisTitleFormat instantiates a new workbookChartAxisTitleFormat and sets the default values.
func NewWorkbookChartAxisTitleFormat()(*WorkbookChartAxisTitleFormat) {
    m := &WorkbookChartAxisTitleFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartAxisTitleFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartAxisTitleFormatFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartAxisTitleFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartAxisTitleFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["font"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFontFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFont(val.(WorkbookChartFontable))
        }
        return nil
    }
    return res
}
// GetFont gets the font property value. Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
func (m *WorkbookChartAxisTitleFormat) GetFont()(WorkbookChartFontable) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
func (m *WorkbookChartAxisTitleFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartAxisTitleFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFont sets the font property value. Represents the font attributes, such as font name, font size, color, etc. of chart axis title object. Read-only.
func (m *WorkbookChartAxisTitleFormat) SetFont(value WorkbookChartFontable)() {
    if m != nil {
        m.font = value
    }
}
