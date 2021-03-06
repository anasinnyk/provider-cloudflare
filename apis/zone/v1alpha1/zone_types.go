/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ZoneParameters are the configurable fields of a Zone.
type ZoneParameters struct {
	// Name is the name of the Zone, which should be a valid
	// domain.
	// +kubebuilder:validation:Format=hostname
	// +kubebuilder:validation:MaxLength=253
	// +immutable
	Name string `json:"name"`

	// JumpStart enables attempting to import existing DNS records
	// when a new Zone is created.
	// +kubebuilder:default=false
	// +immutable
	// +optional
	JumpStart bool `json:"jumpStart,"`

	// Paused indicates if the zone is only using Cloudflare DNS services.
	// +optional
	Paused *bool `json:"paused,omitempty"`

	// PlanID indicates the plan that this Zone will be subscribed
	// to.
	// +optional
	PlanID *string `json:"planId,omitempty"`

	// Type indicates the type of this zone - partial (partner-hosted
	// or CNAME only) or full.
	// +kubebuilder:validation:Enum=full;partial
	// +kubebuilder:default=full
	// +immutable
	// +optional
	Type *string `json:"type,omitempty"`

	// Settings contains a Zone settings that can be applied
	// to this zone.
	// +optional
	Settings ZoneSettings `json:"settings,omitempty"`

	// VanityNameServers lists an array of domains to use for custom
	// nameservers.
	// +optional
	VanityNameServers []string `json:"vanityNameServers,omitempty"`
	// +optional
	AccountRef *xpv1.Reference `json:"accountRef,omitempty"`
	// +optional
	AccountSelector *xpv1.Selector `json:"accountSelector,omitempty"`
}

// MinifySettings represents the minify settings on a Zone
type MinifySettings struct {
	// CSS enables or disables minifying CSS assets
	// +kubebuilder:validation:Enum=off;on
	// +optional
	CSS *string `json:"css,omitempty"`
	// HTML enables or disables minifying HTML assets
	// +kubebuilder:validation:Enum=off;on
	// +optional
	HTML *string `json:"html,omitempty"`
	// JS enables or disables minifying JS assets
	// +kubebuilder:validation:Enum=off;on
	// +optional
	JS *string `json:"js,omitempty"`
}

// MobileRedirectSettings represents the mobile_redirect settings on a Zone
type MobileRedirectSettings struct {
	// Status enables or disables mobile redirection
	// +kubebuilder:validation:Enum=off;on
	// +optional
	Status *string `json:"status,omitempty"`
	// Subdomain defines the subdomain prefix to redirect mobile devices to
	// +optional
	Subdomain *string `json:"subdomain,omitempty"`
	// StripURI defines whether or not to strip the path from the URI when redirecting
	// +optional
	StripURI *bool `json:"stripURI,omitempty"`
}

// StrictTransportSecuritySettings represents the STS settings on a Zone's security headers
type StrictTransportSecuritySettings struct {
	// Enabled enables or disables STS settings
	// +optional
	Enabled *bool `json:"enabled,omitempty"`
	// MaxAge defines the maximum age in seconds of the STS
	// +optional
	MaxAge *int64 `json:"maxAge,omitempty"`
	// IncludeSubdomains defines whether or not to include all subdomains
	// +optional
	IncludeSubdomains *bool `json:"includeSubdomains,omitempty"`
	// NoSniff defines whether or not to include 'X-Content-Type-Options: nosniff' header
	// +optional
	NoSniff *bool `json:"noSniff,omitempty"`
}

// SecurityHeaderSettings represents the security headers on a Zone
type SecurityHeaderSettings struct {
	// StrictTransportSecurity defines the STS settings on a Zone
	// +optional
	StrictTransportSecurity *StrictTransportSecuritySettings `json:"strictTransportSecurity,omitempty"`
}

