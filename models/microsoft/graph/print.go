package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Print struct {
    additionalData map[string]interface{};
    connectors []PrintConnector;
    operations []PrintOperation;
    printers []Printer;
    services []PrintService;
    settings *PrintSettings;
    shares []PrinterShare;
    taskDefinitions []PrintTaskDefinition;
}
func NewPrint()(*Print) {
    m := &Print{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Print) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Print) GetConnectors()([]PrintConnector) {
    if m == nil {
        return nil
    } else {
        return m.connectors
    }
}
func (m *Print) GetOperations()([]PrintOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
func (m *Print) GetPrinters()([]Printer) {
    if m == nil {
        return nil
    } else {
        return m.printers
    }
}
func (m *Print) GetServices()([]PrintService) {
    if m == nil {
        return nil
    } else {
        return m.services
    }
}
func (m *Print) GetSettings()(*PrintSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *Print) GetShares()([]PrinterShare) {
    if m == nil {
        return nil
    } else {
        return m.shares
    }
}
func (m *Print) GetTaskDefinitions()([]PrintTaskDefinition) {
    if m == nil {
        return nil
    } else {
        return m.taskDefinitions
    }
}
func (m *Print) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["connectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintConnector() })
        if err != nil {
            return err
        }
        res := make([]PrintConnector, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintConnector))
        }
        m.SetConnectors(res)
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintOperation() })
        if err != nil {
            return err
        }
        res := make([]PrintOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintOperation))
        }
        m.SetOperations(res)
        return nil
    }
    res["printers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinter() })
        if err != nil {
            return err
        }
        res := make([]Printer, len(val))
        for i, v := range val {
            res[i] = *(v.(*Printer))
        }
        m.SetPrinters(res)
        return nil
    }
    res["services"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintService() })
        if err != nil {
            return err
        }
        res := make([]PrintService, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintService))
        }
        m.SetServices(res)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*PrintSettings))
        return nil
    }
    res["shares"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterShare() })
        if err != nil {
            return err
        }
        res := make([]PrinterShare, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrinterShare))
        }
        m.SetShares(res)
        return nil
    }
    res["taskDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskDefinition() })
        if err != nil {
            return err
        }
        res := make([]PrintTaskDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintTaskDefinition))
        }
        m.SetTaskDefinitions(res)
        return nil
    }
    return res
}
func (m *Print) IsNil()(bool) {
    return m == nil
}
func (m *Print) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConnectors()))
        for i, v := range m.GetConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("connectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPrinters()))
        for i, v := range m.GetPrinters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("printers", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetServices()))
        for i, v := range m.GetServices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("services", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetShares()))
        for i, v := range m.GetShares() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("shares", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTaskDefinitions()))
        for i, v := range m.GetTaskDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("taskDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Print) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Print) SetConnectors(value []PrintConnector)() {
    m.connectors = value
}
func (m *Print) SetOperations(value []PrintOperation)() {
    m.operations = value
}
func (m *Print) SetPrinters(value []Printer)() {
    m.printers = value
}
func (m *Print) SetServices(value []PrintService)() {
    m.services = value
}
func (m *Print) SetSettings(value *PrintSettings)() {
    m.settings = value
}
func (m *Print) SetShares(value []PrinterShare)() {
    m.shares = value
}
func (m *Print) SetTaskDefinitions(value []PrintTaskDefinition)() {
    m.taskDefinitions = value
}
