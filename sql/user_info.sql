-- 1. 创建用户表
CREATE TABLE "user_info" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR(50) NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "status" SMALLINT NOT NULL DEFAULT 1,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT uk_username UNIQUE ("username")
);

-- 2. 为各字段添加注释
COMMENT ON COLUMN "user_info"."id" IS '用户唯一标识（自增主键）';
COMMENT ON COLUMN "user_info"."username" IS '用户名（登录用，唯一）';
COMMENT ON COLUMN "user_info"."password" IS '加密后的密码（BCrypt/SHA256等算法哈希）';
COMMENT ON COLUMN "user_info"."status" IS '账号状态：0-禁用，1-正常';
COMMENT ON COLUMN "user_info"."created_at" IS '创建时间（带时区）';
COMMENT ON COLUMN "user_info"."updated_at" IS '更新时间（带时区）';
COMMENT ON TABLE "user_info" IS '用户登录核心表';

-- 3. 创建更新updated_at的函数和触发器
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW."updated_at" = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_user_updated_at
BEFORE UPDATE ON "user_info"
FOR EACH ROW
EXECUTE FUNCTION update_modified_column();