# Correction step 1

## Characters and count

```graphql
query {
  characters {
    nodes {
      name
      appearsIn
      height
    }
    totalCount
  }
}
```

## Reviews and count

```graphql
query {
  reviews(episode: EMPIRE) {
    nodes {
      commentary
      stars
    }
    totalCount
  }
}
```

## Jedi's heroes' and friends

```graphql
query {
  hero(episode: JEDI) {
    id
    name
    friends {
      nodes {
        id
        name
      }
    }
  }
}
```

## Discover mutation

```graphql
mutation {
  createReview(input: { episode: EMPIRE, stars: 10, commentary: "COMMENT" }) {
    episode
    review {
      id
      commentary
      stars
    }
  }
}
```

## Discover sub query parameters

```graphql
query {
  hero(episode: EMPIRE) {
    id
    name
    height(unit: FOOT)
    friends {
      nodes {
        id
        name
        height(unit: FOOT)
      }
    }
  }
}
```

```graphql
query OrderCharactersAndHero {
  characters(order_by: { height: DESC }) {
    nodes {
      id
      name
      height
    }
  }

  hero(episode: EMPIRE) {
    name
    friends {
      nodes {
        id
        name
      }
    }
  }
}
```