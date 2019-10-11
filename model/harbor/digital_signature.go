/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

// The signature of the chart
type DigitalSignature struct {

	// A flag to indicate if the chart is signed
	Signed bool `json:"signed,omitempty"`

	// The URL of the provance file
	ProvFile string `json:"prov_file,omitempty"`
}
