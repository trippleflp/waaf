################
##### Builder
FROM rust:1.61.0-slim as builder

WORKDIR /usr/src/hello-world

# We want dependencies cached, so copy those first.
COPY . .

## Install target platform (Cross-Compilation) --> Needed for Alpine
RUN rustup target add x86_64-unknown-linux-musl

# This is a dummy build to get the dependencies cached.
RUN cargo build --target x86_64-unknown-linux-musl --release

# Now copy in the rest of the sources
COPY src /usr/src/hello-world/src/

## Touch main.rs to prevent cached release build
RUN touch /usr/src/hello-world/src/main.rs

# This is the actual application build.
RUN cargo build --target x86_64-unknown-linux-musl --release

################
##### Runtime
FROM alpine:3.16.0 AS runtime 

# Copy application binary from builder image
COPY --from=builder /usr/src/hello-world/target/x86_64-unknown-linux-musl/release/hello-world /usr/local/bin


# Run the application
CMD ["/usr/local/bin/hello-world"]