package v1

import (
	"net/http"

	"github.com/artomsopun/mortgage/mortgage-api/internal/service"
	"github.com/artomsopun/mortgage/mortgage-api/pkg/types"
	"github.com/labstack/echo/v4"
)

func (h *Handler) initBanksRoutes(api *echo.Group) {
	banks := api.Group("/banks")
	{
		banks.GET("", h.getAllBanks)
		banks.POST("/mortgage", h.getMortgage)
		profile := banks.Group("/profile", h.checkAuth)
		{
			profile.GET("", h.getBanksByUserID)
			profile.POST("", h.createBank)
			profile.PUT("", h.updateBank)
			profile.DELETE("/:id", h.deleteBank)
		}
	}
}

type bankInfo struct {
	ID         types.BinaryUUID `json:"id"`
	Title      string           `json:"title"`
	Rate       float64          `json:"rate"`
	MaxLoan    uint             `json:"maxLoan"`
	MinPayment uint             `json:"minPayment"`
	LoanTerm   uint             `json:"loanTerm"`
	UserID     types.BinaryUUID `json:"userId"`
}

func (h *Handler) getAllBanks(c echo.Context) error {
	banksServ, err := h.services.Banks.GetAllBanks()
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var banks []bankInfo
	for _, bank := range banksServ {
		banks = append(banks, bankInfo{
			ID:         bank.ID,
			Title:      bank.Title,
			Rate:       bank.Rate,
			MaxLoan:    bank.MaxLoan,
			MinPayment: bank.MinPayment,
			LoanTerm:   bank.LoanTerm,
			UserID:     bank.UserID,
		})
	}
	return c.JSON(http.StatusOK, banks)
}

func (h *Handler) getBanksByUserID(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	banksServ, err := h.services.Banks.GetBanksByUserID(userID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var banks []bankInfo
	for _, bank := range banksServ {
		banks = append(banks, bankInfo{
			ID:         bank.ID,
			Title:      bank.Title,
			Rate:       bank.Rate,
			MaxLoan:    bank.MaxLoan,
			MinPayment: bank.MinPayment,
			LoanTerm:   bank.LoanTerm,
			UserID:     bank.UserID,
		})
	}
	return c.JSON(http.StatusOK, banks)
}

type bankInput struct {
	Title      string  `json:"title"`
	Rate       float64 `json:"rate"`
	MaxLoan    uint    `json:"maxLoan"`
	MinPayment uint    `json:"minPayment"`
	LoanTerm   uint    `json:"loanTerm"`
}

func (h *Handler) createBank(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var input bankInput
	if err := c.Bind(&input); err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.Banks.CreateBank(service.Bank{
		Title:      input.Title,
		Rate:       input.Rate,
		MaxLoan:    input.MaxLoan,
		MinPayment: input.MinPayment,
		LoanTerm:   input.LoanTerm,
		UserID:     userID,
	}); err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusCreated, "bank created")
}

type bankInputUpdate struct {
	ID         types.BinaryUUID `json:"id"`
	Title      string           `json:"title"`
	Rate       float64          `json:"rate"`
	MaxLoan    uint             `json:"maxLoan"`
	MinPayment uint             `json:"minPayment"`
	LoanTerm   uint             `json:"loanTerm"`
	UserID     types.BinaryUUID `json:"userId"`
}

func (h *Handler) updateBank(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	var input bankInputUpdate
	if err := c.Bind(&input); err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := h.services.Banks.UpdateBank(service.Bank{
		ID:         input.ID,
		Title:      input.Title,
		Rate:       input.Rate,
		MaxLoan:    input.MaxLoan,
		MinPayment: input.MinPayment,
		LoanTerm:   input.LoanTerm,
		UserID:     userID,
	}); err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusOK, "bank updated")
}

func (h *Handler) deleteBank(c echo.Context) error {
	userID, err := getUserID(c)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	bankIDStr := c.Param("id")

	bankID := types.ParseUUID(bankIDStr)
	err = h.services.Banks.DeleteBank(userID, bankID)
	if err != nil {
		return newResponse(c, http.StatusInternalServerError, err.Error())
	}
	return newResponse(c, http.StatusOK, "bank deleted")
}

type calculateInput struct {
	Loan    uint             `json:"loan"`
	Payment uint             `json:"payment"`
	BankID  types.BinaryUUID `json:"bankId"`
}

func (h *Handler) getMortgage(c echo.Context) error {
	var input calculateInput
	if err := c.Bind(&input); err != nil {
		return newResponse(c, http.StatusBadRequest, err.Error())
	}

	result := h.services.Banks.CalculateMortgage(service.CalculateInput{
		Loan:    input.Loan,
		Payment: input.Payment,
		BankID:  input.BankID,
	})
	return newResponse(c, http.StatusOK, result)
}
