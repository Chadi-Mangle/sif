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

	users, _ := queries.ListUsers(ctx)

	for i, user := range users {
		fmt.Printf("User%d : %v\n\n", i, user)
	}

	handler := handlers.NewHandler(ctx, *queries, cfg.Server.SecretKey)

	// Web :
	component := templates.Hello("World !")
	http.Handle("/", templ.Handler(component))

	http.HandleFunc("GET /login", handler.GetSignIn)
	http.HandleFunc("POST /login", handler.PostSignIn)

	http.HandleFunc("GET /register", handler.GetSignUp)
	http.HandleFunc("POST /register", handler.PostSingUp)

	// Route temporaire pour le dashboard
	http.HandleFunc("GET /dashboard", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Dashboard("John", "Doe", false)
		templ.Handler(component).ServeHTTP(w, r)
	})

	// Route pour la page de paiement
	http.HandleFunc("GET /payment", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Payment()
		templ.Handler(component).ServeHTTP(w, r)
	})

	// Route pour la page de signature
	http.HandleFunc("GET /signature", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Signature("John", "Doe", 42)
		templ.Handler(component).ServeHTTP(w, r)
	})

	// Route pour la page de profil
	http.HandleFunc("GET /profile", func(w http.ResponseWriter, r *http.Request) {
		// Récupérer les paramètres de l'URL
		paymentMethod := r.URL.Query().Get("payment")
		paymentStatus := r.URL.Query().Get("status")
		
		// Données de préférences simulées
		preferences := map[string]bool{
			"vegetarian":      true,
			"has_license":     true,
			"has_vehicle":     false,
			"halal_meat":      false,
			"wants_to_drive":  true,
			"has_pathologies": false,
		}
		
		component := templates.Profile("John", "Doe", false, "john.doe@example.com", paymentMethod, paymentStatus, 42, preferences)
		templ.Handler(component).ServeHTTP(w, r)
	})

	// Routes pour les pages d'information
	http.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		component := templates.About()
		templ.Handler(component).ServeHTTP(w, r)
	})

	http.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Contact()
		templ.Handler(component).ServeHTTP(w, r)
	})

	http.HandleFunc("GET /faq", func(w http.ResponseWriter, r *http.Request) {
		component := templates.FAQ()
		templ.Handler(component).ServeHTTP(w, r)
	})

	http.HandleFunc("GET /terms", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Terms()
		templ.Handler(component).ServeHTTP(w, r)
	})

	http.HandleFunc("GET /privacy", func(w http.ResponseWriter, r *http.Request) {
		component := templates.Privacy()
		templ.Handler(component).ServeHTTP(w, r)
	})

	serverAddr := cfg.GetServerAddress()
	fmt.Printf("Listening on %s\n", serverAddr)
	http.ListenAndServe(":"+cfg.Server.Port, nil)
}
