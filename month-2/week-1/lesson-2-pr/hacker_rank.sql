
-- 1
SELECT * FROM city WHERE
population > 100000 and CountryCode = 'USA';
-- 2
SELECT name FROM city WHERE
population > 120000 and CountryCode = 'USA';
-- 3
SELECT * FROM city 

-- 4
SELECT * FROM city WHERE
ID = 1661;
-- 5
SELECT name FROM city WHERE
COUNTRYCODE = 'JPN';
-- 6 
SELECT city, state FROM station;

-- 7 
SELECT DISTINCT city FROM station WHERE
ID % 2 = 0;

-- 8
SELECT (COUNT(city) - COUNT(DISTINCT city)) FROM station;

-- 9
SELECT city FROM station WHERE
char_length(city) = (SELECT MIN(char_length(city)) from station) ORDER BY city ASC;

SELECT city FROM station WHERE
char_length(city) = (SELECT MAX(char_length(city)) from station) ORDER BY city ASC;

-- 10
SELECT DISTINCT CITY FROM STATION where SUBSTR(CITY,1,1) NOT IN('A','E','I','O','U') and SUBSTR(CITY,char_length(city),1) IN('A','E','I','O','U');

-- 11
SELECT DISTINCT CITY FROM STATION where SUBSTR(CITY,char_length(city),1) IN('A','E','I','O','U');

SELECT DISTINCT city FROM station WHERE SUBSTR(city, length(city),1) NOT IN('A','E','I','O','U');


SELECT DISTINCT CITY FROM STATION where SUBSTR(CITY,1,1) NOT IN('A','E','I','O','U') OR UPPER(SUBSTR(CITY, length(city),1)) NOT IN('A','E','I','O','U');

-- Query the list of CITY names from STATION that do not start with vowels and do not end with vowels. Your result cannot contain duplicates.

SELECT name FROM STUDENTS WHERE 
marks > 75 ORDER BY SUBSTR(name, -3,3,), ID ASC;

select months*salary, count(*) from employee
group by months*salary
order by months*salary desc
limit 1;