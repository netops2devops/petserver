Rudimentary example of creating http handlers and middlewares in Go using standard `net/http` library.

### Usage

create

```
❯ curl -X POST \
-H "Content-type: application/json" \
-d '{"name":"Bella","kind":"Cat","age":5,"color":"Gray"}' \
localhost:8080/pets/

403 Forbidden

❯ curl -X POST \
-H "Content-type: application/json" \
-H "X-AUTH-TOKEN: secretlifeofpets" \
-d '{"name":"Bella","kind":"Cat","age":5,"color":"Gray"}' \
localhost:8080/pets/

[
  {
    "name": "Juniper",
    "kind": "Cat",
    "age": 5,
    "color": "Orange"
  },
  {
    "name": "Ashby",
    "kind": "Cat",
    "age": 5,
    "color": "Gray"
  },
  {
    "name": "Bruce",
    "kind": "Dog",
    "age": 8,
    "color": "Golden"
  },
  {
    "name": "Bella",
    "kind": "Cat",
    "age": 5,
    "color": "Gray"
  }
]

```

Delete

```
❯ curl -X DELETE localhost:8080/pets/Bella
403 Forbidden


❯ curl -X DELETE -H "X-AUTH-TOKEN: secretlifeofpets" localhost:8080/pets/Bella
{
  "Pets": [
    {
      "name": "Juniper",
      "kind": "Cat",
      "age": 5,
      "color": "Orange"
    },
    {
      "name": "Ashby",
      "kind": "Cat",
      "age": 5,
      "color": "Gray"
    },
    {
      "name": "Bruce",
      "kind": "Dog",
      "age": 8,
      "color": "Golden"
    }
  ]
}
```
