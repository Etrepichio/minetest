# Base Alpine Linux image
FROM alpine:3.6 

# Add binary to bin directory
ADD minesweeper bin/

# Commad to execute at start up
ENTRYPOINT ["bin/minesweeper"]
CMD ["-addr", ":8080"]

# Default port
EXPOSE 8080
