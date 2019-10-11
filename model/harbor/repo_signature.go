/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

type RepoSignature struct {

	// The tag of image.
	Tag string `json:"tag,omitempty"`

	// The JSON object of the hash of the image.
	Hashes *interface{} `json:"hashes,omitempty"`
}
