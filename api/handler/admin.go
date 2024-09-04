package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	md "github.com/mirjalilova/api-gateway_blacklist/api/middleware"
	pb "github.com/mirjalilova/api-gateway_blacklist/internal/genproto/black_list"

	// rd "github.com/mirjalilova/api-gateway_blacklist/pkg/helper"
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
		c.JSON(400, gin.H{"error": err})
		return
	}

	// Clear relevant cache keys
	// h.Redis.Del(c, "allhr:")

	slog.Error("HR approved successfully")
	c.JSON(200, gin.H{"message": "HR approved successfully"})
}

// @Summary 		Get hr
// @Description     Get hr
// @Tags       	    Admin
// @Accept 			json
// @Produce 		json
// @Security 		BearerAuth
// @Param           hr_id query string true "HR Id"
// @Success 200 {object} black_list.Hr
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/hr/{hr_id} [get]
func (h *HandlerStruct) GetHR(c *gin.Context) {
	userId := c.Query("hr_id")

	req := &pb.GetById{
        Id: userId,
    }

	res, err := h.Clients.AdminClient.GetHRById(context.Background(), req)
	if err!= nil {
        slog.Error("failed to get hr: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

	slog.Info("HR retrieved successfully")
	c.JSON(http.StatusOK, res)
}

// @Summary 		Get hr list
// @Description     Get hr list
// @Tags       	    Admin
// @Accept 			json
// @Produce 		json
// @Security 		BearerAuth
// @Param           limit query int false "Limit"
// @Param           offset query int false "Offset"
// @Success 200 {object} black_list.GetAllHRRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/hr_list [get]
func (h *HandlerStruct) ListHR(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	limitValue := 10
	offsetValue := 1

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

	// cacheKey := "allhr:"

	// res := pb.GetAllHRRes{}

	// err := rd.GetCachedData(c, h.Redis, cacheKey, &res)
	// if err == nil {
	// 	slog.Info("Weekly data retrieved from cache")
	// 	c.JSON(200, res)
	// 	return
	// }

	resp, err := h.Clients.AdminClient.ListHR(context.Background(), req)
	if err != nil {
		slog.Error("Error getting HR list: ", err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	// res = *resp

	// rd.CacheData(c, h.Redis, cacheKey, res)

	slog.Info("HR list retrieved successfully")
	c.JSON(200, resp)
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
		c.JSON(400, gin.H{"error": err})
		return
	}

	// Clear relevant cache keys
	// h.Redis.Del(c, "allhr:")

	slog.Info("HR deleted successfully")
	c.JSON(200, "HR deleted successfully")
}

func getuserId(ctx *gin.Context) string {
	user_id, err := md.GetUserId(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return ""
	}
	return user_id
}

// @Summary 		Get all users
// @Description     Get all users
// @Tags       	    Admin
// @Accept 			json
// @Produce 		json
// @Security 		BearerAuth
// @Param           username query string false "UserName"
// @Param           full_name query string false "Full Name"
// @Param           limit query int false "Limit"
// @Param           offset query int false "Offset"
// @Success 200 {object} black_list.ListUserRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/users [get]
func (h *HandlerStruct) GetAllUsers(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")
	username := c.Query("username")
	fullName := c.Query("full_name")

	limitValue := 10
	offsetValue := 1

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

	req := &pb.ListUserReq{
		Username: username,
		FullName: fullName,
		Filter: &pb.Filter{
			Limit:  int32(limitValue),
			Offset: int32(offsetValue),
		},
	}

	res, err := h.Clients.AdminClient.GetAllUsers(context.Background(), req)
	if err != nil {
		slog.Error("failed to get all users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary 			Change role
// @Description		 	Change role of a user
// @Tags 				Admin
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param  				user body black_list.ChangeRoleReq true "Employee"
// @Success 200			{object} string "Employee updated successfully"
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router 				/admin/change_role [PUT]
func (h *HandlerStruct) ChangeRole(c *gin.Context) {
	var req pb.ChangeRoleReq

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err := h.Clients.AdminClient.ChangeRole(context.Background(), &req)
	if err != nil {
		slog.Error("failed to update user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	slog.Info("Role updated successfully")
	c.JSON(http.StatusOK, "Role updated successfully")
}

// @Summary 		Get user
// @Description     Get user
// @Tags       	    Admin
// @Accept 			json
// @Produce 		json
// @Security 		BearerAuth
// @Param           user_id query string true "User Id"
// @Success 200 {object} black_list.UserRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /admin/user/{user_id} [get]
func (h *HandlerStruct) GetUser(c *gin.Context) {
	userId := c.Query("user_id")

	req := &pb.GetById{
        Id: userId,
    }

	res, err := h.Clients.AdminClient.GetUserById(context.Background(), req)
	if err!= nil {
        slog.Error("failed to get user: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
        return
    }

	slog.Info("User retrieved successfully")
	c.JSON(http.StatusOK, res)
}