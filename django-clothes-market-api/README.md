# Django Clothes market API

## Quick start

- Install [Poetry](https://python-poetry.org/docs/)
- Install dependencies `poetry install`
- Run migrations `poetry run python manage.py migrate`
- Create superuser `poetry run python manage.py createsuperuser`
- Run server `poetry run python manage.py runserver`

## Endpoints

> Test the API using [Postman](https://www.postman.com) and the [Example.postman_collection.json](./Example.postman_collection.json)

```
  GET /clothes/article
  Description: Retrieve all articles stored in a database. A query param can be set to filter selection.
  Query Params:
    - Sold: boolean (set to only retrieve sold or in stock articles)
  Return: { number_of_articles: number, articles: Article[] }

  GET /clothes/article/:id/
  Description: Retrieve an article by its identifier
  Return: { Article }

  POST /clothes/article
  Description: Create a new article, take as body required information to create an article and returns the created row from the db.
  Body: { Article (except id, sold_boolean & sold_price) }
  Return: { Article }

  PUT /clothes/article/:id/
  Description: Update article information (can be one or multiple field or the selected article) and returns it with updated information.
  Body: { Article (except id, every props are optional and only these set will be updated) }
  Return: { Article }

  DELETE /clothes/article/:id/
  Description: Delete an article by its identifier.
  Return: { message: "Article {name} successfully deleted." }
```