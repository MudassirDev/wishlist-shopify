FROM debian:stable-slim

# Copy the binary
COPY out /usr/local/bin/out

# Make it executable
RUN chmod +x /usr/local/bin/out

# Run the binary
CMD ["/usr/local/bin/out"]
