/*
 * aid-ukraine
 *
 * Migration aid database and searcher for the Ukrain refugees.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Person struct {

	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Phone string `json:"phone,omitempty"`

	Email string `json:"email,omitempty"`

	NightCall bool `json:"nightCall,omitempty"`

	Status *Status `json:"status,omitempty"`

	StatusComment string `json:"statusComment,omitempty"`
}
