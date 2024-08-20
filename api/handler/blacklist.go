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
// @Param data 			body pb.BlackListBody true "Add employee"
// @Success 201 		{object} string "Add employee to blacklist successfully"
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
		Reason: body.Reason,
		AddedBy: getuserId(c),
	}
	_, err := h.Clients.BlacklistClient.Add(context.Background(), req)
	if err != nil {
		slog.Error("Error while adding employee to blacklist")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

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
// @Success 200         {object} pb.GetAllBlackListRes
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
		Limit: int32(limitValue),
		Offset: int32(offsetValue),
	}

	lists, err := h.Clients.BlacklistClient.GetAll(context.Background(), req)
	if err != nil {
		slog.Error("Error while getting blacklist")
		c.JSON(400, gin.H{"error": err.Error()})
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
// @Success 201 		{object} "Removed employee successfully from blacklist"
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router 				/blacklist/remove/{employee_id} [DELETE]
func (h *HandlerStruct) RemoveEmployee(c *gin.Context) {
	id := c.Param("employee_id")

	req := &pb.RemoveReq{
		EmployeeId: id,
		AddedBy: getuserId(c),
	}

	_, err := h.Clients.BlacklistClient.Remove(context.Background(), req)
	if err != nil {
		slog.Error("Error while remove employee from blacklist")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Removed employee successfully from blacklist")
	c.JSON(200, "Removed employee successfully from blacklist")
}