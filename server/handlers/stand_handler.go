package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScanStandRequest struct {
	EncryptedStandQR string `json:"encrypted_stand_qr" binding:"required"`
}

// ScanStand handles the scanning of a stand's QR code.
// @Summary Scan a bike stand QR code
// @Description Receives an encrypted QR code from a bike stand, validates it, and returns stand details.
// @Tags Stands
// @Accept  json
// @Produce  json
// @Param   body  body   ScanStandRequest  true  "Encrypted Stand QR Payload"
// @Success 200 {object} gin.H{"stand_id": string, "stand_name": string, "message": string}
// @Failure 400 {object} gin.H{"error": string}
// @Router /api/v1/stand/scan [post]
func ScanStand(c *gin.Context) {
	var json ScanStandRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Decrypt and validate the QR code
	// For now, we'll assume the QR is valid and return mock data as per Task.md

	c.JSON(http.StatusOK, gin.H{
		"stand_id":   "stand_bh_01",
		"stand_name": "Boys Hostel - Zone A",
		"message":    "Stand verified. Now scan a bike.",
	})
}
