package app

import (
	"github.com/jmckind/lyceum/app/db"
	"github.com/jmckind/lyceum/app/services"
	"github.com/revel/revel"
	rethinkdb "gopkg.in/gorethink/gorethink.v4"
)

type AppServices struct {
	ArtifactService     *services.ArtifactService
	AuthService         *services.AuthService
	ItemService         *services.ItemService
	LibraryService      *services.LibraryService
	OrganizationService *services.OrganizationService
	RoleService         *services.RoleService
	UserService         *services.UserService
}

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

	// RethinkDBSession rethinkdb session handle for the application
	RethinkDBSession *rethinkdb.Session

	// Services map of application services
	Services *AppServices
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	revel.OnAppStart(RegisterServices)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}

func InitDB() {
	opts := map[string]interface{}{
		"listen_ip":      "",
		"listen_port":    4778,
		"db_url":         "localhost:28015",
		"db_con_initial": 10,
		"db_con_max":     10,
	}
	session, err := db.ConnectRethinkDB(opts)
	if err != nil {
		revel.AppLog.Errorf("unable to connect to database: %v", err)
	}
	RethinkDBSession = session
}

func RegisterServices() {
	Services = &AppServices{
		ArtifactService:     services.NewArtifactService(RethinkDBSession),
		AuthService:         services.NewAuthService(RethinkDBSession),
		ItemService:         services.NewItemService(RethinkDBSession),
		LibraryService:      services.NewLibraryService(RethinkDBSession),
		OrganizationService: services.NewOrganizationService(RethinkDBSession),
		RoleService:         services.NewRoleService(RethinkDBSession),
		UserService:         services.NewUserService(RethinkDBSession),
	}
}
