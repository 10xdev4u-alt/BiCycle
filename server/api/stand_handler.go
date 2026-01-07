package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/princetheprogrammerbtw/TheBiCycleApp/server/crypto"
)

type ScanStandRequest struct {
	EncryptedStandQR string `json:"encrypted_stand_qr" binding:"required"`
}

type StandQRPayload struct {
	Type      string `json:"type"`
	ID        string `json:"id"`
	Location  string `json:"location"`
	Timestamp int64  `json:"timestamp"`
}

// ScanStand handles the scanning of a stand's QR code.
func (a *API) ScanStand(c *gin.Context) {
	var req ScanStandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decrypted, err := crypto.Decrypt(req.EncryptedStandQR, a.Cfg)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid QR code."})
		return
	}

	var payload StandQRPayload
	if err := json.Unmarshal([]byte(decrypted), &payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid QR code payload."})
		return
	}

	if payload.Type != "stand" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid QR code type."})
		return
	}

	var standName string
	err = a.DB.QueryRow("SELECT name FROM stands WHERE id = $1", payload.ID).Scan(&standName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stand not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stand_id":   payload.ID,
		"stand_name": standName,
		"message":    "Stand verified. Now scan a bike.",
	})
}
