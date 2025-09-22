package handlers

import (
	"ave_project/internal/usecase/cipher"
	"encoding/json"
	"net/http"
)

type CipherHandler struct {
	Usecase *cipher.CipherUsecase
}

type CipherRequest struct {
	Text  string `json:"text"`
	Shift int    `json:"shift"`
}

func (h *CipherHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	var req CipherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result := h.Usecase.Encrypt(req.Text, req.Shift)
	json.NewEncoder(w).Encode(map[string]string{"encrypted": result})
}

func (h *CipherHandler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	var req CipherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	result := h.Usecase.Decrypt(req.Text, req.Shift)

	// TODO: Првоверить статус ответа
	//w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"decrypted": result})
}
