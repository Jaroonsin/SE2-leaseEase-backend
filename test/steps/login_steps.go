package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"LeaseEase/internal/dtos"
	"LeaseEase/internal/handlers"
	"LeaseEase/mocks"

	"github.com/cucumber/godog"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

var (
	requestBody         []byte
	recorder            *httptest.ResponseRecorder
	mockAuthService     *mocks.AuthService
	mockService         *mocks.Service
	mockPropertyService *mocks.PropertyService // not used in this example
	mockLesseeService   *mocks.LesseeService   // not used in this example
	mockReviewService   *mocks.ReviewService   // not used in this example
	mockPaymentService  *mocks.PaymentService  // not used in this example
	mockLessorService   *mocks.LessorService   // not used in this example
	mockUserService     *mocks.UserService     // not used in this example
	handler             handlers.Handler
)

// -------------------------------
// Step 1: Initialize Mocks and Handler
// -------------------------------

func setupMocks() {
	mockAuthService = new(mocks.AuthService)
	mockPropertyService = new(mocks.PropertyService) // not used in this example
	mockLesseeService = new(mocks.LesseeService)     // not used in this example
	mockReviewService = new(mocks.ReviewService)     // not used in this example
	mockPaymentService = new(mocks.PaymentService)   // not used in this example
	mockLessorService = new(mocks.LessorService)     // not used in this example
	mockUserService = new(mocks.UserService)         // not used in this example

	mockService = new(mocks.Service)

	mockService.On("Auth").Return(mockAuthService)
	mockService.On("Property").Return(mockPropertyService) // not used in this example
	mockService.On("Lessee").Return(mockLesseeService)     // not used in this example
	mockService.On("Review").Return(mockReviewService)     // not used in this example
	mockService.On("Payment").Return(mockPaymentService)   // not used in this example
	mockService.On("Lessor").Return(mockLessorService)     // not used in this example
	mockService.On("User").Return(mockUserService)         // not used in this example

	// Mock successful login
	mockAuthService.On("Login", &dtos.LoginDTO{
		Email:    "john@example.com",
		Password: "password123",
	}).Return("mocked-token", nil)

	// Mock failed login with wrong password
	mockAuthService.On("Login", &dtos.LoginDTO{
		Email:    "john@example.com",
		Password: "wrongpassword",
	}).Return("", fmt.Errorf("Invalid credentials"))

	// Mock failed login with non-existent user
	mockAuthService.On("Login", &dtos.LoginDTO{
		Email:    "nouser@example.com",
		Password: "password123",
	}).Return("", fmt.Errorf("User not found"))

	// Mock OTP request
	mockAuthService.On("RequestOTP", &dtos.RequestOTPDTO{
		Email: "john@example.com",
	}).Return(nil)

	// Mock Reset Password Request
	mockAuthService.On("RequestPasswordReset", &dtos.ResetPassRequestDTO{
		Email: "john@example.com",
	}).Return("reset-link", nil)

	// Mock Reset Password
	mockAuthService.On("ResetPassword", &dtos.ResetPassDTO{
		Token:    "mock-token",
		Password: "newpassword123",
	}).Return(nil)

	handler = handlers.NewHandler(mockService)
	recorder = httptest.NewRecorder()
}

// -------------------------------
// Step 2: Define Steps
// -------------------------------

func iHaveAValidRegisterPayload(username, password, email string) error {
	req := dtos.RegisterDTO{
		Name:     username,
		Password: password,
		Email:    email,
	}

	var err error
	requestBody, err = json.Marshal(req)
	if err != nil {
		return err
	}

	// Ensure that the mock expects a valid call
	mockAuthService.On("Register", &req).Return(nil)
	return nil
}

func iSendAPostRequestTo(endpoint string) error {
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	recorder = httptest.NewRecorder()

	handlerFunc := setHandler(endpoint)
	if handlerFunc == nil {
		return fmt.Errorf("no handler found for endpoint: %s", endpoint)
	}

	handlerFunc.ServeHTTP(recorder, req)
	return nil
}

func theResponseCodeShouldBe(code int) error {
	if recorder == nil {
		return fmt.Errorf("recorder is nil")
	}

	// Pass t to assert for better output
	assert.Equal(&testing.T{}, code, recorder.Code, "Response code did not match")
	return nil
}

func theResponseShouldContainAnAccessToken() error {
	var res map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &res)
	if err != nil {
		return err
	}

	// Ensure the token is present
	assert.NotEmpty(&testing.T{}, res["data"], "Access token not found in response")
	return nil
}

func theResponseShouldContainAnErrorMessage(message string) error {
	var res map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &res)
	if err != nil {
		return err
	}

	assert.Equal(&testing.T{}, message, res["message"], "Error message did not match")
	return nil
}

func iHaveANonexistentUsernameAndPassword(username, password string) error {
	req := dtos.LoginDTO{
		Email:    username,
		Password: password,
	}

	var err error
	requestBody, err = json.Marshal(req)
	if err != nil {
		return err
	}

	mockAuthService.On("Login", &req).Return("", fmt.Errorf("User not found"))

	return nil
}

