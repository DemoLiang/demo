#!/usr/bin/env bash
set -e

ulimit -n 102400


# update timezone
#cp /jfg/cfg/timezone /etc/timezone
#dpkg-reconfigure -f noninteractive tzdata

start supervisord with no daemon
/usr/bin/supervisord -c /etc/supervisor/supervisord.conf -n
