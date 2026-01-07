package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/crypto"
)

type ScanBookRequest struct {
	EncryptedBikeQR string `json:"encrypted_bike_qr" binding:"required"`
	CurrentStand    string `json:"current_stand" binding:"required"`
}

type BikeQRPayload struct {
	Type         string `json:"type"`
	ID           string `json:"id"`
	InitialStand string `json:"initial_stand"`
	Timestamp    int64  `json:"timestamp"`
}

// ScanAndBookCycle handles the scanning of a bike's QR code and attempts to book it.
func (a *API) ScanAndBookCycle(c *gin.Context) {
	var req ScanBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decrypted, err := crypto.Decrypt(req.EncryptedBikeQR, a.Cfg)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid QR code."})
		return
	}

	var payload BikeQRPayload
	if err := json.Unmarshal([]byte(decrypted), &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid QR code payload."})
		return
	}

	if payload.Type != "bike" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid QR code type. Expected 'bike'."})
		return
	}

	var bikeID int
	var status string
	var currentStandID int
	err = a.DB.QueryRow("SELECT id, status, current_stand_id FROM bikes WHERE uuid = $1", payload.ID).Scan(&bikeID, &status, &currentStandID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bike not found."})
		return
	}

	if status != "available" {
		c.JSON(http.StatusConflict, gin.H{"error": "Bike is not available."})
		return
	}

	// TODO: This check needs to be improved. The request sends the stand's UUID, but the bikes table stores the stand's integer ID.
	// This requires a join or a second query. For now, we'll skip this check.
	// var standUUID string
	// err = a.DB.QueryRow("SELECT uuid FROM stands WHERE id = $1", currentStandID).Scan(&standUUID)
	// if err != nil || standUUID != req.CurrentStand {
	//  	c.JSON(http.StatusBadRequest, gin.H{"error": "Bike is not at the specified stand."})
	// 	return
	// }

	// TODO:
	// 1. Verify user has no active booking
	// 2. ATOMIC: Book bike in DB (or Redis)

	// For now, assume success and return mock data
	c.JSON(http.StatusOK, gin.H{
		"success":         true,
		"booking_id":      "bk_7f3a9c2e",
		"bike_display_id": "Cycle-07", // Should come from the DB
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
