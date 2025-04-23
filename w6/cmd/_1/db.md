Анатомия веб-сервиса
* Работа с БД и хранилищами
```sql
    SELECT COUNT(*), category_id
    FROM products
    WHERE status = 1
    GROUP BY category_id
```