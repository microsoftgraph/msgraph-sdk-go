package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type NetworkConnection struct {
    additionalData map[string]interface{};
    applicationName *string;
    destinationAddress *string;
    destinationDomain *string;
    destinationLocation *string;
    destinationPort *string;
    destinationUrl *string;
    direction *ConnectionDirection;
    domainRegisteredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    localDnsName *string;
    natDestinationAddress *string;
    natDestinationPort *string;
    natSourceAddress *string;
    natSourcePort *string;
    protocol *SecurityNetworkProtocol;
    riskScore *string;
    sourceAddress *string;
    sourceLocation *string;
    sourcePort *string;
    status *ConnectionStatus;
    urlParameters *string;
}
func NewNetworkConnection()(*NetworkConnection) {
    m := &NetworkConnection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *NetworkConnection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *NetworkConnection) GetApplicationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationName
    }
}
func (m *NetworkConnection) GetDestinationAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationAddress
    }
}
func (m *NetworkConnection) GetDestinationDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationDomain
    }
}
func (m *NetworkConnection) GetDestinationLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationLocation
    }
}
func (m *NetworkConnection) GetDestinationPort()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationPort
    }
}
func (m *NetworkConnection) GetDestinationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationUrl
    }
}
func (m *NetworkConnection) GetDirection()(*ConnectionDirection) {
    if m == nil {
        return nil
    } else {
        return m.direction
    }
}
func (m *NetworkConnection) GetDomainRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.domainRegisteredDateTime
    }
}
func (m *NetworkConnection) GetLocalDnsName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localDnsName
    }
}
func (m *NetworkConnection) GetNatDestinationAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.natDestinationAddress
    }
}
func (m *NetworkConnection) GetNatDestinationPort()(*string) {
    if m == nil {
        return nil
    } else {
        return m.natDestinationPort
    }
}
func (m *NetworkConnection) GetNatSourceAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.natSourceAddress
    }
}
func (m *NetworkConnection) GetNatSourcePort()(*string) {
    if m == nil {
        return nil
    } else {
        return m.natSourcePort
    }
}
func (m *NetworkConnection) GetProtocol()(*SecurityNetworkProtocol) {
    if m == nil {
        return nil
    } else {
        return m.protocol
    }
}
func (m *NetworkConnection) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
func (m *NetworkConnection) GetSourceAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceAddress
    }
}
func (m *NetworkConnection) GetSourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceLocation
    }
}
func (m *NetworkConnection) GetSourcePort()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourcePort
    }
}
func (m *NetworkConnection) GetStatus()(*ConnectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *NetworkConnection) GetUrlParameters()(*string) {
    if m == nil {
        return nil
    } else {
        return m.urlParameters
    }
}
func (m *NetworkConnection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationName(val)
        return nil
    }
    res["destinationAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationAddress(val)
        return nil
    }
    res["destinationDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationDomain(val)
        return nil
    }
    res["destinationLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationLocation(val)
        return nil
    }
    res["destinationPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationPort(val)
        return nil
    }
    res["destinationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDestinationUrl(val)
        return nil
    }
    res["direction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionDirection)
        if err != nil {
            return err
        }
        cast := val.(ConnectionDirection)
        m.SetDirection(&cast)
        return nil
    }
    res["domainRegisteredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDomainRegisteredDateTime(val)
        return nil
    }
    res["localDnsName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocalDnsName(val)
        return nil
    }
    res["natDestinationAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNatDestinationAddress(val)
        return nil
    }
    res["natDestinationPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNatDestinationPort(val)
        return nil
    }
    res["natSourceAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNatSourceAddress(val)
        return nil
    }
    res["natSourcePort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNatSourcePort(val)
        return nil
    }
    res["protocol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityNetworkProtocol)
        if err != nil {
            return err
        }
        cast := val.(SecurityNetworkProtocol)
        m.SetProtocol(&cast)
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskScore(val)
        return nil
    }
    res["sourceAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceAddress(val)
        return nil
    }
    res["sourceLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceLocation(val)
        return nil
    }
    res["sourcePort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourcePort(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectionStatus)
        if err != nil {
            return err
        }
        cast := val.(ConnectionStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["urlParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrlParameters(val)
        return nil
    }
    return res
}
func (m *NetworkConnection) IsNil()(bool) {
    return m == nil
}
func (m *NetworkConnection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationName", m.GetApplicationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationAddress", m.GetDestinationAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationDomain", m.GetDestinationDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationLocation", m.GetDestinationLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationPort", m.GetDestinationPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationUrl", m.GetDestinationUrl())
        if err != nil {
            return err
        }
    }
    if m.GetDirection() != nil {
        cast := m.GetDirection().String()
        err := writer.WriteStringValue("direction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("domainRegisteredDateTime", m.GetDomainRegisteredDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("localDnsName", m.GetLocalDnsName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("natDestinationAddress", m.GetNatDestinationAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("natDestinationPort", m.GetNatDestinationPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("natSourceAddress", m.GetNatSourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("natSourcePort", m.GetNatSourcePort())
        if err != nil {
            return err
        }
    }
    if m.GetProtocol() != nil {
        cast := m.GetProtocol().String()
        err := writer.WriteStringValue("protocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceAddress", m.GetSourceAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceLocation", m.GetSourceLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourcePort", m.GetSourcePort())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("urlParameters", m.GetUrlParameters())
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
func (m *NetworkConnection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *NetworkConnection) SetApplicationName(value *string)() {
    m.applicationName = value
}
func (m *NetworkConnection) SetDestinationAddress(value *string)() {
    m.destinationAddress = value
}
func (m *NetworkConnection) SetDestinationDomain(value *string)() {
    m.destinationDomain = value
}
func (m *NetworkConnection) SetDestinationLocation(value *string)() {
    m.destinationLocation = value
}
func (m *NetworkConnection) SetDestinationPort(value *string)() {
    m.destinationPort = value
}
func (m *NetworkConnection) SetDestinationUrl(value *string)() {
    m.destinationUrl = value
}
func (m *NetworkConnection) SetDirection(value *ConnectionDirection)() {
    m.direction = value
}
func (m *NetworkConnection) SetDomainRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.domainRegisteredDateTime = value
}
func (m *NetworkConnection) SetLocalDnsName(value *string)() {
    m.localDnsName = value
}
func (m *NetworkConnection) SetNatDestinationAddress(value *string)() {
    m.natDestinationAddress = value
}
func (m *NetworkConnection) SetNatDestinationPort(value *string)() {
    m.natDestinationPort = value
}
func (m *NetworkConnection) SetNatSourceAddress(value *string)() {
    m.natSourceAddress = value
}
func (m *NetworkConnection) SetNatSourcePort(value *string)() {
    m.natSourcePort = value
}
func (m *NetworkConnection) SetProtocol(value *SecurityNetworkProtocol)() {
    m.protocol = value
}
func (m *NetworkConnection) SetRiskScore(value *string)() {
    m.riskScore = value
}
func (m *NetworkConnection) SetSourceAddress(value *string)() {
    m.sourceAddress = value
}
func (m *NetworkConnection) SetSourceLocation(value *string)() {
    m.sourceLocation = value
}
func (m *NetworkConnection) SetSourcePort(value *string)() {
    m.sourcePort = value
}
func (m *NetworkConnection) SetStatus(value *ConnectionStatus)() {
    m.status = value
}
func (m *NetworkConnection) SetUrlParameters(value *string)() {
    m.urlParameters = value
}