// ZoneSettings represents settings on a Zone
type ZoneSettings struct { // AlwaysOnline enables or disables Always Online
	// +kubebuilder:validation:Enum=off;on
	// +optional
	AlwaysOnline *string `json:"alwaysOnline,omitempty"`

	// AdvancedDDOS enables or disables Advanced DDoS mitigation
	// +kubebuilder:validation:Enum=off;on
	// +optional
	AdvancedDDOS *string `json:"advancedDdos,omitempty"`

	// AlwaysUseHTTPS enables or disables Always use HTTPS
	// +kubebuilder:validation:Enum=off;on
	// +optional
	AlwaysUseHTTPS *string `json:"alwaysUseHttps,omitempty"`

	// AutomaticHTTPSRewrites enables or disables Automatic HTTPS Rewrites
	// +kubebuilder:validation:Enum=off;on
	// +optional
	AutomaticHTTPSRewrites *string `json:"automaticHttpsRewrites,omitempty"`

	// Brotli enables or disables Brotli
	// +kubebuilder:validation:Enum=off;on
	// +optional
	Brotli *string `json:"brotli,omitempty"`

	// BrowserCacheTTL configures the browser cache ttl.
	// 0 means respect existing headers
	// +kubebuilder:validation:Enum=0;30;60;300;1200;1800;3600;7200;10800;14400;18000;28800;43200;57600;72000;86400;172800;259200;345600;432000;691200;1382400;2073600;2678400;5356800;16070400;31536000
	// +optional
	BrowserCacheTTL *int64 `json:"browserCacheTtl,omitempty"`

	// BrowserCheck enables or disables Browser check
	// +kubebuilder:validation:Enum=off;on
	// +optional
	BrowserCheck *string `json:"browserCheck,omitempty"`

	// CacheLevel configures the cache level
	// +kubebuilder:validation:Enum=bypass;basic;simplified;aggressive;cache_everything
	// +optional
	CacheLevel *string `json:"cacheLevel,omitempty"`

	// ChallengeTTL configures the edge cache ttl
	// +kubebuilder:validation:Enum=300;900;1800;2700;3600;7200;10800;14400;28800;57600;86400;604800;2592000;31536000
	// +optional
	ChallengeTTL *int64 `json:"challengeTtl,omitempty"`

	// Ciphers configures which ciphers are allowed for TLS termination
	// +optional
	Ciphers []string `json:"ciphers,omitempty"`

	// CnameFlattening configures CNAME flattening
	// +kubebuilder:validation:Enum=flatten_at_root;flatten_all;flatten_none
	// +optional
	CnameFlattening *string `json:"cnameFlattening,omitempty"`

	// DevelopmentMode enables or disables Development mode
	// +kubebuilder:validation:Enum=off;on
	// +optional
	DevelopmentMode *string `json:"developmentMode,omitempty"`

	// EdgeCacheTTL configures the edge cache ttl
	// +optional
	EdgeCacheTTL *int64 `json:"edgeCacheTtl,omitempty"`

	// EmailObfuscation enables or disables Email obfuscation
	// +kubebuilder:validation:Enum=off;on
	// +optional
	EmailObfuscation *string `json:"emailObfuscation,omitempty"`

	// HotlinkProtection enables or disables Hotlink protection
	// +kubebuilder:validation:Enum=off;on
	// +optional
	HotlinkProtection *string `json:"hotlinkProtection,omitempty"`

	// HTTP2 enables or disables HTTP2
	// +kubebuilder:validation:Enum=off;on
	// +optional
	HTTP2 *string `json:"http2,omitempty"`

	// HTTP3 enables or disables HTTP3
	// +kubebuilder:validation:Enum=off;on
	// +optional
	HTTP3 *string `json:"http3,omitempty"`

	// IPGeolocation enables or disables IP Geolocation
	// +kubebuilder:validation:Enum=off;on
	// +optional
	IPGeolocation *string `json:"ipGeolocation,omitempty"`

	// IPv6 enables or disables IPv6
	// +kubebuilder:validation:Enum=off;on
	// +optional
	IPv6 *string `json:"ipv6,omitempty"`

	// LogToCloudflare enables or disables Logging to cloudflare
	// +kubebuilder:validation:Enum=off;on
	// +optional
	LogToCloudflare *string `json:"logToCloudflare,omitempty"`

	// MaxUpload configures the maximum upload payload size
	// +optional
	MaxUpload *int64 `json:"maxUpload,omitempty"`

	// Minify configures minify settings for certain assets
	// +optional
	Minify *MinifySettings `json:"minify,omitempty"`

	// MinTLSVersion configures the minimum TLS version
	// +kubebuilder:validation:Enum="1.0";"1.1";"1.2";"1.3"
	// +optional
	MinTLSVersion *string `json:"minTLSVersion,omitempty"`

	// Mirage enables or disables Mirage
	// +kubebuilder:validation:Enum=off;on
	// +optional
	Mirage *string `json:"mirage,omitempty"`

	// MobileRedirect configures automatic redirections to mobile-optimized subdomains
	// +optional
	MobileRedirect *MobileRedirectSettings `json:"mobileRedirect,omitempty"`

	// OpportunisticEncryption enables or disables Opportunistic encryption
	// +kubebuilder:validation:Enum=off;on
	// +optional
	OpportunisticEncryption *string `json:"opportunisticEncryption,omitempty"`

	// OpportunisticOnion enables or disables Opportunistic onion
	// +kubebuilder:validation:Enum=off;on
	// +optional
	OpportunisticOnion *string `json:"opportunisticOnion,omitempty"`

	// OrangeToOrange enables or disables Orange to orange
	// +kubebuilder:validation:Enum=off;on
	// +optional
	OrangeToOrange *string `json:"orangeToOrange,omitempty"`

	// OriginErrorPagePassThru enables or disables Mirage
	// +kubebuilder:validation:Enum=off;on
	// +optional
	OriginErrorPagePassThru *string `json:"originErrorPagePassThru,omitempty"`

	// Polish configures the Polish setting
	// +kubebuilder:validation:Enum=off;lossless;lossy
	// +optional
	Polish *string `json:"polish,omitempty"`

	// PrefetchPreload enables or disables Prefetch preload
	// +kubebuilder:validation:Enum=off;on
	// +optional
	PrefetchPreload *string `json:"prefetchPreload,omitempty"`

	// PrivacyPass enables or disables Privacy pass
	// +kubebuilder:validation:Enum=off;on
	// +optional
	PrivacyPass *string `json:"privacyPass,omitempty"`

	// PseudoIPv4 configures the Pseudo IPv4 setting
	// +kubebuilder:validation:Enum=off;add_header;overwrite_header
	// +optional
	PseudoIPv4 *string `json:"pseudoIpv4,omitempty"`

	// ResponseBuffering enables or disables Response buffering
	// +kubebuilder:validation:Enum=off;on
	// +optional
	ResponseBuffering *string `json:"responseBuffering,omitempty"`

	// RocketLoader enables or disables Rocket loader
	// +kubebuilder:validation:Enum=off;on
	// +optional
	RocketLoader *string `json:"rocketLoader,omitempty"`

	// SecurityHeader defines the security headers for a Zone
	// +optional
	SecurityHeader *SecurityHeaderSettings `json:"securityHeader,omitempty"`

	// SecurityLevel configures the Security level
	// +kubebuilder:validation:Enum=off;essentially_off;low;medium;high;under_attack
	// +optional
	SecurityLevel *string `json:"securityLevel,omitempty"`

	// ServerSideExclude enables or disables Server side exclude
	// +kubebuilder:validation:Enum=off;on
	// +optional
	ServerSideExclude *string `json:"serverSideExclude,omitempty"`

	// SortQueryStringForCache enables or disables Sort query string for cache
	// +kubebuilder:validation:Enum=off;on
	// +optional
	SortQueryStringForCache *string `json:"sortQueryStringForCache,omitempty"`

	// SSL configures the SSL mode
	// +kubebuilder:validation:Enum=off;flexible;full;strict;origin_pull
	// +optional
	SSL *string `json:"ssl,omitempty"`

	// TLS13 configures TLS 1.3
	// +kubebuilder:validation:Enum=off;on;zrt
	// +optional
	TLS13 *string `json:"tls13,omitempty"`

	// TLSClientAuth enables or disables TLS client authentication
	// +kubebuilder:validation:Enum=off;on
	// +optional
	TLSClientAuth *string `json:"tlsClientAuth,omitempty"`

	// TrueClientIPHeader enables or disables True client IP Header
	// +kubebuilder:validation:Enum=off;on
	// +optional
	TrueClientIPHeader *string `json:"trueClientIPHeader,omitempty"`

	// VisitorIP enables or disables Visitor IP
	// +kubebuilder:validation:Enum=off;on
	// +optional
	VisitorIP *string `json:"visitorIP,omitempty"`

	// WAF enables or disables the Web application firewall
	// +kubebuilder:validation:Enum=off;on
	// +optional
	WAF *string `json:"waf,omitempty"`

	// WebP enables or disables WebP
	// +kubebuilder:validation:Enum=off;on
	// +optional
	WebP *string `json:"webP,omitempty"`

	// WebSockets enables or disables Web sockets
	// +kubebuilder:validation:Enum=off;on
	// +optional
	WebSockets *string `json:"webSockets,omitempty"`

	// ZeroRTT enables or disables Zero RTT
	// +kubebuilder:validation:Enum=off;on
	// +optional
	ZeroRTT *string `json:"zeroRtt,omitempty"`
}

