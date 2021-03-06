// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AirPrintDestination undocumented
type AirPrintDestination struct {
	// Object is the base model of AirPrintDestination
	Object
	// IPAddress The IP Address of the AirPrint destination.
	IPAddress *string `json:"ipAddress,omitempty"`
	// ResourcePath undocumented
	ResourcePath *string `json:"resourcePath,omitempty"`
	// Port The listening port of the AirPrint destination. If this key is not specified AirPrint will use the default port. Available in iOS 11.0 and later.
	Port *int `json:"port,omitempty"`
	// ForceTLS If true AirPrint connections are secured by Transport Layer Security (TLS). Default is false. Available in iOS 11.0 and later.
	ForceTLS *bool `json:"forceTls,omitempty"`
}
