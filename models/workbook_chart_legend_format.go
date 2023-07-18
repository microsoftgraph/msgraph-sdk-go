package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookChartLegendFormat 
type WorkbookChartLegendFormat struct {
    Entity
}
// NewWorkbookChartLegendFormat instantiates a new workbookChartLegendFormat and sets the default values.
func NewWorkbookChartLegendFormat()(*WorkbookChartLegendFormat) {
    m := &WorkbookChartLegendFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookChartLegendFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookChartLegendFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookChartLegendFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookChartLegendFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFillFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFill(val.(WorkbookChartFillable))
        }
        return nil
    }
    res["font"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookChartFontFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFont(val.(WorkbookChartFontable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetFill gets the fill property value. Represents the fill format of an object, which includes background formating information. Read-only.
func (m *WorkbookChartLegendFormat) GetFill()(WorkbookChartFillable) {
    val, err := m.GetBackingStore().Get("fill")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartFillable)
    }
    return nil
}
// GetFont gets the font property value. Represents the font attributes such as font name, font size, color, etc. of a chart legend. Read-only.
func (m *WorkbookChartLegendFormat) GetFont()(WorkbookChartFontable) {
    val, err := m.GetBackingStore().Get("font")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WorkbookChartFontable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkbookChartLegendFormat) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WorkbookChartLegendFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFill sets the fill property value. Represents the fill format of an object, which includes background formating information. Read-only.
func (m *WorkbookChartLegendFormat) SetFill(value WorkbookChartFillable)() {
    err := m.GetBackingStore().Set("fill", value)
    if err != nil {
        panic(err)
    }
}
// SetFont sets the font property value. Represents the font attributes such as font name, font size, color, etc. of a chart legend. Read-only.
func (m *WorkbookChartLegendFormat) SetFont(value WorkbookChartFontable)() {
    err := m.GetBackingStore().Set("font", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookChartLegendFormat) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// WorkbookChartLegendFormatable 
type WorkbookChartLegendFormatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFill()(WorkbookChartFillable)
    GetFont()(WorkbookChartFontable)
    GetOdataType()(*string)
    SetFill(value WorkbookChartFillable)()
    SetFont(value WorkbookChartFontable)()
    SetOdataType(value *string)()
}
