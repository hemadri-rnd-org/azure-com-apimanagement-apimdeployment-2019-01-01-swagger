package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ApimResource represents the ApimResource schema from the OpenAPI specification
type ApimResource struct {
	Id string `json:"id,omitempty"` // Resource ID.
	Name string `json:"name,omitempty"` // Resource name.
	Tags map[string]interface{} `json:"tags,omitempty"` // Resource tags.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource is set to Microsoft.ApiManagement.
}

// OperationListResult represents the OperationListResult schema from the OpenAPI specification
type OperationListResult struct {
	Nextlink string `json:"nextLink,omitempty"` // URL to get the next set of operation list results if there are any.
	Value []Operation `json:"value,omitempty"` // List of operations supported by the resource provider.
}

// ApiManagementServiceUpdateParameters represents the ApiManagementServiceUpdateParameters schema from the OpenAPI specification
type ApiManagementServiceUpdateParameters struct {
	Id string `json:"id,omitempty"` // Resource ID.
	Name string `json:"name,omitempty"` // Resource name.
	Tags map[string]interface{} `json:"tags,omitempty"` // Resource tags.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource is set to Microsoft.ApiManagement.
}

// ResourceSku represents the ResourceSku schema from the OpenAPI specification
type ResourceSku struct {
	Name string `json:"name,omitempty"` // Name of the Sku.
}

// ResourceSkuResult represents the ResourceSkuResult schema from the OpenAPI specification
type ResourceSkuResult struct {
	Sku ResourceSku `json:"sku,omitempty"` // Describes an available API Management SKU.
	Capacity ResourceSkuCapacity `json:"capacity,omitempty"` // Describes scaling information of a SKU.
	Resourcetype string `json:"resourceType,omitempty"` // The type of resource the SKU applies to.
}

// VirtualNetworkConfiguration represents the VirtualNetworkConfiguration schema from the OpenAPI specification
type VirtualNetworkConfiguration struct {
	Subnetname string `json:"subnetname,omitempty"` // The name of the subnet.
	Vnetid string `json:"vnetid,omitempty"` // The virtual network ID. This is typically a GUID. Expect a null GUID by default.
	Subnetresourceid string `json:"subnetResourceId,omitempty"` // The full resource ID of a subnet in a virtual network to deploy the API Management service in.
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	Properties map[string]interface{} `json:"properties,omitempty"` // The operation properties.
	Display interface{} `json:"display,omitempty"` // The object that describes the operation.
	Name string `json:"name,omitempty"` // Operation name: {provider}/{resource}/{operation}
	Origin string `json:"origin,omitempty"` // The operation origin.
}

// AdditionalLocation represents the AdditionalLocation schema from the OpenAPI specification
type AdditionalLocation struct {
	Sku ApiManagementServiceSkuProperties `json:"sku"` // API Management service resource SKU properties.
	Virtualnetworkconfiguration VirtualNetworkConfiguration `json:"virtualNetworkConfiguration,omitempty"` // Configuration of a virtual network to which API Management service is deployed.
	Gatewayregionalurl string `json:"gatewayRegionalUrl,omitempty"` // Gateway URL of the API Management service in the Region.
	Location string `json:"location"` // The location name of the additional region among Azure Data center regions.
	Privateipaddresses []string `json:"privateIPAddresses,omitempty"` // Private Static Load Balanced IP addresses of the API Management service which is deployed in an Internal Virtual Network in a particular additional location. Available only for Basic, Standard and Premium SKU.
	Publicipaddresses []string `json:"publicIPAddresses,omitempty"` // Public Static Load Balanced IP addresses of the API Management service in the additional location. Available only for Basic, Standard and Premium SKU.
}

// ApiManagementServiceApplyNetworkConfigurationParameters represents the ApiManagementServiceApplyNetworkConfigurationParameters schema from the OpenAPI specification
type ApiManagementServiceApplyNetworkConfigurationParameters struct {
	Location string `json:"location,omitempty"` // Location of the Api Management service to update for a multi-region service. For a service deployed in a single region, this parameter is not required.
}

// ApiManagementServiceNameAvailabilityResult represents the ApiManagementServiceNameAvailabilityResult schema from the OpenAPI specification
type ApiManagementServiceNameAvailabilityResult struct {
	Message string `json:"message,omitempty"` // If reason == invalid, provide the user with the reason why the given name is invalid, and provide the resource naming requirements so that the user can select a valid name. If reason == AlreadyExists, explain that <resourceName> is already in use, and direct them to select a different name.
	Nameavailable bool `json:"nameAvailable,omitempty"` // True if the name is available and can be used to create a new API Management service; otherwise false.
	Reason string `json:"reason,omitempty"` // Invalid indicates the name provided does not match the resource provider’s naming requirements (incorrect length, unsupported characters, etc.) AlreadyExists indicates that the name is already in use and is therefore unavailable.
}

