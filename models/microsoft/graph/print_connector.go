package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new printConnector and sets the default values.
func NewPrintConnector()(*PrintConnector) {
    m := &PrintConnector{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appVersion property value. The connector's version.
func (m *PrintConnector) GetAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appVersion
    }
}
// Gets the displayName property value. The name of the connector.
func (m *PrintConnector) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the fullyQualifiedDomainName property value. The connector machine's hostname.
func (m *PrintConnector) GetFullyQualifiedDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fullyQualifiedDomainName
    }
}
// Gets the location property value. The physical and/or organizational location of the connector.
func (m *PrintConnector) GetLocation()(*PrinterLocation) {
    if m == nil {
        return nil
    } else {
        return m.location
    }
}
// Gets the operatingSystem property value. The connector machine's operating system version.
func (m *PrintConnector) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// Gets the registeredDateTime property value. The DateTimeOffset when the connector was registered.
func (m *PrintConnector) GetRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registeredDateTime
    }
}
// The deserialization information for the current model
func (m *PrintConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppVersion(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["fullyQualifiedDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFullyQualifiedDomainName(val)
        return nil
    }
    res["location"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrinterLocation() })
        if err != nil {
            return err
        }
        m.SetLocation(val.(*PrinterLocation))
        return nil
    }
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystem(val)
        return nil
    }
    res["registeredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRegisteredDateTime(val)
        return nil
    }
    return res
}
func (m *PrintConnector) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the appVersion property value. The connector's version.
// Parameters:
//  - value : Value to set for the appVersion property.
func (m *PrintConnector) SetAppVersion(value *string)() {
    m.appVersion = value
}
// Sets the displayName property value. The name of the connector.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PrintConnector) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the fullyQualifiedDomainName property value. The connector machine's hostname.
// Parameters:
//  - value : Value to set for the fullyQualifiedDomainName property.
func (m *PrintConnector) SetFullyQualifiedDomainName(value *string)() {
    m.fullyQualifiedDomainName = value
}
// Sets the location property value. The physical and/or organizational location of the connector.
// Parameters:
//  - value : Value to set for the location property.
func (m *PrintConnector) SetLocation(value *PrinterLocation)() {
    m.location = value
}
// Sets the operatingSystem property value. The connector machine's operating system version.
// Parameters:
//  - value : Value to set for the operatingSystem property.
func (m *PrintConnector) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// Sets the registeredDateTime property value. The DateTimeOffset when the connector was registered.
// Parameters:
//  - value : Value to set for the registeredDateTime property.
func (m *PrintConnector) SetRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registeredDateTime = value
}
