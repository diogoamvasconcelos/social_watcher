# DynamoDb Table

- single table

## Parameters:

### Main entry
- date: when it happened
- source: twitter, instagram, reddit, hackernews, facebook, etc
- id: uuid
- hash: for duplication checks: could be externalId or hashOfMessage
- data: OBJECT?
- link: url to navigate to original content
  
### Entry aggregates
  Could be comments, etc.. should be a list
  
  - entryId
  - date
  - data
  - type: comment, reply, etc...
    - type=0 -> main entry
    - type>0 -> entry aggregate

## Data access - queries

- query by "date"
- query by "source"
- query by "id"

# Table Design:

pk: id
sk: index

## Indices:

- all:
  - pk: "ALL"
  - sk: date
- source:
  - pk: source
  - sk: date
- day:
  - pk: day
  - sk: source