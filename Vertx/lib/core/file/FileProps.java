---------------------------
FileProps
---------------------------
	# �ļ����Խӿڣ� interface FileProps

	long creationTime();
	long lastAccessTime();
	long lastModifiedTime();
	boolean isDirectory();
	boolean isOther();
	boolean isRegularFile();
	boolean isSymbolicLink();
	long size();