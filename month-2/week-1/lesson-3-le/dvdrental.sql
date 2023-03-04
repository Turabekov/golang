
-- 1. TASK
SELECT  
    a.first_name as name,
    ARRAY_AGG(f.title) as films
FROM film_actor as fa
JOIN actor as a ON a.actor_id = fa.actor_id
JOIN film as f ON f.film_id = fa.film_id
GROUP BY a.first_name
ORDER BY a.first_name;

-- 2. TASK
SELECT  
    a.first_name as name,
    ARRAY_AGG(c.name) as categories
FROM film_actor as fa
JOIN actor as a ON a.actor_id = fa.actor_id
JOIN film as f ON f.film_id = fa.film_id
JOIN film_category as fc ON f.film_id = fc.film_id
JOIN category as c ON c.category_id = fc.category_id
GROUP BY a.first_name
ORDER BY a.first_name;
