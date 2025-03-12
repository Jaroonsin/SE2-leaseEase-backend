Feature: User Login

  Scenario: Successful login with valid credentials
    Given I have a valid username "testuser" and password "password123"
    When I send a POST request to "/auth/login"
    Then the response code should be 200
    And the response should contain an access token

  Scenario: Failed login with incorrect password
    Given I have a valid username "testuser" and an invalid password "wrongpassword"
    When I send a POST request to "/auth/login"
    Then the response code should be 401
    And the response should contain an error message "Invalid credentials"

  Scenario: Failed login with non-existent user
    Given I have a non-existent username "nouser" and password "password123"
    When I send a POST request to "/auth/login"
    Then the response code should be 404
    And the response should contain an error message "User not found"
