CREATE TABLE IF NOT EXISTS "gifts" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "type" VARCHAR(50),
    "cpf" VARCHAR(14) NOT NULL,
    "amount" DECIMAL(10, 2),
    "pix_key_type" VARCHAR(50),
    "pix_key" VARCHAR(255),
    "message" TEXT,
    "status" VARCHAR(50) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
