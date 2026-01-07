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

type ApprovePickupRequest struct {
	TicketToken string `json:"ticket_token" binding:"required"`
	GuardID     string `json:"guard_id" binding:"required"`
	Action      string `json:"action" binding:"required"`
}

// ApprovePickup handles the guard's approval of a bike pickup.
// @Summary Approve a bike pickup
// @Description Receives a ticket token, guard ID, and an action ('approve' or 'reject') to confirm a bike pickup.
// @Tags Guards
// @Accept  json
// @Produce  json
// @Param   body  body   ApprovePickupRequest  true  "Approve Pickup Payload"
// @Success 200 {object} gin.H{"success": bool}
// @Failure 400 {object} gin.H{"error": string}
// @Router /api/v1/guard/approve-pickup [post]
func ApprovePickup(c *gin.Context) {
	var json ApprovePickupRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// 1. Verify Guard JWT
	// 2. Update booking: status='picked_up', picked_up_at=NOW()
	// 3. Update bike: status='in_use'
	// 4. Clear Redis expiry

	// For now, assume success
	c.JSON(http.StatusOK, gin.H{"success": true})
}

type ApproveReturnRequest struct {
	TicketToken string `json:"ticket_token" binding:"required"`
	GuardID     string `json:"guard_id" binding:"required"`
	Action      string `json:"action" binding:"required"`
	Damage      bool   `json:"damage"`
}

// ApproveReturn handles the guard's approval of a bike return.
// @Summary Approve a bike return
// @Description Receives a ticket token, guard ID, an action ('approve' or 'reject'), and a damage flag to confirm a bike return.
// @Tags Guards
// @Accept  json
// @Produce  json
// @Param   body  body   ApproveReturnRequest  true  "Approve Return Payload"
// @Success 200 {object} gin.H{"success": bool}
// @Failure 400 {object} gin.H{"error": string}
// @Router /api/v1/guard/approve-return [post]
func ApproveReturn(c *gin.Context) {
	var json ApproveReturnRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// 1. Update booking: status='returned', returned_at=NOW()
	// 2. Update bike: status='available', current_stand='stand_mg_12'
	// 3. Clear Redis: DEL user:active:2023CS0113
	// 4. INCR stand:available:stand_mg_12
	// 5. Check violations (late, fast, no-show)

	// For now, assume success
	c.JSON(http.StatusOK, gin.H{"success": true})
}
