package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementExchangeConnector provides operations to manage the deviceManagement singleton.
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
// NewDeviceManagementExchangeConnector instantiates a new deviceManagementExchangeConnector and sets the default values.
func NewDeviceManagementExchangeConnector()(*DeviceManagementExchangeConnector) {
    m := &DeviceManagementExchangeConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementExchangeConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementExchangeConnectorFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementExchangeConnector(), nil
}
// GetConnectorServerName gets the connectorServerName property value. The name of the server hosting the Exchange Connector.
func (m *DeviceManagementExchangeConnector) GetConnectorServerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorServerName
    }
}
// GetExchangeAlias gets the exchangeAlias property value. An alias assigned to the Exchange server
func (m *DeviceManagementExchangeConnector) GetExchangeAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAlias
    }
}
// GetExchangeConnectorType gets the exchangeConnectorType property value. The type of Exchange Connector Configured. Possible values are: onPremises, hosted, serviceToService, dedicated.
func (m *DeviceManagementExchangeConnector) GetExchangeConnectorType()(*DeviceManagementExchangeConnectorType) {
    if m == nil {
        return nil
    } else {
        return m.exchangeConnectorType
    }
}
// GetExchangeOrganization gets the exchangeOrganization property value. Exchange Organization to the Exchange server
func (m *DeviceManagementExchangeConnector) GetExchangeOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeOrganization
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeConnector) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectorServerName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorServerName(val)
        }
        return nil
    }
    res["exchangeAlias"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAlias(val)
        }
        return nil
    }
    res["exchangeConnectorType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeConnectorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeConnectorType(val.(*DeviceManagementExchangeConnectorType))
        }
        return nil
    }
    res["exchangeOrganization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeOrganization(val)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["primarySmtpAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimarySmtpAddress(val)
        }
        return nil
    }
    res["serverName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerName(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeConnectorStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceManagementExchangeConnectorStatus))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last sync time for the Exchange Connector
func (m *DeviceManagementExchangeConnector) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetPrimarySmtpAddress gets the primarySmtpAddress property value. Email address used to configure the Service To Service Exchange Connector.
func (m *DeviceManagementExchangeConnector) GetPrimarySmtpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.primarySmtpAddress
    }
}
// GetServerName gets the serverName property value. The name of the Exchange server.
func (m *DeviceManagementExchangeConnector) GetServerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serverName
    }
}
// GetStatus gets the status property value. Exchange Connector Status. Possible values are: none, connectionPending, connected, disconnected.
func (m *DeviceManagementExchangeConnector) GetStatus()(*DeviceManagementExchangeConnectorStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetVersion gets the version property value. The version of the ExchangeConnectorAgent
func (m *DeviceManagementExchangeConnector) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *DeviceManagementExchangeConnector) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetExchangeConnectorType()).String()
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
        cast := (*m.GetStatus()).String()
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
// SetConnectorServerName sets the connectorServerName property value. The name of the server hosting the Exchange Connector.
func (m *DeviceManagementExchangeConnector) SetConnectorServerName(value *string)() {
    if m != nil {
        m.connectorServerName = value
    }
}
// SetExchangeAlias sets the exchangeAlias property value. An alias assigned to the Exchange server
func (m *DeviceManagementExchangeConnector) SetExchangeAlias(value *string)() {
    if m != nil {
        m.exchangeAlias = value
    }
}
// SetExchangeConnectorType sets the exchangeConnectorType property value. The type of Exchange Connector Configured. Possible values are: onPremises, hosted, serviceToService, dedicated.
func (m *DeviceManagementExchangeConnector) SetExchangeConnectorType(value *DeviceManagementExchangeConnectorType)() {
    if m != nil {
        m.exchangeConnectorType = value
    }
}
// SetExchangeOrganization sets the exchangeOrganization property value. Exchange Organization to the Exchange server
func (m *DeviceManagementExchangeConnector) SetExchangeOrganization(value *string)() {
    if m != nil {
        m.exchangeOrganization = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last sync time for the Exchange Connector
func (m *DeviceManagementExchangeConnector) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetPrimarySmtpAddress sets the primarySmtpAddress property value. Email address used to configure the Service To Service Exchange Connector.
func (m *DeviceManagementExchangeConnector) SetPrimarySmtpAddress(value *string)() {
    if m != nil {
        m.primarySmtpAddress = value
    }
}
// SetServerName sets the serverName property value. The name of the Exchange server.
func (m *DeviceManagementExchangeConnector) SetServerName(value *string)() {
    if m != nil {
        m.serverName = value
    }
}
// SetStatus sets the status property value. Exchange Connector Status. Possible values are: none, connectionPending, connected, disconnected.
func (m *DeviceManagementExchangeConnector) SetStatus(value *DeviceManagementExchangeConnectorStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetVersion sets the version property value. The version of the ExchangeConnectorAgent
func (m *DeviceManagementExchangeConnector) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
