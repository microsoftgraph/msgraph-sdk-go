package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChartDataLabelFormat provides operations to manage the drive singleton.
type WorkbookChartDataLabelFormat struct {
    Entity
    // Represents the fill format of the current chart data label. Read-only.
    fill WorkbookChartFillable;
    // Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
    font WorkbookChartFontable;
}
// NewWorkbookChartDataLabelFormat instantiates a new workbookChartDataLabelFormat and sets the default values.
func NewWorkbookChartDataLabelFormat()(*WorkbookChartDataLabelFormat) {
    m := &WorkbookChartDataLabelFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartDataLabelFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartDataLabelFormatFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewWorkbookChartDataLabelFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartDataLabelFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFillFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(WorkbookChartFillable))
        }
        return nil
    }
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
// GetFill gets the fill property value. Represents the fill format of the current chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) GetFill()(WorkbookChartFillable) {
    if m == nil {
        return nil
    } else {
        return m.fill
    }
}
// GetFont gets the font property value. Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) GetFont()(WorkbookChartFontable) {
    if m == nil {
        return nil
    } else {
        return m.font
    }
}
func (m *WorkbookChartDataLabelFormat) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChartDataLabelFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fill", m.GetFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("font", m.GetFont())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFill sets the fill property value. Represents the fill format of the current chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) SetFill(value WorkbookChartFillable)() {
    if m != nil {
        m.fill = value
    }
}
// SetFont sets the font property value. Represents the font attributes (font name, font size, color, etc.) for a chart data label. Read-only.
func (m *WorkbookChartDataLabelFormat) SetFont(value WorkbookChartFontable)() {
    if m != nil {
        m.font = value
    }
}
