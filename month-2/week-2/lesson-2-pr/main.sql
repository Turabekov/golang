-- 1. Create a function that takes two strings as parameters and 
-- returns a new string that is the concatenation of the two strings, 
-- with a space between them.

CREATE OR REPLACE FUNCTION concat_str(str1 VARCHAR, str2 VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$ 
    BEGIN
        return CONCAT(str1, ' ', str2);
    END;
$$;

-- 2. Create a function that takes a string as a parameter and returns a new string that is the reverse of the original string.

CREATE OR REPLACE FUNCTION reverse_str(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS 
$$ 
    DECLARE
        reverse VARCHAR = '';
        i integer = LENGTH(str);
    BEGIN


        LOOP
            IF i < 0 THEN 
                EXIT;
            END IF;

            reverse = reverse || substring(str, i, 1);
            i = i - 1;
        END LOOP;

        return reverse;
    END;
$$;

-- 3. Create a function that takes a string as a parameter and returns a new string with all the vowels removed.
CREATE OR REPLACE FUNCTION remove_all_vowels(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$ 
    DECLARE
        result VARCHAR = '';
    BEGIN


        FOR i IN 1..LENGTH(str) LOOP
            IF LOWER(substring(str, i, 1)) = 'a' OR LOWER(substring(str, i, 1)) = 'e' OR LOWER(substring(str, i, 1)) = 'u' OR LOWER(substring(str, i, 1)) = 'i' OR LOWER(substring(str, i, 1)) = 'o' THEN 
                
            ELSE 
                result = result || substring(str, i, 1);
            END IF;

        END LOOP;

        return result;
    END;
$$;

-- 4. Create a function that takes a string as a parameter and returns a new string with all the consonants removed.

CREATE OR REPLACE FUNCTION remove_all_consonants(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$ 
    DECLARE
        result VARCHAR = '';
    BEGIN


        FOR i IN 1..LENGTH(str) LOOP
            IF LOWER(substring(str, i, 1)) = 'a' OR LOWER(substring(str, i, 1)) = 'e' OR LOWER(substring(str, i, 1)) = 'u' OR LOWER(substring(str, i, 1)) = 'i' OR LOWER(substring(str, i, 1)) = 'o' THEN 
                result = result || substring(str, i, 1);
            END IF;

        END LOOP;

        return result;
    END;
$$;

-- 5. Create a function that takes a string as a parameter and returns a new string with the first letter of each word capitalized.
CREATE OR REPLACE FUNCTION first_letter_capitalize(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$ 
    DECLARE
        result VARCHAR = '';
    BEGIN


        FOR i IN 1..LENGTH(str) LOOP
            IF i = 1 THEN 
                result = result || UPPER(substring(str, i, 1));
            END IF;

            IF LOWER(substring(str, i, 1)) = ' ' THEN
                result = result || UPPER(substring(str, i, 1));
                result = result || UPPER(substring(str, i + 1, 1));
            ELSE 
                result = result || substring(str, i, 1);
            END IF;

        END LOOP;

        return result;
    END;
$$;
-- 6. Create a function that takes a string as a parameter and returns a new string with all the characters converted to uppercase.
CREATE OR REPLACE FUNCTION uppercase(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$ 
    BEGIN

        return UPPER(str);
    END;
$$;

-- 7. Create a function that takes a string as a parameter and returns a new string with all the characters converted to lowercase.
CREATE OR REPLACE FUNCTION lowercase(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$ 
    BEGIN

        return LOWER(str);
    END;
$$;

-- 8. Create a function that takes a string as a parameter and returns a new string with all occurrences of a given substring replaced with another substring. The function should take three parameters: the original string, the substring to be replaced, and the replacement substring.

CREATE OR REPLACE FUNCTION replace_Str(str_original VARCHAR, replaced VARCHAR, replacement VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$ 
    BEGIN
        return REPLACE(str_original, replaced, replacement);
    END;
$$;

-- 9 Create a function that takes a string as a parameter and returns a new string with all non-alphabetic characters removed.
CREATE OR REPLACE FUNCTION remove_nonalphabetic(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$  
    BEGIN
        return regexp_replace(str, '\W+', '', 'g');
    END;
$$;

-- 10 Create a function that takes a string as a parameter and returns a new string with all digits removed.
CREATE OR REPLACE FUNCTION remove_digits(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$  
    BEGIN
        return REGEXP_REPLACE(str,'[[:digit:]]','','g');
    END;
$$;
-- 11 Create a function that takes a string as a parameter and returns a new string with all whitespace characters removed.
CREATE OR REPLACE FUNCTION remove_space(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$  
    BEGIN
        return REPLACE(str, ' ','');
    END;
$$;
-- 12 Create a function that takes a string as a parameter and returns a new string with all leading and trailing whitespace characters removed.
CREATE OR REPLACE FUNCTION trim_str(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$  
    BEGIN
        return TRIM(str, ' ');
    END;
$$;

-- 13 Create a function that takes a string as a parameter and returns a new string with all occurrences of a given character removed. The function should take two parameters: the original string and the character to be removed.
CREATE OR REPLACE FUNCTION remove_char(str VARCHAR, c VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$
    DECLARE
        res VARCHAR = '';
    BEGIN
        FOR i IN 1..LENGTH(str) LOOP
            IF substring(str, i, 1) <> c THEN 
                res = res || substring(str, i, 1);
            END IF;
        END LOOP;

        return res;
    END;
$$;

-- 14 Create a function that takes a string as a parameter and returns a new string with all consecutive duplicate characters removed.
CREATE OR REPLACE FUNCTION remove_dupl(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$
    DECLARE
        res VARCHAR = '';
        bool INT = 0;
    BEGIN
        FOR i IN 1..LENGTH(str) LOOP
            IF substring(str, i, 1) = substring(str, i + 1, 1) THEN 
                bool = 1;
            END IF;

            IF substring(str, i, 1) <> substring(str, i + 1, 1) THEN
                IF bool = 1 THEN
                    bool = 0;
                ELSE 
                    res = res || substring(str, i, 1);
                END IF;
            END IF;
        END LOOP;

        return res;
    END;
$$;

-- 15 Create a function that takes a string as a parameter and returns a new string with all consecutive whitespace characters collapsed into a single space.
CREATE OR REPLACE FUNCTION remove_dupl_space(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$
    DECLARE
        res VARCHAR = '';
        bool INT = 0;
    BEGIN
        FOR i IN 1..LENGTH(str) LOOP
            IF substring(str, i, 1) = ' ' AND substring(str, i + 1, 1) = ' ' THEN 
                bool = 1;
                CONTINUE;
            END IF;

            IF substring(str, i, 1) <> ' ' AND substring(str, i + 1, 1) <> ' ' THEN
                IF bool = 1 THEN
                    res = res || ' ';
                    bool = 0;
                ELSE 
                    res = res || substring(str, i, 1);
                END IF;
            END IF;
        END LOOP;

        return res;
    END;
$$;

CREATE OR REPLACE FUNCTION remove_dupl_space(str VARCHAR) RETURNS VARCHAR LANGUAGE PLPGSQL 
    AS
$$
    BEGIN
        return regexp_replace(str, '( ){1,}', ' ', 'g');
    END;
$$;


-- 16 Create a function that takes an integer as a parameter and returns the first N Fibonacci numbers as an array.
CREATE OR REPLACE FUNCTION fibanacci(N INT) RETURNS INT[] LANGUAGE PLPGSQL 
    AS
$$
    DECLARE
        nums INT[] = ARRAY[0, 1];
        a INT = 0;
        b INT = 1;
        c INT = b;
    BEGIN

        LOOP
            c = b;
            b = a + b;
            IF b >= N THEN
                EXIT;
            END IF;
            a = c;
            nums = array_append(nums, b);
        END LOOP;

        return nums;
    END;
$$;

-- func reverseWords(s string) string {  
--     words := strings.Split(s, " ");


--     for i, word := range words {
--         j := len(word) - 1
--         var reverse string
--         for j >= 0 {
--             reverse += string(word[j])
--             j--
--         }
--         words[i] = reverse
--     }

--     return strings.Join(words, " ")
-- }