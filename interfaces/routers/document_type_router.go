package routers

import (
	DocumentTypeV1Point00 "cargo-rest-api/interfaces/handler/v1.0/document_type"
	"cargo-rest-api/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func documentTypeRoutes(e *gin.Engine, r *Router, rg *RouterAuthGateway) {
	DocumentTypeV1 := DocumentTypeV1Point00.NewDocumentTypes(r.dbService.DocumentType)

	guard := middleware.Guard(rg.authGateway)
	v1 := e.Group("/api/v1/external")

	v1.GET("/documentTypes", guard.Authenticate(), DocumentTypeV1.GetDocumentTypes)
	v1.POST("/documentTypes", guard.Authenticate(), DocumentTypeV1.SaveDocumentType)
	v1.GET("/documentTypes/:uuid", guard.Authenticate(), DocumentTypeV1.GetDocumentType)
	v1.PUT("/documentTypes/:uuid", guard.Authenticate(), DocumentTypeV1.UpdateDocumentType)
	v1.DELETE("/documentTypes/:uuid", guard.Authenticate(), DocumentTypeV1.DeleteDocumentType)
}
