package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintUsageByPrinter struct {
    PrintUsage
    printerId *string;
}
func NewPrintUsageByPrinter()(*PrintUsageByPrinter) {
    m := &PrintUsageByPrinter{
        PrintUsage: *NewPrintUsage(),
    }
    return m
}
func (m *PrintUsageByPrinter) GetPrinterId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.printerId
    }
}
func (m *PrintUsageByPrinter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PrintUsage.GetFieldDeserializers()
    res["printerId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrinterId(val)
        return nil
    }
    return res
}
func (m *PrintUsageByPrinter) IsNil()(bool) {
    return m == nil
}
func (m *PrintUsageByPrinter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PrintUsage.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("printerId", m.GetPrinterId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PrintUsageByPrinter) SetPrinterId(value *string)() {
    m.printerId = value
}
