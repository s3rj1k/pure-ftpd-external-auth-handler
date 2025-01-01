CREATE TABLE `users` (
  `UserName` varchar(255) NOT NULL COMMENT 'Account user name',
  `Password` varchar(255) NOT NULL COMMENT 'Account password in plain text',
  `UID` int(11) NOT NULL DEFAULT '1034' COMMENT 'uid:xxx - The system uid to be assigned to that user. Must be > 0.',
  `GID` int(11) NOT NULL DEFAULT '1034' COMMENT 'gid:xxx - The primary system gid. Must be > 0.',
  `HomeDirectory` varchar(255) NOT NULL COMMENT 'dir:xxx - The absolute path to the home directory. Can contain /./ for a chroot jail.',
  `UploadBandwidth` bigint(20) NOT NULL DEFAULT '1000000000' COMMENT 'throttling_bandwidth_ul:xxx - The allocated bandwidth for uploads, in bytes per second.',
  `DownloadBandwidth` bigint(20) NOT NULL DEFAULT '1000000000' COMMENT 'throttling_bandwidth_dl:xxx - The allocated bandwidth for downloads, in bytes per second.',
  `MaxNumberOfConnections` int(11) NOT NULL DEFAULT '50' COMMENT 'per_user_max:xxx - The maximal authorized number of concurrent sessions.',
  `FilesQuota` bigint(20) NOT NULL DEFAULT '549755813888' COMMENT 'user_quota_files:xxx - The maximal number of files for this account.',
  `SizeQuota` bigint(20) NOT NULL DEFAULT '200' COMMENT 'user_quota_size:xxx - The maximal total size for this account, in bytes.',
  `AuthorizedClientIPs` text NOT NULL COMMENT 'Comma separated list of allowed IPs',
  `RefuzedClientIPs` text NOT NULL COMMENT 'Comma separated list of denied IPs',
  `Comments` varchar(255) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

ALTER TABLE `users` ADD PRIMARY KEY(`UserName`);
