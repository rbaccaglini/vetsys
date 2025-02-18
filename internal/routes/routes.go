package routes

import (
	"github.com/gin-gonic/gin"
	holiday_handler "github.com/rbaccaglini/vetsys/internal/handlers/holiday"
)

func InitRouter(r *gin.RouterGroup, handler holiday_handler.HolidayHandlerInterface) {

	holidayRoutes := r.Group("/holiday")
	{
		holidayRoutes.GET("", handler.FindAllHoliday)
		holidayRoutes.GET("/:date", handler.FindHolidayByDate)
		holidayRoutes.POST("", handler.CreateHoliday)
	}

	speciesRoutes := r.Group("/species")
	{
		speciesRoutes.GET("")
	}

	breedRoutes := r.Group("/breed")
	{
		breedRoutes.GET("")
	}

	sampleRoutes := r.Group("/sample")
	{
		sampleRoutes.GET("")
	}

	examTypeRoutes := r.Group("/exam-type")
	{
		examTypeRoutes.GET("")
	}

	clinicRoutes := r.Group("/clinic")
	{
		clinicRoutes.GET("")
	}

	veterinarianRoutes := r.Group("/veterinarian")
	{
		veterinarianRoutes.GET("")
	}

	customerRoutes := r.Group("/customer")
	{
		customerRoutes.GET("")
	}

	animalRoutes := r.Group("/animal")
	{
		animalRoutes.GET("")
	}

	examRoutes := r.Group("/exam")
	{
		examRoutes.GET("")
	}

}
