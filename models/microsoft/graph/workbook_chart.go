package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WorkbookChart 
type WorkbookChart struct {
    Entity
    // Represents chart axes. Read-only.
    axes *WorkbookChartAxes;
    // Represents the datalabels on the chart. Read-only.
    dataLabels *WorkbookChartDataLabels;
    // Encapsulates the format properties for the chart area. Read-only.
    format *WorkbookChartAreaFormat;
    // Represents the height, in points, of the chart object.
    height *float64;
    // The distance, in points, from the left side of the chart to the worksheet origin.
    left *float64;
    // Represents the legend for the chart. Read-only.
    legend *WorkbookChartLegend;
    // Represents the name of a chart object.
    name *string;
    // Represents either a single series or collection of series in the chart. Read-only.
    series []WorkbookChartSeries;
    // Represents the title of the specified chart, including the text, visibility, position and formating of the title. Read-only.
    title *WorkbookChartTitle;
    // Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart).
    top *float64;
    // Represents the width, in points, of the chart object.
    width *float64;
    // The worksheet containing the current chart. Read-only.
    worksheet *WorkbookWorksheet;
}
// NewWorkbookChart instantiates a new workbookChart and sets the default values.
func NewWorkbookChart()(*WorkbookChart) {
    m := &WorkbookChart{
        Entity: *NewEntity(),
    }
    return m
}
// GetAxes gets the axes property value. Represents chart axes. Read-only.
func (m *WorkbookChart) GetAxes()(*WorkbookChartAxes) {
    if m == nil {
        return nil
    } else {
        return m.axes
    }
}
// GetDataLabels gets the dataLabels property value. Represents the datalabels on the chart. Read-only.
func (m *WorkbookChart) GetDataLabels()(*WorkbookChartDataLabels) {
    if m == nil {
        return nil
    } else {
        return m.dataLabels
    }
}
// GetFormat gets the format property value. Encapsulates the format properties for the chart area. Read-only.
func (m *WorkbookChart) GetFormat()(*WorkbookChartAreaFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetHeight gets the height property value. Represents the height, in points, of the chart object.
func (m *WorkbookChart) GetHeight()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// GetLeft gets the left property value. The distance, in points, from the left side of the chart to the worksheet origin.
func (m *WorkbookChart) GetLeft()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.left
    }
}
// GetLegend gets the legend property value. Represents the legend for the chart. Read-only.
func (m *WorkbookChart) GetLegend()(*WorkbookChartLegend) {
    if m == nil {
        return nil
    } else {
        return m.legend
    }
}
// GetName gets the name property value. Represents the name of a chart object.
func (m *WorkbookChart) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSeries gets the series property value. Represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChart) GetSeries()([]WorkbookChartSeries) {
    if m == nil {
        return nil
    } else {
        return m.series
    }
}
// GetTitle gets the title property value. Represents the title of the specified chart, including the text, visibility, position and formating of the title. Read-only.
func (m *WorkbookChart) GetTitle()(*WorkbookChartTitle) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetTop gets the top property value. Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart).
func (m *WorkbookChart) GetTop()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// GetWidth gets the width property value. Represents the width, in points, of the chart object.
func (m *WorkbookChart) GetWidth()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
// GetWorksheet gets the worksheet property value. The worksheet containing the current chart. Read-only.
func (m *WorkbookChart) GetWorksheet()(*WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChart) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["axes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxes() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAxes(val.(*WorkbookChartAxes))
        }
        return nil
    }
    res["dataLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartDataLabels() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataLabels(val.(*WorkbookChartDataLabels))
        }
        return nil
    }
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAreaFormat() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*WorkbookChartAreaFormat))
        }
        return nil
    }
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["left"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLeft(val)
        }
        return nil
    }
    res["legend"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartLegend() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLegend(val.(*WorkbookChartLegend))
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
    res["series"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartSeries() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkbookChartSeries, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkbookChartSeries))
            }
            m.SetSeries(res)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartTitle() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val.(*WorkbookChartTitle))
        }
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    res["worksheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorksheet(val.(*WorkbookWorksheet))
        }
        return nil
    }
    return res
}
func (m *WorkbookChart) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WorkbookChart) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("axes", m.GetAxes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("dataLabels", m.GetDataLabels())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("left", m.GetLeft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("legend", m.GetLegend())
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
    if m.GetSeries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSeries()))
        for i, v := range m.GetSeries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("series", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("top", m.GetTop())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("worksheet", m.GetWorksheet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAxes sets the axes property value. Represents chart axes. Read-only.
func (m *WorkbookChart) SetAxes(value *WorkbookChartAxes)() {
    if m != nil {
        m.axes = value
    }
}
// SetDataLabels sets the dataLabels property value. Represents the datalabels on the chart. Read-only.
func (m *WorkbookChart) SetDataLabels(value *WorkbookChartDataLabels)() {
    if m != nil {
        m.dataLabels = value
    }
}
// SetFormat sets the format property value. Encapsulates the format properties for the chart area. Read-only.
func (m *WorkbookChart) SetFormat(value *WorkbookChartAreaFormat)() {
    if m != nil {
        m.format = value
    }
}
// SetHeight sets the height property value. Represents the height, in points, of the chart object.
func (m *WorkbookChart) SetHeight(value *float64)() {
    if m != nil {
        m.height = value
    }
}
// SetLeft sets the left property value. The distance, in points, from the left side of the chart to the worksheet origin.
func (m *WorkbookChart) SetLeft(value *float64)() {
    if m != nil {
        m.left = value
    }
}
// SetLegend sets the legend property value. Represents the legend for the chart. Read-only.
func (m *WorkbookChart) SetLegend(value *WorkbookChartLegend)() {
    if m != nil {
        m.legend = value
    }
}
// SetName sets the name property value. Represents the name of a chart object.
func (m *WorkbookChart) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSeries sets the series property value. Represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChart) SetSeries(value []WorkbookChartSeries)() {
    if m != nil {
        m.series = value
    }
}
// SetTitle sets the title property value. Represents the title of the specified chart, including the text, visibility, position and formating of the title. Read-only.
func (m *WorkbookChart) SetTitle(value *WorkbookChartTitle)() {
    if m != nil {
        m.title = value
    }
}
// SetTop sets the top property value. Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart).
func (m *WorkbookChart) SetTop(value *float64)() {
    if m != nil {
        m.top = value
    }
}
// SetWidth sets the width property value. Represents the width, in points, of the chart object.
func (m *WorkbookChart) SetWidth(value *float64)() {
    if m != nil {
        m.width = value
    }
}
// SetWorksheet sets the worksheet property value. The worksheet containing the current chart. Read-only.
func (m *WorkbookChart) SetWorksheet(value *WorkbookWorksheet)() {
    if m != nil {
        m.worksheet = value
    }
}
