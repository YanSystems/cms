Feature: Content Update
    As a user
    I want to update content
    So that I can manage content in the repository

    Background:
        Given a connection to the database
        And a ContentRepository instance
        And a unique testsCollection is created

    Scenario: Successful Update
        Given a content with valid attributes is stored in the repository
        When I update the content with new attributes
        Then the content should be updated in the repository
        And the updated content should have the new attributes
        And the updated content should have a new UpdatedAt timestamp
        And the CreatedAt timestamp should remain unchanged

    Scenario: Partial Update
        Given a content with valid attributes is stored in the repository
        When I partially update the content with new attributes
        Then the content should be partially updated in the repository
        And the unchanged attributes should remain the same
        And the updated content should have a new UpdatedAt timestamp
        And the CreatedAt timestamp should remain unchanged

    Scenario: Non-Existent Content Update
        Given the repository does not contain a content with a specific ID
        When I attempt to update the content by this non-existent ID
        Then an error should be returned indicating the content does not exist

    Scenario: Invalid Content ID Update
        Given I provide an invalid ID format
        When I attempt to update the content by this invalid ID
        Then an error should be returned indicating the invalid ID

        Cleanup:
        Given the testsCollection is deleted
