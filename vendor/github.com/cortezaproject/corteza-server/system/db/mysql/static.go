// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package mysql

var Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1a\x00	\x0020180704080000.base.up.sqlUT\x05\x00\x01\x80Cm8-- all known organisations (crust instances) and our relation towards them\nCREATE TABLE organisations (\n  id               BIGINT UNSIGNED NOT NULL,\n  fqn              TEXT            NOT NULL, -- fully qualified name of the organisation\n  name             TEXT            NOT NULL, -- display name of the organisation\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  archived_at      DATETIME            NULL,\n  deleted_at       DATETIME            NULL, -- organisation soft delete\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE TABLE settings (\n  name  VARCHAR(200) NOT NULL   COMMENT 'Unique set of setting keys',\n  value TEXT                    COMMENT 'Setting value',\n\n  PRIMARY KEY (name)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\n-- Keeps all known users, home and external organisation\n--   changes are stored in audit log\nCREATE TABLE users (\n  id               BIGINT UNSIGNED NOT NULL,\n  email            TEXT            NOT NULL,\n  username         TEXT            NOT NULL,\n  password         TEXT            NOT NULL,\n  name             TEXT            NOT NULL,\n  handle           TEXT            NOT NULL,\n  meta             JSON            NOT NULL,\n  satosa_id        CHAR(36)            NULL,\n\n  rel_organisation BIGINT UNSIGNED NOT NULL,\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  suspended_at     DATETIME            NULL,\n  deleted_at       DATETIME            NULL, -- user soft delete\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE UNIQUE INDEX uid_satosa ON users (satosa_id);\n\n-- Keeps all known teams\nCREATE TABLE teams (\n  id               BIGINT UNSIGNED NOT NULL,\n  name             TEXT            NOT NULL, -- display name of the team\n  handle           TEXT            NOT NULL, -- team handle string\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  archived_at      DATETIME            NULL,\n  deleted_at       DATETIME            NULL, -- team soft delete\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\n-- Keeps team memberships\nCREATE TABLE team_members (\n  rel_team         BIGINT UNSIGNED NOT NULL REFERENCES organisation(id),\n  rel_user         BIGINT UNSIGNED NOT NULL,\n\n  PRIMARY KEY (rel_team, rel_user)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\nPK\x07\x08\xedzU\x8am	\x00\x00m	\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00.\x00	\x0020181124181811.rename_and_prefix_tables.up.sqlUT\x05\x00\x01\x80Cm8ALTER TABLE teams RENAME TO sys_team;\nALTER TABLE organisations RENAME TO sys_organisation;\nALTER TABLE team_members RENAME TO sys_team_member;\nALTER TABLE users RENAME TO sys_user;PK\x07\x08\xf2\xc4\x87\xe8\xb5\x00\x00\x00\xb5\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00-\x00	\x0020181125100429.add_user_kind_and_owner.up.sqlUT\x05\x00\x01\x80Cm8# add field to manage user type (bot support)\nALTER TABLE `sys_user` ADD `kind` VARCHAR(8) NOT NULL DEFAULT '' AFTER `handle`;\n\n# add field to manage \"ownership\" (get all bots created by user)\nALTER TABLE `sys_user` ADD `rel_user_id` BIGINT UNSIGNED NOT NULL AFTER `rel_organisation`, ADD INDEX (`rel_user_id`);\nPK\x07\x089\xa0\xdat8\x01\x00\x008\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00-\x00	\x0020181125153544.satosa_index_not_unique.up.sqlUT\x05\x00\x01\x80Cm8ALTER TABLE `sys_user` DROP INDEX `uid_satosa`, ADD INDEX `uid_satosa` (`satosa_id`) USING BTREE;PK\x07\x08\x0d\xf9\xd3ga\x00\x00\x00a\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00!\x00	\x0020181208140000.credentials.up.sqlUT\x05\x00\x01\x80Cm8-- Keeps all known users, home and external organisation\n--   changes are stored in audit log\nCREATE TABLE sys_credentials (\n  id               BIGINT UNSIGNED NOT NULL,\n  rel_owner        BIGINT UNSIGNED NOT NULL REFERENCES sys_users(id),\n  label            TEXT            NOT NULL COMMENT 'something we can differentiate credentials by',\n  kind             VARCHAR(128)    NOT NULL COMMENT 'hash, facebook, gplus, github, linkedin ...',\n  credentials      TEXT            NOT NULL COMMENT 'crypted/hashed passwords, secrets, social profile ID',\n  meta             JSON            NOT NULL,\n  expires_at       DATETIME            NULL,\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  deleted_at       DATETIME            NULL, -- user soft delete\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE INDEX idx_owner ON sys_credentials (rel_owner);\nPK\x07\x08f\x1f\x08\xd0\x9a\x03\x00\x00\x9a\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00)\x00	\x0020190103203201.users-password-null.up.sqlUT\x05\x00\x01\x80Cm8ALTER TABLE `sys_user` MODIFY `password` TEXT NULL;\nPK\x07\x080V\x13\x0f4\x00\x00\x004\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1b\x00	\x0020190116102104.rules.up.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE `sys_rules` (\n  `rel_team` BIGINT UNSIGNED NOT NULL,\n  `resource` VARCHAR(128) NOT NULL,\n  `operation` VARCHAR(128) NOT NULL,\n  `value` TINYINT(1) NOT NULL,\n\n  PRIMARY KEY (`rel_team`, `resource`, `operation`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\nPK\x07\x08\x05\x10[\x91\x05\x01\x00\x00\x05\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00)\x00	\x0020190221001051.rename-team-to-role.up.sqlUT\x05\x00\x01\x80Cm8ALTER TABLE sys_team RENAME TO sys_role;\nALTER TABLE sys_team_member RENAME TO sys_role_member;\n\nALTER TABLE `sys_role_member` CHANGE COLUMN `rel_team` `rel_role` BIGINT UNSIGNED NOT NULL;\nALTER TABLE `sys_rules` CHANGE COLUMN `rel_team` `rel_role` BIGINT UNSIGNED NOT NULL;\nPK\x07\x08s-\x98\xd0\x13\x01\x00\x00\x13\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00,\x00	\x0020190226160000.system_roles_and_rules.up.sqlUT\x05\x00\x01\x80Cm8REPLACE INTO `sys_role` (`id`, `name`, `handle`) VALUES\n  (1, 'Everyone', 'everyone'),\n  (2, 'Administrators', 'admins');\n\nPK\x07\x08\x06RHi{\x00\x00\x00{\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\"\x00	\x0020190306205033.applications.up.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE sys_application (\n  id               BIGINT UNSIGNED NOT NULL,\n  rel_owner        BIGINT UNSIGNED NOT NULL REFERENCES sys_users(id),\n  name             TEXT            NOT NULL COMMENT 'something we can differentiate application by',\n  enabled          BOOL            NOT NULL,\n\n  unify            JSON                NULL COMMENT 'unify specific settings',\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  deleted_at       DATETIME            NULL, -- user soft delete\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\n\nREPLACE INTO `sys_application` (`id`, `name`, `enabled`, `rel_owner`, `unify`) VALUES\n( 1, 'Crust Messaging', true, 0,\n  '{\"logo\": \"/applications/crust.jpg\", \"icon\": \"/applications/crust_favicon.png\", \"url\": \"/messaging/\", \"listed\": true}'\n),\n( 2, 'Crust CRM', true, 0,\n  '{\"logo\": \"/applications/crust.jpg\", \"icon\": \"/applications/crust_favicon.png\", \"url\": \"/crm/\", \"listed\": true}'\n),\n( 3, 'Crust Admin Area', true, 0,\n  '{\"logo\": \"/applications/crust.jpg\", \"icon\": \"/applications/crust_favicon.png\", \"url\": \"/admin/\", \"listed\": true}'\n),\n( 4, 'Corteza Jitsi Bridge', true, 0,\n  '{\"logo\": \"/applications/jitsi.png\", \"icon\": \"/applications/jitsi_icon.png\", \"url\": \"/bridge/jitsi/\", \"listed\": true}'\n),\n( 5, 'Google Maps', true, 0,\n  '{\"logo\": \"/applications/google_maps.png\", \"icon\": \"/applications/google_maps_icon.png\", \"url\": \"/bridge/google-maps/\", \"listed\": true}'\n);\n\nPK\x07\x08Oi\xd5\xd3\xc6\x05\x00\x00\xc6\x05\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1e\x00	\x0020190326122000.settings.up.sqlUT\x05\x00\x01\x80Cm8DROP TABLE IF EXISTS `settings`;\n\nCREATE TABLE IF NOT EXISTS `sys_settings` (\n  rel_owner        BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT 'Value owner, 0 for global settings',\n  name             VARCHAR(200)    NOT NULL               COMMENT 'Unique set of setting keys',\n  value            JSON                                   COMMENT 'Setting value',\n\n  updated_at       DATETIME        NOT NULL DEFAULT NOW() COMMENT 'When was the value updated',\n  updated_by       BIGINT UNSIGNED NOT NULL DEFAULT 0     COMMENT 'Who created/updated the value',\n\n  PRIMARY KEY (name, rel_owner)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\nPK\x07\x08`\xcb\x1b\x81t\x02\x00\x00t\x02\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00#\x00	\x0020190403113201.users-cleanup.up.sqlUT\x05\x00\x01\x80Cm8ALTER TABLE `sys_user` DROP `password`;\nALTER TABLE `sys_user` DROP `satosa_id`;\nALTER TABLE `sys_credentials` ADD `last_used_at` DATETIME NULL;\nPK\x07\x088\x92\x0fs\x91\x00\x00\x00\x91\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00#\x00	\x0020190405090000.internal-auth.up.sqlUT\x05\x00\x01\x80Cm8ALTER TABLE `sys_user` ADD `email_confirmed` BOOLEAN NOT NULL DEFAULT FALSE;\nPK\x07\x08\x8fQs\x8cM\x00\x00\x00M\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00!\x00	\x0020190506090000.compose-app.up.sqlUT\x05\x00\x01\x80Cm8UPDATE `sys_application`\n   SET `name`  = 'Crust Compose',\n       `unify` = '{\"logo\": \"/applications/default_logo.jpg\", \"icon\": \"/applications/default_icon.png\", \"url\": \"/compose/\", \"listed\": true}'\n WHERE id = 2;\nPK\x07\x089\x0b\xb8\xf9\xd6\x00\x00\x00\xd6\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00!\x00	\x0020190506090000.permissions.up.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE IF NOT EXISTS sys_permission_rules (\n  rel_role   BIGINT UNSIGNED NOT NULL,\n  resource   VARCHAR(128)    NOT NULL,\n  operation  VARCHAR(128)    NOT NULL,\n  access     TINYINT(1)      NOT NULL,\n\n  PRIMARY KEY (rel_role, resource, operation)\n) ENGINE=InnoDB;\n\nCREATE TABLE IF NOT EXISTS messaging_permission_rules (\n  rel_role   BIGINT UNSIGNED NOT NULL,\n  resource   VARCHAR(128)    NOT NULL,\n  operation  VARCHAR(128)    NOT NULL,\n  access     TINYINT(1)      NOT NULL,\n\n  PRIMARY KEY (rel_role, resource, operation)\n) ENGINE=InnoDB;\n\nCREATE TABLE IF NOT EXISTS compose_permission_rules (\n  rel_role   BIGINT UNSIGNED NOT NULL,\n  resource   VARCHAR(128)    NOT NULL,\n  operation  VARCHAR(128)    NOT NULL,\n  access     TINYINT(1)      NOT NULL,\n\n  PRIMARY KEY (rel_role, resource, operation)\n) ENGINE=InnoDB;\n\nREPLACE sys_permission_rules\n    (rel_role, resource, operation, access)\n    SELECT rel_role, resource, operation, `value` - 1 FROM sys_rules WHERE resource LIKE 'system%';\n\nREPLACE compose_permission_rules\n    (rel_role, resource, operation, access)\n    SELECT rel_role, resource, operation, `value` - 1 FROM sys_rules WHERE resource LIKE 'compose%';\n\nREPLACE messaging_permission_rules\n    (rel_role, resource, operation, access)\n    SELECT rel_role, resource, operation, `value` - 1 FROM sys_rules WHERE resource LIKE 'messaging%';\n\nDROP TABLE sys_rules;\nPK\x07\x08\x08\xd4\xe0+e\x05\x00\x00e\x05\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00*\x00	\x0020190826085348.migrate-gplus-google.up.sqlUT\x05\x00\x01\x80Cm8/* migrates existing credentials */\nUPDATE sys_credentials SET kind = 'google' WHERE kind = 'gplus';\n\n/* migrates existing settings. */\nUPDATE sys_settings SET name = REPLACE(name, '.gplus.', '.google.') WHERE name LIKE 'auth.external.providers.gplus.%';\nPK\x07\x08<\xac\xedE\xff\x00\x00\x00\xff\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00 \x00	\x0020190902080000.automation.up.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE IF NOT EXISTS sys_automation_script (\n    `id`            BIGINT(20)  UNSIGNED NOT NULL,\n    `rel_namespace` BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0         COMMENT 'For compatibility only, not used',\n    `name`          VARCHAR(64)          NOT NULL DEFAULT 'unnamed' COMMENT 'The name of the script',\n    `source`        TEXT                 NOT NULL                   COMMENT 'Source code for the script',\n    `source_ref`    VARCHAR(200)         NOT NULL                   COMMENT 'Where is the script located (if remote)',\n    `async`         BOOLEAN              NOT NULL DEFAULT FALSE     COMMENT 'Do we run this script asynchronously?',\n    `rel_runner`    BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0         COMMENT 'Who is running the script? 0 for invoker',\n    `run_in_ua`     BOOLEAN              NOT NULL DEFAULT FALSE     COMMENT 'Run this script inside user-agent environment',\n    `timeout`       INT         UNSIGNED NOT NULL DEFAULT 0         COMMENT 'Any explicit timeout set for this script (milliseconds)?',\n    `critical`      BOOLEAN              NOT NULL DEFAULT TRUE      COMMENT 'Is it critical that this script is executed successfully',\n    `enabled`       BOOLEAN              NOT NULL DEFAULT TRUE      COMMENT 'Is this script enabled?',\n\n    `created_by`    BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `created_at`    DATETIME             NOT NULL DEFAULT CURRENT_TIMESTAMP,\n    `updated_by`    BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `updated_at`    DATETIME                 NULL DEFAULT NULL,\n    `deleted_by`    BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `deleted_at`    DATETIME                 NULL DEFAULT NULL,\n\n    PRIMARY KEY (`id`)\n\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE TABLE IF NOT EXISTS sys_automation_trigger (\n    `id`         BIGINT(20)  UNSIGNED NOT NULL,\n    `rel_script` BIGINT(20)  UNSIGNED NOT NULL              COMMENT 'Script that is triggered',\n\n    `resource`   VARCHAR(128)         NOT NULL              COMMENT 'Resource triggering the event',\n    `event`      VARCHAR(128)         NOT NULL              COMMENT 'Event triggered',\n    `event_condition`\n                 TEXT                 NOT NULL              COMMENT 'Trigger condition',\n    `enabled`    BOOLEAN              NOT NULL DEFAULT TRUE COMMENT 'Trigger enabled?',\n\n    `weight`     INT                  NOT NULL DEFAULT 0,\n\n    `created_by` BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `created_at` DATETIME             NOT NULL DEFAULT CURRENT_TIMESTAMP,\n    `updated_by` BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `updated_at` DATETIME                 NULL DEFAULT NULL,\n    `deleted_by` BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `deleted_at` DATETIME                 NULL DEFAULT NULL,\n\n    CONSTRAINT `fk_sys_automation_script` FOREIGN KEY (`rel_script`) REFERENCES `sys_automation_script` (`id`),\n\n    PRIMARY KEY (`id`)\n\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\nPK\x07\x08\xac\xbb\x1b\x07i\x0b\x00\x00i\x0b\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1f\x00	\x0020190924093443.reminders.up.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE IF NOT EXISTS sys_reminder (\n    `id`           BIGINT(20)   UNSIGNED NOT NULL,\n    `resource`     VARCHAR(128)          NOT NULL                           COMMENT 'Resource, that this reminder is bound to',\n    `payload`      JSON                  NOT NULL                           COMMENT 'Payload for this reminder',\n    `snooze_count` INT                   NOT NULL DEFAULT 0                 COMMENT 'Number of times this reminder was snoozed',\n\n    `assigned_to`  BIGINT(20)   UNSIGNED NOT NULL DEFAULT 0                 COMMENT 'Assignee for this reminder',\n    `assigned_by`  BIGINT(20)   UNSIGNED NOT NULL DEFAULT 0                 COMMENT 'User that assigned this reminder',\n    `assigned_at`  DATETIME              NOT NULL                           COMMENT 'When the reminder was assigned',\n\n    `dismissed_by` BIGINT(20)   UNSIGNED NOT NULL DEFAULT 0                 COMMENT 'User that dismissed this reminder',\n    `dismissed_at` DATETIME                  NULL DEFAULT NULL              COMMENT 'Time the reminder was dismissed',\n\n    `remind_at`    DATETIME                  NULL DEFAULT NULL              COMMENT 'Time the user should be reminded',\n\n    `created_by`   BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `created_at`   DATETIME             NOT NULL DEFAULT CURRENT_TIMESTAMP,\n    `updated_by`   BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `updated_at`   DATETIME                 NULL DEFAULT NULL,\n    `deleted_by`   BIGINT(20)  UNSIGNED NOT NULL DEFAULT 0,\n    `deleted_at`   DATETIME                 NULL DEFAULT NULL,\n\n    PRIMARY KEY (`id`)\n\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\nPK\x07\x08\n\x10\"\x05X\x06\x00\x00X\x06\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00&\x00	\x0020191023213030.settings-cleanup.up.sqlUT\x05\x00\x01\x80Cm8UPDATE `sys_settings` SET `name` = 'general.mail.logo'      WHERE `rel_owner` = 0 AND `name` = 'system.defaultLogo';\nUPDATE `sys_settings` SET `name` = 'general.mail.header.en' WHERE `rel_owner` = 0 AND `name` = 'system.mail.header.en';\nUPDATE `sys_settings` SET `name` = 'general.mail.footer.en' WHERE `rel_owner` = 0 AND `name` = 'system.mail.footer.en';\nPK\x07\x08\x98\xd0\xdcje\x01\x00\x00e\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00 \x00	\x0020200419125927.attachment.up.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE IF NOT EXISTS sys_attachment (\n  id               BIGINT UNSIGNED NOT NULL,\n  rel_owner        BIGINT UNSIGNED NOT NULL,\n\n  kind             VARCHAR(32) NOT NULL,\n\n  url              VARCHAR(512),\n  preview_url      VARCHAR(512),\n\n  size             INT    UNSIGNED,\n  mimetype         VARCHAR(255),\n  name             TEXT,\n\n  meta             JSON,\n\n  created_at       DATETIME        NOT NULL DEFAULT NOW(),\n  updated_at       DATETIME            NULL,\n  deleted_at       DATETIME            NULL,\n\n  PRIMARY KEY (id)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nPK\x07\x08\xca\xba\xa1l=\x02\x00\x00=\x02\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x1f\x00	\x0020200508070000.actionlog.up.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE IF NOT EXISTS sys_actionlog (\n  ts               DATETIME        NOT NULL DEFAULT NOW(),\n  actor_ip_addr    VARCHAR(15)     NOT NULL,\n  actor_id         BIGINT          UNSIGNED,\n  request_origin   VARCHAR(32)     NOT NULL,\n  request_id       VARCHAR(64)     NOT NULL,\n  resource         VARCHAR(128)    NOT NULL,\n  `action`         VARCHAR(64)     NOT NULL,\n  `error`          VARCHAR(64)     NOT NULL,\n  severity         SMALLINT        NOT NULL,\n  description      TEXT,\n  meta             JSON\n\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nCREATE INDEX ts             ON sys_actionlog (ts DESC);\nCREATE INDEX request_origin ON sys_actionlog (request_origin);\nCREATE INDEX actor_id       ON sys_actionlog (actor_id);\nCREATE INDEX resource       ON sys_actionlog (resource);\nCREATE INDEX `action`       ON sys_actionlog (`action`);\nPK\x07\x08>\xed!\xdbI\x03\x00\x00I\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00migrations.sqlUT\x05\x00\x01\x80Cm8CREATE TABLE IF NOT EXISTS `migrations` (\n `project` varchar(16) NOT NULL COMMENT 'sam, crm, ...',\n `filename` varchar(255) NOT NULL COMMENT 'yyyymmddHHMMSS.sql',\n `statement_index` int(11) NOT NULL COMMENT 'Statement number from SQL file',\n `status` TEXT NOT NULL COMMENT 'ok or full error message',\n PRIMARY KEY (`project`,`filename`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8;\n\nPK\x07\x08\x0d\xa5T2x\x01\x00\x00x\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x00\x00!(\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x06\x00	\x00new.shUT\x05\x00\x01\x80Cm8#!/bin/bash\ntouch $(date +%Y%m%d%H%M%S).up.sqlPK\x07\x08s\xd4N*.\x00\x00\x00.\x00\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\xedzU\x8am	\x00\x00m	\x00\x00\x1a\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x0020180704080000.base.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\xf2\xc4\x87\xe8\xb5\x00\x00\x00\xb5\x00\x00\x00.\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xbe	\x00\x0020181124181811.rename_and_prefix_tables.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(9\xa0\xdat8\x01\x00\x008\x01\x00\x00-\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xd8\n\x00\x0020181125100429.add_user_kind_and_owner.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x0d\xf9\xd3ga\x00\x00\x00a\x00\x00\x00-\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81t\x0c\x00\x0020181125153544.satosa_index_not_unique.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(f\x1f\x08\xd0\x9a\x03\x00\x00\x9a\x03\x00\x00!\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x819\x0d\x00\x0020181208140000.credentials.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(0V\x13\x0f4\x00\x00\x004\x00\x00\x00)\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81+\x11\x00\x0020190103203201.users-password-null.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x05\x10[\x91\x05\x01\x00\x00\x05\x01\x00\x00\x1b\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xbf\x11\x00\x0020190116102104.rules.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(s-\x98\xd0\x13\x01\x00\x00\x13\x01\x00\x00)\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x16\x13\x00\x0020190221001051.rename-team-to-role.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x06RHi{\x00\x00\x00{\x00\x00\x00,\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x89\x14\x00\x0020190226160000.system_roles_and_rules.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(Oi\xd5\xd3\xc6\x05\x00\x00\xc6\x05\x00\x00\"\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81g\x15\x00\x0020190306205033.applications.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(`\xcb\x1b\x81t\x02\x00\x00t\x02\x00\x00\x1e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x86\x1b\x00\x0020190326122000.settings.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(8\x92\x0fs\x91\x00\x00\x00\x91\x00\x00\x00#\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81O\x1e\x00\x0020190403113201.users-cleanup.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x8fQs\x8cM\x00\x00\x00M\x00\x00\x00#\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81:\x1f\x00\x0020190405090000.internal-auth.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(9\x0b\xb8\xf9\xd6\x00\x00\x00\xd6\x00\x00\x00!\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xe1\x1f\x00\x0020190506090000.compose-app.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x08\xd4\xe0+e\x05\x00\x00e\x05\x00\x00!\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x0f!\x00\x0020190506090000.permissions.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(<\xac\xedE\xff\x00\x00\x00\xff\x00\x00\x00*\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xcc&\x00\x0020190826085348.migrate-gplus-google.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\xac\xbb\x1b\x07i\x0b\x00\x00i\x0b\x00\x00 \x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81,(\x00\x0020190902080000.automation.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\n\x10\"\x05X\x06\x00\x00X\x06\x00\x00\x1f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xec3\x00\x0020190924093443.reminders.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x98\xd0\xdcje\x01\x00\x00e\x01\x00\x00&\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x9a:\x00\x0020191023213030.settings-cleanup.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\xca\xba\xa1l=\x02\x00\x00=\x02\x00\x00 \x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\\<\x00\x0020200419125927.attachment.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(>\xed!\xdbI\x03\x00\x00I\x03\x00\x00\x1f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xf0>\x00\x0020200508070000.actionlog.up.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(\x0d\xa5T2x\x01\x00\x00x\x01\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x8fB\x00\x00migrations.sqlUT\x05\x00\x01\x80Cm8PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x00\x00!(s\xd4N*.\x00\x00\x00.\x00\x00\x00\x06\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xed\x81LD\x00\x00new.shUT\x05\x00\x01\x80Cm8PK\x05\x06\x00\x00\x00\x00\x17\x00\x17\x00\xf7\x07\x00\x00\xb7D\x00\x00\x00\x00"
