package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type WorkbookChart struct {
    Entity
    axes *WorkbookChartAxes;
    dataLabels *WorkbookChartDataLabels;
    format *WorkbookChartAreaFormat;
    height *float64;
    left *float64;
    legend *WorkbookChartLegend;
    name *string;
    series []WorkbookChartSeries;
    title *WorkbookChartTitle;
    top *float64;
    width *float64;
    worksheet *WorkbookWorksheet;
}
func NewWorkbookChart()(*WorkbookChart) {
    m := &WorkbookChart{
        Entity: *NewEntity(),
    }
    return m
}
func (m *WorkbookChart) GetAxes()(*WorkbookChartAxes) {
    if m == nil {
        return nil
    } else {
        return m.axes
    }
}
func (m *WorkbookChart) GetDataLabels()(*WorkbookChartDataLabels) {
    if m == nil {
        return nil
    } else {
        return m.dataLabels
    }
}
func (m *WorkbookChart) GetFormat()(*WorkbookChartAreaFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
func (m *WorkbookChart) GetHeight()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
func (m *WorkbookChart) GetLeft()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.left
    }
}
func (m *WorkbookChart) GetLegend()(*WorkbookChartLegend) {
    if m == nil {
        return nil
    } else {
        return m.legend
    }
}
func (m *WorkbookChart) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *WorkbookChart) GetSeries()([]WorkbookChartSeries) {
    if m == nil {
        return nil
    } else {
        return m.series
    }
}
func (m *WorkbookChart) GetTitle()(*WorkbookChartTitle) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
func (m *WorkbookChart) GetTop()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
func (m *WorkbookChart) GetWidth()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
func (m *WorkbookChart) GetWorksheet()(*WorkbookWorksheet) {
    if m == nil {
        return nil
    } else {
        return m.worksheet
    }
}
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
func (m *WorkbookChart) SetAxes(value *WorkbookChartAxes)() {
    m.axes = value
}
func (m *WorkbookChart) SetDataLabels(value *WorkbookChartDataLabels)() {
    m.dataLabels = value
}
func (m *WorkbookChart) SetFormat(value *WorkbookChartAreaFormat)() {
    m.format = value
}
func (m *WorkbookChart) SetHeight(value *float64)() {
    m.height = value
}
func (m *WorkbookChart) SetLeft(value *float64)() {
    m.left = value
}
func (m *WorkbookChart) SetLegend(value *WorkbookChartLegend)() {
    m.legend = value
}
func (m *WorkbookChart) SetName(value *string)() {
    m.name = value
}
func (m *WorkbookChart) SetSeries(value []WorkbookChartSeries)() {
    m.series = value
}
func (m *WorkbookChart) SetTitle(value *WorkbookChartTitle)() {
    m.title = value
}
func (m *WorkbookChart) SetTop(value *float64)() {
    m.top = value
}
func (m *WorkbookChart) SetWidth(value *float64)() {
    m.width = value
}
func (m *WorkbookChart) SetWorksheet(value *WorkbookWorksheet)() {
    m.worksheet = value
}
