package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScanBookRequest struct {
	EncryptedBikeQR string `json:"encrypted_bike_qr" binding:"required"`
	CurrentStand    string `json:"current_stand" binding:"required"`
}

// ScanAndBookCycle handles the scanning of a bike's QR code and attempts to book it.
func (a *API) ScanAndBookCycle(c *gin.Context) {
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
func (a *API) UploadPickupPhoto(c *gin.Context) {
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
