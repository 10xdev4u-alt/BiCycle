package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScanBookRequest struct {
	EncryptedBikeQR string `json:"encrypted_bike_qr" binding:"required"`
	CurrentStand    string `json:"current_stand" binding:"required"`
}

// ScanAndBookCycle handles the scanning of a bike's QR code and attempts to book it.
// @Summary Scan and book a cycle
// @Description Receives an encrypted bike QR and the current stand ID to book a cycle.
// @Tags Cycles
// @Accept  json
// @Produce  json
// @Param   body  body   ScanBookRequest  true  "Scan and Book Payload"
// @Success 200 {object} gin.H{"success": bool, "booking_id": string, "bike_display_id": string, "message": string}
// @Failure 400 {object} gin.H{"error": string}
// @Failure 409 {object} gin.H{"error": string, "message": string}
// @Router /api/v1/cycles/scan-book [post]
func ScanAndBookCycle(c *gin.Context) {
	var json ScanBookRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// 1. Decrypt bike QR to get bike_uuid
	// 2. Verify bike.current_stand === json.CurrentStand
	// 3. Check bike.status === 'available' (Redis)
	// 4. Verify user has no active booking
	// 5. ATOMIC: Book bike in Redis

	// For now, assume success and return mock data as per Task.md
	c.JSON(http.StatusOK, gin.H{
		"success":         true,
		"booking_id":      "bk_7f3a9c2e",
		"bike_display_id": "Cycle-07",
		"message":         "Bike booked. Take pickup photo.",
	})
}

type PhotoPickupRequest struct {
	BookingID    string `json:"booking_id" binding:"required"`
	PhotoBase64  string `json:"photo_base64" binding:"required"`
}

// UploadPickupPhoto handles the upload of a pickup photo for a booking.
// @Summary Upload pickup photo
// @Description Receives a booking ID and a base64 encoded photo of the bike at pickup.
// @Tags Cycles
// @Accept  json
// @Produce  json
// @Param   body  body   PhotoPickupRequest  true  "Photo Pickup Payload"
// @Success 200 {object} gin.H{"success": bool, "ticket_qr": string, "guard_code": string, "expires_in": int}
// @Failure 400 {object} gin.H{"error": string}
// @Router /api/v1/cycles/photo-pickup [post]
func UploadPickupPhoto(c *gin.Context) {
	var json PhotoPickupRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO:
	// 1. Verify booking belongs to user & status='booked'
	// 2. Store photo: pickups/{booking_id}.jpg
	// 3. Optional: AI check (MobileNet) -> "Bike detected: 95%"
	// 4. AI fails -> Reject: "Retake photo"
	// 5. Success -> Generate Pickup Ticket

	// For now, assume success and return mock data as per Task.md
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"ticket_qr":   "<base64>",
		"guard_code":  "3847",
		"expires_in":  300,
	})
}
