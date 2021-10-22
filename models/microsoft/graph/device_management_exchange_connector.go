package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementExchangeConnector struct {
    Entity
    connectorServerName *string;
    exchangeAlias *string;
    exchangeConnectorType *DeviceManagementExchangeConnectorType;
    exchangeOrganization *string;
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    primarySmtpAddress *string;
    serverName *string;
    status *DeviceManagementExchangeConnectorStatus;
    version *string;
}
func NewDeviceManagementExchangeConnector()(*DeviceManagementExchangeConnector) {
    m := &DeviceManagementExchangeConnector{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementExchangeConnector) GetConnectorServerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorServerName
    }
}
func (m *DeviceManagementExchangeConnector) GetExchangeAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeAlias
    }
}
func (m *DeviceManagementExchangeConnector) GetExchangeConnectorType()(*DeviceManagementExchangeConnectorType) {
    if m == nil {
        return nil
    } else {
        return m.exchangeConnectorType
    }
}
func (m *DeviceManagementExchangeConnector) GetExchangeOrganization()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeOrganization
    }
}
func (m *DeviceManagementExchangeConnector) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
func (m *DeviceManagementExchangeConnector) GetPrimarySmtpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.primarySmtpAddress
    }
}
func (m *DeviceManagementExchangeConnector) GetServerName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serverName
    }
}
func (m *DeviceManagementExchangeConnector) GetStatus()(*DeviceManagementExchangeConnectorStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *DeviceManagementExchangeConnector) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
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
func (m *DeviceManagementExchangeConnector) SetConnectorServerName(value *string)() {
    m.connectorServerName = value
}
func (m *DeviceManagementExchangeConnector) SetExchangeAlias(value *string)() {
    m.exchangeAlias = value
}
func (m *DeviceManagementExchangeConnector) SetExchangeConnectorType(value *DeviceManagementExchangeConnectorType)() {
    m.exchangeConnectorType = value
}
func (m *DeviceManagementExchangeConnector) SetExchangeOrganization(value *string)() {
    m.exchangeOrganization = value
}
func (m *DeviceManagementExchangeConnector) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
func (m *DeviceManagementExchangeConnector) SetPrimarySmtpAddress(value *string)() {
    m.primarySmtpAddress = value
}
func (m *DeviceManagementExchangeConnector) SetServerName(value *string)() {
    m.serverName = value
}
func (m *DeviceManagementExchangeConnector) SetStatus(value *DeviceManagementExchangeConnectorStatus)() {
    m.status = value
}
func (m *DeviceManagementExchangeConnector) SetVersion(value *string)() {
    m.version = value
}