// ApiManagementServiceGetSsoTokenResult represents the ApiManagementServiceGetSsoTokenResult schema from the OpenAPI specification
type ApiManagementServiceGetSsoTokenResult struct {
	Redirecturi string `json:"redirectUri,omitempty"` // Redirect URL to the Publisher Portal containing the SSO token.
}

// HostnameConfiguration represents the HostnameConfiguration schema from the OpenAPI specification
type HostnameConfiguration struct {
	Certificate CertificateInformation `json:"certificate,omitempty"` // SSL certificate information.
	Certificatepassword string `json:"certificatePassword,omitempty"` // Certificate Password.
	Defaultsslbinding bool `json:"defaultSslBinding,omitempty"` // Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate. If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The setting only applied to Proxy Hostname Type.
	Encodedcertificate string `json:"encodedCertificate,omitempty"` // Base64 Encoded certificate.
	Hostname string `json:"hostName"` // Hostname to configure on the Api Management service.
	Keyvaultid string `json:"keyVaultId,omitempty"` // Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided, auto-update of ssl certificate will not work. This requires Api Management service to be configured with MSI. The secret should be of type *application/x-pkcs12*
	Negotiateclientcertificate bool `json:"negotiateClientCertificate,omitempty"` // Specify true to always negotiate client certificate on the hostname. Default Value is false.
	TypeField string `json:"type"` // Hostname type.
}

// ApiManagementServiceListResult represents the ApiManagementServiceListResult schema from the OpenAPI specification
type ApiManagementServiceListResult struct {
	Value []ApiManagementServiceResource `json:"value"` // Result of the List API Management services operation.
	Nextlink string `json:"nextLink,omitempty"` // Link to the next set of results. Not empty if Value contains incomplete list of API Management services.
}

// ApiManagementServiceProperties represents the ApiManagementServiceProperties schema from the OpenAPI specification
type ApiManagementServiceProperties struct {
	Hostnameconfigurations []HostnameConfiguration `json:"hostnameConfigurations,omitempty"` // Custom hostname configuration of the API Management service.
	Gatewayregionalurl string `json:"gatewayRegionalUrl,omitempty"` // Gateway URL of the API Management service in the Default Region.
	Createdatutc string `json:"createdAtUtc,omitempty"` // Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Publicipaddresses []string `json:"publicIPAddresses,omitempty"` // Public Static Load Balanced IP addresses of the API Management service in Primary region. Available only for Basic, Standard and Premium SKU.
	Managementapiurl string `json:"managementApiUrl,omitempty"` // Management API endpoint URL of the API Management service.
	Portalurl string `json:"portalUrl,omitempty"` // Publisher portal endpoint Url of the API Management service.
	Notificationsenderemail string `json:"notificationSenderEmail,omitempty"` // Email address from which the notification will be sent.
	Virtualnetworkconfiguration VirtualNetworkConfiguration `json:"virtualNetworkConfiguration,omitempty"` // Configuration of a virtual network to which API Management service is deployed.
	Certificates []CertificateConfiguration `json:"certificates,omitempty"` // List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
	Enableclientcertificate bool `json:"enableClientCertificate,omitempty"` // Property only meant to be used for Consumption SKU Service. This enforces a client certificate to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the policy on the gateway.
	Privateipaddresses []string `json:"privateIPAddresses,omitempty"` // Private Static Load Balanced IP addresses of the API Management service in Primary region which is deployed in an Internal Virtual Network. Available only for Basic, Standard and Premium SKU.
	Provisioningstate string `json:"provisioningState,omitempty"` // The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	Customproperties map[string]interface{} `json:"customProperties,omitempty"` // Custom properties of the API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11` can be used to disable just TLS 1.1 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10` can be used to disable TLS 1.0 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2` can be used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH operation will reset omitted properties' values to their defaults. For all the settings except Http2 the default value is `True` if the service was created on or before April 1st 2018 and `False` otherwise. Http2 setting's default value is `False`.</br></br>You can disable any of next ciphers by using settings `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]`: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example, `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256`:`false`. The default value is `true` for them.
	Targetprovisioningstate string `json:"targetProvisioningState,omitempty"` // The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	Additionallocations []AdditionalLocation `json:"additionalLocations,omitempty"` // Additional datacenter locations of the API Management service.
	Gatewayurl string `json:"gatewayUrl,omitempty"` // Gateway URL of the API Management service.
	Scmurl string `json:"scmUrl,omitempty"` // SCM endpoint URL of the API Management service.
	Virtualnetworktype string `json:"virtualNetworkType,omitempty"` // The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
}

