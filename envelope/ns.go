package envelope

const (
	NS_ADDRESSING  = "http://schemas.xmlsoap.org/ws/2004/08/addressing"        // a:
	NS_CIMBINDING  = "http://schemas.dmtf.org/wbem/wsman/1/cimbinding.xsd"     // b:
	NS_ENUM        = "http://schemas.xmlsoap.org/ws/2004/09/enumeration"       // n:
	NS_SCHEMA      = "http://www.w3.org/2001/XMLSchema"                        // xsd:
	NS_SCHEMA_INST = "http://www.w3.org/2001/XMLSchema-instance"               // xsi:
	NS_SOAP_ENV    = "http://www.w3.org/2003/05/soap-envelope"                 // s:
	NS_TRANSFER    = "http://schemas.xmlsoap.org/ws/2004/09/transfer"          // x:
	NS_WIN_SHELL   = "http://schemas.microsoft.com/wbem/wsman/1/windows/shell" // rsp:
	NS_WSMAN_CONF  = "http://schemas.microsoft.com/wbem/wsman/1/config"        // cfg:
	NS_WSMAN_DMTF  = "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"          // w:
	NS_WSMAN_MSFT  = "http://schemas.microsoft.com/wbem/wsman/1/wsman.xsd"     // p:
	NS_WSMAN_FAULT = "http://schemas.microsoft.com/wbem/wsman/1/wsmanfault"    // f:
	NS_EVENTING    = "http://schemas.xmlsoap.org/ws/2004/08/eventing"

	NS_WMI       = "http://schemas.microsoft.com/wbem/wsman/1/wmi"
	NS_WMI_CIMV2 = "http://schemas.microsoft.com/wbem/wsman/1/wmi/root/cimv2"
	NS_CIMV2     = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2"
	NS_WINDOWS   = "http://schemas.microsoft.com/wbem/wsman/1/windows"
)

// wmi      http://schemas.microsoft.com/wbem/wsman/1/wmi
// wmicimv2 http://schemas.microsoft.com/wbem/wsman/1/wmi/root/cimv2
// cimv2    http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2
// winrm    http://schemas.microsoft.com/wbem/wsman/1
// wsman    http://schemas.microsoft.com/wbem/wsman/1
// shell    http://schemas.microsoft.com/wbem/wsman/1/windows/shell
