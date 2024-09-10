#!/bin/false

# simply source this file as shown below, and all invocations of docker and docker-compose be via sudo
# please note that this will affect your entire shell session, and not just sandbox commands
# also, the sudo invocation is with -E, which passes through all environment variables
#
# . docker-is-sudo-docker.sh

if mkdir /tmp/.docker-is-sudo-docker 2>/dev/null
then
  cat >/tmp/.docker-is-sudo-docker/docker-is-sudo-docker <<'EOF'
#!/usr/bin/env bash
true
EOF
  cat >/tmp/.docker-is-sudo-docker/docker <<'EOF'
#!/usr/bin/env bash
sudo -E docker "$@"
EOF
  cat >/tmp/.docker-is-sudo-docker/docker-compose <<'EOF'
#!/usr/bin/env bash
sudo -E docker-compose "$@"
EOF
  chmod a+x /tmp/.docker-is-sudo-docker/docker{-is-sudo-docker,,-compose}
fi

if ! which docker-is-sudo-docker &>/dev/null
then
  export PATH=/tmp/.docker-is-sudo-docker:"$PATH"
fi

if ! which docker-is-sudo-docker &>/dev/null
then
  echo oops
else
  echo done
fi
