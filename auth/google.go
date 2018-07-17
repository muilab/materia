package auth

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	jsoniter "github.com/json-iterator/go"
	"github.com/muilab/materia/mui"
	"github.com/muilab/materia/mui/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// GoogleUser is the user data we receive from Google
type GoogleUser struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

// InstallGoogleAuth enables Google login for the app.
func InstallGoogleAuth(app *aero.Application) {
	config := &oauth2.Config{
		ClientID:     mui.APIKeys.Google.ID,
		ClientSecret: mui.APIKeys.Google.Secret,
		RedirectURL:  "https://" + app.Config.Domain + "/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			// "https://www.googleapis.com/auth/plus.me",
			// "https://www.googleapis.com/auth/plus.login",
		},
		Endpoint: google.Endpoint,
	}

	app.Get("/auth/google", func(ctx *aero.Context) string {
		state := ctx.Session().ID()
		url := config.AuthCodeURL(state)
		return ctx.Redirect(url)
	})

	app.Get("/auth/google/callback", func(ctx *aero.Context) string {
		if !ctx.HasSession() {
			return ctx.Error(http.StatusUnauthorized, "Google login failed", errors.New("Session does not exist"))
		}

		session := ctx.Session()

		if session.ID() != ctx.Query("state") {
			return ctx.Error(http.StatusUnauthorized, "Google login failed", errors.New("Incorrect state"))
		}

		// Handle the exchange code to initiate a transport
		token, err := config.Exchange(context.Background(), ctx.Query("code"))

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Could not obtain OAuth token", err)
		}

		// Construct the OAuth client
		client := config.Client(context.Background(), token)

		// Fetch user data from Google
		resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed requesting user data from Google", err)
		}

		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)

		// Construct a GoogleUser object
		var googleUser GoogleUser
		err = jsoniter.Unmarshal(data, &googleUser)

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Failed parsing user data (JSON)", err)
		}

		if googleUser.Sub == "" {
			return ctx.Error(http.StatusBadRequest, "Failed retrieving Google data", errors.New("Empty ID"))
		}

		// Change googlemail.com to gmail.com
		googleUser.Email = strings.Replace(googleUser.Email, "googlemail.com", "gmail.com", 1)

		// Is this an existing user connecting another social account?
		user := mui.GetUserFromContext(ctx)

		if user != nil {
			// Add GoogleToUser reference
			user.ConnectGoogle(googleUser.Sub)

			// Save in DB
			user.Save()

			// Log
			authLog.Info("Added Google ID to existing account", user.ID, user.Nick, ctx.RealIP(), user.Accounts.Email.Address, user.RealName())

			return ctx.Redirect("/")
		}

		var getErr error

		// Try to find an existing user via the Google user ID
		user, getErr = mui.GetUserByGoogleID(googleUser.Sub)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Google ID", user.ID, user.Nick, ctx.RealIP(), user.Accounts.Email.Address, user.RealName())

			user.LastLogin = utils.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect("/")
		}

		// Try to find an existing user via the associated e-mail address
		user, getErr = mui.GetUserByEmail(googleUser.Email)

		if getErr == nil && user != nil {
			authLog.Info("User logged in via Email", user.ID, user.Nick, ctx.RealIP(), user.Accounts.Email.Address, user.RealName())

			user.LastLogin = utils.DateTimeUTC()
			user.Save()

			session.Set("userId", user.ID)
			return ctx.Redirect("/")
		}

		// Register new user
		user = mui.NewUser()
		user.Nick = "g" + googleUser.Sub
		user.FirstName = googleUser.GivenName
		user.LastName = googleUser.FamilyName
		user.LastLogin = utils.DateTimeUTC()

		// Connect email
		user.ConnectEmail(googleUser.Email)

		// Save basic user info already to avoid data inconsistency problems
		user.Save()

		// Register user
		mui.RegisterUser(user)

		// Connect account to a Google account
		user.ConnectGoogle(googleUser.Sub)

		// Save user object again with updated data
		user.Save()

		// Login
		session.Set("userId", user.ID)

		// Log
		authLog.Info("Registered new user via Google", user.ID, user.Nick, ctx.RealIP(), user.Accounts.Email.Address, user.RealName())

		// Redirect to starting page for new users
		return ctx.Redirect(newUserStartRoute)
	})
}
