#! /usr/bin/env bash

if [ "$EUID" -ne 0 ]
  then echo "ERROR: Please run as root" && exit 1
fi

options=$(getopt -o d:gl -l dns:,google,local -- "$@")

[ $? -eq 0 ] || {
    echo "Incorrect options provided"
    exit 1
}

eval set -- "$options"
while true; do
  case "$1" in
    -d|--dns) shift; dns=${1} ;;
    -g|--google) dns="google" ;;
    -l|--local) dns="local" ;;
    --) shift; break ;;
  esac
  shift
done

if [ "${dns}" != "google" ] && [ "${dns}" != "local" ]
  then echo 'ERROR: Should be "google" or "local"' && exit 2
fi

chattr -i /etc/resolv.conf

if [ "${dns}" = "google" ]
  then sed -i "s/nameserver.*/nameserver 8.8.8.8/" /etc/resolv.conf
  echo 'SUCCESS: Your DNS is now configured as 8.8.8.8'
elif [ "${dns}" = "local" ]
  then sed -i "s/nameserver.*/nameserver 127.0.0.1/" /etc/resolv.conf
  echo 'SUCCESS: Your DNS is now configured as 127.0.0.1'
fi

chattr +i /etc/resolv.conf
exit 0
