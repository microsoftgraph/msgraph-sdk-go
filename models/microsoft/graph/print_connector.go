package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// printConnector 
type PrintConnector struct {
    Entity
    // The connector's version.
    appVersion *string;
    // The name of the connector.
    displayName *string;
    // The connector machine's hostname.
    fullyQualifiedDomainName *string;
    // The physical and/or organizational location of the connector.
    location *PrinterLocation;
    // The connector machine's operating system version.
    operatingSystem *string;
    // The DateTimeOffset when the connector was registered.
    registeredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewPrintConnector instantiates a new printConnector and sets the default values.
func NewPrintConnector()(*PrintConnector) {
    m := &PrintConnector{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppVersion gets the appVersion property value. The connector's version.
func (m *PrintConnector) GetAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appVersion
    }
}
// GetDisplayName gets the displayName property value. The name of the connector.
func (m *PrintConnector) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFullyQualifiedDomainName gets the fullyQualifiedDomainName property value. The connector machine's hostname.
func (m *PrintConnector) GetFullyQualifiedDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fullyQualifiedDomainName
    }
}
// GetLocation gets the location property value. The physical and/or organizational location of the connector.
func (m *PrintConnector) GetLocation()(*PrinterLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// GetOperatingSystem gets the operatingSystem property value. The connector machine's operating system version.
func (m *PrintConnector) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// GetRegisteredDateTime gets the registeredDateTime property value. The DateTimeOffset when the connector was registered.
func (m *PrintConnector) GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registeredDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppVersion(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["fullyQualifiedDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullyQualifiedDomainName(val)
        }
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterLocation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(*PrinterLocation))
        }
        return nil
    }
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["registeredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegisteredDateTime(val)
        }
        return nil
    }
    return res
}
func (m *PrintConnector) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrintConnector) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appVersion", m.GetAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fullyQualifiedDomainName", m.GetFullyQualifiedDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registeredDateTime", m.GetRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppVersion sets the appVersion property value. The connector's version.
func (m *PrintConnector) SetAppVersion(value *string)() {
    m.appVersion = value
}
// SetDisplayName sets the displayName property value. The name of the connector.
func (m *PrintConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFullyQualifiedDomainName sets the fullyQualifiedDomainName property value. The connector machine's hostname.
func (m *PrintConnector) SetFullyQualifiedDomainName(value *string)() {
    m.fullyQualifiedDomainName = value
}
// SetLocation sets the location property value. The physical and/or organizational location of the connector.
func (m *PrintConnector) SetLocation(value *PrinterLocation)() {
    m.location = value
}
// SetOperatingSystem sets the operatingSystem property value. The connector machine's operating system version.
func (m *PrintConnector) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetRegisteredDateTime sets the registeredDateTime property value. The DateTimeOffset when the connector was registered.
func (m *PrintConnector) SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registeredDateTime = value
}
