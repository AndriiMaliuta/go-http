goget:
	go get github.com/lib/pq
	go get github.com/jmoiron/sqlx
	go get github.com/go-gorm/gorm
#You can’t even run it in the background using the & or
 #bg command because once you log out, the web service gets killed.
	# nohup ./ws-s &