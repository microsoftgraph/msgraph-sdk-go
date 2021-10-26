package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementExchangeConnector struct {
    Entity
    // The name of the server hosting the Exchange Connector.
    connectorServerName *string;
    // An alias assigned to the Exchange server
    exchangeAlias *string;
    // The type of Exchange Connector Configured. Possible values are: onPremises, hosted, serviceToService, dedicated.
    exchangeConnectorType *DeviceManagementExchangeConnectorType;
    // Exchange Organization to the Exchange server
    exchangeOrganization *string;
    // Last sync time for the Exchange Connector
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Email address used to configure the Service To Service Exchange Connector.
    primarySmtpAddress *string;
    // The name of the Exchange server.
    serverName *string;
    // Exchange Connector Status. Possible values are: none, connectionPending, connected, disconnected.
    status *DeviceManagementExchangeConnectorStatus;
    // The version of the ExchangeConnectorAgent
    version *string;
}
// Instantiates a new deviceManagementExchangeConnector and sets the default values.
func NewDeviceManagementExchangeConnector()(*DeviceManagementExchangeConnector) {
    m := &DeviceManagementExchangeConnector{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the connectorServerName property value. The name of the server hosting the Exchange Connector.
func (m *DeviceManagementExchangeConnector) GetConnectorServerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorServerName
    }
}
// Gets the exchangeAlias property value. An alias assigned to the Exchange server
func (m *DeviceManagementExchangeConnector) GetExchangeAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAlias
    }
}
// Gets the exchangeConnectorType property value. The type of Exchange Connector Configured. Possible values are: onPremises, hosted, serviceToService, dedicated.
func (m *DeviceManagementExchangeConnector) GetExchangeConnectorType()(*DeviceManagementExchangeConnectorType) {
    if m == nil {
        return nil
    } else {
        return m.exchangeConnectorType
    }
}
// Gets the exchangeOrganization property value. Exchange Organization to the Exchange server
func (m *DeviceManagementExchangeConnector) GetExchangeOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeOrganization
    }
}
// Gets the lastSyncDateTime property value. Last sync time for the Exchange Connector
func (m *DeviceManagementExchangeConnector) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// Gets the primarySmtpAddress property value. Email address used to configure the Service To Service Exchange Connector.
func (m *DeviceManagementExchangeConnector) GetPrimarySmtpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.primarySmtpAddress
    }
}
// Gets the serverName property value. The name of the Exchange server.
func (m *DeviceManagementExchangeConnector) GetServerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serverName
    }
}
// Gets the status property value. Exchange Connector Status. Possible values are: none, connectionPending, connected, disconnected.
func (m *DeviceManagementExchangeConnector) GetStatus()(*DeviceManagementExchangeConnectorStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the version property value. The version of the ExchangeConnectorAgent
func (m *DeviceManagementExchangeConnector) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceManagementExchangeConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectorServerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConnectorServerName(val)
        return nil
    }
    res["exchangeAlias"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExchangeAlias(val)
        return nil
    }
    res["exchangeConnectorType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeConnectorType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExchangeConnectorType)
        m.SetExchangeConnectorType(&cast)
        return nil
    }
    res["exchangeOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExchangeOrganization(val)
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSyncDateTime(val)
        return nil
    }
    res["primarySmtpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPrimarySmtpAddress(val)
        return nil
    }
    res["serverName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServerName(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeConnectorStatus)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementExchangeConnectorStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *DeviceManagementExchangeConnector) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementExchangeConnector) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("connectorServerName", m.GetConnectorServerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("exchangeAlias", m.GetExchangeAlias())
        if err != nil {
            return err
        }
    }
    if m.GetExchangeConnectorType() != nil {
        cast := m.GetExchangeConnectorType().String()
        err = writer.WriteStringValue("exchangeConnectorType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("exchangeOrganization", m.GetExchangeOrganization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primarySmtpAddress", m.GetPrimarySmtpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serverName", m.GetServerName())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the connectorServerName property value. The name of the server hosting the Exchange Connector.
// Parameters:
//  - value : Value to set for the connectorServerName property.
func (m *DeviceManagementExchangeConnector) SetConnectorServerName(value *string)() {
    m.connectorServerName = value
}
// Sets the exchangeAlias property value. An alias assigned to the Exchange server
// Parameters:
//  - value : Value to set for the exchangeAlias property.
func (m *DeviceManagementExchangeConnector) SetExchangeAlias(value *string)() {
    m.exchangeAlias = value
}
// Sets the exchangeConnectorType property value. The type of Exchange Connector Configured. Possible values are: onPremises, hosted, serviceToService, dedicated.
// Parameters:
//  - value : Value to set for the exchangeConnectorType property.
func (m *DeviceManagementExchangeConnector) SetExchangeConnectorType(value *DeviceManagementExchangeConnectorType)() {
    m.exchangeConnectorType = value
}
// Sets the exchangeOrganization property value. Exchange Organization to the Exchange server
// Parameters:
//  - value : Value to set for the exchangeOrganization property.
func (m *DeviceManagementExchangeConnector) SetExchangeOrganization(value *string)() {
    m.exchangeOrganization = value
}
// Sets the lastSyncDateTime property value. Last sync time for the Exchange Connector
// Parameters:
//  - value : Value to set for the lastSyncDateTime property.
func (m *DeviceManagementExchangeConnector) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// Sets the primarySmtpAddress property value. Email address used to configure the Service To Service Exchange Connector.
// Parameters:
//  - value : Value to set for the primarySmtpAddress property.
func (m *DeviceManagementExchangeConnector) SetPrimarySmtpAddress(value *string)() {
    m.primarySmtpAddress = value
}
// Sets the serverName property value. The name of the Exchange server.
// Parameters:
//  - value : Value to set for the serverName property.
func (m *DeviceManagementExchangeConnector) SetServerName(value *string)() {
    m.serverName = value
}
// Sets the status property value. Exchange Connector Status. Possible values are: none, connectionPending, connected, disconnected.
// Parameters:
//  - value : Value to set for the status property.
func (m *DeviceManagementExchangeConnector) SetStatus(value *DeviceManagementExchangeConnectorStatus)() {
    m.status = value
}
// Sets the version property value. The version of the ExchangeConnectorAgent
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceManagementExchangeConnector) SetVersion(value *string)() {
    m.version = value
}
