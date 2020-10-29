gnome-terminal --window-with-profile=dgraph-start -e "/usr/local/bin/dgraph alpha --lru_mb 1024"
gnome-terminal --window-with-profile=dgraph-start -e "/usr/local/bin/dgraph zero"
gnome-terminal --window-with-profile=dgraph-start -e "/usr/local/bin/dgraph-ratel"
