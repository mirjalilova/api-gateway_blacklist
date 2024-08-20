package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	md "github.com/mirjalilova/api-gateway_blacklist/api/middleware"
	pb "github.com/mirjalilova/api-gateway_blacklist/internal/genproto/black_list"
	"golang.org/x/exp/slog"
)

// @Summary 			Create a new hr
// @Description		    Create a new hr
// @Tags 				Admin
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    user_id path string true "USER ID"
// @Success 200 {object} string "HR created successfully"
// @Failure 400 {object} string "error": "error message"
// @Router /admin/approve/{user_id} [post]
func (h *HandlerStruct) Approve(c *gin.Context) {
	user_id := c.Param("user_id")
	approved_by := getuserId(c)

	req := &pb.CreateHR{
		UserId:     user_id,
		ApprovedBy: approved_by,
	}

	_, err := h.Clients.AdminClient.Approve(context.Background(), req)

	if err != nil {
		slog.Error("Error approving HR: ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slog.Error("HR approved successfully")
	c.JSON(200, gin.H{"message": "HR approved successfully"})
}

// @Summary 		Get hr list
// @Description     Get hr list
// @Tags       	    Admin
// @Accept 			json
// @Produce 		json
// @Security 		BearerAuth
// @Param           limit query int false "Limit"
// @Param           offset query int false "Offset"
// @Success 200 {object} auth.GetAllHRRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/hr_list [get]
func (h *HandlerStruct) ListHR(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	limitValue := 10
	offsetValue := 0

	if limit != "" {
		parsedLimit, err := strconv.Atoi(limit)
		if err != nil {
			slog.Error("Invalid limit value", err)
			c.JSON(400, gin.H{"error": "Invalid limit value"})
			return
		}
		limitValue = parsedLimit
	}

	if offset != "" {
		parsedOffset, err := strconv.Atoi(offset)
		if err != nil {
			slog.Error("Invalid offset value", err)
			c.JSON(400, gin.H{"error": "Invalid offset value"})
			return
		}
		offsetValue = parsedOffset
	}

	req := &pb.Filter{
		Limit:  int32(limitValue),
		Offset: int32(offsetValue),
	}

	admins, err := h.Clients.AdminClient.ListHR(context.Background(), req)
	if err != nil {
		slog.Error("Error getting HR list: ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slog.Info("HR list retrieved successfully")
	c.JSON(200, admins)
}

// @Summary			 Delete HR
// @Description 	 Delete HR
// @Tags			 Admin
// @Accept 			 json
// @Produce 		 json
// @Security		 BearerAuth
// @Param hr_id query string true "HR ID"
// @Success 200 {object} string "HR deleted successfully"
// @Failure 400 {object} string "error": "error message"
// @Router /admin/delete_hr/{hr_id} [delete]
func (h *HandlerStruct) DeleteHR(c *gin.Context) {
	hr_id := c.Query("hr_id")

	req := &pb.GetById{
		Id: hr_id,
	}

	_, err := h.Clients.AdminClient.Delete(context.Background(), req)
	if err != nil {
		slog.Error("Error deleting Genetic Data: ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slog.Info("HR deleted successfully")
	c.JSON(200, "HR deleted successfully")
}

func getuserId(ctx *gin.Context) string {
	user_id, err := md.GetUserId(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return ""
	}
	return user_id
}