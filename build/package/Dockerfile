# =============================================================================
#
#
# Project:
#     https://github.com/jotka/ac_rest_go.git
#     Home Asisstant IP climate middleware for Samsung AC devices.
#
#
# =============================================================================
#
# -----------------------------------------------------------------------------
# Base image
# -----------------------------------------------------------------------------
FROM alpine:3.10

# -----------------------------------------------------------------------------
# Install system packages
# -----------------------------------------------------------------------------
RUN apk update
RUN apk add --no-cache openssl bash ca-certificates tzdata
RUN update-ca-certificates 2>/dev/null || true

# -----------------------------------------------------------------------------
# Set image metadata
# -----------------------------------------------------------------------------
LABEL name="jotka/ac_rest_go" \
      maintainer="Jarek Krochmalski <jkrochmalski@gmail.com>" \
      summary="Samsung AC middleware for Home Assistant." \
      description="Home Asisstant IP climate middleware for Samsung AC devices."

# -----------------------------------------------------------------------------
# Stage application
# -----------------------------------------------------------------------------
# Agent binaries directory should be set as build context
COPY ac-rest-go /usr/bin/

# -----------------------------------------------------------------------------
# Provide a non-root user to run the process.
# -----------------------------------------------------------------------------
RUN addgroup -g 1001 -S acrestgo && \
    adduser -u 1001 -S acrestgo -G acrestgo && \
    chown -R 1001:1001 /usr/bin/ac-rest-go && \
    chmod -R 755 /usr/bin/ac-rest-go && \
    chmod -R 644 /etc/ac-rest-go

# -----------------------------------------------------------------------------
# Define execution user
# -----------------------------------------------------------------------------
USER acrestgo

# -----------------------------------------------------------------------------
# Expose service ports
# -----------------------------------------------------------------------------
EXPOSE 8080

# -----------------------------------------------------------------------------
# Run container
# -----------------------------------------------------------------------------
ENTRYPOINT ["/usr/bin/ac-rest-go"]
