package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Print 
type Print struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The list of available print connectors.
    connectors []PrintConnector;
    // The list of print long running operations.
    operations []PrintOperation;
    // The list of printers registered in the tenant.
    printers []Printer;
    // The list of available Universal Print service endpoints.
    services []PrintService;
    // Tenant-wide settings for the Universal Print service.
    settings *PrintSettings;
    // The list of printer shares registered in the tenant.
    shares []PrinterShare;
    // List of abstract definition for a task that can be triggered when various events occur within Universal Print.
    taskDefinitions []PrintTaskDefinition;
}
// NewPrint instantiates a new Print and sets the default values.
func NewPrint()(*Print) {
    m := &Print{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Print) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConnectors gets the connectors property value. The list of available print connectors.
func (m *Print) GetConnectors()([]PrintConnector) {
    if m == nil {
        return nil
    } else {
        return m.connectors
    }
}
// GetOperations gets the operations property value. The list of print long running operations.
func (m *Print) GetOperations()([]PrintOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPrinters gets the printers property value. The list of printers registered in the tenant.
func (m *Print) GetPrinters()([]Printer) {
    if m == nil {
        return nil
    } else {
        return m.printers
    }
}
// GetServices gets the services property value. The list of available Universal Print service endpoints.
func (m *Print) GetServices()([]PrintService) {
    if m == nil {
        return nil
    } else {
        return m.services
    }
}
// GetSettings gets the settings property value. Tenant-wide settings for the Universal Print service.
func (m *Print) GetSettings()(*PrintSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetShares gets the shares property value. The list of printer shares registered in the tenant.
func (m *Print) GetShares()([]PrinterShare) {
    if m == nil {
        return nil
    } else {
        return m.shares
    }
}
// GetTaskDefinitions gets the taskDefinitions property value. List of abstract definition for a task that can be triggered when various events occur within Universal Print.
func (m *Print) GetTaskDefinitions()([]PrintTaskDefinition) {
    if m == nil {
        return nil
    } else {
        return m.taskDefinitions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Print) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["connectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintConnector() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintConnector, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintConnector))
            }
            m.SetConnectors(res)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["printers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Printer, len(val))
            for i, v := range val {
                res[i] = *(v.(*Printer))
            }
            m.SetPrinters(res)
        }
        return nil
    }
    res["services"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintService() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintService, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintService))
            }
            m.SetServices(res)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*PrintSettings))
        }
        return nil
    }
    res["shares"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterShare() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrinterShare, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrinterShare))
            }
            m.SetShares(res)
        }
        return nil
    }
    res["taskDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTaskDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTaskDefinition, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintTaskDefinition))
            }
            m.SetTaskDefinitions(res)
        }
        return nil
    }
    return res
}
func (m *Print) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Print) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetConnectors sets the connectors property value. The list of available print connectors.
func (m *Print) SetConnectors(value []PrintConnector)() {
    m.connectors = value
}
// SetOperations sets the operations property value. The list of print long running operations.
func (m *Print) SetOperations(value []PrintOperation)() {
    m.operations = value
}
// SetPrinters sets the printers property value. The list of printers registered in the tenant.
func (m *Print) SetPrinters(value []Printer)() {
    m.printers = value
}
// SetServices sets the services property value. The list of available Universal Print service endpoints.
func (m *Print) SetServices(value []PrintService)() {
    m.services = value
}
// SetSettings sets the settings property value. Tenant-wide settings for the Universal Print service.
func (m *Print) SetSettings(value *PrintSettings)() {
    m.settings = value
}
// SetShares sets the shares property value. The list of printer shares registered in the tenant.
func (m *Print) SetShares(value []PrinterShare)() {
    m.shares = value
}
// SetTaskDefinitions sets the taskDefinitions property value. List of abstract definition for a task that can be triggered when various events occur within Universal Print.
func (m *Print) SetTaskDefinitions(value []PrintTaskDefinition)() {
    m.taskDefinitions = value
}
