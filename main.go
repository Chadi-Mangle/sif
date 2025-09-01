package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jackc/pgx/v5"

	assetsfs "github.com/Chadi-Mangle/templ-hmr-setup/assets"
	"github.com/Chadi-Mangle/templ-hmr-setup/internal/config"
	"github.com/Chadi-Mangle/templ-hmr-setup/internal/handlers"
	"github.com/Chadi-Mangle/templ-hmr-setup/internal/models"
	"github.com/Chadi-Mangle/templ-hmr-setup/templates"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	var httpRoot http.FileSystem
	if assetsfs.IsEmbedded {
		httpRoot = http.FS(assetsfs.Box)
	} else {
		httpRoot = http.Dir("assets")
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(httpRoot)))

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.GetDatabaseURL())
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	queries := models.New(conn)

	handler := handlers.NewHandler(ctx, *queries)

	// Web :
	component := templates.Hello("World !")
	http.Handle("/", templ.Handler(component))

	http.HandleFunc("/login", handler.SignIn)

	serverAddr := cfg.GetServerAddress()
	fmt.Printf("Listening on %s\n", serverAddr)
	http.ListenAndServe(":"+cfg.Server.Port, nil)
}
