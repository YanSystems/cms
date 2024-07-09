Feature: Content Deletion
    As a user
    I want to delete content
    So that I can manage content in the repository

    Background:
        Given a connection to the database
        And a ContentRepository instance
        And a unique testsCollection is created
        And a valid initial content is stored in the repository

    Scenario: Successful Content Deletion
        Given a content with valid attributes is stored in the repository
        When I delete the content by its ID
        Then the content should be removed from the repository
        And the content should no longer be retrievable by its ID

    Scenario: Non-Existent Content Deletion
        Given the repository does not contain a content with a specific ID
        When I delete the content by this non-existent ID
        Then no error should be returned
        And the ID of the non-existent content should be returned

    Scenario: Invalid Content ID Deletion
        Given I provide an invalid ID format
        When I delete the content by this invalid ID
        Then no error should be returned
        And the ID of the invalid content should be returned

    Scenario: Successful Class Deletion
        Given multiple contents with the same class are stored in the repository
        When I delete all contents by their class
        Then the contents should be removed from the repository
        And the contents should no longer be retrievable by their IDs

    Scenario: Non-Existent Class Deletion
        Given the repository does not contain contents with a specific class
        When I delete all contents by this non-existent class
        Then no error should be returned
        And an empty list of IDs should be returned

    Scenario: Invalid Collection for Class Deletion
        Given I provide an invalid collection name
        When I delete all contents by their class in this invalid collection
        Then an error should be returned
        And no IDs should be returned

    Scenario: Successful Collection Deletion
        Given multiple contents are stored in a collection
        When I delete the collection
        Then the collection should be removed from the database
        And the contents should no longer be retrievable by their IDs
        And the collection should no longer exist in the database

    Scenario: Empty Collection Deletion
        Given a collection is empty
        When I delete the collection
        Then no error should be returned
        And an empty list of IDs should be returned
        And the collection should no longer exist in the database

    Scenario: Invalid Collection Deletion
        Given I provide an invalid collection name
        When I delete the collection
        Then an error should be returned
        And no IDs should be returned

        Cleanup:
        Given the testsCollection is deleted