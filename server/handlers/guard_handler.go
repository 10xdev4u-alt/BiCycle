package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GuardScanRequest struct {
	TicketQR string `json:"ticket_qr" binding:"required"`
}

// ScanTicket handles the scanning of a student's ticket by a guard.
// @Summary Scan a student ticket
// @Description Receives a ticket QR from a student, validates it, and returns booking details.
// @Tags Guards
// @Accept  json
// @Produce  json
// @Param   body  body   GuardScanRequest  true  "Ticket QR Payload"
// @Success 200 {object} gin.H{"type": string, "student": object, "bike": object, "photos": object, "duration": int, "guard_code": string}
// @Failure 400 {object} gin.H{"error": string}
// @Router /api/v1/guard/scan [post]
func ScanTicket(c *gin.Context) {
	var json GuardScanRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// 1. Decode JWT to get ticket details
	// 2. Verify signature and expiry
	// 3. Return booking details

	// For now, assume success and return mock data as per Task.md
	c.JSON(http.StatusOK, gin.H{
		"type": "pickup",
		"student": gin.H{
			"name":   "Prince",
			"reg_no": "2023CS0113",
		},
		"bike": gin.H{
			"display_id": "Cycle-07",
		},
		"photos": gin.H{
			"pickup_url": "/photos/pickups/bk_7f3a9c2e.jpg",
			"return_url": nil,
		},
		"duration":   0,
		"guard_code": "3847",
	})
}
