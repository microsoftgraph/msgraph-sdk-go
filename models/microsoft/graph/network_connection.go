package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type NetworkConnection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the application managing the network connection (for example, Facebook or SMTP).
    applicationName *string;
    // Destination IP address (of the network connection).
    destinationAddress *string;
    // Destination domain portion of the destination URL. (for example 'www.contoso.com').
    destinationDomain *string;
    // Location (by IP address mapping) associated with the destination of a network connection.
    destinationLocation *string;
    // Destination port (of the network connection).
    destinationPort *string;
    // Network connection URL/URI string - excluding parameters. (for example 'www.contoso.com/products/default.html')
    destinationUrl *string;
    // Network connection direction. Possible values are: unknown, inbound, outbound.
    direction *ConnectionDirection;
    // Date when the destination domain was registered. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    domainRegisteredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The local DNS name resolution as it appears in the host's local DNS cache (for example, in case the 'hosts' file was tampered with).
    localDnsName *string;
    // Network Address Translation destination IP address.
    natDestinationAddress *string;
    // Network Address Translation destination port.
    natDestinationPort *string;
    // Network Address Translation source IP address.
    natSourceAddress *string;
    // Network Address Translation source port.
    natSourcePort *string;
    // Network protocol. Possible values are: unknown, ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII.
    protocol *SecurityNetworkProtocol;
    // Provider generated/calculated risk score of the network connection. Recommended value range of 0-1, which equates to a percentage.
    riskScore *string;
    // Source (i.e. origin) IP address (of the network connection).
    sourceAddress *string;
    // Location (by IP address mapping) associated with the source of a network connection.
    sourceLocation *string;
    // Source (i.e. origin) IP port (of the network connection).
    sourcePort *string;
    // Network connection status. Possible values are: unknown, attempted, succeeded, blocked, failed.
    status *ConnectionStatus;
    // Parameters (suffix) of the destination URL.
    urlParameters *string;
}
// Instantiates a new networkConnection and sets the default values.
func NewNetworkConnection()(*NetworkConnection) {
    m := &NetworkConnection{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkConnection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the applicationName property value. Name of the application managing the network connection (for example, Facebook or SMTP).
func (m *NetworkConnection) GetApplicationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationName
    }
}
// Gets the destinationAddress property value. Destination IP address (of the network connection).
func (m *NetworkConnection) GetDestinationAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationAddress
    }
}
// Gets the destinationDomain property value. Destination domain portion of the destination URL. (for example 'www.contoso.com').
func (m *NetworkConnection) GetDestinationDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationDomain
    }
}
// Gets the destinationLocation property value. Location (by IP address mapping) associated with the destination of a network connection.
func (m *NetworkConnection) GetDestinationLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationLocation
    }
}
// Gets the destinationPort property value. Destination port (of the network connection).
func (m *NetworkConnection) GetDestinationPort()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationPort
    }
}
// Gets the destinationUrl property value. Network connection URL/URI string - excluding parameters. (for example 'www.contoso.com/products/default.html')
func (m *NetworkConnection) GetDestinationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationUrl
    }
}
// Gets the direction property value. Network connection direction. Possible values are: unknown, inbound, outbound.
func (m *NetworkConnection) GetDirection()(*ConnectionDirection) {
    if m == nil {
        return nil
    } else {
        return m.direction
    }
}
// Gets the domainRegisteredDateTime property value. Date when the destination domain was registered. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *NetworkConnection) GetDomainRegisteredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.domainRegisteredDateTime
    }
}
// Gets the localDnsName property value. The local DNS name resolution as it appears in the host's local DNS cache (for example, in case the 'hosts' file was tampered with).
func (m *NetworkConnection) GetLocalDnsName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.localDnsName
    }
}
// Gets the natDestinationAddress property value. Network Address Translation destination IP address.
func (m *NetworkConnection) GetNatDestinationAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.natDestinationAddress
    }
}
// Gets the natDestinationPort property value. Network Address Translation destination port.
func (m *NetworkConnection) GetNatDestinationPort()(*string) {
    if m == nil {
        return nil
    } else {
        return m.natDestinationPort
    }
}
// Gets the natSourceAddress property value. Network Address Translation source IP address.
func (m *NetworkConnection) GetNatSourceAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.natSourceAddress
    }
}
// Gets the natSourcePort property value. Network Address Translation source port.
func (m *NetworkConnection) GetNatSourcePort()(*string) {
    if m == nil {
        return nil
    } else {
        return m.natSourcePort
    }
}
// Gets the protocol property value. Network protocol. Possible values are: unknown, ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII.
func (m *NetworkConnection) GetProtocol()(*SecurityNetworkProtocol) {
    if m == nil {
        return nil
    } else {
        return m.protocol
    }
}
// Gets the riskScore property value. Provider generated/calculated risk score of the network connection. Recommended value range of 0-1, which equates to a percentage.
func (m *NetworkConnection) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
// Gets the sourceAddress property value. Source (i.e. origin) IP address (of the network connection).
func (m *NetworkConnection) GetSourceAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceAddress
    }
}
// Gets the sourceLocation property value. Location (by IP address mapping) associated with the source of a network connection.
func (m *NetworkConnection) GetSourceLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceLocation
    }
}
// Gets the sourcePort property value. Source (i.e. origin) IP port (of the network connection).
func (m *NetworkConnection) GetSourcePort()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourcePort
    }
}
// Gets the status property value. Network connection status. Possible values are: unknown, attempted, succeeded, blocked, failed.
func (m *NetworkConnection) GetStatus()(*ConnectionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the urlParameters property value. Parameters (suffix) of the destination URL.
func (m *NetworkConnection) GetUrlParameters()(*string) {
    if m == nil {
        return nil
    } else {
        return m.urlParameters
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *NetworkConnection) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the applicationName property value. Name of the application managing the network connection (for example, Facebook or SMTP).
// Parameters:
//  - value : Value to set for the applicationName property.
func (m *NetworkConnection) SetApplicationName(value *string)() {
    m.applicationName = value
}
// Sets the destinationAddress property value. Destination IP address (of the network connection).
// Parameters:
//  - value : Value to set for the destinationAddress property.
func (m *NetworkConnection) SetDestinationAddress(value *string)() {
    m.destinationAddress = value
}
// Sets the destinationDomain property value. Destination domain portion of the destination URL. (for example 'www.contoso.com').
// Parameters:
//  - value : Value to set for the destinationDomain property.
func (m *NetworkConnection) SetDestinationDomain(value *string)() {
    m.destinationDomain = value
}
// Sets the destinationLocation property value. Location (by IP address mapping) associated with the destination of a network connection.
// Parameters:
//  - value : Value to set for the destinationLocation property.
func (m *NetworkConnection) SetDestinationLocation(value *string)() {
    m.destinationLocation = value
}
// Sets the destinationPort property value. Destination port (of the network connection).
// Parameters:
//  - value : Value to set for the destinationPort property.
func (m *NetworkConnection) SetDestinationPort(value *string)() {
    m.destinationPort = value
}
// Sets the destinationUrl property value. Network connection URL/URI string - excluding parameters. (for example 'www.contoso.com/products/default.html')
// Parameters:
//  - value : Value to set for the destinationUrl property.
func (m *NetworkConnection) SetDestinationUrl(value *string)() {
    m.destinationUrl = value
}
// Sets the direction property value. Network connection direction. Possible values are: unknown, inbound, outbound.
// Parameters:
//  - value : Value to set for the direction property.
func (m *NetworkConnection) SetDirection(value *ConnectionDirection)() {
    m.direction = value
}
// Sets the domainRegisteredDateTime property value. Date when the destination domain was registered. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the domainRegisteredDateTime property.
func (m *NetworkConnection) SetDomainRegisteredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.domainRegisteredDateTime = value
}
// Sets the localDnsName property value. The local DNS name resolution as it appears in the host's local DNS cache (for example, in case the 'hosts' file was tampered with).
// Parameters:
//  - value : Value to set for the localDnsName property.
func (m *NetworkConnection) SetLocalDnsName(value *string)() {
    m.localDnsName = value
}
// Sets the natDestinationAddress property value. Network Address Translation destination IP address.
// Parameters:
//  - value : Value to set for the natDestinationAddress property.
func (m *NetworkConnection) SetNatDestinationAddress(value *string)() {
    m.natDestinationAddress = value
}
// Sets the natDestinationPort property value. Network Address Translation destination port.
// Parameters:
//  - value : Value to set for the natDestinationPort property.
func (m *NetworkConnection) SetNatDestinationPort(value *string)() {
    m.natDestinationPort = value
}
// Sets the natSourceAddress property value. Network Address Translation source IP address.
// Parameters:
//  - value : Value to set for the natSourceAddress property.
func (m *NetworkConnection) SetNatSourceAddress(value *string)() {
    m.natSourceAddress = value
}
// Sets the natSourcePort property value. Network Address Translation source port.
// Parameters:
//  - value : Value to set for the natSourcePort property.
func (m *NetworkConnection) SetNatSourcePort(value *string)() {
    m.natSourcePort = value
}
// Sets the protocol property value. Network protocol. Possible values are: unknown, ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6, ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6, ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII.
// Parameters:
//  - value : Value to set for the protocol property.
func (m *NetworkConnection) SetProtocol(value *SecurityNetworkProtocol)() {
    m.protocol = value
}
// Sets the riskScore property value. Provider generated/calculated risk score of the network connection. Recommended value range of 0-1, which equates to a percentage.
// Parameters:
//  - value : Value to set for the riskScore property.
func (m *NetworkConnection) SetRiskScore(value *string)() {
    m.riskScore = value
}
// Sets the sourceAddress property value. Source (i.e. origin) IP address (of the network connection).
// Parameters:
//  - value : Value to set for the sourceAddress property.
func (m *NetworkConnection) SetSourceAddress(value *string)() {
    m.sourceAddress = value
}
// Sets the sourceLocation property value. Location (by IP address mapping) associated with the source of a network connection.
// Parameters:
//  - value : Value to set for the sourceLocation property.
func (m *NetworkConnection) SetSourceLocation(value *string)() {
    m.sourceLocation = value
}
// Sets the sourcePort property value. Source (i.e. origin) IP port (of the network connection).
// Parameters:
//  - value : Value to set for the sourcePort property.
func (m *NetworkConnection) SetSourcePort(value *string)() {
    m.sourcePort = value
}
// Sets the status property value. Network connection status. Possible values are: unknown, attempted, succeeded, blocked, failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *NetworkConnection) SetStatus(value *ConnectionStatus)() {
    m.status = value
}
// Sets the urlParameters property value. Parameters (suffix) of the destination URL.
// Parameters:
//  - value : Value to set for the urlParameters property.
func (m *NetworkConnection) SetUrlParameters(value *string)() {
    m.urlParameters = value
}
