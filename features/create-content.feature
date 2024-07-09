Feature: Content Creation
    As a user
    I want to create content
    So that I can manage content in the repository

    Scenario: Successful Content Creation
        Given a new content with valid attributes
        When I create the content
        Then the content should be stored in the repository
        And the content should be retrievable by its ID
        And the content should match the attributes provided

    Scenario: Content Creation with Missing or Incorrect Fields
        Given a new content with missing or incorrect attributes
        When I attempt to create the content
        Then an error should be returned indicating the issue

    Scenario Outline: Content Creation with Specific Missing or Incorrect Fields
        Given a new content with a missing or incorrect "<field>"
        When I attempt to create the content
        Then an error should be returned indicating the issue with the "<field>"

        Examples:
            | field     |
            | Id        |
            | Class     |
            | CreatorId |
            | UpdatedAt |
            | CreatedAt |

    Scenario: Content Creation with Duplicate IDs
        Given a new content with a unique ID
        And the content is successfully created
        When I attempt to create another content with the same ID
        Then an error should be returned indicating a duplicate ID
        And the original content should remain unchanged

    Scenario: Cleanup After Tests
        Given a collection of test contents
        When the tests are completed
        Then the test collection should be deleted
        And there should be no residual test data