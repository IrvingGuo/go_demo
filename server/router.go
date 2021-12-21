package server

import (
	"resource-plan-improvement/api"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(logger(), cors())

	apiRouter := router.Group("/api")
	{
		apiRouter.POST("/login", api.Signin)

		userGroup := apiRouter.Group("/user")
		userGroup.Use(TokenVerifier())
		{
			userGroup.POST("", api.SaveUser)
			userGroup.GET("", api.GetAllUsers)
			userGroup.PUT("", api.SaveUser)

			userGroup.GET("/current", api.AutoSignin)

			userGroup.GET("/:id", api.GetUserById)
			userGroup.DELETE("/:id", api.DeleteUserById)

			userGroup.GET("/subor", api.GetUsersUnderLeader)
		}

		programGroup := apiRouter.Group("/program")
		programGroup.Use(TokenVerifier())
		{
			programGroup.POST("", api.SaveProgram)
			programGroup.GET("", api.GetAllPrograms)
			programGroup.PUT("", api.SaveProgram)

			programGroup.GET("/:id", api.GetProgramById)
			programGroup.DELETE("/:id", api.DeleteProgramById)

			programGroup.GET("/user", api.GetProgramsUnderUser)
		}

		subprogramGroup := apiRouter.Group("/subprogram")
		subprogramGroup.Use(TokenVerifier())
		{
			subprogramGroup.POST("", api.SaveSubprogram)
			subprogramGroup.GET("", api.GetAllSubprograms)
			subprogramGroup.PUT("", api.SaveSubprogram)
			subprogramGroup.DELETE("/:id", api.DeleteSubprogramById)
		}

		activityGroup := apiRouter.Group("/activity")
		activityGroup.Use(TokenVerifier())
		{
			activityGroup.POST("", api.SaveActivity)
			activityGroup.GET("", api.GetAllActivities)
			activityGroup.PUT("", api.SaveActivity)
			activityGroup.DELETE("/:id", api.DeleteActivityById)
		}

		assignmentGroup := apiRouter.Group("/assignment")
		assignmentGroup.Use(TokenVerifier())
		{
			assignmentGroup.POST("", api.SaveAssignments)
			assignmentGroup.GET("", api.GetAllAssignments)
			assignmentGroup.PUT("", api.UpdateAssignments)
			assignmentGroup.DELETE("", api.DeleteAssignmentsByIds)

			assignmentGroup.PUT("/status", api.UpdateAssignmentsStatus)

			assignmentGroup.GET("/leader", api.GetAssignmentsUnderLeader)
			assignmentGroup.GET("/tpm", api.GetAssignmentsUnderTpm)
			assignmentGroup.GET("/program/:id", api.GetAssignmentsByProgramId)

			assignmentGroup.GET("/excel", api.GetMasterResPlanExcelFilename)
		}

		deptGroup := apiRouter.Group("/department")
		deptGroup.Use(TokenVerifier())
		{
			deptGroup.POST("", api.SaveDepartment)
			deptGroup.GET("", api.GetAllDepartments)
			deptGroup.PUT("", api.SaveDepartment)

			deptGroup.GET("/:id", api.GetDepartmentById)
			deptGroup.DELETE("/:id", api.DeleteDepartmentById)

			deptGroup.GET("/user", api.GetDepartmentsUnderLoginUser)
		}

		otherGroup := apiRouter.Group("")
		otherGroup.Use(TokenVerifier())
		{
			otherGroup.GET("/download/:filename", api.DownloadFile)
		}
	}

	return router
}