// ZoneObservation are the observable fields of a Zone.
type ZoneObservation struct {
	// AccountID is the account ID that this zone exists under
	AccountID string `json:"accountId,omitempty"`

	// AccountName is the account name that this zone exists under
	Account string `json:"accountName,omitempty"`

	// DevModeTimer indicates the number of seconds left
	// in dev mode (if positive), otherwise the number
	// of seconds since dev mode expired.
	DevModeTimer int `json:"devModeTimer,omitempty"`

	// OriginalNS lists the original nameservers when
	// this Zone was created.
	OriginalNS []string `json:"originalNameServers,omitempty"`

	// OriginalRegistrar indicates the original registrar
	// when this Zone was created.
	OriginalRegistrar string `json:"originalRegistrar,omitempty"`

	// OriginalDNSHost indicates the original DNS host
	// when this Zone was created.
	OriginalDNSHost string `json:"originalDNSHost,omitempty"`

	// NameServers lists the Name servers that are assigned
	// to this Zone.
	NameServers []string `json:"nameServers,omitempty"`

	// PlanID indicates the billing plan ID assigned
	// to this Zone.
	PlanID string `json:"planId,omitempty"`

	// Plan indicates the name of the plan assigned
	// to this Zone.
	Plan string `json:"plan,omitempty"`

	// PlanPendingID indicates the ID of the pending plan
	// assigned to this Zone.
	PlanPendingID string `json:"planPendingId,omitempty"`

	// PlanPending indicates the name of the pending plan
	// assigned to this Zone.
	PlanPending string `json:"planPending,omitempty"`

	// Status indicates the status of this Zone.
	Status string `json:"status,omitempty"`

	// Betas indicates the betas available on this Zone.
	Betas []string `json:"betas,omitempty"`

	// DeactReason indicates the deactivation reason on
	// this Zone.
	DeactReason string `json:"deactivationReason,omitempty"`

	// VerificationKey indicates the Verification key set
	// on this Zone.
	VerificationKey string `json:"verificationKey,omitempty"`

	// VanityNameServers lists the currently assigned vanity
	// name server addresses.
	VanityNameServers []string `json:"vanityNameServers,omitempty"`
}

// A ZoneSpec defines the desired state of a Zone.
type ZoneSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ZoneParameters `json:"forProvider"`
}

// A ZoneStatus represents the observed state of a Zone.
type ZoneStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Zone is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudfalre}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ZoneSpec   `json:"spec"`
	Status ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zone
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Zone type metadata.
var (
	ZoneKind             = reflect.TypeOf(Zone{}).Name()
	ZoneGroupKind        = schema.GroupKind{Group: Group, Kind: ZoneKind}.String()
	ZoneKindAPIVersion   = ZoneKind + "." + SchemeGroupVersion.String()
	ZoneGroupVersionKind = SchemeGroupVersion.WithKind(ZoneKind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
