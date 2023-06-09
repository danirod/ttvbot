README.txt: docs/ttvbot.8
	groff -man docs/ttvbot.8 -Tutf8 -P-c -P-b -P-u > README.txt
