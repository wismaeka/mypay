CREATE TABLE IF NOT EXISTS transfers
(
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sender_number   TEXT                           NOT NULL,
    sender_name     TEXT                           NOT NULL,
    sender_bank     TEXT                           NOT NULL,
    receiver_number TEXT                           NOT NULL,
    receiver_name   TEXT                           NOT NULL,
    receiver_bank   TEXT                           NOT NULL,
    amount          TEXT                           NOT NULL,
    status          TEXT                           NOT NULL,
    currency        TEXT                           NOT NULL,
    description     TEXT                           NOT NULL,
    created         TIMESTAMPTZ      DEFAULT NOW() NOT NULL,
    updated         TIMESTAMPTZ      DEFAULT NOW() NOT NULL
);
