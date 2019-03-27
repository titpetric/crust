package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/pkg/errors"
	"github.com/titpetric/factory/resputil"

	"github.com/crusttech/crust/internal/auth"
	"github.com/crusttech/crust/internal/db"
	"github.com/crusttech/crust/internal/mail"
	"github.com/crusttech/crust/internal/metrics"
	"github.com/crusttech/crust/internal/settings"
	migrate "github.com/crusttech/crust/system/db"
	"github.com/crusttech/crust/system/internal/auth/social"
	"github.com/crusttech/crust/system/internal/repository"
	"github.com/crusttech/crust/system/service"
)

var (
	jwtVerifier      func(http.Handler) http.Handler
	jwtAuthenticator func(http.Handler) http.Handler
	jwtEncoder       auth.TokenEncoder
)

func Init() error {
	// validate configuration
	if err := flags.Validate(); err != nil {
		return err
	}
	// JWT Auth
	if jwtAuth, err := auth.JWT(); err != nil {
		return errors.Wrap(err, "Error creating JWT Auth object")
	} else {
		jwtEncoder = jwtAuth
		jwtVerifier = jwtAuth.Verifier()
		jwtAuthenticator = jwtAuth.Authenticator()
	}

	mail.SetupDialer(flags.smtp)

	InitDatabase()

	// Load settings from the database,
	// for now, only at start-up time.
	ctx := context.Background()
	// settingsRepository := internalRepository.NewSettings(repository.DB(ctx), "sys_settings")
	// if ss, err := settingsRepository.Find(types.SettingsFilter{}); err != nil {
	// 	panic(err)
	// } else {
	// 	spew.Dump(ss.KV())
	// }
	settingService := settings.NewService(settings.NewRepository(repository.DB(ctx), "sys_settings"))

	// configure resputil options
	resputil.SetConfig(resputil.Options{
		Pretty: flags.http.Pretty,
		Trace:  flags.http.Tracing,
		Logger: func(err error) {
			// @todo: error logging
		},
	})

	// Don't change this, it needs database connection
	service.Init()

	// Setup goth/social authentication
	social.Init(flags.social, settingService.With(ctx))

	return nil
}

func InitDatabase() error {
	// start/configure database connection
	db, err := db.TryToConnect("system", flags.db.DSN, flags.db.Profiler)
	if err != nil {
		return errors.Wrap(err, "could not connect to database")
	}

	// migrate database schema
	if err := migrate.Migrate(db); err != nil {
		return err
	}

	return nil
}

func StartRestAPI(ctx context.Context) error {
	log.Println("Starting http server on address " + flags.http.Addr)
	listener, err := net.Listen("tcp", flags.http.Addr)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Can't listen on addr %s", flags.http.Addr))
	}

	if flags.monitor.Interval > 0 {
		go metrics.NewMonitor(flags.monitor.Interval)
	}

	go http.Serve(listener, Routes(ctx))
	<-ctx.Done()

	return nil
}
