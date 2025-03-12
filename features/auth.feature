Feature: Auth Service

  Scenario: Successful user registration
    Given I have a valid register payload with username "john_doe", password "password123", and email "john@example.com"
    When I send a POST request to "/auth/register"
    Then the response code should be 201

  Scenario: Failed registration with missing data
    Given I have a valid register payload with username "john_doe", password "", and email "john@example.com"
    When I send a POST request to "/auth/register"
    Then the response code should be 400
    And the response should contain an error message "Invalid request payload"

  Scenario: Successful login
    Given I have a valid register payload with username "john_doe", password "password123", and email "john@example.com"
    When I send a POST request to "/auth/login"
    Then the response code should be 201
    And the response should contain an access token

  Scenario: Failed login with incorrect password
    Given I have a valid register payload with username "john_doe", password "wrongpassword", and email "john@example.com"
    When I send a POST request to "/auth/login"
    Then the response code should be 401
    And the response should contain an error message "Invalid credentials"

  Scenario: Request OTP
    Given I have a valid OTP request payload with email "john@example.com"
    When I send a POST request to "/auth/request-otp"
    Then the response code should be 201

  Scenario: Reset Password Request
    Given I have a valid reset password request with email "john@example.com"
    When I send a POST request to "/auth/forgot-password"
    Then the response code should be 200

  Scenario: Reset Password
    Given I have a valid reset password payload with token "mock-token" and new password "newpassword123"
    When I send a POST request to "/auth/reset-password"
    Then the response code should be 200
