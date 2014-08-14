#!/bin/sh

set -e

if [ ! -d /etc/inspeqtor ]; then
  mkdir /etc/inspeqtor
fi

if [ ! -d /etc/inspeqtor/conf.d ]; then
  cp -r /usr/share/inspeqtor/conf.d /etc/inspeqtor

  # on Ubuntu, allow any adms to edit inqs
  if ! chown -R root:adm /etc/inspeqtor/conf.d ; then
    true # ignore if adm doesn't exist
  fi
fi

if [ ! -f /etc/inspeqtor/inspeqtor.conf ]; then
  cp /usr/share/inspeqtor/inspeqtor.conf /etc/inspeqtor
  # passwords in here, root readable only.
  chmod 600 /etc/inspeqtor/inspeqtor.conf
fi


if [ ! -e /etc/service/inspeqtor ]; then
  ln -s /etc/sv/inspeqtor /etc/service/inspeqtor
else
  sv restart inspeqtor
fi


cat <<"TXT"
 _                            _
(_)_ __  ___ _ __   ___  __ _| |_ ___  _ __
| | '_ \/ __| '_ \ / _ \/ _  | __/ _ \| '__|
| | | | \__ \ |_) |  __/ (_| | || (_) | |
|_|_| |_|___/ .__/ \___|\__, |\__\___/|_|
            |_|            |_|


Thank you for installing Inspeqtor!

Please configure your notification settings in /etc/inspeqtor/inspeqtor.conf and
then restart Inspeqtor with 'sudo sv restart inspeqtor'.
TXT
