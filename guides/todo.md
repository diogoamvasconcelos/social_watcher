- twitterLambda(keyword) -> searches twitter (last hour/day) and adds to ddb

  - searches last hour
    - writes to ddb (batches in parallel?)

- Go improvements:

  - Use the AWS SDK v2: https://aws.github.io/aws-sdk-go-v2/docs/

- CI/Dev

  - Run lambda locally
  - Unit Tests
  - Integration Tests
  - Environment Tests
  - Test coverage
  - Vulnerability scans
  - Behaviour test (cucumber?)

- Reddit search

  - run once a day (very few posts)
  - or native https://www.reddit.com/dev/api/#GET_search (not sure if searhes for comments)
  - use https://redditsearch.io/ (not sure if legal)
