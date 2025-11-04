-- 用户表：适配 PostgreSQL 的语法
CREATE TABLE "user" (
  "id" BIGSERIAL PRIMARY KEY COMMENT '用户唯一标识（自增主键）',
  "username" VARCHAR(50) NOT NULL COMMENT '用户名（登录用，唯一）',
  "password" VARCHAR(255) NOT NULL COMMENT '加密后的密码（BCrypt/SHA256等算法哈希）',
  "status" SMALLINT NOT NULL DEFAULT 1 COMMENT '账号状态：0-禁用，1-正常',
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间（带时区）',
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间（带时区）',
  CONSTRAINT uk_username UNIQUE ("username") COMMENT '用户名唯一索引'
);

-- 为 updated_at 添加自动更新触发器（PostgreSQL 无 ON UPDATE 语法，需手动实现）
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW."updated_at" = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_user_updated_at
BEFORE UPDATE ON "user"
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();