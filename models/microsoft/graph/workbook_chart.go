package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new workbookChart and sets the default values.
func NewWorkbookChart()(*WorkbookChart) {
    m := &WorkbookChart{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the axes property value. Represents chart axes. Read-only.
func (m *WorkbookChart) GetAxes()(*WorkbookChartAxes) {
    if m == nil {
        return nil
    } else {
        return m.axes
    }
}
// Gets the dataLabels property value. Represents the datalabels on the chart. Read-only.
func (m *WorkbookChart) GetDataLabels()(*WorkbookChartDataLabels) {
    if m == nil {
        return nil
    } else {
        return m.dataLabels
    }
}
// Gets the format property value. Encapsulates the format properties for the chart area. Read-only.
func (m *WorkbookChart) GetFormat()(*WorkbookChartAreaFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// Gets the height property value. Represents the height, in points, of the chart object.
func (m *WorkbookChart) GetHeight()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// Gets the left property value. The distance, in points, from the left side of the chart to the worksheet origin.
func (m *WorkbookChart) GetLeft()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.left
    }
}
// Gets the legend property value. Represents the legend for the chart. Read-only.
func (m *WorkbookChart) GetLegend()(*WorkbookChartLegend) {
    if m == nil {
        return nil
    } else {
        return m.legend
    }
}
// Gets the name property value. Represents the name of a chart object.
func (m *WorkbookChart) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the series property value. Represents either a single series or collection of series in the chart. Read-only.
func (m *WorkbookChart) GetSeries()([]WorkbookChartSeries) {
    if m == nil {
        return nil
    } else {
        return m.series
    }
}
// Gets the title property value. Represents the title of the specified chart, including the text, visibility, position and formating of the title. Read-only.
func (m *WorkbookChart) GetTitle()(*WorkbookChartTitle) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Gets the top property value. Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart).
func (m *WorkbookChart) GetTop()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// Gets the width property value. Represents the width, in points, of the chart object.
func (m *WorkbookChart) GetWidth()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
// Gets the worksheet property value. The worksheet containing the current chart. Read-only.
func (m *WorkbookChart) GetWorksheet()(*WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
// The deserialization information for the current model
func (m *WorkbookChart) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["axes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAxes() })
        if err != nil {
            return err
        }
        m.SetAxes(val.(*WorkbookChartAxes))
        return nil
    }
    res["dataLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartDataLabels() })
        if err != nil {
            return err
        }
        m.SetDataLabels(val.(*WorkbookChartDataLabels))
        return nil
    }
    res["format"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartAreaFormat() })
        if err != nil {
            return err
        }
        m.SetFormat(val.(*WorkbookChartAreaFormat))
        return nil
    }
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetHeight(val)
        return nil
    }
    res["left"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetLeft(val)
        return nil
    }
    res["legend"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartLegend() })
        if err != nil {
            return err
        }
        m.SetLegend(val.(*WorkbookChartLegend))
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
    res["series"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartSeries() })
        if err != nil {
            return err
        }
        res := make([]WorkbookChartSeries, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkbookChartSeries))
        }
        m.SetSeries(res)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookChartTitle() })
        if err != nil {
            return err
        }
        m.SetTitle(val.(*WorkbookChartTitle))
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetTop(val)
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetWidth(val)
        return nil
    }
    res["worksheet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkbookWorksheet() })
        if err != nil {
            return err
        }
        m.SetWorksheet(val.(*WorkbookWorksheet))
        return nil
    }
    return res
}
func (m *WorkbookChart) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
// Sets the axes property value. Represents chart axes. Read-only.
// Parameters:
//  - value : Value to set for the axes property.
func (m *WorkbookChart) SetAxes(value *WorkbookChartAxes)() {
    m.axes = value
}
// Sets the dataLabels property value. Represents the datalabels on the chart. Read-only.
// Parameters:
//  - value : Value to set for the dataLabels property.
func (m *WorkbookChart) SetDataLabels(value *WorkbookChartDataLabels)() {
    m.dataLabels = value
}
// Sets the format property value. Encapsulates the format properties for the chart area. Read-only.
// Parameters:
//  - value : Value to set for the format property.
func (m *WorkbookChart) SetFormat(value *WorkbookChartAreaFormat)() {
    m.format = value
}
// Sets the height property value. Represents the height, in points, of the chart object.
// Parameters:
//  - value : Value to set for the height property.
func (m *WorkbookChart) SetHeight(value *float64)() {
    m.height = value
}
// Sets the left property value. The distance, in points, from the left side of the chart to the worksheet origin.
// Parameters:
//  - value : Value to set for the left property.
func (m *WorkbookChart) SetLeft(value *float64)() {
    m.left = value
}
// Sets the legend property value. Represents the legend for the chart. Read-only.
// Parameters:
//  - value : Value to set for the legend property.
func (m *WorkbookChart) SetLegend(value *WorkbookChartLegend)() {
    m.legend = value
}
// Sets the name property value. Represents the name of a chart object.
// Parameters:
//  - value : Value to set for the name property.
func (m *WorkbookChart) SetName(value *string)() {
    m.name = value
}
// Sets the series property value. Represents either a single series or collection of series in the chart. Read-only.
// Parameters:
//  - value : Value to set for the series property.
func (m *WorkbookChart) SetSeries(value []WorkbookChartSeries)() {
    m.series = value
}
// Sets the title property value. Represents the title of the specified chart, including the text, visibility, position and formating of the title. Read-only.
// Parameters:
//  - value : Value to set for the title property.
func (m *WorkbookChart) SetTitle(value *WorkbookChartTitle)() {
    m.title = value
}
// Sets the top property value. Represents the distance, in points, from the top edge of the object to the top of row 1 (on a worksheet) or the top of the chart area (on a chart).
// Parameters:
//  - value : Value to set for the top property.
func (m *WorkbookChart) SetTop(value *float64)() {
    m.top = value
}
// Sets the width property value. Represents the width, in points, of the chart object.
// Parameters:
//  - value : Value to set for the width property.
func (m *WorkbookChart) SetWidth(value *float64)() {
    m.width = value
}
// Sets the worksheet property value. The worksheet containing the current chart. Read-only.
// Parameters:
//  - value : Value to set for the worksheet property.
func (m *WorkbookChart) SetWorksheet(value *WorkbookWorksheet)() {
    m.worksheet = value
}
