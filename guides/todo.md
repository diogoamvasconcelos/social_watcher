- cloudwatch event trigger for lambda
- twitterLambda(keyword) -> searches twitter (last hour/day) and adds to ddb
  - v1: searches last hour
    - writes to ddb (batches in parallel?)
    - make `keyword` part of the stored item data (and queriable, and post it on slack)
  - v2: also searches for replies on main entries that are active (having a reply in the last x days)
- ddbstream consumer lambda -> consumes and posts discord

  - v3: translates if not english?

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