// ApiManagementServiceUpdateProperties represents the ApiManagementServiceUpdateProperties schema from the OpenAPI specification
type ApiManagementServiceUpdateProperties struct {
	Gatewayurl string `json:"gatewayUrl,omitempty"` // Gateway URL of the API Management service.
	Scmurl string `json:"scmUrl,omitempty"` // SCM endpoint URL of the API Management service.
	Virtualnetworktype string `json:"virtualNetworkType,omitempty"` // The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	Hostnameconfigurations []HostnameConfiguration `json:"hostnameConfigurations,omitempty"` // Custom hostname configuration of the API Management service.
	Gatewayregionalurl string `json:"gatewayRegionalUrl,omitempty"` // Gateway URL of the API Management service in the Default Region.
	Createdatutc string `json:"createdAtUtc,omitempty"` // Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Publicipaddresses []string `json:"publicIPAddresses,omitempty"` // Public Static Load Balanced IP addresses of the API Management service in Primary region. Available only for Basic, Standard and Premium SKU.
	Managementapiurl string `json:"managementApiUrl,omitempty"` // Management API endpoint URL of the API Management service.
	Portalurl string `json:"portalUrl,omitempty"` // Publisher portal endpoint Url of the API Management service.
	Notificationsenderemail string `json:"notificationSenderEmail,omitempty"` // Email address from which the notification will be sent.
	Virtualnetworkconfiguration VirtualNetworkConfiguration `json:"virtualNetworkConfiguration,omitempty"` // Configuration of a virtual network to which API Management service is deployed.
	Certificates []CertificateConfiguration `json:"certificates,omitempty"` // List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
	Enableclientcertificate bool `json:"enableClientCertificate,omitempty"` // Property only meant to be used for Consumption SKU Service. This enforces a client certificate to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the policy on the gateway.
	Privateipaddresses []string `json:"privateIPAddresses,omitempty"` // Private Static Load Balanced IP addresses of the API Management service in Primary region which is deployed in an Internal Virtual Network. Available only for Basic, Standard and Premium SKU.
	Provisioningstate string `json:"provisioningState,omitempty"` // The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	Customproperties map[string]interface{} `json:"customProperties,omitempty"` // Custom properties of the API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11` can be used to disable just TLS 1.1 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10` can be used to disable TLS 1.0 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2` can be used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH operation will reset omitted properties' values to their defaults. For all the settings except Http2 the default value is `True` if the service was created on or before April 1st 2018 and `False` otherwise. Http2 setting's default value is `False`.</br></br>You can disable any of next ciphers by using settings `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]`: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example, `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256`:`false`. The default value is `true` for them.
	Targetprovisioningstate string `json:"targetProvisioningState,omitempty"` // The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	Additionallocations []AdditionalLocation `json:"additionalLocations,omitempty"` // Additional datacenter locations of the API Management service.
}

// ApiManagementServiceSkuProperties represents the ApiManagementServiceSkuProperties schema from the OpenAPI specification
type ApiManagementServiceSkuProperties struct {
	Capacity int `json:"capacity,omitempty"` // Capacity of the SKU (number of deployed units of the SKU).
	Name string `json:"name"` // Name of the Sku.
}

// ApiManagementServiceCheckNameAvailabilityParameters represents the ApiManagementServiceCheckNameAvailabilityParameters schema from the OpenAPI specification
type ApiManagementServiceCheckNameAvailabilityParameters struct {
	Name string `json:"name"` // The name to check for availability.
}

// ResourceSkuCapacity represents the ResourceSkuCapacity schema from the OpenAPI specification
type ResourceSkuCapacity struct {
	Scaletype string `json:"scaleType,omitempty"` // The scale type applicable to the sku.
	DefaultField int `json:"default,omitempty"` // The default capacity.
	Maximum int `json:"maximum,omitempty"` // The maximum capacity that can be set.
	Minimum int `json:"minimum,omitempty"` // The minimum capacity.
}

// ResourceSkuResults represents the ResourceSkuResults schema from the OpenAPI specification
type ResourceSkuResults struct {
	Nextlink string `json:"nextLink,omitempty"` // The uri to fetch the next page of API Management service Skus.
	Value []ResourceSkuResult `json:"value"` // The list of skus available for the service.
}

// ApiManagementServiceBackupRestoreParameters represents the ApiManagementServiceBackupRestoreParameters schema from the OpenAPI specification
type ApiManagementServiceBackupRestoreParameters struct {
	Containername string `json:"containerName"` // Azure Cloud Storage blob container name used to place/retrieve the backup.
	Storageaccount string `json:"storageAccount"` // Azure Cloud Storage account (used to place/retrieve the backup) name.
	Accesskey string `json:"accessKey"` // Azure Cloud Storage account (used to place/retrieve the backup) access key.
	Backupname string `json:"backupName"` // The name of the backup file to create.
}