func iHaveAValidOTPRequestPayloadWithEmail(email string) error {
	req := dtos.RequestOTPDTO{
		Email: email,
	}

	var err error
	requestBody, err = json.Marshal(req)
	if err != nil {
		return err
	}

	mockAuthService.On("RequestOTP", &req).Return(nil)

	return nil
}

func iHaveAValidResetPasswordPayloadWithTokenAndNewPassword(token, newPassword string) error {
	req := dtos.ResetPassDTO{
		Token:    token,
		Password: newPassword,
	}

	var err error
	requestBody, err = json.Marshal(req)
	if err != nil {
		return err
	}

	mockAuthService.On("ResetPassword", &req).Return(nil)

	return nil
}

func iHaveAValidResetPasswordRequestWithEmail(email string) error {
	req := dtos.ResetPassRequestDTO{
		Email: email,
	}

	var err error
	requestBody, err = json.Marshal(req)
	if err != nil {
		return err
	}

	mockAuthService.On("RequestPasswordReset", &req).Return("reset-link", nil)

	return nil
}

func iHaveAValidUsernameAndPassword(username, password string) error {
	req := dtos.LoginDTO{
		Email:    username,
		Password: password,
	}

	var err error
	requestBody, err = json.Marshal(req)
	if err != nil {
		return err
	}

	mockAuthService.On("Login", &req).Return("mocked-token", nil)

	return nil
}

func iHaveAValidUsernameAndAnInvalidPassword(username, password string) error {
	req := dtos.LoginDTO{
		Email:    username,
		Password: password,
	}

	var err error
	requestBody, err = json.Marshal(req)
	if err != nil {
		return err
	}

	mockAuthService.On("Login", &req).Return("", fmt.Errorf("Invalid credentials"))

	return nil
}

// -------------------------------
// Step 3: Dynamically Select Handler
// -------------------------------

func setHandler(endpoint string) http.HandlerFunc {
	if handler == nil {
		return nil
	}

	switch endpoint {
	case "/auth/register":
		return fiberToHttpHandler(handler.Auth().Register)
	case "/auth/login":
		return fiberToHttpHandler(handler.Auth().Login)
	case "/auth/logout":
		return fiberToHttpHandler(handler.Auth().Logout)
	case "/auth/request-otp":
		return fiberToHttpHandler(handler.Auth().RequestOTP)
	case "/auth/verify-otp":
		return fiberToHttpHandler(handler.Auth().VerifyOTP)
	case "/auth/forgot-password":
		return fiberToHttpHandler(handler.Auth().ResetPasswordRequest)
	case "/auth/reset-password":
		return fiberToHttpHandler(handler.Auth().ResetPassword)
	default:
		return nil
	}
}

func fiberToHttpHandler(handler fiber.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app := fiber.New()

		// Create fasthttp.RequestCtx
		reqCtx := &fasthttp.RequestCtx{}
		c := app.AcquireCtx(reqCtx)
		defer app.ReleaseCtx(c)

		// Set method and URI
		c.Request().Header.SetMethod(r.Method)
		c.Request().SetRequestURI(r.RequestURI)

		// Read body correctly from request
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Request().SetBody(body)

		// Pass headers correctly
		for k, v := range r.Header {
			for _, vv := range v {
				c.Request().Header.Set(k, vv)
			}
		}

		// Execute the handler
		if err := handler(c); err != nil {
			w.WriteHeader(c.Response().StatusCode())
			w.Write(c.Response().Body())
			return
		}

		// Write back response to HTTP
		w.WriteHeader(c.Response().StatusCode())
		w.Write(c.Response().Body())
	}
}

// -------------------------------
// Step 4: Register All Steps in Godog
// -------------------------------

func InitializeHandler(ctx *godog.ScenarioContext) {
	setupMocks()

	ctx.Step(`^I have a valid register payload with username "([^"]*)", password "([^"]*)", and email "([^"]*)"$`, iHaveAValidRegisterPayload)
	ctx.Step(`^I have a valid username "([^"]*)" and password "([^"]*)"$`, iHaveAValidUsernameAndPassword)
	ctx.Step(`^I have a valid username "([^"]*)" and an invalid password "([^"]*)"$`, iHaveAValidUsernameAndAnInvalidPassword)
	ctx.Step(`^I have a non-existent username "([^"]*)" and password "([^"]*)"$`, iHaveANonexistentUsernameAndPassword)
	ctx.Step(`^I have a valid OTP request payload with email "([^"]*)"$`, iHaveAValidOTPRequestPayloadWithEmail)
	ctx.Step(`^I have a valid reset password request with email "([^"]*)"$`, iHaveAValidResetPasswordRequestWithEmail)
	ctx.Step(`^I have a valid reset password payload with token "([^"]*)" and new password "([^"]*)"$`, iHaveAValidResetPasswordPayloadWithTokenAndNewPassword)
	ctx.Step(`^I send a POST request to "([^"]*)"$`, iSendAPostRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response should contain an access token$`, theResponseShouldContainAnAccessToken)
	ctx.Step(`^the response should contain an error message "([^"]*)"$`, theResponseShouldContainAnErrorMessage)
}
