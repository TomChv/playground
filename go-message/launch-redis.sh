docker run --name "$REDIS_NAME" -e REDIS_PASSWORD="$REDIS_PASS" -p "$REDIS_PORT":"$REDIS_PORT" -d bitnami/redis:latest