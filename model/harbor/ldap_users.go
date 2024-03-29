/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

type LdapUsers struct {

	// search ldap user name based on ldapconf.
	LdapUsername string `json:"ldap_username,omitempty"`

	// system will try to guess the user realname form \"uid\" or \"cn\" attribute.
	LdapRealname string `json:"ldap_realname,omitempty"`

	// system will try to guess the user email address form \"mail\" or \"email\" attribute.
	LdapEmail string `json:"ldap_email,omitempty"`
}
