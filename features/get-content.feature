Feature: Content Retrieval
    As a user
    I want to retrieve content
    So that I can view the stored content in the repository

    Background:
        Given a connection to the database
        And a ContentRepository instance
        And a unique testsCollection is created
        And a valid initial content is stored in the repository

    Scenario: Successful Content Retrieval
        When I retrieve the content by its ID
        Then the content should be returned
        And the content attributes should match the stored content

    Scenario: Non-Existent Content
        Given the repository does not contain a content with a specific ID
        When I attempt to retrieve the content by this non-existent ID
        Then an error should be returned indicating "content not found"
        And no content should be returned

    Scenario: Invalid Content ID
        Given I provide an invalid ID format
        When I attempt to retrieve the content by this invalid ID
        Then an error should be returned indicating the invalid ID
        And no content should be returned

    Scenario: Database Connection Nil
        Given the database connection is nil
        When I attempt to retrieve a content
        Then an error should be returned indicating the database connection is nil
        And no content should be returned

    Scenario: Successful Collection Retrieval
        When I retrieve all contents in the collection
        Then the contents should be returned
        And the contents should match the stored contents

    Scenario: Empty Collection
        Given the collection is empty
        When I retrieve all contents in the collection
        Then an empty list should be returned

    Scenario: Invalid Collection
        Given I provide an invalid collection name
        When I attempt to retrieve all contents in the collection
        Then an error should be returned indicating the invalid collection
        And no contents should be returned

    Scenario: Successful Class Retrieval
        When I retrieve all contents in the class
        Then the contents should be returned
        And the contents should match the stored contents

    Scenario: Empty Class
        Given the class is empty
        When I retrieve all contents in the class
        Then an empty list should be returned

    Scenario: Invalid Collection for Class Retrieval
        Given I provide an invalid collection name
        When I attempt to retrieve all contents in the class
        Then an error should be returned indicating the invalid collection
        And no contents should be returned

        Cleanup:
        Given the testsCollection is deleted