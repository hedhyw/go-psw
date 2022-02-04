Feature: Generator of passwords
    Scenario: User wants to generate a password
        When app is called without arguments
        Then a random password is printed
    
    Scenario: User wants to generate a password of given length
        When app is called with a single argument <password_length>
        Then a random password of length <password_length> is printed
        Examples:
        | <password_length> |
        | 10                |
        | 100               |
    
    Scenario: User gives an invalid password length argument
        When app is called with a single argument <password_length>
        Then app returns an error
        Examples:
        | <password_length> |
        | 0                 |
        | -1                |
        | invalid           |

    Scenario: User wants to use only specific groups of chars
        When app is called with <password_length> = 1 and <chars>
        Then a random password of length <password_length> is printed
        And password contains only chars of a <group>.
        Examples:
        | <chars> | group        |
        | a       | LowerLetters |
        | A       | UpperLetters |
        | 1       | Digits       |
        | .       | Symbols      |
        | ё       | Single       |
