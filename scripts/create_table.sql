-- Table: public.list_node

-- DROP TABLE IF EXISTS public.list_node;

CREATE TABLE IF NOT EXISTS public.list_node
(
    id integer NOT NULL DEFAULT nextval('list_node_id_seq'::regclass),
    list_1 bigint[] NOT NULL,
    list_2 bigint[] NOT NULL,
    merged_list bigint[],
    CONSTRAINT list_node_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.list_node
    OWNER to postgres;