package main

import (
	"fields/internal/field"
	fieldAdapter "fields/internal/field/adapter"
	"fields/internal/owner"
	ownerAdapter "fields/internal/owner/adapter"
	"fields/internal/reservation"
	reservationAdapter "fields/internal/reservation/adapter"
	"fields/pkg/auth"
	"fields/pkg/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Hello World!")

	err := database.InitPool("app/migrations")
	if err != nil {
		fmt.Printf("error initializing database: %v\n", err)
		return
	}

	fmt.Println("Database initialized")

	app := fiber.New()
	fieldRepo := fieldAdapter.NewFieldRepositoryDB(database.GetPool())
	ownerRepo := ownerAdapter.NewOwnerRepositoryDB(database.GetPool())
	resvRepo := reservationAdapter.NewReservationRepositoryDB(database.GetPool())

	fieldService := field.NewFieldService(fieldRepo)
	ownerService := owner.NewOwnerService(ownerRepo)
	resvService := reservation.NewReservationService(resvRepo, fieldService)

	fieldHandler := field.NewFieldHandler(fieldService)
	ownerHandler := owner.NewOwnerHandler(ownerService)
	resvHandler := reservation.NewReservationHandler(resvService)

	ConfigRoutes(app, resvHandler, fieldHandler, ownerHandler)

	app.Listen(":8080")

}

func ConfigRoutes(app *fiber.App, r *reservation.ReservationHandler, f *field.FieldHandler, o *owner.OwnerHandler) {
	api := app.Group("/api/v1")
	reservations := api.Group("/reservations")
	reservations.Post("/", auth.JWTProtected(), r.CreateReservation)
	reservations.Get("/:id", r.GetReservation)
	reservations.Get("/", r.ListReservation)
	reservations.Get("/booker/:bookerId", auth.JWTProtected(), r.ListReservationByBookerId)
	reservations.Get("/field/:fieldId", r.ListReservationByFieldId)

	// Field routes
	fields := api.Group("/fields")
	fields.Get("/:id", f.GetField)
	fields.Get("/", f.ListFields)
	fields.Get("/owner/:ownerId", f.ListFieldsByOwnerId)

	// Owner routes
	owners := api.Group("/owners")
	owners.Get("/:id", o.GetOwner)

}
