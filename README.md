# go-mobile-log-elk
Sample Application showing how to we write remote logs using 'go' generated mobile libraries (android and ios)

# Start ELK Stack

docker-compose up -d

# Generate .aar/.framework file
Issue command to generate library file

```
gomobile bind -tags mobile -target=android .
gomobile bind -tags mobile -target=ios .

```

# Use newly generated library in android's MyApplicationTest app
