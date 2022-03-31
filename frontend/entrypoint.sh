#!/bin/sh

echo "Check that we have NEXT_PUBLIC_API_URL vars"
test -n "$NEXT_PUBLIC_API_URL"
test -n "$NEXTGATEWAY_PORT"

find /app/.next \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i "s#APP_NEXT_PUBLIC_API_URL#$NEXT_PUBLIC_API_URL#g"
find /app/.next \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i "s#APP_NEXTGATEWAY_PORT#$NEXTGATEWAY_PORT#g"

echo "Starting Nextjs"
exec "$@"