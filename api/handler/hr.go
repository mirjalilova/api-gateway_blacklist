package handler

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	pb "github.com/mirjalilova/api-gateway_blacklist/internal/genproto/black_list"
	"golang.org/x/exp/slog"
)

// @Summary 			CREATE EMPLOYEE
// @Description		 	This api create employee
// @Tags 				Employee
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param data 			body black_list.EmployeeCreate true "Employee"
// @Success 201 		{object} string "Employee created successfully"
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router 				/employee/create [POST]
func (h *HandlerStruct) CreateEmployee(c *gin.Context) {
	req := &pb.EmployeeCreate{}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("Failed to bind JSON: ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := h.Clients.HrClient.Create(context.Background(), req)
	if err != nil {
		slog.Error("Error while create employee: ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Employee created successfully")
	c.JSON(200, "Employee created successfully")
}

// @Summary 			GET EMPLOYEE
// @Description		 	This api get employee by id
// @Tags 				Employee
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    employee_id path string true "Employee ID"
// @Success 200			{object} black_list.Employee
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router 				/employee/{employee_id} [GET]
func (h *HandlerStruct) GetEmployee(c *gin.Context) {
	id := c.Param("employee_id")

	req := &pb.GetById{
		Id: id,
	}

	employee, err := h.Clients.HrClient.Get(context.Background(), req)
	if err != nil {
		slog.Error("Error while getting employee: ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Employee getting successfully")
	c.JSON(200, gin.H{"employee": employee})
}

// @Summary             GET ALL EMPLOYEES
// @Description         This API gets all employees
// @Tags                Employee
// @Accept              json
// @Produce             json
// @Security            BearerAuth
// @Param               position query string false "Position"
// @Param               limit query int false "Limit"
// @Param               offset query int false "Offset"
// @Success 200         {object} black_list.ListEmployeeRes
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router              /employee/all [GET]
func (h *HandlerStruct) GetAllEmployees(c *gin.Context) {
	position := c.Query("position")
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

	req := &pb.ListEmployeeReq{
		Position: position,
		Filter: &pb.Filter{
			Limit:  int32(limitValue),
			Offset: int32(offsetValue),
		},
	}

	employees, err := h.Clients.HrClient.GetAll(context.Background(), req)
	if err != nil {
		slog.Error("Error while getting employees")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Employees getting successfully")
	c.JSON(200, employees)
}

// @Summary 			UPDATES EMPLOYEE
// @Description		 	This api updatedes employee
// @Tags 				Employee
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param  				employee_id  path string true "Employee Id"
// @Param  				employee  body black_list.UpdateReqBody true "Employee"
// @Success 200			{object} string "Employee updated successfully"
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router 				/employee/update/{employee_id} [PUT]
func (h *HandlerStruct) UpdateEmployee(c *gin.Context) {
	id := c.Query("employee_id")

	var body pb.UpdateReqBody

	if err := c.ShouldBindJSON(&body); err != nil {
		slog.Error("Failed to bind JSON: ", err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	req := &pb.UpdateReq{
		Id:       id,
		Position: body.Position,
		HrId:     body.HrId,
	}
	_, err := h.Clients.HrClient.Update(context.Background(), req)
	if err != nil {
		slog.Error("Error while updating employee")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	slog.Info("Employee updated successfully")
	c.JSON(200, gin.H{"message": "Employee updated successfully"})
}

// @Summary 			DELETES EMPLOYEE
// @Description		 	This api deletes employee
// @Tags 				Employee
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param  				employee_id  path string true "Employee"
// @Success 200			{object} string "Employee deleted successfully"
// @Failure 400         {string} Error "Bad Request"
// @Failure 404         {string} Error "Not Found"
// @Failure 500         {string} Error "Internal Server Error"
// @Router 				/employee/{employee_id} [DELETE]
func (h *HandlerStruct) DeleteEmployee(c *gin.Context) {
	id := c.Param("employee_id")
	
	req := &pb.GetById{
		Id: id,
	}

	_, err := h.Clients.HrClient.Delete(context.Background(), req)
	if err != nil {
		slog.Error("Error while deleting employee")
		c.JSON(400, gin.H{"error": err})
		return
	}

	slog.Info("Employee deleted successfully")
	c.JSON(200, gin.H{"message": "Employee deleted successfully"})
}
