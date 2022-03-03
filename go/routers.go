/*
 * aid-ukraine
 *
 * Migration aid database and searcher for the Ukrain refugees.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"AddAccomodation",
		strings.ToUpper("Post"),
		"/v1/accomodation",
		AddAccomodation,
	},

	Route{
		"DeleteAccomodation",
		strings.ToUpper("Delete"),
		"/v1/accomodation/{accomodationId}",
		DeleteAccomodation,
	},

	Route{
		"FilterAccomodations",
		strings.ToUpper("Get"),
		"/v1/accomodation/filter",
		FilterAccomodations,
	},

	Route{
		"FindAccomodationsByStatus",
		strings.ToUpper("Get"),
		"/v1/accomodation/findByStatus",
		FindAccomodationsByStatus,
	},

	Route{
		"GetAccomodationById",
		strings.ToUpper("Get"),
		"/v1/accomodation/{accomodationId}",
		GetAccomodationById,
	},

	Route{
		"ListAllAccomodation",
		strings.ToUpper("Get"),
		"/v1/accomodation",
		ListAllAccomodation,
	},

	Route{
		"UpdateAccomodation",
		strings.ToUpper("Put"),
		"/v1/accomodation",
		UpdateAccomodation,
	},

	Route{
		"AddPerson",
		strings.ToUpper("Post"),
		"/v1/people",
		AddPerson,
	},

	Route{
		"DeletePerson",
		strings.ToUpper("Delete"),
		"/v1/people/{personId}",
		DeletePerson,
	},

	Route{
		"FindPeopleByStatus",
		strings.ToUpper("Get"),
		"/v1/people/findByStatus",
		FindPeopleByStatus,
	},

	Route{
		"GetPersonById",
		strings.ToUpper("Get"),
		"/v1/people/{personId}",
		GetPersonById,
	},

	Route{
		"ListAllPeople",
		strings.ToUpper("Get"),
		"/v1/people",
		ListAllPeople,
	},

	Route{
		"UpdatePerson",
		strings.ToUpper("Put"),
		"/v1/people",
		UpdatePerson,
	},

	Route{
		"AddTransportation",
		strings.ToUpper("Post"),
		"/v1/transport",
		AddTransportation,
	},

	Route{
		"DeleteTransportation",
		strings.ToUpper("Delete"),
		"/v1/transport/{transportationId}",
		DeleteTransportation,
	},

	Route{
		"FilterTransportations",
		strings.ToUpper("Get"),
		"/v1/transport/filter",
		FilterTransportations,
	},

	Route{
		"FindTransportationsByStatus",
		strings.ToUpper("Get"),
		"/v1/transport/findByStatus",
		FindTransportationsByStatus,
	},

	Route{
		"GetTransportationById",
		strings.ToUpper("Get"),
		"/v1/transport/{transportationId}",
		GetTransportationById,
	},

	Route{
		"ListAllTransportation",
		strings.ToUpper("Get"),
		"/v1/transport",
		ListAllTransportation,
	},

	Route{
		"UpdateTransportation",
		strings.ToUpper("Put"),
		"/v1/transport",
		UpdateTransportation,
	},
}