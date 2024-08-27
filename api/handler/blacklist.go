package handler

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/mirjalilova/api-gateway_blacklist/internal/genproto/black_list"
)

// @Summary 			Add employee to black list
// @Description		 	Add employee to black list
// @Tags 				Black List
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param data 			body black_list.BlackListBody true "Add employee"
// @Success 200 		{string} string "Add employee to blacklist successfully"
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router 				/blacklist/add [POST]
func (h *HandlerStruct) AddEmployee(c *gin.Context) {
	var body pb.BlackListBody
	if err := c.ShouldBindJSON(&body); err != nil {
		slog.Error("Failed to bind JSON")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req := &pb.BlackListCreate{
		EmployeeId: body.EmployeeId,
		Reason:     body.Reason,
		AddedBy:    getuserId(c),
	}
	_, err := h.Clients.BlacklistClient.Add(context.Background(), req)
	if err != nil {
		slog.Error("Error while adding employee to blacklist")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Clear relevant cache keys
	h.Redis.Del(c, "daily:", "weekly:", "monthly:")

	slog.Info("Add employee to blacklist successfully")
	c.JSON(200, gin.H{"message": "Add employee to blacklist successfully"})
}

// @Summary             Get all employee from blacklist
// @Description         Get all employee from blacklist
// @Tags                Black List
// @Accept              json
// @Produce             json
// @Security            BearerAuth
// @Param               limit query int false "Limit"
// @Param               offset query int false "Offset"
// @Success 200         {object} black_list.GetAllBlackListRes
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router              /blacklist/all [GET]
func (h *HandlerStruct) GetAll(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	limitValue := 10
	offsetValue := 0

	if limit != "" {
		parsedLimit, err := strconv.Atoi(limit)
		if err != nil {
			slog.Error("Invalid limit value")
			c.JSON(400, gin.H{"error": "Invalid limit value"})
			return
		}
		limitValue = parsedLimit
	}

	if offset != "" {
		parsedOffset, err := strconv.Atoi(offset)
		if err != nil {
			slog.Error("Invalid offset value")
			c.JSON(400, gin.H{"error": "Invalid offset value"})
			return
		}
		offsetValue = parsedOffset
	}

	req := &pb.Filter{
		Limit:  int32(limitValue),
		Offset: int32(offsetValue),
	}

	lists, err := h.Clients.BlacklistClient.GetAll(context.Background(), req)
	if err != nil {
		slog.Error("Error while getting blacklist")
		c.JSON(400, gin.H{"error": err})
		return
	}

	slog.Info("Employee getting successfully from blacklist")
	c.JSON(200, lists)
}

// @Summary 			Remove employee from blacklist
// @Description		 	Remove employee from blacklist
// @Tags 				Black List
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    employee_id path string true "Employee Id"
// @Success 200 		{string} string "Removed employee successfully from blacklist"
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router 				/blacklist/remove/{employee_id} [DELETE]
func (h *HandlerStruct) RemoveEmployee(c *gin.Context) {
	id := c.Param("employee_id")

	req := &pb.RemoveReq{
		EmployeeId: id,
		AddedBy:    getuserId(c),
	}

	_, err := h.Clients.BlacklistClient.Remove(context.Background(), req)
	if err != nil {
		slog.Error("Error while remove employee from blacklist")
		c.JSON(400, gin.H{"error": err})
		return
	}

	// Clear relevant cache keys
	h.Redis.Del(c, "daily:", "weekly:", "monthly:")

	slog.Info("Removed employee successfully from blacklist")
	c.JSON(200, "Removed employee successfully from blacklist")
}

// @Summary             Monitoring daily blacklist
// @Description         Monitoring daily blacklist
// @Tags                Monitoring
// @Accept              json
// @Produce             json
// @Security            BearerAuth
// @Success 200         {object} black_list.Reports
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router              /blacklist/daily [GET]
func (h *HandlerStruct) GetDaily(c *gin.Context) {
	// cacheKey := "daily:"

	// res := pb.Reports{}

	// slog.Info("ssssssssssssssss")

	// err := rd.GetCachedData(c, h.Redis, cacheKey, &res)
	// slog.Info("ssssssssssssssssssssssss, err:", err)
	// if err == nil {
	// 	slog.Info("Weekly data retrieved from cache")
	// 	c.JSON(200, res)
	// 	return
	// }

	resp, err := h.Clients.BlacklistClient.MonitoringDailyReport(context.Background(), &pb.Void{})
	slog.Info("ssssssssssssssssssssssss, err:", err)
	if err != nil {
		slog.Error("Error while getting daily blacklist")
		c.JSON(400, gin.H{"error": err})
		return
	}

	// slog.Info("ssssssssssssssssssssssss, resp:", resp)
	// res = *resp
	// slog.Info("ssssssssssssssssssssssss, res:", res)

	// rd.CacheData(c, h.Redis, cacheKey, res)
	// slog.Info("ssssssssssssssssssssssss, res:", res)

	slog.Info("Daily blacklist retrieved successfully")
	c.JSON(200, resp)
}

// @Summary             Monitoring weekly blacklist
// @Description         Monitoring weekly blacklist
// @Tags                Monitoring
// @Accept              json
// @Produce             json
// @Security            BearerAuth
// @Success 200         {object} black_list.Reports
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router              /blacklist/weekly [GET]
func (h *HandlerStruct) GetWeekly(c *gin.Context) {
	// cacheKey := "weekly:"

	// res := pb.Reports{}

	// err := rd.GetCachedData(c, h.Redis, cacheKey, &res)
	// if err == nil {
	// 	slog.Info("Weekly data retrieved from cache")
	// 	c.JSON(200, res)
	// 	return
	// }

	resp, err := h.Clients.BlacklistClient.MonitoringWeeklyReport(context.Background(), &pb.Void{})
	if err != nil {
		slog.Error("Error while getting daily blacklist", err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	// res = *resp

	// rd.CacheData(c, h.Redis, cacheKey, res)

	slog.Info("Daily blacklist retrieved successfully")
	c.JSON(200, resp)
}

// @Summary             Monitoring monthly blacklist
// @Description         Monitoring monthly blacklist
// @Tags                Monitoring
// @Accept              json
// @Produce             json
// @Security            BearerAuth
// @Success 200         {object} black_list.Reports
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router              /blacklist/monthly [GET]
func (h *HandlerStruct) GetMonthly(c *gin.Context) {
	// cacheKey := "monthly:"

	// res := pb.Reports{}

	// err := rd.GetCachedData(c, h.Redis, cacheKey, &res)
	// if err == nil {
	// 	slog.Info("Weekly data retrieved from cache")
	// 	c.JSON(200, res)
	// 	return
	// }

	resp, err := h.Clients.BlacklistClient.MonitoringMonthlyReport(context.Background(), &pb.Void{})
	if err != nil {
		slog.Error("Error while getting daily blacklist")
		c.JSON(400, gin.H{"error": err})
		return
	}

	// res = *resp

	// rd.CacheData(c, h.Redis, cacheKey, res)

	slog.Info("Daily blacklist retrieved successfully")
	c.JSON(200, resp)
}
