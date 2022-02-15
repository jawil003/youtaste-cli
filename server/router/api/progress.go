package api

import (
	"bs-to-scrapper/server/services"
	"errors"
	"github.com/gin-gonic/gin"
)

func RegisterProgress(api *gin.RouterGroup) {
	progress := api.Group("/progress")
	{
		progress.GET("", func(context *gin.Context) {
			treeService := services.DB().ProgressTree()

			context.JSON(200, gin.H{
				"progress": treeService.Tree.Root.Value,
			})

		})

		progress.PUT("", func(context *gin.Context) {
			treeService := services.DB().ProgressTree()

			if treeService.Tree.Root.Steps == nil || len(treeService.Tree.Root.Steps) == 0 {
				context.JSON(400, gin.H{"error": errors.New("no steps left").Error()})
			}

			tree, err := treeService.Next(treeService.Tree.Root.Steps[0].Value)

			if err != nil {
				return
			}

			context.JSON(200, gin.H{
				"progress": tree.Root.Value,
			})

		})
	}
}
