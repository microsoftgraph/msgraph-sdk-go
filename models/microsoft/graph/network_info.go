package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/callrecords"
)

// 
type NetworkInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
    bandwidthLowEventRatio *float32;
    // The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
    basicServiceSetIdentifier *string;
    // Type of network used by the media endpoint. Possible values are: unknown, wired, wifi, mobile, tunnel, unknownFutureValue.
    connectionType *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType;
    // Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
    delayEventRatio *float32;
    // DNS suffix associated with the network adapter of the media endpoint.
    dnsSuffix *string;
    // IP address of the media endpoint.
    ipAddress *string;
    // Link speed in bits per second reported by the network adapter used by the media endpoint.
    linkSpeed *int64;
    // The media access control (MAC) address of the media endpoint's network device.
    macAddress *string;
    // Network port number used by media endpoint.
    port *int32;
    // Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
    receivedQualityEventRatio *float32;
    // IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
    reflexiveIPAddress *string;
    // IP address of the media relay server allocated by the media endpoint.
    relayIPAddress *string;
    // Network port number allocated on the media relay server by the media endpoint.
    relayPort *int32;
    // Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
    sentQualityEventRatio *float32;
    // Subnet used for media stream by the media endpoint.
    subnet *string;
    // WiFi band used by the media endpoint. Possible values are: unknown, frequency24GHz, frequency50GHz, frequency60GHz, unknownFutureValue.
    wifiBand *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand;
    // Estimated remaining battery charge in percentage reported by the media endpoint.
    wifiBatteryCharge *int32;
    // WiFi channel used by the media endpoint.
    wifiChannel *int32;
    // Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
    wifiMicrosoftDriver *string;
    // Version of the Microsoft WiFi driver used by the media endpoint.
    wifiMicrosoftDriverVersion *string;
    // Type of WiFi radio used by the media endpoint. Possible values are: unknown, wifi80211a, wifi80211b, wifi80211g, wifi80211n, wifi80211ac, wifi80211ax, unknownFutureValue.
    wifiRadioType *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType;
    // WiFi signal strength in percentage reported by the media endpoint.
    wifiSignalStrength *int32;
    // Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
    wifiVendorDriver *string;
    // Version of the WiFi driver used by the media endpoint.
    wifiVendorDriverVersion *string;
}
// Instantiates a new networkInfo and sets the default values.
func NewNetworkInfo()(*NetworkInfo) {
    m := &NetworkInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the bandwidthLowEventRatio property value. Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
func (m *NetworkInfo) GetBandwidthLowEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.bandwidthLowEventRatio
    }
}
// Gets the basicServiceSetIdentifier property value. The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
func (m *NetworkInfo) GetBasicServiceSetIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.basicServiceSetIdentifier
    }
}
// Gets the connectionType property value. Type of network used by the media endpoint. Possible values are: unknown, wired, wifi, mobile, tunnel, unknownFutureValue.
func (m *NetworkInfo) GetConnectionType()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType) {
    if m == nil {
        return nil
    } else {
        return m.connectionType
    }
}
// Gets the delayEventRatio property value. Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
func (m *NetworkInfo) GetDelayEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.delayEventRatio
    }
}
// Gets the dnsSuffix property value. DNS suffix associated with the network adapter of the media endpoint.
func (m *NetworkInfo) GetDnsSuffix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dnsSuffix
    }
}
// Gets the ipAddress property value. IP address of the media endpoint.
func (m *NetworkInfo) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// Gets the linkSpeed property value. Link speed in bits per second reported by the network adapter used by the media endpoint.
func (m *NetworkInfo) GetLinkSpeed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.linkSpeed
    }
}
// Gets the macAddress property value. The media access control (MAC) address of the media endpoint's network device.
func (m *NetworkInfo) GetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macAddress
    }
}
// Gets the port property value. Network port number used by media endpoint.
func (m *NetworkInfo) GetPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.port
    }
}
// Gets the receivedQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
func (m *NetworkInfo) GetReceivedQualityEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.receivedQualityEventRatio
    }
}
// Gets the reflexiveIPAddress property value. IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
func (m *NetworkInfo) GetReflexiveIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reflexiveIPAddress
    }
}
// Gets the relayIPAddress property value. IP address of the media relay server allocated by the media endpoint.
func (m *NetworkInfo) GetRelayIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relayIPAddress
    }
}
// Gets the relayPort property value. Network port number allocated on the media relay server by the media endpoint.
func (m *NetworkInfo) GetRelayPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.relayPort
    }
}
// Gets the sentQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
func (m *NetworkInfo) GetSentQualityEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.sentQualityEventRatio
    }
}
// Gets the subnet property value. Subnet used for media stream by the media endpoint.
func (m *NetworkInfo) GetSubnet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnet
    }
}
// Gets the wifiBand property value. WiFi band used by the media endpoint. Possible values are: unknown, frequency24GHz, frequency50GHz, frequency60GHz, unknownFutureValue.
func (m *NetworkInfo) GetWifiBand()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand) {
    if m == nil {
        return nil
    } else {
        return m.wifiBand
    }
}
// Gets the wifiBatteryCharge property value. Estimated remaining battery charge in percentage reported by the media endpoint.
func (m *NetworkInfo) GetWifiBatteryCharge()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiBatteryCharge
    }
}
// Gets the wifiChannel property value. WiFi channel used by the media endpoint.
func (m *NetworkInfo) GetWifiChannel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiChannel
    }
}
// Gets the wifiMicrosoftDriver property value. Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) GetWifiMicrosoftDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMicrosoftDriver
    }
}
// Gets the wifiMicrosoftDriverVersion property value. Version of the Microsoft WiFi driver used by the media endpoint.
func (m *NetworkInfo) GetWifiMicrosoftDriverVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMicrosoftDriverVersion
    }
}
// Gets the wifiRadioType property value. Type of WiFi radio used by the media endpoint. Possible values are: unknown, wifi80211a, wifi80211b, wifi80211g, wifi80211n, wifi80211ac, wifi80211ax, unknownFutureValue.
func (m *NetworkInfo) GetWifiRadioType()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType) {
    if m == nil {
        return nil
    } else {
        return m.wifiRadioType
    }
}
// Gets the wifiSignalStrength property value. WiFi signal strength in percentage reported by the media endpoint.
func (m *NetworkInfo) GetWifiSignalStrength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiSignalStrength
    }
}
// Gets the wifiVendorDriver property value. Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) GetWifiVendorDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiVendorDriver
    }
}
// Gets the wifiVendorDriverVersion property value. Version of the WiFi driver used by the media endpoint.
func (m *NetworkInfo) GetWifiVendorDriverVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiVendorDriverVersion
    }
}
// The deserialization information for the current model
func (m *NetworkInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bandwidthLowEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetBandwidthLowEventRatio(val)
        return nil
    }
    res["basicServiceSetIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBasicServiceSetIdentifier(val)
        return nil
    }
    res["connectionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseNetworkConnectionType)
        if err != nil {
            return err
        }
        cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType)
        m.SetConnectionType(&cast)
        return nil
    }
    res["delayEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetDelayEventRatio(val)
        return nil
    }
    res["dnsSuffix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDnsSuffix(val)
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpAddress(val)
        return nil
    }
    res["linkSpeed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLinkSpeed(val)
        return nil
    }
    res["macAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMacAddress(val)
        return nil
    }
    res["port"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPort(val)
        return nil
    }
    res["receivedQualityEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetReceivedQualityEventRatio(val)
        return nil
    }
    res["reflexiveIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReflexiveIPAddress(val)
        return nil
    }
    res["relayIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRelayIPAddress(val)
        return nil
    }
    res["relayPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRelayPort(val)
        return nil
    }
    res["sentQualityEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetSentQualityEventRatio(val)
        return nil
    }
    res["subnet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubnet(val)
        return nil
    }
    res["wifiBand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseWifiBand)
        if err != nil {
            return err
        }
        cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand)
        m.SetWifiBand(&cast)
        return nil
    }
    res["wifiBatteryCharge"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWifiBatteryCharge(val)
        return nil
    }
    res["wifiChannel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWifiChannel(val)
        return nil
    }
    res["wifiMicrosoftDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiMicrosoftDriver(val)
        return nil
    }
    res["wifiMicrosoftDriverVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiMicrosoftDriverVersion(val)
        return nil
    }
    res["wifiRadioType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseWifiRadioType)
        if err != nil {
            return err
        }
        cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType)
        m.SetWifiRadioType(&cast)
        return nil
    }
    res["wifiSignalStrength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWifiSignalStrength(val)
        return nil
    }
    res["wifiVendorDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiVendorDriver(val)
        return nil
    }
    res["wifiVendorDriverVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiVendorDriverVersion(val)
        return nil
    }
    return res
}
func (m *NetworkInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *NetworkInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat32Value("bandwidthLowEventRatio", m.GetBandwidthLowEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("basicServiceSetIdentifier", m.GetBasicServiceSetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := m.GetConnectionType().String()
        err := writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("delayEventRatio", m.GetDelayEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dnsSuffix", m.GetDnsSuffix())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("linkSpeed", m.GetLinkSpeed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("receivedQualityEventRatio", m.GetReceivedQualityEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reflexiveIPAddress", m.GetReflexiveIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("relayIPAddress", m.GetRelayIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("relayPort", m.GetRelayPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("sentQualityEventRatio", m.GetSentQualityEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnet", m.GetSubnet())
        if err != nil {
            return err
        }
    }
    if m.GetWifiBand() != nil {
        cast := m.GetWifiBand().String()
        err := writer.WriteStringValue("wifiBand", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiBatteryCharge", m.GetWifiBatteryCharge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiChannel", m.GetWifiChannel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMicrosoftDriver", m.GetWifiMicrosoftDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMicrosoftDriverVersion", m.GetWifiMicrosoftDriverVersion())
        if err != nil {
            return err
        }
    }
    if m.GetWifiRadioType() != nil {
        cast := m.GetWifiRadioType().String()
        err := writer.WriteStringValue("wifiRadioType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiSignalStrength", m.GetWifiSignalStrength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiVendorDriver", m.GetWifiVendorDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiVendorDriverVersion", m.GetWifiVendorDriverVersion())
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
func (m *NetworkInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the bandwidthLowEventRatio property value. Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
// Parameters:
//  - value : Value to set for the bandwidthLowEventRatio property.
func (m *NetworkInfo) SetBandwidthLowEventRatio(value *float32)() {
    m.bandwidthLowEventRatio = value
}
// Sets the basicServiceSetIdentifier property value. The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
// Parameters:
//  - value : Value to set for the basicServiceSetIdentifier property.
func (m *NetworkInfo) SetBasicServiceSetIdentifier(value *string)() {
    m.basicServiceSetIdentifier = value
}
// Sets the connectionType property value. Type of network used by the media endpoint. Possible values are: unknown, wired, wifi, mobile, tunnel, unknownFutureValue.
// Parameters:
//  - value : Value to set for the connectionType property.
func (m *NetworkInfo) SetConnectionType(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType)() {
    m.connectionType = value
}
// Sets the delayEventRatio property value. Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
// Parameters:
//  - value : Value to set for the delayEventRatio property.
func (m *NetworkInfo) SetDelayEventRatio(value *float32)() {
    m.delayEventRatio = value
}
// Sets the dnsSuffix property value. DNS suffix associated with the network adapter of the media endpoint.
// Parameters:
//  - value : Value to set for the dnsSuffix property.
func (m *NetworkInfo) SetDnsSuffix(value *string)() {
    m.dnsSuffix = value
}
// Sets the ipAddress property value. IP address of the media endpoint.
// Parameters:
//  - value : Value to set for the ipAddress property.
func (m *NetworkInfo) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// Sets the linkSpeed property value. Link speed in bits per second reported by the network adapter used by the media endpoint.
// Parameters:
//  - value : Value to set for the linkSpeed property.
func (m *NetworkInfo) SetLinkSpeed(value *int64)() {
    m.linkSpeed = value
}
// Sets the macAddress property value. The media access control (MAC) address of the media endpoint's network device.
// Parameters:
//  - value : Value to set for the macAddress property.
func (m *NetworkInfo) SetMacAddress(value *string)() {
    m.macAddress = value
}
// Sets the port property value. Network port number used by media endpoint.
// Parameters:
//  - value : Value to set for the port property.
func (m *NetworkInfo) SetPort(value *int32)() {
    m.port = value
}
// Sets the receivedQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
// Parameters:
//  - value : Value to set for the receivedQualityEventRatio property.
func (m *NetworkInfo) SetReceivedQualityEventRatio(value *float32)() {
    m.receivedQualityEventRatio = value
}
// Sets the reflexiveIPAddress property value. IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
// Parameters:
//  - value : Value to set for the reflexiveIPAddress property.
func (m *NetworkInfo) SetReflexiveIPAddress(value *string)() {
    m.reflexiveIPAddress = value
}
// Sets the relayIPAddress property value. IP address of the media relay server allocated by the media endpoint.
// Parameters:
//  - value : Value to set for the relayIPAddress property.
func (m *NetworkInfo) SetRelayIPAddress(value *string)() {
    m.relayIPAddress = value
}
// Sets the relayPort property value. Network port number allocated on the media relay server by the media endpoint.
// Parameters:
//  - value : Value to set for the relayPort property.
func (m *NetworkInfo) SetRelayPort(value *int32)() {
    m.relayPort = value
}
// Sets the sentQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
// Parameters:
//  - value : Value to set for the sentQualityEventRatio property.
func (m *NetworkInfo) SetSentQualityEventRatio(value *float32)() {
    m.sentQualityEventRatio = value
}
// Sets the subnet property value. Subnet used for media stream by the media endpoint.
// Parameters:
//  - value : Value to set for the subnet property.
func (m *NetworkInfo) SetSubnet(value *string)() {
    m.subnet = value
}
// Sets the wifiBand property value. WiFi band used by the media endpoint. Possible values are: unknown, frequency24GHz, frequency50GHz, frequency60GHz, unknownFutureValue.
// Parameters:
//  - value : Value to set for the wifiBand property.
func (m *NetworkInfo) SetWifiBand(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand)() {
    m.wifiBand = value
}
// Sets the wifiBatteryCharge property value. Estimated remaining battery charge in percentage reported by the media endpoint.
// Parameters:
//  - value : Value to set for the wifiBatteryCharge property.
func (m *NetworkInfo) SetWifiBatteryCharge(value *int32)() {
    m.wifiBatteryCharge = value
}
// Sets the wifiChannel property value. WiFi channel used by the media endpoint.
// Parameters:
//  - value : Value to set for the wifiChannel property.
func (m *NetworkInfo) SetWifiChannel(value *int32)() {
    m.wifiChannel = value
}
// Sets the wifiMicrosoftDriver property value. Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
// Parameters:
//  - value : Value to set for the wifiMicrosoftDriver property.
func (m *NetworkInfo) SetWifiMicrosoftDriver(value *string)() {
    m.wifiMicrosoftDriver = value
}
// Sets the wifiMicrosoftDriverVersion property value. Version of the Microsoft WiFi driver used by the media endpoint.
// Parameters:
//  - value : Value to set for the wifiMicrosoftDriverVersion property.
func (m *NetworkInfo) SetWifiMicrosoftDriverVersion(value *string)() {
    m.wifiMicrosoftDriverVersion = value
}
// Sets the wifiRadioType property value. Type of WiFi radio used by the media endpoint. Possible values are: unknown, wifi80211a, wifi80211b, wifi80211g, wifi80211n, wifi80211ac, wifi80211ax, unknownFutureValue.
// Parameters:
//  - value : Value to set for the wifiRadioType property.
func (m *NetworkInfo) SetWifiRadioType(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType)() {
    m.wifiRadioType = value
}
// Sets the wifiSignalStrength property value. WiFi signal strength in percentage reported by the media endpoint.
// Parameters:
//  - value : Value to set for the wifiSignalStrength property.
func (m *NetworkInfo) SetWifiSignalStrength(value *int32)() {
    m.wifiSignalStrength = value
}
// Sets the wifiVendorDriver property value. Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
// Parameters:
//  - value : Value to set for the wifiVendorDriver property.
func (m *NetworkInfo) SetWifiVendorDriver(value *string)() {
    m.wifiVendorDriver = value
}
// Sets the wifiVendorDriverVersion property value. Version of the WiFi driver used by the media endpoint.
// Parameters:
//  - value : Value to set for the wifiVendorDriverVersion property.
func (m *NetworkInfo) SetWifiVendorDriverVersion(value *string)() {
    m.wifiVendorDriverVersion = value
}
