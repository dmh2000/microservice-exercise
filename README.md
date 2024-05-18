## Pro:

### development

- small codebase
- multiple language friendly
- simpler testing
- simpler API
- distributed development

### devops

- horizontal scaling
- faster build and deployment
- independent updates
- fault isolation
- cost optimization

### Cons

- resource overhead
- complex debugging
- transactions?
- observability
- complex integration testing

## api

- rating service

  - generic records with rating
  - internal api
  - actions
    - store rating
    - get aggregated data by id and type
  - data
    - record id
    - user id
    - type
    - rating value

- movie metadata service

  - internal api
  - actions
    - get metadata by id
    - store metadata
  - data (static)
    - movie id
    - title
    - year
    - description
    - directory
    - actors[]

- movie service

  - client facing API
  - actions
    - get complete information including metadata and rating
      - for one movie
      - for multiple movies
  - data
    - movie metadata
    - movie rating

### structure

<pre>
                    +-> rating srevice
    client +-> movie service +
                    +-> metadata service
    - metadata service
    - rating service
</pre>

<pre>
    API handler -> business logic (controller) -> repository -> DB
<pre>
