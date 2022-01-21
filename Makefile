goget:
	go get github.com/jmoiron/sqlx
#You can’t even run it in the background using the & or
 #bg command because once you log out, the web service gets killed.
	# nohup ./ws-s &