package routes

/*
type MyEventsRouter struct {
	myBookingsController *controllers.MyEventsController
	authStrategy middlewares.AuthStrategy
}

func NewMyEventsRouter(myBookingsController *controllers.MyEventsController, authStrategy middlewares.AuthStrategy) *MyBookingsRouter{
	return &MyBookingsRouter{
		myBookingsController: myBookingsController,
		authStrategy: authStrategy,
	}
}

func (mbr *MyBookingsRouter) RegisterRoutes(echo *echo.Echo){
	//creating group
	myBookingsGroup := echo.Group("/my-bookings", middlewares.AuthMiddleware(mbr.authStrategy))

	//suscribing controller
	myBookingsGroup.GET("", mbr.myBookingsController.GetAll)
}*/
