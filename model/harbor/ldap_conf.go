/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

type LdapConf struct {

	// The url of ldap service.
	LdapUrl string `json:"ldap_url,omitempty"`

	// The search dn of ldap service.
	LdapSearchDn string `json:"ldap_search_dn,omitempty"`

	// The search password of ldap service.
	LdapSearchPassword string `json:"ldap_search_password,omitempty"`

	// The base dn of ldap service.
	LdapBaseDn string `json:"ldap_base_dn,omitempty"`

	// The serach filter of ldap service.
	LdapFilter string `json:"ldap_filter,omitempty"`

	// The serach uid from ldap service attributes.
	LdapUid string `json:"ldap_uid,omitempty"`

	// The serach scope of ldap service.
	LdapScope int64 `json:"ldap_scope,omitempty"`

	// The connect timeout of ldap service(second).
	LdapConnectionTimeout int64 `json:"ldap_connection_timeout,omitempty"`
}
