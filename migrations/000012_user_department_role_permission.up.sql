CREATE TABLE user_department_role_permission (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    tenant_id     BIGINT NOT NULL,
    user_id       BIGINT NOT NULL,
    dept_id       BIGINT NOT NULL,
    role_id       BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,

    created_by    BIGINT,
    updated_by    BIGINT,
    deleted_by    BIGINT,

    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    is_deleted    BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at    TIMESTAMPTZ,

    CONSTRAINT fk_udrp_tenant
        FOREIGN KEY (tenant_id)
        REFERENCES tenant(id),

    CONSTRAINT fk_udrp_user
        FOREIGN KEY (user_id)
        REFERENCES user_table(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_udrp_department
        FOREIGN KEY (dept_id)
        REFERENCES department_master(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_udrp_role
        FOREIGN KEY (role_id)
        REFERENCES role_master(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_udrp_permission
        FOREIGN KEY (permission_id)
        REFERENCES permission(id)
        ON DELETE CASCADE,

    CONSTRAINT uk_udrp_user_dept_role
        UNIQUE (
            tenant_id,
            user_id,
            dept_id,
            role_id
        )
);