// ApiManagementServiceResource represents the ApiManagementServiceResource schema from the OpenAPI specification
type ApiManagementServiceResource struct {
	Id string `json:"id,omitempty"` // Resource ID.
	Name string `json:"name,omitempty"` // Resource name.
	Tags map[string]interface{} `json:"tags,omitempty"` // Resource tags.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource is set to Microsoft.ApiManagement.
}

// CertificateConfiguration represents the CertificateConfiguration schema from the OpenAPI specification
type CertificateConfiguration struct {
	Encodedcertificate string `json:"encodedCertificate,omitempty"` // Base64 Encoded certificate.
	Storename string `json:"storeName"` // The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and CertificateAuthority are valid locations.
	Certificate CertificateInformation `json:"certificate,omitempty"` // SSL certificate information.
	Certificatepassword string `json:"certificatePassword,omitempty"` // Certificate Password.
}

// CertificateInformation represents the CertificateInformation schema from the OpenAPI specification
type CertificateInformation struct {
	Expiry string `json:"expiry"` // Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Subject string `json:"subject"` // Subject of the certificate.
	Thumbprint string `json:"thumbprint"` // Thumbprint of the certificate.
}

// ApiManagementServiceBaseProperties represents the ApiManagementServiceBaseProperties schema from the OpenAPI specification
type ApiManagementServiceBaseProperties struct {
	Privateipaddresses []string `json:"privateIPAddresses,omitempty"` // Private Static Load Balanced IP addresses of the API Management service in Primary region which is deployed in an Internal Virtual Network. Available only for Basic, Standard and Premium SKU.
	Provisioningstate string `json:"provisioningState,omitempty"` // The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
	Customproperties map[string]interface{} `json:"customProperties,omitempty"` // Custom properties of the API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11` can be used to disable just TLS 1.1 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10` can be used to disable TLS 1.0 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2` can be used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH operation will reset omitted properties' values to their defaults. For all the settings except Http2 the default value is `True` if the service was created on or before April 1st 2018 and `False` otherwise. Http2 setting's default value is `False`.</br></br>You can disable any of next ciphers by using settings `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]`: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example, `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256`:`false`. The default value is `true` for them.
	Targetprovisioningstate string `json:"targetProvisioningState,omitempty"` // The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
	Additionallocations []AdditionalLocation `json:"additionalLocations,omitempty"` // Additional datacenter locations of the API Management service.
	Gatewayurl string `json:"gatewayUrl,omitempty"` // Gateway URL of the API Management service.
	Scmurl string `json:"scmUrl,omitempty"` // SCM endpoint URL of the API Management service.
	Virtualnetworktype string `json:"virtualNetworkType,omitempty"` // The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
	Hostnameconfigurations []HostnameConfiguration `json:"hostnameConfigurations,omitempty"` // Custom hostname configuration of the API Management service.
	Gatewayregionalurl string `json:"gatewayRegionalUrl,omitempty"` // Gateway URL of the API Management service in the Default Region.
	Createdatutc string `json:"createdAtUtc,omitempty"` // Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Publicipaddresses []string `json:"publicIPAddresses,omitempty"` // Public Static Load Balanced IP addresses of the API Management service in Primary region. Available only for Basic, Standard and Premium SKU.
	Managementapiurl string `json:"managementApiUrl,omitempty"` // Management API endpoint URL of the API Management service.
	Portalurl string `json:"portalUrl,omitempty"` // Publisher portal endpoint Url of the API Management service.
	Notificationsenderemail string `json:"notificationSenderEmail,omitempty"` // Email address from which the notification will be sent.
	Virtualnetworkconfiguration VirtualNetworkConfiguration `json:"virtualNetworkConfiguration,omitempty"` // Configuration of a virtual network to which API Management service is deployed.
	Certificates []CertificateConfiguration `json:"certificates,omitempty"` // List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
	Enableclientcertificate bool `json:"enableClientCertificate,omitempty"` // Property only meant to be used for Consumption SKU Service. This enforces a client certificate to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the policy on the gateway.
}

// ApiManagementServiceIdentity represents the ApiManagementServiceIdentity schema from the OpenAPI specification
type ApiManagementServiceIdentity struct {
	TypeField string `json:"type"` // The identity type. Currently the only supported type is 'SystemAssigned'.
	Principalid string `json:"principalId,omitempty"` // The principal id of the identity.
	Tenantid string `json:"tenantId,omitempty"` // The client tenant id of the identity.
}
