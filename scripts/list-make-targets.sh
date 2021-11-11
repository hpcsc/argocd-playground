#!/bin/bash

MAKEFILES=$@
awk 'BEGIN {
    FS = ":.*##";
    printf "\nUsage:\n  make \033[36m<target>\033[0m\n"
}
/^[a-zA-Z0-9_-]+:.*?##/ {
    # show task name and description
    printf "  \033[36m%-25s\033[0m %s\n", $1, $2
}
/^##@/ {
    # show section header
    printf "\n\033[1m%s\033[0m\n", substr($0, 5)
} ' ${MAKEFILES}
