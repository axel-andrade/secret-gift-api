CREATE TABLE IF NOT EXISTS "authorized_gifts" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "gift_id" UUID NOT NULL REFERENCES "gifts" ("id") ON DELETE CASCADE,
    "expiration_date" TIMESTAMP,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
