if test git && test go; then
  cd /tmp
  git clone https://github.com/maateen/battery-monitor.git
  cd battery-monitor
  go mod tidy
  GOOS=linux GOARCH=amd64 go build cmd/battery-monitor/battery-monitor.go
  mv battery-monitor /usr/local/bin/
  mkdir -p /etc/battery-monitor
  mv assets /etc/battery-monitor/
  mv configs/config.toml /etc/battery-monitor/
  sed -i "s/userNameHere/$(whoami)/g" scripts/battery-monitor
  mv scripts/battery-monitor /etc/init.d/
else
  echo "Dependencies not met."
fi